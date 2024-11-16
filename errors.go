package golem

import "errors"

var (
	ErrZeroDet = errors.New("Determinat is Zero")
	ErrZeroMag = errors.New("Magnitude Is Zero")

	ErrZeroLen = errors.New("Length is Zero")
	ErrZeroDiv = errors.New("Cant Divide By Zero")

	ErrMaxOrder  = errors.New("Error: Cannot Have More than 3 axis")
	ErrRepeatRot = errors.New("Cant Repeat Rotation")

	ErrInSigAngle = errors.New("Insignificant Angle For Rotation")
	ErrInvalidLen = errors.New("Invalid Len Slice")

	ErrNormalizeError   = errors.New("Cant Normalize the Result")
	ErrInvalidOperation = errors.New("Invalid Operation: May result in inconsistent Result")

	ErrInvalidOrderString  = errors.New("Invalid Order String")
	ErrUnsupportedRotOrder = errors.New("unsupported rotation order: must include 'X', 'Y', and 'Z'")

	ErrInvalidInterPolParam = errors.New("Invalid Interpolation Parameter")
)
