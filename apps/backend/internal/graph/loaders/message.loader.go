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
