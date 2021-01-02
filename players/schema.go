package players

// Queries used in queries.go
var Queries = `
	player(id : String!): Player
	players(first:Int, offset: Int) : [Player]
`

// Mutations used mutations.go
var Mutations = `
	createPlayer(input :CreatePlayerInput!) :Player
	updatePlayer(input :UpdatePlayerInput!) :Player
	deletePlayer(input :DeletePlayerInput!) :Player
`

// Inputs used in Mutations
var Inputs = `
	input CreatePlayerInput {
		player: PlayerInput!
	}
	input UpdatePlayerInput {
		id : String!
		patch : PlayerInput
	}
	input DeletePlayerInput {
		id: String!
	}
	# Used in creating/updating a new player
	input PlayerInput {
		firstname : String
		lastname  : String
		location: LocationInput
	}
	# Used in creating/updating a new Location
	input LocationInput {
		longitude : Float
		latitude: Float
	}
`

// Interface definitions used in Types
var Interfaces = `
	interface IPlayer {
		location(): Location!
	}
`

// Type definitions linked in player.go
var Types = `
	type Location {
		longitude : Float!
		latitude : Float!
	}

	# A Player Entity
	type Player implements Entity, IPlayer {
		# The ID of the channel
		id: String!
		# What this human calls themselves
		firstname: String!
		# What this human calls themselves
		lastname: String!
		# The updated Location of the entity
		location: Location!
		# The created time of the entity
		created: String!
		# The updated time of the entity
		updated: String!
	}
`
