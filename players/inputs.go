package players

// Used as inputs to perform Mutations, Queries and Subscriptions
type PlayerInput struct {
	FirstName *string
	LastName  *string
	Location  *LocationInput
}

// Used in mutations.go
type CreatePlayerInput struct {
	Player *PlayerInput
}

// Used in mutations.go
type UpdatePlayerInput struct {
	ID    string
	Patch *PlayerInput
}

// Used in mutations.go
type DeletePlayerInput struct {
	ID string
}

// Not currently in use
type LocationInput struct {
	Longitude *float64
	Latitude  *float64
}

// Not currently in use
type ModifyLocationInput struct {
	Longitude *float64
	Latitude  *float64
}

// Used in queries.go
type PaginationInput struct {
	First  *int32
	Offset *int32
}
