package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

type orderedMap [][2]string

func (a orderedMap) Len() int           { return len(a) }
func (a orderedMap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a orderedMap) Less(i, j int) bool { return a[i][0] < a[j][0] }

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
	oMap := make(orderedMap, 0)
	for k, v := range env {
		oMap = append(oMap, [2]string{k, v})
	}
	sort.Sort(oMap)
	for _, e := range oMap {
		_, err = fmt.Fprintf(outF, "%s: \"%s\"\n", strings.ToLower(e[0]), e[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
