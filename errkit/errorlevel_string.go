// Code generated by "stringer -type=ErrorLevel"; DO NOT EDIT.

package errkit

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_Unknown - -1]
	_ = x[FATAL-0]
	_ = x[ERROR-1]
	_ = x[WARNING-2]
	_ = x[DEBUG-3]
	_ = x[INFO-4]
}

const _ErrorLevel_name = "_UnknownFATALERRORWARNINGDEBUGINFO"

var _ErrorLevel_index = [...]uint8{0, 8, 13, 18, 25, 30, 34}

func (i ErrorLevel) String() string {
	i -= -1
	if i < 0 || i >= ErrorLevel(len(_ErrorLevel_index)-1) {
		return "ErrorLevel(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _ErrorLevel_name[_ErrorLevel_index[i]:_ErrorLevel_index[i+1]]
}