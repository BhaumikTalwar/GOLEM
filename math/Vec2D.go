package math

import (
	"errors"
	"math"
)

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

func (v *Vec2D) Normalize() (float64, error) {
	// returns the lengtyh
	l := v.Length()
	if l == 0 {
		return -1, errors.New("The Length Is Zero")
	}

	v.x = v.x / l
	v.y = v.y / l

	return l, nil
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

func (v *Vec2D) Dot(vec Vec2D) float64 {
	return (v.x * vec.x) + (v.y * vec.y)
}

func (v *Vec2D) Cross2D(vec Vec2D) float64 {
	// Gives the Magnitude of the CrossVec
	return (v.x * vec.y) - (v.y * vec.x)
}

func (v *Vec2D) Rotate(theta float64) {
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	v.x = (cos * v.x) - (sin * v.y)
	v.y = (sin * v.x) + (cos * v.y)
}

func (v Vec2D) RotateOf(theta, x, y float64) Vec2D {
	v.x -= x
	v.y -= y

	v.Rotate(theta)

	v.x += x
	v.y += y

	return v
}

func (v Vec2D) Projection(vec Vec2D) Vec2D {
	p := v.Dot(vec) / (math.Pow(vec.Length(), 2))
	v.ScalerMul(p)

	return v
}

func (v Vec2D) Reflection(Nvec Vec2D) Vec2D {
	dot := v.Dot(Nvec)
	magSq := Nvec.Dot(Nvec)
	Nvec.ScalerMul(2 * (dot / magSq))

	v.Sub(Nvec)
	return v
}

func (v Vec2D) AngleBetween(vec Vec2D) float64 {
	cos := v.Dot(vec) / (v.Length() * vec.Length())

	return math.Acos(cos)
}

func (v Vec2D) CosAngleBetween(vec Vec2D) float64 {
	return v.Dot(vec) / (v.Length() * vec.Length())
}

func (v Vec2D) LeftPerpendicular() Vec2D {
	return Vec2D{
		x: -1 * v.y,
		y: v.x,
	}
}

func (v Vec2D) RightPerpendicular() Vec2D {
	return Vec2D{
		x: v.y,
		y: -1 * v.x,
	}
}
