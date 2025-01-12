package portproto

import "strconv"

// Type is a TCP/IP port protocol type.
type Type uint32

// String returns a string representation of the type.
func (t Type) String() string {
	switch t {
	case RawTCP:
		return "RAW"
	case LPR:
		return "LPR"
	default:
		return "unknown-" + strconv.Itoa(int(t))
	}
}
