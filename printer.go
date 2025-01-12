package winprint

import (
	"fmt"
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winprint/accessoptions"
	"github.com/gentlemanautomaton/winprint/accessrights"
	"github.com/gentlemanautomaton/winprint/objname"
	"github.com/gentlemanautomaton/winprint/spoolerapi"
)

// Printer is an open connection to a printer.
type Printer struct {
	mutex  sync.RWMutex
	handle syscall.Handle
}

// OpenPrinter opens a connection to a printer with the given name.
//
// It is the caller's responsibility to close the printer object when finished
// with it.
func OpenPrinter(name string, desiredAccess accessrights.Mask, options ...accessoptions.Option) (*Printer, error) {
	// Check to make sure we were provided a simple printer name
	if t := objname.DetectType(name); t != objname.Unspecified {
		return nil, fmt.Errorf("winprint.OpenPrinter was unexpectedly provided a %s name", t)
	}

	handle, err := spoolerapi.Open(name, desiredAccess, options...)
	if err != nil {
		return nil, err
	}
	return &Printer{handle: handle}, nil
}

// Close releases all resources and system handles associated with p.
func (p *Printer) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return spoolerapi.Close(p.handle)
}

// Delete requests that the printer p be deleted. The deletion will not take
// place until all existing print jobs have finished and p is closed.
func (p *Printer) Delete() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return spoolerapi.DeletePrinter(p.handle)
}
