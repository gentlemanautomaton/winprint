package portdata

import (
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/tcpipprinter/portproto"
)

// Level2 holds level 2 TCP/IP printer port information.
type Level2 struct {
	Name                string
	Protocol            portproto.Type
	HostAddress         string
	SNMPCommunity       string
	DoubleSpool         bool
	Queue               string
	PortNumber          uint32
	SNMPEnabled         bool
	SNMPDeviceIndex     uint32
	PortMonitorMIBIndex uint32
}

// MarshalBinary marshals the given level 2 port data into a binary format
// expected by API calls.
func (info *Level2) MarshalBinary() (data []byte, err error) {
	utf16Name, err := syscall.UTF16FromString(info.Name)
	if err != nil {
		return nil, newMarshalError(2, "port name is invalid")
	} else if len(utf16Name) > MaxPortNameLength {
		return nil, newMarshalError(2, "port name is too long (%d characters, max is %d)", len(utf16Name), MaxPortNameLength)
	}

	utf16HostAddress, err := syscall.UTF16FromString(info.HostAddress)
	if err != nil {
		return nil, newMarshalError(2, "port host address is invalid")
	} else if len(utf16HostAddress) > MaxHostAddressLength2 {
		return nil, newMarshalError(2, "port host address is too long (%d characters, max is %d)", len(utf16HostAddress), MaxHostAddressLength2)
	}

	utf16SNMPCommunity, err := syscall.UTF16FromString(info.SNMPCommunity)
	if err != nil {
		return nil, newMarshalError(2, "SNMP community is invalid")
	} else if len(utf16SNMPCommunity) > MaxSNMPCommunityLength {
		return nil, newMarshalError(2, "SNMP community is too long (%d characters, max is %d)", len(utf16SNMPCommunity), MaxSNMPCommunityLength)
	}

	utf16Queue, err := syscall.UTF16FromString(info.Queue)
	if err != nil {
		return nil, newMarshalError(2, "SNMP community is invalid")
	} else if len(utf16Queue) > MaxQueueNameLength {
		return nil, newMarshalError(2, "SNMP community is too long (%d characters, max is %d)", len(utf16Queue), MaxQueueNameLength)
	}

	raw := rawLevel2{
		Version:             2,
		Protocol:            info.Protocol,
		Size:                uint32(Size[Level2]()),
		PortNumber:          info.PortNumber,
		SNMPDeviceIndex:     info.SNMPDeviceIndex,
		PortMonitorMIBIndex: info.PortMonitorMIBIndex,
	}
	if info.DoubleSpool {
		raw.DoubleSpool = 1
	}
	if info.SNMPEnabled {
		raw.SNMPEnabled = 1
	}
	copy(raw.Name[:], utf16Name)
	copy(raw.HostAddress[:], utf16HostAddress)
	copy(raw.SNMPCommunity[:], utf16SNMPCommunity)
	copy(raw.Queue[:], utf16Queue)

	return unsafe.Slice((*byte)(unsafe.Pointer(&raw)), unsafe.Sizeof(raw)), nil
}

// UnmarshalBinary unmarshals the given level 1 port data into info.
func (info *Level2) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level2]() {
		return newUnmarshalError(2, "insufficient data buffer supplied for unmarshaling")
	}

	raw := (*rawLevel2)(unsafe.Pointer(&data[0]))

	if raw.Version != 2 {
		return newUnmarshalError(2, "provided port data is version %d (expected version 2)", raw.Version)
	}

	info.Name = syscall.UTF16ToString(raw.Name[:])
	info.Protocol = raw.Protocol
	info.HostAddress = syscall.UTF16ToString(raw.HostAddress[:])
	info.SNMPCommunity = syscall.UTF16ToString(raw.SNMPCommunity[:])
	info.DoubleSpool = raw.DoubleSpool != 0
	info.Queue = syscall.UTF16ToString(raw.Queue[:])
	info.PortNumber = raw.PortNumber
	info.SNMPEnabled = raw.SNMPEnabled != 0
	info.SNMPDeviceIndex = raw.SNMPDeviceIndex
	info.PortMonitorMIBIndex = raw.PortMonitorMIBIndex

	return nil
}

// https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_port_data_2
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/2d256a4b-819c-4b84-8e8f-b1edb1dcf0c2
type rawLevel2 struct {
	Name                [MaxPortNameLength]uint16
	Version             uint32
	Protocol            portproto.Type
	Size                uint32
	_                   uint32 // reserved
	HostAddress         [MaxHostAddressLength2]uint16
	SNMPCommunity       [MaxSNMPCommunityLength]uint16
	DoubleSpool         uint32
	Queue               [MaxQueueNameLength]uint16
	_                   [514]byte // reserved
	PortNumber          uint32
	SNMPEnabled         uint32
	SNMPDeviceIndex     uint32
	PortMonitorMIBIndex uint32
}
