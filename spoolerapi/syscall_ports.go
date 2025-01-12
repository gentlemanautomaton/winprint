package spoolerapi

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
	"github.com/gentlemanautomaton/winprint/portinfo"
)

var (
	procEnumPorts = modwinspool.NewProc("EnumPortsW")
)

// EnumPorts returns information about printer ports installed on
// the local system. It calls the EnumPorts windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/enumports
func EnumPorts[T portinfo.Level](server string) (ports []T, err error) {
	// Marshal arguments as UTF-16.
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
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
		level    = portinfo.ID[T]()
		infoSize = portinfo.Size[T]()
	)

	for attempt := 0; ; attempt++ {
		var bufferSize, entries uint32
		result, _, errno := syscall.SyscallN(
			procEnumPorts.Addr(),
			uintptr(unsafe.Pointer(utf16Server)),
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
			return nil, fmt.Errorf("failed to enumerate printer ports: the returned buffer size is insufficient (needed %d but got %d)", needed, len(buffer))
		}

		// Unmarshal the raw data into a slice of printer ports of the requested
		// info level.
		ports = make([]T, entries)
		for i := 0; i < len(ports); i++ {
			data := buffer[i*infoSize : (i+1)*infoSize]
			if err := portinfo.Unmarshal(data, &ports[i]); err != nil {
				return nil, fmt.Errorf("failed to enumerate printer ports: entry %d: %w", i, err)
			}
		}

		return ports, nil
	}
}
