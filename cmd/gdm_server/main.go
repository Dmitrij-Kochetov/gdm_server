package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

func main() {
	res, err := git.PlainClone("some_prj", false, &git.CloneOptions{
		URL:      "https://gitlab.com/ImpError/httpserv",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("[!] Error: %v", err)
	}

	fmt.Printf("%v", res)
}
