package phy

import "math"

type Vec2D struct {
	x, y float64
}

func (v *Vec2D) Set(x, y float64) {
	v.x = x
	v.y = y
}

func (v *Vec2D) SetZero() {
	v.x = 0.0
	v.y = 0.0
}

func (v *Vec2D) Add(vec Vec2D) {
	v.x += vec.x
	v.y += vec.y
}

func (v Vec2D) AddVec(vec Vec2D) Vec2D {
	v.Add(vec)
	return v
}

func (v *Vec2D) Sub(vec Vec2D) {
	v.x = v.x - vec.x
	v.y = v.y - vec.y
}

func (v Vec2D) SubVec(vec Vec2D) Vec2D {
	v.Sub(vec)
	return v
}

func (v *Vec2D) ScalerMul(x float64) {
	v.x *= x
	v.y *= x
}

func (v *Vec2D) ScalerDiv(x float64) {
	if x == 0 {
		return
	}

	v.x /= x
	v.y /= x
}

func (v *Vec2D) IsEqual(vec Vec2D) bool {
	return (v.x == vec.x) && (v.y == vec.y)
}

func (v *Vec2D) IsNotEqual(vec Vec2D) bool {
	return (v.x != vec.x) || (v.y != vec.y)
}

func (v *Vec2D) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y))
}

func (v *Vec2D) Dist(vec Vec2D) float64 {
	x := v.x - vec.x
	y := v.y - vec.y

	return math.Sqrt((x * x) + (y * y))
}

func (v *Vec2D) Normalize() float64 {
	// returns the lengtyh
	l := v.Length()
	v.x = v.x / l
	v.y = v.y / l

	return l
}

func (v Vec2D) Directon() Vec2D {
	v.Normalize()
	return v
}

func (v *Vec2D) Swap() {
	v.x = v.x + v.y
	v.y = v.x - v.y
	v.x = v.x - v.y
}

func (v *Vec2D) Reverse() {
	v.x *= -1
	v.y *= -1
}
