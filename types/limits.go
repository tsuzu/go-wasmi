package types

// Limits represents the size range of resizeable storage
type Limits struct {
	Min       uint32
	Max       uint32
	IgnoreMax bool
}
