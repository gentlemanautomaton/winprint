package spoolerapi

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/monitorinfo"
)

var (
	procEnumMonitors = modwinspool.NewProc("EnumMonitorsW")
	procXcvData      = modwinspool.NewProc("XcvDataW")
)

// EnumMonitors returns information about port monitors installed on
// the local system. It calls the EnumMonitors windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/enummonitors
func EnumMonitors[T monitorinfo.Level](server string) (monitors []T, err error) {
	// Marshal arguments as UTF-16.
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
	if err != nil {
		return nil, err
	}

	// Make up to 3 attempts to get the port monitor list:
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported size (first attempt)
	// 3: Using a dynamic buffer based on the reported size (second attempt)
	//
	// It is unlikely, but feasible, that the length could change between
	// calls if the port monitor list changes.

	var (
		scratch  [4096]byte
		buffer   = scratch[:]
		level    = monitorinfo.ID[T]()
		infoSize = monitorinfo.Size[T]()
	)

	for attempt := 0; ; attempt++ {
		var bufferSize, entries uint32
		result, _, errno := syscall.SyscallN(
			procEnumMonitors.Addr(),
			uintptr(unsafe.Pointer(utf16Server)),
			uintptr(level),
			uintptr(unsafe.Pointer(&buffer[0])),
			uintptr(len(buffer)),
			uintptr(unsafe.Pointer(&bufferSize)),
			uintptr(unsafe.Pointer(&entries)))

		// If the function fails, the return value is zero.
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
			return nil, fmt.Errorf("failed to enumerate printer port monitors: the returned buffer size is insufficient (needed %d but got %d)", needed, len(buffer))
		}

		// Unmarshal the raw data into a slice of printer port monitors of
		// the requested info level.
		monitors = make([]T, entries)
		for i := 0; i < len(monitors); i++ {
			data := buffer[i*infoSize : (i+1)*infoSize]
			if err := monitorinfo.Unmarshal(data, &monitors[i]); err != nil {
				return nil, fmt.Errorf("failed to enumerate printer port monitors: entry %d: %w", i, err)
			}
		}

		return monitors, nil
	}
}

// XcvData sends commands and receives data from a port monitor identified
// by the given port monitor handle. It calls the XcvData windows API function.
//
// The handle can be aquired by calling the OpenPrinter2 API function in a
// special way.
//
// https://learn.microsoft.com/en-us/previous-versions/ff564255(v=vs.85)
func XcvData(monitor syscall.Handle, operation string, input []byte) (output []byte, err error) {
	// Marshal arguments as UTF-16.
	utf16Operation, err := utf16conv.StringToOptionalPtr(operation)
	if err != nil {
		return nil, err
	}

	var inputPtr *byte
	if len(input) > 0 {
		inputPtr = &input[0]
	}

	// Make up to 3 attempts to get the port monitor list:
	//
	// 1: Using a fixed buffer allocated on the stack
	// 2: Using a dynamic buffer based on the reported size (first attempt)
	// 3: Using a dynamic buffer based on the reported size (second attempt)
	//
	// It is unlikely, but feasible, that the length could change between
	// calls if the port monitor list changes.

	var (
		scratch [4096]byte
		buffer  = scratch[:]
	)

	for attempt := 0; ; attempt++ {
		var bufferSize uint32
		var status uint32
		result, _, errno := syscall.SyscallN(
			procXcvData.Addr(),
			uintptr(monitor),
			uintptr(unsafe.Pointer(utf16Operation)),
			uintptr(unsafe.Pointer(inputPtr)),
			uintptr(len(input)),
			uintptr(unsafe.Pointer(&buffer[0])),
			uintptr(len(buffer)),
			uintptr(unsafe.Pointer(&bufferSize)),
			uintptr(unsafe.Pointer(&status)))

		// If the function fails, the return value is 0.
		if result == 0 {
			if attempt > 2 || errno != syscall.ERROR_INSUFFICIENT_BUFFER {
				return nil, errno
			}
			buffer = make([]byte, bufferSize)
			continue
		}

		// If the function call itself succeeds, the status returned holds the
		// result of the requested operation.
		if status != 0 {
			return nil, syscall.Errno(status)
		}

		// Trim the buffer down to the size that was filled with data.
		if bufferSize == 0 {
			output = nil
		} else {
			output = buffer[:bufferSize]
		}

		return output, nil
	}
}
