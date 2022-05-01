package db

import "websocket/models"

// region Team creation

var Lakers = models.Team{
	Name:    "Los Angles Lakers",
	Players: GetPlayersOfLakers(),
}

var Celtics = models.Team{
	Name:    "Boston Celtics",
	Players: GetPlayersOfCeltics(),
}

var Bulls = models.Team{
	Name:    "Chicago Bulls",
	Players: GetPlayersOfBulls(),
}

var Jazz = models.Team{
	Name:    "Utah Jazz",
	Players: GetPlayersOfJazz(),
}

var Suns = models.Team{
	Name:    "Phoenix Suns",
	Players: GetPlayersOfSuns(),
}

var Warriors = models.Team{
	Name:    "Golden State Warriors",
	Players: GetPlayersOfWarriors(),
}

// endregion

// singleton pattern
var teamInstance []*models.Team

func InitTeams() {
	if teamInstance == nil {
		teams := []*models.Team{}
		teams = append(teams, &Lakers)
		teams = append(teams, &Celtics)
		teams = append(teams, &Bulls)
		teams = append(teams, &Jazz)
		teams = append(teams, &Suns)
		teams = append(teams, &Warriors)
		teamInstance = teams
	}
}

func GetTeams() []*models.Team {
	return teamInstance
}
