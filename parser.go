package main

import (
	"github.com/alexflint/go-arg"
	"github.com/pgilad/jtl-parse/jtl"
	"github.com/pgilad/jtl-parse/output"
)

const DefaultOutputFormat = "json"

func main() {
	var args struct {
		Filename string `arg:"positional,required"`
		Output   string `arg:"-o" help:"specify the output type, valid options: csv,xml,json"`
	}
	args.Output = DefaultOutputFormat
	arg.MustParse(&args)

	outputStream := getOutputStream(args.Output)

	data := make(chan interface{})
	go jtl.Decode(args.Filename, data)
	outputStream(data)
}

func getOutputStream(outputType string) func(data <-chan interface{}) {
	switch outputType {
	case "csv":
		return output.CSV
	case "xml":
		return output.XML
	case "json":
		return output.JSON
	default:
		panic("Unknown output type " + outputType)
	}
}
