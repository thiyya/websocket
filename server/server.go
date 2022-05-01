package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"websocket/frontend"
	"websocket/handlers"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", frontend.HomePage())
	})
	r.Path("/match-result").HandlerFunc(handlers.MatchResultHandler)
	r.Path("/leauge-table").HandlerFunc(handlers.LeaugeTableHandler)
	r.Path("/top-scorer-assister-table").HandlerFunc(handlers.TopPlayerTableHandler)

	fileserver := http.FileServer(http.Dir("./frontend"))
	r.PathPrefix("/").Handler(fileserver)
	return r
}
