package GOLEM

import "math"

type RotMat2D struct {
	Mat2D
}

func (r *RotMat2D) Set(angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	r.Mat2D[0][0] = cos
	r.Mat2D[0][1] = -sin
	r.Mat2D[1][0] = sin
	r.Mat2D[1][1] = cos
}

func (r *RotMat2D) SetFrom2Vec(v1, v2 Vec2D) {
	angle := NormalizeAngle(math.Atan2(v2.y, v2.x) - math.Atan2(v1.y, v1.x))
	r.Set(angle)
}

func (r RotMat2D) EulerAngle() float64 {
	return math.Atan2(r.Mat2D[1][0], r.Mat2D[0][0])
}

func (r RotMat2D) EulerAngleDeg() float64 {
	return ToDegrees(math.Atan2(r.Mat2D[1][0], r.Mat2D[0][0]))
}

func (r RotMat2D) RotateVec2D(vec Vec2D) Vec2D {
	return Vec2D{
		x: (vec.x * r.Mat2D[0][0]) + (vec.y * r.Mat2D[0][1]),
		y: (vec.x * r.Mat2D[1][0]) + (vec.y * r.Mat2D[1][1]),
	}
}

func (r RotMat2D) RotateArndPoint(vec, center Vec2D) Vec2D {
	vec.Sub(center)
	vec = r.RotateVec2D(vec)
	vec.Add(center)

	return vec
}

func (r *RotMat2D) ReflectX() {
	r.Mat2D[1][1] *= -1
}

func (r *RotMat2D) ReflectY() {
	r.Mat2D[0][0] *= -1
}

func (r RotMat2D) SlerpR(target RotMat2D, t float64) RotMat2D {
	angle := r.EulerAngle()
	targetAngle := target.EulerAngle()

	interpolatedAngle := NormalizeAngle(angle + t*(targetAngle-angle))

	out := RotMat2D{}
	out.Set(interpolatedAngle)

	return out
}
