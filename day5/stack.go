package main

type Stack struct {
	Crates string
}

func NewStack(crates string) *Stack {
	ps := new(Stack)
	ps.Crates = crates
	return ps
}

// Returns true if the stack is empty
func (ps *Stack) Empty() bool {
	return ps.Length() <= 0
}

// Returns the length of the stack
func (ps *Stack) Length() int {
	return len(ps.Crates)
}

// Peeks the top element without popping it
func (ps *Stack) Peek() rune {
	if ps.Empty() {
		panic("Stack empty")
	}
	n := ps.Length()
	topCrate := rune(ps.Crates[n-1])
	return topCrate
}

// Pops a crate
func (ps *Stack) Pop() rune {
	if ps.Empty() {
		panic("Stack empty")
	}
	n := ps.Length()
	topCrate := rune(ps.Crates[n-1])
	ps.Crates = ps.Crates[:n-1]
	return topCrate
}

// Pushes a crate
func (ps *Stack) Push(crate rune) {
	ps.Crates += string(crate)
}
