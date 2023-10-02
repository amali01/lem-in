package funcs

import (
	"fmt"
	"sort"
	"strconv"
)

func QuickestPath(allPaths []Path, numOfAnts int) []Ant {
	Ants := createAnts(numOfAnts)
	fmt.Println(Ants, numOfAnts)
	// pathLength := 1
	// for _, path := range allPaths {
	// 	if len(path) == pathLength {

	// 	}
	// }
	SortPaths(allPaths)
	// prevlen := -1
	// for i := 0; i < len(allPaths); i++ {
	// 	for antIdx := range Ants {

	for i, antIdx := 0, 0; i < len(allPaths) && antIdx < len(Ants); i, antIdx = i+1, antIdx+1 {
		// if i != 0 {
		// 	prevlen = len(allPaths[i-1])
		// }
		if i == 0 {
			Ants[antIdx].AntPath = allPaths[i]
			// prevlen = len(allPaths[i])
			if len(allPaths) == 1 {
				WaitPaths(allPaths, 1)
				i = -1
			}
		// } else if prevlen != -1 && len(allPaths[i])-prevlen > 1 {
		// 	Ants[antIdx].AntPath = allPaths[i]
		// 	WaitPaths(allPaths, i+1)
		// 	// prevlen = len(allPaths[i])
		// 	i = -1
		} else {
			Ants[antIdx].AntPath = allPaths[i]
			// prevlen = len(allPaths[i])
			if len(allPaths) == i+1 || (i+1 < len(allPaths) && len(allPaths[i+1])-len(allPaths[i]) > 1) {
				WaitPaths(allPaths, i+1)
				i = -1
			}
		}
	}
	return Ants
}

func WaitPaths(allPaths []Path, i int) {
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

// func maxPath(paths []Path) int {
// 	max := 0
// 	for _, path := range paths {
// 		if len(path) > max {
// 			max = len(path)
// 		}
// 	}
// 	return max
// }

// func minPath(paths []Path) int {
// 	min := len(paths[0])
// 	for _, path := range paths {
// 		if len(path) < min {
// 			min = len(path)
// 		}
// 	}
// 	return min
// }
