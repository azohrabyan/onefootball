package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTeam_Unmarshall(t *testing.T) {
	jsonString, err := os.Open("test-fixtures/team.json")
	assert.Nil(t, err)

	actual := TeamPlayers{}

	jsonParser := json.NewDecoder(jsonString)
	err = jsonParser.Decode(&actual)

	assert.Nil(t, err)
	assert.Equal(t, TeamPlayers{
		Team: Team{
			ID:   96,
			Name: "Germany",
		},
		Players: []Player{
			{
				ID:   1,
				Name: "foo",
				Age:  24,
			},
			{
				ID:   2,
				Name: "bar",
				Age:  26,
			},
		},
	}, actual)
}
