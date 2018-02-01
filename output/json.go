package output

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/pgilad/jtl-parse/jtl"
)

func JSON(output <-chan interface{}) {
	os.Stdout.Write([]byte("{"))
	first := true
	for {
		element, more := <-output
		if !more {
			os.Stdout.Write([]byte("]}}"))
			break
		}
		switch elementType := element.(type) {
		case jtl.TestResults:
			// add opening testResults tag
			str := fmt.Sprintf(`"testResults": { "version": "%v", "samples": [`, elementType.Version)
			os.Stdout.Write([]byte(str))
		default:
			parsed, err := json.MarshalIndent(element, "", "    ")
			if err != nil {
				fmt.Printf("Could not marshal xml: %v\n", err)
			}
			if !first {
				os.Stdout.Write([]byte(","))
			} else {
				first = false
			}
			os.Stdout.Write(parsed)
		}
	}
}
