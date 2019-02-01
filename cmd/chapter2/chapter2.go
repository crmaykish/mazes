package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 3
const height = 3

func main() {
	fmt.Println("Chapter 2: Binary Tree Algorithm")

	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	var g = grid.GridInit(width, height)

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			var currentCell = grid.CellAt(g, i, j)

			var neighbors []*cell.Cell

			var northNeighbor = grid.CellAt(g, j, j).North
			var eastNeighbor = grid.CellAt(g, i, j).East

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
