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

package repositories

import (
	"buggeon/internal/db"
	"buggeon/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepo struct {
	collection *mongo.Collection
}

func NewNotificationRepo() *NotificationRepo {
	return &NotificationRepo{
		collection: db.GetCollection("buggeon", "notifications"),
	}
}

func NewNotificationRepoWithDbName() *NotificationRepo {
	return &NotificationRepo{
		collection: db.GetCollection("buggeon", "notifications"),
	}
}

func (r *NotificationRepo) CreateNotification(ctx context.Context, Notification *models.Notification) (primitive.ObjectID, error) {

	Notification.ID = primitive.NewObjectID()
	Notification.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, Notification)
	return Notification.ID, err

}

func (r *NotificationRepo) UpdateNotification(ctx context.Context, notificationID primitive.ObjectID, newNotificationData *models.Notification) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": notificationID},
		bson.M{"$set": newNotificationData},
	)

	return err

}

func (r *NotificationRepo) DeleteNotification(ctx context.Context, notificationID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": notificationID})

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("Failed to delete notification")
	}

	return nil

}

func (r *NotificationRepo) DeleteNotifications(ctx context.Context, userID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"user_id": userID})

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("Failed to delete notification")
	}

	return nil

}

func (r *NotificationRepo) GetNotification(ctx context.Context, notificationID primitive.ObjectID) (models.Notification, error) {

	var notification models.Notification

	err := r.collection.FindOne(ctx, bson.M{"_id": notificationID}).Decode(&notification)

	return notification, err

}

func (r *NotificationRepo) GetNotificationsByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Notification, error) {

	var notifications []models.Notification

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil

}

func (r *NotificationRepo) GetNotificationsByIDs(ctx context.Context, notificationIDs []primitive.ObjectID) ([]models.Notification, error) {

	if len(notificationIDs) == 0 {
		return []models.Notification{}, nil
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{
			"_id": bson.M{
				"$in": notificationIDs,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var notifications []models.Notification

	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil

}
