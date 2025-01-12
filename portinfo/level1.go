package portinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
)

// Level1 holds level 1 printer port information.
type Level1 struct {
	Name string
}

// UnmarshalBinary unmarshals the given level 1 port data into info.
func (info *Level1) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level1]() {
		return errors.New("insufficient data for portinfo.Level1 unmarshaling")
	}

	raw := (*rawLevel1)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.Name)

	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/port-info-1
type rawLevel1 struct {
	Name *uint16
}
