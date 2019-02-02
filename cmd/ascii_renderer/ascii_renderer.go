package main

import (
	"fmt"

	"github.com/crmaykish/mazes/pkg/algorithm"
	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 8
const height = 4

func main() {
	var g = grid.GridInit(width, height)

	algorithm.SideWinder(g)

	var output = "+"

	for i := 0; i < g.Width; i++ {
		output += "---+"
	}

	output += "\n"

	for y := g.Height - 1; y >= 0; y-- {
		var top = "|"
		var bottom = "+"

		for x := 0; x < g.Width; x++ {
			var currentCell = grid.CellAt(g, x, y)

			var body = "   "
			var eastBoundary string
			var southBoundary string

			if cell.CellsLinked(currentCell, currentCell.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}

			top += (body + eastBoundary)

			if cell.CellsLinked(currentCell, currentCell.South) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}

			var corner = "+"

			bottom += (southBoundary + corner)
		}

		output += (top + "\n")
		output += (bottom + "\n")
	}

	fmt.Println(output)
}
