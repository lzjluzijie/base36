package base36_test

import (
	"fmt"

	"github.com/lzjluzijie/base36"
)

func ExampleEncode() {
	//Encode example data with base36 encoding.
	data := []byte("Hello World!")
	encoded := base36.Encode(data)

	//Print encoded data.
	fmt.Println(encoded)

	//Output: 2678lx5gvmsv1dro9b5
}

func ExampleDecode() {
	//Decode example data with base36 encoding.
	encoded := "2678lx5gvmsv1dro9b5"
	decoded := base36.Decode(encoded)

	//Check errors.
	// if decoded == nil {
	// 	fmt.Println("Base36 decode error")
	// 	return
	// }

	//Print decoded data.
	fmt.Println(string(decoded))

	//Output: Hello World!
}
