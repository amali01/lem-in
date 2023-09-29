package funcs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func SetColony(filePath string) {
	input, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lineType := CheckLineRoomFormat(line)
		switch lineType {
		case "room":
			AddRoom(line)
		case "tunnel":
			AddTunnel(line)
			// default:
			// 	fmt.Println("Unknown line type:", lineType)
		}
		if scanner.Text() == "##start" {
			isStart = true
		} else if scanner.Text() == "##end" {
			isEnd = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// add the end room
	ColonyRooms = append(ColonyRooms, EndRoom)
	fmt.Println(ColonyRooms)
}

func CheckLineRoomFormat(line string) string {
	if roomMatch, _ := regexp.MatchString(`^\S+\s\d+\s\d+$`, line); roomMatch {
		return "room"
	} else if tunnel, _ := regexp.MatchString(`\w+-\w+`, line); tunnel {
		return "tunnel"
	}
	return "wrong format"
}
