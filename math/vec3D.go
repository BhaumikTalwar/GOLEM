package math

import (
	"errors"
	"math"
)

type Vec3D struct {
	x float64
	y float64
	z float64
}

func (v *Vec3D) Set(x, y, z float64) {
	v.x = x
	v.y = y
	v.z = z
}

func (v *Vec3D) SetZero() {
	v.x = 0.0
	v.y = 0.0
	v.z = 0.0
}

func (v *Vec3D) Add(vec Vec3D) {
	v.x += vec.x
	v.y += vec.y
	v.z += vec.z
}

func (v Vec3D) AddVec(vec Vec3D) Vec3D {
	v.Add(vec)
	return v
}

func (v *Vec3D) Sub(vec Vec3D) {
	v.x = v.x - vec.x
	v.y = v.y - vec.y
	v.z = v.z - vec.z
}

func (v Vec3D) SubVec(vec Vec3D) Vec3D {
	v.Sub(vec)
	return v
}

func (v *Vec3D) ScalerMul(x float64) {
	v.x *= x
	v.y *= x
	v.z *= x
}

func (v *Vec3D) ScalerDiv(x float64) {
	if x == 0 {
		return
	}

	v.x /= x
	v.y /= x
	v.z /= x
}

func (v *Vec3D) IsEqual(vec Vec3D) bool {
	return ((v.x == vec.x) && (v.y == vec.y) && (v.z == vec.z))
}

func (v *Vec3D) IsNotEqual(vec Vec3D) bool {
	return (v.x != vec.x) || (v.y != vec.y) || (v.z != vec.z)
}

func (v *Vec3D) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v *Vec3D) Dist(vec Vec3D) float64 {
	x := v.x - vec.x
	y := v.y - vec.y
	z := v.z - vec.z

	return math.Sqrt((x * x) + (y * y) + (z * z))
}

func (v *Vec3D) Normalize() float64 {
	// returns the lengtyh
	l := v.Length()
	v.x = v.x / l
	v.y = v.y / l
	v.z = v.z / l

	return l
}

func (v Vec3D) Directon() Vec3D {
	v.Normalize()
	return v
}

func (v *Vec3D) Reverse() {
	v.x *= -1
	v.y *= -1
	v.z *= -1
}

func (v *Vec3D) Dot(vec Vec3D) float64 {
	return (v.x * vec.x) + (v.y * vec.y) + (v.z * vec.z)
}

func (v Vec3D) Cross(vec Vec3D) Vec3D {
	return Vec3D{
		x: (v.y * vec.z) - (v.z * vec.y),
		y: (v.z * vec.x) - (v.x * vec.z),
		z: (v.x * vec.y) - (v.y * vec.x),
	}
}

func (v Vec3D) ProjectionOnto(vec Vec3D) Vec3D {
	p := v.Dot(vec) / (math.Pow(vec.Length(), 2))
	v.ScalerMul(p)

	return v
}

func (v Vec3D) Reflection(Nvec Vec3D) Vec3D {
	dot := v.Dot(Nvec)
	magSq := Nvec.Dot(Nvec)
	Nvec.ScalerMul(2 * (dot / magSq))

	v.Sub(Nvec)
	return v
}

func (v Vec3D) AngleBetween(vec Vec3D) (float64, error) {
	lenM := v.Length() * vec.Length()
	if lenM == 0 {
		return -1, errors.New("Cant divide by zero")
	}

	cos := v.Dot(vec) / (lenM)
	return math.Acos(cos), nil
}

func (v Vec3D) CosAngleBetween(vec Vec3D) (float64, error) {
	lenM := v.Length() * vec.Length()
	if lenM == 0 {
		return -1, errors.New("Cant divide by zero")
	}

	return v.Dot(vec) / (lenM), nil
}

func (v *Vec3D) Rotate(theta float64) {
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	v.x = (cos * v.x) - (sin * v.y)
	v.y = (sin * v.x) + (cos * v.y)
}

func (v Vec3D) RotateOf(theta, x, y float64) Vec3D {
	v.x -= x
	v.y -= y

	v.Rotate(theta)

	v.x += x
	v.y += y

	return v
}

func OrthoGraphicProjection(point Vec3D) Vec3D {
	return Vec3D{
		x: point.x,
		y: point.y,
		z: 0,
	}
}

func PerspectiveProjection(point Vec3D, focalLen float64) Vec3D {
	return Vec3D{
		x: (point.x * focalLen) / point.z,
		y: (point.y * focalLen) / point.z,
		z: 0,
	}
}