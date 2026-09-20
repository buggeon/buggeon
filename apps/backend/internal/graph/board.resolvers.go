package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"buggeon/internal/graph/loaders"
	"context"
)

func (r *boardResolver) Cards(
	ctx context.Context,
	obj *gqlmodel.Board,
) ([]*gqlmodel.Card, error) {

	loader := loaders.FromContext(ctx)

	result := make([]*gqlmodel.Card, len(obj.Cards))

	for i, c := range obj.Cards {

		card, err := loader.CardLoader.Load(ctx, c)

		if err != nil {
			return nil, err
		}

		result[i] = gqlmodel.NewCard(*card)

	}

	return result, nil

}
