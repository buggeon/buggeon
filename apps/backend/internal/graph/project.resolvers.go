package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"buggeon/internal/graph/loaders"
	"context"
)

func (r *projectResolver) Members(
	ctx context.Context,
	obj *gqlmodel.Project,
) ([]*gqlmodel.Member, error) {

	loaders := loaders.FromContext(ctx)

	result := make([]*gqlmodel.Member, len(obj.Members))

	for i, m := range obj.Members {

		member, err := loaders.MemberLoader.Load(ctx, m)

		if err != nil {
			return nil, err
		}

		result[i] = gqlmodel.NewMember(*member)

	}

	return result, nil

}

func (r *projectResolver) Lead(
	ctx context.Context,
	obj *gqlmodel.Project,
) (*gqlmodel.Member, error) {

	member, err := loaders.FromContext(ctx).MemberLoader.Load(ctx, obj.LeadID)

	if err != nil {
		return nil, err
	}

	return gqlmodel.NewMember(*member), nil

}

func (r *projectResolver) Boards(
	ctx context.Context,
	obj *gqlmodel.Project,
) ([]*gqlmodel.Board, error) {

	loaders := loaders.FromContext(ctx)

	result := make([]*gqlmodel.Board, len(obj.Boards))

	for i, b := range obj.Boards {

		board, err := loaders.BoardLoader.Load(ctx, b)

		if err != nil {
			return nil, err
		}

		result[i] = gqlmodel.NewBoard(*board)

	}

	return result, nil

}
