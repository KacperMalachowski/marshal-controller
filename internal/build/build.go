package build

import (
	"fmt"
	"time"
)

// Holds data about current binary build
var s = time.Now()
var Version = fmt.Sprintf("dev-%d.%.02d.%.02d%.02d%.02d%.02d", s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second())
var Time = time.Now().Format(time.UnixDate)
