// Buggeon - SelfHosted service for bug and task tracking
// Copyright (C) 2026 DEVE corp.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
