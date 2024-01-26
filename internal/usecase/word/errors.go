package word

import "fmt"

var (
	ErrNothingUpdated = fmt.Errorf("nothing to update")
	ErrNothingDeleted = fmt.Errorf("nothing to delete")
)
