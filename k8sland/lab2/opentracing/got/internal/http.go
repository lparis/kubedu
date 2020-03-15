package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

type (
	// Response a generic Quest response
	Response struct {
		Status string `json:"status"`
	}

	// Quest represents a knight's quest
	Quest struct {
		Knight string `json:"knight"`
	}
)

// Call a remote http service
func Call(ctx context.Context, method, url string, payload io.Reader, res interface{}) error {
	p := opentracing.SpanFromContext(ctx)
	s := opentracing.StartSpan(FuncName(0), opentracing.ChildOf(p.Context()))
	defer s.Finish()
	s.SetTag("http.method", method)
	s.SetTag("http.url", url)

	sx := opentracing.ContextWithSpan(ctx, s)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return SpanError(sx, err)
	}
	req.Header.Add("Content-Type", "application/json")

	opentracing.GlobalTracer().Inject(
		p.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return SpanError(sx, fmt.Errorf("remote call failed %s", err))
	}

	if resp.StatusCode != http.StatusOK {
		e, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return SpanError(sx, err)
		}
		return SpanError(sx, fmt.Errorf("call `%s failed (%d)-%s", url, resp.StatusCode, string(e)))
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return SpanError(sx, err)
	}
	return nil
}

// SpanError decorates span with error
func SpanError(ctx context.Context, err error) error {
	s := opentracing.SpanFromContext(ctx)
	s.SetTag("error", true)
	s.LogKV("event", "error", "message", err)
	return err
}

// WriteErrOut formulate err response and decorte span
func WriteErrOut(ctx context.Context, w http.ResponseWriter, err error) {
	s := opentracing.SpanFromContext(ctx)
	// Tag this span with an error. Tag error, http.status_code and log event/message
	!!YOUR_CODE!!
	http.Error(w, err.Error(), http.StatusExpectationFailed)
}
