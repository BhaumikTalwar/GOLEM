package golem

import (
	"math"
)

type Vec3D struct {
	x float64
	y float64
	z float64
}

func (v *Vec3D) Set(x, y, z float64) {
	v.x = x
	v.y = y
	v.z = z
}

func (v *Vec3D) SetZero() {
	v.x = 0.0
	v.y = 0.0
	v.z = 0.0
}

func (v *Vec3D) Add(vec Vec3D) {
	v.x += vec.x
	v.y += vec.y
	v.z += vec.z
}

func (v Vec3D) AddVec(vec Vec3D) Vec3D {
	v.Add(vec)
	return v
}

func (v *Vec3D) Sub(vec Vec3D) {
	v.x = v.x - vec.x
	v.y = v.y - vec.y
	v.z = v.z - vec.z
}

func (v Vec3D) SubVec(vec Vec3D) Vec3D {
	v.Sub(vec)
	return v
}

func (v *Vec3D) ScalerMul(x float64) {
	v.x *= x
	v.y *= x
	v.z *= x
}

func (v *Vec3D) ScalerDiv(x float64) {
	if x == 0 {
		return
	}

	v.x /= x
	v.y /= x
	v.z /= x
}

func (v *Vec3D) IsEqual(vec Vec3D) bool {
	return ((v.x == vec.x) && (v.y == vec.y) && (v.z == vec.z))
}

func (v *Vec3D) IsNotEqual(vec Vec3D) bool {
	return (v.x != vec.x) || (v.y != vec.y) || (v.z != vec.z)
}

func (v *Vec3D) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v *Vec3D) Dist(vec Vec3D) float64 {
	x := v.x - vec.x
	y := v.y - vec.y
	z := v.z - vec.z

	return math.Sqrt((x * x) + (y * y) + (z * z))
}

func (v *Vec3D) Normalize() (float64, error) {
	// returns the lengtyh
	l := v.Length()
	if l == 0 {
		return -1, ErrZeroLen
	}

	v.x = v.x / l
	v.y = v.y / l
	v.z = v.z / l

	return l, nil
}

func (v Vec3D) Directon() Vec3D {
	v.Normalize()
	return v
}

func (v *Vec3D) Reverse() {
	v.x *= -1
	v.y *= -1
	v.z *= -1
}

func (v *Vec3D) Dot(vec Vec3D) float64 {
	return (v.x * vec.x) + (v.y * vec.y) + (v.z * vec.z)
}

func (v *Vec3D) Cross(vec Vec3D) {
	*v = v.CrossV(vec)
}

func (v Vec3D) CrossV(vec Vec3D) Vec3D {
	return Vec3D{
		x: (v.y * vec.z) - (v.z * vec.y),
		y: (v.z * vec.x) - (v.x * vec.z),
		z: (v.x * vec.y) - (v.y * vec.x),
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
		w: 0,
		x: v.x,
		y: v.y,
		z: v.z,
	}

	qRot.Multiply(qVec)
	qRot.Multiply(qInv)

	v.x = qRot.x
	v.y = qRot.y
	v.z = qRot.z

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
	_, err := a.axis.Normalize()
	if err != nil {
		return Vec3D{}, err
	}

	cos := math.Cos(a.angle)
	sin := math.Sin(a.angle)

	dot := v.Dot(a.axis)
	cross := v.CrossV(a.axis)

	v.ScalerMul(cos)
	cross.ScalerMul(sin)
	a.axis.ScalerMul(dot * (1 - cos))

	v.Add(cross)
	v.Add(a.axis)

	return v, nil
}

func OrthoGraphicProjection(point Vec3D) Vec3D {
	return Vec3D{
		x: point.x,
		y: point.y,
		z: 0,
	}
}

func PerspectiveProjection(point Vec3D, focalLen float64) Vec3D {
	return Vec3D{
		x: (point.x * focalLen) / point.z,
		y: (point.y * focalLen) / point.z,
		z: 0,
	}
}

func (v Vec3D) LerpV(vec Vec3D, t float64) (Vec3D, error) {
	if t < 0 || t > 1 {
		return Vec3D{}, ErrInvalidInterPolParam
	}

	return Vec3D{
		x: v.x + (t * (vec.x - v.x)),
		y: v.y + (t * (vec.y - v.y)),
		z: v.z + (t * (vec.z - v.z)),
	}, nil
}

func (v *Vec3D) Lerp(vec Vec3D, t float64) error {
	if t < 0 || t > 1 {
		return ErrInvalidInterPolParam
	}

	v.x = v.x + (t * (vec.x - v.x))
	v.y = v.y + (t * (vec.y - v.y))
	v.z = v.z + (t * (vec.z - v.z))

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

		if math.Abs(v.x) < math.Abs(v.y) && math.Abs(v.x) < math.Abs(v.z) {
			axis = Vec3D{1, 0, 0}.CrossV(v)
		} else if math.Abs(v.y) < math.Abs(v.z) {
			axis = Vec3D{0, 1, 0}.CrossV(v)
		} else {
			axis = Vec3D{0, 0, 1}.CrossV(v)
		}

		axisAng := AxisAngle{
			axis:  axis,
			angle: math.Pi * t,
		}

		_, err := axisAng.axis.Normalize()
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
		x: (s1 * v.x) + (s2 * vec.x),
		y: (s1 * v.y) + (s2 * vec.y),
		z: (s1 * v.z) + (s2 * vec.z),
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
