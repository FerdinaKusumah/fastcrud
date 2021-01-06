package utils


// SendLogError is print error message to console
func SendLogError(text string, err error) error {
	ErrorLog.Errorf(text, err)
	return nil
}

// SendLogError is print info message to console
func SendLogInfo(text string) error {
	InfoLog.InfoF(text)
	return nil
}

// SendLogError is print debug message to console
func DebugLogInfo(text string) error {
	DebugLog.DebugF(text)
	return nil
}
