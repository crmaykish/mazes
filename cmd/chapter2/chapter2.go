package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"

	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 10
const height = 4

func main() {
	fmt.Println("Chapter 2: Binary Tree Algorithm")

	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	var g = grid.GridInit(width, height)

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			var currentCell = grid.CellAt(g, i, j)
			var northNeighbor = grid.CellAt(g, i, j).North
			var eastNeighbor = grid.CellAt(g, i, j).East

			if r.Intn(2) == 0 {
				if northNeighbor != nil {
					cell.CellLink(currentCell, northNeighbor)
				}
			} else {
				if eastNeighbor != nil {
					cell.CellLink(currentCell, eastNeighbor)
				}
			}
		}
	}

	grid.Print(g)
}
