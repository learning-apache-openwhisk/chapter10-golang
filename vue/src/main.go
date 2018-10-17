package main

import "app"

/*
func Hello() map[string]interface{} {
	map := make(map[string]interface{})
	return JSONResponse(map)
}*/

// Main is the main action
func Main(args map[string]interface{}) map[string]interface{} {
	// get the path
	path, ok := args["__ow_path"].(string)
	if ok && path != "" {
		return app.WebResponse(path)
	}
	return map[string]interface{}{
		"body": `
	<script>
	  location.href += "/index.html"
	</script>
		`,
	}
}
