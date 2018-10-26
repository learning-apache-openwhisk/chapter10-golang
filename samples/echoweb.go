package main

// Main is an echo function returing the arguments in JSON format
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["body"] = args
	res["status"] = "200"
	res["headers"] = map[string]string{
		"Content-Type": "application/json",
	}
	return res
}
