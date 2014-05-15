package word

// Id represents word identity.
type Id uint32

const (
	NIL Id = ^Id(0) // An invalid word.
)
