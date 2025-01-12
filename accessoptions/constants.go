package accessoptions

// Printer access option flags.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-option-flags
const (
	NoCache      Flags = 0x00000001 // PRINTER_OPTION_NO_CACHE
	Cache        Flags = 0x00000002 // PRINTER_OPTION_CACHE
	ClientChange Flags = 0x00000004 // PRINTER_OPTION_CLIENT_CHANGE
	NoClientData Flags = 0x00000008 // PRINTER_OPTION_NO_CLIENT_DATA
)
