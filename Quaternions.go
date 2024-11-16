package golem

import (
	"math"
)

type Quaternion struct {
	W, X, Y, Z float64
}

func (q *Quaternion) Set(w, x, y, z float64) {
	q.W = w
	q.X = x
	q.Y = y
	q.Z = z
}

func (q *Quaternion) SetZero() {
	q.W = 0.0
	q.X = 0.0
	q.Y = 0.0
	q.Z = 0.0
}

func (q *Quaternion) SetFromAxisAngle(a AxisAngle) {
	a.Axis.Normalize()
	sinHF, cosHF := math.Sincos(a.Angle / 2)

	q.W = cosHF
	q.X = a.Axis.X * sinHF
	q.Y = a.Axis.Y * sinHF
	q.Z = a.Axis.Z * sinHF
}

func (q *Quaternion) SetFromEulerAngles(e EulerAngle) {
	cosR, sinR := math.Sincos(e.Roll / 2)
	cosP, sinP := math.Sincos(e.Pitch / 2)
	cosY, sinY := math.Sincos(e.Yaw / 2)

	q.W = (cosR * cosP * cosY) + (sinR * sinP * sinY)
	q.X = (sinR * cosP * cosY) - (cosR * sinP * sinY)
	q.Y = (cosR * sinP * cosY) + (sinR * cosP * sinY)
	q.Z = (cosR * cosP * sinY) - (sinR * sinP * cosY)
}

// Trace Method Or Shephard's Method
func (q *Quaternion) SetFromRotMat3D(r RotMat3D) {

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
}

// Creates a Pure Quarternion from a Vec3d
func (q *Quaternion) SetFromVec3D(v Vec3D) {
	q.W = 0
	q.X = v.X
	q.Y = v.Y
	q.Z = v.Z
}

// Add qt to q
func (q *Quaternion) Add(qt Quaternion) {
	q.W += qt.W
	q.X += qt.X
	q.Y += qt.Y
	q.Z += qt.Z
}

// returns a new Quaternions after addition
func (q Quaternion) AddQt(qt Quaternion) Quaternion {
	q.W += qt.W
	q.X += qt.X
	q.Y += qt.Y
	q.Z += qt.Z

	return q
}

// Sub qt from q => q - qt
func (q *Quaternion) Sub(qt Quaternion) {
	q.W -= qt.W
	q.X -= qt.X
	q.Y -= qt.Y
	q.Z -= qt.Z
}

// returns a new Quaternions after Subtraction
func (q Quaternion) SubQt(qt Quaternion) Quaternion {
	q.W -= qt.W
	q.X -= qt.X
	q.Y -= qt.Y
	q.Z -= qt.Z

	return q
}

func (q *Quaternion) ScaleBy(fac float64) {
	q.W *= fac
	q.X *= fac
	q.Y *= fac
	q.Z *= fac
}

func (q Quaternion) ScaleByQt(fac float64) Quaternion {
	q.W *= fac
	q.X *= fac
	q.Y *= fac
	q.Z *= fac

	return q
}

func (q Quaternion) Magnitude() float64 {
	return math.Sqrt((q.W * q.W) + (q.X * q.X) + (q.Y * q.Y) + (q.Z * q.Z))
}

// returns the initial magnitude after normalizing
func (q *Quaternion) Normalize() (float64, error) {
	m := q.Magnitude()
	if m == 0 {
		return -1, ErrZeroMag
	}

	q.W = q.W / m
	q.X = q.X / m
	q.Y = q.Y / m
	q.Z = q.Z / m

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
	q.W *= -1
	q.X *= -1
	q.Y *= -1
	q.Z *= -1
}

func (q Quaternion) NegateQt() Quaternion {
	q.Negate()
	return q
}

func (q *Quaternion) Conjugate() {
	q.X *= -1
	q.Y *= -1
	q.Z *= -1
}

func (q Quaternion) ConjugateQt() Quaternion {
	q.Conjugate()
	return q
}

func (q *Quaternion) Inverse() error {
	magSq := q.Dot(*q)
	if magSq == 0 {
		return ErrZeroMag
	}

	q.Conjugate()
	q.ScaleBy(1 / magSq)

	return nil
}

func (q Quaternion) InverseQt() (Quaternion, error) {
	magSq := q.Dot(q)
	if magSq == 0 {
		return Quaternion{}, ErrZeroMag
	}

	q.Conjugate()
	q.ScaleBy(1 / magSq)

	return q, nil
}

func (q *Quaternion) Dot(qt Quaternion) float64 {
	return (q.X * qt.X) + (q.Y * qt.Y) + (q.Z * qt.Z)
}

func (q *Quaternion) Multiply(qt Quaternion) {
	w, x, y, z := q.W, q.X, q.Y, q.Z

	q.W = w*qt.W - x*qt.X - y*qt.Y - z*qt.Z
	q.X = w*qt.X + x*qt.W + y*qt.Z - z*qt.Y
	q.Y = w*qt.Y - x*qt.Z + y*qt.W + z*qt.X
	q.Z = w*qt.Z + x*qt.Y - y*qt.X + z*qt.W
}

func (q *Quaternion) MultiplyQt(qt Quaternion) Quaternion {
	return Quaternion{
		W: q.W*qt.W - q.X*qt.X - q.Y*qt.Y - q.Z*qt.Z,
		X: q.W*qt.X + q.X*qt.W + q.Y*qt.Z - q.Z*qt.Y,
		Y: q.W*qt.Y - q.X*qt.Z + q.Y*qt.W + q.Z*qt.X,
		Z: q.W*qt.Z + q.X*qt.Y - q.Y*qt.X + q.Z*qt.W,
	}

}

