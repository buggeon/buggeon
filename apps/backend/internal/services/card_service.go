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
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardService struct {
	cardRepo             *repositories.CardRepo
	boardRepo            *repositories.BoardRepo
	messageRepo          *repositories.MessageRepo
	notififactionService *NotificationService
	memberRepo           *repositories.MemberRepo
}

func NewCardService(
	cardRepo *repositories.CardRepo,
	boardRepo *repositories.BoardRepo,
	messageRepo *repositories.MessageRepo,
	notificationService *NotificationService,
	memberRepo *repositories.MemberRepo,
) *CardService {
	return &CardService{
		cardRepo:             cardRepo,
		boardRepo:            boardRepo,
		messageRepo:          messageRepo,
		notififactionService: notificationService,
		memberRepo:           memberRepo,
	}
}

func (s *CardService) CreateCard(ctx context.Context, dto *dto.CreateCardDto) (*models.Card, error) {

	fmt.Println("Создаём карточку...")

	boardID, err := primitive.ObjectIDFromHex(dto.BoardID)
	var assigneeObjIDs []primitive.ObjectID

	for _, assigneeID := range dto.Assignees {

		assigneeObjID, err := primitive.ObjectIDFromHex(assigneeID)

		if err != nil {
			continue
		}

		assigneeObjIDs = append(assigneeObjIDs, assigneeObjID)
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
		Assignees: assigneeObjIDs,
		Status:    "inProgress",
		BoardID:   boardID,
		Priority:  dto.Priority,
		DueDate:   dueDateTime,
		Messages:  []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	cardID, err := s.cardRepo.CreateCard(ctx, card)

	err = s.boardRepo.AddCard(ctx, boardID, cardID)

	if err != nil {
		return &models.Card{}, err
	}

	members, err := s.memberRepo.GetMembersByIDs(ctx, card.Assignees)

	for _, member := range members {

		fmt.Println("Оповещаем " + member.UserID.Hex())

		_, err := s.notififactionService.NewNotification(
			ctx,
			member.UserID,
			"Appending in card",
			"You was append in card",
			map[string]any{
				"cardId": cardID,
			},
		)

		if err != nil {
			continue
		}

	}

	return card, nil

}

func (s *CardService) UpdateCard(ctx context.Context, cardID string, newCardData *models.Card) (*models.Card, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return &models.Card{}, err
	}

	return newCardData, s.cardRepo.UpdateCard(ctx, cardObjID, newCardData)

}

func (s *CardService) GetCard(ctx context.Context, cardID string) (models.Card, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return models.Card{}, err
	}

	card, err := s.cardRepo.GetCard(ctx, cardObjID)

	return card, err
}

func (s *CardService) DeleteCard(ctx context.Context, boardID, cardID string) (bool, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return false, err
	}

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	if err != nil {
		return false, err
	}

	err = s.boardRepo.DeleteCard(ctx, boardObjID, cardObjID)

	if err != nil {
		return false, err
	}

	err = s.cardRepo.DeleteCard(ctx, cardObjID)

	if err != nil {
		return false, err
	}

	err = s.messageRepo.DeleteMessages(ctx, cardObjID)

	return true, err
}

func (s *CardService) GetCards(ctx context.Context, boardID string) ([]models.Card, error) {

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	cards, err := s.cardRepo.GetCardsByBoardID(ctx, boardObjID)

	return cards, err
}

func (s *CardService) CloseCard() {

}

func (s *CardService) UpdateCardLocation(ctx context.Context, cardID, oldBoardID, newBoardID string) error {

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

	if err := s.boardRepo.DeleteCard(ctx, oldBoardObjID, cardObjID); err != nil {
		return err
	}

	if err := s.boardRepo.AddCard(ctx, newBoardObjID, cardObjID); err != nil {
		return err
	}

	newBoard, err := s.boardRepo.GetBoard(ctx, newBoardObjID)

	if err != nil {
		return err
	}

	return s.cardRepo.UpdateCardLocation(ctx, cardObjID, newBoardObjID, newBoard.CardsStatus)

}
