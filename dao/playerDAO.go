package dao

// playerDAO persists player data in database
type PlayerDAO struct{}

// NewplayerDAO creates a new playerDAO
func NewplayerDAO() *PlayerDAO {
	return &PlayerDAO{}
}

//// Get reads the player with the specified ID from the database.
//func (dao *PlayerDAO) Get(rs app.RequestScope, id int) (*model.Player) {
//	var player model.Player
//	return &player
//}



//// Create saves a new player record in the database.
//// The player.Id field will be populated with an automatically generated ID upon successful saving.
//func (dao *PlayerDAO) Create(rs app.RequestScope, player *model.Player) error {
//	player.Id = 0
//	return rs.Tx().Model(player).Insert()
//}
//
//// Update saves the changes to an player in the database.
//func (dao *PlayerDAO) Update(rs app.RequestScope, id int, player *model.Player) error {
//	if _, err := dao.Get(rs, id); err != nil {
//		return err
//	}
//	player.ID = id
//	return rs.Tx().Model(player).Exclude("Id").Update()
//}
//
//// Delete deletes an player with the specified ID from the database.
//func (dao *PlayerDAO) Delete(rs app.RequestScope, id int) error {
//	player, err := dao.Get(rs, id)
//	if err != nil {
//		return err
//	}
//	return rs.Tx().Model(player).Delete()
//}
