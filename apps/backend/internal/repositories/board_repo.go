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

type BoardRepo struct {
	collection *mongo.Collection
}

func NewBoardRepo() *BoardRepo {
	return &BoardRepo{
		collection: db.GetCollection("buggeon", "boards"),
	}
}

func NewBoardRepoWithDbName(dbName string) *BoardRepo {
	return &BoardRepo{
		collection: db.GetCollection(dbName, "boards"),
	}
}

func (r *BoardRepo) CreateBoard(ctx context.Context, board *models.Board) (primitive.ObjectID, error) {

	board.ID = primitive.NewObjectID()
	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, board)
	return board.ID, err

}

func (r *BoardRepo) UpdateBoard(ctx context.Context, boardID primitive.ObjectID, newBoardData *models.Board) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": boardID},
		bson.M{"$set": newBoardData},
	)

	return err

}

func (r *BoardRepo) DeleteBoard(ctx context.Context, boardID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": boardID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete board")
	}

	return nil

}

func (r *BoardRepo) DeleteBoards(ctx context.Context, projectID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"project_id": projectID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete boards")
	}

	return nil

}

func (r *BoardRepo) GetBoard(ctx context.Context, boardID primitive.ObjectID) (models.Board, error) {

	var board models.Board

	err := r.collection.FindOne(ctx, bson.M{"_id": boardID}).Decode(&board)

	return board, err

}

func (r *BoardRepo) GetBoardsByProjectID(ctx context.Context, projectID primitive.ObjectID) ([]models.Board, error) {

	var boards []models.Board

	cursor, err := r.collection.Find(ctx, bson.M{"project_id": projectID})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	return boards, nil

}

func (r *BoardRepo) GetBoardsByIDs(ctx context.Context, boardIDs []primitive.ObjectID) ([]models.Board, error) {

	if len(boardIDs) == 0 {
		return []models.Board{}, nil
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{
			"_id": bson.M{
				"$in": boardIDs,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var boards []models.Board

	if err := cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	return boards, nil

}

func (r *BoardRepo) AddCard(ctx context.Context, boardID, cardID primitive.ObjectID) error {

	filter := bson.M{"_id": boardID}
	update := bson.M{"$push": bson.M{"cards": cardID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *BoardRepo) DeleteCard(ctx context.Context, boardID, cardID primitive.ObjectID) error {

	filter := bson.M{"_id": boardID}
	update := bson.M{"$pull": bson.M{"cards": cardID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}
