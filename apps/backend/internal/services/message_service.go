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
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	messageRepo *repositories.MessageRepo
	cardRepo    *repositories.CardRepo
}

func NewMessageService(
	messageRepo *repositories.MessageRepo,
	cardRepo *repositories.CardRepo,
) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
		cardRepo:    cardRepo,
	}
}

func (s *MessageService) CreateMessage(ctx context.Context, createMessageDto dto.NewMessageDto) (*models.Message, error) {

	var message models.Message

	senderID, err := primitive.ObjectIDFromHex(createMessageDto.SenderID)

	if err != nil {
		return &models.Message{}, err
	}

	cardID, err := primitive.ObjectIDFromHex(createMessageDto.CardID)

	if err != nil {
		return &models.Message{}, err
	}

	if createMessageDto.ReplyTo != "" {

		replyToObjID, err := primitive.ObjectIDFromHex(createMessageDto.ReplyTo)

		if err != nil {
			return &models.Message{}, err
		}

		message.ReplyTo = replyToObjID

	}

	message.CardID = cardID
	message.Content = createMessageDto.Content
	message.SenderID = senderID
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()

	messageID, err := s.messageRepo.CreateMessage(ctx, &message)

	if err != nil {
		return &models.Message{}, err
	}

	message.ID = messageID

	return &message, s.cardRepo.AddMessage(ctx, cardID, messageID)

}

func (s *MessageService) UpdateMessage(ctx context.Context, messageID string, newMessageData *models.Message) (*models.Message, error) {

	messageObjID, err := primitive.ObjectIDFromHex(messageID)

	if err != nil {
		return &models.Message{}, err
	}

	return newMessageData, s.messageRepo.UpdateMessage(ctx, messageObjID, newMessageData)

}

func (s *MessageService) DeleteMessage(ctx context.Context, cardID string, userID string) {

}

func (s *MessageService) GetMessages(ctx context.Context, cardID string) ([]models.Message, error) {

	cardObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return nil, err
	}

	return s.messageRepo.GetMessagesByCardID(ctx, cardObjID)

}

func (s *MessageService) GetMessage(ctx context.Context, messageID string) (models.Message, error) {

	messageObjID, err := primitive.ObjectIDFromHex(messageID)

	if err != nil {
		return models.Message{}, nil
	}

	return s.messageRepo.GetMessage(ctx, messageObjID)

}
