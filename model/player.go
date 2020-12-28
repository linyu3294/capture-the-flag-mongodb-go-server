package model

type Player struct {
	id   int
	name string
	flag int
}

func NewUser(id int , name string, flag int) *Player {
	return &Player{
		id:   id,
		name: name,
		flag: flag,
	}
}

func (u *Player) GetID() int {
	return u.id
}

func (u *Player) GetName() string {
	return u.name
}

func (u *Player) GetFlag() int {
	return u.flag
}
