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

type MessageRepo struct {
	collection *mongo.Collection
}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{
		collection: db.GetCollection("buggeon", "messages"),
	}
}

func NewMessageRepoWithDbName(dbName string) *MessageRepo {
	return &MessageRepo{
		collection: db.GetCollection(dbName, "messages"),
	}
}

func (r *MessageRepo) NewMessage(ctx context.Context, message *models.Message) (primitive.ObjectID, error) {

	message.ID = primitive.NewObjectID()
	message.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, message)
	return message.ID, err

}

func (r *MessageRepo) UpdateMessage(ctx context.Context, messageID primitive.ObjectID, newMessageData *models.Message) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": messageID},
		bson.M{"$set": newMessageData},
	)

	return err

}

func (r *MessageRepo) DeleteMessage(ctx context.Context, messageID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": messageID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete message")
	}

	return nil

}

func (r *MessageRepo) DeleteMessages(ctx context.Context, cardID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"card_id": cardID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete messages")
	}

	return nil

}

func (r *MessageRepo) GetMessage(ctx context.Context, messageID primitive.ObjectID) (models.Message, error) {

	var message models.Message

	err := r.collection.FindOne(ctx, bson.M{"_id": messageID}).Decode(&message)

	return message, err

}

func (r *MessageRepo) GetMessagesByCardID(ctx context.Context, cardID primitive.ObjectID) ([]models.Message, error) {

	var members []models.Message

	cursor, err := r.collection.Find(ctx, bson.M{"card_id": cardID})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	return members, nil

}

func (r *MessageRepo) GetMessagesByIDs(ctx context.Context, messageIDs []primitive.ObjectID) ([]models.Message, error) {

	if len(messageIDs) == 0 {
		return []models.Message{}, nil
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{
			"_id": bson.M{
				"$in": messageIDs,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var messages []models.Message

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil

}
