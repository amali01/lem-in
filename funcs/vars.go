package funcs

var (
	ColonyRooms []Room
	EndRoom     Room
	isEnd       bool
	isStart     bool
	Tunnels     = make(map[string]bool)
)

type Room struct {
	Name    string
	Coord_x int
	Coord_y int
	Connections []*Room
}