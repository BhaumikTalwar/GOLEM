package golem

import (
	"math"
)

func ToDegrees(angleRad float64) float64 {
	return angleRad * (180 / math.Pi)
}

func ToRadians(angleDeg float64) float64 {
	return angleDeg * (math.Pi / 180)
}

// for nomalizing to [-Pi, Pi]
func NormalizeAngle(angle float64) float64 {
	angle = math.Mod(angle, 2*math.Pi)

	if angle > math.Pi {
		angle -= 2 * math.Pi
	} else if angle < -math.Pi {
		angle += 2 * math.Pi
	}

	return angle
}

// for Normalizing angle to [0, 2*Pi]
func NormalizeAngleTo2Pi(angle float64) float64 {
	angle = math.Mod(angle, (2 * math.Pi))

	if angle < 0 {
		angle += 2 * math.Pi
	}

	return angle
}

func Clamp(f, low, high float64) float64 {
	if f > high {
		return high
	}

	if f < low {
		return low
	}

	return f
}
