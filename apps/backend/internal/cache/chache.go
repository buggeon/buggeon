package cache

import (
	"sync"
	"time"
)

type CachedUser struct {
	ID        string
	Name      string
	Login     string
	Email     string
	AvatarUrl string
	ExpiresAt time.Time
}

type UserCache struct {
	mu    sync.RWMutex
	users map[string]*CachedUser
	ttl   time.Duration
}

func NewUserCache(ttl time.Duration) *UserCache {

	return &UserCache{
		users: make(map[string]*CachedUser),
		ttl:   ttl,
	}

}

func (c *UserCache) Get(userID string) *CachedUser {

	c.mu.RLock()
	defer c.mu.RUnlock()

	u, ok := c.users[userID]

	if !ok || time.Now().After(u.ExpiresAt) {
		return nil
	}

	return u

}

func (c *UserCache) Set(u *CachedUser) {

	c.mu.Lock()
	defer c.mu.Unlock()

	u.ExpiresAt = time.Now().Add(c.ttl)
	c.users[u.ID] = u

}

func (c *UserCache) Invalidate(userID string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.users, userID)

}

func (c *UserCache) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.users = make(map[string]*CachedUser)

}
