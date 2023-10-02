package funcs

import (
	"sort"
	"strconv"
)

func QuickestPath(allPaths []Path, numOfAnts int) []Ant {
	Ants := createAnts(numOfAnts)
	SortPaths(allPaths)
	for i, antIdx := 0, 0; i < len(allPaths) && antIdx < len(Ants); i, antIdx = i+1, antIdx+1 {
		if CheckConflict(Ants, allPaths[i]) {
			antIdx--
		} else {
			Ants[antIdx].AntPath = allPaths[i]
		}
		// wait in start room if we need to
		if len(allPaths) == i+1 || (i+1 < len(allPaths) && len(allPaths[i+1])-len(allPaths[i]) > 1) {
			CrossGate(allPaths, i+1)
			i = -1
		}
	}
	return Ants
}

// CheckConflict would check if new path have a room that have been used at the same time (position) and same thing with the tunnels.
func CheckConflict(ants []Ant, path Path) bool {
	for _, ant := range ants {
		for i, room := range path {
			// Check if the room is already occupied by another ant at the same time
			if i < len(ant.AntPath) && room.Name == ant.AntPath[i].Name && !isEndRoom(room, path) {
				return true
			}
			// check if the tunnel is used at the same time 
			// if i+1 < len(ant.AntPath) && room.Name == ant.AntPath[i+1].Name {
			// 	return true
			// }
		}
	}
	return false // No conflicts found
}

func isEndRoom(room Room, path Path) bool {
	start := path[0]
	end := path[len(path)-1]
	return room.Name == start.Name || room.Name == end.Name
}

func CrossGate(allPaths []Path, i int) {
	for pathIdx := range allPaths {
		if pathIdx == i {
			break
		}
		// add the first element again at the beginning
		firstRoom := allPaths[pathIdx][0]
		allPaths[pathIdx] = append([]Room{firstRoom}, allPaths[pathIdx]...)
	}
}

func createAnts(numOfAnts int) []Ant {
	Ants := make([]Ant, numOfAnts)
	for i := range Ants {
		Ants[i].Name = "L" + strconv.Itoa(i+1)
	}
	return Ants
}

func SortPaths(allPaths []Path) {
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})
}
