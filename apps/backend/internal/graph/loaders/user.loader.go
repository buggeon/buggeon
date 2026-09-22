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

package loaders

import (
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserLoader struct {
	loader *dataloader.Loader[primitive.ObjectID, *models.User]
}

func NewUserLoader(repo *repositories.UserRepo) *UserLoader {

	batchFn := func(
		ctx context.Context,
		ids []primitive.ObjectID,
	) []*dataloader.Result[*models.User] {
		users, err := repo.GetUsersByIDs(ctx, ids)

		if err != nil {
			results := make([]*dataloader.Result[*models.User], len(ids))

			for i := range results {
				results[i] = &dataloader.Result[*models.User]{
					Error: err,
				}
			}

			return results
		}

		usersByID := make(
			map[primitive.ObjectID]*models.User,
			len(users),
		)

		for i := range users {
			user := &users[i]
			usersByID[user.ID] = user
		}

		results := make([]*dataloader.Result[*models.User], len(ids))

		for i, id := range ids {
			user, ok := usersByID[id]

			if !ok {
				results[i] = &dataloader.Result[*models.User]{
					Error: mongo.ErrNoDocuments,
				}
				continue
			}

			results[i] = &dataloader.Result[*models.User]{
				Data: user,
			}
		}

		return results

	}

	return &UserLoader{
		loader: dataloader.NewBatchedLoader(batchFn),
	}

}

func (l *UserLoader) Load(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.User, error) {
	return l.loader.Load(ctx, id)()
}
