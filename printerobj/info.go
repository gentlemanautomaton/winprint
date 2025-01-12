package printerobj

/*
// Info holds information about printer-related objects.
//
// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-1
type Info struct {
	Flags       printerenum.Flags
	Description string
	Name        string
	Comment     string
}

// UnmarshalBinary unmarshals the given level 1 printer data into info.
func (info *Info) UnmarshalBinary(data []byte) error {
	if len(data) < int(unsafe.Sizeof(rawInfo{})) {
		return errors.New("insufficient data for printerinfo.Level1 unmarshaling")
	}

	raw := (*rawInfo)(unsafe.Pointer(&data[0]))

	info.Flags = raw.Flags
	info.Description = utf16conv.PtrToString(raw.Description)
	info.Name = utf16conv.PtrToString(raw.Name)
	info.Comment = utf16conv.PtrToString(raw.Comment)

	return nil
}

type rawInfo struct {
	Flags       printerenum.Flags
	Description *uint16
	Name        *uint16
	Comment     *uint16
}
*/
