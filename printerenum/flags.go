package printerenum

import "strings"

// Flags holds a set of printer enumeration flags that can be passed to
// spoolerapi.EnumPrinters and returned in printerinfo.Level1.
type Flags uint32

// Match returns true if v contains all of the flags specified by c.
func (v Flags) Match(c Flags) bool {
	return v&c == c
}

// String returns a string representation of the flags using a default
// separator and format.
func (v Flags) String() string {
	return v.Join("|", FormatGo)
}

// Join returns a string representation of the flags using the given
// separator and format.
func (v Flags) Join(sep string, format Format) string {
	if s, ok := format[v]; ok {
		return s
	}

	var matched []string
	for i := 0; i < 32; i++ {
		flag := Flags(1 << uint32(i))
		if v.Match(flag) {
			if s, ok := format[flag]; ok {
				matched = append(matched, s)
			}
		}
	}

	return strings.Join(matched, sep)
}
