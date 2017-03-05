package repository

import (
	"github.com/azohrabyan/oftask/api"
)

type Team struct {
	ID      int
	Name    string
	players []*Player
}

func NewTeam(at *api.Team) *Team {
	t := new(Team)
	t.ID = at.ID
	t.Name = at.Name

	return t
}

func (t *Team) AddPlayer(p *Player) {
	p.teams = append(p.teams, t)
	t.players = append(t.players, p)
}

type Player struct {
	ID    int
	Name  string
	Age   int
	teams []*Team
}

func NewPlayer(ap *api.Player) *Player {
	p := new(Player)
	p.ID = ap.ID
	p.Name = ap.Name
	p.Age = ap.Age

	return p
}

func (p *Player) Teams() []*Team {
	return p.teams
}

type Repository struct {
	players map[int]*Player
	teams   map[int]*Team
}

func NewRepository() Repository {
	return Repository{
		players: map[int]*Player{},
		teams:   map[int]*Team{},
	}
}

func (r *Repository) ExtractPlayersFrom(apiTeam *api.Team) {
	team := r.newTeamIfNotExists(apiTeam)
	for _, apiPlayer := range apiTeam.Players {
		player := r.newPlayerIfNotExists(&apiPlayer)
		team.AddPlayer(player)
	}
}

func (r *Repository) newTeamIfNotExists(apiTeam *api.Team) *Team {
	team, ok := r.teams[apiTeam.ID]
	if !ok {
		team = NewTeam(apiTeam)
		r.teams[team.ID] = team
	}
	return team
}

func (r *Repository) newPlayerIfNotExists(apiPlayer *api.Player) *Player {
	player, ok := r.players[apiPlayer.ID]
	if !ok {
		player = NewPlayer(apiPlayer)
		r.players[player.ID] = player
	}
	return player
}

func (r *Repository) Players() map[int]*Player {
	return r.players
}

func (r *Repository) Player(id int) *Player {
	return r.players[id]
}
