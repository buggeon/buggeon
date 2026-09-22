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

type BoardLoader struct {
	loader *dataloader.Loader[primitive.ObjectID, *models.Board]
}

func NewBoardLoader(repo *repositories.BoardRepo) *BoardLoader {

	batchFn := func(
		ctx context.Context,
		ids []primitive.ObjectID,
	) []*dataloader.Result[*models.Board] {
		boards, err := repo.GetBoardsByIDs(ctx, ids)

		if err != nil {
			results := make([]*dataloader.Result[*models.Board], len(ids))

			for i := range results {
				results[i] = &dataloader.Result[*models.Board]{
					Error: err,
				}
			}

			return results
		}

		boardsByID := make(
			map[primitive.ObjectID]*models.Board,
			len(boards),
		)

		for i := range boards {
			board := &boards[i]
			boardsByID[board.ID] = board
		}

		results := make([]*dataloader.Result[*models.Board], len(ids))

		for i, id := range ids {
			board, ok := boardsByID[id]

			if !ok {
				results[i] = &dataloader.Result[*models.Board]{
					Error: mongo.ErrNoDocuments,
				}
				continue
			}

			results[i] = &dataloader.Result[*models.Board]{
				Data: board,
			}
		}

		return results

	}

	return &BoardLoader{
		loader: dataloader.NewBatchedLoader(batchFn),
	}

}

func (l *BoardLoader) Load(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.Board, error) {
	return l.loader.Load(ctx, id)()
}
