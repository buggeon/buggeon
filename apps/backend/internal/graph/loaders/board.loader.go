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
