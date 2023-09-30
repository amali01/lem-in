package main

import (
	"fmt"
	"os"
	"lem-in/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage")
		return
	}
	funcs.SetColony(os.Args[1])
	fmt.Println(funcs.ColonyRooms)
}
