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

import "buggeon/internal/services"

type Resolver struct {
	ProjectService *services.ProjectService
	SchemaService  *services.SchemaService
	BoardService   *services.BoardService
	CardService    *services.CardService
	MemberService  *services.MemberService
	UserService    *services.UserService
	MessageService *services.MessageService
}

type queryResolver struct {
	*Resolver
}

type mutationResolver struct {
	*Resolver
}

type boardResolver struct {
	*Resolver
}

type memberResolver struct {
	*Resolver
}

type cardResolver struct {
	*Resolver
}

type messageResolver struct {
	*Resolver
}

type projectResolver struct {
	*Resolver
}

type schemaResolver struct {
	*Resolver
}

type userResolver struct {
	*Resolver
}

func (r *Resolver) Board() BoardResolver {
	return &boardResolver{r}
}

func (r *Resolver) Card() CardResolver {
	return &cardResolver{r}
}

func (r *Resolver) Member() MemberResolver {
	return &memberResolver{r}
}

func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}

func (r *Resolver) Project() ProjectResolver {
	return &projectResolver{r}
}

func (r *Resolver) Schema() SchemaResolver {
	return &schemaResolver{r}
}

func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
