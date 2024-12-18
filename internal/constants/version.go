package constants

import "fmt"

const (
	Major uint = 1
	Minor uint = 0
	Patch uint = 0
)

var Version = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
var APIVersion = fmt.Sprintf("/api/v%d", Major)
