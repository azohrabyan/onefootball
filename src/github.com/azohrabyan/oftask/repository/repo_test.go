package repository

import (
	"github.com/azohrabyan/oftask/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeam_AddPlayer(tst *testing.T) {
	t := Team{ID: 1}

	p := Player{ID: 11}
	t.AddPlayer(&p)

	assert.Len(tst, t.players, 1)
	assert.Len(tst, p.teams, 1)
}

func TestRepository_ExtractFrom(tst *testing.T) {
	at := api.Team{
		ID: 1,
		Players: []api.Player{
			{ID: 11},
			{ID: 12},
		},
	}

	r := NewRepository()

	r.ExtractPlayersFrom(&at)

	assert.Len(tst, r.teams, 1)
	assert.Len(tst, r.players, 2)

	assert.Equal(tst, r.teams[1].players[0].ID, 11)
	assert.Equal(tst, r.teams[1].players[1].ID, 12)

	assert.Equal(tst, r.players[11].ID, 11)
	assert.Equal(tst, r.players[12].ID, 12)
}

func TestRepository_ExtractFrom_PlayerDifferentTeams(tst *testing.T) {
	at1 := api.Team{
		ID: 1,
		Players: []api.Player{
			{ID: 11},
			{ID: 12},
		},
	}
	at2 := api.Team{
		ID: 2,
		Players: []api.Player{
			{ID: 11},
			{ID: 21},
		},
	}

	r := NewRepository()

	r.ExtractPlayersFrom(&at1)
	r.ExtractPlayersFrom(&at2)

	assert.Len(tst, r.teams, 2)
	assert.Len(tst, r.players, 3)

	assert.Equal(tst, 1, r.players[11].teams[0].ID)
	assert.Equal(tst, 2, r.players[11].teams[1].ID)
}
