package funcs

import (
	"os"
	"strconv"
	"strings"
)

func AddRoom(line string) {
	room := MakeRoom(line)
	switch {
	case isStart:
		ColonyRooms = append([]Room{room}, ColonyRooms...)
		isStart = false
	case isEnd:
		EndRoom = room
		isEnd = false
	default:
		ColonyRooms = append(ColonyRooms, room)
	}
}

func MakeRoom(line string) Room {
	data := strings.Split(line, " ")
	coX, err1 := strconv.Atoi(data[1])
	coY, err2 := strconv.Atoi(data[2])
	if err1 != nil || err2 != nil {
		os.Exit(1)
	}
	return Room{Name: data[0], Coord_x: coX, Coord_y: coY}
}
