package graph

import "github.com/Crushtain/publicationService.git/storage"

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Storage storage.Storage
}

func NewResolver(storage storage.Storage) *Resolver {
	return &Resolver{Storage: storage}
}
