package graph

import (
	"capture_the_flag_mongodb_go_server.com/m/DB"
	"capture_the_flag_mongodb_go_server.com/m/players"
)

func RegisterNodes() {
	s, _ := DB.Session.GetSession()

	// Register each module (collection) here
	s.Register(&players.Player{}, "Player")
}
