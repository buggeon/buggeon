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

package services

import (
	"buggeon/internal/dto"
	"buggeon/internal/models"
	"buggeon/internal/repositories"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BoardService struct {
	boardRepo   *repositories.BoardRepo
	projectRepo *repositories.ProjectRepo
	cardRepo    *repositories.CardRepo
	messageRepo *repositories.MessageRepo
}

func NewBoardService(
	boardRepo *repositories.BoardRepo,
	projectRepo *repositories.ProjectRepo,
	cardRepo *repositories.CardRepo,
	messageRepo *repositories.MessageRepo,
) *BoardService {
	return &BoardService{
		boardRepo:   boardRepo,
		projectRepo: projectRepo,
		cardRepo:    cardRepo,
		messageRepo: messageRepo,
	}
}

func (s *BoardService) CreateBoard(ctx context.Context, dto *dto.CreateBoardDto) (*models.Board, error) {

	projectID, err := primitive.ObjectIDFromHex(dto.ProjectID)

	if err != nil {
		return &models.Board{}, err
	}

	board := &models.Board{
		Name:        dto.Name,
		CardsStatus: dto.CardsStatus,
		ProjectID:   projectID,
		Direction:   dto.Direction,
		Cards:       []primitive.ObjectID{},
	}

	boardID, err := s.boardRepo.CreateBoard(ctx, board)

	if err != nil {
		return &models.Board{}, err
	}

	return board, s.projectRepo.AddBoard(ctx, projectID, boardID)

}

func (s *BoardService) UpdateBoard(ctx context.Context, boardID string, newBoardData *models.Board) (*models.Board, error) {

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	if err != nil {
		return &models.Board{}, err
	}

	return newBoardData, s.boardRepo.UpdateBoard(ctx, boardObjID, newBoardData)

}

func (s *BoardService) GetBoard(ctx context.Context, boardID string) (models.Board, error) {

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	if err != nil {
		return models.Board{}, err
	}

	boards, err := s.boardRepo.GetBoard(ctx, boardObjID)

	return boards, err
}

func (s *BoardService) GetBoards(ctx context.Context, projectID string) ([]models.Board, error) {

	boardObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return nil, err
	}

	return s.boardRepo.GetBoardsByProjectID(ctx, boardObjID)
}

func (s *BoardService) DeleteBoard(ctx context.Context, projectID, boardID string) (bool, error) {

	boardObjID, err := primitive.ObjectIDFromHex(boardID)

	if err != nil {
		return false, err
	}

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return false, err
	}

	err = s.projectRepo.DeleteBoard(ctx, projectObjID, boardObjID)

	if err != nil {
		return false, err
	}

	err = s.boardRepo.DeleteBoard(ctx, boardObjID)

	if err != nil {
		return false, err
	}

	cards, err := s.cardRepo.GetCardsByBoardID(ctx, boardObjID)

	for _, card := range cards {

		s.cardRepo.DeleteCard(ctx, card.ID)

		err := s.messageRepo.DeleteMessages(ctx, card.ID)

		if err != nil {
			continue
		}

	}

	return true, nil
}
