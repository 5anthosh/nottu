package note

import "store/security/xss"

const (
	nameMaxLen    = 100
	contentMaxLen = 5000
)

var (
	nameXSSPolicy = xss.StrictPolicy
)
