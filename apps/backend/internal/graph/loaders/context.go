package loaders

import "context"

type Loaders struct {
	MemberLoader  *MemberLoader
	CardLoader    *CardLoader
	UserLoader    *UserLoader
	MessageLoader *MessageLoader
	BoardLoader   *BoardLoader
}

type contextKey struct{}

func WithLoaders(ctx context.Context, loaders *Loaders) context.Context {
	return context.WithValue(ctx, contextKey{}, loaders)
}

func FromContext(ctx context.Context) *Loaders {
	return ctx.Value(contextKey{}).(*Loaders)
}
