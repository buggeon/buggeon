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

type CardHub struct {
	hub *Hub
}

func NewCardHub(hub *Hub) *CardHub {
	return &CardHub{
		hub: hub,
	}
}

func (ch *CardHub) Join(c *Client, cardID string) {
	c.RoomID = "card:" + cardID
	ch.hub.Register(c)
}

func (ch *CardHub) Leave(c *Client) {
	ch.hub.Unregister(c)
}

func (ch *CardHub) Broadcast(msg []byte, cardID string) {
	ch.hub.Broadcast("card:"+cardID, msg)
}

func (ch *CardHub) BroadcastExcept(msg []byte, cardID string, ex *Client) {
	ch.hub.BroadcastExcept("card:"+cardID, msg, ex)
}
