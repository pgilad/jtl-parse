package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

type TestResults struct {
	XMLName xml.Name `xml:"testResults"`
	Version string   `xml:"version,attr"`
}

type InnerChild struct {
	Value string `xml:",chardata"`
	Class string `xml:"class,attr"`
}

type AssertionResult struct {
	Name           string `xml:"name,omitempty"`
	Failure        string `xml:"failure,omitempty"`
	Error          string `xml:"error,omitempty"`
	FailureMessage []byte `xml:"failureMessage,omitempty"`
}

type SampleData struct {
	// sample attributes
	Bytes           *int    `xml:"by,attr,omitempty"`
	ConnectTime     *int    `xml:"ct,attr,omitempty"`
	DataEncoding    *string `xml:"de,attr,omitempty"`
	DataType        *string `xml:"dt,attr,omitempty"`
	ElapsedTime     *int    `xml:"t,attr,omitempty"`
	ErrorCount      *int    `xml:"ec,attr,omitempty"`
	HostName        *string `xml:"hn,attr,omitempty"`
	IdleTime        *int    `xml:"it,attr,omitempty"`
	Label           *string `xml:"lb,attr,omitempty"`
	Latency         *int    `xml:"lt,attr,omitempty"`
	NA              *int    `xml:"na,attr,omitempty"`
	NG              *int    `xml:"ng,attr,omitempty"`
	ResponseCode    *string `xml:"rc,attr,omitempty"`
	ResponseMessage *string `xml:"rm,attr,omitempty"`
	SampleCount     *int    `xml:"sc,attr,omitempty"`
	SentBytes       *int    `xml:"sby,attr,omitempty"`
	Success         *bool   `xml:"s,attr,omitempty"`
	ThreadName      *string `xml:"tn,attr,omitempty"`
	Timestamp       *int    `xml:"ts,attr,omitempty"`

	// nested elements
	AssertionResults []AssertionResult `xml:"assertionResult"`
	Cookies          *InnerChild       `xml:"cookies,omitempty"`
	JavaNetURL       *string           `xml:"java.net.URL,omitempty"`
	Method           *InnerChild       `xml:"method,omitempty"`
	QueryString      *InnerChild       `xml:"queryString,omitempty"`
	RedirectLocation *InnerChild       `xml:"redirectLocation,omitempty"`
	RequestHeader    *InnerChild       `xml:"requestHeader,omitempty"`
	ResponseData     *InnerChild       `xml:"responseData,omitempty"`
	ResponseFile     *InnerChild       `xml:"responseFile,omitempty"`
	ResponseHeader   *InnerChild       `xml:"responseHeader,omitempty"`
	SamplerData      *InnerChild       `xml:"samplerData,omitempty"`

	// catch all other attributes
	Attributes []xml.Attr `xml:",any,attr"`
}

type NestedSamples struct {
	Samples     []Sample     `xml:"sample"`
	HttpSamples []HttpSample `xml:"httpSample"`
}

type Sample struct {
	XMLName xml.Name `xml:"sample"`
	SampleData
	NestedSamples
}

type HttpSample struct {
	XMLName xml.Name `xml:"httpSample"`
	SampleData
	NestedSamples
}

func outputCSV(output <-chan interface{}) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// write csv header
	writer.Write([]string{"label", "timestamp"})

	for {
		element, more := <-output
		if !more {
			break
		}
		switch elementType := element.(type) {
		case TestResults:
			continue
		case HttpSample:
			writer.Write([]string{*elementType.Label, string(*elementType.Timestamp)})
		case Sample:
			writer.Write([]string{*elementType.Label, string(*elementType.Timestamp)})
		default:
			// unknown element
		}
	}
}

func outputXML(output <-chan interface{}) {
	os.Stdout.Write([]byte(Header))
	for {
		element, more := <-output
		if !more {
			os.Stdout.Write([]byte("</testResults>"))
			break
		}
		switch elementType := element.(type) {
		case TestResults:
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

func streamDecodeJTL(filename string, output chan<- interface{}) {
	defer close(output)

	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		// Inspect the type of the token just read.
		switch elementType := token.(type) {
		case xml.StartElement:
			name := elementType.Name.Local
			if name == "testResults" {
				version := elementType.Attr[0].Value
				// send testResult tag
				testResults := TestResults{Version: version, XMLName: xml.Name{Space: "", Local: "testResults"}}
				output <- testResults
			} else if name == "sample" {
				var sample Sample
				err := decoder.DecodeElement(&sample, &elementType)
				if err != nil {
					log.Println("Could not decode element", err)
					// log this, but just skip
					break
				}

				// send sample
				output <- sample
			} else if name == "httpSample" {
				var sample HttpSample
				err := decoder.DecodeElement(&sample, &elementType)
				if err != nil {
					log.Println("Could not decode element", err)
					// log this, but just skip
					break
				}

				// send sample
				output <- sample
			} else {
				fmt.Println("Found an unknown element", name)
			}
		default:
		}
	}
}

func main() {
	filename := "data/short-assertion.jtl"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	output := make(chan interface{})
	go streamDecodeJTL(filename, output)
	outputXML(output)
}
