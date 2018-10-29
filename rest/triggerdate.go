package main

import (
	"os"
	"fmt"
	"net/http"
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

func fireThenRetrieve(trigger string, args map[string]interface{}) map[string]interface{} {
	res := whiskTrigger(trigger, args)
	// check if we have the activationId
	if _, ok := res["activationId"]; !ok {
		return mkErr("cannot invoke trigger")
	}
	// call myself to retrieve the result
	me := os.Getenv("__OW_ACTION_NAME")
	return whiskInvoke(me, res, true, true)
}

// FireTrigger invoke date using triggers and retrieve the result
// It can be invoked with the trigger to fire it
// And with the activationId to retrieve the result
func FireTrigger(args map[string]interface{}) map[string]interface{} {
	id, ok := args["activationId"].(string)
	if ok {
		return whiskRetrieve(id)
	}
	trigger, ok := args["trigger"].(string)
	if ok {
		return fireThenRetrieve(trigger, args)
	}
	return mkErr("no tirgger defined")
}
