package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"buggeon/internal/graph/loaders"
	"context"
)

func (r *schemaResolver) Author(
	ctx context.Context,
	obj *gqlmodel.Schema,
) (*gqlmodel.Member, error) {

	member, err := loaders.FromContext(ctx).MemberLoader.Load(ctx, obj.AuthorID)

	if err != nil {
		return nil, err
	}

	return gqlmodel.NewMember(*member), nil

}
