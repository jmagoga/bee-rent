package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/jmagoga/new-equimper-go-graphql/graph/model"
	"github.com/jmagoga/new-equimper-go-graphql/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	requests []*model.Request
	bees []*model.Bee

	// these are added here so they can be accessed inside of schema.resolvers.go
	RequestsRepo postgres.RequestsRepo
	BeesRepo postgres.BeesRepo
}
