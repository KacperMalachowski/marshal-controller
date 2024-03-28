package build

import (
	"time"
)

// Holds data about current binary build
var Version = "development"
var Time = time.Now().Format(time.UnixDate)
