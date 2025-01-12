package printerenum

// Printer enumeration flags.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/enumprinters
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-1
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/2a1fe8a4-e8be-4cf3-8b37-8d19f9a2edcd
const (
	Default            = 0x00000001 // PRINTER_ENUM_DEFAULT
	Local              = 0x00000002 // PRINTER_ENUM_LOCAL
	Connections        = 0x00000004 // PRINTER_ENUM_CONNECTIONS
	Name               = 0x00000008 // PRINTER_ENUM_NAME
	Remote             = 0x00000010 // PRINTER_ENUM_REMOTE (level 1 only)
	Shared             = 0x00000020 // PRINTER_ENUM_SHARED
	Network            = 0x00000040 // PRINTER_ENUM_NETWORK (level 1 only)
	ExpansionSupported = 0x00004000 // PRINTER_ENUM_EXPAND
	Container          = 0x00008000 // PRINTER_ENUM_CONTAINER
	Icon1              = 0x00010000 // PRINTER_ENUM_ICON1
	Icon2              = 0x00020000 // PRINTER_ENUM_ICON2
	Icon3              = 0x00040000 // PRINTER_ENUM_ICON3
	Icon4              = 0x00080000 // PRINTER_ENUM_ICON4 (reserved)
	Icon5              = 0x00100000 // PRINTER_ENUM_ICON5 (reserved)
	Icon6              = 0x00200000 // PRINTER_ENUM_ICON6 (reserved)
	Icon7              = 0x00400000 // PRINTER_ENUM_ICON7 (reserved)
	Icon8              = 0x00800000 // PRINTER_ENUM_ICON8
	Hide               = 0x01000000 // PRINTER_ENUM_HIDE
	CategoryAll        = 0x02000000 // PRINTER_ENUM_CATEGORY_ALL
	Category3D         = 0x04000000 // PRINTER_ENUM_CATEGORY_3D
)
