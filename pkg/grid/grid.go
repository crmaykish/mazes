package grid

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crmaykish/mazes/pkg/cell"
)

type Grid struct {
	Rows    int
	Columns int
	Cells   [][]cell.Cell
}

func GridInit(r, c int) *Grid {
	if r <= 0 || c <= 0 {
		return nil
	}

	var grid = new(Grid)

	grid.Rows = r
	grid.Columns = c

	prepare(grid)
	configureCells(grid)

	return grid
}

func CellAt(grid *Grid, row, column int) *cell.Cell {
	return &grid.Cells[row][column]
}

func RandomCell(grid *Grid) *cell.Cell {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	return &grid.Cells[r.Intn(grid.Rows)][r.Intn(grid.Columns)]
}

func Size(grid *Grid) int {
	return grid.Rows * grid.Columns
}

func Print(grid *Grid) {
	var output = "+"

	for i := 0; i < grid.Columns; i++ {
		output += "---+"
	}

	output += "\n"

	for i := 0; i < grid.Rows; i++ {
		var top = "|"
		var bottom = "+"

		for j := 0; j < grid.Columns; j++ {
			var currentCell = CellAt(grid, i, j)

			var body = "   "
			var eastBoundary string
			var southBoundary string

			if cell.CellsLinked(currentCell, currentCell.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}

			top += (body + eastBoundary)

			if cell.CellsLinked(CellAt(grid, i, j), CellAt(grid, i, j).South) {
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

func prepare(grid *Grid) {
	// create the outer slice
	grid.Cells = make([][]cell.Cell, grid.Rows)

	for i := 0; i < grid.Rows; i++ {
		// great each inner slice
		grid.Cells[i] = make([]cell.Cell, grid.Columns)

		for j := 0; j < grid.Columns; j++ {
			// populate each cell in the grid
			grid.Cells[i][j] = cell.CellInit(i, j)
		}
	}
}

func configureCells(grid *Grid) {
	for i := 0; i < grid.Rows; i++ {
		for j := 0; j < grid.Columns; j++ {

			if i > 0 {
				grid.Cells[i][j].North = &grid.Cells[i-1][j]
			}

			if i < grid.Rows-1 {
				grid.Cells[i][j].South = &grid.Cells[i+1][j]
			}

			if j > 0 {
				grid.Cells[i][j].West = &grid.Cells[i][j-1]
			}
			if j < grid.Columns-1 {
				grid.Cells[i][j].East = &grid.Cells[i][j+1]
			}
		}
	}
}
