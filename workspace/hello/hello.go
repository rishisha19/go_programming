package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello"))

	fmt.Print("using new upper method: ")
	fmt.Println(stringutil.ToUpper("hello"))
}
