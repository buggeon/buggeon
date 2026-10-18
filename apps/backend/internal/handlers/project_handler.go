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
	_ "buggeon/internal/models"
	"buggeon/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *services.ProjectService
	schemaService  *services.SchemaService
	boardService   *services.BoardService
	cardService    *services.CardService
	memberService  *services.MemberService
	messageService *services.MessageService
}

func NewProjectHandler(
	projectService *services.ProjectService,
	schemaService *services.SchemaService,
	cardService *services.CardService,
	boardServices *services.BoardService,
	memberService *services.MemberService,
	messageService *services.MessageService,
) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		boardService:   boardServices,
		schemaService:  schemaService,
		cardService:    cardService,
		memberService:  memberService,
		messageService: messageService,
	}
}

// CreateProject godoc
// @Summary      Create a project
// @Description  Create a new project with the current user as lead
// @Tags         projects
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      dto.CreateProjectDto  true  "Project data"
// @Success      201    {object}  models.Project  "Created project"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      401    {object}  map[string]string  "Unauthorized"
// @Router       /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {

	// leadID := c.Request.FormValue("leadId")
	// name := c.Request.FormValue("name")
	// description := c.Request.FormValue("description")
	// //members := c.Request.FormValue("members")
	// logo, _ := c.FormFile("logo")

	// err := h.projectService.CreateProject(&dto.CreateProjectDto{
	// 	LeadID:      leadID,
	// 	Name:        name,
	// 	Description: description,
	// 	Logo:        logo,
	// 	Progress:    0,
	// 	//Members: string(membersJSON),
	// })

	// if err != nil {
	// 	c.Status(403)
	// 	c.Abort()
	// 	return
	// }

	// c.Status(201)

}

// SetProjectLogo godoc
// @Summary      Set project logo
// @Description  Upload a new logo file for the project
// @Tags         projects
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Param        logo        formData  file    true  "New logo file"
// @Success      200         {object}  map[string]string  "ok"
// @Failure      403         {object}  map[string]string  "New logo file was not provided"
// @Failure      500         {object}  map[string]string  "Failed to upload new logo file"
// @Router       /projects/{project_id}/logo [patch]
func (h *ProjectHandler) SetProjectLogo(c *gin.Context) {

	file, err := c.FormFile("logo")
	projectID := c.Param("project_id")

	if err != nil {
		c.JSON(403, "New logo file was not provide")
		return
	}

	err = h.projectService.SetProjectLogo(c, projectID, file)

	if err != nil {
		c.JSON(500, "Failed to upload new logo file")
		return
	}

}

// DeleteProject godoc
// @Summary      Delete a project
// @Description  Delete a project by its ID
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path  string  true  "Project ID"
// @Success      200  {object}  map[string]string  "ok"
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {

	projectID := c.Param("project_id")

	_, err := h.projectService.DeleteProject(c, projectID)

	if err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	c.Status(200)
}

// GetProject godoc
// @Summary      Get a project
// @Description  Get a project by its ID
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Param        project_id  path      string  true  "Project ID"
// @Success      200         {object}  models.Project
// @Failure      403         {object}  map[string]string  "Forbidden"
// @Router       /projects/{project_id} [get]
func (h *ProjectHandler) GetProject(c *gin.Context) {

	projectID := c.Param("project_id")

	if project, err := h.projectService.GetProject(c, projectID); err != nil {
		c.Status(403)
		c.Abort()
		return
	} else {
		c.JSON(200, project)
	}
}

// GetProjects godoc
// @Summary      List projects
// @Description  Get all projects of the current user
// @Tags         projects
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Project
// @Failure      403  {object}  map[string]string  "Forbidden"
// @Router       /projects [get]
func (h *ProjectHandler) GetProjects(c *gin.Context) {

	userID, _ := c.Get("userID")

	projects, err := h.projectService.GetProjects(c, userID.(string))

	if err != nil {
		fmt.Println(err)
		c.Status(403)
		c.Abort()
		return

	} else {
		c.JSON(200, projects)
	}

}
