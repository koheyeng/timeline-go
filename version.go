package version

import (
	"fmt"
	"runtime"
)

const Version = "v0.0.1"

var UserAgent = fmt.Sprintf("timeline distance aggregator/%s (%s)", Version, runtime.Version())
