package funcs

type Path []Room

func FindAllPaths(AllRooms []Room) []Path {
	startRoom := AllRooms[0]
	endRoom := AllRooms[len(AllRooms)-1]

	// if startRoom == nil || endRoom == nil {
	// 	return nil
	// }

	var allPaths []Path
	currentPath := make(Path, 0)
	visited := make(map[string]bool)
	findPathsDFS(startRoom, endRoom, currentPath, visited, &allPaths)

	return allPaths
}

func findPathsDFS(currentRoom, endRoom Room, currentPath Path, visited map[string]bool, allPaths *[]Path) {
	// Mark the current room as visited
	visited[currentRoom.Name] = true
	// Add the current room to the current path
	currentPath = append(currentPath, currentRoom)
	// Check if the current room is the end room
	if currentRoom.Name == endRoom.Name {
		// Add the current path to the list of all paths
		*allPaths = append(*allPaths, currentPath)
	}

	// Explore the neighboring rooms recursively
	for _, neighbor := range currentRoom.Connections {
		if !visited[neighbor.Name] {
			findPathsDFS(*neighbor, endRoom, currentPath, visited, allPaths)
		}
	}

	// Backtrack: Remove the current room from the current path and mark it as unvisited
	currentPath = currentPath[:len(currentPath)-1]
	visited[currentRoom.Name] = false
}
