// Code generated by "stringer -type=GameState"; DO NOT EDIT.

package sqrgame

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NextMove-0]
	_ = x[WrongMove-1]
	_ = x[GameEnd-2]
	_ = x[Cheating-3]
}

const _GameState_name = "NextMoveWrongMoveGameEndCheating"

var _GameState_index = [...]uint8{0, 8, 17, 24, 32}

func (i GameState) String() string {
	if i >= GameState(len(_GameState_index)-1) {
		return "GameState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GameState_name[_GameState_index[i]:_GameState_index[i+1]]
}