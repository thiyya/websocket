package main

import (
	"net/http"
	"websocket/db"
	"websocket/server"
)

func main() {
	// repo olusturulur.
	// Gerçek bir db olmadıgı için in memory kullanıldı. Bu nedenle concurrent calısamaz.
	// Eğer concurrent olması isteniyorsa gerçek bir db kullanılması gerekiyor.
	db.InitPlayers()
	db.InitTeams()

	r := server.NewServer()
	http.ListenAndServe(":8080", r)
}
