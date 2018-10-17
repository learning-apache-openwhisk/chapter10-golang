package main

// Hello is the main
func Hello(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "world"
	}
	res := make(map[string]interface{})
	res["hello"] = "Hello, " + name + " !"
	return res
}
