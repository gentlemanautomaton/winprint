package printerinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/printerattr"
)

// Level4 holds level 4 printer information.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-4
type Level4 struct {
	Name       string
	Server     string
	Attributes printerattr.Value
}

// UnmarshalBinary unmarshals the given level 4 printer data into info.
func (info *Level4) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level4]() {
		return errors.New("insufficient data for printerinfo.Level4 unmarshaling")
	}

	raw := (*rawLevel4)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.Name)
	info.Server = utf16conv.PtrToString(raw.Server)
	info.Attributes = raw.Attributes

	return nil
}

type rawLevel4 struct {
	Name       *uint16
	Server     *uint16
	Attributes printerattr.Value
}
