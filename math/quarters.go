package math

import (
	"errors"
	"math"
)

type Quaternion struct {
	w, x, y, z float64
}

func (q *Quaternion) Set(w, x, y, z float64) {
	q.w = w
	q.x = x
	q.y = y
	q.z = z
}

func (q *Quaternion) SetZero() {
	q.w = 0.0
	q.x = 0.0
	q.y = 0.0
	q.z = 0.0
}

func (q *Quaternion) Add(qt Quaternion) {
	// Add qt to q
	q.w += qt.w
	q.x += qt.x
	q.y += qt.y
	q.z += qt.z
}

func (q Quaternion) AddQt(qt Quaternion) Quaternion {
	// returns a new Quaternions after addition
	q.w += qt.w
	q.x += qt.x
	q.y += qt.y
	q.z += qt.z

	return q
}

func (q *Quaternion) ScaleBy(fac float64) {
	q.w *= fac
	q.x *= fac
	q.y *= fac
	q.z *= fac
}

func (q Quaternion) Magnitude() float64 {
	return math.Sqrt((q.w * q.w) + (q.x * q.x) + (q.y * q.y) + (q.z * q.z))
}

func (q *Quaternion) Normalize() (float64, error) {
	//returns the initial magnitude after normalizing
	m := q.Magnitude()
	if m == 0 {
		return -1, errors.New("Magnitude is Zero")
	}

	q.w = q.w / m
	q.x = q.x / m
	q.y = q.y / m
	q.z = q.z / m

	return m, nil
}

func (q Quaternion) Direction() Quaternion {
	q.Normalize()
	return q
}

func (q *Quaternion) Conjugate() {
	q.x *= -1
	q.y *= -1
	q.z *= -1
}

func (q Quaternion) ConjugateQt() Quaternion {
	q.Conjugate()
	return q
}

func (q *Quaternion) Dot(qt Quaternion) float64 {
	return (q.x * qt.x) + (q.y * qt.y) + (q.z * qt.z)
}
