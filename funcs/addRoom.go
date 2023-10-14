package funcs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// AddRoom adds a room to the colony
func AddRoom(line string, isStart bool, isEnd bool) {
	room, err := MakeRoom(line)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check for duplicate room names
	for _, chroom := range ColonyRooms {
		if chroom.Name == room.Name {
			log.Fatal("Duplicate room name detected")
		}
	}

	// Add the room to the appropriate list based on its properties
	switch {
	case isStart:
		// Prepend the room to the ColonyRooms list if it is the start room
		ColonyRooms = append([]Room{room}, ColonyRooms...)
	case isEnd:
		// Set the room as the end room
		EndRoom = room
	default:
		// Append the room to the ColonyRooms list if it is a regular room
		ColonyRooms = append(ColonyRooms, room)
	}
}

// MakeRoom creates a new room based on the provided line of data
func MakeRoom(line string) (Room, error) {
	data := strings.Split(line, " ")

	// Validate room name
	// A room will never start with the letter L or with # and must have no spaces.
	if strings.HasPrefix(data[0], "#") || strings.HasPrefix(data[0], "L") || strings.Contains(data[0], " ") {
		return Room{}, fmt.Errorf("error: invalid room Name")
	}

	// Parse room coordinates
	coX, err1 := strconv.Atoi(data[1])
	coY, err2 := strconv.Atoi(data[2])
	if err1 != nil || err2 != nil {
		fmt.Println("error: invalid room coordinates")
		os.Exit(1)
	}

	// Create and return the room
	return Room{Name: data[0], Coord_x: coX, Coord_y: coY}, nil
}
