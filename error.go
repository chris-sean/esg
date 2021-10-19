package esg

type ErrorType interface {
	error
	ErrorCode() interface{}
	StatusCode() int
	Extra() interface{}
}

var NoError = noError{}

type noError struct {
}

func (e noError) Error() string {
	return ""
}

func (e noError) ErrorCode() interface{} {
	return "NoError"
}

func (e noError) StatusCode() int {
	return 200
}

func (e noError) Extra() interface{} {
	return nil
}
