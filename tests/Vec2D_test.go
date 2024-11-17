package tests

import (
	m "golem"
	"math"
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

func TestIsNotEqual(t *testing.T) {
	tests := []struct {
		name string
		v1   m.Vec2D
		v2   m.Vec2D
		res  bool
	}{
		{"Equal Case", m.Vec2D{X: 1, Y: 0}, m.Vec2D{X: 1, Y: 0}, false},
		{"UnEqual Case", m.Vec2D{X: -26, Y: 75}, m.Vec2D{X: 13, Y: 56}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.v1.IsNotEqual(tt.v2) != tt.res {
				t.Errorf("Expected %v, Got %v", tt.v1, tt.v2)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name string
		v    m.Vec2D
		len  float64
	}{
		{"Case 1", m.Vec2D{X: 1, Y: 0}, 1},
		{"Case 2", m.Vec2D{X: 0, Y: 0}, 0},
		{"Case 3", m.Vec2D{X: -4, Y: 3}, 5},
		{"Case 4", m.Vec2D{X: 10, Y: 2}, math.Sqrt(104)},
		{"Case 5", m.Vec2D{X: -9.8, Y: -4}, math.Sqrt(112.040)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.v.Length()
			if l != tt.len {
				t.Errorf("Expected %v, Got %v", tt.len, l)
			}
		})
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		name string
		v1   m.Vec2D
		v2   m.Vec2D
		res  float64
	}{
		{"Case 1", m.Vec2D{X: 1, Y: 0}, m.Vec2D{X: 1, Y: 0}, 1},
		{"Case 2", m.Vec2D{X: 1, Y: 0}, m.Vec2D{X: 0, Y: 1}, 0},
		{"Case 3", m.Vec2D{X: 1, Y: 1}, m.Vec2D{X: 1, Y: 1}, 2},
		{"Case 4", m.Vec2D{X: 3, Y: 4}, m.Vec2D{X: 3, Y: 4}, 5},
		{"Case 5", m.Vec2D{X: 1, Y: 1}, m.Vec2D{X: -1, Y: -1}, -2},
		{"Case 6", m.Vec2D{X: 1, Y: 1}, m.Vec2D{X: 0.5, Y: 0.5}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dot := tt.v1.Dot(tt.v2)
			if dot != tt.res {
				t.Errorf("Expected %v, Got %v", tt.res, dot)
			}
		})
	}
}

func TestDist(t *testing.T) {
	tests := []struct {
		name string
		v1   m.Vec2D
		v2   m.Vec2D
		res  float64
	}{
		{"Equal Case", m.Vec2D{X: 1, Y: 0}, m.Vec2D{X: 1, Y: 0}, 0},
		{"Same Axis Case", m.Vec2D{X: 3, Y: 0}, m.Vec2D{X: 1, Y: 0}, 2},
		{"Right Tri Case", m.Vec2D{X: 3, Y: 4}, m.Vec2D{X: 0, Y: 0}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := tt.v1.Dist(tt.v2)
			d2 := tt.v2.Dist(tt.v1)

			if d1 != d2 && d1 != tt.res {
				t.Errorf("Expected %v, Got %v and %v", tt.res, d1, d2)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		v    m.Vec2D
		len  float64
	}{
		{"Case 1", m.Vec2D{X: 1, Y: 0}, 1},
		{"Case 2", m.Vec2D{X: 0, Y: 0}, math.Sqrt(0)},
		{"Case 3", m.Vec2D{X: -4, Y: 3}, 5},
		{"Case 4", m.Vec2D{X: 10, Y: 2}, math.Sqrt(104)},
		{"Case 5", m.Vec2D{X: 10, Y: 5}, math.Sqrt(125)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := tt.v.Normalize()

			if l != tt.len || err != nil {
				t.Errorf("Expected %v, Got %v", tt.len, l)
			} else if tt.v.Length() != 1 {
				t.Errorf("Didnt Normalized Correctly Even though Expected %v, Got %v", tt.len, l)
			}
		})
	}
}
