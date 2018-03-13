package jtl

import (
	"encoding/json"
	"encoding/xml"
)

type TestResults struct {
	XMLName xml.Name `xml:"testResults"`
	Version string   `xml:"version,attr"`
}

type InnerChild struct {
	Value string `xml:",chardata" json:"value"`
	Class string `xml:"class,attr" json:"class"`
}

type AssertionResult struct {
	Name           string   `xml:"name,omitempty" json:"name,omitempty"`
	Failure        string   `xml:"failure,omitempty" json:"failure,omitempty"`
	Error          string   `xml:"error,omitempty" json:"error,omitempty"`
	FailureMessage []string `xml:"failureMessage,omitempty" json:"failureMessage,omitempty"`
}

type XmlAttr []xml.Attr

type SampleData struct {
	// sample attributes
	Bytes           *int    `xml:"by,attr,omitempty" json:"by,omitempty"`
	ConnectTime     *int    `xml:"ct,attr,omitempty" json:"ct,omitempty"`
	DataEncoding    *string `xml:"de,attr,omitempty" json:"de,omitempty"`
	DataType        *string `xml:"dt,attr,omitempty" json:"dt,omitempty"`
	ElapsedTime     *int    `xml:"t,attr,omitempty" json:"t,omitempty"`
	ErrorCount      *int    `xml:"ec,attr,omitempty" json:"ec,omitempty"`
	HostName        *string `xml:"hn,attr,omitempty" json:"hn,omitempty"`
	IdleTime        *int    `xml:"it,attr,omitempty" json:"it,omitempty"`
	Label           *string `xml:"lb,attr,omitempty" json:"lb,omitempty"`
	Latency         *int    `xml:"lt,attr,omitempty" json:"lt,omitempty"`
	NA              *int    `xml:"na,attr,omitempty" json:"na,omitempty"`
	NG              *int    `xml:"ng,attr,omitempty" json:"ng,omitempty"`
	ResponseCode    *string `xml:"rc,attr,omitempty" json:"rc,omitempty"`
	ResponseMessage *string `xml:"rm,attr,omitempty" json:"rm,omitempty"`
	SampleCount     *int    `xml:"sc,attr,omitempty" json:"sc,omitempty"`
	SentBytes       *int    `xml:"sby,attr,omitempty" json:"sby,omitempty"`
	Success         *bool   `xml:"s,attr,omitempty" json:"s,omitempty"`
	ThreadName      *string `xml:"tn,attr,omitempty" json:"tn,omitempty"`
	Timestamp       *uint64 `xml:"ts,attr,omitempty" json:"ts,omitempty"`

	// nested elements
	AssertionResults []AssertionResult `xml:"assertionResult" json:"assertionResults,omitempty"`
	Cookies          *InnerChild       `xml:"cookies,omitempty" json:"cookies,omitempty"`
	JavaNetURL       *string           `xml:"java.net.URL,omitempty" json:"java.net.URL,omitempty"`
	Method           *InnerChild       `xml:"method,omitempty" json:"method,omitempty"`
	QueryString      *InnerChild       `xml:"queryString,omitempty" json:"queryString,omitempty"`
	RedirectLocation *InnerChild       `xml:"redirectLocation,omitempty" json:"redirectLocation,omitempty"`
	RequestHeader    *InnerChild       `xml:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	ResponseData     *InnerChild       `xml:"responseData,omitempty" json:"responseData,omitempty"`
	ResponseFile     *InnerChild       `xml:"responseFile,omitempty" json:"responseFile,omitempty"`
	ResponseHeader   *InnerChild       `xml:"responseHeader,omitempty" json:"responseHeader,omitempty"`
	SamplerData      *InnerChild       `xml:"samplerData,omitempty" json:"samplerData,omitempty"`

	// catch all other attributes
	Attributes []XmlAttr `xml:",any,attr" json:"customVariables"`
}

type NestedSamples struct {
	Samples     []Sample     `xml:"sample" json:"samples,omitempty"`
	HttpSamples []HttpSample `xml:"httpSample" json:"httpSamples,omitempty"`
}

type Sample struct {
	XMLName xml.Name `xml:"sample" json:"-"`
	SampleData
	NestedSamples
}
type HttpSample struct {
	XMLName xml.Name `xml:"httpSample" json:"-"`
	SampleData
	NestedSamples
}

func (sample Sample) MarshalJSON() ([]byte, error) {
	type Alias Sample
	return json.Marshal(&struct {
		Type string `json:"type"`
		Alias
	}{
		Type:  "sample",
		Alias: (Alias)(sample),
	})
}

func (sample HttpSample) MarshalJSON() ([]byte, error) {
	type Alias HttpSample
	return json.Marshal(&struct {
		Type string `json:"type"`
		Alias
	}{
		Type:  "httpSample",
		Alias: (Alias)(sample),
	})
}

func (s XmlAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{
		Name:  s[0].Name.Local,
		Value: s[0].Value,
	})
}
