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

package gqlmodel

import (
	"buggeon/internal/models"
	"time"
)

func NewUser(user models.User) *User {
	return &User{
		User:      user,
		ID:        user.ID.Hex(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func NewProject(project models.Project) *Project {
	return &Project{
		Project:   project,
		ID:        project.ID.Hex(),
		CreatedAt: project.CreatedAt.Format(time.RFC3339),
		UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		LogoURL:   project.LogoUrl,
		Progress:  int32(project.Progress),
	}
}

func NewMessage(message models.Message) *Message {
	return &Message{
		Message:   message,
		ID:        message.ID.Hex(),
		CreatedAt: message.CreatedAt.Format(time.RFC3339),
		UpdatedAt: message.UpdatedAt.Format(time.RFC3339),
		CardID:    message.CardID.Hex(),
	}
}

func NewMember(member models.Member) *Member {
	return &Member{
		Member:    member,
		ID:        member.ID.Hex(),
		CreatedAt: member.CreatedAt.Format(time.RFC3339),
		ProjectID: member.ProjectID.Hex(),
	}
}

func NewCard(card models.Card) *Card {
	return &Card{
		Card:      card,
		ID:        card.ID.Hex(),
		CreatedAt: card.CreatedAt.Format(time.RFC3339),
		UpdatedAt: card.UpdatedAt.Format(time.RFC3339),
		BoardID:   card.BoardID.Hex(),
		DueDate:   card.DueDate.Format(time.RFC3339),
	}
}

func NewBoard(board models.Board) *Board {
	return &Board{
		Board:       board,
		ID:          board.ID.Hex(),
		CreatedAt:   board.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   board.UpdatedAt.Format(time.RFC3339),
		ProjectID:   board.ProjectID.Hex(),
		CardsStatus: board.CardsStatus,
	}
}

func NewSchema(schema models.Schema) *Schema {
	return &Schema{
		Schema:    schema,
		ID:        schema.ID.Hex(),
		CreatedAt: schema.CreatedAt.Format(time.RFC3339),
		UpdatedAt: schema.UpdatedAt.Format(time.RFC3339),
	}
}
