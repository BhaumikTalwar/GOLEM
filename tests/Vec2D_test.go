package tests

import (
	m "golem"
	"testing"
)

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name string
		v1   m.Vec2D
		v2   m.Vec2D
		res  bool
	}{
		{"Equal Case", m.Vec2D{X: 1, Y: 0}, m.Vec2D{X: 1, Y: 0}, true},
		{"UnEqual Case", m.Vec2D{X: -26, Y: 75}, m.Vec2D{X: 13, Y: 56}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.v1.IsEqual(tt.v2) != tt.res {
				t.Errorf("Expected %v, Got %v", tt.v1, tt.v2)
			}
		})
	}
}
