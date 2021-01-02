package players

type Location struct {
	Longitude *float64 `json:"longitude" bson:"longitude"`
	Latitude  *float64 `json:"latitude" bson:"latitude"`
}

type LocationResolver struct {
	l *Location
}

func (r *PlayerResolver) Location() (*LocationResolver, error) {
	return &LocationResolver{l: &r.p.Location}, nil
}

func (r *LocationResolver) Longitude () float64 {
	return *r.l.Longitude
}

func (r *LocationResolver) Latitude () float64 {
	return *r.l.Latitude
}
