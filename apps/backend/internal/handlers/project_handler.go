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

package handlers

import (
	"buggeon/internal/dto"
	_ "buggeon/internal/models"
	"buggeon/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *services.ProjectService
	schemaService  *services.SchemaService
	boardService   *services.BoardService
	cardService    *services.CardService
	memberService  *services.MemberService
	messageService *services.MessageService
}

func NewProjectHandler(
	projectService *services.ProjectService,
	schemaService *services.SchemaService,
	cardService *services.CardService,
	boardServices *services.BoardService,
	memberService *services.MemberService,
	messageService *services.MessageService,
) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		boardService:   boardServices,
		schemaService:  schemaService,
		cardService:    cardService,
		memberService:  memberService,
		messageService: messageService,
	}
}

// CreateProject godoc
// @Summary      Create a project
// @Description  Create a new project with the current user as lead
// @Tags         projects
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      dto.CreateProjectDto  true  "Project data"
// @Success      201    {object}  models.Project  "Created project"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      401    {object}  map[string]string  "Unauthorized"
// @Router       /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {

	// leadID := c.Request.FormValue("leadId")
	// name := c.Request.FormValue("name")
	// description := c.Request.FormValue("description")
	// //members := c.Request.FormValue("members")
	// logo, _ := c.FormFile("logo")

	// err := h.projectService.CreateProject(&dto.CreateProjectDto{
	// 	LeadID:      leadID,
	// 	Name:        name,
	// 	Description: description,
	// 	Logo:        logo,
	// 	Progress:    0,
	// 	//Members: string(membersJSON),
	// })

	// if err != nil {
	// 	c.Status(403)
	// 	c.Abort()
	// 	return
	// }

	// c.Status(201)

}

// SetProjectLogo godoc
// @Summary      Set project logo
// @Description  Upload a new logo file for the project
// @Tags         projects
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        logo        formData  file    true  "New logo file"
// @Success      200         {object}  map[string]string  "ok"
// @Failure      403         {object}  map[string]string  "New logo file was not provided"
// @Failure      500         {object}  map[string]string  "Failed to upload new logo file"
// @Router       /projects/{project_id}/logo [patch]
func (h *ProjectHandler) SetProjectLogo(c *gin.Context) {

	file, err := c.FormFile("logo")
	projectID := c.Param("project_id")

	if err != nil {
		c.JSON(403, "New logo file was not provide")
		return
	}

	err = h.projectService.SetProjectLogo(c, projectID, file)

	if err != nil {
		c.JSON(500, "Failed to upload new logo file")
		return
	}

}

// DeleteProject godoc
// @Summary      Delete a project
// @Description  Delete a project by its ID
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {

	projectID := c.Param("project_id")

	_, err := h.projectService.DeleteProject(c, projectID)

	if err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	c.Status(200)
}

// GetProject godoc
// @Summary      Get a project
// @Description  Get a project by its ID
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {object}  models.Project
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id} [get]
func (h *ProjectHandler) GetProject(c *gin.Context) {

	projectID := c.Param("project_id")

	if project, err := h.projectService.GetProject(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, project)
	}
}

// GetProjects godoc
// @Summary      List projects
// @Description  Get all projects of the current user
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Project
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects [get]
func (h *ProjectHandler) GetProjects(c *gin.Context) {

	userID, _ := c.Get("userID")

	projects, err := h.projectService.GetProjects(c, userID.(string))

	if err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return

	} else {
		c.JSON(200, projects)
	}

}

// CreateBoard godoc
// @Summary      Create a board
// @Description  Create a new board inside the project
// @Tags         boards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string                true  "Project ID"
// @Param        input       body      dto.CreateBoardDto    true  "Board data"
// @Success      200         {object}  map[string]string     "ok"
// @Failure      403         {object}  map[string]string     "Forbidden"
// @Router       /projects/{project_id}/boards [post]
func (h *ProjectHandler) CreateBoard(c *gin.Context) {

	var dto dto.CreateBoardDto
	projectID := c.Param("project_id")

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	}

	dto.ProjectID = projectID

	if _, err := h.boardService.CreateBoard(c, &dto); err != nil {

		fmt.Println(err)

		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}

}

