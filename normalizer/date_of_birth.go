package normalizer

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/metakeule/fmtdate"
)

const (
	// FRENCH_FORMAT ...
	FRENCH_FORMAT = "DD/MM/YYYY"

	// ISO_FORMAT ...
	ISO_FORMAT = "YYYYMMDD"

	// TIMESTAMP ...
	TIMESTAMP = "timestamp"

	// TIMESTAMP_MILLIS ...
	TIMESTAMP_MILLIS = "timestamp_millis"
)

// DateOfBirth returns a normalized date using the `params` arguments,
// the latter being a list of optional arguments to use to format the output appropriately:
// - the first item is the string format of the input string (defaut to ISO format: `YYYYMMDD`);
// - the second item is the string format for the output (default to French date: `DD/MM/YYYY`).
// The input format could be a `timestamp` or a `timestamp_millis`.
var DateOfBirth VariadicNormalizer = func(input string, params ...string) (string, error) {
	input = strings.TrimSpace(input)
	inputFormat := ISO_FORMAT
	if len(params) > 0 {
		switch strings.ToLower(params[0]) {
		case TIMESTAMP:
			inputFormat = TIMESTAMP
		case TIMESTAMP_MILLIS:
			inputFormat = TIMESTAMP_MILLIS
		default:
			inputFormat = strings.ToUpper(params[0])
		}
	}
	outputFormat := FRENCH_FORMAT
	if len(params) > 1 {
		outputFormat = strings.ToUpper(params[1])
	}
	var d time.Time
	if inputFormat == TIMESTAMP {
		i, e := strconv.ParseInt(input, 10, 64)
		if e != nil {
			return "", e
		}
		d = time.Unix(i, 0)
	} else if inputFormat == TIMESTAMP_MILLIS {
		millis := len(input) - 3
		if millis < 0 {
			return "", errors.New("invalid timestamp in milliseconds")
		}
		nanos, e := strconv.ParseInt(input[millis:], 10, 64)
		if e != nil {
			return "", e
		}
		secs, e := strconv.ParseInt(input[:millis], 10, 64)
		if e != nil {
			return "", e
		}
		d = time.Unix(secs, nanos)
	} else {
		parsed, e := fmtdate.Parse(inputFormat, input)
		if e != nil {
			return "", e
		}
		d = parsed
	}
	return fmtdate.Format(outputFormat, d), nil
}
