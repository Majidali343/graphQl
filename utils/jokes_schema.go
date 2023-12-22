package utils

import (
	jokesData "graphQl/assets"

	"github.com/graphql-go/graphql"
)

var JokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"text":     &graphql.Field{Type: graphql.String},
			"type":     &graphql.Field{Type: graphql.String},
			"strength": &graphql.Field{Type: graphql.Int},
		},
	},
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"jokes": &graphql.Field{
				Type: graphql.NewList(JokeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					return jokesData.JokesData, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: RootQuery,
	},
)
