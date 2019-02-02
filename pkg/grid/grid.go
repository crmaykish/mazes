package grid

import (
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
)

type Grid struct {
	Width  int
	Height int
	cells  [][]cell.Cell
}

func GridInit(w, h int) *Grid {
	if w <= 0 || h <= 0 {
		return nil
	}

	var grid = new(Grid)

	grid.Width = w
	grid.Height = h

	prepare(grid)
	configureCells(grid)

	return grid
}

func CellAt(grid *Grid, x, y int) *cell.Cell {
	return &grid.cells[x][y]
}

func RandomCell(grid *Grid) *cell.Cell {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	return &grid.cells[r.Intn(grid.Width)][r.Intn(grid.Height)]
}

func Size(grid *Grid) int {
	return grid.Width * grid.Height
}

func prepare(grid *Grid) {
	// Create the outer slices
	grid.cells = make([][]cell.Cell, grid.Width)

	for x := 0; x < grid.Width; x++ {
		// Create the inner slices
		grid.cells[x] = make([]cell.Cell, grid.Height)

		for y := 0; y < grid.Height; y++ {
			// populate each cell in the grid
			grid.cells[x][y] = cell.CellInit(x, y)
		}
	}
}

func configureCells(grid *Grid) {
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			var c = CellAt(grid, x, y)

			if y < grid.Height-1 {
				c.North = CellAt(grid, x, y+1)
			}

			if y > 0 {
				c.South = CellAt(grid, x, y-1)
			}

			if x < grid.Width-1 {
				c.East = CellAt(grid, x+1, y)
			} else {
			}

			if x > 0 {
				c.West = CellAt(grid, x-1, y)
			}
		}
	}
}
