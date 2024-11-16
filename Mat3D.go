package golem

type Mat3D [3][3]float64

func (m *Mat3D) Set(mat [][]float64) error {
	if len(mat) < 3 || len(mat[0]) < 3 || len(mat[1]) < 3 || len(mat[2]) < 3 {
		return ErrInvalidLen
	}

	m[0][0] = mat[0][0]
	m[0][1] = mat[0][1]
	m[0][2] = mat[0][2]

	m[1][0] = mat[1][0]
	m[1][1] = mat[1][1]
	m[1][2] = mat[1][2]

	m[2][0] = mat[2][0]
	m[2][1] = mat[2][1]
	m[2][2] = mat[2][2]

	return nil
}

func (m *Mat3D) SetZero() {
	m[0][0] = 0
	m[0][1] = 0
	m[0][2] = 0

	m[1][0] = 0
	m[1][1] = 0
	m[1][2] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 0
}

func (m *Mat3D) SetIdentity() {
	m[0][0] = 1
	m[0][1] = 0
	m[0][2] = 0

	m[1][0] = 0
	m[1][1] = 1
	m[1][2] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
}

func (m *Mat3D) Add(mat Mat3D) {
	m[0][0] += mat[0][0]
	m[0][1] += mat[0][1]
	m[0][2] += mat[0][2]

	m[1][0] += mat[1][0]
	m[1][1] += mat[1][1]
	m[1][2] += mat[1][2]

	m[2][0] += mat[2][0]
	m[2][1] += mat[2][1]
	m[2][2] += mat[2][2]
}

func (m Mat3D) AddMat(mat Mat3D) Mat3D {
	m.Add(mat)
	return m
}

func (m *Mat3D) Sub(mat Mat3D) {
	m[0][0] -= mat[0][0]
	m[0][1] -= mat[0][1]
	m[0][2] -= mat[0][2]

	m[1][0] -= mat[1][0]
	m[1][1] -= mat[1][1]
	m[1][2] -= mat[1][2]

	m[2][0] -= mat[2][0]
	m[2][1] -= mat[2][1]
	m[2][2] -= mat[2][2]
}

func (m Mat3D) SubMat(mat Mat3D) Mat3D {
	m.Sub(mat)
	return m
}
func (m *Mat3D) Scale(fac float64) {
	m[0][0] *= fac
	m[0][1] *= fac
	m[0][2] *= fac

	m[1][0] *= fac
	m[1][1] *= fac
	m[1][2] *= fac

	m[2][0] *= fac
	m[2][1] *= fac
	m[2][2] *= fac
}

func (m Mat3D) ScaleMat(fac float64) Mat3D {
	m.Scale(fac)
	return m
}

func (m *Mat3D) ScaleByVec2D(vec Vec3D) {
	m[0][0] *= vec.x
	m[0][1] *= vec.y
	m[0][2] *= vec.z

	m[1][0] *= vec.x
	m[1][1] *= vec.y
	m[1][2] *= vec.z

	m[2][0] *= vec.x
	m[2][1] *= vec.y
	m[2][2] *= vec.z
}

func (m *Mat3D) Transpose() {
	m[0][1], m[1][0] = m[1][0], m[0][1]
	m[0][2], m[2][0] = m[2][0], m[0][2]
	m[2][1], m[1][2] = m[1][2], m[2][1]
}

func (m Mat3D) TranposeMat() Mat3D {
	m.Transpose()
	return m
}

func (m Mat3D) Det() float64 {
	m1 := (m[1][1] * m[2][2]) - (m[2][1] * m[1][2])
	m2 := (m[1][0] * m[2][2]) - (m[2][0] * m[1][2])
	m3 := (m[1][0] * m[2][1]) - (m[2][0] * m[1][1])

	return (m[0][0] * m1) - (m[0][1] * m2) + (m[0][2] * m3)
}

func (m Mat3D) AdjointMat() Mat3D {
	adj := Mat3D{}

	adj[0][0] = m[1][1]*m[2][2] - m[1][2]*m[2][1]
	adj[0][1] = -(m[1][0]*m[2][2] - m[1][2]*m[2][0])
	adj[0][2] = m[1][0]*m[2][1] - m[1][1]*m[2][0]

	adj[1][0] = -(m[0][1]*m[2][2] - m[0][2]*m[2][1])
	adj[1][1] = m[0][0]*m[2][2] - m[0][2]*m[2][0]
	adj[1][2] = -(m[0][0]*m[2][1] - m[0][1]*m[2][0])

	adj[2][0] = m[0][1]*m[1][2] - m[0][2]*m[1][1]
	adj[2][1] = -(m[0][0]*m[1][2] - m[0][2]*m[1][0])
	adj[2][2] = m[0][0]*m[1][1] - m[0][1]*m[1][0]

	return adj

}

func (m *Mat3D) ToAdjoint() {
	*m = m.AdjointMat()
}

func (m *Mat3D) Inverse() error {
	det := m.Det()
	if det == 0 {
		return ErrZeroDet
	}

	m.ToAdjoint()
	m.Scale(1 / det)

	return nil
}

func (m Mat3D) InverseMat() Mat3D {
	m.Inverse()
	return m
}

func (m Mat3D) Multiply(mat Mat3D) Mat3D {
	out := Mat3D{}

	for k := 0; k < 3; k++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				out[i][j] += (m[i][k] * mat[k][j])
			}
		}
	}

	return out
}

func (m *Mat3D) IsEqual(mat Mat3D) bool {
	return (m[0][0] == mat[0][0] && m[0][1] == mat[0][1] && m[0][2] == mat[0][2] &&
		m[1][0] == mat[1][0] && m[1][1] == mat[1][1] && m[1][2] == mat[1][2] &&
		m[2][0] == mat[2][0] && m[2][1] == mat[2][1] && m[2][2] == mat[2][2])

}

func (m *Mat3D) IsIdentity() bool {
	return (m[0][0] == 1 && m[0][1] == 0 && m[0][2] == 0 &&
		m[1][0] == 0 && m[1][1] == 1 && m[1][2] == 0 &&
		m[2][0] == 0 && m[2][1] == 0 && m[2][2] == 1)

}

func (m *Mat3D) Trace() float64 {
	return m[0][0] + m[1][1] + m[2][2]
}
