package graph

import (
	"buggeon/graph/model"
	"buggeon/internal/dto"
	"context"
	"errors"
	"time"
)

func (r *mutationResolver) CreateCard(
	ctx context.Context,
	boardID string,
	input model.CreateCardInput,
) (*model.Card, error) {

	card, err := r.CardService.CreateCard(&dto.CreateCardDto{
		BoardID:   boardID,
		Title:     input.Title,
		Content:   input.Content,
		Assignees: []string{""},
		DueDate:   input.DueDate,
		Priority:  input.Priority,
	})

	return r.toGraphQLCard(context.Background(), card), err

}

func (r *mutationResolver) UpdateCard(
	ctx context.Context,
	cardID string,
	input model.UpdateCardInput,
) (*model.Card, error) {

	card, err := r.CardService.GetCard(cardID)

	if err != nil {
		return &model.Card{}, err
	}

	if input.Title != nil {
		card.Title = *input.Title
	}

	if input.Content != nil {
		card.Content = *input.Content
	}

	// if input.Assignees != nil {

	// 	newAssignees

	// 	card.Assignees = *input.Assignees
	// }

	if input.DueDate != nil {

		dueDateTime, err := time.Parse("02.01.2006", *input.DueDate)

		if err != nil {
			return &model.Card{}, errors.New("Invalid dueDate format")
		}

		card.DueDate = dueDateTime
	}

	if input.Status != nil {
		card.Status = *input.Status
	}

	if input.Priority != nil {
		card.Priority = *input.Priority
	}

	updatedCard, err := r.CardService.UpdateCard(cardID, &card)

	return r.toGraphQLCard(context.Background(), updatedCard), err

}

func (r *mutationResolver) DeleteCard(
	ctx context.Context,
	boardID string,
	cardID string,
) (bool, error) {
	panic("Not implemented")
}

func (r *mutationResolver) CreateBoard(
	ctx context.Context,
	projectID string,
	input model.CreateBoardInput,
) (*model.Board, error) {

	board, err := r.BoardService.CreateBoard(&dto.CreateBoardDto{
		ProjectID:   projectID,
		Name:        input.Name,
		Direction:   input.Direction,
		CardsStatus: input.CardsStatus,
		ThemeColor:  *input.ThemeColor,
	})

	return r.toGraphQLBoard(context.Background(), board), err

}

func (r *mutationResolver) UpdateBoard(
	ctx context.Context,
	boardID string,
	input model.UpdateBoardInput,
) (*model.Board, error) {

	board, err := r.BoardService.GetBoard(boardID)

	if err != nil {
		return &model.Board{}, err
	}

	if input.Name != nil {
		board.Name = *input.Name
	}

	if input.Direction != nil {
		board.Direction = *input.Direction
	}

	if input.ThemeColor != nil {
		board.ThemeColor = *input.ThemeColor
	}

	if input.CardsStatus != nil {
		board.CardsStatus = *input.CardsStatus
	}

	updatedBoard, err := r.BoardService.UpdateBoard(boardID, &board)

	return r.toGraphQLBoard(context.Background(), updatedBoard), err

}

func (r *mutationResolver) DeleteBoard(
	ctx context.Context,
	projectID string,
	boardID string,
) (bool, error) {

	return r.BoardService.DeleteBoard(projectID, boardID)

}

func (r *mutationResolver) UpdateProject(
	ctx context.Context,
	projectID string,
	input model.UpdateProjectInput,
) (*model.Project, error) {

	project, err := r.ProjectService.GetProject(projectID)

	if err != nil {
		return &model.Project{}, err
	}

	if input.Name != nil {
		project.Name = *input.Name
	}

	if input.Description != nil {
		project.Description = *input.Description
	}

	if input.Progress != nil {
		project.Progress = int(*input.Progress)
	}

	updatedProject, err := r.ProjectService.UpdateProject(projectID, &project)

	return r.toGraphQLProject(context.Background(), updatedProject), err

}

func (r *mutationResolver) CreateProject(
	ctx context.Context,
	input model.CreateProjectInput,
) (*model.Project, error) {

	project, err := r.ProjectService.CreateProject(&dto.CreateProjectDto{
		LeadID:      input.LeadID,
		Name:        input.Name,
		Description: *input.Description,
		//Members:   input.Members
	})

	projectGraph := r.toGraphQLProject(context.Background(), project)

	return projectGraph, err

}

func (r *mutationResolver) DeleteProject(
	ctx context.Context,
	projectID string,
) (bool, error) {

	return r.ProjectService.DeleteProject(projectID)

}
