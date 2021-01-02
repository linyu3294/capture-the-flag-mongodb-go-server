package players

import "github.com/zebresel-com/mongodm"

// Player struct
type Player struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	FirstName            *string  `json:"firstname" bson:"firstname"`
	LastName             *string  `json:"lastname" bson:"lastname"`
	Location             Location `json:"location" bson:"location"`
}

// Player resolver contains a reference to a Player
type PlayerResolver struct {
	p *Player
}

// Setters and Getters of Players
func (r *PlayerResolver) FirstName() string {
	return *r.p.FirstName
}
func (r *PlayerResolver) LastName() string {
	return *r.p.LastName
}
func (r *LocationResolver) Location() Location {
	return r.Location()
}

// Mongo defined values and operations - ID, CreatedAt, and UpdatedAt
func (r *PlayerResolver) ID() string {
	return r.p.Id.Hex()
}

func (r *PlayerResolver) Created() string {
	return r.p.CreatedAt.String()
}

func (r *PlayerResolver) Updated() string {
	return r.p.UpdatedAt.String()
}
