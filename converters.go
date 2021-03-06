// Package goutils contains a collection of useful Golang utility methods and libraries
package goutils

import (
	// Standard lib
	"strconv"

	// Third-party
	log "github.com/sirupsen/logrus"
)

// Bool2String converts a bool to a string
func Bool2String(v bool) string {
	return strconv.FormatBool(v)
}

// Float642String converts a float64 to a string
func Float642String(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// IntSlice2StringSlice converts a slice of ints to a slice of strings
func IntSlice2StringSlice(s []int) []string {
	// Form return value
	ret := make([]string, 0)

	// Check for empty input
	if len(s) == 0 {
		return ret
	}

	// Loop through values
	for _, i := range s {
		// Only append non-empty strings
		if v := Int2String(i); v != "" {
			ret = append(ret, v)
		}
	}

	return ret
}

// Int2String converts an int to a string
func Int2String(v int) string {
	return strconv.Itoa(v)
}

// Int642String converts an int64 to a string
func Int642String(v int64) string {
	return strconv.Itoa(int(v))
}

// Interface2String attempts to determine the underlying type of an interface and returns it as a string
func Interface2String(i interface{}) string {
	// Attempt to cast attribute based on it's underlying type
	switch t := i.(type) {
	case float64:
		return Float642String(i.(float64))
	case int:
		return Int2String(i.(int))
	case int64:
		return Int642String(i.(int64))
	case string:
		return i.(string)
	default:
		// Log unsupported type
		log.WithField("type", t).Warn("Interface is of unsupported type")

		return ""
	}
}

// MapFromInterface type-asserts interfaces as a map[string]interface{}
// so that other methods can more-easily access it's properties
// NOTE: Other underlying types not currently supported
func MapFromInterface(i interface{}) map[string]interface{} {
	// TO-DO: Add recovery

	return i.(map[string]interface{})
}

// String2Bool converts a string to a bool
func String2Bool(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		// Log conversion error
		log.WithFields(log.Fields{
			"string": v,
			"error":  err.Error(),
		}).Warn("Error converting string to bool")

		return false
	}

	return b
}

// String2Float64 converts a string to a float64
func String2Float64(v string) float64 {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		// Log conversion error
		log.WithFields(log.Fields{
			"string": v,
			"error":  err.Error(),
		}).Warn("Error converting string to float64")

		return 0.0
	}

	return f
}

// String2Int converts a string to an int
func String2Int(v string) int {
	return int(String2Int64(v))
}

// String2Int64 converts a string to an int64
func String2Int64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		// Log conversion error
		log.WithFields(log.Fields{
			"string": v,
			"error":  err.Error(),
		}).Warn("Error converting string to int64")

		return 0
	}

	return i
}
