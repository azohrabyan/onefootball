package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAPIResponse_Unmarshall(t *testing.T) {
	jsonString, err := os.Open("test-fixtures/api-response.json")
	assert.Nil(t, err)

	actual := NewAPIResponse("team", &Team{})

	jsonParser := json.NewDecoder(jsonString)
	err = jsonParser.Decode(&actual)

	assert.Nil(t, err)
	assert.Equal(t, APIResponse{
		Status: "ok",
		Code:   0,
		Data: map[string]Entity{
			"team": &Team{
				ID:   96,
				Name: "Germany",
			},
		},
		Message: "Team feed successfully generated. Api Version: 1",
	}, *actual)
}
