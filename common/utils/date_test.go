package utils_test

import (
	"testing"
	"time"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestDateLayoutJava2Go ...
func TestDateLayoutJava2Go(t *testing.T) {
	java := "EEE, d MMM yyyy HH:mm:ss z"
	goLang := utils.DateLayoutJava2Go(java)
	assert.Equal(t, goLang, "Mon, 2 Jan 2006 15:04:05 MST")

	java = "yyyy-MM-dd'T'HH:mm:ssZ"
	goLang = utils.DateLayoutJava2Go(java)
	assert.Equal(t, goLang, "2006-01-02T15:04:05-0700")

	java = "yyyyMMddHHmmss"
	goLang = utils.DateLayoutJava2Go(java)
	assert.Equal(t, goLang, "20060102150405")
}

func TestDateFormat(t *testing.T) {
	java := "yyyyMMddHHmmss"
	thisTime, err := time.Parse("2006-01-02 15:04:05", "2021-10-01 12:56:13")
	assert.NilError(t, err)
	thisTimeFormatted := utils.DateFormat(thisTime, java)
	assert.Equal(t, thisTimeFormatted, "20211001125613")

	java = "HH"
	thisTime, err = time.Parse("2006-01-02 15:04:05", "2021-10-01 12:56:13")
	assert.NilError(t, err)
	thisTimeFormatted = utils.DateFormat(thisTime, java)
	assert.Equal(t, thisTimeFormatted, "12")
}
