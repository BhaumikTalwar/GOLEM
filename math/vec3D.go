package math

import "math"

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

	return math.Sqrt((x * x) + (y * y))
}

func (v *Vec3D) Normalize() float64 {
	// returns the lengtyh
	l := v.Length()
	v.x = v.x / l
	v.y = v.y / l

	return l
}

func (v Vec3D) Directon() Vec3D {
	v.Normalize()
	return v
}
