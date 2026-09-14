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
	"buggeon/config"
	"buggeon/internal/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func testConfig() config.Config {
	return config.Config{
		AccessTokenSecret:  "test-access-secret-key-12345",
		RefreshTokenSecret: "test-refresh-secret-key-67890",
		AccessTokenExpire:  15,
		RefreshTokenExpire: 24,
		TokensIssuer:       "buggeon-test",
	}
}

func testUser() *models.User {

	userID, err := primitive.ObjectIDFromHex("6a9ffd23ff85e2540dbeec42")

	if err != nil {
		return &models.User{}
	}

	return &models.User{
		ID:    userID,
		Name:  "Test",
		Login: "@test",
		Email: "test@gmail.com",
	}
}

var (
	service = NewTokenService(testConfig())
	user    = testUser()
)

func validateToken(t *testing.T, token string, validate func(string) (*TokenClaims, error)) {

	if token == "" {
		t.Errorf("Expected token, got empty string")
	}

	claims, err := validate(token)

	if err != nil {
		t.Errorf("Expected claims, got %v", err)
	}

	if claims.UserEmail != user.Email {
		t.Errorf("Expected value '%v' on field 'email', got '%v'", user.Email, claims.UserEmail)
	}

	if claims.UserName != user.Name {
		t.Errorf("Expected value '%v' on field 'name', got '%v'", user.Name, claims.UserName)
	}

	if claims.UserLogin != user.Login {
		t.Errorf("Expected value '%v' on field 'login', got '%v'", user.Login, claims.UserLogin)
	}

	if claims.UserID != user.ID.Hex() {
		t.Errorf("Expected value '%v' on field 'id', got '%v'", user.ID.Hex(), claims.UserID)
	}
}

func TestTokenService_GenerateAccessToken_Success(t *testing.T) {

	token, err := service.GenerateAccessToken(user)

	if err != nil {
		t.Errorf("Expected token, got %v", err)
	}

	validateToken(t, token, service.ValidateAccessToken)

}

func TestTokenService_GenerateRefreshToken_Success(t *testing.T) {

	token, err := service.GenerateRefreshToken(user)

	if err != nil {
		t.Errorf("Expected token, got %v", err)
	}

	validateToken(t, token, service.ValidateRefreshToken)

}

func TestTokenService_RefreshAccessToken_Success(t *testing.T) {

	tokenPair, err := service.GenerateTokensPair(testUser())

	if err != nil {
		t.Errorf("Expected token pair, got %v", err)
	}

	newTokenPair, err := service.RefreshAccessToken(tokenPair.RefreshToken)

	validateToken(t, newTokenPair.AccessToken, service.ValidateAccessToken)

}
