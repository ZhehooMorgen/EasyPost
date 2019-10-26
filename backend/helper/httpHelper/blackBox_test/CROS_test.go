package blackBox_test

import (
	"backend/helper/httpHelper"
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

func (writer StubHttpResponseWriter)Header() http.Header{
	return writer.header
}

func TestCORS(t *testing.T) {
	httpHelper.CORS(StubHttpResponseWriter{map[string][]string{}})
}