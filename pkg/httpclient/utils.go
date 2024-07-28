package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

func CloneRequest(r *http.Request) (*http.Request, error) {
	// Sao chép cấu trúc cơ bản của http.Request
	clone := new(http.Request)
	*clone = *r

	// Sao chép URL
	if r.URL != nil {
		urlCopy := *r.URL
		clone.URL = &urlCopy
	}

	// Sao chép Header
	clone.Header = make(http.Header, len(r.Header))
	for k, v := range r.Header {
		clone.Header[k] = append([]string(nil), v...)
	}

	// Sao chép Body
	if r.Body != nil {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(r.Body); err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		clone.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
	}
	return clone, nil
}

func LogInfoRequest(ctx context.Context,
	end time.Duration,
	req http.Request,
	res http.Response,
	body []byte,
	err error,
) {

}
