package math

type EulerAngle struct {
	roll  float64
	pitch float64
	yaw   float64
}

func (e *EulerAngle) Set(roll, pitch, yaw float64) {
	e.roll = roll
	e.pitch = pitch
	e.yaw = yaw
}

func (e *EulerAngle) SetZero() {
	e.roll = 0
	e.pitch = 0
	e.yaw = 0
}
