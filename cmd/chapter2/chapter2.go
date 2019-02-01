package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 20
const height = 8

func main() {
	fmt.Println("Chapter 2: Binary Tree Algorithm")

	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	var g = grid.GridInit(width, height)

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			var currentCell = grid.CellAt(g, x, y)

			var neighbors []*cell.Cell

			var northNeighbor = grid.CellAt(g, x, y).North
			var eastNeighbor = grid.CellAt(g, x, y).East

			if northNeighbor != nil {
				neighbors = append(neighbors, northNeighbor)
			}

			if eastNeighbor != nil {
				neighbors = append(neighbors, eastNeighbor)
			}

			if len(neighbors) > 0 {
				var randomNeighbor = neighbors[r.Intn(len(neighbors))]

				if randomNeighbor != nil {
					cell.CellLink(currentCell, randomNeighbor)
				}
			}
		}
	}

	grid.Print(g)
}
