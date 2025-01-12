package driverinfo

import "unsafe"

// Level is a type parameter constraint that maches driver info levels.
type Level interface {
	Level1 | Level2
}

// ID returns an integer identifying the layout of a driver info structure
// of the indicated level.
func ID[T Level]() int {
	switch any((*T)(nil)).(type) {
	case *Level1:
		return 1
	case *Level2:
		return 2
	default:
		panic("unexpected type passed to driverinfo.ID")
	}
}

// Size returns the number of bytes needed to marshal a driver info
// structure of the indicated level.
func Size[T Level]() int {
	switch any((*T)(nil)).(type) {
	case *Level1:
		return int(unsafe.Sizeof(rawLevel1{}))
	case *Level2:
		return int(unsafe.Sizeof(rawLevel2{}))
	default:
		panic("unexpected type passed to driverinfo.Size")
	}
}

// Unmarshal attempts to unmarshal data into info.
func Unmarshal[T Level](data []byte, info *T) error {
	switch v := any(info).(type) {
	case *Level1:
		return v.UnmarshalBinary(data)
	case *Level2:
		return v.UnmarshalBinary(data)
	default:
		return nil
	}
}
