package algorithm

import (
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
)

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

func SideWinder(g *grid.Grid) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	// Loop through the grid by rows
	for y := 0; y < g.Height; y++ {
		var run []*cell.Cell
		// Loop through each cell in the row
		for x := 0; x < g.Width; x++ {
			var currentCell = grid.CellAt(g, x, y)

			run = append(run, currentCell)

			var atEasternBoundary bool
			var atNorthernBoundary bool

			if currentCell.East == nil {
				atEasternBoundary = true
			}

			if currentCell.North == nil {
				atNorthernBoundary = true
			}

			var shouldCloseOut = atEasternBoundary || !atNorthernBoundary && r.Intn(2) == 0

			if shouldCloseOut {
				var member = run[r.Intn(len(run))]

				if member.North != nil {
					cell.CellLink(member, member.North)
				}

				// Clear the run of cells
				run = nil
			} else {
				cell.CellLink(currentCell, currentCell.East)
			}
		}
	}
}
