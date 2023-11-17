package main

type Point struct {
	row int
	col int
}

// NewPoint creates a new Point with the specified row and column.
func NewPoint(row, col int) *Point {
	return &Point{row, col}
}

// Touches returns true if the square of the Euclidean distance between
// this point and that point is less than or equal to 2.
func (this *Point) Touches(that *Point) bool {
	rowdif := this.row - that.row
	coldif := this.col - that.col
	return rowdif*rowdif+coldif*coldif <= 2
}
