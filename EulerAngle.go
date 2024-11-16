package golem

import "math"

type EulerAngle struct {
	Roll  float64
	Pitch float64
	Yaw   float64
}

func (e *EulerAngle) Set(roll, pitch, yaw float64) {
	e.Roll = roll
	e.Pitch = pitch
	e.Yaw = yaw
}

func (e *EulerAngle) SetZero() {
	e.Roll = 0
	e.Pitch = 0
	e.Yaw = 0
}

func (e EulerAngle) ToDegrees() (float64, float64, float64) {
	e.Roll *= (180 / math.Pi)
	e.Pitch *= (180 / math.Pi)
	e.Yaw *= (180 / math.Pi)

	return e.Roll, e.Pitch, e.Yaw
}

// for [ -Pi , Pi ]
func (e *EulerAngle) Normalize() {
	e.Roll = NormalizeAngle(e.Roll)
	e.Pitch = NormalizeAngle(e.Pitch)
	e.Yaw = NormalizeAngle(e.Yaw)
}

// for [ 0 , Pi ]"
func (e *EulerAngle) NormalizeTo2Pi() {
	e.Roll = NormalizeAngleTo2Pi(e.Roll)
	e.Pitch = NormalizeAngleTo2Pi(e.Pitch)
	e.Yaw = NormalizeAngleTo2Pi(e.Yaw)
}

func (e EulerAngle) ToQuaternion() Quaternion {
	cosR, sinR := math.Sincos(e.Roll / 2)
	cosP, sinP := math.Sincos(e.Pitch / 2)
	cosY, sinY := math.Sincos(e.Yaw / 2)

	return Quaternion{
		W: (cosR * cosP * cosY) + (sinR * sinP * sinY),
		X: (sinR * cosP * cosY) - (cosR * sinP * sinY),
		Y: (cosR * sinP * cosY) + (sinR * cosP * sinY),
		Z: (cosR * cosP * sinY) - (sinR * sinP * cosY),
	}
}

func (e EulerAngle) ToRotMat3D(order string) (RotMat3D, error) {
	out := RotMat3D{}

	err := out.SetRot(order, e.Roll, e.Pitch, e.Yaw)
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
