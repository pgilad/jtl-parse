package main

import (
	"github.com/alexflint/go-arg"
	"github.com/pgilad/jtl-parse/jtl"
	"github.com/pgilad/jtl-parse/output"
)

const DefaultOutputFormat = "json"

type OutputType string

const (
	CSV  OutputType = "csv"
	XML  OutputType = "xml"
	JSON OutputType = "json"
)

func main() {
	var args struct {
		Filename string `arg:"positional,required"`
		Output   string `arg:"-o" help:"specify the output type, valid options: csv,xml,json"`
	}
	args.Output = DefaultOutputFormat
	arg.MustParse(&args)

	outputType := OutputType(args.Output)
	outputStream := getOutputStream(outputType)

	data := make(chan interface{})
	go jtl.Decode(args.Filename, data)
	outputStream(data)
}

func getOutputStream(outputType OutputType) func(data <-chan interface{}) {
	switch outputType {
	case CSV:
		return output.CSV
	case XML:
		return output.XML
	case JSON:
		return output.JSON
	default:
		panic("Unknown output type " + outputType)
	}
}
