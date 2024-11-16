package golem

import (
	"math"
)

type Vec2D struct {
	X, Y float64
}

func (v *Vec2D) Set(x, y float64) {
	v.X = x
	v.Y = y
}

func (v *Vec2D) SetZero() {
	v.X = 0.0
	v.Y = 0.0
}

func (v *Vec2D) Add(vec Vec2D) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v Vec2D) AddVec(vec Vec2D) Vec2D {
	v.Add(vec)
	return v
}

func (v *Vec2D) Sub(vec Vec2D) {
	v.X = v.X - vec.X
	v.Y = v.Y - vec.Y
}

func (v Vec2D) SubVec(vec Vec2D) Vec2D {
	v.Sub(vec)
	return v
}

func (v *Vec2D) ScalerMul(x float64) {
	v.X *= x
	v.Y *= x
}

func (v *Vec2D) ScalerDiv(x float64) {
	if x == 0 {
		return
	}

	v.X /= x
	v.Y /= x
}

func (v *Vec2D) IsEqual(vec Vec2D) bool {
	return (v.X == vec.X) && (v.Y == vec.Y)
}

func (v *Vec2D) IsNotEqual(vec Vec2D) bool {
	return (v.X != vec.X) || (v.Y != vec.Y)
}

func (v *Vec2D) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v *Vec2D) Dist(vec Vec2D) float64 {
	x := v.X - vec.X
	y := v.Y - vec.Y

	return math.Sqrt((x * x) + (y * y))
}

// returns the length and err in case of len == 0
func (v *Vec2D) Normalize() (float64, error) {
	l := v.Length()
	if l == 0 {
		return -1, ErrZeroLen
	}

	v.X = v.X / l
	v.Y = v.Y / l

	return l, nil
}

func (v Vec2D) Directon() Vec2D {
	v.Normalize()
	return v
}

func (v *Vec2D) Swap() {
	v.X = v.X + v.Y
	v.Y = v.X - v.Y
	v.X = v.X - v.Y
}

func (v *Vec2D) Reverse() {
	v.X *= -1
	v.Y *= -1
}

func (v *Vec2D) Dot(vec Vec2D) float64 {
	return (v.X * vec.X) + (v.Y * vec.Y)
}

// Gives the Magnitude of the CrossVec
func (v *Vec2D) Cross2D(vec Vec2D) float64 {
	return (v.X * vec.Y) - (v.Y * vec.X)
}

func (v *Vec2D) Rotate(theta float64) {
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	v.X = (cos * v.X) - (sin * v.Y)
	v.Y = (sin * v.X) + (cos * v.Y)
}

func (v Vec2D) RotateOf(theta, x, y float64) Vec2D {
	v.X -= x
	v.Y -= y

	v.Rotate(theta)

	v.X += x
	v.Y += y

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
		X: -1 * v.Y,
		Y: v.X,
	}
}

func (v Vec2D) RightPerpendicular() Vec2D {
	return Vec2D{
		X: v.Y,
		Y: -1 * v.X,
	}
}

func (v Vec2D) LerpV(vec Vec2D, t float64) (Vec2D, error) {
	if t < 0 || t > 1 {
		return Vec2D{}, ErrInvalidInterPolParam
	}

	return Vec2D{
		X: v.X + (t * (vec.X - v.X)),
		Y: v.Y + (t * (vec.Y - v.Y)),
	}, nil
}

func (v *Vec2D) Lerp(vec Vec2D, t float64) error {

	if t < 0 || t > 1 {
		return ErrInvalidInterPolParam
	}

	v.X = v.X + (t * (vec.X - v.X))
	v.Y = v.Y + (t * (vec.Y - v.Y))

	return nil
}
