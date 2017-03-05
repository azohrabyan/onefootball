package main

import (
	"fmt"
	"github.com/azohrabyan/oftask/api"
	"github.com/azohrabyan/oftask/repository"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	baseUri := "https://vintagemonster.onefootball.com"
	fmt.Printf("Connecting to %s.\n", baseUri)

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	apiClient := api.NewClient(httpClient, baseUri)

	teams := []int{96, 61, 45, 34, 2, 9, 21, 6, 5, 26}
	repo := repository.NewRepository()
	for _, id := range teams {
		team, err := apiClient.Team(id)
		if err != nil {
			log.Fatal("Error when retreiving data for team ", id)
			continue
		}
		repo.ExtractPlayersFrom(team)
	}

	sortedIds := repository.NewPlayerSorter(repo.Players()).SortBy(repository.Name)

	for i, id := range sortedIds {
		fmt.Println(formatPlayer(i, repo.Player(id)))
	}
}

func formatPlayer(i int, p *repository.Player) string {
	var teams []string
	for _, t := range p.Teams() {
		teams = append(teams, t.Name)
	}

	return fmt.Sprintf("%d. %s; %d; %s", i, p.Name, p.Age, strings.Join(teams, ", "))
}
