package math

import (
	"errors"
	"math"
	"strings"
)

type RotMat3D struct {
	Mat3D
	order string
}

func RotMatX(angle float64) RotMat3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	r := RotMat3D{}

	r.Mat3D[0][0] = 1
	r.Mat3D[0][1] = 0
	r.Mat3D[0][2] = 0
	r.Mat3D[1][0] = 0
	r.Mat3D[1][1] = cos
	r.Mat3D[1][2] = -sin
	r.Mat3D[2][0] = 0
	r.Mat3D[2][1] = sin
	r.Mat3D[2][2] = cos

	return r
}

func RotMatY(angle float64) RotMat3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	r := RotMat3D{}

	r.Mat3D[0][0] = cos
	r.Mat3D[0][1] = 0
	r.Mat3D[0][2] = sin
	r.Mat3D[1][0] = 0
	r.Mat3D[1][1] = 1
	r.Mat3D[1][2] = 0
	r.Mat3D[2][0] = -sin
	r.Mat3D[2][1] = 0
	r.Mat3D[2][2] = cos

	return r
}

func RotMatZ(angle float64) RotMat3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	r := RotMat3D{}

	r.Mat3D[0][0] = cos
	r.Mat3D[0][1] = -sin
	r.Mat3D[0][2] = 0
	r.Mat3D[1][0] = sin
	r.Mat3D[1][1] = cos
	r.Mat3D[1][2] = 0
	r.Mat3D[2][0] = 0
	r.Mat3D[2][1] = 0
	r.Mat3D[2][2] = 1

	return r
}

func (r *RotMat3D) SetRot(order string, roll, pitch, yaw float64) error {
	if len(order) != 3 {
		return errors.New("Invalid String")
	}

	order = strings.ToUpper(order)
	if !strings.Contains(order, "X") || !strings.Contains(order, "Y") || !strings.Contains(order, "Z") {
		return errors.New("Invalid String")
	}

	r.order = order
	r.SetIdentity()

	for _, axis := range order {
		switch axis {

		case 'X':
			r.Multiply(RotMatX(roll).Mat3D)

		case 'Y':
			r.Multiply(RotMatY(pitch).Mat3D)

		case 'Z':
			r.Multiply(RotMatZ(yaw).Mat3D)
		}
	}

	return nil
}

func (r RotMat3D) EulerAngle() float64 {
	return math.Atan2(r.Mat3D[1][0], r.Mat3D[0][0])
}

func (r RotMat3D) EulerAngleDeg() float64 {
	return ToDegrees(math.Atan2(r.Mat3D[1][0], r.Mat3D[0][0]))
}

func (r RotMat3D) RotateVec2D(vec Vec2D) Vec2D {
	return Vec2D{
		x: (vec.x * r.Mat3D[0][0]) + (vec.y * r.Mat3D[0][1]),
		y: (vec.x * r.Mat3D[1][0]) + (vec.y * r.Mat3D[1][1]),
	}
}

func (r RotMat3D) RotateArndPoint(vec, center Vec2D) Vec2D {
	vec.Sub(center)
	vec = r.RotateVec2D(vec)
	vec.Add(center)

	return vec
}

func (r *RotMat3D) ReflectX() {
	r.Mat3D[1][1] *= -1
}

func (r *RotMat3D) ReflectY() {
	r.Mat3D[0][0] *= -1
}

func (r RotMat3D) SlerpR(target RotMat3D, t float64) RotMat3D {
	angle := r.EulerAngle()
	targetAngle := target.EulerAngle()

	interpolatedAngle := NormalizeAngle(angle + t*(targetAngle-angle))

	out := RotMat3D{}
	out.Set(interpolatedAngle)

	return out
}
