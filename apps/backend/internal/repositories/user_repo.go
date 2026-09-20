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
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		collection: db.GetCollection("buggeon", "users"),
	}
}

func NewUserRepoWithDbName(dbName string) *UserRepo {
	return &UserRepo{
		collection: db.GetCollection(dbName, "users"),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepo) UpdateUser(ctx context.Context, userID primitive.ObjectID, newUserData *models.User) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$set": newUserData},
	)

	return err

}

func (r *UserRepo) GetByLogin(ctx context.Context, userLogin string) (*models.User, error) {

	var result models.User

	err := r.collection.FindOne(ctx, bson.M{"login": userLogin}).Decode(&result)

	return &result, err
}

func (r *UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error) {

	var result models.User

	err := r.collection.FindOne(ctx, bson.M{"refresh_tokens": refreshToken}).Decode(&result)

	return &result, err
}

func (r *UserRepo) GetUser(ctx context.Context, userID string) (models.User, error) {

	userObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return models.User{}, err
	}

	var result models.User

	err = r.collection.FindOne(ctx, bson.M{"_id": userObjID}).Decode(&result)

	return result, err
}

func (r *UserRepo) GetUserPasswordHash(ctx context.Context, userLogin string) (string, error) {

	var result struct {
		Password string `bson:"password"`
	}

	err := r.collection.FindOne(ctx, bson.M{"login": userLogin}).Decode(&result)

	return result.Password, err

}

func (r *UserRepo) GetUsersByIDs(ctx context.Context, userIDs []primitive.ObjectID) ([]models.User, error) {

	if len(userIDs) == 0 {
		return []models.User{}, nil
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{
			"_id": bson.M{
				"$in": userIDs,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []models.User

	if err := cursor.All(ctx, users); err != nil {
		return nil, err
	}

	return users, nil

}

func (r *UserRepo) GetAllUsers(ctx context.Context) ([]models.User, error) {

	var users []models.User

	cursor, err := r.collection.Find(ctx, bson.M{"role": "user"})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil

}

func (r *UserRepo) UpdateRefreshToken(ctx context.Context, userID primitive.ObjectID, oldRefreshToken, newRefreshToken string) error {

	filter := bson.M{"_id": userID}
	pullUpdate := bson.M{"$pull": bson.M{"refresh_tokens": oldRefreshToken}}

	_, err := r.collection.UpdateOne(ctx, filter, pullUpdate)

	if err != nil {
		return err
	}

	pushUpdate := bson.M{"$push": bson.M{"refresh_tokens": newRefreshToken}}

	_, err = r.collection.UpdateOne(ctx, filter, pushUpdate)

	return err
}

func (r *UserRepo) AddRefreshToken(ctx context.Context, userID primitive.ObjectID, refreshToken string) error {

	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"refresh_tokens": refreshToken}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *UserRepo) SetAvatar(ctx context.Context, userID primitive.ObjectID, avatarUrl string) error {

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"avatar_url": avatarUrl}})

	return err

}
