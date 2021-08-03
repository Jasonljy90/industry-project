package server

import (
	"encoding/json"
	"net/http"
)

// ALLRooms is the global hashmap for the server
var ALLRooms RoomMap

// Create a room and return roomID
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
	roomID := ALLRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

// Join the client in a particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
}
