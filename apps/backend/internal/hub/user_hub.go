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

package hub

type UserHub struct {
	hub *Hub
}

func NewUserHub(hub *Hub) *UserHub {
	return &UserHub{
		hub: hub,
	}
}

func (ch *UserHub) Join(c *Client, userID string) {
	c.RoomID = "user:" + userID
	ch.hub.Register(c)
}

func (ch *UserHub) Leave(c *Client) {
	ch.hub.Unregister(c)
}

func (ch *UserHub) Broadcast(msg []byte, userID string) {
	ch.hub.Broadcast("user:"+userID, msg)
}

func (ch *UserHub) BroadcastExcept(msg []byte, userID string, ex *Client) {
	ch.hub.BroadcastExcept("user:"+userID, msg, ex)
}
