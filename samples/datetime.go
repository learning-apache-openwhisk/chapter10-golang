package main

import "time"

// Datetime format date/time with a format string
func Datetime(args map[string]interface{}) map[string]interface{} {
	now := time.Now()
	fmt, ok := args["format"].(string)
	if ok {
		return mkMap("result", now.Format(fmt))
	}
	return mkMap("error", "no format")
}
