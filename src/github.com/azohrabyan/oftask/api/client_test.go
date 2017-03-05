package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Team(t *testing.T) {
	ts := registerTestHandler(t, "/api/teams/en/11.json", `{ "data": { "team":  {"id": 555} } }`)
	defer ts.Close()

	h := http.Client{}
	client := NewClient(h, ts.URL)
	team, err := client.Team(11)

	assert.Nil(t, err)
	assert.Equal(t, team.ID, 555)
}

func TestClient_TeamInvalidJSON(t *testing.T) {
	ts := registerTestHandler(t, "/api/teams/en/11.json", `"id": 555`)
	defer ts.Close()

	h := http.Client{}
	client := NewClient(h, ts.URL)
	team, err := client.Team(11)

	assert.Error(t, err)
	assert.Nil(t, team)
}

func registerTestHandler(t *testing.T, expectedUrl string, response string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, expectedUrl, r.URL.RequestURI())
		fmt.Fprint(w, response)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))

	return ts
}
