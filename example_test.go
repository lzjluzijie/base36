package base36_test

import (
	"base36"
	"fmt"
)

func ExampleEncode() {
	//Encode example data with base36 encoding.
	data := []byte("Hello World!")
	encoded := base36.Encode(data)

	//Print encoded data.
	fmt.Println(string(encoded))

	//Output: 2678lx5gvmsv1dro9b5
}

func ExampleDecode() {
	//Decode example data with base36 encoding.
	encoded := "2678lx5gvmsv1dro9b5"
	decoded, err := base36.Decode(encoded)

	//Check errors.
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Print decoded data.
	fmt.Println(string(decoded))

	//Output: Hello World!
}
