package printerstatus

// Printer statuses.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/1625e9d9-29e4-48f4-b83d-3bd0fdaea787
const (
	Paused                 = 0x00000001 // PRINTER_STATUS_PAUSED
	Error                  = 0x00000002 // PRINTER_STATUS_ERROR
	PendingDeletion        = 0x00000004 // PRINTER_STATUS_PENDING_DELETION
	PaperJam               = 0x00000008 // PRINTER_STATUS_PAPER_JAM
	PaperOut               = 0x00000010 // PRINTER_STATUS_PAPER_OUT
	ManualFeed             = 0x00000020 // PRINTER_STATUS_MANUAL_FEED
	PaperProblem           = 0x00000040 // PRINTER_STATUS_PAPER_PROBLEM
	Offline                = 0x00000080 // PRINTER_STATUS_OFFLINE
	ActiveIO               = 0x00000100 // PRINTER_STATUS_IO_ACTIVE
	Busy                   = 0x00000200 // PRINTER_STATUS_BUSY
	Printing               = 0x00000400 // PRINTER_STATUS_PRINTING
	OutputBinFull          = 0x00000800 // PRINTER_STATUS_OUTPUT_BIN_FULL
	NotAvailable           = 0x00001000 // PRINTER_STATUS_NOT_AVAILABLE
	Waiting                = 0x00002000 // PRINTER_STATUS_WAITING
	Processing             = 0x00004000 // PRINTER_STATUS_PROCESSING
	Initializing           = 0x00008000 // PRINTER_STATUS_INITIALIZING
	WarmingUp              = 0x00010000 // PRINTER_STATUS_WARMING_UP
	TonerLow               = 0x00020000 // PRINTER_STATUS_TONER_LOW
	NoToner                = 0x00040000 // PRINTER_STATUS_NO_TONER
	PageNotPrintable       = 0x00080000 // PRINTER_STATUS_PAGE_PUNT
	UserInterventionNeeded = 0x00100000 // PRINTER_STATUS_USER_INTERVENTION
	OutOfMemory            = 0x00200000 // PRINTER_STATUS_OUT_OF_MEMORY
	DoorOpen               = 0x00400000 // PRINTER_STATUS_DOOR_OPEN
	ServerUnknown          = 0x00800000 // PRINTER_STATUS_SERVER_UNKNOWN
	PowerSave              = 0x01000000 // PRINTER_STATUS_POWER_SAVE
	ServerOffline          = 0x02000000 // PRINTER_STATUS_SERVER_OFFLINE
	DriverUpdateNeeded     = 0x04000000 // PRINTER_STATUS_DRIVER_UPDATE_NEEDED
)
