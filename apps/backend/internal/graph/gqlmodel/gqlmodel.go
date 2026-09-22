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

import "buggeon/internal/models"

type User struct {
	models.User

	ID        string
	CreatedAt string
}

type Project struct {
	models.Project

	ID        string
	CreatedAt string
	UpdatedAt string
	LogoURL   string
	Progress  int32
}

type Message struct {
	models.Message

	ID        string
	CreatedAt string
	UpdatedAt string
	CardID    string
}

type Member struct {
	models.Member

	ID        string
	CreatedAt string
	ProjectID string
}

type Card struct {
	models.Card

	ID        string
	CreatedAt string
	UpdatedAt string
	BoardID   string
	DueDate   string
}

type Board struct {
	models.Board

	ID          string
	CreatedAt   string
	UpdatedAt   string
	ProjectID   string
	ThemeColor  string
	CardsStatus string
}

type Schema struct {
	models.Schema

	ID        string
	CreatedAt string
	UpdatedAt string
}
