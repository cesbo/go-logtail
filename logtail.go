package logtail

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog"
)

type Logtail struct {
	url   *url.URL
	token string
}

var (
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:    10 * time.Second,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    4,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        90 * time.Second,
			ResponseHeaderTimeout:  2 * time.Second,
			ExpectContinueTimeout:  1 * time.Second,
			MaxResponseHeaderBytes: 2 * 1024,
			ForceAttemptHTTP2:      true,
		},
	}
)

func NewLogtail(token string) *Logtail {
	return &Logtail{
		url: &url.URL{
			Scheme: "https",
			Host:   "in.logtail.com",
		},
		token: "Bearer " + token,
	}
}

func (l *Logtail) Write(body []byte) (int, error) {
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
	request.URL = l.url
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", l.token)

	response, err := client.Do(request)
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return 0, fmt.Errorf("log send: %w", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != 202 {
		return 0, fmt.Errorf("log send: %s", response.Status)
	}

	return len(body), nil
}

func (l *Logtail) WriteLevel(level zerolog.Level, body []byte) (int, error) {
	return l.Write(body)
}

func (l *Logtail) NewLogger() zerolog.Logger {
	logger := zerolog.New(l)
	logger = logger.Hook(&LogtailTimestamp{})
	return logger
}
