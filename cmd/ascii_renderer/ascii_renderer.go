package main

import (
	"github.com/crmaykish/mazes/pkg/algorithm"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 20
const height = 8

func main() {
	var g = grid.GridInit(width, height)

	algorithm.SideWinder(g)

	grid.Print(g)
}
