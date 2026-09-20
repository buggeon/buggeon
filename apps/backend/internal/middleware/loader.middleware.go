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

package middleware

import (
	"buggeon/internal/graph/loaders"
	"buggeon/internal/repositories"

	"github.com/gin-gonic/gin"
)

type LoaderMiddleware struct {
	userRepo    *repositories.UserRepo
	cardRepo    *repositories.CardRepo
	boardRepo   *repositories.BoardRepo
	messageRepo *repositories.MessageRepo
	memberRepo  *repositories.MemberRepo
}

func NewLoaderMiddleware(
	userRepo *repositories.UserRepo,
	cardRepo *repositories.CardRepo,
	boardRepo *repositories.BoardRepo,
	messageRepo *repositories.MessageRepo,
	memberRepo *repositories.MemberRepo,
) *LoaderMiddleware {
	return &LoaderMiddleware{
		userRepo:    userRepo,
		cardRepo:    cardRepo,
		boardRepo:   boardRepo,
		messageRepo: messageRepo,
		memberRepo:  memberRepo,
	}
}

func (m *LoaderMiddleware) SetLoaderMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ls := &loaders.Loaders{
			UserLoader:    loaders.NewUserLoader(m.userRepo),
			BoardLoader:   loaders.NewBoardLoader(m.boardRepo),
			CardLoader:    loaders.NewCardLoader(m.cardRepo),
			MemberLoader:  loaders.NewMemberLoader(m.memberRepo),
			MessageLoader: loaders.NewMessageLoader(m.messageRepo),
		}

		requectCtx := loaders.WithLoaders(ctx, ls)

		ctx.Request = ctx.Request.WithContext(requectCtx)

		ctx.Next()
	}
}
