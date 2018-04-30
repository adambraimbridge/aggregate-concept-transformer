package concept

import (
	"bytes"
	"io/ioutil"
	"net/http"

	logger "github.com/Financial-Times/go-logger"
)

type mockHTTPClient struct {
	resp       string
	statusCode int
	err        error
	called     []string
}

func init() {
	logger.InitLogger("test-aggregate-concept-transformer", "debug")
}

func (c *mockHTTPClient) Do(req *http.Request) (resp *http.Response, err error) {
	c.called = append(c.called, req.URL.String())
	cb := ioutil.NopCloser(bytes.NewReader([]byte(c.resp)))
	return &http.Response{Body: cb, StatusCode: c.statusCode}, c.err
}