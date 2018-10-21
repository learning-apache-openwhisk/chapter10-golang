package parser

import "fmt"

func print(op string, a, b int, err error) {
	if err == nil {
		fmt.Printf("%s(%d,%d)\n", op, a, b)
	} else {
		fmt.Printf("err: %s\n", err.Error())
	}
}
func ExampleParse() {
	print(Parse("2 + 2"))
	print(Parse("2 + 3"))
	print(Parse("2"))
	print(Parse("3 * a"))
	print(Parse("3 / 2"))
	// Output:
	// +(2,2)
	// +(2,3)
	// err: not enough args
	// err: strconv.Atoi: parsing "a": invalid syntax
	// /(3,2)
}
