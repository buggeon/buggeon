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
