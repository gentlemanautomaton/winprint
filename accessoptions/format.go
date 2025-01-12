package accessoptions

// Format maps flags to their string representations.
type Format map[Flags]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	NoCache:      "NoCache",
	Cache:        "Cache",
	ClientChange: "ClientChange",
	NoClientData: "NoClientData",
}
