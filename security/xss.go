package security

import (
	"github.com/microcosm-cc/bluemonday"
)

var (
	StrictXSSPolicy = bluemonday.StrictPolicy()
)
