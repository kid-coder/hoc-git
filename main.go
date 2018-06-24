package main

import (
	"fmt"

	"github.com/kid-coder/stringutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(stringutil.Reverse("dlroW olleH"))
	city := "Ha Noi"
	country, continent := location(city)
	fmt.Printf("I live in %s %s\n", country, continent)
}
