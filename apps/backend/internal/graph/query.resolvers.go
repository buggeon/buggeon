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

func (r *queryResolver) Project(ctx context.Context, projectID string) (*gqlmodel.Project, error) {

	project, err := r.ProjectService.GetProject(ctx, projectID)

	return gqlmodel.NewProject(project), err
}

func (r *queryResolver) Projects(ctx context.Context, userID string) ([]*gqlmodel.Project, error) {

	projects, err := r.ProjectService.GetProjects(ctx, userID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Project, len(projects))

	for i, p := range projects {

		result[i] = gqlmodel.NewProject(p)
	}

	return result, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*gqlmodel.User, error) {

	user, err := r.UserService.GetUser(ctx, userID)

	return gqlmodel.NewUser(user), err

}

func (r *queryResolver) Board(ctx context.Context, boardID string) (*gqlmodel.Board, error) {

	board, err := r.BoardService.GetBoard(ctx, boardID)

	return gqlmodel.NewBoard(board), err

}

func (r *queryResolver) Boards(ctx context.Context, projectID string) ([]*gqlmodel.Board, error) {

	boards, err := r.BoardService.GetBoards(ctx, projectID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Board, len(boards))

	for i, b := range boards {

		result[i] = gqlmodel.NewBoard(b)

	}

	return result, nil

}

func (r *queryResolver) Card(ctx context.Context, cardID string) (*gqlmodel.Card, error) {

	card, err := r.CardService.GetCard(ctx, cardID)

	return gqlmodel.NewCard(card), err

}

func (r *queryResolver) Cards(ctx context.Context, boardID string) ([]*gqlmodel.Card, error) {

	cards, err := r.CardService.GetCards(ctx, boardID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Card, len(cards))

	for i, c := range cards {

		result[i] = gqlmodel.NewCard(c)

	}

	return result, nil

}

func (r *queryResolver) Message(ctx context.Context, messageID string) (*gqlmodel.Message, error) {

	message, err := r.MessageService.GetMessage(ctx, messageID)

	return gqlmodel.NewMessage(message), err

}

func (r *queryResolver) Schema(ctx context.Context, schemaID string) (*gqlmodel.Schema, error) {

	schema, err := r.SchemaService.GetSchema(ctx, schemaID)

	return gqlmodel.NewSchema(schema), err
}

func (r *queryResolver) Schemas(ctx context.Context, projectID string) ([]*gqlmodel.Schema, error) {

	schemas, err := r.SchemaService.GetSchemas(ctx, projectID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Schema, len(schemas))

	for i, s := range schemas {

		result[i] = gqlmodel.NewSchema(s)

	}

	return result, nil
}

func (r *queryResolver) Messages(ctx context.Context, cardID string) ([]*gqlmodel.Message, error) {

	messages, err := r.MessageService.GetMessages(ctx, cardID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Message, len(messages))

	for i, m := range messages {

		result[i] = gqlmodel.NewMessage(m)

	}

	return result, nil

}

func (r *queryResolver) Member(ctx context.Context, memberID string) (*gqlmodel.Member, error) {

	member, err := r.MemberService.GetMember(ctx, memberID)

	return gqlmodel.NewMember(member), err

}

func (r *queryResolver) Members(ctx context.Context, projectID string) ([]*gqlmodel.Member, error) {

	members, err := r.MemberService.GetMembers(ctx, projectID)

	if err != nil {
		return nil, err
	}

	result := make([]*gqlmodel.Member, len(members))

	for i, m := range members {

		result[i] = gqlmodel.NewMember(m)

	}

	return result, nil

}
