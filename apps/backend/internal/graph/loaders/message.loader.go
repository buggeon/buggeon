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

type MessageLoader struct {
	loader *dataloader.Loader[primitive.ObjectID, *models.Message]
}

func NewMessageLoader(repo *repositories.MessageRepo) *MessageLoader {

	batchFn := func(
		ctx context.Context,
		ids []primitive.ObjectID,
	) []*dataloader.Result[*models.Message] {
		messages, err := repo.GetMessagesByIDs(ctx, ids)

		if err != nil {
			results := make([]*dataloader.Result[*models.Message], len(ids))

			for i := range results {
				results[i] = &dataloader.Result[*models.Message]{
					Error: err,
				}
			}

			return results
		}

		messagesByID := make(
			map[primitive.ObjectID]*models.Message,
			len(messages),
		)

		for i := range messages {
			message := &messages[i]
			messagesByID[message.ID] = message
		}

		results := make([]*dataloader.Result[*models.Message], len(ids))

		for i, id := range ids {
			message, ok := messagesByID[id]

			if !ok {
				results[i] = &dataloader.Result[*models.Message]{
					Error: mongo.ErrNoDocuments,
				}
				continue
			}

			results[i] = &dataloader.Result[*models.Message]{
				Data: message,
			}
		}

		return results

	}

	return &MessageLoader{
		loader: dataloader.NewBatchedLoader(batchFn),
	}

}

func (l *MessageLoader) Load(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.Message, error) {
	return l.loader.Load(ctx, id)()
}
