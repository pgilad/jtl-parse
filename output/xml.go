package output

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/pgilad/jtl-parse/jtl"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

func XML(output <-chan interface{}) {
	os.Stdout.Write([]byte(Header))
	for {
		element, more := <-output
		if !more {
			os.Stdout.Write([]byte("</testResults>"))
			break
		}
		switch elementType := element.(type) {
		case jtl.TestResults:
			// add opening testResults tag
			os.Stdout.Write([]byte(`<testResults version="` + elementType.Version + `">`))
		default:
			parsed, err := xml.MarshalIndent(element, "", "    ")
			if err != nil {
				fmt.Printf("Could not marshal xml: %v\n", err)
			}
			os.Stdout.Write(parsed)
		}
	}
}
