package objname

import "strings"

// Type identifies an object name type.
type Type string

// DetectType returns the type of printer object described by name.
//
// If name does not contain an explicit or recognized object type, the
// Unspecified type is returned.
func DetectType(name string) Type {
	switch {
	case strings.Contains(name, ",XcvMonitor"):
		return Monitor
	case strings.Contains(name, ",XcvPort"):
		return Port
	case strings.Contains(name, ",Job"):
		return Job
	case strings.HasPrefix(name, `\`):
		return Server
	default:
		return Unspecified
	}
}
