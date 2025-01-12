package accessoptions

import "unsafe"

// Data holds a set of access flags that can be passed to spoolerapi.Open.
type Data struct {
	Flags Flags
}

// Bytes returns d in a format that can be consumed by spooler API calls.
func (d Data) Bytes() []byte {
	raw := rawData{
		Size:  uint32(unsafe.Sizeof(rawData{})),
		Flags: d.Flags,
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&raw)), unsafe.Sizeof(rawData{}))
}

type rawData struct {
	Size  uint32
	Flags Flags
}
