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
		Name:   "refreshToken",
		Value:  result.Tokens.RefreshToken,
		MaxAge: 60 * 60 * 24 * 30,
		Path:   "auth/refreshtoken",
		Domain: "localhost",
		//Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{"accessToken": result.Tokens.AccessToken, "userData": dto.UserAuthResponseDto{
		Name:      result.Name,
		Email:     result.Email,
		Login:     result.Login,
		AvatarUrl: result.AvatarUrl,
	}})

}

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

	} else {
		cookie := &http.Cookie{
			Name:   "refreshToken",
			Value:  result.Tokens.RefreshToken,
			MaxAge: 60 * 60 * 24 * 30,
			Path:   "auth/refreshtoken",
			Domain: "localhost",
			//Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(200, gin.H{"accessToken": result.Tokens.AccessToken, "userData": dto.UserAuthResponseDto{
			Name:      result.Name,
			Login:     result.Login,
			Email:     result.Email,
			AvatarUrl: result.AvatarUrl,
		}})
	}

}

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
		Name:   "refreshToken",
		Value:  tokenPair.RefreshToken,
		MaxAge: 60 * 60 * 24 * 30,
		Path:   "auth/refreshtoken",
		Domain: "localhost",
		//Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{"accessToken": tokenPair.AccessToken})
}

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
