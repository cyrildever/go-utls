package utils

import (
	"strings"
	"time"
)

const (
	ISO_8601 = "yyyy-MM-dd'T'HH:mm:ss.SSSZ"
)

// DateFormat formats the passed time to the Java-notated layout, eg. `yyyy-MM-dd`
func DateFormat(t time.Time, layout string) string {
	goLayout := DateLayoutJava2Go(layout)
	return t.Format(goLayout)
}

// DateLayoutJava2Go transforms a Java notation to its Golang equivalent
func DateLayoutJava2Go(layout string) string {
	var goLayout string
	// Year
	if strings.Contains(layout, "yyyy") {
		goLayout = strings.Replace(layout, "yyyy", "2006", 1)
	} else if strings.Contains(layout, "yy") {
		goLayout = strings.Replace(layout, "yy", "06", 1)
	} else {
		goLayout = layout
	}
	// Month
	if strings.Contains(goLayout, "MMMM") {
		goLayout = strings.Replace(goLayout, "MMMM", "January", 1)
	} else if strings.Contains(goLayout, "MMM") {
		goLayout = strings.Replace(goLayout, "MMM", "Jan", 1)
	} else if strings.Contains(goLayout, "MM") {
		goLayout = strings.Replace(goLayout, "MM", "01", 1)
	}
	// Day of month
	if strings.Contains(goLayout, "dd") {
		goLayout = strings.Replace(goLayout, "dd", "02", 1)
	} else if strings.Contains(goLayout, "d") {
		goLayout = strings.Replace(goLayout, "d", "2", 1)
	}
	// Day of week
	if strings.Contains(goLayout, "EEEE") {
		goLayout = strings.Replace(goLayout, "EEEE", "Monday", 1)
	} else if strings.Contains(goLayout, "EEE") {
		goLayout = strings.Replace(goLayout, "EEE", "Mon", 1)
	}
	// Hour
	if strings.Contains(goLayout, "HH") {
		goLayout = strings.Replace(goLayout, "HH", "15", 1)
	} else if strings.Contains(goLayout, "KK") {
		goLayout = strings.Replace(goLayout, "KK", "03", 1)
	} else if strings.Contains(goLayout, "K") {
		goLayout = strings.Replace(goLayout, "K", "3", 1)
	}
	if strings.Contains(goLayout, " a") {
		goLayout = strings.Replace(goLayout, " a", " PM", 1)
	}
	// Minute
	if strings.Contains(goLayout, "mm") {
		goLayout = strings.Replace(goLayout, "mm", "04", 1)
	}
	// Second
	if strings.Contains(goLayout, "ss") {
		goLayout = strings.Replace(goLayout, "ss", "05", 1)
	}
	// Millisecond
	if strings.Contains(goLayout, "S") {
		goLayout = strings.ReplaceAll(goLayout, "S", "0")
	}
	// Time zone
	if strings.Contains(goLayout, "'T'") {
		goLayout = strings.Replace(goLayout, "'T'", "T", 1)
	}
	if strings.Contains(goLayout, "Z") {
		goLayout = strings.Replace(goLayout, "Z", "-0700", 1)
	} else if strings.Contains(goLayout, " z") {
		goLayout = strings.Replace(goLayout, " z", " MST", 1)
	}
	return goLayout
}
