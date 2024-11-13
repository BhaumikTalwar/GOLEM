package math

import "errors"

type Mat2D [2][2]float64

func (m *Mat2D) Set(mat [][]float64) error {
	if len(mat) < 2 || len(mat[0]) < 2 || len(mat[1]) < 2 {
		return ErrInvalidLen
	}

	m[0][0] = mat[0][0]
	m[0][1] = mat[0][1]
	m[1][0] = mat[1][0]
	m[1][1] = mat[1][1]

	return nil
}

func (m *Mat2D) SetZero() {
	m[0][0] = 0
	m[0][1] = 0
	m[1][0] = 0
	m[1][1] = 0
}

func (m *Mat2D) SetIdentity() {
	m[0][0] = 1
	m[0][1] = 0
	m[1][0] = 0
	m[1][1] = 1
}

func (m *Mat2D) Add(mat Mat2D) {
	m[0][0] += mat[0][0]
	m[0][1] += mat[0][1]
	m[1][0] += mat[1][0]
	m[1][1] += mat[1][1]
}

func (m Mat2D) AddMat(mat Mat2D) Mat2D {
	m.Add(mat)
	return m
}

func (m *Mat2D) Sub(mat Mat2D) {
	m[0][0] -= mat[0][0]
	m[0][1] -= mat[0][1]
	m[1][0] -= mat[1][0]
	m[1][1] -= mat[1][1]
}

func (m Mat2D) SubMat(mat Mat2D) Mat2D {
	m.Sub(mat)
	return m
}
func (m *Mat2D) Scale(fac float64) {
	m[0][0] *= fac
	m[0][1] *= fac
	m[1][0] *= fac
	m[1][1] *= fac
}

func (m Mat2D) ScaleMat(fac float64) Mat2D {
	m.Scale(fac)
	return m
}

func (m *Mat2D) ScaleByVec2D(vec Vec2D) {
	m[0][0] *= vec.x
	m[0][1] *= vec.y
	m[1][0] *= vec.x
	m[1][1] *= vec.y
}

func (m *Mat2D) Transpose() {
	m[0][1], m[1][0] = m[1][0], m[0][1]
}

func (m Mat2D) TranposeMat() Mat2D {
	m.Transpose()
	return m
}

func (m Mat2D) Det() float64 {
	return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
}

func (m *Mat2D) ToAdjoint() {
	m[0][0], m[1][1] = m[1][1], m[0][0]

	m[0][1] *= -1
	m[1][0] *= -1
}

func (m *Mat2D) AdjointMat() Mat2D {
	return Mat2D{
		{m[1][1], -m[0][1]},
		{-m[1][0], m[0][0]},
	}
}

func (m *Mat2D) Inverse() error {
	det := m.Det()
	if det == 0 {
		return ErrZeroDet
	}

	m.ToAdjoint()
	m.Scale(1 / det)

	return nil
}

func (m Mat2D) InverseMat() Mat2D {
	m.Inverse()
	return m
}

func (m Mat2D) Multiply(mat Mat2D) Mat2D {
	out := Mat2D{}

	for k := 0; k < 2; k++ {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				out[i][j] += (m[i][k] * mat[k][j])
			}
		}
	}

	return out
}

func (m *Mat2D) IsEqual(mat Mat2D) bool {
	return m[0][0] == mat[0][0] && m[0][1] == mat[0][1] &&
		m[1][0] == mat[1][0] && m[1][1] == mat[1][1]
}

func (m *Mat2D) Trace() float64 {
	return m[0][0] + m[1][1]
}
