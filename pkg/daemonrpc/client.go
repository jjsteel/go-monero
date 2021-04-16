package daemonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	EndpointJsonRPC = "/json_rpc"
	VersionJsonRPC  = "2.0"
)

type Client struct {
	http *http.Client
	url  *url.URL
}

type ClientOptions struct {
	HTTPClient *http.Client
}

type ClientOption func(o *ClientOptions)

func WithHTTPClient(v *http.Client) func(o *ClientOptions) {
	return func(o *ClientOptions) {
		o.HTTPClient = v
	}
}

func NewHTTPClient(verbose bool) *http.Client {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	if verbose == true {
		client.Transport = NewDumpTransport(http.DefaultTransport)
	}

	return client

}

func NewClient(address string, opts ...ClientOption) (*Client, error) {
	options := &ClientOptions{
		HTTPClient: NewHTTPClient(false),
	}

	for _, opt := range opts {
		opt(options)
	}

	parsedAddress, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("url parse: %w", err)
	}

	return &Client{
		url:  parsedAddress,
		http: options.HTTPClient,
	}, nil
}

// ResponseEnvelope wraps all responses from the RPC server.
//
type ResponseEnvelope struct {
	Id      string      `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// RequestEnvelope wraps all requests made to the RPC server.
//
type RequestEnvelope struct {
	Id      string      `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func (c *Client) JsonRPC(method string, params interface{}, response interface{}) error {
	url := *c.url
	url.Path = EndpointJsonRPC

	b, err := json.Marshal(&RequestEnvelope{
		Id:      "0",
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequest("GET", url.String(), bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("new req '%s': %w", url.String(), err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("non-2xx status code: %d", resp.StatusCode)
	}

	rpcResponseBody := &ResponseEnvelope{
		Result: response,
	}

	if err := json.NewDecoder(resp.Body).Decode(rpcResponseBody); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	if rpcResponseBody.Error.Code != 0 || rpcResponseBody.Error.Message != "" {
		return fmt.Errorf("rpc error: code=%d message=%s",
			rpcResponseBody.Error.Code,
			rpcResponseBody.Error.Message,
		)
	}

	return nil
}
