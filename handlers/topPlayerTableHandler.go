package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"websocket/constants"
	"websocket/db"
	"websocket/models"
)

func TopPlayerTableHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("TopPlayer Connection error : ", err)
		return
	}
	log.Println("TopPlayer Connection established.")
	defer conn.Close()

	for {
		ticker := time.NewTicker(constants.RefreshTime * time.Second)
		done := make(chan bool)

		messageType, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("TopPlayer Could not read message from websocket, error : ", err)
			done <- true
			break
		}

		// Her refreshTime sn de bir lig tablosunun yeni hali gönderiyor.
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					topPlayer := models.TopPlayer{
						TopScorer:   findTopScorer(),
						TopAssister: findTopAssister(),
					}
					databytes, err := json.Marshal(topPlayer)
					if err = conn.WriteMessage(messageType, databytes); err != nil {
						log.Println("TopPlayer Could not write message to websocket, error", err)
						return
					}
				}
			}
		}()
	}
}

func findTopAssister() string {
	var players = db.GetPlayers()
	topAssister := players[0]
	for _, player := range players {
		if player.TotalAssist > topAssister.TotalAssist {
			topAssister = player
		}
	}
	return topAssister.Name
}

func findTopScorer() string {
	var players = db.GetPlayers()
	topScorer := players[0]
	for _, player := range players {
		if player.TotalScore > topScorer.TotalScore {
			topScorer = player
		}
	}
	return topScorer.Name
}
