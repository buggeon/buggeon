// Buggeon - SelfHosted service for bug and task tracking
// Copyright (C) 2026 DEVE corp.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"buggeon/config"
	_ "buggeon/docs"
	"buggeon/internal/cache"
	"buggeon/internal/db"
	"buggeon/internal/graph"
	"buggeon/internal/handlers"
	"buggeon/internal/hub"
	"buggeon/internal/middleware"
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	s3storage "buggeon/internal/s3Storage"
	"buggeon/internal/security"
	"buggeon/internal/services"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createAdmin(tokenService *services.TokenService) error {

	admin_password, err := security.HashPassword(os.Getenv("ADMIN_PASSWORD"))

	if err != nil {
		return err
	}

	admin_login := "@" + os.Getenv("ADMIN_LOGIN")
	admin_name := os.Getenv("ADMIN_NAME")
	admin_email := os.Getenv("ADMIN_EMAIL")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	users := db.GetCollection("buggeon", "users")

	count, err := users.CountDocuments(ctx, bson.M{"role": "admin"})

	if err != nil {
		return err
	}

	if count == 0 {

		admin := models.User{
			ID:        primitive.NewObjectID(),
			Name:      admin_name,
			Login:     admin_login,
			Password:  admin_password,
			AvatarKey: "",
			Email:     admin_email,
			Role:      "admin",
			CreatedAt: time.Now(),
		}

		refreshToken, err := tokenService.GenerateRefreshToken(admin.ID.Hex())

		if err != nil {
			return err
		}

		admin.RefreshTokens = []string{refreshToken}

		_, err = users.InsertOne(context.TODO(), admin)

		return err
	}

	return nil

}

// @title           Buggeon API
// @version         0.1.0
// @description     Self-hosted bug and task tracking service.
// @termsOfService  https://buggeon.com/terms

// @contact.name   Buggeon Support
// @contact.url    https://buggeon.com/support
// @contact.email  support@buggeon.com

// @license.name  AGPL-3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.html

// @host      localhost:9090
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введи "Bearer <token>"

func main() {

	godotenv.Load()

	mongoHost := os.Getenv("DB_HOST")
	mongoPort := os.Getenv("DB_PORT")
	mongoUser := os.Getenv("DB_USER")
	mongoPassword := os.Getenv("DB_PASSWORD")

	if err := db.Connect(mongoHost, mongoPort, mongoUser, mongoPassword); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	s3Storage := s3storage.NewS3Storage()

	userCache := cache.NewUserCache(5 * time.Minute)
	permissionsCache := cache.NewPermissionCache(5 * time.Minute)

	wsHub := hub.NewHub()

	go wsHub.Run()

	cardHub := hub.NewCardHub(wsHub)
	userHub := hub.NewUserHub(wsHub)

	userRepo := repositories.NewUserRepo()
	projectRepo := repositories.NewProjectRepo()
	memberRepo := repositories.NewMemberRepo()
	boardRepo := repositories.NewBoardRepo()
	cardRepo := repositories.NewCardRepo()
	messageRepo := repositories.NewMessageRepo()
	schemaRepo := repositories.NewSchemaRepo()
	notificationRepo := repositories.NewNotificationRepo()

	notificationService := services.NewNotificationService(notificationRepo, userHub)
	tokenService := services.NewTokenService(config.LoadConfig())
	boardService := services.NewBoardService(boardRepo, projectRepo, cardRepo, messageRepo)
	cardService := services.NewCardService(cardRepo, boardRepo, messageRepo, notificationService, memberRepo)
	memberService := services.NewMemberService(memberRepo, projectRepo, *notificationService)
	schemaService := services.NewSchemaService(projectRepo, schemaRepo, s3Storage)
	projectService := services.NewProjectService(projectRepo, memberService, memberRepo, boardRepo, cardRepo, messageRepo, s3Storage)
	userService := services.NewUserService(userRepo, tokenService, s3Storage, userCache)
	messageService := services.NewMessageService(messageRepo, cardRepo)

	projectHandler := handlers.NewProjectHandler(projectService, schemaService, cardService, boardService, memberService, messageService)
	userHandler := handlers.NewUserHandler(userService)
	cardHandler := handlers.NewCardHandler(cardService, boardService)
	boardHandler := handlers.NewBoardHandler(cardService, boardService)
	memberHandler := handlers.NewMemberHandler(memberService)
	cardWSHandler := handlers.NewCardWSHandler(messageService, tokenService, userService, cardHub)
	userWSHandler := handlers.NewUserWSHandler(tokenService, userService, userHub)

	authMiddleware := middleware.NewAuthMiddleware(tokenService)
	permissionsMiddleware := middleware.NewPermissionsMiddleware(permissionsCache)
	loaderMiddleware := middleware.NewLoaderMiddleware(userRepo, cardRepo, boardRepo, messageRepo, memberRepo)

	router.Use(loaderMiddleware.SetLoaderMiddleware())

	if err := createAdmin(tokenService); err != nil {
		fmt.Println("Failed to create admin")
	}

	graphqlResolver := &graph.Resolver{
		ProjectService: projectService,
		BoardService:   boardService,
		CardService:    cardService,
		UserService:    userService,
		MemberService:  memberService,
		MessageService: messageService,
	}

	gqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: graphqlResolver},
		),
	)

	setupRoutes(
		router,
		projectHandler,
		memberHandler,
		cardHandler,
		boardHandler,
		userHandler,
		cardWSHandler,
		userWSHandler,
		authMiddleware,
		permissionsMiddleware,
		gqlHandler,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9187"
	}
	log.Printf("Server starting on :%s", port)
	router.Run(fmt.Sprintf(":%s", port))
}

