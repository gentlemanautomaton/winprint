package spoolerapi

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/accessoptions"
	"github.com/gentlemanautomaton/winprint/accessrights"
	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
)

var (
	procOpenPrinter2 = modwinspool.NewProc("OpenPrinter2W")
	procClosePrinter = modwinspool.NewProc("ClosePrinter")
)

type printerDefaults struct {
	DataType      *uint16
	DevMode       *byte
	DesiredAccess accessrights.Mask
}

type printerOptions struct {
	Size  uint32
	Flags accessoptions.Flags
}

// API use examples can be found in the cloudprint project:
// https://chromium.googlesource.com/chromium/src/+/02d5cc9941cf621f195871e960a82928a19a3076/cloud_print/virtual_driver/win/install/setup.cc

// Open attempts to open a handle to a printer, port or monitor described
// by name. It calls the OpenPrinter2 windows API function.
//
// If successful, it returns a handle that can be used to communicate with
// the named object.
func Open(name string, desiredAccess accessrights.Mask, options ...accessoptions.Option) (h syscall.Handle, err error) {
	// Collect options and apply them to data.
	var data accessoptions.Data
	for i := range options {
		options[i].Apply(&data)
	}

	// Marshal arguments as UTF-16.
	utf16Name, err := utf16conv.StringToPtr(name)
	if err != nil {
		return syscall.InvalidHandle, err
	}

	// Prepare a PRINTER_DEFAULTS structure.
	defaults := printerDefaults{
		DesiredAccess: desiredAccess,
	}

	// Prepare a PRINTER_OPTIONS structure.
	rawOptions := printerOptions{
		Flags: data.Flags,
	}
	rawOptions.Size = uint32(unsafe.Sizeof(rawOptions))

	// Perform the system call.
	result, _, errno := syscall.SyscallN(
		procOpenPrinter2.Addr(),
		uintptr(unsafe.Pointer(utf16Name)),
		uintptr(unsafe.Pointer(&h)),
		uintptr(unsafe.Pointer(&defaults)),
		uintptr(unsafe.Pointer(&rawOptions)))

	// If the function fails, the return value is zero.
	if result == 0 {
		return syscall.InvalidHandle, errno
	}

	return
}

// Closes releases any system resources associated with the given handle and
// invalidates it, preventing future use. The given handle should be one
// return by a previous call to Open.
func Close(h syscall.Handle) error {
	// Perform the system call.
	result, _, _ := syscall.SyscallN(
		procClosePrinter.Addr(),
		uintptr(h))

	// If the function fails, the return value is zero.
	if result == 0 {
		return errors.New("spoolerapi: failed to close handle")
	}

	return nil
}
