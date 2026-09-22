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
