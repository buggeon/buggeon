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

package handlers

import (
	"buggeon/internal/dto"
	"buggeon/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Registration godoc
// @Summary      User registration
// @Description  Create a new user account and return access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.UserRegistrationDto  true  "Registration data"
// @Success      200    {object}  map[string]string  "accessToken and userData"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      409    {object}  map[string]string  "User already exists"
// @Router       /auth/register [post]
func (h *UserHandler) Registration(c *gin.Context) {

	var registData dto.UserRegistrationDto

	if err := c.ShouldBindJSON(&registData); err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	result, err := h.userService.Register(c, &registData)

	if err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    result.Tokens.RefreshToken,
		MaxAge:   60 * 60 * 24 * 30,
		Path:     "/api/auth/refreshtoken",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{
		"accessToken": result.Tokens.AccessToken,
		"userData": dto.UserAuthResponseDto{
			Name:      result.Name,
			Email:     result.Email,
			Login:     result.Login,
			AvatarUrl: result.AvatarUrl,
		},
	})
}

// Login godoc
// @Summary      System entry
// @Description  Authentification by password and login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.UserLoginDto  true  "Entry data"
// @Success      200    {object}  map[string]string  "accessToken and userData"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      401    {object}  map[string]string  "Invalid credentials"
// @Router       /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {

	var loginDto dto.UserLoginDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	result, err := h.userService.Login(c, &loginDto)

	if err != nil {
		c.JSON(401, gin.H{"message": "Invalid credentials"})
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    result.Tokens.RefreshToken,
		MaxAge:   60 * 60 * 24 * 30,
		Path:     "/api/auth/refreshtoken",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{
		"accessToken": result.Tokens.AccessToken,
		"userData": dto.UserAuthResponseDto{
			Name:      result.Name,
			Login:     result.Login,
			Email:     result.Email,
			AvatarUrl: result.AvatarUrl,
		},
	})
}

// RefreshAccessToken godoc
// @Summary      Access token refreshing
// @Description  Refresh access token by refresh token from httpOnly cookie
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string  "accessToken"
// @Failure      401  {object}  map[string]string  "Invalid credentials"
// @Router       /auth/refreshtoken [post]
func (h *UserHandler) RefreshAccessToken(c *gin.Context) {

	refreshToken, err := c.Cookie("refreshToken")

	if err != nil {
		fmt.Println("Couldnt find refresh token from cookie")
		c.Status(403)
		c.Abort()
		return
	}

	tokenPair, err := h.userService.RefreshAccessToken(c, refreshToken)

	if err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    tokenPair.RefreshToken,
		MaxAge:   60 * 60 * 24 * 30,
		Path:     "/api/auth/refreshtoken",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{"accessToken": tokenPair.AccessToken})
}

// SetAvatar godoc
// @Summary      Set user avatar
// @Description  Upload a new avatar file for the current user
// @Tags         users
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        user_id  path      string  true  "User ID"
// @Param        avatar   formData  file    true  "New avatar file"
// @Success      200      {object}  map[string]string  "avatarUrl"
// @Failure      401      {object}  map[string]string  "Unauthorized"
// @Failure      500      {object}  map[string]string  "Failed to upload avatar"
// @Router       /users/{user_id}/avatar [patch]
func (h *UserHandler) SetAvatar(c *gin.Context) {

	avatar, err := c.FormFile("avatar")

	if err != nil {
		c.JSON(500, "Failed to upload avatar file")
		return
	}

	userID, _ := c.Get("userID")

	avatarUrl, err := h.userService.SetAvatar(c, userID.(string), avatar)

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"avatarUrl": avatarUrl})
}
