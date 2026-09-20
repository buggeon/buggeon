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

type CardRepo struct {
	collection *mongo.Collection
}

func NewCardRepo() *CardRepo {
	return &CardRepo{
		collection: db.GetCollection("buggeon", "cards"),
	}
}

func NewCardRepoWithDbName() *CardRepo {
	return &CardRepo{
		collection: db.GetCollection("buggeon", "cards"),
	}
}

func (r *CardRepo) CreateCard(ctx context.Context, card *models.Card) (primitive.ObjectID, error) {

	card.ID = primitive.NewObjectID()
	card.CreatedAt = time.Now()
	card.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, card)
	return card.ID, err

}

func (r *CardRepo) UpdateCard(ctx context.Context, cardID primitive.ObjectID, newCardData *models.Card) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": cardID},
		bson.M{"$set": newCardData},
	)

	return err

}

func (r *CardRepo) DeleteCard(ctx context.Context, cardID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": cardID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete card")
	}

	return nil

}

func (r *CardRepo) DeleteCards(ctx context.Context, boardID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"board_id": boardID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete card")
	}

	return nil

}

func (r *CardRepo) GetCard(ctx context.Context, cardID primitive.ObjectID) (models.Card, error) {

	var card models.Card

	err := r.collection.FindOne(ctx, bson.M{"_id": cardID}).Decode(&card)

	return card, err

}

func (r *CardRepo) GetCardsByBoardID(ctx context.Context, boardID primitive.ObjectID) ([]models.Card, error) {

	var cards []models.Card

	cursor, err := r.collection.Find(ctx, bson.M{"board_id": boardID})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &cards); err != nil {
		return nil, err
	}

	return cards, nil

}

func (r *CardRepo) GetCardsByIDs(ctx context.Context, cardIDs []primitive.ObjectID) ([]models.Card, error) {

	if len(cardIDs) == 0 {
		return []models.Card{}, nil
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{
			"_id": bson.M{
				"$in": cardIDs,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var cards []models.Card

	if err := cursor.All(ctx, &cards); err != nil {
		return nil, err
	}

	return cards, nil

}

func (r *CardRepo) AddMessage(ctx context.Context, cardID, messageID primitive.ObjectID) error {

	filter := bson.M{"_id": cardID}
	update := bson.M{"$push": bson.M{"messages": messageID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *CardRepo) UpdateCardLocation(ctx context.Context, cardID, boardID primitive.ObjectID, newStatus string) error {

	filter := bson.M{"_id": cardID}
	update := bson.M{"$set": bson.M{"board_id": boardID, "status": newStatus}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}
