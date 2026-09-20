package graph

import (
	"buggeon/internal/graph/gqlmodel"
	"context"
)

func (r *messageResolver) Sender(
	ctx context.Context,
	obj *gqlmodel.Message,
) (*gqlmodel.User, error) {

	user, err := r.UserService.GetUser(ctx, obj.SenderID.Hex())

	if err != nil {
		return nil, err
	}

	return gqlmodel.NewUser(user), err

}

func (r *messageResolver) ReplyTo(
	ctx context.Context,
	obj *gqlmodel.Message,
) (*gqlmodel.Message, error) {

	message, err := r.MessageService.GetMessage(ctx, obj.ID)

	if err != nil {
		return nil, err
	}

	return gqlmodel.NewMessage(message), err

}

func (r *messageResolver) Replies(
	ctx context.Context,
	obj *gqlmodel.Message,
) ([]*gqlmodel.Message, error) {

	result := make([]*gqlmodel.Message, 0)

	return result, nil

}
