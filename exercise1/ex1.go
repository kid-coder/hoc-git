package main

import "fmt"

type User struct {
	ID             int
	Name, Location string
}

type Player struct {
	*User
	GameID int
}

func (u *User) Gretting() string {
	return fmt.Sprintf("Hello from %s %s", u.Name, u.Location)
}

func NewPlayer(id int, name, location string, gameid int) *Player {
	return &Player{
		User:   &User{id, name, location},
		GameID: gameid,
	}
}

func main() {
	p := NewPlayer(10, "Tung", "HaNoi", 1)
	fmt.Printf(p.Gretting())
}
