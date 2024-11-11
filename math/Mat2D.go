package math

import "errors"

var (
	ErrZeroDet    = errors.New("Determinat is Zero")
	ErrInvalidLen = errors.New("Invalid Len Slice")
)

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

func (m *Mat2D) Transpose() {
	m[0][1] = m[0][1] + m[1][0]
	m[1][0] = m[0][1] - m[1][0]
	m[0][1] = m[0][1] - m[1][0]
}

func (m Mat2D) TranposeMat() Mat2D {
	m.Transpose()
	return m
}

func (m Mat2D) Det() float64 {
	return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
}

func (m *Mat2D) ToAdjoint() {
	m[0][0] = m[0][0] + m[1][1]
	m[1][1] = m[0][0] - m[1][1]
	m[0][0] = m[0][0] - m[1][1]

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

	}

	m.ToAdjoint()
	m.Scale(det)

	return nil
}

func (m Mat2D) InvserseMat() Mat2D {
	m.Inverse()
	return m
}

func (m Mat2D) Multiply(mat Mat2D) Mat2D {

}
