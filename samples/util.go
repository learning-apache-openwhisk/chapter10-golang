package main

func mkMap(key string, any interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res[key] = any
	return res
}
