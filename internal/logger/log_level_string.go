// Code generated by "stringer -type=LogLevel -trimprefix=LogLevel -output=log_level_string.go"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LogLevelNotSpecified-0]
	_ = x[LogLevelDebug-1]
	_ = x[LogLevelInfo-2]
	_ = x[LogLevelWarn-3]
	_ = x[LogLevelError-4]
	_ = x[LogLevelFatal-5]
}

const _LogLevel_name = "NotSpecifiedDebugInfoWarnErrorFatal"

var _LogLevel_index = [...]uint8{0, 12, 17, 21, 25, 30, 35}

func (i LogLevel) String() string {
	if i >= LogLevel(len(_LogLevel_index)-1) {
		return "LogLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LogLevel_name[_LogLevel_index[i]:_LogLevel_index[i+1]]
}