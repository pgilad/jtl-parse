package main

import (
	"os"

	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/pgilad/jtl-parser/jtl"
	"github.com/pgilad/jtl-parser/stream"
)

func main() {
	var args struct {
		Filename string `arg:"positional,required"`
		Output   string `arg:"-o" help:"specify the output type, valid options: csv,xml"`
	}
	args.Output = "csv"
	arg.MustParse(&args)

	output := make(chan interface{})
	go jtl.Decode(args.Filename, output)
	switch args.Output {
	case "csv":
		stream.OutputCSV(output)
	case "xml":
		stream.OutputXML(output)
	default:
		fmt.Println("Unknown output type " + args.Output)
		os.Exit(1)
	}
}
