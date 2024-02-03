package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/proxy"
)

type GetTorProxyRotateUseCase struct {
	UrlCheck  string
	ProxyAddr string
}

func NewGetTorProxyRotateUseCase() *GetTorProxyRotateUseCase {
	return &GetTorProxyRotateUseCase{
		UrlCheck:  "http://httpbin.org/ip",
		ProxyAddr: "tor:9050",
	}
}

func (u *GetTorProxyRotateUseCase) Execute() (map[string]interface{}, error) {
	dialer, err := proxy.SOCKS5("tcp", u.ProxyAddr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	resp, err := httpClient.Get(u.UrlCheck)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
