package logger

type mapLoggerData struct {
	Key   string
	Value map[string]any
}

func (e mapLoggerData) GetKey() string {
	return e.Key
}

func (e mapLoggerData) GetData() map[string]any {
	return e.Value
}

func Obj(key string, value map[string]any) mapLoggerData {
	return mapLoggerData{
		Key:   key,
		Value: value,
	}
}
