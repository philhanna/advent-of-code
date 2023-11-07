package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMove(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		that    Move
		wantErr bool
	}{
		{"simple", "move 1 from 3 to 5", Move{1, 3, 5}, false},
		{">9 count", "move 11 from 1 to 9", Move{11, 1, 9}, false},
		{"empty", "", Move{}, true},
		{"bad format", "move 1 from 1", Move{}, true},
		{"bogus", "bogus", Move{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := ParseMove(tt.line)
			if tt.wantErr == false {
				assert.Nil(t, err)
				assert.True(t, this.Equal(tt.that))
			} else {
				assert.NotNil(t, err)				
			}
		})
	}
}
