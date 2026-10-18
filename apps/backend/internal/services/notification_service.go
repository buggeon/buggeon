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
	"buggeon/internal/hub"
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationService struct {
	notificationRepo *repositories.NotificationRepo
	userHub          *hub.UserHub
}

func NewNotificationService(
	notificationRepo *repositories.NotificationRepo,
	userHub *hub.UserHub,
) *NotificationService {
	return &NotificationService{
		notificationRepo: notificationRepo,
		userHub:          userHub,
	}
}

func (s *NotificationService) NewNotification(ctx context.Context, userID primitive.ObjectID, title, content string, payload map[string]any) (*models.Notification, error) {

	notificationID := primitive.NewObjectID()

	notification := &models.Notification{
		ID:        notificationID,
		Title:     title,
		Content:   content,
		Payload:   payload,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	_, err := s.notificationRepo.CreateNotification(ctx, notification)

	notificationBytes, err := json.Marshal(notification)

	if err != nil {
		return &models.Notification{}, err
	}

	go s.userHub.Broadcast(notificationBytes, userID.Hex())

	return notification, err

}

func (s *NotificationService) SendNotificationByEmail(ctx context.Context, email string, notification string) {

}

func (s *NotificationService) SendNotificationByTelegram(ctx context.Context, chatID string, notification models.Notification) {

}
