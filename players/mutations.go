package players

import (
	"capture_the_flag_mongodb_go_server.com/m/DB"
	"github.com/pkg/errors"
	"github.com/zebresel-com/mongodm"
	"gopkg.in/mgo.v2/bson"
)

func CreatePlayer(input *CreatePlayerInput) (*PlayerResolver, error) {
	db, _ := DB.Session.GetSession()
	collection := db.Model("Player")

	player := &Player{}
	collection.New(player)
	player.FirstName = input.Player.FirstName
	player.LastName = input.Player.LastName
	player.Location = Location(*input.Player.Location)

	err := player.Save()
	return &PlayerResolver{player}, err
}

func UpdatePlayer(input *UpdatePlayerInput) (*PlayerResolver, error) {
	player := &Player{}
	db, _ := DB.Session.GetSession()
	collection := db.Model("Player")
	//err := collection.Update(bson.M{"_id": input.ID}, input.patch)

	if !bson.IsObjectIdHex(input.ID) {
		return nil, errors.New("invalid ObjectID")
	}
	err := collection.FindOne(
		bson.M{"_id": bson.ObjectIdHex(input.ID)}).
		Exec(player)

	player.FirstName = input.Patch.FirstName
	player.LastName = input.Patch.LastName
	player.Location = Location(*input.Patch.Location)

	player.Save()
	return &PlayerResolver{player}, err
}

func DeletePlayer(input *DeletePlayerInput) (*PlayerResolver, error) {
	player := &Player{}
	db, _ := DB.Session.GetSession()
	collection := db.Model("Player")

	if !bson.IsObjectIdHex(input.ID) {
		return nil, errors.New("invalid ObjectID")
	}
	err := collection.Remove(
		bson.M{"_id": bson.ObjectIdHex(input.ID)})
	if _, ok := err.(*mongodm.NotFoundError); ok {
		return nil, err
	}


	return &PlayerResolver{player}, err
}
