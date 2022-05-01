package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"websocket/constants"
	"websocket/db"
)

func LeaugeTableHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("LeaugeTable Connection error : ", err)
		return
	}
	log.Println("LeaugeTable Connection established.")
	defer conn.Close()

	for {
		ticker := time.NewTicker(constants.RefreshTime * time.Second)
		done := make(chan bool)

		messageType, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("LeaugeTable Could not read message from websocket, error : ", err)
			done <- true
			break
		}

		// Ilk acıldıgında boş gorunmesin diye initial halini gönderiyor.
		var teams = db.GetTeams()
		databytes, err := json.Marshal(teams)
		if err = conn.WriteMessage(messageType, databytes); err != nil {
			log.Println("LeaugeTable Could not write message to websocket, error", err)
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
					databytes, err := json.Marshal(teams)
					if err = conn.WriteMessage(messageType, databytes); err != nil {
						log.Println("LeaugeTable Could not write message to websocket, error", err)
						return
					}
				}
			}
		}()
	}
}
