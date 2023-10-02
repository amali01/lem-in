package funcs

func FindAllPaths(AllRooms []Room) ([]Path, bool) {
	startRoom := &AllRooms[0]
	endRoom := &AllRooms[len(AllRooms)-1]

	if startRoom == nil || endRoom == nil {
		return nil, false
	}

	var allPaths []Path
	currentPath := make(Path, 0)
	visited := make(map[string]bool)
	findPaths(*startRoom, *endRoom, currentPath, visited, &allPaths)
	return allPaths, len(allPaths) != 0
}


func findPaths(currentRoom, endRoom Room, currentPath Path, visited map[string]bool, allPaths *[]Path) {
	// Mark the current room as visited
	visited[currentRoom.Name] = true
	// Add the current room to the current path
	currentPath = append(currentPath, currentRoom)
	// Check if the current room is the end room
	if currentRoom.Name == endRoom.Name {
		// Add the current path to the list of all paths
		// Create a copy of the current path and add it to the list of all paths
		copiedPath := make(Path, len(currentPath))
		copy(copiedPath, currentPath)
		*allPaths = append(*allPaths, copiedPath)
	}
	// Explore the neighboring rooms recursively
	for _, neighbor := range currentRoom.Connections {
		if !visited[neighbor.Name] {
			findPaths(*neighbor, endRoom, currentPath, visited, allPaths)
		}
	}

	// Backtrack: Remove the current room from the current path and mark it as unvisited
	currentPath = currentPath[:len(currentPath)-1]
	visited[currentRoom.Name] = false
}
