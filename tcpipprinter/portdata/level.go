package portdata

import "unsafe"

// Level is a type parameter constraint that maches port data levels.
type Level interface {
	Level1 | Level2
}

// Size returns the number of bytes needed to marshal a port data
// structure of the indicated level.
func Size[T Level]() int {
	switch any((*T)(nil)).(type) {
	case *Level1:
		return int(unsafe.Sizeof(rawLevel1{}))
	case *Level2:
		return int(unsafe.Sizeof(rawLevel2{}))
	default:
		panic("unexpected type passed to portdata.Size")
	}
}
