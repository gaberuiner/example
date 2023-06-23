package clnwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClientWrapper[Req any, Res any] struct {
	path    string
	ctx     context.Context
	request Req
	fn      func(ctx context.Context, req Req) (Res, error)
}

func New[Req any, Res any](ctx context.Context, path string, request Req, fn func(ctx context.Context, req Req) (Res, error)) *ClientWrapper[Req, Res] {
	return &ClientWrapper[Req, Res]{path: path, ctx: ctx, request: request, fn: fn}
}

func (c *ClientWrapper[Req, Res]) Wrap() (Res, error) {
	var Ret Res
	rawData, err := json.Marshal(&c.request)
	if err != nil {
		return Ret, fmt.Errorf("encode request %w", err)
	}
	ctx, fnCancel := context.WithTimeout(c.ctx, 5*time.Second)
	defer fnCancel()
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, c.path, bytes.NewBuffer(rawData))
	if err != nil {
		return Ret, fmt.Errorf("prepare request: %w", err)
	}
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return Ret, fmt.Errorf("do request: %w", err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return Ret, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}
	var response Res
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if httpResponse.StatusCode != http.StatusOK {
		return Ret, fmt.Errorf("decode request: %w", err)
	}

	return response, nil
}
