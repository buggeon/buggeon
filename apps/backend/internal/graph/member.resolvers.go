package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"buggeon/internal/graph/loaders"
	"context"
)

func (r *memberResolver) User(
	ctx context.Context,
	obj *gqlmodel.Member,
) (*gqlmodel.User, error) {

	user, err := loaders.FromContext(ctx).UserLoader.Load(ctx, obj.UserID)

	if err != nil {
		return nil, err
	}

	return gqlmodel.NewUser(*user), nil

}
