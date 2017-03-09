package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	http    http.Client
	baseUri string
}

func NewClient(http http.Client, baseUri string) *Client {
	return &Client{http, baseUri}
}

func (c *Client) Team(id int) (*Team, error) {
	apiResponse := NewAPIResponse("team", &Team{})

	url := fmt.Sprintf("%s/api/teams/en/%d.json", c.baseUri, id)
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	buf, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err := json.Unmarshal(buf, apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.Data["team"].(*Team), nil
}

func (c *Client) TeamPlayers(id int) (*TeamPlayers, error) {
	apiResponse := NewAPIResponse("team", &TeamPlayers{})

	url := fmt.Sprintf("%s/api/teams/en/%d.json", c.baseUri, id)
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	buf, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err := json.Unmarshal(buf, apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.Data["team"].(*TeamPlayers), nil
}
