package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	in, out := os.Args[1], os.Args[2]
	f, err := os.Open(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()
	env, err := godotenv.Parse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	outF, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for k, v := range env {
		_, err = fmt.Fprintf(outF, "%s: \"%s\"\n", strings.ToLower(k), v)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
