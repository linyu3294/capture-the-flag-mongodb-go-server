package graph

import (
	"capture_the_flag_mongodb_go_server.com/m/players"
	"github.com/graph-gophers/graphql-go"
)

var GQLSchema *graphql.Schema

func init() {
	GQLSchema = graphql.MustParseSchema(Schema, &Resolver{})
}

/*
	Schema implement queries, mutations, resolvers and inputs from each individual module
*/

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		` + players.Queries + `
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		` + players.Mutations + `
	}

	# A character from the Star Wars universe
	interface Entity {
		# The ID of the entity
		id: String!
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
	}

	` + players.Interfaces + `
	` + players.Inputs + `
	` + players.Types + `

	union SearchResult = Player
`