package main

import (
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 20
const height = 8

func main() {
	var g = grid.GridInit(width, height)

	BinaryTree(g)

	grid.Print(g)
}

func BinaryTree(g *grid.Grid) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

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
}

func SideWinder(g grid.Grid) {

}
