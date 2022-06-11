package main

import (
	"flag"
	"fmt"
	"github.com/buger/jsonparser"
	"log"
	"os"
	"strings"
)

var readFile = flag.String("f", "package.json", "read file path")
var jsonPath = flag.String("p", "version", "json path")

func main() {
	flag.Parse()
	bs, err := os.ReadFile(*readFile)
	if err != nil {
		log.Fatal(err)
	}
	p := *jsonPath
	r, err := jsonparser.GetString(bs, strings.Split(p, ".")...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(r)
}
