package main

import "fmt"

type Vertex struct {
	Lat, Long int
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40, -74,
	},
	"Google": Vertex{
		37, -122,
	},
}

func main() {
	fmt.Println(m)
}
