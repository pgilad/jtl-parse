package jtl

import (
	"encoding/xml"
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
