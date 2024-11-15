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

func (e EulerAngle) ToDegrees() (float64, float64, float64) {
	e.roll *= (180 / math.Pi)
	e.pitch *= (180 / math.Pi)
	e.yaw *= (180 / math.Pi)

	return e.roll, e.pitch, e.yaw
}

// for [ -Pi , Pi ]
func (e *EulerAngle) Normalize() {
	e.roll = NormalizeAngle(e.roll)
	e.pitch = NormalizeAngle(e.pitch)
	e.yaw = NormalizeAngle(e.yaw)
}

// for [ 0 , Pi ]"
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

func (e EulerAngle) ToRotMat3D(order string) (RotMat3D, error) {
	out := RotMat3D{}

	err := out.SetRot(order, e.roll, e.pitch, e.yaw)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (e EulerAngle) ToAxisAngle(order string) (AxisAngle, error) {
	rmat, err := e.ToRotMat3D(order)
	if err != nil {
		return AxisAngle{}, err
	}

	return rmat.ToAxisAngle(), nil
}
