package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"lem-in/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage")
		return
	}

	// Set the colony based on the command-line argument
	funcs.SetColony(os.Args[1])

	// Find all paths in the colony
	allPaths, pathnum := funcs.FindAllPaths(funcs.ColonyRooms)

	// Check if there are any valid paths
	if !pathnum {
		log.Fatal("UnValid Path Detected")
	}

	// Print the contents of the input file
	Printfile(os.Args[1])

	start := funcs.ColonyRooms[0].Name

	// Find the quickest path for each ant
	ants := funcs.QuickestPath(allPaths, funcs.NumOfAnts)

	numOfLines := 0

	// Print the ant paths
	for i := 0; i < maxAntPath(ants); i++ { // lines
		for _, ant := range ants { // ants room in line
			if i < len(ant.AntPath) && ant.AntPath[i].Name != start {
				fmt.Print(ant.Name, "-", ant.AntPath[i].Name, " ")
			}
		}
		fmt.Println()
		numOfLines++
	}

	fmt.Println(numOfLines)
}

// maxAntPath returns the maximum length of ant paths
func maxAntPath(ants []funcs.Ant) (max int) {
	for _, ant := range ants {
		if len(ant.AntPath) > max {
			max = len(ant.AntPath)
		}
	}
	return max
}

// Printfile prints the contents of a file specified by the filePath
func Printfile(filePath string) {
	// Run the cat command with the -e option
	cmd := exec.Command("cat", filePath)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))
}
