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

type CardHandler struct {
	boardService *services.BoardService
	cardService  *services.CardService
}

func NewCardHandler(
	cardService *services.CardService,
	boardServices *services.BoardService,
) *CardHandler {
	return &CardHandler{
		boardService: boardServices,
		cardService:  cardService,
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
func (h *CardHandler) CreateCard(c *gin.Context) {
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
func (h *CardHandler) DeleteCard(c *gin.Context) {

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
func (h *CardHandler) GetCard(c *gin.Context) {

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
func (h *CardHandler) GetCards(c *gin.Context) {

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
func (h *CardHandler) UpdateCardLocation(c *gin.Context) {

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
