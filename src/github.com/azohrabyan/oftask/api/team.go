package api

import (
	"encoding/json"
	"strconv"
)

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
}

type TeamPlayers struct {
	Team
	Players []Player `json:"players"`
}

type Player struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Age  int    `json:"age,string"`
}
