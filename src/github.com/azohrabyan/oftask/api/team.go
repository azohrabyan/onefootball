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

func (p *Player) UnmarshalJSON(b []byte) error {
	tmp := struct {
		ID json.RawMessage
		Name string
		Age json.RawMessage
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	unquoted, err := strconv.Unquote(string(tmp.ID))
	if err != nil {
		return err
	}
	p.ID, err = strconv.Atoi(unquoted)
	if err != nil {
		return err
	}

	p.Name = tmp.Name

	unquoted, err = strconv.Unquote(string(tmp.Age))
	if err == strconv.ErrSyntax {
		unquoted = string(tmp.Age)
	}
	p.Age, err = strconv.Atoi(unquoted)
	if err != nil {
		return err
	}

	return nil
}
