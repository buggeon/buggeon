package cache

import (
	"sync"
	"time"
)

type ProjectPermissions struct {
	Members   map[string][]string
	ExpiresAt time.Time
}

type PermissionsCache struct {
	mu       sync.RWMutex
	projects map[string]*ProjectPermissions
	ttl      time.Duration
}

func NewPermissionCache(ttl time.Duration) *PermissionsCache {
	return &PermissionsCache{
		projects: make(map[string]*ProjectPermissions),
		ttl:      ttl,
	}
}

func (c *PermissionsCache) Get(projectID, userID string) []string {

	c.mu.RLock()
	defer c.mu.RUnlock()

	p, ok := c.projects[projectID]

	if !ok || time.Now().After(p.ExpiresAt) {
		return nil
	}

	return p.Members[userID]

}

func (c *PermissionsCache) Set(projectID, userID string, permissions []string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.projects[projectID].Members[userID] = permissions

}

func (c *PermissionsCache) Invalidate(userID string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.projects, userID)

}

func (c *PermissionsCache) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.projects = make(map[string]*ProjectPermissions)

}
