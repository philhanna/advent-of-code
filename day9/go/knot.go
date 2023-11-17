package main

type Knot struct {
	name  string
	point *Point
}

// NewKnot creates a new Knot with the specified name and location
func NewKnot(name string, point *Point) *Knot {
	p := new(Knot)
	p.name = name
	p.point = point
	return p
}

// Move takes the knot in the specified direction
func (k *Knot) Move(direction string) {
	switch direction {
	case "U":
		k.point.row--
	case "D":
		k.point.row++
	case "L":
		k.point.col--
	case "R":
		k.point.col++
	}
}

// Follow adjusts the position of this knot to be still touching the
// specified other knot
func (k *Knot) Follow(head *Knot, direction string) {
	stillTouches := k.point.Touches(head.point)
	if !stillTouches {
		k.Move(direction)
		switch direction {
		case "U", "D":
			k.point.col = head.point.col
		case "L", "R":
			k.point.row = head.point.row
		}
	}
}
