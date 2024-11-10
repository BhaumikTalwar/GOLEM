package math

import "math"

type EulerAngle struct {
	roll  float64
	pitch float64
	yaw   float64
}

func (e *EulerAngle) Set(roll, pitch, yaw float64) {
	e.roll = roll
	e.pitch = pitch
	e.yaw = yaw
}

func (e *EulerAngle) SetZero() {
	e.roll = 0
	e.pitch = 0
	e.yaw = 0
}

func (e *EulerAngle) ToDegrees() {
	e.roll *= (180 / math.Pi)
	e.pitch *= (180 / math.Pi)
	e.yaw *= (180 / math.Pi)
}

// for [-Pi , Pi ]
func (e *EulerAngle) Normalize() {
	e.roll = NormalizeAngle(e.roll)
	e.pitch = NormalizeAngle(e.pitch)
	e.yaw = NormalizeAngle(e.yaw)
}

// for [-Pi , Pi ]"
func (e *EulerAngle) NormalizeTo2Pi() {
	e.roll = NormalizeAngleTo2Pi(e.roll)
	e.pitch = NormalizeAngleTo2Pi(e.pitch)
	e.yaw = NormalizeAngleTo2Pi(e.yaw)
}

func (e EulerAngle) ToQuaternion() Quaternion {
	cosR, sinR := math.Sincos(e.roll / 2)
	cosP, sinP := math.Sincos(e.pitch / 2)
	cosY, sinY := math.Sincos(e.yaw / 2)

	return Quaternion{
		w: (cosR * cosP * cosY) + (sinR * sinP * sinY),
		x: (sinR * cosP * cosY) - (cosR * sinP * sinY),
		y: (cosR * sinP * cosY) + (sinR * cosP * sinY),
		z: (cosR * cosP * sinY) - (sinR * sinP * cosY),
	}
}
