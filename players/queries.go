package players

import (
	"capture_the_flag_mongodb_go_server.com/m/DB"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

func GetPlayer(args struct{ ID string }) (*PlayerResolver, error) {
	db, _ := DB.Session.GetSession()
	collection := db.Model("Player")
	player := &Player{}
	if !bson.IsObjectIdHex(args.ID) {
		return nil, errors.New("invalid ObjectID")
	}
	collection.FindOne(
		bson.M{"_id": bson.ObjectIdHex(args.ID)}).
		Exec(player)

	return &PlayerResolver{p: player}, nil
}



func GetPlayers(args PaginationInput) (*[]*PlayerResolver, error) {
	c, _ := DB.Session.GetSession()
	collection := c.Model("Player")
	var players []*Player

	limit := 0
	offset := 0

	if args.First != nil {
		limit = int(*args.First)
	}
	if args.Offset != nil {
		offset = int(*args.Offset)
	}
	collection.Find().
		Limit(int(limit)).
		Skip(int(offset)).
		Exec(&players)

	var ret []*PlayerResolver
	for _, player := range players {
		ret = append(ret, &PlayerResolver{p: player})
	}

	return &ret, nil
}
