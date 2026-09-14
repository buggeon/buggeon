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

package services

import (
	"buggeon/internal/dto"
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardService struct {
	cardRepo    *repositories.CardRepo
	boardRepo   *repositories.BoardRepo
	messageRepo *repositories.MessageRepo
}

func NewCardService(
	cardRepo *repositories.CardRepo,
	boardRepo *repositories.BoardRepo,
	messageRepo *repositories.MessageRepo,
) *CardService {
	return &CardService{
		cardRepo:    cardRepo,
		boardRepo:   boardRepo,
		messageRepo: messageRepo,
	}
}

func (s *CardService) CreateCard(dto *dto.CreateCardDto) (*models.Card, error) {

	boardID, err := primitive.ObjectIDFromHex(dto.BoardID)
	var assigneeIDs []primitive.ObjectID

	for _, assigneeID := range dto.Assignees {
		cardObjID, err := primitive.ObjectIDFromHex(assigneeID)

		if err != nil {
			continue
		}

		assigneeIDs = append(assigneeIDs, cardObjID)
	}

	if err != nil {
		return &models.Card{}, err
	}

	dueDateTime, err := time.Parse("02.01.2006", dto.DueDate)

	if err != nil {
		fmt.Println(dto.DueDate)
		return &models.Card{}, errors.New("Invalid dueDate format")
	}

	card := &models.Card{
		Title:     dto.Title,
		Content:   dto.Content,
		Assignees: assigneeIDs,
		Status:    "inProgress",
		BoardID:   boardID,
		Priority:  dto.Priority,
		DueDate:   dueDateTime,
		Messages:  []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	cardID, err := s.cardRepo.CreateCard(card)

	return card, s.boardRepo.AddCard(boardID, cardID)
}

func (s *CardService) UpdateCard(cardID string, newCardData *models.Card) (*models.Card, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return &models.Card{}, err
	}

	return newCardData, s.cardRepo.UpdateCard(cardObjID, newCardData)

}

func (s *CardService) GetCard(cardID string) (models.Card, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return models.Card{}, err
	}

	card, err := s.cardRepo.GetCard(cardObjID)

	return card, err
}

func (s *CardService) DeleteCard(boardID, cardID string) (bool, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return false, err
	}

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	if err != nil {
		return false, err
	}

	err = s.boardRepo.DeleteCard(boardObjID, cardObjID)

	if err != nil {
		return false, err
	}

	err = s.cardRepo.DeleteCard(cardObjID)

	if err != nil {
		return false, err
	}

	err = s.messageRepo.DeleteMessages(cardObjID)

	return true, err
}

func (s *CardService) GetCards(boardID string) ([]models.Card, error) {

	cardObjID, err := primitive.ObjectIDFromHex(boardID)

	cards, err := s.cardRepo.GetCards(cardObjID)

	return cards, err
}

func (s *CardService) CloseCard() {

}

func (s *CardService) UpdateCardLocation(cardID, oldBoardID, newBoardID string) error {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return err
	}

	oldBoardObjID, err := primitive.ObjectIDFromHex(oldBoardID)

	if err != nil {
		return err
	}

	newBoardObjID, err := primitive.ObjectIDFromHex(newBoardID)

	if err != nil {
		return err
	}

	if err := s.boardRepo.DeleteCard(oldBoardObjID, cardObjID); err != nil {
		return err
	}

	if err := s.boardRepo.AddCard(newBoardObjID, cardObjID); err != nil {
		return err
	}

	newBoard, err := s.boardRepo.GetBoard(newBoardObjID)

	if err != nil {
		return err
	}

	return s.cardRepo.UpdateCardLocation(cardObjID, newBoardObjID, newBoard.CardsStatus)

}
