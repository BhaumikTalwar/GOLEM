package golem

import "math"

type AxisAngle struct {
	Axis  Vec3D
	Angle float64
}

func NewAxisAngle(axis Vec3D, angle float64) (AxisAngle, error) {
	_, err := axis.Normalize()
	if err != nil {
		return AxisAngle{}, err
	}

	return AxisAngle{
		Axis:  axis,
		Angle: angle,
	}, nil
}

func (a *AxisAngle) Set(axis Vec3D, angle float64) {
	a.Axis = axis
	a.Angle = angle
}

func (a *AxisAngle) SetAxis(axis Vec3D) {
	a.Axis = axis
}

func (a *AxisAngle) SetAngle(angle float64) {
	a.Angle = angle
}

func (a *AxisAngle) SetAxisX(angle float64) {
	a.Axis = Vec3D{X: 1, Y: 0, Z: 0}
	a.Angle = angle
}

func (a *AxisAngle) SetAxisY(angle float64) {
	a.Axis = Vec3D{X: 0, Y: 1, Z: 0}
	a.Angle = angle
}

func (a *AxisAngle) SetAxisZ(angle float64) {
	a.Axis = Vec3D{X: 0, Y: 0, Z: 1}
	a.Angle = angle
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
	_, err := a.Axis.Normalize()
	if err != nil {
		return Quaternion{}, err
	}

	sinHF, cosHF := math.Sincos(a.Angle / 2)

	return Quaternion{
		W: cosHF,
		X: a.Axis.X * sinHF,
		Y: a.Axis.Y * sinHF,
		Z: a.Axis.Z * sinHF,
	}, nil
}

func (a AxisAngle) ToRotMat3D() (RotMat3D, error) {
	_, err := a.Axis.Normalize()
	if err != nil {
		return RotMat3D{
			Mat3D: Mat3D{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		}, err
	}

	cos := math.Cos(a.Angle)
	sin := math.Sin(a.Angle)
	cosm1 := 1 - cos

	return RotMat3D{
		Mat3D: Mat3D{
			{
				cos + (a.Axis.X * a.Axis.X * cosm1),
				(a.Axis.X * a.Axis.Y * cosm1) - (a.Axis.Z * sin),
				(a.Axis.X * a.Axis.Z * cosm1) + (a.Axis.Y * sin),
			},
			{
				(a.Axis.Y * a.Axis.X * cosm1) + (a.Axis.Z * sin),
				cos + (a.Axis.Y * a.Axis.Y * cosm1),
				(a.Axis.Y * a.Axis.Z * cosm1) - (a.Axis.X * sin),
			},
			{
				(a.Axis.Z * a.Axis.X * cosm1) - (a.Axis.Y * sin),
				(a.Axis.Z * a.Axis.Y * cosm1) + (a.Axis.X * sin),
				cos + (a.Axis.Z * a.Axis.Z * cosm1),
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
