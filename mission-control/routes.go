package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerRoutes(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/missions", getMissionsHandler)
	router.HandlerFunc(http.MethodPost, "/missions", createMissionHandler)
	router.HandlerFunc(http.MethodDelete, "/missions/:slug", deleteMissionHandler)
	router.HandlerFunc(http.MethodPost, "/missions/:slug/drones", assignDroneToMissionHandler)
	router.HandlerFunc(http.MethodPost, "/missions/:slug/backlog", addTaskToMissionBacklogHandler)
	router.HandlerFunc(http.MethodGet, "/missions/:slug/backlog", getMissionBacklogHandler)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not marshal data to json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