func setupRoutes(
	router *gin.Engine,
	projectHandler *handlers.ProjectHandler,
	memberHandler *handlers.MemberHandler,
	cardHandler *handlers.CardHandler,
	boardHandler *handlers.BoardHandler,
	userHandler *handlers.UserHandler,
	cardWSHandler *handlers.CardWSHandler,
	userWSHandler *handlers.UserWSHandler,
	authMiddleware *middleware.AuthMiddleware,
	permissionsMiddleware *middleware.PermissionsMiddleware,
	gqlHandler *handler.Server,
) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/test", handlers.Test)
	router.GET("/ws/card/:cardId", cardWSHandler.ServeWS)
	router.GET("/ws/user", userWSHandler.ServeWS)

	router.Static("/assets", "static/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("static/index.html")
	})

	router.Any("/files/*path", handlers.S3ProxyHandler(os.Getenv("S3_ENDPOINT")))

	api := router.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Registration)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refreshtoken", userHandler.RefreshAccessToken)
		}

		protected := api.Group("")
		protected.Use(authMiddleware.AuthRequired())
		{
			users := protected.Group("/users")
			{
				users.GET("/:user_id", userHandler.GetUser)
				users.GET("/", userHandler.GetUsers)
				users.PATCH("/:user_id/avatar", userHandler.SetAvatar)
			}

			projects := protected.Group("/projects")
			{
				projects.GET("", projectHandler.GetProjects)
				projects.GET("/:project_id", projectHandler.GetProject)
				projects.POST("", projectHandler.CreateProject)
				projects.DELETE("/:project_id", projectHandler.DeleteProject)
				projects.PATCH("/:project_id/logo", projectHandler.SetProjectLogo)

				members := projects.Group("/:project_id/members")
				{
					members.GET("", memberHandler.GetMembers)
					members.GET("/:member_id", memberHandler.GetMember)
					members.POST("", memberHandler.CreateMember)
					members.DELETE("/:member_id", memberHandler.DeleteMember)
				}

				boards := projects.Group("/:project_id/boards")
				{
					boards.GET("", boardHandler.GetBoards)
					boards.GET("/:board_id", boardHandler.GetBoard)
					boards.POST("", boardHandler.CreateBoard)
					boards.DELETE("/:board_id", boardHandler.DeleteBoard)

					cards := boards.Group("/:board_id/cards")
					{
						cards.GET("", cardHandler.GetCards)
						cards.GET("/:card_id", cardHandler.GetCard)
						cards.POST("", cardHandler.CreateCard)
						cards.DELETE("/:card_id", cardHandler.DeleteCard)
						cards.PUT("/:card_id/updatelocation", cardHandler.UpdateCardLocation)
					}
				}
			}

			queryGroup := protected.Group("/query")
			queryGroup.POST("", gin.WrapH(gqlHandler))
		}
	}

}
