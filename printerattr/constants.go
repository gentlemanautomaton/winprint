package printerattr

// Printer attributes.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/1625e9d9-29e4-48f4-b83d-3bd0fdaea787
const (
	Queued                        = 0x00000001 // PRINTER_ATTRIBUTE_QUEUED
	Direct                        = 0x00000002 // PRINTER_ATTRIBUTE_DIRECT
	Default                       = 0x00000004 // PRINTER_ATTRIBUTE_DEFAULT
	Shared                        = 0x00000008 // PRINTER_ATTRIBUTE_SHARED
	Network                       = 0x00000010 // PRINTER_ATTRIBUTE_NETWORK
	Hidden                        = 0x00000020 // PRINTER_ATTRIBUTE_HIDDEN
	Local                         = 0x00000040 // PRINTER_ATTRIBUTE_LOCAL
	EnableDevQ                    = 0x00000080 // PRINTER_ATTRIBUTE_ENABLE_DEVQ
	KeepPrintedJobs               = 0x00000100 // PRINTER_ATTRIBUTE_KEEPPRINTEDJOBS
	DoCompleteFirst               = 0x00000200 // PRINTER_ATTRIBUTE_DO_COMPLETE_FIRST
	WorkOffline                   = 0x00000400 // PRINTER_ATTRIBUTE_WORK_OFFLINE
	BidirectionalCommunication    = 0x00000800 // PRINTER_ATTRIBUTE_ENABLE_BIDI
	RawOnly                       = 0x00001000 // PRINTER_ATTRIBUTE_RAW_ONLY
	Published                     = 0x00002000 // PRINTER_ATTRIBUTE_PUBLISHED
	Fax                           = 0x00004000 // PRINTER_ATTRIBUTE_FAX
	Redirected                    = 0x00008000 // PRINTER_ATTRIBUTE_TS
	PushedUser                    = 0x00020000 // PRINTER_ATTRIBUTE_PUSHED_USER
	PushedMachine                 = 0x00040000 // PRINTER_ATTRIBUTE_PUSHED_MACHINE
	Machine                       = 0x00080000 // PRINTER_ATTRIBUTE_MACHINE
	FriendlyName                  = 0x00100000 // PRINTER_ATTRIBUTE_FRIENDLY_NAME
	TerminalServicesGenericDriver = 0x00200000 // PRINTER_ATTRIBUTE_TS_GENERIC_DRIVER
	PerUser                       = 0x00400000 // PRINTER_ATTRIBUTE_PER_USER
	EnterpriseCloud               = 0x00800000 // PRINTER_ATTRIBUTE_ENTERPRISE_CLOUD
)
