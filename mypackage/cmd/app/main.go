package main

import (
	"fmt"

	"github.com/tojixll/mypackage/mypackage"
)

func main() {
    if sum := mypackage.Add(1, 2); sum != 3 {
        panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
    }

    fmt.Println("Well done!")
}