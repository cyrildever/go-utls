package content_type

import (
	"github.com/cyrildever/go-utls/common/utils"
)

// TODO Enrich as needed
var authorized = []string{APPLICATION_JSON, TEXT_PLAIN}

const (
	// APPLICATION_JSON ...
	APPLICATION_JSON = "application/json"

	// TEXT_PLAIN ...
	TEXT_PLAIN = "text/plain; charset=utf-8"
)

// IsAuthorized ...
func IsAuthorized(contentType string) bool {
	return utils.ContainString(authorized, contentType)
}
