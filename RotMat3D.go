package golem

import (
	"math"
	"strings"
)

const (
	QtSet = "SET" // For RotMat3D formed by Quaternion as there is no order associated there
)

type RotMat3D struct {
	Mat3D
	Order string
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

	r.Order = "X"

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

	r.Order = "Y"

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

	r.Order = "Z"

	return r
}

func (r *RotMat3D) SetRot(order string, roll, pitch, yaw float64) error {
	if len(order) != 3 {
		return ErrInvalidOrderString
	}

	order = strings.ToUpper(order)
	if !strings.Contains(order, "X") || !strings.Contains(order, "Y") || !strings.Contains(order, "Z") {
		return ErrInvalidOrderString
	}

	r.Order = order
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

// For Creating RotMat from inidividual RotMat's of a particular Axis
// For casual multiplication use Mat3D.Multiply
func (r *RotMat3D) MultiplyRotMat(rmat RotMat3D) error {
	if len(r.Order) >= 3 {
		return ErrMaxOrder
	}

	if len(rmat.Order) != 1 {
		return ErrInvalidOperation
	}

	if strings.Contains(r.Order, rmat.Order) {
		return ErrRepeatRot
	}

	r.Multiply(rmat.Mat3D)
	r.Order += rmat.Order

	return nil
}

func (r *RotMat3D) Clear() {
	r.SetIdentity()
	r.Order = ""
}

func (r *RotMat3D) ToEulerAngles() (EulerAngle, error) {
	e := EulerAngle{}

	if len(r.Order) != 3 || !strings.Contains(r.Order, "X") || !strings.Contains(r.Order, "Y") || !strings.Contains(r.Order, "Z") {
		return e, ErrUnsupportedRotOrder
	}

	switch r.Order {
	case "XYZ":
		e.Pitch = math.Asin(-r.Mat3D[2][0])
		if math.Abs(r.Mat3D[2][0]) < 0.99999 {
			e.Roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[2][2])
			e.Yaw = math.Atan2(r.Mat3D[1][0], r.Mat3D[0][0])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(-r.Mat3D[1][2], r.Mat3D[1][1])
		}

	case "XZY":
		e.Pitch = math.Asin(r.Mat3D[1][0])
		if math.Abs(r.Mat3D[1][0]) < 0.99999 {
			e.Roll = math.Atan2(-r.Mat3D[1][2], r.Mat3D[1][1])
			e.Yaw = math.Atan2(-r.Mat3D[2][0], r.Mat3D[0][0])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[2][2])
		}

	case "YXZ":
		e.Pitch = math.Asin(r.Mat3D[2][1])
		if math.Abs(r.Mat3D[2][1]) < 0.99999 {
			e.Roll = math.Atan2(-r.Mat3D[2][0], r.Mat3D[2][2])
			e.Yaw = math.Atan2(-r.Mat3D[0][1], r.Mat3D[1][1])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(r.Mat3D[0][2], r.Mat3D[0][0])
		}

	case "YZX":
		e.Pitch = math.Asin(-r.Mat3D[0][1])
		if math.Abs(r.Mat3D[0][1]) < 0.99999 {
			e.Roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[1][1])
			e.Yaw = math.Atan2(r.Mat3D[0][2], r.Mat3D[0][0])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(-r.Mat3D[2][0], r.Mat3D[2][2])
		}

	case "ZXY":
		e.Pitch = math.Asin(-r.Mat3D[1][2])
		if math.Abs(r.Mat3D[1][2]) < 0.99999 {
			e.Roll = math.Atan2(r.Mat3D[1][0], r.Mat3D[1][1])
			e.Yaw = math.Atan2(r.Mat3D[0][2], r.Mat3D[2][2])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(-r.Mat3D[0][1], r.Mat3D[0][0])
		}

	case "ZYX":
		e.Pitch = math.Asin(-r.Mat3D[0][2])
		if math.Abs(r.Mat3D[0][2]) < 0.99999 {
			e.Roll = math.Atan2(r.Mat3D[1][2], r.Mat3D[2][2])
			e.Yaw = math.Atan2(r.Mat3D[0][1], r.Mat3D[0][0])
		} else {
			e.Yaw = 0
			e.Roll = math.Atan2(-r.Mat3D[1][0], r.Mat3D[1][1])
		}

	default:
		return e, ErrUnsupportedRotOrder
	}

	return e, nil
}

