package logger

type errorLoggerData struct {
	Error error
}

func (e errorLoggerData) GetKey() string {
	return "error"
}

func (e errorLoggerData) GetData() map[string]any {
	return map[string]any{
		"message": e.Error.Error(),
	}
}

func ObjError(err error) errorLoggerData {
	return errorLoggerData{
		Error: err,
	}
}
