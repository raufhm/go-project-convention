package logger

// Common field keys
const (
	FieldKeyError      = "error"
	FieldKeyDuration   = "duration"
	FieldKeyMethod     = "method"
	FieldKeyPath       = "path"
	FieldKeyRequestID  = "request_id"
	FieldKeyUserID     = "user_id"
	FieldKeyStatusCode = "status_code"
)

func NewField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func Error(err error) Field {
	return Field{Key: FieldKeyError, Value: err}
}

func Duration(value interface{}) Field {
	return Field{Key: FieldKeyDuration, Value: value}
}

func Method(value string) Field {
	return Field{Key: FieldKeyMethod, Value: value}
}

func Path(value string) Field {
	return Field{Key: FieldKeyPath, Value: value}
}

func RequestID(value string) Field {
	return Field{Key: FieldKeyRequestID, Value: value}
}

func UserID(value string) Field {
	return Field{Key: FieldKeyUserID, Value: value}
}

func StatusCode(value int) Field {
	return Field{Key: FieldKeyStatusCode, Value: value}
}
