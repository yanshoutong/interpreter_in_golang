package main

import (
	"fmt"
	"interpreter_in_golang/repl"
	"os"
	"os/user"
)

func main() {
	me, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", me.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