func (h *ProjectHandler) EditBoard(c *gin.Context) {

}

// DeleteBoard godoc
// @Summary      Delete a board
// @Description  Delete a board from the project
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Param        board_id    path  string  true  "Board ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id} [delete]
func (h *ProjectHandler) DeleteBoard(c *gin.Context) {

	boardID := c.Param("board_id")
	projectID := c.Param("project_id")

	if _, err := h.boardService.DeleteBoard(c, projectID, boardID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// GetBoard godoc
// @Summary      Get a board
// @Description  Get a board by its ID
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        board_id    path      string  true  "Board ID"
// @Success      200         {object}  models.Board
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id} [get]
func (h *ProjectHandler) GetBoard(c *gin.Context) {

	boardID := c.Param("board_id")

	if board, err := h.boardService.GetBoard(c, boardID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, board)
	}
}

// GetBoards godoc
// @Summary      List boards
// @Description  Get all boards of the project
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {array}   models.Board
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards [get]
func (h *ProjectHandler) GetBoards(c *gin.Context) {

	projectID := c.Param("project_id")

	if boards, err := h.boardService.GetBoards(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, boards)
	}
}

// CreateCard godoc
// @Summary      Create a card
// @Description  Create a new card inside the board
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string               true  "Project ID"
// @Param        board_id    path      string               true  "Board ID"
// @Param        input       body      dto.CreateCardDto    true  "Card data"
// @Success      200         {object}  map[string]string    "ok"
// @Failure      403         {object}  map[string]string    "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id}/cards [post]
func (h *ProjectHandler) CreateCard(c *gin.Context) {
	var dto dto.CreateCardDto

	boardID := c.Param("board_id")

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	}

	fmt.Println(dto)

	dto.BoardID = boardID

	if _, err := h.cardService.CreateCard(c, &dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// DeleteCard godoc
// @Summary      Delete a card
// @Description  Delete a card from the board
// @Tags         cards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Param        board_id    path  string  true  "Board ID"
// @Param        card_id     path  string  true  "Card ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id}/cards/{card_id} [delete]
func (h *ProjectHandler) DeleteCard(c *gin.Context) {

	cardID := c.Param("card_id")
	boardID := c.Param("board_id")

	if _, err := h.cardService.DeleteCard(c, boardID, cardID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// GetCard godoc
// @Summary      Get a card
// @Description  Get a card by its ID
// @Tags         cards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        board_id    path      string  true  "Board ID"
// @Param        card_id     path      string  true  "Card ID"
// @Success      200         {object}  models.Card
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id}/cards/{card_id} [get]
func (h *ProjectHandler) GetCard(c *gin.Context) {

	cardID := c.Param("card_id")

	if card, err := h.cardService.GetCard(c, cardID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, card)
	}
}

// GetCards godoc
// @Summary      List cards
// @Description  Get all cards of the board
// @Tags         cards
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        board_id    path      string  true  "Board ID"
// @Success      200         {array}   models.Card
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id}/cards [get]
func (h *ProjectHandler) GetCards(c *gin.Context) {

	boardID := c.Param("board_id")

	if cards, err := h.cardService.GetCards(c, boardID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, cards)
	}
}

// UpdateCardLocation godoc
// @Summary      Move a card to another board
// @Description  Update the board a card belongs to
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        board_id    path      string  true  "Old board ID"
// @Param        card_id     path      string  true  "Card ID"
// @Param        input       body      object  true  "New board"  SchemaExample({"newBoardId": "63abc..."})
// @Success      200  {object}  map[string]string  "ok"
// @Failure      400  {object}  map[string]string  "Invalid input"
// @Failure      500  {object}  map[string]string  "Failed to move card"
// @Router       /projects/{project_id}/boards/{board_id}/cards/{card_id}/updatelocation [put]
func (h *ProjectHandler) UpdateCardLocation(c *gin.Context) {

	type Body struct {
		NewBoardId string
	}

	var body Body

	oldBoardID := c.Param("board_id")
	cardID := c.Param("card_id")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(400)
		return
	}

	if err := h.cardService.UpdateCardLocation(c, cardID, oldBoardID, body.NewBoardId); err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

