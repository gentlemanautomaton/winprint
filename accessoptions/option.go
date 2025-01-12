package accessoptions

// Option is an access option that can be supplied to spoolerapi.Open.
type Option interface {
	Apply(*Data)
}
