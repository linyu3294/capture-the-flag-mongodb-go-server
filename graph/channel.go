package graph

import (
	"capture_the_flag_mongodb_go_server.com/m/players"
)

/*
	Resolver binding for channel module.
*/

func (r *Resolver) Player(args struct{ ID string }) (*players.PlayerResolver, error) {
	return players.GetPlayer(args)
}

func (r *Resolver) Players(args players.PaginationInput) (*[]*players.PlayerResolver, error) {
	return players.GetPlayers(args)
}

func (r *Resolver) CreatePlayer(args struct{Input *players.CreatePlayerInput }) (*players.PlayerResolver, error) {
	return players.CreatePlayer(args.Input)
}

func (r *Resolver) UpdatePlayer(args struct{Input *players.UpdatePlayerInput }) (*players.PlayerResolver, error) {
	print(args.Input.Patch)
	return players.UpdatePlayer(args.Input)
}

func (r *Resolver) DeletePlayer(args struct{Input *players.DeletePlayerInput }) (*players.PlayerResolver, error) {
	return players.DeletePlayer(args.Input)
}