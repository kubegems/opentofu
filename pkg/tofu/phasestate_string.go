// Code generated by "stringer -type phaseState"; DO NOT EDIT.

package tofu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[workingState-0]
	_ = x[refreshState-1]
	_ = x[prevRunState-2]
}

const _phaseState_name = "workingStaterefreshStateprevRunState"

var _phaseState_index = [...]uint8{0, 12, 24, 36}

func (i phaseState) String() string {
	if i < 0 || i >= phaseState(len(_phaseState_index)-1) {
		return "phaseState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _phaseState_name[_phaseState_index[i]:_phaseState_index[i+1]]
}
