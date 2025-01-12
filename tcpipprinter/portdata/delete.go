package portdata

import (
	"syscall"
	"unsafe"
)

// Delete1 holds level 1 TCP/IP printer port deletion request data.
type Delete1 struct {
	Name string
}

// MarshalBinary marshals the level 1 port deletion data into a binary format
// expected by API calls.
func (info *Delete1) MarshalBinary() (data []byte, err error) {
	utf16Name, err := syscall.UTF16FromString(info.Name)
	if err != nil {
		return nil, newMarshalError(2, "port name is invalid")
	} else if len(utf16Name) > MaxPortNameLength {
		return nil, newMarshalError(2, "port name is too long (%d characters, max is %d)", len(utf16Name), MaxPortNameLength)
	}

	raw := rawDelete1{
		Version: 1,
	}

	copy(raw.Name[:], utf16Name)

	return unsafe.Slice((*byte)(unsafe.Pointer(&raw)), unsafe.Sizeof(raw)), nil
}

// https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_delete_port_data_1
type rawDelete1 struct {
	Name    [MaxPortNameLength]uint16
	_       [98]byte // reserved
	Version uint32
	_       uint32 // reserved
}
