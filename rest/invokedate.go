package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func url(operation string) string {
	return fmt.Sprintf("%s/api/v1/namespaces/%s/%s",
		os.Getenv("__OW_API_HOST"),
		os.Getenv("__OW_NAMESPACE"),
		operation)
}

func auth() (string, string) {
	up := strings.Split(os.Getenv("__OW_API_KEY"), ":")
	return up[0], up[1]
}

func post(action string,
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

func whiskInvoke(action string, args map[string]interface{},
	blocking bool, result bool) (map[string]interface{}, error) {
	invoke := fmt.Sprintf("actions/%s?blocking=%t&result=%t",
		action, blocking, result)
	req, err := post(invoke, args)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// encode answer
	var objmap map[string]interface{}
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		return nil, err
	}
	return objmap, nil
}

// Main invoke date using
func Main(args map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	action, ok := args["action"].(string)
	if !ok {
		out["error"] = "no action"
		return out
	}
	message, ok := args["message"].(string)
	if !ok {
		out["error"] = "no message"
		return out
	}
	res, err := whiskInvoke(action, out, true, true)
	if err != nil {
		out["error"] = err.Error()
		return out
	}
	date, ok := res["date"].(string)
	if !ok {
		out["error"] = "no date"
		return out
	}
	out["result"] = fmt.Sprintf("%s %s", message, date)
	return out
}
