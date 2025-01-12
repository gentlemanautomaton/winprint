package portproto

// TCP/IP port protocol types.
//
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn/2d256a4b-819c-4b84-8e8f-b1edb1dcf0c2
const (
	RawTCP = 0x00000001 // PROTOCOL_RAWTCP_TYPE
	LPR    = 0x00000002 // PROTOCOL_LPR_TYPE
)
