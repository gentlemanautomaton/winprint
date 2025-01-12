package printerinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/printerattr"
)

// Level5 holds level 5 printer information.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-5
type Level5 struct {
	Name       string
	Ports      []string
	Attributes printerattr.Value
}

// UnmarshalBinary unmarshals the given level 5 printer data into info.
func (info *Level5) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level5]() {
		return errors.New("insufficient data for printerinfo.Level5 unmarshaling")
	}

	raw := (*rawLevel5)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.Name)
	info.Ports = utf16conv.PtrToStringSlice(raw.Ports)
	info.Attributes = raw.Attributes

	return nil
}

type rawLevel5 struct {
	Name       *uint16
	Ports      *uint16
	Attributes printerattr.Value
}
