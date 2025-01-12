package printerinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/printerenum"
)

// Level1 holds level 1 printer information.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-1
type Level1 struct {
	Flags       printerenum.Flags
	Description string
	Name        string
	Comment     string
}

// UnmarshalBinary unmarshals the given level 1 printer data into info.
func (info *Level1) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level1]() {
		return errors.New("insufficient data for printerinfo.Level1 unmarshaling")
	}

	raw := (*rawLevel1)(unsafe.Pointer(&data[0]))

	info.Flags = raw.Flags
	info.Description = utf16conv.PtrToString(raw.Description)
	info.Name = utf16conv.PtrToString(raw.Name)
	info.Comment = utf16conv.PtrToString(raw.Comment)

	return nil
}

type rawLevel1 struct {
	Flags       printerenum.Flags
	Description *uint16
	Name        *uint16
	Comment     *uint16
}
