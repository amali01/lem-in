package funcs

import (
	"fmt"
	"strings"
)

func AddTunnel(line string) {
	tunnel := strings.Split(line, "-")
	// if the tunnel already exists don't add it
	if Tunnels[line] {
		return
	}
	Tunnels[line] = true

	// Add the tunnel connection to the corresponding rooms
	room1 := GetRoomByName(tunnel[0])
	room2 := GetRoomByName(tunnel[1])
	if room1 != nil && room2 != nil {
		room1.AddConnection(room2)
	} else {
		fmt.Println("Error: Invalid room in tunnel connection")
	}
}

func GetRoomByName(name string) *Room {
	for i := range ColonyRooms {
		if ColonyRooms[i].Name == name {
			return &ColonyRooms[i]
		}
	}
	return nil
}

func (r *Room) AddConnection(room *Room) {
	r.Connections = append(r.Connections, room)
}
