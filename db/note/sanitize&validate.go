package note

import (
	"strings"
)

func validate(title string, content string) string {
	if len(title) > nameMaxLen {
		return titleMaxLenValidation
	}
	if len(content) > contentMaxLen {
		return contentMaxLenValidation
	}
	return ok
}

func sanitize(title *string, content *string) {
	*title = nameXSSPolicy.Sanitize(strings.TrimSpace(*title))
	*content = nameXSSPolicy.Sanitize(strings.TrimSpace(*content))
}

func proccessInputs(title string, content string) string {
	sanitize(&title, &content)
	return validate(title, content)
}
