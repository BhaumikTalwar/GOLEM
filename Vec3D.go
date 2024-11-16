package golem

import (
	"math"
)

type Vec3D struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3D) Set(x, y, z float64) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v *Vec3D) SetZero() {
	v.X = 0.0
	v.Y = 0.0
	v.Z = 0.0
}

func (v *Vec3D) Add(vec Vec3D) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

func (v Vec3D) AddVec(vec Vec3D) Vec3D {
	v.Add(vec)
	return v
}

func (v *Vec3D) Sub(vec Vec3D) {
	v.X = v.X - vec.X
	v.Y = v.Y - vec.Y
	v.Z = v.Z - vec.Z
}

func (v Vec3D) SubVec(vec Vec3D) Vec3D {
	v.Sub(vec)
	return v
}

func (v *Vec3D) ScalerMul(x float64) {
	v.X *= x
	v.Y *= x
	v.Z *= x
}

func (v *Vec3D) ScalerDiv(x float64) {
	if x == 0 {
		return
	}

	v.X /= x
	v.Y /= x
	v.Z /= x
}

func (v *Vec3D) IsEqual(vec Vec3D) bool {
	return ((v.X == vec.X) && (v.Y == vec.Y) && (v.Z == vec.Z))
}

func (v *Vec3D) IsNotEqual(vec Vec3D) bool {
	return (v.X != vec.X) || (v.Y != vec.Y) || (v.Z != vec.Z)
}

func (v *Vec3D) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v *Vec3D) Dist(vec Vec3D) float64 {
	x := v.X - vec.X
	y := v.Y - vec.Y
	z := v.Z - vec.Z

	return math.Sqrt((x * x) + (y * y) + (z * z))
}

func (v *Vec3D) Normalize() (float64, error) {
	// returns the lengtyh
	l := v.Length()
	if l == 0 {
		return -1, ErrZeroLen
	}

	v.X = v.X / l
	v.Y = v.Y / l
	v.Z = v.Z / l

	return l, nil
}

func (v Vec3D) Directon() Vec3D {
	v.Normalize()
	return v
}

func (v *Vec3D) Reverse() {
	v.X *= -1
	v.Y *= -1
	v.Z *= -1
}

func (v *Vec3D) Dot(vec Vec3D) float64 {
	return (v.X * vec.X) + (v.Y * vec.Y) + (v.Z * vec.Z)
}

func (v *Vec3D) Cross(vec Vec3D) {
	*v = v.CrossV(vec)
}

func (v Vec3D) CrossV(vec Vec3D) Vec3D {
	return Vec3D{
		X: (v.Y * vec.Z) - (v.Z * vec.Y),
		Y: (v.Z * vec.X) - (v.X * vec.Z),
		Z: (v.X * vec.Y) - (v.Y * vec.X),
	}
}

func (v Vec3D) ProjectionOnto(vec Vec3D) Vec3D {
	p := v.Dot(vec) / (math.Pow(vec.Length(), 2))
	v.ScalerMul(p)

	return v
}

func (v Vec3D) Reflection(Nvec Vec3D) Vec3D {
	dot := v.Dot(Nvec)
	magSq := Nvec.Dot(Nvec)
	Nvec.ScalerMul(2 * (dot / magSq))

	v.Sub(Nvec)
	return v
}

func (v Vec3D) AngleBetween(vec Vec3D) (float64, error) {
	lenM := v.Length() * vec.Length()
	if lenM == 0 {
		return -1, ErrZeroDiv
	}

	cos := v.Dot(vec) / (lenM)
	return math.Acos(cos), nil
}

func (v Vec3D) CosAngleBetween(vec Vec3D) (float64, error) {
	lenM := v.Length() * vec.Length()
	if lenM == 0 {
		return -1, ErrZeroDiv
	}

	return v.Dot(vec) / (lenM), nil
}

func (v *Vec3D) Rotate(qRot Quaternion) error {
	_, err := qRot.Normalize()
	if err != nil {
		return err
	}

	// As qRot is Normalized So conjugate == inverse
	qInv := qRot.ConjugateQt()

	qVec := Quaternion{
		W: 0,
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}

	qRot.Multiply(qVec)
	qRot.Multiply(qInv)

	v.X = qRot.X
	v.Y = qRot.Y
	v.Z = qRot.Z

	return nil
}

func (v Vec3D) RotateVec(qRot Quaternion) (Vec3D, error) {
	err := v.Rotate(qRot)
	if err != nil {
		return Vec3D{}, err
	}

	return v, nil
}

