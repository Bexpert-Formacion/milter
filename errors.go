package milter

import "errors"

// pre-defined errors
var (
	ErrCloseSession = errors.New("stop current milter processing")
	ErrMacroNoData  = errors.New("macro definition with no data")
	ErrNoListenAddr = errors.New("no listen addr specified")
)
