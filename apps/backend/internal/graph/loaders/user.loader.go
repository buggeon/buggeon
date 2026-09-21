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
