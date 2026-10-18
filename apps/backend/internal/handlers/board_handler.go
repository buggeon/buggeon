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

type BoardHandler struct {
	boardService *services.BoardService
	cardService  *services.CardService
}

func NewBoardHandler(
	cardService *services.CardService,
	boardServices *services.BoardService,
) *BoardHandler {
	return &BoardHandler{
		boardService: boardServices,
		cardService:  cardService,
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
func (h *BoardHandler) CreateBoard(c *gin.Context) {

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

func (h *BoardHandler) EditBoard(c *gin.Context) {

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
func (h *BoardHandler) DeleteBoard(c *gin.Context) {

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
func (h *BoardHandler) GetBoard(c *gin.Context) {

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
func (h *BoardHandler) GetBoards(c *gin.Context) {

	projectID := c.Param("project_id")

	if boards, err := h.boardService.GetBoards(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, boards)
	}
}
