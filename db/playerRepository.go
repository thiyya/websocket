package db

import "websocket/models"

// region Lakers Players
var leBron = models.Player{
	Name: "LeBron James",
}
var carmeloAnthony = models.Player{
	Name: "Carmelo Anthony",
}
var malikMonk = models.Player{
	Name: "Malik Monk",
}
var westbrook = models.Player{
	Name: "Russell Westbrook",
}
var howard = models.Player{
	Name: "Dwight Howard",
}

// endregion

// region Celtic Players
var payton = models.Player{
	Name: "Payton Pritchard",
}
var williams = models.Player{
	Name: "Grant Williams",
}
var thomas = models.Player{
	Name: "Brodric Thomas",
}
var tatum = models.Player{
	Name: "Jayson Tatum",
}
var nesmith = models.Player{
	Name: "Aaron Nesmith",
}

// endregion

// region Bulls Players
var demar = models.Player{
	Name: "DeMar DeRozan",
}
var lavine = models.Player{
	Name: "Zach LaVine",
}
var lonzo = models.Player{
	Name: "Lonzo Ball",
}
var caruso = models.Player{
	Name: "Alex Caruso",
}
var vucevic = models.Player{
	Name: "Nikola Vucevic",
}

// endregion

// region Jazz Players
var donovan = models.Player{
	Name: "Donovan Mitchell",
}
var rudy = models.Player{
	Name: "Rudy Gobert",
}
var bojan = models.Player{
	Name: "Bojan Bogdanovic",
}
var clarkson = models.Player{
	Name: "Jordan Clarkson",
}
var conley = models.Player{
	Name: "Mike Conley Jr.",
}

// endregion

// region Warriors Players
var curry = models.Player{
	Name: "Stephen Curry",
}
var klay = models.Player{
	Name: "Klay Thompson",
}
var jordan = models.Player{
	Name: "Jordan Poole",
}
var green = models.Player{
	Name: "Draymond Green",
}
var wiseman = models.Player{
	Name: "James Wiseman",
}

// endregion

// region Suns Players
var devin = models.Player{
	Name: "Devin Booker",
}
var paul = models.Player{
	Name: "Chris Paul",
}
var deandre = models.Player{
	Name: "Deandre Ayton",
}
var bridges = models.Player{
	Name: "Mikal Bridges",
}
var johnson = models.Player{
	Name: "Cameron Johnson",
}

// endregion

func GetPlayersOfLakers() (players []*models.Player) {
	players = append(players, &leBron)
	players = append(players, &carmeloAnthony)
	players = append(players, &malikMonk)
	players = append(players, &westbrook)
	players = append(players, &howard)
	return
}

func GetPlayersOfBulls() (players []*models.Player) {
	players = append(players, &demar)
	players = append(players, &lavine)
	players = append(players, &lonzo)
	players = append(players, &caruso)
	players = append(players, &vucevic)
	return
}

func GetPlayersOfCeltics() (players []*models.Player) {
	players = append(players, &payton)
	players = append(players, &williams)
	players = append(players, &thomas)
	players = append(players, &tatum)
	players = append(players, &nesmith)
	return
}

func GetPlayersOfJazz() (players []*models.Player) {
	players = append(players, &donovan)
	players = append(players, &rudy)
	players = append(players, &bojan)
	players = append(players, &clarkson)
	players = append(players, &conley)
	return
}

func GetPlayersOfWarriors() (players []*models.Player) {
	players = append(players, &curry)
	players = append(players, &klay)
	players = append(players, &jordan)
	players = append(players, &green)
	players = append(players, &wiseman)
	return
}

func GetPlayersOfSuns() (players []*models.Player) {
	players = append(players, &devin)
	players = append(players, &paul)
	players = append(players, &deandre)
	players = append(players, &bridges)
	players = append(players, &johnson)
	return
}

// singleton pattern
var playerInstance []*models.Player

func InitPlayers() {
	if playerInstance == nil {
		players := []*models.Player{}
		players = append(players, GetPlayersOfLakers()...)
		players = append(players, GetPlayersOfCeltics()...)
		players = append(players, GetPlayersOfJazz()...)
		players = append(players, GetPlayersOfBulls()...)
		players = append(players, GetPlayersOfWarriors()...)
		players = append(players, GetPlayersOfSuns()...)
		playerInstance = players

	}
}

func GetPlayers() (players []*models.Player) {
	return playerInstance
}
