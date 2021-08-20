package main

import (
	"fmt"
	"os"
	"os/user"
	"vore/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	} 
	fmt.Printf("Hello %s, this is the Vore Programming Language\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}