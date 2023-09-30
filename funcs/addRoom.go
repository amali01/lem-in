package funcs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddRoom(line string, isStart bool, isEnd bool) {
	room, err := MakeRoom(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch {
	case isStart:
		ColonyRooms = append([]Room{room}, ColonyRooms...)
	case isEnd:
		EndRoom = room
	default:
		ColonyRooms = append(ColonyRooms, room)
	}
}

func MakeRoom(line string) (Room, error) {
	data := strings.Split(line, " ")
	// A room will never start with the letter L or with # and must have no spaces.
	if strings.HasPrefix(data[0], "#") || strings.HasPrefix(data[0], "L") || strings.Contains(data[0], " ") {
		return Room{}, fmt.Errorf("Error: Invalid room Name")
	}
	coX, err1 := strconv.Atoi(data[1])
	coY, err2 := strconv.Atoi(data[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid room coordinates")
		os.Exit(1)
	}
	return Room{Name: data[0], Coord_x: coX, Coord_y: coY}, nil
}
