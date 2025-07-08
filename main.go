package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nicolesoto-dev/json-formatter-cli/formatter"
)

type Options struct {
	file     string
	validate bool
	minify   bool
	prettify bool
	output   bool
}

func main() {
	var options Options

	flag.StringVar(&options.file, "file", "", "Path to json file")
	flag.StringVar(&options.file, "f", "", "Shortcut for --file")

	flag.BoolVar(&options.minify, "minify", false, "Minify json output")
	flag.BoolVar(&options.minify, "m", false, "Shortcut for --minify")

	flag.BoolVar(&options.validate, "validate", false, "Validate json input")
	flag.BoolVar(&options.validate, "v", false, "Shortcut for --validate")

	flag.BoolVar(&options.output, "output", false, "Output json to stdout")
	flag.BoolVar(&options.output, "o", false, "Shortcut for --output")

	flag.BoolVar(&options.prettify, "prettify", false, "Prettify json output")
	flag.BoolVar(&options.prettify, "p", false, "Shortcut for --prettify")

	flag.Parse()

	if options.file == "" {
		log.Fatal("Please provide a --file path")
	}

	result, err := formatter.ReadFile(options.file, formatter.FlagOpts{
		Validate: options.validate,
		Minify:   options.minify,
		Prettify: options.prettify,
		Output:   options.output,
	})

	if err != nil {
		log.Fatal(err)
	}

	if options.output {
		fmt.Println(result)
	}
}
