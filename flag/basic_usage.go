package main

import (
	"flag"
	"fmt"
	"os"
)

var count int
var total int

func init() {
	flag.IntVar(&count, "count", 1, "help message for count")
	flag.IntVar(&total, "total", 10, "help message for total")
}

func main() {

	flag.Parse()

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	flag.PrintDefaults()

}
