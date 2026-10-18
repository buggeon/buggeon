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
	_ "buggeon/internal/models"
	"buggeon/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	memberService *services.MemberService
}

func NewMemberHandler(
	memberService *services.MemberService,
) *MemberHandler {
	return &MemberHandler{
		memberService: memberService,
	}
}

// CreateMember godoc
// @Summary      Add a member
// @Description  Add a new member to the project
// @Tags         members
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string                 true  "Project ID"
// @Param        input       body      dto.CreateMemberDto    true  "Member data"
// @Success      200         {object}  map[string]string      "ok"
// @Failure      403         {object}  map[string]string      "Forbidden"
// @Router       /projects/{project_id}/members [post]
func (h *MemberHandler) CreateMember(c *gin.Context) {
	var dto dto.CreateMemberDto

	projectID := c.Param("project_id")

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	}

	dto.ProjectID = projectID

	if _, err := h.memberService.CreateMember(c, &dto); err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// DeleteMember godoc
// @Summary      Remove a member
// @Description  Remove a member from the project
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Param        member_id   path  string  true  "Member ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members/{member_id} [delete]
func (h *MemberHandler) DeleteMember(c *gin.Context) {

	memberID := c.Param("member_id")

	if err := h.memberService.DeleteMember(c, memberID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.Status(200)
	}
}

// GetMember godoc
// @Summary      Get a member
// @Description  Get a member by ID
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        member_id   path      string  true  "Member ID"
// @Success      200         {object}  models.Member
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members/{member_id} [get]
func (h *MemberHandler) GetMember(c *gin.Context) {

	memberID := c.Param("member_id")

	if member, err := h.memberService.GetMember(c, memberID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, member)
	}
}

// GetMembers godoc
// @Summary      List members
// @Description  Get all members of the project
// @Tags         members
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {array}   models.Member
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id}/members [get]
func (h *MemberHandler) GetMembers(c *gin.Context) {

	projectID := c.Param("project_id")

	if members, err := h.memberService.GetMembers(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, members)
	}
}
