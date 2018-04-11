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
	writer.Write([]string{"Label", "Timestamp", "Response Time", "Latency", "Users"})

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

func createCSVRow(sample interface{}) []string {
	// TODO: remove duplication
	switch s := sample.(type) {
	case jtl.Sample:
		return []string{
			*s.Label,
			strconv.FormatUint(*s.Timestamp, 10),
			strconv.Itoa(*s.ElapsedTime),
			strconv.Itoa(*s.Latency),
			strconv.Itoa(*s.NA),
		}
	case jtl.HttpSample:
		return []string{
			*s.Label,
			strconv.FormatUint(*s.Timestamp, 10),
			strconv.Itoa(*s.ElapsedTime),
			strconv.Itoa(*s.Latency),
			strconv.Itoa(*s.NA),
		}
	default:
		return []string{}
	}
}

func writeSample(sample jtl.Sample, writer *csv.Writer) {
	data := createCSVRow(sample)
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}

func writeHttpSample(sample jtl.HttpSample, writer *csv.Writer) {
	data := createCSVRow(sample)
	writer.Write(data)
	for _, row := range sample.Samples {
		writeSample(row, writer)
	}
	for _, row := range sample.HttpSamples {
		writeHttpSample(row, writer)
	}
}
