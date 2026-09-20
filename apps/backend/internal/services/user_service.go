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
	"buggeon/internal/security"
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo     *repositories.UserRepo
	tokenService *TokenService
	s3Storage    *s3storage.S3Storage
}

func NewUserService(
	userRepo *repositories.UserRepo,
	tokenService *TokenService,
	s3Storage *s3storage.S3Storage,
) *UserService {
	return &UserService{
		userRepo:     userRepo,
		tokenService: tokenService,
		s3Storage:    s3Storage,
	}
}

func (s *UserService) Register(ctx context.Context, dto *dto.UserRegistrationDto) (models.TokenResponse, error) {

	passwordHash, err := security.HashPassword(dto.Password)

	if err != nil {
		return models.TokenResponse{}, err
	}

	userID := primitive.NewObjectID()

	user := &models.User{
		ID:        userID,
		Name:      dto.Name,
		Email:     dto.Email,
		AvatarUrl: "",
		Login:     dto.Login,
		Password:  passwordHash,
		Role:      "user",
	}

	tokenPair, err := s.tokenService.GenerateTokensPair(userID.Hex())

	if err != nil {
		return models.TokenResponse{}, err
	}

	user.RefreshTokens = []string{tokenPair.RefreshToken}

	s.userRepo.CreateUser(ctx, user)

	return tokenPair, nil

}

func (s *UserService) Login(ctx context.Context, dto *dto.UserLoginDto) (models.TokenResponse, error) {

	user, err := s.userRepo.GetByLogin(ctx, dto.Login)

	if err != nil {
		return models.TokenResponse{}, err
	}

	verificationResult, _ := security.VerifyPassword(dto.Password, user.Password)

	if verificationResult == true {

		tokensPair, err := s.tokenService.GenerateTokensPair(user.ID.Hex())

		if err != nil {
			return models.TokenResponse{}, err
		}

		return tokensPair, s.userRepo.AddRefreshToken(ctx, user.ID, tokensPair.RefreshToken)

	}

	return models.TokenResponse{}, errors.New("Unathorized")

}

func (s *UserService) UpdateUser(ctx context.Context, userID string, newUserData *models.User) (*models.User, error) {

	userObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return &models.User{}, err
	}

	return newUserData, s.userRepo.UpdateUser(ctx, userObjID, newUserData)

}

func (s *UserService) RefreshAccessToken(ctx context.Context, refreshToken string) (models.TokenResponse, error) {

	user, err := s.userRepo.GetByRefreshToken(ctx, refreshToken)

	if err != nil {
		return models.TokenResponse{}, err
	}

	tokenPair, err := s.tokenService.RefreshAccessToken(refreshToken)

	if err != nil {
		return models.TokenResponse{}, err
	}

	return tokenPair, s.userRepo.UpdateRefreshToken(ctx, user.ID, refreshToken, tokenPair.RefreshToken)

}

func (s *UserService) GetUser(ctx context.Context, userID string) (models.User, error) {

	user, err := s.userRepo.GetUser(ctx, userID)

	if err != nil {
		return models.User{}, nil
	}

	return user, nil

}

func (s *UserService) SetAvatar(ctx context.Context, userID string, avatar *multipart.FileHeader) (string, error) {

	src, err := avatar.Open()

	if err != nil {
		return "", err
	}

	url, err := s.s3Storage.Upload(context.TODO(), fmt.Sprintf("users/%s/avatar/%s", userID, avatar.Filename), src, avatar.Header.Get("Content-Type"))

	if err != nil {
		return "", err
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return "", err
	}

	return url, s.userRepo.SetAvatar(ctx, userObjID, url)

}
