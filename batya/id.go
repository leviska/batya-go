package batya

import "fmt"

// returns source network name
type Sourcer interface {
	Source() string
}

type ID interface {
	Sourcer
	fmt.Stringer
}
