package funcs

import (
	"fmt"
	"strings"
)

func AddTunnel(line string) {
	tunnel := strings.Split(line, "-")
	Tunnels[line] = true
	fmt.Println(line)

	// Add the tunnel connection to the corresponding rooms
	room1 := GetRoomByName(tunnel[0])
	room2 := GetRoomByName(tunnel[1])
	fmt.Println(room1,room2)
	if room1 != nil && room2 != nil {
		room1.AddConnection(room2)
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
