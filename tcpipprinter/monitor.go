package tcpipprinter

import (
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winprint/accessoptions"
	"github.com/gentlemanautomaton/winprint/accessrights"
	"github.com/gentlemanautomaton/winprint/spoolerapi"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/portdata"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/tcpipmonapi"
)

// Monitor is an open connection to a TCP/IP printer port monitor.
type Monitor struct {
	mutex  sync.RWMutex
	handle syscall.Handle
}

// OpenMonitor opens a connection to the TCP/IP port monitor.
//
// It is the caller's responsibility to close the monitor object when finished
// with it.
func OpenMonitor(desiredAccess accessrights.Mask, options ...accessoptions.Option) (*Monitor, error) {
	handle, err := spoolerapi.Open(",XcvMonitor Standard TCP/IP Port", desiredAccess, options...)
	if err != nil {
		return nil, err
	}
	return &Monitor{handle: handle}, nil
}

// AddPort adds a new TCP/IP printer port.
//
// TODO: Switch to Level2, and then consider supplying a dedicated struct
// from this package
func (m *Monitor) AddPort(data portdata.Level2) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return tcpipmonapi.AddPort(m.handle, data)
}

// DeletePort deletes the TCP/IP port with the given name.
func (m *Monitor) DeletePort(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return tcpipmonapi.DeletePort(m.handle, name)
}

// Close releases all resources and system handles associated with m.
func (m *Monitor) Close() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return spoolerapi.Close(m.handle)
}
