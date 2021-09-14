package esg

type ErrorType interface {
	error
	ErrorCode() interface{}
	StatusCode() int
}

var NoError = noError{}

type noError struct {
}

func (e noError) Error() string {
	return ""
}

func (e noError) ErrorCode() string {
	return "NoError"
}

func (e noError) StatusCode() int {
	return 200
}
