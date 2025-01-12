package accessrights

// Format maps flags to their string representations.
type Format map[Mask]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	AdministerPrinter:        "AdministerPrinter",
	UsePrinter:               "UsePrinter",
	LimitedPrinterManagement: "LimitedPrinterManagement",
	FullPrinterAccess:        "FullPrinterAccess",
	AdministerJob:            "AdministerJob",
	ReadSpooledJobData:       "ReadSpooledJobData",
	FullJobAccess:            "FullJobAccess",
	AdministerServer:         "AdministerServer",
	EnumerateServer:          "EnumerateServer",
	FullServerAccess:         "FullServerAccess",
}