func (q Quaternion) RotateVec(vec Vec3D) (Vec3D, error) {
	p := Quaternion{0, vec.X, vec.Y, vec.Z}

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

	return Vec3D{q.X, q.Y, q.Z}, nil
}

func (q Quaternion) ToAxisAngle() (AxisAngle, error) {
	_, err := q.Normalize()
	if err != nil {
		return AxisAngle{}, err
	}

	//TODO: To handle the edge case when q.w == 1 || q.w == -1
	// Maybe I should normalize the angle here hmm will see

	angle := 2 * math.Acos(q.W)
	sinHF := math.Sqrt(1 - (q.W * q.W))

	if angle < 1e-10 {
		// returns the arbitary axis of rotation
		return AxisAngle{
			Axis:  Vec3D{1, 0, 0},
			Angle: 0}, ErrInSigAngle
	}

	axis := Vec3D{
		X: q.X / sinHF,
		Y: q.Y / sinHF,
		Z: q.Z / sinHF,
	}

	_, err = axis.Normalize()
	if err != nil {
		return AxisAngle{Angle: angle}, err
	}

	return AxisAngle{Axis: axis, Angle: angle}, nil
}

func (q Quaternion) ToEulerAngles() EulerAngle {
	e := EulerAngle{}

	e.Roll = math.Atan2(2*((q.W*q.X)+(q.Y*q.Z)), 1-(2*((q.X*q.X)+(q.Y*q.Y))))

	sinp := 2 * (q.W*q.Y - q.Z*q.X)
	if math.Abs(sinp) >= 1 {
		e.Pitch = math.Copysign(math.Pi/2, sinp)
	} else {
		e.Pitch = math.Asin(sinp)
	}

	e.Yaw = math.Atan2(2*(q.W*q.Z+q.X*q.Y), 1-2*(q.Y*q.Y+q.Z*q.Z))

	return e
}

func (q Quaternion) ToRotMat3D() RotMat3D {
	return RotMat3D{
		Order: QtSet,
		Mat3D: Mat3D{
			{
				1 - (2 * ((q.Y * q.Y) + (q.Z * q.Z))),
				2 * ((q.X * q.Y) - (q.Z * q.W)),
				2 * ((q.X * q.Z) + (q.Y * q.W)),
			},
			{
				2 * ((q.X * q.Y) + (q.Z * q.W)),
				1 - (2 * ((q.X * q.X) + (q.Z * q.Z))),
				2 * ((q.Y * q.Z) - (q.W * q.X)),
			},
			{
				2 * ((q.X * q.Z) - (q.W * q.Y)),
				2 * ((q.Y * q.Z) + (q.W * q.X)),
				1 - (2 * ((q.X * q.X) + (q.Y * q.Y))),
			},
		},
	}
}

func (q *Quaternion) IsZero() bool {
	return q.W == 0 && q.X == 0 && q.Y == 0 && q.Z == 0
}

func (q *Quaternion) IsEqual(qt Quaternion) bool {
	return q.W == qt.W && q.X == qt.X && q.Y == qt.Y && q.Z == qt.Z
}

func (q *Quaternion) Slerp(qt Quaternion, t float64) error {

	result, err := q.SlerpQt(qt, t)
	if err != nil {
		return err
	}

	*q = result

	return nil
}

func (q Quaternion) SlerpQt(qt Quaternion, t float64) (Quaternion, error) {
	if t < 0 || t > 1 {
		return Quaternion{}, ErrInvalidInterPolParam
	}

	if _, err := q.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	if _, err := qt.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	dot := q.Dot(qt)
	dot = Clamp(dot, -1, 1)

	if dot < 0 {
		qt.Negate()
		dot = -dot
	}

	if dot > 0.9995 {
		return q.LerpQt(qt, t)
	}

	theta := math.Acos(dot)
	sin := math.Sin(theta)

	s1 := math.Sin(((1 - t) * theta)) / sin
	s2 := math.Sin((t * theta)) / sin

	result := Quaternion{
		W: (s1 * q.W) + (s2 * qt.W),
		X: (s1 * q.X) + (s2 * qt.X),
		Y: (s1 * q.Y) + (s2 * qt.Y),
		Z: (s1 * q.Z) + (s2 * qt.Z),
	}

	if _, err := result.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	return result, nil
}

func (q *Quaternion) Lerp(qt Quaternion, t float64) error {
	result, err := q.LerpQt(qt, t)
	if err != nil {
		return err
	}

	*q = result

	return nil
}

func (q Quaternion) LerpQt(qt Quaternion, t float64) (Quaternion, error) {
	if t < 0 || t > 1 {
		return Quaternion{}, ErrInvalidInterPolParam
	}

	if _, err := q.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	if _, err := qt.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	result := Quaternion{
		W: (1-t)*q.W + t*qt.W,
		X: (1-t)*q.X + t*qt.X,
		Y: (1-t)*q.Y + t*qt.Y,
		Z: (1-t)*q.Z + t*qt.Z,
	}

	if _, err := result.Normalize(); err != nil {
		return Quaternion{}, ErrNormalizeError
	}

	return result, nil
}
