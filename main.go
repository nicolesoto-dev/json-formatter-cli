package main

import (
	"flag"
	"fmt"
)

type Options struct {
	file     string
	validate bool
	minify   bool
}

func main() {
	var options Options

	flag.StringVar(&options.file, "file", "", "Path to json file")
	flag.BoolVar(&options.minify, "minify", false, "Minify json output")
	flag.BoolVar(&options.validate, "validate", false, "Validate json input")

	flag.Parse()
	input := flag.Arg(0)

	fmt.Println("Flags parsed:", options.file, input)
}