// CreateMember godoc
// @Summary      Add a member
// @Description  Add a new member to the project
// @Tags         members
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string                 true  "Project ID"
// @Param        input       body      dto.CreateMemberDto    true  "Member data"
// @Success      200         {object}  map[string]string      "ok"
// @Failure      403         {object}  map[string]string      "Forbidden"
// @Router       /projects/{project_id}/members [post]
func (h *ProjectHandler) CreateMember(c *gin.Context) {
	var dto dto.CreateMemberDto

	projectID := c.Param("project_id")

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	}

	dto.ProjectID = projectID

	if _, err := h.memberService.CreateMember(c, &dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// DeleteMember godoc
// @Summary      Remove a member
// @Description  Remove a member from the project
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Param        member_id   path  string  true  "Member ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members/{member_id} [delete]
func (h *ProjectHandler) DeleteMember(c *gin.Context) {

	memberID := c.Param("member_id")

	if err := h.memberService.DeleteMember(c, memberID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// GetMember godoc
// @Summary      Get a member
// @Description  Get a member by ID
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        member_id   path      string  true  "Member ID"
// @Success      200         {object}  models.Member
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members/{member_id} [get]
func (h *ProjectHandler) GetMember(c *gin.Context) {

	memberID := c.Param("member_id")

	if member, err := h.memberService.GetMember(c, memberID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, member)
	}
}

// GetMembers godoc
// @Summary      List members
// @Description  Get all members of the project
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {array}   models.Member
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members [get]
func (h *ProjectHandler) GetMembers(c *gin.Context) {

	projectID := c.Param("project_id")

	if members, err := h.memberService.GetMembers(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, members)
	}
}

// NewMessage godoc
// @Summary      Send a message
// @Description  Post a new message to the card
// @Tags         messages
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string              true  "Project ID"
// @Param        board_id    path      string              true  "Board ID"
// @Param        card_id     path      string              true  "Card ID"
// @Param        input       body      dto.NewMessageDto   true  "Message data"
// @Success      201         {object}  map[string]string   "created"
// @Failure      403         {object}  map[string]string   "Forbidden"
// @Router       /projects/{project_id}/boards/{board_id}/cards/{card_id}/messages [post]
func (h *ProjectHandler) NewMessage(c *gin.Context) {

	var dto dto.NewMessageDto

	if err := c.ShouldBindJSON(&dto); err != nil {

		c.Status(403)
		c.Abort()
		return
	}

	userID, _ := c.Get("userID")
	cardID := c.Param("card_id")

	dto.CardID = cardID
	dto.SenderID = userID.(string)

	if _, err := h.messageService.CreateMessage(c, dto); err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	c.Status(201)

}

// GetProjectSchemas godoc
// @Summary      List project schemas
// @Description  Get all schemas attached to the project
// @Tags         schemas
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {array}   models.Schema
// @Failure      500         {object}  map[string]string  "Internal error"
// @Router       /projects/{project_id}/schemas [get]
func (h *ProjectHandler) GetProjectSchemas(c *gin.Context) {

	projectID := c.Param("project_id")

	schemas, err := h.schemaService.GetSchemas(c, projectID)

	fmt.Println("===SCHEMAS===")
	fmt.Println(schemas)

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, schemas)

}

// CreateProjectSchema godoc
// @Summary      Create a project schema
// @Description  Attach a new schema to the project
// @Tags         schemas
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string                true  "Project ID"
// @Param        input       body      dto.CreateSchemaDto   true  "Schema data"
// @Success      200         {object}  map[string]string     "ok"
// @Failure      400         {object}  map[string]string     "Invalid input"
// @Failure      500         {object}  map[string]string     "Internal error"
// @Router       /projects/{project_id}/schemas [post]
func (h *ProjectHandler) CreateProjectSchema(c *gin.Context) {

	var body dto.CreateSchemaDto

	userID, _ := c.Get("userID")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(400)
		return
	}

	projectID := c.Param("project_id")

	_, err := h.schemaService.CreateSchema(c, projectID, body.Direction, body.Url, body.Name, userID.(string))

	if err != nil {
		c.Status(500)
		return
	}

}
