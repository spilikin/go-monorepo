package main

import (
	"fmt"

	"github.com/spilikin/go-monorepo/module"
)

func main() {
	fmt.Println("This is v2")
	fmt.Println(module.MessageFromModule())
}
