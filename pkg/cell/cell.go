package cell

// TODO: maybe build a map with the key being the coords for the Linked Cells?
// would save a lot of trouble messing around with slices

type Cell struct {
	Row    int
	Column int
	links  map[string]*Cell
	North  *Cell
	South  *Cell
	East   *Cell
	West   *Cell
}

func CellInit(row, column int) Cell {
	var cell Cell

	cell.Row = row
	cell.Column = column
	cell.links = make(map[string]*Cell)

	return cell
}

func CellLink(cell1, cell2 *Cell) {
	var key1 = string(cell1.Row) + string(cell1.Column)
	var key2 = string(cell2.Row) + string(cell2.Column)

	cell1.links[key2] = cell2
	cell2.links[key1] = cell1
}

func CellUnlink(cell1, cell2 *Cell) {
	var key1 = string(cell1.Row) + string(cell1.Column)
	var key2 = string(cell2.Row) + string(cell2.Column)

	delete(cell1.links, key2)
	delete(cell2.links, key1)
}

func CellLinks(cell *Cell) {
}

func CellsLinked(cell1, cell2 *Cell) bool {
	if cell2 == nil {
		return false
	}

	var key2 = string(cell2.Row) + string(cell2.Column)

	return cell1.links[key2] != nil
}

func CellNeighbors(cell *Cell) []Cell {
	var neighbors []Cell

	if cell.North != nil {
		neighbors = append(neighbors, *cell.North)
	}

	if cell.South != nil {
		neighbors = append(neighbors, *cell.South)
	}

	if cell.East != nil {
		neighbors = append(neighbors, *cell.East)
	}

	if cell.West != nil {
		neighbors = append(neighbors, *cell.West)
	}

	return neighbors
}
