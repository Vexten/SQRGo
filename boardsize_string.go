// Code generated by "stringer -type=BoardSize"; DO NOT EDIT.

package sqrgame

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Small-30]
	_ = x[Medium-50]
	_ = x[Large-70]
	_ = x[ExtraLarge-90]
}

const (
	_BoardSize_name_0 = "Small"
	_BoardSize_name_1 = "Medium"
	_BoardSize_name_2 = "Large"
	_BoardSize_name_3 = "ExtraLarge"
)

func (i BoardSize) String() string {
	switch {
	case i == 30:
		return _BoardSize_name_0
	case i == 50:
		return _BoardSize_name_1
	case i == 70:
		return _BoardSize_name_2
	case i == 90:
		return _BoardSize_name_3
	default:
		return "BoardSize(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
