package driverinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
)

// Level2 holds level 2 printer driver information.
type Level2 struct {
	Version     int
	Name        string
	Environment string
	DriverPath  string
	DataFile    string
	ConfigFile  string
}

// UnmarshalBinary unmarshals the given data into info.
func (info *Level2) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level2]() {
		return errors.New("insufficient data for driverinfo.Level2 unmarshaling")
	}

	raw := (*rawLevel2)(unsafe.Pointer(&data[0]))

	info.Version = int(raw.Version)
	info.Name = utf16conv.PtrToString(raw.Name)
	info.Environment = utf16conv.PtrToString(raw.Environment)
	info.DriverPath = utf16conv.PtrToString(raw.DriverPath)
	info.DataFile = utf16conv.PtrToString(raw.DataFile)
	info.ConfigFile = utf16conv.PtrToString(raw.ConfigFile)

	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/driver-info-1
type rawLevel2 struct {
	Version     uint32
	Name        *uint16
	Environment *uint16
	DriverPath  *uint16
	DataFile    *uint16
	ConfigFile  *uint16
}
