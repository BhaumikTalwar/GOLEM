package math

import "math"

func ToDegrees(angleRad float64) float64 {
	return angleRad * (180 / math.Pi)
}

func ToRadians(angleDeg float64) float64 {
	return angleDeg * (math.Pi / 180)
}

func NormalizeAngle(angle float64) float64 {
	// for nomalizing to [-Pi, Pi]
	angle = math.Mod(angle, 2*math.Pi)

	if angle > math.Pi {
		angle -= 2 * math.Pi
	} else if angle < -math.Pi {
		angle += 2 * math.Pi
	}

	return angle
}

func NormalizeAngleTo2Pi(angle float64) float64 {
	// for Normalizing angle to [0, 2*Pi]
	angle = math.Mod(angle, (2 * math.Pi))

	if angle < 0 {
		angle += 2 * math.Pi
	}

	return angle
}