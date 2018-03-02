// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package common

const (
	// ErrInvalidBlock 非法区块
	ErrInvalidBlock = 1001
)

// Error ...
type Error int

// Error ...
func (err Error) Error() string {
	return errMap[err]
}

var errMap = map[Error]string{
	ErrInvalidBlock: "非法区块",
}

