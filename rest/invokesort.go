package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func url(operation string) string {
	if operation[:1] == "/" {
		return fmt.Sprintf("%s/api/v1/namespaces%s",
			os.Getenv("__OW_API_HOST"),
			operation)
	}
	return fmt.Sprintf("%s/api/v1/namespaces/%s/%s",
		os.Getenv("__OW_API_HOST"),
		os.Getenv("__OW_NAMESPACE"),
		operation)
}

func auth() (string, string) {
	up := strings.Split(os.Getenv("__OW_API_KEY"), ":")
	return up[0], up[1]
}

func mkMap(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func addMap(data map[string]interface{}, key string, value interface{}) map[string]interface{} {
	data[key] = value
	return data
}

func mkErr(err interface{}) map[string]interface{} {
	switch v := err.(type) {
	case error:
		return mkMap("error", v.Error())
	case string:
		return mkMap("error", v)
	default:
		return mkMap("error", fmt.Sprintf("%v", err))
	}
}

func mkPost(action string,
	args map[string]interface{}) (*http.Request, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url(action),
		bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(auth())
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func doCall(req *http.Request) map[string]interface{} {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return mkErr(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return mkErr(err)
	}
	// encode answer
	var objmap map[string]interface{}
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		return mkErr(err)
	}
	return objmap
}

func whiskInvoke(action string, args map[string]interface{},
	blocking bool, result bool) map[string]interface{} {
	invoke := fmt.Sprintf("actions/%s?blocking=%t&result=%t",
		action, blocking, result)
	req, err := mkPost(invoke, args)
	if err != nil {
		return mkErr(err)
	}
	return doCall(req)
}

// Invoke invokes the sort using the action parameter specified
func Invoke(args map[string]interface{}) map[string]interface{} {

	// retrieve action
	action, ok := args["action"].(string)
	if !ok {
		return mkErr("no action")
	}

	// prepare args
	text, ok := args["text"].(string)
	if !ok {
		return mkErr("no text")
	}
	input := strings.Split(text, ",")

	// invoke action
	res := whiskInvoke(action, mkMap("lines", input), true, true)
	log.Printf("%v", res)
	lines, ok := res["lines"].([]interface{})
	if !ok {
		return mkErr("cannot retrieve result")
	}

	// retrieve message
	result, ok := args["message"].(string)
	if !ok {
		result = ">>>"
	}
	for _, v := range lines {
		result += " " + v.(string)
	}
	return mkMap("result", result)
}
