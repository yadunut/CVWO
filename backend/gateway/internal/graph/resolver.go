package graph

import (
	"github.com/yadunut/CVWO/backend/proto/auth"
	"github.com/yadunut/CVWO/backend/proto/thread"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthClient   auth.AuthServiceClient
	ThreadClient thread.ThreadServiceClient
}
