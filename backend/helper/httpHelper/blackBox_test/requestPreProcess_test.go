package blackBox

import (
	"backend/helper/httpHelper"
	"bytes"
	"errors"
	"net/http"
	"testing"
)

type StubHttpResponseWriter struct {
	header http.Header
}

func (writer StubHttpResponseWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func (writer StubHttpResponseWriter) WriteHeader(statusCode int) {
	return
}

func (writer StubHttpResponseWriter) Header() http.Header {
	return writer.header
}

type StubBody struct {
	bytes.Reader
	CloseError error
}

func (body *StubBody) Close() error {
	return body.CloseError
}

type StubBodyReadError struct {
}

func (body StubBodyReadError) Read(p []byte) (n int, err error) {
	return 0, errors.New("fsfsdf")
}

func (body StubBodyReadError) Close() error {
	panic("implement me")
}

func TestReqPreProcess(t *testing.T) {
	type A struct {
		A int
	}
	type B struct {
		A string
	}
	var tests = []struct {
		method       string
		actualMethod string
		data         string
		target       interface{}
		errorCode    int
	}{
		{http.MethodPost, http.MethodPost, `{"A":10}`, &A{}, 0},
		{http.MethodPost, http.MethodPost, `{"A":10}`, A{}, 500},
		{http.MethodPost, http.MethodPost, `{"A":10}`, nil, 0},
		{http.MethodPost, http.MethodDelete, `{"A":10}`, &A{}, 400},
		{http.MethodPost, http.MethodPost, `{"A":10}`, &B{}, 400},
	}

	for _, v := range tests {
		stubReq := http.Request{
			Method: v.actualMethod,
			Header: nil,
			Body: &StubBody{Reader: *bytes.NewReader([]byte(v.data)),
			}}

		_, rawData, err := httpHelper.ReqPreProcess(v.method, StubHttpResponseWriter{header: map[string][]string{}}, &stubReq, true, v.target)
		if !(bytes.Compare(rawData, []byte(v.data)) == 0) {
			t.Fatal("raw data not match")
		}
		if err == nil {
			if v.errorCode != 0 {
				t.Fatal("miss error")
			}
		} else if err.ErrorCode() != v.errorCode {
			t.Fatal("error code not match")
		}
	}

	//test when cannot read message
	stubReq := http.Request{
		Method: http.MethodPost,
		Header: nil,
		Body:   &StubBodyReadError{},
	}
	_, _, err := httpHelper.ReqPreProcess(http.MethodPost, StubHttpResponseWriter{header: map[string][]string{}}, &stubReq, true, &A{})
	if err.ErrorCode()!=500{
		t.Fatal("error code not match")
	}
}
