package handlers

import (
	"buggeon/internal/services"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	aiService *services.AIService
}

func NewAIHandler(aiService *services.AIService) *AIHandler {
	return &AIHandler{
		aiService: aiService,
	}
}

func (h *AIHandler) FullAnalytic(c *gin.Context) {

	var body struct {
		Request string `json:"request"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(403)
		c.Abort()
		return
	}

	response, err := h.aiService.FullAnalytic(c, c.Param("project_id"), body.Request)

	if err != nil {
		c.Status(500)
		c.Abort()
		return
	}

	c.JSON(200, response)

}
