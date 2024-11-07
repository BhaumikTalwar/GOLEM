package phy

type Vec2D struct {
	x, y float32
}

func (v *Vec2D) Set(x, y float32) {
	v.x = x
	v.y = y
}

func (v *Vec2D) SetZero() {
	v.x = 0.0
	v.y = 0.0
}
