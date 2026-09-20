package loaders

import (
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberLoader struct {
	loader *dataloader.Loader[primitive.ObjectID, *models.Member]
}

func NewMemberLoader(repo *repositories.MemberRepo) *MemberLoader {

	batchFn := func(
		ctx context.Context,
		ids []primitive.ObjectID,
	) []*dataloader.Result[*models.Member] {
		members, err := repo.GetMembersByIDs(ctx, ids)

		if err != nil {
			results := make([]*dataloader.Result[*models.Member], len(ids))

			for i := range results {
				results[i] = &dataloader.Result[*models.Member]{
					Error: err,
				}
			}

			return results
		}

		membersByID := make(
			map[primitive.ObjectID]*models.Member,
			len(members),
		)

		for i := range members {
			member := &members[i]
			membersByID[member.ID] = member
		}

		results := make([]*dataloader.Result[*models.Member], len(ids))

		for i, id := range ids {
			member, ok := membersByID[id]

			if !ok {
				results[i] = &dataloader.Result[*models.Member]{
					Error: mongo.ErrNoDocuments,
				}
				continue
			}

			results[i] = &dataloader.Result[*models.Member]{
				Data: member,
			}
		}

		return results

	}

	return &MemberLoader{
		loader: dataloader.NewBatchedLoader(batchFn),
	}

}

func (l *MemberLoader) Load(
	ctx context.Context,
	id primitive.ObjectID,
) (*models.Member, error) {
	return l.loader.Load(ctx, id)()
}
