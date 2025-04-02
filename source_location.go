package slogdriver

import "context"

func WithoutSourceLocation(ctx context.Context) context.Context {
	return context.WithValue(ctx, sourceLocationContextKeyT{}, false)
}

func WithSourceLocation(ctx context.Context) context.Context {
	return context.WithValue(ctx, sourceLocationContextKeyT{}, true)
}

func shouldAddSourceLocation(ctx context.Context) bool {
	value, ok := ctx.Value(sourceLocationContextKeyT{}).(bool)
	return !ok || value
}

type sourceLocationContextKeyT struct{}
