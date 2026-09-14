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
	"buggeon/graph"
	"buggeon/internal/db"
	"buggeon/internal/handlers"
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
	"github.com/gin-contrib/cors"
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
			AvatarUrl: "",
			Email:     admin_email,
			Role:      "admin",
			CreatedAt: time.Now(),
		}

		refreshToken, err := tokenService.GenerateRefreshToken(&admin)

		if err != nil {
			return err
		}

		admin.RefreshTokens = []string{refreshToken}

		_, err = users.InsertOne(context.TODO(), admin)

		return err
	}

	return nil

}

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

	corsConfig := cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"OPTIONS",
			"HEAD",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	router := gin.Default()
	router.Use(corsConfig)

	s3Storage := s3storage.NewS3Storage()

	userRepo := repositories.NewUserRepo()
	projectRepo := repositories.NewProjectRepo()
	memberRepo := repositories.NewMemberRepo()
	boardRepo := repositories.NewBoardRepo()
	cardRepo := repositories.NewCardRepo()
	messageRepo := repositories.NewMessageRepo()
	schemaRepo := repositories.NewSchemaRepo()

	boardService := services.NewBoardService(boardRepo, projectRepo, cardRepo, messageRepo)
	cardService := services.NewCardService(cardRepo, boardRepo, messageRepo)
	memberService := services.NewMemberService(memberRepo, projectRepo)
	schemaService := services.NewSchemaService(projectRepo, schemaRepo, s3Storage)
	projectService := services.NewProjectService(projectRepo, memberService, memberRepo, boardRepo, cardRepo, messageRepo, s3Storage)
	tokenService := services.NewTokenService(config.LoadConfig())
	userService := services.NewUserService(userRepo, tokenService)
	systemService := services.NewSystemService(userRepo)
	messageService := services.NewMessageService(messageRepo, cardRepo)

	projectHandler := handlers.NewProjectHandler(projectService, schemaService, cardService, boardService, memberService, messageService)
	userHandler := handlers.NewUserHandler(userService)
	systemHandler := handlers.NewSystemHandler(systemService)
	chatHandler := handlers.NewChatHandler(messageService, tokenService)

	authMiddleware := middleware.NewAuthMiddleware(tokenService)

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

	setupRoutes(router, projectHandler, userHandler, systemHandler, chatHandler, authMiddleware, gqlHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	router.Run(":" + port)
}

func setupRoutes(
	router *gin.Engine,
	projectHandler *handlers.ProjectHandler,
	userHandler *handlers.UserHandler,
	systemHandler *handlers.SystemHandler,
	chatHandler *handlers.ChatHandler,
	authMiddleware *middleware.AuthMiddleware,
	gqlHandler *handler.Server,
) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/auth/register", userHandler.Registration)
	router.POST("/auth/login", userHandler.Login)
	router.POST("/auth/refreshtoken", userHandler.RefreshAccessToken)
	router.GET("/test", handlers.Test)
	router.GET("/ws/:cardId", chatHandler.ServeWS)

	api := router.Group("/api")
	api.Use(authMiddleware.AuthRequired())
	{
		projects := api.Group("/projects")
		{
			projects.GET("", projectHandler.GetProjects)
			projects.GET("/:project_id", projectHandler.GetProject)
			projects.POST("", projectHandler.CreateProject)
			projects.DELETE("/:project_id", projectHandler.DeleteProject)
			projects.PATCH("/:project_id/logo", projectHandler.SetProjectLogo)

			schemas := projects.Group("/:project_id/schemas")
			{
				schemas.POST("", projectHandler.CreateProjectSchema)
				schemas.GET("", projectHandler.GetProjectSchemas)
			}

			members := projects.Group("/:project_id/members")
			{
				members.GET("", projectHandler.GetMembers)
				members.GET("/:member_id", projectHandler.GetMember)
				members.POST("", projectHandler.CreateMember)
				members.DELETE("/:member_id", projectHandler.DeleteMember)
			}

			boards := projects.Group("/:project_id/boards")
			{
				boards.GET("", projectHandler.GetBoards)
				boards.GET("/:board_id", projectHandler.GetBoard)
				boards.POST("", projectHandler.CreateBoard)
				boards.DELETE("/:board_id", projectHandler.DeleteBoard)

				cards := boards.Group("/:board_id/cards")
				{
					cards.GET("", projectHandler.GetCards)
					cards.GET("/:card_id", projectHandler.GetCard)
					cards.POST("", projectHandler.CreateCard)
					cards.DELETE("/:card_id", projectHandler.DeleteCard)
					cards.PUT("/:card_id/updatelocation", projectHandler.UpdateCardLocation)

					messages := cards.Group("/:card_id/messages")
					{
						messages.POST("", projectHandler.NewMessage)
					}
				}
			}
		}

		queryGroup := api.Group("/query")
		queryGroup.POST("", gin.WrapH(gqlHandler))

		api.GET("/users", systemHandler.GetAllUsers)
	}
}
