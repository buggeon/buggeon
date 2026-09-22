package middleware

import (
	"buggeon/internal/cache"
	"slices"

	"github.com/gin-gonic/gin"
)

type PermissionsMiddleware struct {
	permissionsCache *cache.PermissionsCache
}

func NewPermissionsMiddleware(permissionsCache *cache.PermissionsCache) *PermissionsMiddleware {
	return &PermissionsMiddleware{
		permissionsCache: permissionsCache,
	}
}

func (m *PermissionsMiddleware) HasPermission(permission string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		userID, _ := ctx.Get("userID")
		projectID := ctx.Param("project_id")

		if ok := slices.Contains(m.permissionsCache.Get(projectID, userID.(string)), permission); !ok {
			ctx.JSON(403, "You don't have permission")
			return
		}

		ctx.Next()

	}

}
