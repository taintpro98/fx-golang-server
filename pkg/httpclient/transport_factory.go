package httpclient

import (
	"net/http"
	"sync"
	"time"
)

// manage http client transport per service

type HttpTransportFactory struct {
	transportMap map[string]*http.Transport
}

var instance *HttpTransportFactory
var once sync.Once

func GetHttpTransportFactoryInstance() *HttpTransportFactory {
	once.Do(
		func() {
			instance = &HttpTransportFactory{
				transportMap: make(map[string]*http.Transport),
			}
		},
	)
	return instance
}

func (f *HttpTransportFactory) GetTransport(service string) *http.Transport {
	if transport, ok := f.transportMap[service]; ok {
		return transport
	}
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
	f.transportMap[service] = transport
	return transport
}

func (f *HttpTransportFactory) Close() {
	for _, transport := range f.transportMap {
		transport.CloseIdleConnections()
	}
}
