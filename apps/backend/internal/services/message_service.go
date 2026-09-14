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

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	messageRepo *repositories.MessageRepo
	cardRepo    *repositories.CardRepo
}

func NewMessageService(messageRepo *repositories.MessageRepo, cardRepo *repositories.CardRepo) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
		cardRepo:    cardRepo,
	}
}

func (s *MessageService) CreateMessage(message dto.NewMessageDto) (*models.Message, error) {

	senderID, err := primitive.ObjectIDFromHex(message.SenderID)

	if err != nil {
		return &models.Message{}, err
	}

	cardID, err := primitive.ObjectIDFromHex(message.CardID)
	var messageID primitive.ObjectID

	if err != nil {
		return &models.Message{}, err
	}

	if message.ReplyTo != "" {
		replyTo, err := primitive.ObjectIDFromHex(message.ReplyTo)

		if err != nil {
			return &models.Message{}, err
		}

		id, err := s.messageRepo.NewMessage(&models.Message{
			SenderID: senderID,
			CardID:   cardID,
			ReplyTo:  replyTo,
			Content:  message.Content,
		})

		if err != nil {
			return &models.Message{}, err
		}

		messageID = id

	} else {
		id, err := s.messageRepo.NewMessage(&models.Message{
			SenderID: senderID,
			CardID:   cardID,
			Content:  message.Content,
		})

		if err != nil {
			return &models.Message{}, err
		}

		messageID = id
	}

	return &models.Message{}, s.cardRepo.AddMessage(cardID, messageID)

}

func (s *MessageService) UpdateMessage(messageID string, newMessageData *models.Message) (*models.Message, error) {

	messageObjID, err := primitive.ObjectIDFromHex(messageID)

	if err != nil {
		return &models.Message{}, err
	}

	return newMessageData, s.messageRepo.UpdateMessage(messageObjID, newMessageData)

}

func (s *MessageService) DeleteMessage(cardID string, userID string) {

}

func (s *MessageService) GetMessages(cardID string) ([]models.Message, error) {

	messageObjID, err := primitive.ObjectIDFromHex(cardID)

	if err != nil {
		return nil, err
	}

	return s.messageRepo.GetMessages(messageObjID)

}

func (s *MessageService) GetMessage(messageID string) (models.Message, error) {

	messageObjID, err := primitive.ObjectIDFromHex(messageID)

	if err != nil {
		return models.Message{}, nil
	}

	return s.messageRepo.GetMessage(messageObjID)

}
