package printerinfo

import "unsafe"

// Level is a type parameter constraint that maches printer info levels.
type Level interface {
	Level1 | Level2 | Level4 | Level5
}

// ID returns an integer identifying the layout of a printer info structure
// of the indicated level.
func ID[T Level]() int {
	switch any((*T)(nil)).(type) {
	case *Level1:
		return 1
	case *Level2:
		return 2
	case *Level4:
		return 4
	case *Level5:
		return 5
	default:
		panic("unexpected type passed to printerinfo.ID")
	}
}

// Size returns the number of bytes needed to marshal a printer info
// structure of the indicated level.
func Size[T Level]() int {
	switch any((*T)(nil)).(type) {
	case *Level1:
		return int(unsafe.Sizeof(rawLevel1{}))
	case *Level2:
		return int(unsafe.Sizeof(rawLevel2{}))
	case *Level4:
		return int(unsafe.Sizeof(rawLevel4{}))
	case *Level5:
		return int(unsafe.Sizeof(rawLevel5{}))
	default:
		panic("unexpected type passed to printerinfo.Size")
	}
}

// Unmarshal attempts to unmarshal data into info.
func Unmarshal[T Level](data []byte, info *T) error {
	switch v := any(info).(type) {
	case *Level1:
		return v.UnmarshalBinary(data)
	case *Level2:
		return v.UnmarshalBinary(data)
	case *Level4:
		return v.UnmarshalBinary(data)
	case *Level5:
		return v.UnmarshalBinary(data)
	default:
		panic("unexpected type passed to printerinfo.Size")
	}
}
