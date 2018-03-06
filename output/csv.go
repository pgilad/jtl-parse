package output

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/pgilad/jtl-parse/jtl"
)

func CSV(output <-chan interface{}) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// write csv header
	writer.Write([]string{"Label", "Timestamp", "Latency", "Users"})

	for {
		element, more := <-output
		if !more {
			break
		}
		switch elementType := element.(type) {
		case jtl.TestResults:
			continue
		case jtl.Sample:
			writeSample(elementType, writer)
		case jtl.HttpSample:
			writeHttpSample(elementType, writer)
		default:
			// unknown element
		}
	}
}

func writeHttpSample(sample jtl.HttpSample, writer *csv.Writer) {
	data := []string{*sample.Label, strconv.FormatUint(*sample.Timestamp, 10), strconv.Itoa(*sample.Latency), strconv.Itoa(*sample.NA)}
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}

func writeSample(sample jtl.Sample, writer *csv.Writer) {
	data := []string{*sample.Label, strconv.FormatUint(*sample.Timestamp, 10), strconv.Itoa(*sample.Latency), strconv.Itoa(*sample.NA)}
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}
