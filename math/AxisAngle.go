package math

import "math"

type AxisAngle struct {
	axis  Vec3D
	angle float64
}

func (a *AxisAngle) Set(axis Vec3D, angle float64) {
	a.axis = axis
	a.angle = angle
}

func (a *AxisAngle) SetAxis(axis Vec3D) {
	a.axis = axis
}

func (a *AxisAngle) SetAngle(angle float64) {
	a.angle = angle
}

func (a *AxisAngle) SetAxisX(angle float64) {
	a.axis = Vec3D{x: 1, y: 0, z: 0}
	a.angle = angle
}

func (a *AxisAngle) SetAxisY(angle float64) {
	a.axis = Vec3D{x: 0, y: 1, z: 0}
	a.angle = angle
}

func (a *AxisAngle) SetAxisZ(angle float64) {
	a.axis = Vec3D{x: 0, y: 0, z: 1}
	a.angle = angle
}

func (a *AxisAngle) SetFromRotMat3D(r RotMat3D) {
	*a = r.ToAxisAngle()
}

func (a *AxisAngle) SetFromQuaternion(qt Quaternion) error {
	out, err := qt.ToAxisAngle()
	if err != nil {
		return err
	}

	*a = out
	return nil
}

func (a *AxisAngle) SetFromEulerAngle(e EulerAngle, order string) error {
	out, err := e.ToAxisAngle(order)
	if err != nil {
		return err
	}

	*a = out
	return nil
}

func (a AxisAngle) ToQuaternion() (Quaternion, error) {
	_, err := a.axis.Normalize()
	if err != nil {
		return Quaternion{}, err
	}

	sinHF, cosHF := math.Sincos(a.angle / 2)

	return Quaternion{
		w: cosHF,
		x: a.axis.x * sinHF,
		y: a.axis.y * sinHF,
		z: a.axis.z * sinHF,
	}, nil
}

func (a AxisAngle) ToRotMat3D() (RotMat3D, error) {
	_, err := a.axis.Normalize()
	if err != nil {
		return RotMat3D{
			Mat3D: Mat3D{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		}, err
	}

	cos := math.Cos(a.angle)
	sin := math.Sin(a.angle)
	cosm1 := 1 - cos

	return RotMat3D{
		Mat3D: Mat3D{
			{
				cos + (a.axis.x * a.axis.x * cosm1),
				(a.axis.x * a.axis.y * cosm1) - (a.axis.z * sin),
				(a.axis.x * a.axis.z * cosm1) + (a.axis.y * sin),
			},
			{
				(a.axis.y * a.axis.x * cosm1) + (a.axis.z * sin),
				cos + (a.axis.y * a.axis.y * cosm1),
				(a.axis.y * a.axis.z * cosm1) - (a.axis.x * sin),
			},
			{
				(a.axis.z * a.axis.x * cosm1) - (a.axis.y * sin),
				(a.axis.z * a.axis.y * cosm1) + (a.axis.x * sin),
				cos + (a.axis.z * a.axis.z * cosm1),
			},
		},
	}, nil
}

func (a AxisAngle) ToEulerAngle() (EulerAngle, error) {
	rm, err := a.ToRotMat3D()
	if err != nil {
		return EulerAngle{0, 0, 0}, err
	}

	return rm.ToEulerAngles()
}
