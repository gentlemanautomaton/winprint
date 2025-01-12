package accessrights

// Printer access rights.
const (
	AdministerPrinter        = 0x00000004 // PRINTER_ACCESS_ADMINISTER
	UsePrinter               = 0x00000008 // PRINTER_ACCESS_USE
	LimitedPrinterManagement = 0x00000040 // PRINTER_ACCESS_MANAGE_LIMITED
	FullPrinterAccess        = 0x000F000C // PRINTER_ALL_ACCESS
)

// Job access rights.
const (
	AdministerJob      = 0x00000010 // JOB_ACCESS_ADMINISTER
	ReadSpooledJobData = 0x00000020 // JOB_ACCESS_READ
	FullJobAccess      = 0x000F0030 // JOB_ALL_ACCESS
)

// Print server access rights.
const (
	AdministerServer = 0x00000001 // SERVER_ACCESS_ADMINISTER
	EnumerateServer  = 0x00000002 // SERVER_ACCESS_ENUMERATE
	FullServerAccess = 0x000F0003 // SERVER_ALL_ACCESS
)
