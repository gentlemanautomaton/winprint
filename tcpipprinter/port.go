package tcpipprinter

import (
	"fmt"
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winprint/accessoptions"
	"github.com/gentlemanautomaton/winprint/accessrights"
	"github.com/gentlemanautomaton/winprint/spoolerapi"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/portdata"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/tcpipmonapi"
)

// Port is an open connection to a TCP/IP printer port.
type Port struct {
	mutex  sync.RWMutex
	handle syscall.Handle
}

// OpenPort opens a connection to a TCP/IP port.
//
// It is the caller's responsibility to close the port object when finished
// with it.
func OpenPort(name string, desiredAccess accessrights.Mask, options ...accessoptions.Option) (*Port, error) {
	path := fmt.Sprintf(",XcvPort %s", name)
	handle, err := spoolerapi.Open(path, desiredAccess, options...)
	if err != nil {
		return nil, err
	}
	return &Port{handle: handle}, nil
}

// Configure configures an existing TCP/IP printer port.
//
// TODO: Switch to Level2, and then consider supplying a dedicated struct
// from this package
func (p *Port) Configure(data portdata.Level2) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return tcpipmonapi.ConfigPort(p.handle, data)
}

// Close releases all resources and system handles associated with m.
func (p *Port) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return spoolerapi.Close(p.handle)
}
