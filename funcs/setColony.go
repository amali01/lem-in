package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SetColony(filePath string) {
	input, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var line string

	// Read the first line containing the number of ants
	if scanner.Scan() {
		line = scanner.Text()
		if NumOfAnts, err := strconv.Atoi(line); err == nil {
			fmt.Println("Number of ants:", NumOfAnts)
		} else {
			log.Fatal("Invalid number of ants")
		}
	}

	// Process the lines until a non-room line or non-comment li	ne is encountered
	var isEnd, isStart bool
	for scanner.Scan() {
		line = scanner.Text()
		if !RoomLineFormat(line) && !strings.HasPrefix(line, "#") {
			break
		}
		if RoomLineFormat(line) {
			AddRoom(line,isStart,isEnd)
			isStart = false
			isEnd = false
		} else if line == "##start" {
			isStart = true
		} else if line == "##end" {
			isEnd = true
		}
	}
	ColonyRooms = append(ColonyRooms, EndRoom)

	// Check and process the tunnel lines
	if TunnelLineFormat(line) {
		AddTunnel(line)
	}

	// Process the remaining tunnel lines
	for scanner.Scan() {
		line = scanner.Text()
		if TunnelLineFormat(line) {
			AddTunnel(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// RoomLineFormat checks if a line matches the format for a room line
func RoomLineFormat(line string) bool {
	roomMatch, _ := regexp.MatchString(`^\S+\s\d+\s\d+$`, line)
	return roomMatch
}

// TunnelLineFormat checks if a line matches the format for a tunnel line
func TunnelLineFormat(line string) bool {
	tunnelMatch, _ := regexp.MatchString(`\w+-\w+`, line)
	return tunnelMatch
}
