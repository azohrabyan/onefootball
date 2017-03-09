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

const MaxTeamId = 300

func main() {
	baseUri := "https://vintagemonster.onefootball.com"
	fmt.Printf("Connecting to %s.\n", baseUri)

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	apiClient := api.NewClient(httpClient, baseUri)

	teamNames :=[]string{
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester Utd",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"FC Bayern Munich",
	}
	repo := repository.NewRepository()
	for id := 1; len(teamNames) > 0 && id < MaxTeamId; id++ {
		team, err := apiClient.Team(id)
		if err != nil {
			log.Fatal("Error when retreiving data for team ", id, err)
			continue
		}
		if index := inArray(teamNames, team.Name); index != -1 {
			fmt.Println("Found team", team.Name, "extracting players");
			team, err := apiClient.TeamPlayers(id)
			if err != nil {
				log.Fatal("Error when fetching team players", id, err)
			}
			repo.ExtractPlayersFrom(team)
			teamNames = append(teamNames[:index], teamNames[index+1:]...)
		}
	}

	sortedIds := repository.NewPlayerSorter(repo.Players()).SortBy(repository.Name)

	for i, id := range sortedIds {
		fmt.Println(formatPlayer(i, repo.Player(id)))
	}
}

func inArray(a []string, needed string) int {
	for index, value := range a {
		if value == needed {
			return index
		}
	}
	return -1
}

func formatPlayer(i int, p *repository.Player) string {
	var teams []string
	for _, t := range p.Teams() {
		teams = append(teams, t.Name)
	}

	return fmt.Sprintf("%d. %s; %d; %s", i, p.Name, p.Age, strings.Join(teams, ", "))
}
