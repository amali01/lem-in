package funcs

import (
	"log"
	"strings"
)

// AddTunnel adds a tunnel connection between two rooms
func AddTunnel(line string) {
	tunnel := strings.Split(line, "-")

	// If the tunnel already exists, don't add it
	if Tunnels[line] {
		return
	}
	Tunnels[line] = true

	// Add the tunnel connection to the corresponding rooms
	room1 := GetRoomByName(tunnel[0])
	room2 := GetRoomByName(tunnel[1])
	if room1 != nil && room2 != nil {
		// Add room2 as a connection to room1
		room1.AddConnection(room2)
		// Add room1 as a connection to room2
		room2.AddConnection(room1)
	} else {
		log.Fatal("Error: Invalid room in tunnel")
	}
}

// GetRoomByName retrieves a room from ColonyRooms based on its name
func GetRoomByName(name string) *Room {
	for i := range ColonyRooms {
		if ColonyRooms[i].Name == name {
			return &ColonyRooms[i]
		}
	}
	return nil
}

// AddConnection adds a room as a connection to the current room
func (r *Room) AddConnection(room *Room) {
	r.Connections = append(r.Connections, room)
}