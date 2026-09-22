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
