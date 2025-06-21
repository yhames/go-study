package errors

import "fmt"

const (
	NotFoundUser = iota
	DatabaseError
)

var errorMessage = map[int64]string{
	NotFoundUser:  "user not found",
	DatabaseError: "database error: ",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errorMessage[code]; ok {
		return fmt.Errorf(message, args...)
	}
	return fmt.Errorf("internal server error")
}
