package main

import (
	"math"

	"github.com/crmaykish/mazes/pkg/algorithm"
	"github.com/crmaykish/mazes/pkg/cell"
	"github.com/crmaykish/mazes/pkg/grid"
	"github.com/fogleman/gg"
)

// Maze settings
const width = 20
const height = 14
const cellSize = 16
const wallThickness = 1
const outsidePadding = cellSize

// Colors
const wallR, wallG, wallB = 0x00, 0x00, 0x00
const bgR, bgG, bgB = 0xFF, 0xFF, 0xFF

var outputFileName = "maze.png"

func main() {
	var g = grid.GridInit(width, height)

	algorithm.SideWinder(g)

	var imageWidth = (width * cellSize) + (2 * outsidePadding) + wallThickness
	var imageHeight = (height * cellSize) + (2 * outsidePadding) + wallThickness

	dc := gg.NewContext(imageWidth, imageHeight)

	// Fill the background
	dc.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
	dc.SetRGB255(bgR, bgG, bgB)
	dc.Fill()

	var length = float64(cellSize)
	var thickness = float64(wallThickness)

	// Draw left border
	dc.DrawRectangle(outsidePadding, outsidePadding, thickness, height*cellSize)

	// Draw right border
	dc.DrawRectangle(outsidePadding, outsidePadding, width*cellSize, thickness)

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			var maxY = (height * cellSize)

			var c = grid.CellAt(g, x, y)

			if !cell.CellsLinked(c, c.South) {
				var southX = float64((x * cellSize) + outsidePadding)
				var southY = float64((maxY - y*cellSize) + outsidePadding)

				dc.DrawRectangle(southX, southY, length+wallThickness, thickness)
			}

			if !cell.CellsLinked(c, c.East) {
				var eastX = float64(((x + 1) * cellSize) + outsidePadding)
				var eastY = float64((maxY - (y+1)*cellSize) + outsidePadding)

				dc.DrawRectangle(eastX, eastY, thickness, length)
			}
		}
	}

	dc.SetRGB255(wallR, wallG, wallB)
	dc.Fill()

	// Add entry and exit
	var entryX float64 = outsidePadding
	var entryY = (math.Floor(height/2) * cellSize) + outsidePadding + thickness
	var exitX float64 = outsidePadding + (width * cellSize)
	var exitY = entryY

	dc.DrawRectangle(entryX, entryY, thickness, cellSize-thickness)
	dc.DrawRectangle(exitX, exitY, thickness, cellSize-thickness)

	dc.SetRGB255(bgR, bgG, bgB)
	dc.Fill()

	// Output image file
	dc.SavePNG(outputFileName)
}
