package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func whiskTrigger(trigger string,
	args map[string]interface{}) map[string]interface{} {
	invoke := fmt.Sprintf("triggers/%s", trigger)
	req, err := mkPost(invoke, args)
	if err != nil {
		return mkErr(err)
	}
	return doCall(req)
}

func mkGet(action string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url(action), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(auth())
	return req, nil
}

func whiskRetrieve(id string) map[string]interface{} {
	invoke := fmt.Sprintf("activations/%s", id)
	req, err := mkGet(invoke)
	if err != nil {
		return mkErr(err)
	}
	return doCall(req)
}

// Fire invoke sort using triggers then retrieve the result.
// It can be invoked with "trigger" to fire that trigger,
// and with the activationId to retrieve the result
func Fire(args map[string]interface{}) map[string]interface{} {

	// get the trigger
	trigger, ok := args["trigger"].(string)
	if !ok {
		return mkErr("no trigger defined")
	}

	// read the text argument
	text, ok := args["text"].(string)
	if !ok {
		return mkErr("no text")
	}

	// fire the trigger
	input := mkMap("lines", strings.Split(text, ","))
	log.Printf("invoking trigger=%s", trigger)
	res := whiskTrigger(trigger, input)

	// check if we have the activationId
	id, ok := res["activationId"]
	if !ok {
		return mkErr("trigger did not return an activationId")
	}

	// invoke the action specified to retrieve the result
	out := whiskRetrieve(res)
	response, ok := out["response"].(map[string]interface{})
	if !ok {
		return mkErr("no response")
	}
	return response["result"].(map[string]interface{})
}
