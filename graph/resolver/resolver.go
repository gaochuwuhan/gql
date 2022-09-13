package resolver

import (
	"github.com/gaochuwuhan/gql/graph/generated"
	"github.com/gaochuwuhan/gql/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run -mod=mod github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *mongo.Client
}

func New(mongo_uri string) (*generated.Config, error) {
	client, err := pkg.MongoClient(mongo_uri)
	if client != nil && err == nil {
		return &generated.Config{
			Resolvers: &Resolver{
				DB: client,
			},
		}, nil
	}
	return nil, err
}
