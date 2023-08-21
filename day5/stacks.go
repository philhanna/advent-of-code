package main

import "fmt"

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Stacks struct {
	CrateStacks []Stack
}

func (ps *Stacks) ApplyMove(move Move) error {
	ns := len(ps.CrateStacks)

	// Convert to zero-based stack numbers after checking they are within bounds

	if move.FromStackIndex < 1 || move.FromStackIndex > ns {
		return fmt.Errorf("From stack number must be between %d and %d, not %d", 1, ns, move.FromStackIndex)
	}
	fx := move.FromStackIndex - 1

	if move.ToStackIndex < 1 || move.ToStackIndex > ns {
		return fmt.Errorf("To stack number must be between %d and %d, not %d", 1, ns, move.ToStackIndex)
	}
	tx := move.ToStackIndex - 1

	// Make the move

	for i := 0; i < move.Count; i++ {
		if ps.CrateStacks[fx].Empty() {
			return fmt.Errorf("From stack (%d) is empty", move.FromStackIndex)
		}
		crate := ps.CrateStacks[fx].Pop()
		ps.CrateStacks[tx].Push(crate)
	}

	// Everything ok
	
	return nil
}
