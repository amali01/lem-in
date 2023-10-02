package funcs

var (
	ColonyRooms []Room
	NumOfAnts   int
	EndRoom     Room
	Tunnels     = make(map[string]bool)
)

type Room struct {
	Name        string
	Coord_x     int
	Coord_y     int
	Connections []*Room
}

type Path []Room

type Ant struct {
	Name string
	AntPath Path
}