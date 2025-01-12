package printerinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/printerattr"
	"github.com/gentlemanautomaton/winprint/printerstatus"
)

// Level2 holds level 2 printer information.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
// https://learn.microsoft.com/en-us/troubleshoot/windows/win32/printer-print-job-status
type Level2 struct {
	Name            string
	Server          string
	ShareName       string
	Port            string
	Driver          string
	Comment         string
	Location        string
	SeparatorFile   string
	PrintProcessor  string
	DataType        string
	Parameters      string
	Attributes      printerattr.Value
	Priority        uint32
	DefaultPriority uint32
	StartTime       uint32
	UntilTime       uint32
	Status          printerstatus.Value
	Jobs            uint32
	AveragePPM      uint32
}

// UnmarshalBinary unmarshals the given level 2 printer data into info.
func (info *Level2) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level2]() {
		return errors.New("insufficient data for printerinfo.Level2 unmarshaling")
	}

	raw := (*rawLevel2)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.PrinterName)
	info.Server = utf16conv.PtrToString(raw.ServerName)
	info.ShareName = utf16conv.PtrToString(raw.ShareName)
	info.Port = utf16conv.PtrToString(raw.PortName)
	info.Driver = utf16conv.PtrToString(raw.DriverName)
	info.Comment = utf16conv.PtrToString(raw.Comment)
	info.Location = utf16conv.PtrToString(raw.Location)
	info.SeparatorFile = utf16conv.PtrToString(raw.SepFile)
	info.PrintProcessor = utf16conv.PtrToString(raw.PrintProcessor)
	info.DataType = utf16conv.PtrToString(raw.Datatype)
	info.Parameters = utf16conv.PtrToString(raw.Parameters)
	info.Attributes = raw.Attributes

	return nil
}

type rawLevel2 struct {
	ServerName      *uint16
	PrinterName     *uint16
	ShareName       *uint16
	PortName        *uint16
	DriverName      *uint16
	Comment         *uint16
	Location        *uint16
	devmode         uintptr
	SepFile         *uint16
	PrintProcessor  *uint16
	Datatype        *uint16
	Parameters      *uint16
	secdesc         uintptr
	Attributes      printerattr.Value
	Priority        uint32
	DefaultPriority uint32
	StartTime       uint32
	UntilTime       uint32
	Status          printerstatus.Value
	Jobs            uint32
	AveragePPM      uint32
}
