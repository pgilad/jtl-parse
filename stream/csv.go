package stream

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/pgilad/jtl-parse/jtl"
)

func OutputCSV(output <-chan interface{}) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// write csv header
	writer.Write([]string{"Label", "Timestamp", "Latency"})

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
	data := []string{*sample.Label, strconv.Itoa(*sample.Timestamp), strconv.Itoa(*sample.Latency)}
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}

func writeSample(sample jtl.Sample, writer *csv.Writer) {
	data := []string{*sample.Label, strconv.Itoa(*sample.Timestamp), strconv.Itoa(*sample.Latency)}
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}
