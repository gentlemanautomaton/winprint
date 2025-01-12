package tcpipmonapi

import (
	"unsafe"
)

// Config1 holds level 1 TCP/IP printer port configuration request data.
type Config1 struct {
}

// MarshalBinary marshals the given level 2 port data into a binary format
// expected by API calls.
func (request *Config1) MarshalBinary() (data []byte, err error) {
	raw := config1{
		Version: 2,
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&raw)), unsafe.Sizeof(raw)), nil
}

// https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_config_info_data_1
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/9764f825-47a1-440c-973a-32328d204296
type config1 struct {
	_       [128]byte // reserved
	Version uint32
}
