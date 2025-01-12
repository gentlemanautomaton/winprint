package spoolerapi

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winprint/driverinfo"
	"github.com/gentlemanautomaton/winprint/internal/utf16conv"
)

var (
	procEnumPrinterDrivers              = modwinspool.NewProc("EnumPrinterDriversW")
	procUploadPrinterDriverPackage      = modspoolss.NewProc("UploadPrinterDriverPackage")
	procInstallPrinterDriverFromPackage = modspoolss.NewProc("InstallPrinterDriverFromPackage")
	procDeletePrinterDriver             = modwinspool.NewProc("DeletePrinterDriverExW")
)

// EnumPrinterDrivers returns information about printer drivers installed on
// the local system. It calls the EnumPrinterDrivers windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/enumprinterdrivers
func EnumPrinterDrivers[T driverinfo.Level](server, environment string) (drivers []T, err error) {
	// Marshal arguments as UTF-16.
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
	if err != nil {
		return nil, err
	}

	utf16Environment, err := utf16conv.StringToOptionalPtr(environment)
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
		level    = driverinfo.ID[T]()
		infoSize = driverinfo.Size[T]()
	)

	for attempt := 0; ; attempt++ {
		var bufferSize, entries uint32
		result, _, errno := syscall.SyscallN(
			procEnumPrinterDrivers.Addr(),
			uintptr(unsafe.Pointer(utf16Server)),
			uintptr(unsafe.Pointer(utf16Environment)),
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
			return nil, fmt.Errorf("failed to enumerate print drivers: the returned buffer size is insufficient (needed %d but got %d)", needed, len(buffer))
		}

		// Unmarshal the raw data into a slice of drivers of the requested
		// info level.
		drivers = make([]T, entries)
		for i := 0; i < len(drivers); i++ {
			data := buffer[i*infoSize : (i+1)*infoSize]
			if err := driverinfo.Unmarshal(data, &drivers[i]); err != nil {
				return nil, fmt.Errorf("failed to enumerate print drivers: entry %d: %w", i, err)
			}
		}

		return drivers, nil
	}
}

// UploadPrinterDriverPackage adds a printer driver package to the windows
// driver store using the provided INF file. It calls the
// UploadPrinterDriverPackage windows API function.
//
// The server, environment, and window arguments are optional.
//
// If server is empty, the driver package will be added to the driver store
// on the local machine.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/uploadprinterdriverpackage
func UploadPrinterDriverPackage(server, infPath, environment string, flags uint32, window syscall.Handle) (driverStoreInfPath string, err error) {
	// Marshal arguments as UTF-16.
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
	if err != nil {
		return "", err
	}

	utf16InfPath, err := utf16conv.StringToPtr(infPath)
	if err != nil {
		return "", err
	}

	utf16Environment, err := utf16conv.StringToOptionalPtr(environment)
	if err != nil {
		return "", err
	}

	// Allocate a buffer on the stack to receive the new inf path.
	var buffer [4096]uint16
	length := uint32(len(buffer))

	// Perform the system call.
	result, _, _ := syscall.SyscallN(
		procUploadPrinterDriverPackage.Addr(),
		uintptr(unsafe.Pointer(utf16Server)),
		uintptr(unsafe.Pointer(utf16InfPath)),
		uintptr(unsafe.Pointer(utf16Environment)),
		uintptr(flags),
		uintptr(window),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&length)))

	// If the function fails, the result is a non-zero error number.
	if result != 0 {
		return "", syscall.Errno(result)
	}

	return syscall.UTF16ToString(buffer[:length]), nil
}

// InstallPrinterDriverFromPackage adds a printer driver package to the windows
// driver store using the provided INF file. It calls the
// InstallPrinterDriverFromPackage windows API function.
//
// The server, environment, and window arguments are optional.
//
// If server is empty, the driver package will be added to the driver store
// on the local machine.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/installprinterdriverfrompackage
func InstallPrinterDriverFromPackage(server, infPath, driverName, environment string, flags uint32) (err error) {
	// Marshal arguments as UTF-16.
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
	if err != nil {
		return err
	}

	utf16InfPath, err := utf16conv.StringToOptionalPtr(infPath)
	if err != nil {
		return err
	}

	utf16DriverName, err := utf16conv.StringToPtr(driverName)
	if err != nil {
		return err
	}

	utf16Environment, err := utf16conv.StringToOptionalPtr(environment)
	if err != nil {
		return err
	}

	// Perform the system call.
	result, _, _ := syscall.SyscallN(
		procInstallPrinterDriverFromPackage.Addr(),
		uintptr(unsafe.Pointer(utf16Server)),
		uintptr(unsafe.Pointer(utf16InfPath)),
		uintptr(unsafe.Pointer(utf16DriverName)),
		uintptr(unsafe.Pointer(utf16Environment)),
		uintptr(flags))

	// If the function fails, the result is a non-zero error number.
	if result != 0 {
		return syscall.Errno(result)
	}

	return nil
}

// DeletePrinterDriver removes a printer driver package from the windows
// driver store. It calls the
// DeletePrinterDriverEx windows API function.
//
// The server and environment arguments are optional.
//
// If server is empty, the driver package will be removed from the driver store
// on the local machine.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/deleteprinterdriverex
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/21d247cd-66e2-4373-a187-b858c907e13f
func DeletePrinterDriver(server, environment, driverName string, flags, version uint32) (err error) {
	// Marshal arguments as UTF-16
	utf16Server, err := utf16conv.StringToOptionalPtr(server)
	if err != nil {
		return err
	}

	utf16Environment, err := utf16conv.StringToOptionalPtr(environment)
	if err != nil {
		return err
	}

	utf16DriverName, err := utf16conv.StringToPtr(driverName)
	if err != nil {
		return err
	}

	// Perform the system call.
	result, _, errno := syscall.SyscallN(
		procDeletePrinterDriver.Addr(),
		uintptr(unsafe.Pointer(utf16Server)),
		uintptr(unsafe.Pointer(utf16Environment)),
		uintptr(unsafe.Pointer(utf16DriverName)),
		uintptr(flags),
		uintptr(version))

	// If the function fails, the result is zero.
	if result == 0 {
		return errno
	}

	return nil
}
