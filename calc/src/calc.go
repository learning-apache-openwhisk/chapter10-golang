package main

import (
	"ops"
	"parser"
)

// Main is an action that calculates an expressions
func Main(args map[string]interface{}) map[string]interface{} {
	expr, ok := args["expr"].(string)
	if !ok {
		return mkMap("error", "no parameter expr")
	}
	op, a, b, err := parser.Parse(expr)
	if err != nil {
		return mkMap("error", err.Error())
	}
	switch op {
	case "+":
		return mkMap("result", ops.Add(a, b))
	case "*":
		return mkMap("result", ops.Mul(a, b))
	default:
		return mkMap("error", "Unsupported Operation")
	}
}
