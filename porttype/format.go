package porttype

// Format maps flags to their string representations.
type Format map[Value]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	Write:       "Write",
	Read:        "Read",
	Redirected:  "Redirected",
	NetAttached: "NetAttached",
}
