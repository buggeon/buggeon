package graph

import (
	"buggeon/internal/dto"
	"buggeon/internal/graph/gqlinput"
	"buggeon/internal/graph/gqlmodel"
	"context"
	"errors"
	"time"
)

func (r *mutationResolver) CreateCard(
	ctx context.Context,
	boardID string,
	input gqlinput.CreateCardInput,
) (*gqlmodel.Card, error) {

	card, err := r.CardService.CreateCard(ctx, &dto.CreateCardDto{
		BoardID:   boardID,
		Title:     input.Title,
		Content:   input.Content,
		Assignees: []string{""},
		DueDate:   input.DueDate,
		Priority:  input.Priority,
	})

	return gqlmodel.NewCard(*card), err

}

func (r *mutationResolver) UpdateCard(
	ctx context.Context,
	cardID string,
	input gqlinput.UpdateCardInput,
) (*gqlmodel.Card, error) {

	card, err := r.CardService.GetCard(ctx, cardID)

	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		card.Title = *input.Title
	}

	if input.Content != nil {
		card.Content = *input.Content
	}

	if input.DueDate != nil {

		dueDateTime, err := time.Parse("02.01.2006", *input.DueDate)

		if err != nil {
			return nil, errors.New("Invalid dueDate format")
		}

		card.DueDate = dueDateTime
	}

	if input.Status != nil {
		card.Status = *input.Status
	}

	if input.Priority != nil {
		card.Priority = *input.Priority
	}

	updatedCard, err := r.CardService.UpdateCard(ctx, cardID, &card)

	return gqlmodel.NewCard(*updatedCard), err

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
	input gqlinput.CreateBoardInput,
) (*gqlmodel.Board, error) {

	board, err := r.BoardService.CreateBoard(ctx, &dto.CreateBoardDto{
		ProjectID:   projectID,
		Name:        input.Name,
		Direction:   input.Direction,
		CardsStatus: input.CardsStatus,
		ThemeColor:  *input.ThemeColor,
	})

	return gqlmodel.NewBoard(*board), err

}

func (r *mutationResolver) UpdateBoard(
	ctx context.Context,
	boardID string,
	input gqlinput.UpdateBoardInput,
) (*gqlmodel.Board, error) {

	board, err := r.BoardService.GetBoard(ctx, boardID)

	if err != nil {
		return nil, err
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

	updatedBoard, err := r.BoardService.UpdateBoard(ctx, boardID, &board)

	return gqlmodel.NewBoard(*updatedBoard), err

}

func (r *mutationResolver) DeleteBoard(
	ctx context.Context,
	projectID string,
	boardID string,
) (bool, error) {

	return r.BoardService.DeleteBoard(ctx, projectID, boardID)

}

func (r *mutationResolver) UpdateProject(
	ctx context.Context,
	projectID string,
	input gqlinput.UpdateProjectInput,
) (*gqlmodel.Project, error) {

	project, err := r.ProjectService.GetProject(ctx, projectID)

	if err != nil {
		return nil, err
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

	updatedProject, err := r.ProjectService.UpdateProject(ctx, projectID, &project)

	return gqlmodel.NewProject(*updatedProject), err

}

func (r *mutationResolver) CreateProject(
	ctx context.Context,
	input gqlinput.CreateProjectInput,
) (*gqlmodel.Project, error) {

	project, err := r.ProjectService.CreateProject(ctx, &dto.CreateProjectDto{
		LeadID:      input.LeadID,
		Name:        input.Name,
		Description: *input.Description,
		//Members:   input.Members
	})

	return gqlmodel.NewProject(*project), err

}

func (r *mutationResolver) DeleteProject(
	ctx context.Context,
	projectID string,
) (bool, error) {

	return r.ProjectService.DeleteProject(ctx, projectID)

}

func (r *mutationResolver) UpdateUser(
	ctx context.Context,
	userID string,
	input gqlinput.UpdateUserInput,
) (*gqlmodel.User, error) {

	user, err := r.UserService.GetUser(ctx, userID)

	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Login != nil {
		user.Login = *input.Login
	}

	updatedProject, err := r.UserService.UpdateUser(ctx, userID, &user)

	return gqlmodel.NewUser(*updatedProject), err

}
