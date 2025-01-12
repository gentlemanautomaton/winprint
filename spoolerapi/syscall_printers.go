package spoolerapi

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/printerenum"
	"github.com/gentlemanautomaton/winprint/printerinfo"
)

var (
	procEnumPrinters  = modwinspool.NewProc("EnumPrintersW")
	procAddPrinter    = modwinspool.NewProc("AddPrinter")
	procDeletePrinter = modwinspool.NewProc("DeletePrinter")
)

// EnumPrinters returns information about printers, printer servers, domains
// and print providers. It calls the EnumPrinters windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/enumprinters
func EnumPrinters[T printerinfo.Level](flags printerenum.Flags, name string) (printers []T, err error) {
	// Marshal arguments as UTF-16.
	utf16Name, err := utf16conv.StringToOptionalPtr(name)
	if err != nil {
		return nil, err
	}

	// Make up to 3 attempts to get the driver list:
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported size (first attempt)
	// 3: Using a dynamic buffer based on the reported size (second attempt)
	//
	// It is unlikely, but feasible, that the length could change between
	// calls if the driver list changes.

	var (
		scratch  [4096]byte
		buffer   = scratch[:]
		level    = printerinfo.ID[T]()
		infoSize = printerinfo.Size[T]()
	)

	for attempt := 0; ; attempt++ {
		var bufferSize, entries uint32
		result, _, errno := syscall.SyscallN(
			procEnumPrinters.Addr(),
			uintptr(flags),
			uintptr(unsafe.Pointer(utf16Name)),
			uintptr(level),
			uintptr(unsafe.Pointer(&buffer[0])),
			uintptr(len(buffer)),
			uintptr(unsafe.Pointer(&bufferSize)),
			uintptr(unsafe.Pointer(&entries)))

		// If the function fails, the return value is 0.
		if result == 0 {
			if attempt > 2 || errno != syscall.ERROR_INSUFFICIENT_BUFFER {
				return nil, errno
			}
			buffer = make([]byte, bufferSize)
			continue
		}

		// Trim the buffer down to the size that was filled with data.
		buffer = buffer[:bufferSize]

		// Verify that the buffer is sized correctly.
		if needed := int(entries) * infoSize; len(buffer) < needed {
			return nil, fmt.Errorf("failed to enumerate printers: the returned buffer size is insufficient (needed %d but got %d)", needed, len(buffer))
		}

		// Unmarshal the raw data into a slice of printers of the requested
		// info level.
		printers = make([]T, entries)
		for i := 0; i < len(printers); i++ {
			data := buffer[i*infoSize : (i+1)*infoSize]
			if err := printerinfo.Unmarshal(data, &printers[i]); err != nil {
				return nil, fmt.Errorf("failed to enumerate printers: entry %d: %w", i, err)
			}
		}

		return printers, nil
	}
}

// AddPrinter adds a printer with the given configuration.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/deleteprinter
func AddPrinter[T printerinfo.Level2](name string, config T) (err error) {
	// Marshal arguments as UTF-16.
	utf16Name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	// Perform the system call.
	result, _, errno := syscall.SyscallN(
		procAddPrinter.Addr(),
		uintptr(unsafe.Pointer(utf16Name)),
		uintptr(printerinfo.ID[T]()),
		0)

	// If the function fails, the result is 0.
	if result == 0 {
		fmt.Printf("Error: %d\n", errno)
		return errno
	}

	return nil
}

// DeletePrinter marks the printer with the given handle for deletion.
// It calls the DeletePrinter windows API function.
//
// The printer will not be deleted until all handles to it have been closed,
// including the handle passed to this function. This function will not
// close the provided handle.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/deleteprinter
func DeletePrinter(printer syscall.Handle) (err error) {
	result, _, errno := syscall.SyscallN(
		procDeletePrinter.Addr(),
		uintptr(printer))

	// If the function fails, the result is 0.
	if result == 0 {
		fmt.Printf("Error: %d\n", errno)
		return errno
	}

	return nil
}
