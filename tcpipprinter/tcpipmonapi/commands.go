package tcpipmonapi

import (
	"syscall"

	"github.com/gentlemanautomaton/winprint/spoolerapi"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/portdata"
)

// https://stackoverflow.com/questions/1325485/how-to-create-a-new-port-and-assign-it-to-a-printer
// https://learn.microsoft.com/en-us/windows-hardware/drivers/print/tcpmon-xcv-commands#addport-command
func AddPort(monitor syscall.Handle, data portdata.Level2) error {
	b, err := data.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = spoolerapi.XcvData(monitor, "AddPort", b)
	return err
}

// https://stackoverflow.com/questions/1325485/how-to-create-a-new-port-and-assign-it-to-a-printer
// https://learn.microsoft.com/en-us/windows-hardware/drivers/print/tcpmon-xcv-commands#configport-command
func ConfigPort(port syscall.Handle, data portdata.Level2) error {
	b, err := data.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = spoolerapi.XcvData(port, "ConfigPort", b)
	return err
}

// https://stackoverflow.com/questions/1325485/how-to-create-a-new-port-and-assign-it-to-a-printer
// https://learn.microsoft.com/en-us/windows-hardware/drivers/print/tcpmon-xcv-commands#addport-command
func DeletePort(monitor syscall.Handle, name string) error {
	data := portdata.Delete1{
		Name: name,
	}
	b, err := data.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = spoolerapi.XcvData(monitor, "DeletePort", b)
	return err
}
