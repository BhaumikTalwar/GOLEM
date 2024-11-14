package math

import (
	"errors"
	"math"
)

var (
	ZeroMagERR    = errors.New("Magnitude Is Zero")
	InSigAngleERR = errors.New("Insignificant Angle For Rotation")
)

type Quaternion struct {
	w, x, y, z float64
}

func (q *Quaternion) Set(w, x, y, z float64) {
	q.w = w
	q.x = x
	q.y = y
	q.z = z
}

func (q *Quaternion) SetZero() {
	q.w = 0.0
	q.x = 0.0
	q.y = 0.0
	q.z = 0.0
}

func (q *Quaternion) SetFromAxisAngle(a AxisAngle) {
	a.axis.Normalize()
	sinHF, cosHF := math.Sincos(a.angle / 2)

	q.w = cosHF
	q.x = a.axis.x * sinHF
	q.y = a.axis.y * sinHF
	q.z = a.axis.z * sinHF
}

func (q *Quaternion) SetFromEulerAngles(e EulerAngle) {
	cosR, sinR := math.Sincos(e.roll / 2)
	cosP, sinP := math.Sincos(e.pitch / 2)
	cosY, sinY := math.Sincos(e.yaw / 2)

	q.w = (cosR * cosP * cosY) + (sinR * sinP * sinY)
	q.x = (sinR * cosP * cosY) - (cosR * sinP * sinY)
	q.y = (cosR * sinP * cosY) + (sinR * cosP * sinY)
	q.z = (cosR * cosP * sinY) - (sinR * sinP * cosY)
}

// Creates a Pure Quarternion from a Vec3d
func (q *Quaternion) SetFromVec3D(v Vec3D) {
	q.w = 0
	q.x = v.x
	q.y = v.y
	q.z = v.z
}

func (q *Quaternion) Add(qt Quaternion) {
	// Add qt to q
	q.w += qt.w
	q.x += qt.x
	q.y += qt.y
	q.z += qt.z
}

// returns a new Quaternions after addition
func (q Quaternion) AddQt(qt Quaternion) Quaternion {
	q.w += qt.w
	q.x += qt.x
	q.y += qt.y
	q.z += qt.z

	return q
}

func (q *Quaternion) ScaleBy(fac float64) {
	q.w *= fac
	q.x *= fac
	q.y *= fac
	q.z *= fac
}

func (q Quaternion) Magnitude() float64 {
	return math.Sqrt((q.w * q.w) + (q.x * q.x) + (q.y * q.y) + (q.z * q.z))
}

// returns the initial magnitude after normalizing
func (q *Quaternion) Normalize() (float64, error) {
	m := q.Magnitude()
	if m == 0 {
		return -1, ZeroMagERR
	}

	q.w = q.w / m
	q.x = q.x / m
	q.y = q.y / m
	q.z = q.z / m

	return m, nil
}

func (q Quaternion) Direction() (Quaternion, error) {
	_, err := q.Normalize()
	if err != nil {
		return Quaternion{}, err
	}

	return q, nil
}

func (q *Quaternion) Negate() {
	q.w *= -1
	q.x *= -1
	q.y *= -1
	q.z *= -1
}

func (q Quaternion) NegateQt() Quaternion {
	q.Negate()
	return q
}

func (q *Quaternion) Conjugate() {
	q.x *= -1
	q.y *= -1
	q.z *= -1
}

func (q Quaternion) ConjugateQt() Quaternion {
	q.Conjugate()
	return q
}

func (q *Quaternion) Inverse() error {
	magSq := q.Dot(*q)
	if magSq == 0 {
		return ZeroMagERR
	}

	q.Conjugate()
	q.ScaleBy(1 / magSq)

	return nil
}

func (q Quaternion) InverseQt() (Quaternion, error) {
	magSq := q.Dot(q)
	if magSq == 0 {
		return Quaternion{}, ZeroMagERR
	}

	q.Conjugate()
	q.ScaleBy(1 / magSq)

	return q, nil
}

func (q *Quaternion) Dot(qt Quaternion) float64 {
	return (q.x * qt.x) + (q.y * qt.y) + (q.z * qt.z)
}

func (q *Quaternion) Multiply(qt Quaternion) {
	w, x, y, z := q.w, q.x, q.y, q.z

	q.w = w*qt.w - x*qt.x - y*qt.y - z*qt.z
	q.x = w*qt.x + x*qt.w + y*qt.z - z*qt.y
	q.y = w*qt.y - x*qt.z + y*qt.w + z*qt.x
	q.z = w*qt.z + x*qt.y - y*qt.x + z*qt.w
}

func (q *Quaternion) MultiplyQt(qt Quaternion) Quaternion {
	return Quaternion{
		w: q.w*qt.w - q.x*qt.x - q.y*qt.y - q.z*qt.z,
		x: q.w*qt.x + q.x*qt.w + q.y*qt.z - q.z*qt.y,
		y: q.w*qt.y - q.x*qt.z + q.y*qt.w + q.z*qt.x,
		z: q.w*qt.z + q.x*qt.y - q.y*qt.x + q.z*qt.w,
	}

}

func (q Quaternion) RotateVec(vec Vec3D) (Vec3D, error) {
	p := Quaternion{0, vec.x, vec.y, vec.z}

	_, err := q.Normalize()
	if err != nil {
		return vec, err
	}

	qInv, err := q.InverseQt()
	if err != nil {
		return vec, err
	}

	q.Multiply(p)
	q.Multiply(qInv)

	return Vec3D{q.x, q.y, q.z}, nil
}

func (q Quaternion) ToAxisAngle() (AxisAngle, error) {
	_, err := q.Normalize()
	if err != nil {
		return AxisAngle{}, err
	}

	//TODO: To handle the edge case when q.w == 1 || q.w == -1
	// Maybe I should normalize the angle here hmm will see

	angle := 2 * math.Acos(q.w)
	sinHF := math.Sqrt(1 - (q.w * q.w))

	if angle < 1e-10 {
		// returns the arbitary axis of rotation
		return AxisAngle{
			axis:  Vec3D{1, 0, 0},
			angle: 0}, InSigAngleERR
	}

	axis := Vec3D{
		x: q.x / sinHF,
		y: q.y / sinHF,
		z: q.z / sinHF,
	}

	_, err = axis.Normalize()
	if err != nil {
		return AxisAngle{angle: angle}, err
	}

	return AxisAngle{axis: axis, angle: angle}, nil
}

func (q Quaternion) ToEulerAngles() EulerAngle {
	e := EulerAngle{}

	e.roll = math.Atan2(2*((q.w*q.x)+(q.y*q.z)), 1-(2*((q.x*q.x)+(q.y*q.y))))

	sinp := 2 * (q.w*q.y - q.z*q.x)
	if math.Abs(sinp) >= 1 {
		e.pitch = math.Copysign(math.Pi/2, sinp)
	} else {
		e.pitch = math.Asin(sinp)
	}

	e.yaw = math.Atan2(2*(q.w*q.z+q.x*q.y), 1-2*(q.y*q.y+q.z*q.z))

	return e
}

func (q *Quaternion) IsZero() bool {
	return q.w == 0 && q.x == 0 && q.y == 0 && q.z == 0
}

func (q *Quaternion) IsEqual(qt Quaternion) bool {
	return q.w == qt.w && q.x == qt.x && q.y == qt.y && q.z == qt.z
}
