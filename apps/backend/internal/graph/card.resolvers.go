package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"buggeon/internal/graph/loaders"
	"context"
)

func (r *cardResolver) Assignees(
	ctx context.Context,
	obj *gqlmodel.Card,
) ([]*gqlmodel.Member, error) {

	loader := loaders.FromContext(ctx)

	result := make([]*gqlmodel.Member, len(obj.Assignees))

	for i, a := range obj.Assignees {

		member, err := loader.MemberLoader.Load(ctx, a)

		if err != nil {
			return nil, err
		}

		result[i] = gqlmodel.NewMember(*member)

	}

	return result, nil
}

func (r *cardResolver) Messages(
	ctx context.Context,
	obj *gqlmodel.Card,
) ([]*gqlmodel.Message, error) {

	loaders := loaders.FromContext(ctx)

	result := make([]*gqlmodel.Message, len(obj.Messages))

	for i, m := range obj.Messages {

		message, err := loaders.MessageLoader.Load(ctx, m)

		if err != nil {
			return nil, err
		}

		result[i] = gqlmodel.NewMessage(*message)

	}

	return result, nil

}
