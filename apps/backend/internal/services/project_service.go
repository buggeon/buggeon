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
	s3storage "buggeon/internal/s3Storage"
	"context"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService struct {
	projectRepo   *repositories.ProjectRepo
	memberRepo    *repositories.MemberRepo
	boardRepo     *repositories.BoardRepo
	cardRepo      *repositories.CardRepo
	messageRepo   *repositories.MessageRepo
	memberService *MemberService
	s3Storage     *s3storage.S3Storage
}

func NewProjectService(
	projectRepo *repositories.ProjectRepo,
	memberService *MemberService,
	memberRepo *repositories.MemberRepo,
	boardRepo *repositories.BoardRepo,
	cardRepo *repositories.CardRepo,
	messageRepo *repositories.MessageRepo,
	s3Storage *s3storage.S3Storage,
) *ProjectService {
	return &ProjectService{
		projectRepo:   projectRepo,
		memberService: memberService,
		memberRepo:    memberRepo,
		boardRepo:     boardRepo,
		cardRepo:      cardRepo,
		messageRepo:   messageRepo,
		s3Storage:     s3Storage,
	}
}

func (s *ProjectService) CreateProject(projectData *dto.CreateProjectDto) (*models.Project, error) {

	projectID := primitive.NewObjectID()

	project := &models.Project{
		ID:          projectID,
		Name:        projectData.Name,
		Description: projectData.Description,
		Members:     []primitive.ObjectID{},
		Boards:      []primitive.ObjectID{},
		Schemas:     []string{},
		Progress:    projectData.Progress,
	}

	if err := s.projectRepo.CreateProject(project); err != nil {
		return &models.Project{}, err
	}

	projectData.Members = append(projectData.Members, dto.MemberDto{
		UserID:     projectData.LeadID,
		Role:       "Lead",
		Directions: []string{"managment"},
	})

	for _, member := range projectData.Members {

		createdMember, _ := s.memberService.CreateMember(&dto.CreateMemberDto{
			UserID:     member.UserID,
			ProjectID:  projectID.Hex(),
			Role:       member.Role,
			Directions: member.Directions,
		})

		if createdMember.UserID.Hex() == projectData.LeadID {
			s.projectRepo.AddLead(projectID, createdMember.ID)
			project.LeadID = createdMember.ID
		}

	}

	return project, nil

}

func (s *ProjectService) UpdateProject(projectID string, newProjectData *models.Project) (*models.Project, error) {

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return &models.Project{}, err
	}

	return newProjectData, s.projectRepo.UpdateProject(projectObjID, newProjectData)

}

func (s *ProjectService) GetProjects(userID string) ([]models.Project, error) {

	projectObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	projectsIDs, err := s.memberRepo.GetProjectsIDsByUser(projectObjID)

	if err != nil {
		return nil, err
	}

	return s.projectRepo.GetProjectsByIDs(projectsIDs)
}

func (s *ProjectService) GetProject(projectID string) (models.Project, error) {

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return models.Project{}, err
	}

	project, err := s.projectRepo.GetProject(projectObjID)

	if err != nil {
		return models.Project{}, err
	}

	return project, nil

}

func (s *ProjectService) DeleteProject(projectID string) (bool, error) {

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return false, err
	}

	err = s.projectRepo.DeleteProject(projectObjID)

	if err != nil {
		return false, err
	}

	s.memberRepo.DeleteMembers(projectObjID)

	boards, err := s.boardRepo.GetBoards(projectObjID)

	if err != nil {
		return false, err
	}

	for _, board := range boards {

		s.boardRepo.DeleteBoard(board.ID)

		cards, err := s.cardRepo.GetCards(board.ID)

		if err != nil {
			continue
		}

		for _, card := range cards {

			s.cardRepo.DeleteCard(card.ID)

			err := s.messageRepo.DeleteMessages(card.ID)

			if err != nil {
				continue
			}

		}

	}

	return true, nil

}

func (s *ProjectService) SetProjectLogo(projectID string, logo *multipart.FileHeader) error {

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return err
	}

	src, err := logo.Open()

	if err != nil {
		return err
	}

	url, err := s.s3Storage.Upload(context.Background(), fmt.Sprintf("projects/%s/logo/%s", projectID, logo.Filename), src, logo.Header.Get("Content-Type"))

	if err != nil {
		return err
	}

	s.projectRepo.SetProjectLogoUrl(projectObjID, url)

	return nil

}
