package main

import "fmt"

func main() {
	var x interface{}
	x = "abc"
	x = 100
	x = true
	x = 99.99
	x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = "Aliqua pariatur sint eu proident amet incididunt veniam veniam labore non."
	// fmt.Println(x.(int) * 2)
	if val, ok := x.(int); ok {
		fmt.Println(val * 2)
	} else {
		fmt.Println("x is not int. Operation not allowed")
	}

	// type switch
	// x = 99
	// x = true
	x = 99.99
	switch val := x.(type) {
	case int:
		fmt.Println("x is a int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a boolean, !x =", !val)
	default:
		fmt.Println("unknown type : val = ", val)
	}

}
