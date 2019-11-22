package internal

import (
	"crypto/tls"
	"errors"
	"net/http"
)

//ClientTLS 请求封装
func ClientTLS(r *http.Request, tlsConfig *tls.Config) (*http.Response, error) {
	if tlsConfig == nil {
		return nil, errors.New("need call SetTLSConfig")
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := http.Client{
		Transport: transport,
	}
	return client.Do(r)
}
