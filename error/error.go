package error

import (
	pb "example/communication"
	"strconv"
)

// Define GRError
type GRError struct {
	Rawerror pb.GreetReplyValidationError
	Message  string
	Code     int
}

// Implement Error interface
func (e GRError) Error() string {
	return strconv.Itoa(e.Code) + e.Message
}

type LRError struct {
	Rawerror pb.AccessReplyValidationError
	Message  string
	Code     int
}

// Implement Error interface
func (e LRError) Error() string {
	return /*strconv.Itoa(e.Code) + */e.Message
}