func (v *Vec3D) RotateByEuler(e EulerAngle) error {
	qRot := Quaternion{}
	qRot.SetFromEulerAngles(e)

	err := v.Rotate(qRot)
	if err != nil {
		return err
	}

	return nil
}

func (v Vec3D) RotateVecByEuler(e EulerAngle) (Vec3D, error) {
	err := v.RotateByEuler(e)
	if err != nil {
		return Vec3D{}, err
	}

	return v, nil
}

func (v Vec3D) RotateByAxisAngle(a AxisAngle) (Vec3D, error) {
	_, err := a.Axis.Normalize()
	if err != nil {
		return Vec3D{}, err
	}

	cos := math.Cos(a.Angle)
	sin := math.Sin(a.Angle)

	dot := v.Dot(a.Axis)
	cross := v.CrossV(a.Axis)

	v.ScalerMul(cos)
	cross.ScalerMul(sin)
	a.Axis.ScalerMul(dot * (1 - cos))

	v.Add(cross)
	v.Add(a.Axis)

	return v, nil
}

func (v Vec3D) RotateByRotMat3D(r RotMat3D) Vec3D {
	return r.RotateVec3D(v)
}

func OrthoGraphicProjection(point Vec3D) Vec3D {
	return Vec3D{
		X: point.X,
		Y: point.Y,
		Z: 0,
	}
}

func PerspectiveProjection(point Vec3D, focalLen float64) Vec3D {
	return Vec3D{
		X: (point.X * focalLen) / point.Z,
		Y: (point.Y * focalLen) / point.Z,
		Z: 0,
	}
}

func (v Vec3D) LerpV(vec Vec3D, t float64) (Vec3D, error) {
	if t < 0 || t > 1 {
		return Vec3D{}, ErrInvalidInterPolParam
	}

	return Vec3D{
		X: v.X + (t * (vec.X - v.X)),
		Y: v.Y + (t * (vec.Y - v.Y)),
		Z: v.Z + (t * (vec.Z - v.Z)),
	}, nil
}

func (v *Vec3D) Lerp(vec Vec3D, t float64) error {
	if t < 0 || t > 1 {
		return ErrInvalidInterPolParam
	}

	v.X = v.X + (t * (vec.X - v.X))
	v.Y = v.Y + (t * (vec.Y - v.Y))
	v.Z = v.Z + (t * (vec.Z - v.Z))

	return nil
}

func (v Vec3D) SlerpV(vec Vec3D, t float64) (Vec3D, error) {
	if t < 0 || t > 1 {
		return Vec3D{}, ErrInvalidInterPolParam
	}

	if _, err := v.Normalize(); err != nil {
		return Vec3D{}, ErrNormalizeError
	}

	if _, err := vec.Normalize(); err != nil {
		return Vec3D{}, ErrNormalizeError
	}

	dot := v.Dot(vec)
	dot = Clamp(dot, -1, 1)

	if dot > 0.9995 {
		return v.LerpV(vec, t)
	}

	if dot < -0.9995 {
		axis := Vec3D{}

		if math.Abs(v.X) < math.Abs(v.Y) && math.Abs(v.X) < math.Abs(v.Z) {
			axis = Vec3D{1, 0, 0}.CrossV(v)
		} else if math.Abs(v.Y) < math.Abs(v.Z) {
			axis = Vec3D{0, 1, 0}.CrossV(v)
		} else {
			axis = Vec3D{0, 0, 1}.CrossV(v)
		}

		axisAng := AxisAngle{
			Axis:  axis,
			Angle: math.Pi * t,
		}

		_, err := axisAng.Axis.Normalize()
		if err != nil {
			return Vec3D{}, ErrNormalizeError
		}

		return v.RotateByAxisAngle(axisAng)
	}

	theta := math.Acos(dot)
	sin := math.Sin(theta)

	s1 := math.Sin(((1 - t) * theta)) / sin
	s2 := math.Sin((t * theta)) / sin

	result := Vec3D{
		X: (s1 * v.X) + (s2 * vec.X),
		Y: (s1 * v.Y) + (s2 * vec.Y),
		Z: (s1 * v.Z) + (s2 * vec.Z),
	}

	if _, err := result.Normalize(); err != nil {
		return Vec3D{}, ErrNormalizeError
	}

	return result, nil
}

func (v *Vec3D) Slerp(vec Vec3D, t float64) error {
	result, err := v.SlerpV(vec, t)

	if err != nil {
		return err
	}

	*v = result
	return nil
}
