package accessrights

import "strings"

// Mask holds a set of access control flags that describe a desired level of
// access when accessing a print server, printer, port or monitor.
type Mask uint32

// Match returns true if v contains all of the flags specified by c.
func (v Mask) Match(c Mask) bool {
	return v&c == c
}

// String returns a string representation of the flags using a default
// separator and format.
func (v Mask) String() string {
	return v.Join("|", FormatGo)
}

// Join returns a string representation of the flags using the given
// separator and format.
func (v Mask) Join(sep string, format Format) string {
	if s, ok := format[v]; ok {
		return s
	}

	var matched []string
	for i := 0; i < 32; i++ {
		flag := Mask(1 << uint32(i))
		if v.Match(flag) {
			if s, ok := format[flag]; ok {
				matched = append(matched, s)
			}
		}
	}

	return strings.Join(matched, sep)
}
