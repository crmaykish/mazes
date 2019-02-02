package main

import (
	"fmt"

	"github.com/crmaykish/mazes/pkg/cell"

	"github.com/crmaykish/mazes/pkg/algorithm"
	"github.com/crmaykish/mazes/pkg/grid"
)

const width = 12
const height = 7

const cellBody = "   "

func main() {
	var g = grid.GridInit(width, height)
	algorithm.SideWinder(g)

	var output string

	// Create top border
	var topBorder = "┏"

	for i := 0; i < g.Width; i++ {
		// TODO: this will only work if the top row is totally open like in Binary Tree
		// If the top row can have walls, this will have to add in T box characters
		// as appropriate
		if i > 0 {
			topBorder += "━━━━"
		} else {
			topBorder += "━━━"
		}
	}

	topBorder += "┓\n"

	output += topBorder

	// Loop through each row
	for y := g.Height - 1; y >= 0; y-- {
		var midLine, bottomLine string

		midLine = "┃"

		// Figure out the bottom left corner character for this row
		if y == 0 {
			// Bottom row will always show a corner character
			bottomLine = "┗"
		} else if cell.CellsLinked(grid.CellAt(g, 0, y), grid.CellAt(g, 0, y).South) {
			// If this edge cell is linked to the one below, use a straight character
			bottomLine = "┃"
		} else {
			// If there is no link, use the T character
			bottomLine = "┣"
		}

		// Loop through each cell in the row
		for x := 0; x < g.Width; x++ {
			// Always add the cell body to the middle line of this row
			midLine += cellBody

			var currentCell = grid.CellAt(g, x, y)

			if cell.CellsLinked(currentCell, currentCell.East) {
				// If this cell is linked the cell east of it, open the passage
				midLine += " "
			} else {
				// If there is no link, draw a vertical wall
				midLine += "┃"
			}

			if cell.CellsLinked(currentCell, currentCell.South) {
				bottomLine += cellBody
			} else {
				bottomLine += "━━━"
			}

			// Figure out which corner piece to render
			// ╋  ┻  ┳  ┫  ┣  ┛  ┗  ┓  ┏

			// If cell and cell.east are not linked
			// AND cell and cell.south are not linked
			// AND cell.east and cell.east.south are not linked
			// AND cell.south and cell.south.east are not linked
			if currentCell.South != nil && currentCell.East != nil &&
				!cell.CellsLinked(currentCell, currentCell.East) &&
				!cell.CellsLinked(currentCell, currentCell.South) &&
				!cell.CellsLinked(currentCell.East, currentCell.East.South) &&
				!cell.CellsLinked(currentCell.South, currentCell.South.East) {
				bottomLine += "╋"
			} else if currentCell.East != nil &&
				!cell.CellsLinked(currentCell, currentCell.East) &&
				!cell.CellsLinked(currentCell, currentCell.South) &&
				!cell.CellsLinked(currentCell.East, currentCell.East.South) &&
				(currentCell.South == nil || cell.CellsLinked(currentCell.South, currentCell.South.East)) {
				bottomLine += "┻"
			} else if currentCell.South != nil && currentCell.East != nil &&
				cell.CellsLinked(currentCell, currentCell.East) &&
				!cell.CellsLinked(currentCell, currentCell.South) &&
				!cell.CellsLinked(currentCell.East, currentCell.East.South) &&
				!cell.CellsLinked(currentCell.South, currentCell.South.East) {
				bottomLine += "┳"
			} else if currentCell.South != nil &&
				!cell.CellsLinked(currentCell, currentCell.East) &&
				!cell.CellsLinked(currentCell, currentCell.South) &&
				!cell.CellsLinked(currentCell.South, currentCell.South.East) &&
				(currentCell.East == nil || cell.CellsLinked(currentCell.East, currentCell.East.South)) {
				bottomLine += "┫"
			} else if currentCell.South != nil && currentCell.East != nil &&
				!cell.CellsLinked(currentCell, currentCell.East) &&
				cell.CellsLinked(currentCell, currentCell.South) &&
				!cell.CellsLinked(currentCell.South, currentCell.South.East) &&
				!cell.CellsLinked(currentCell.East, currentCell.East.South) {
				bottomLine += "┣"
			} else {
				bottomLine += "+"
			}
		}

		// Write the middle and bottom lines of the row
		output += (midLine + "\n")
		output += (bottomLine + "\n")
	}

	fmt.Println(output)

}
