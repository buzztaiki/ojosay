package main

import (
	"fmt"
	"os"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/jiro4989/ojosama"
)

func main() {
	ojo, err := ojosama.Convert(strings.Join(os.Args[1:], " "), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}

	say, err := cowsay.Say(
		ojo,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(say)
}
