package loaders

import (
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardLoader struct {
	loader *dataloader.Loader[primitive.ObjectID, *models.Card]
}

func NewCardLoader(repo *repositories.CardRepo) *CardLoader {

	batchFn := func(
		ctx context.Context,
		ids []primitive.ObjectID,
	) []*dataloader.Result[*models.Card] {
		cards, err := repo.GetCardsByIDs(ctx, ids)

		if err != nil {
			results := make([]*dataloader.Result[*models.Card], len(ids))

			for i := range results {
				results[i] = &dataloader.Result[*models.Card]{
					Error: err,
				}
			}

			return results
		}

		cardsByID := make(
			map[primitive.ObjectID]*models.Card,
			len(cards),
		)

		for i := range cards {
			card := &cards[i]
			cardsByID[card.ID] = card
		}

		results := make([]*dataloader.Result[*models.Card], len(ids))

		for i, id := range ids {
			Card, ok := cardsByID[id]

			if !ok {
				results[i] = &dataloader.Result[*models.Card]{
					Error: mongo.ErrNoDocuments,
				}
				continue
			}

			results[i] = &dataloader.Result[*models.Card]{
				Data: Card,
			}
		}

		return results

	}

	return &CardLoader{
		loader: dataloader.NewBatchedLoader(batchFn),
	}

}

func (l *CardLoader) Load(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.Card, error) {
	return l.loader.Load(ctx, id)()
}
