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
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemberService struct {
	memberRepo  *repositories.MemberRepo
	projectRepo *repositories.ProjectRepo
}

func NewMemberService(
	memberRepo *repositories.MemberRepo,
	projectRepo *repositories.ProjectRepo,
) *MemberService {
	return &MemberService{
		memberRepo:  memberRepo,
		projectRepo: projectRepo,
	}
}

func (s *MemberService) CreateMember(ctx context.Context, dto *dto.CreateMemberDto) (*models.Member, error) {

	userID, err := primitive.ObjectIDFromHex(dto.UserID)

	if err != nil {
		return &models.Member{}, err
	}

	projectID, err := primitive.ObjectIDFromHex(dto.ProjectID)

	if err != nil {
		return &models.Member{}, err
	}

	member := &models.Member{
		UserID:     userID,
		Role:       dto.Role,
		ProjectID:  projectID,
		Directions: dto.Directions,
		CreatedAt:  time.Now(),
	}

	memberID, err := s.memberRepo.CreateMember(ctx, member)

	return member, s.projectRepo.AddMember(ctx, projectID, memberID)
}

func (s *MemberService) UpdateMember(ctx context.Context, memberID string, newMemberData *models.Member) (*models.Member, error) {

	memberObjID, err := primitive.ObjectIDFromHex(memberID)

	if err != nil {
		return &models.Member{}, err
	}

	return newMemberData, s.memberRepo.UpdateMember(ctx, memberObjID, newMemberData)

}

func (s *MemberService) GetMember(ctx context.Context, memberID string) (models.Member, error) {

	memberObjID, err := primitive.ObjectIDFromHex(memberID)

	if err != nil {
		fmt.Println(err)

		return models.Member{}, err
	}

	member, err := s.memberRepo.GetMember(ctx, memberObjID)

	return member, err

}

func (s *MemberService) DeleteMember(ctx context.Context, memberID string) error {

	memberObjID, err := primitive.ObjectIDFromHex(memberID)

	if err != nil {
		return err
	}

	err = s.memberRepo.DeleteMember(ctx, memberObjID)

	return err

}

func (s *MemberService) GetMembers(ctx context.Context, projectID string) ([]models.Member, error) {

	projectObjID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return nil, err
	}

	members, err := s.memberRepo.GetMembersByProjectID(ctx, projectObjID)

	return members, err

}
