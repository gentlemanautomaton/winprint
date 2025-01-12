package utf16conv

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// StringToPtr converts the given string s to a null-terminated UTF-16 string
// and returns a pointer to it.
//
// If s is empty it returns a pointer to an empty UTF-16 string.
//
// If s contains a NUL byte at any location, it returns an error.
func StringToPtr(s string) (*uint16, error) {
	return syscall.UTF16PtrFromString(s)
}

// StringToOptionalPtr converts the given string s to a null-terminated UTF-16
// string and returns a pointer to it.
//
// If s is empty it returns a nil pointer.
//
// If s contains a NUL byte at any location, it returns an error.
func StringToOptionalPtr(s string) (*uint16, error) {
	if s == "" {
		return nil, nil
	}
	return syscall.UTF16PtrFromString(s)
}

// PtrToString converts the null-terminated UTF-16 pointer p to a string
// and returns it. The string returned is limited to 65535 characters.
func PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	return syscall.UTF16ToString(unsafe.Slice(p, 65535))
}

// PtrToStringSlice converts a series of null-terminated UTF-16 strings
// at pointer p to a slice of strings and returns it. There series is
// terminated by an empty UTF-16 string. The total lenght of the series
// is limited to 65535 characters.
func PtrToStringSlice(p *uint16) []string {
	if p == nil {
		return nil
	}
	return utf16ToSplitString(unsafe.Slice(p, 65535))
}

// utf16ToSplitString splits a set of null-separated UTF-16 characters and
// returns a slice of substrings between those separators.
func utf16ToSplitString(s []uint16) []string {
	var values []string
	cut := 0
	for i, v := range s {
		if v == 0 {
			if i-cut > 0 {
				values = append(values, string(utf16.Decode(s[cut:i])))
			}
			cut = i + 1
		}
	}
	if cut < len(s) {
		values = append(values, string(utf16.Decode(s[cut:])))
	}
	return values
}