// Trace Method Or Shephard's Method
func (r RotMat3D) ToQuaternion() Quaternion {
	q := Quaternion{}

	if t := r.Trace(); t > 0 {
		s := math.Sqrt(t+1) * 2

		q.W = 0.25 * s
		q.X = (r.Mat3D[2][1] - r.Mat3D[1][2]) / s
		q.Y = (r.Mat3D[0][2] - r.Mat3D[2][0]) / s
		q.Z = (r.Mat3D[1][0] - r.Mat3D[0][1]) / s

	} else if r.Mat3D[0][0] > r.Mat3D[1][1] && r.Mat3D[0][0] > r.Mat3D[2][2] {
		s := math.Sqrt(1.0+r.Mat3D[0][0]-r.Mat3D[1][1]-r.Mat3D[2][2]) * 2

		q.W = (r.Mat3D[2][1] - r.Mat3D[1][2]) / s
		q.X = 0.25 * s
		q.Y = (r.Mat3D[0][1] + r.Mat3D[1][0]) / s
		q.Z = (r.Mat3D[0][2] + r.Mat3D[2][0]) / s

	} else if r.Mat3D[1][1] > r.Mat3D[2][2] {
		s := math.Sqrt(1.0+r.Mat3D[1][1]-r.Mat3D[0][0]-r.Mat3D[2][2]) * 2

		q.W = (r.Mat3D[0][2] - r.Mat3D[2][0]) / s
		q.X = (r.Mat3D[0][1] + r.Mat3D[1][0]) / s
		q.Y = 0.25 * s
		q.Z = (r.Mat3D[1][2] + r.Mat3D[2][1]) / s

	} else {
		s := math.Sqrt(1.0+r.Mat3D[2][2]-r.Mat3D[0][0]-r.Mat3D[1][1]) * 2

		q.W = (r.Mat3D[1][0] - r.Mat3D[0][1]) / s
		q.X = (r.Mat3D[0][2] + r.Mat3D[2][0]) / s
		q.Y = (r.Mat3D[1][2] + r.Mat3D[2][1]) / s
		q.Z = 0.25 * s

	}

	return q
}

func (r RotMat3D) ToAxisAngle() AxisAngle {
	out := AxisAngle{}
	out.Angle = math.Acos((r.Trace() - 1) / 2)

	if math.Abs(out.Angle) < 1e-6 {
		out.Axis = Vec3D{X: 1, Y: 0, Z: 0}

	} else {
		sin := math.Sin(out.Angle)
		out.Axis = Vec3D{
			X: (r.Mat3D[2][1] - r.Mat3D[1][2]) / (2 * sin),
			Y: (r.Mat3D[0][2] - r.Mat3D[2][0]) / (2 * sin),
			Z: (r.Mat3D[1][0] - r.Mat3D[0][1]) / (2 * sin),
		}
	}

	return out
}

func (r RotMat3D) RotateVec3D(vec Vec3D) Vec3D {
	return Vec3D{
		X: (vec.X * r.Mat3D[0][0]) + (vec.Y * r.Mat3D[0][1]) + (vec.Z * r.Mat3D[0][2]),
		Y: (vec.X * r.Mat3D[1][0]) + (vec.Y * r.Mat3D[1][1]) + (vec.Z * r.Mat3D[1][2]),
		Z: (vec.X * r.Mat3D[2][0]) + (vec.Y * r.Mat3D[2][1]) + (vec.Z * r.Mat3D[2][2]),
	}
}

func (r RotMat3D) RotateArndPoint(vec, center Vec3D) Vec3D {
	vec.Sub(center)
	vec = r.RotateVec3D(vec)
	vec.Add(center)

	return vec
}

func (r *RotMat3D) ReflectX() {
	r.Mat3D[1][1] *= -1
	r.Mat3D[2][2] *= -1
}

func (r *RotMat3D) ReflectY() {
	r.Mat3D[0][0] *= -1
	r.Mat3D[2][2] *= -1
}

func (r *RotMat3D) ReflectZ() {
	r.Mat3D[0][0] *= -1
	r.Mat3D[1][1] *= -1
}

func (r *RotMat3D) ReflectXY() {
	r.Mat3D[2][2] *= -1
}

func (r *RotMat3D) ReflectYZ() {
	r.Mat3D[0][0] *= -1
}

func (r *RotMat3D) ReflectXZ() {
	r.Mat3D[1][1] *= -1
}

func (r RotMat3D) SlerpR(target RotMat3D, t float64) RotMat3D {
	q := r.ToQuaternion()
	q.Slerp(target.ToQuaternion(), t)

	return q.ToRotMat3D()
}
