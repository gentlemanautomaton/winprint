package monitorinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
)

// Level2 holds level 2 port monitor information.
type Level2 struct {
	Name        string
	Environment string
	Library     string
}

// UnmarshalBinary unmarshals the given level 2 port monitor data into info.
func (info *Level2) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level2]() {
		return errors.New("insufficient data for monitorinfo.Level2 unmarshaling")
	}

	raw := (*rawLevel2)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.Name)
	info.Environment = utf16conv.PtrToString(raw.Environment)
	info.Library = utf16conv.PtrToString(raw.DllName)

	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/monitor-info-2
type rawLevel2 struct {
	Name        *uint16
	Environment *uint16
	DllName     *uint16
}
