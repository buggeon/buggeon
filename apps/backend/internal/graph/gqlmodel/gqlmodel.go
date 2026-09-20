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
