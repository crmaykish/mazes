package cell

type Cell struct {
	X     int
	Y     int
	links map[string]*Cell
	North *Cell
	South *Cell
	East  *Cell
	West  *Cell
}

func CellInit(x, y int) Cell {
	var cell Cell

	cell.X = x
	cell.Y = y
	cell.links = make(map[string]*Cell)

	return cell
}

func CellLink(cell1, cell2 *Cell) {
	var key1 = key(cell1)
	var key2 = key(cell2)

	cell1.links[key2] = cell2
	cell2.links[key1] = cell1
}

func CellUnlink(cell1, cell2 *Cell) {
	var key1 = key(cell1)
	var key2 = key(cell2)

	delete(cell1.links, key2)
	delete(cell2.links, key1)
}

func CellLinks(cell *Cell) {
}

func CellsLinked(cell1, cell2 *Cell) bool {
	if cell2 == nil {
		return false
	}

	var key2 = key(cell2)

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

func key(cell *Cell) string {
	return string(cell.X) + string(cell.Y)
}
