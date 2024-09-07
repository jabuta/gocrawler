package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	rawBaseURL := getArgs()
	fmt.Print(getHTML(rawBaseURL))
}

func getArgs() string {
	if len(os.Args) < 2 {
		fmt.Print("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Print("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("- 'starting crawl'\n- '%s   ", os.Args[1])
	return os.Args[1]
}
