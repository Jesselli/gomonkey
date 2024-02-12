package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Jesselli/gomonkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
