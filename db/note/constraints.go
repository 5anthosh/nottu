package note

import "github.com/5anthosh/nottu/security"

const (
	nameMaxLen    = 100
	contentMaxLen = 5000
)

var (
	nameXSSPolicy = security.StrictXSSPolicy
)
