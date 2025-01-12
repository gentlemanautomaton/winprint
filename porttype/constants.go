package porttype

// Printer port types.
const (
	Write       = 0x0001 // PORT_TYPE_WRITE
	Read        = 0x0002 // PORT_TYPE_READ
	Redirected  = 0x0004 // PORT_TYPE_REDIRECTED
	NetAttached = 0x0008 // PORT_TYPE_NET_ATTACHED
)
