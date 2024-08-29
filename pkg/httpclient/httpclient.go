package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"fx-golang-server/pkg/constants"
	"fx-golang-server/pkg/e"

	"io"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type CallbackFunc func(body []byte) error

type DoRequestParam struct {
	Request      *http.Request
	Headers      map[string]string
	ErrorHandler CallbackFunc
}

type HttpClient struct {
	Client  *http.Client
	Headers map[string]string
}

func (c HttpClient) DoRequest(ctx context.Context, param DoRequestParam, output interface{}, backupOutput *string) error {
	param.Request.Header.Add("Content-Type", "application/json")
	for key, value := range c.Headers {
		param.Request.Header.Set(key, value)
	}
	for key, value := range param.Headers {
		param.Request.Header.Set(key, value)
	}
	
	requestID := ctx.Value(constants.TraceID)
	if requestID != nil {
		param.Request.Header.Add(constants.XRequestID, fmt.Sprintf("%s", requestID))
	}

	reqClone, err := CloneRequest(param.Request)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("clone request error")
	}
	start := time.Now()
	res, err := c.Client.Do(param.Request)
	end := time.Since(start)

	if err != nil {
		LogInfoRequest(ctx, end, *reqClone, http.Response{}, nil, nil)
		if backupOutput != nil {
			tmp := err.Error()
			*backupOutput = tmp
		}
		var netErr net.Error
		ok := errors.As(err, &netErr)
		if ok && netErr.Timeout() {
			return e.ErrTimeout
		}
		return err
	}
	if res == nil {
		LogInfoRequest(ctx, end, *reqClone, http.Response{}, nil, nil)
		if backupOutput != nil {
			*backupOutput = e.ErrNilResponse.Msg
		}
		return e.ErrNilResponse
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("close body error")
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	LogInfoRequest(ctx, end, *reqClone, *res, body, err)
	if backupOutput != nil {
		tmp := string(body)
		*backupOutput = tmp
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusAccepted {
		var errorOutput e.CustomErr

		if param.ErrorHandler != nil {
			tmp := param.ErrorHandler(body)
			okErr := errors.As(tmp, &errorOutput)
			if !okErr {
				return tmp
			}
		}
		errorOutput.HttpStatusCode = res.StatusCode
		return errorOutput
	}
	if output != nil {
		if err := json.Unmarshal(body, output); err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("unmarshal response body error")
			return err
		}
	}
	return nil
}
