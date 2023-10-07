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
	funcs.SetColony(os.Args[1])
	Printfile(os.Args[1])
	// fmt.Println("all rooms:", funcs.ColonyRooms)
	allPaths, pathnum := funcs.FindAllPaths(funcs.ColonyRooms)
	if !pathnum {
		log.Fatal("No Viable Path Detected")
	}
	start := funcs.ColonyRooms[0].Name
	ants := funcs.QuickestPath(allPaths, funcs.NumOfAnts)
	for i := 0; i < maxAntPath(ants); i++ { // lines
		for _, ant := range ants { // ants room in line
			if i < len(ant.AntPath) && ant.AntPath[i].Name != start {
				fmt.Print(ant.Name, "-", ant.AntPath[i].Name, " ")
			}
		}
		fmt.Println()
	}
}

func maxAntPath(ants []funcs.Ant) (max int) {
	for _, ant := range ants {
		if len(ant.AntPath) > max {
			max = len(ant.AntPath)
		}
	}
	return max
}

func Printfile(filePath string) {
	// Run the cat command with the -e option
	cmd := exec.Command("cat", filePath)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))
}
