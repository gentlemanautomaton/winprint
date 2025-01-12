package portdata

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/tcpipprinter/portproto"
)

// Level1 holds level 1 TCP/IP printer port information.
type Level1 struct {
	Name            string
	Protocol        portproto.Type
	HostAddress     string
	SNMPCommunity   string
	DoubleSpool     bool
	Queue           string
	IPAddr          string
	PortNumber      uint32
	SNMPEnabled     bool
	SNMPDeviceIndex uint32
}

// UnmarshalBinary unmarshals the given level 1 port data into info.
func (info *Level1) UnmarshalBinary(data []byte) error {
	if len(data) < Size[Level1]() {
		return errors.New("portinfo.Level1.UnmarshalBinary: insufficient data buffer supplied for unmarshaling")
	}

	raw := (*rawLevel1)(unsafe.Pointer(&data[0]))

	if raw.Version != 1 {
		return fmt.Errorf("tcpipport: portdata: Level2.UnmarshalBinary: provided port data is version %d (expected version 1)", raw.Version)
	}

	info.Name = syscall.UTF16ToString(raw.Name[:])
	info.Protocol = raw.Protocol
	info.HostAddress = syscall.UTF16ToString(raw.HostAddress[:])
	info.SNMPCommunity = syscall.UTF16ToString(raw.SNMPCommunity[:])
	info.DoubleSpool = raw.DoubleSpool != 0
	info.Queue = syscall.UTF16ToString(raw.Queue[:])
	info.IPAddr = syscall.UTF16ToString(raw.IPAddr[:])
	info.PortNumber = raw.PortNumber
	info.SNMPEnabled = raw.SNMPEnabled != 0
	info.SNMPDeviceIndex = raw.SNMPDeviceIndex

	return nil
}

// https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_port_data_1
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/bd1e99bb-8a27-47c3-b276-e38fda1f9d34
type rawLevel1 struct {
	Name            [MaxPortNameLength]uint16
	Version         uint32
	Protocol        portproto.Type
	Size            uint32
	_               uint32 // reserved
	HostAddress     [MaxHostAddressLength1]uint16
	SNMPCommunity   [MaxSNMPCommunityLength]uint16
	DoubleSpool     uint32
	Queue           [MaxQueueNameLength]uint16
	IPAddr          [MaxIPAddrStringLength]uint16
	_               [540]byte // reserved
	PortNumber      uint32
	SNMPEnabled     uint32
	SNMPDeviceIndex uint32
}
