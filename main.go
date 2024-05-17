package main

import (
	"fmt"
	"interface-struct-mock-mocktest-golang/internal"
)

func main() {
	x := internal.Phone{
		Markasi: "Iphone",
		Modeli:  "13 Pro",
	}

	fmt.Print(x.Gelis())
	fmt.Print(x.İcerik())
}
