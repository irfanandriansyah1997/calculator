package customerror

type InvalidArgError struct {
	Message string
}

func (error InvalidArgError) Error() string {
	return "[Invalid Error]: " + error.Message
}
