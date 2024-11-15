package GOLEM

import (
	"errors"
	"math"
	"strings"
)

const (
	QtSet = "SET" // For RotMat3D formed by Quaternion as there is no order associated there
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

	r.order = "X"

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

	r.order = "Y"

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

	r.order = "Z"

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

// For Creating RotMat from inidividual RotMat's of a particular Axis
// For casual multiplication use Mat3D.Multiply
func (r *RotMat3D) MultiplyRotMat(rmat RotMat3D) error {
	if len(r.order) >= 3 {
		errors.New("Cannot Have More than 3 axis")
	}

	if len(rmat.order) != 1 {
		errors.New("Invalid Operation: May result in inconsistent Result")
	}

	if strings.Contains(r.order, rmat.order) {
		return errors.New("Cant Repeat Rotation")
	}

	r.Multiply(rmat.Mat3D)
	r.order += rmat.order

	return nil
}

func (r *RotMat3D) Clear() {
	r.SetIdentity()
	r.order = ""
}

func (r *RotMat3D) ToEulerAngles() (EulerAngle, error) {
	e := EulerAngle{}

	if len(r.order) != 3 || !strings.Contains(r.order, "X") || !strings.Contains(r.order, "Y") || !strings.Contains(r.order, "Z") {
		return e, errors.New("unsupported rotation order: must include 'X', 'Y', and 'Z'")
	}

	switch r.order {
	case "XYZ":
		e.pitch = math.Asin(-r.Mat3D[2][0])
		if math.Abs(r.Mat3D[2][0]) < 0.99999 {
			e.roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[2][2])
			e.yaw = math.Atan2(r.Mat3D[1][0], r.Mat3D[0][0])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(-r.Mat3D[1][2], r.Mat3D[1][1])
		}

	case "XZY":
		e.pitch = math.Asin(r.Mat3D[1][0])
		if math.Abs(r.Mat3D[1][0]) < 0.99999 {
			e.roll = math.Atan2(-r.Mat3D[1][2], r.Mat3D[1][1])
			e.yaw = math.Atan2(-r.Mat3D[2][0], r.Mat3D[0][0])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[2][2])
		}

	case "YXZ":
		e.pitch = math.Asin(r.Mat3D[2][1])
		if math.Abs(r.Mat3D[2][1]) < 0.99999 {
			e.roll = math.Atan2(-r.Mat3D[2][0], r.Mat3D[2][2])
			e.yaw = math.Atan2(-r.Mat3D[0][1], r.Mat3D[1][1])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(r.Mat3D[0][2], r.Mat3D[0][0])
		}

	case "YZX":
		e.pitch = math.Asin(-r.Mat3D[0][1])
		if math.Abs(r.Mat3D[0][1]) < 0.99999 {
			e.roll = math.Atan2(r.Mat3D[2][1], r.Mat3D[1][1])
			e.yaw = math.Atan2(r.Mat3D[0][2], r.Mat3D[0][0])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(-r.Mat3D[2][0], r.Mat3D[2][2])
		}

	case "ZXY":
		e.pitch = math.Asin(-r.Mat3D[1][2])
		if math.Abs(r.Mat3D[1][2]) < 0.99999 {
			e.roll = math.Atan2(r.Mat3D[1][0], r.Mat3D[1][1])
			e.yaw = math.Atan2(r.Mat3D[0][2], r.Mat3D[2][2])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(-r.Mat3D[0][1], r.Mat3D[0][0])
		}

	case "ZYX":
		e.pitch = math.Asin(-r.Mat3D[0][2])
		if math.Abs(r.Mat3D[0][2]) < 0.99999 {
			e.roll = math.Atan2(r.Mat3D[1][2], r.Mat3D[2][2])
			e.yaw = math.Atan2(r.Mat3D[0][1], r.Mat3D[0][0])
		} else {
			e.yaw = 0
			e.roll = math.Atan2(-r.Mat3D[1][0], r.Mat3D[1][1])
		}

	default:
		return e, errors.New("unsupported rotation order")
	}

	return e, nil
}

// Trace Method Or Shephard's Method
func (r RotMat3D) ToQuaternion() Quaternion {
	q := Quaternion{}

	if t := r.Trace(); t > 0 {
		s := math.Sqrt(t+1) * 2

		q.w = 0.25 * s
		q.x = (r.Mat3D[2][1] - r.Mat3D[1][2]) / s
		q.y = (r.Mat3D[0][2] - r.Mat3D[2][0]) / s
		q.z = (r.Mat3D[1][0] - r.Mat3D[0][1]) / s

	} else if r.Mat3D[0][0] > r.Mat3D[1][1] && r.Mat3D[0][0] > r.Mat3D[2][2] {
		s := math.Sqrt(1.0+r.Mat3D[0][0]-r.Mat3D[1][1]-r.Mat3D[2][2]) * 2

		q.w = (r.Mat3D[2][1] - r.Mat3D[1][2]) / s
		q.x = 0.25 * s
		q.y = (r.Mat3D[0][1] + r.Mat3D[1][0]) / s
		q.z = (r.Mat3D[0][2] + r.Mat3D[2][0]) / s

	} else if r.Mat3D[1][1] > r.Mat3D[2][2] {
		s := math.Sqrt(1.0+r.Mat3D[1][1]-r.Mat3D[0][0]-r.Mat3D[2][2]) * 2

		q.w = (r.Mat3D[0][2] - r.Mat3D[2][0]) / s
		q.x = (r.Mat3D[0][1] + r.Mat3D[1][0]) / s
		q.y = 0.25 * s
		q.z = (r.Mat3D[1][2] + r.Mat3D[2][1]) / s

	} else {
		s := math.Sqrt(1.0+r.Mat3D[2][2]-r.Mat3D[0][0]-r.Mat3D[1][1]) * 2

		q.w = (r.Mat3D[1][0] - r.Mat3D[0][1]) / s
		q.x = (r.Mat3D[0][2] + r.Mat3D[2][0]) / s
		q.y = (r.Mat3D[1][2] + r.Mat3D[2][1]) / s
		q.z = 0.25 * s

	}

	return q
}

func (r RotMat3D) ToAxisAngle() AxisAngle {
	out := AxisAngle{}
	out.angle = math.Acos((r.Trace() - 1) / 2)

	if math.Abs(out.angle) < 1e-6 {
		out.axis = Vec3D{x: 1, y: 0, z: 0}

	} else {
		sin := math.Sin(out.angle)
		out.axis = Vec3D{
			x: (r.Mat3D[2][1] - r.Mat3D[1][2]) / (2 * sin),
			y: (r.Mat3D[0][2] - r.Mat3D[2][0]) / (2 * sin),
			z: (r.Mat3D[1][0] - r.Mat3D[0][1]) / (2 * sin),
		}
	}

	return out
}

func (r RotMat3D) RotateVec3D(vec Vec3D) Vec3D {
	return Vec3D{
		x: (vec.x * r.Mat3D[0][0]) + (vec.y * r.Mat3D[0][1]) + (vec.z * r.Mat3D[0][2]),
		y: (vec.x * r.Mat3D[1][0]) + (vec.y * r.Mat3D[1][1]) + (vec.z * r.Mat3D[1][2]),
		z: (vec.x * r.Mat3D[2][0]) + (vec.y * r.Mat3D[2][1]) + (vec.z * r.Mat3D[2][2]),
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
