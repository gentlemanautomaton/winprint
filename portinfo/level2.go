package portinfo

import (
	"errors"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/porttype"
)

// Level2 holds level 2 printer port information.
type Level2 struct {
	Name        string
	Monitor     string
	Description string
	Type        porttype.Value
}

// UnmarshalBinary unmarshals the given level 2 port data into info.
func (info *Level2) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level2]() {
		return errors.New("insufficient data for portinfo.Level2 unmarshaling")
	}

	raw := (*rawLevel2)(unsafe.Pointer(&data[0]))

	info.Name = utf16conv.PtrToString(raw.PortName)
	info.Monitor = utf16conv.PtrToString(raw.MonitorName)
	info.Description = utf16conv.PtrToString(raw.Description)
	info.Type = raw.PortType

	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/port-info-2
type rawLevel2 struct {
	PortName    *uint16
	MonitorName *uint16
	Description *uint16
	PortType    porttype.Value
	reserved    uint32
}
