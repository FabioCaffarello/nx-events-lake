package entity

import (
	"errors"
	"libs/golang/shared/go-id/config"
	configuuid "libs/golang/shared/go-id/config/uuid"
	"time"
)

var (
	ErrHttpGatewayParamsServiceEmpty        = errors.New("http gateway params service is empty")
	ErrHttpGatewayParamsSourceEmpty         = errors.New("http gateway params source is empty")
	ErrHttpGatewayParamsContextEmpty        = errors.New("http gateway params context is empty")
	ErrHttpGatewayParamsBaseUrlEmpty        = errors.New("http gateway params base url is empty")
	ErrHttpGatewayParamsHeadersEmpty        = errors.New("http gateway params headers is empty")
	ErrHttpGatewayParamsProxyLoadersEmpty   = errors.New("http gateway params proxy loaders is empty")
	ErrHttpGatewayParamsCaptchaSolversEmpty = errors.New("http gateway params captcha solvers is empty")
)

type UrlDomain struct {
	Url  string `bson:"url"`
	Name string `bson:"name"`
}

type ProxyLoader struct {
	Name     string `bson:"name"`
	Priority int    `bson:"priority"`
}

type CaptchaSolver struct {
	Name     string `bson:"name"`
	Priority int    `bson:"priority"`
}

type HttpGatewayParams struct {
	ID             config.ID         `bson:"id"`
	Service        string            `bson:"service"`
	Source         string            `bson:"source"`
	Context        string            `bson:"context"`
	BaseUrl        string            `bson:"base_url"`
	JobParamsID    configuuid.ID     `bson:"job_params_id"`
	UrlDomains     []UrlDomain       `bson:"url_domains"`
	Headers        map[string]string `bson:"headers"`
	EnableProxy    bool              `bson:"enable_proxy"`
	ProxyLoaders   []ProxyLoader     `bson:"proxy_loaders"`
	EnableCaptcha  bool              `bson:"enable_captcha"`
	CaptchaSolvers []CaptchaSolver   `bson:"captcha_solvers"`
	CreatedAt      string            `bson:"created_at"`
	UpdatedAt      string            `bson:"updated_at"`
}

func NewHttpGatewayParams(
	service string,
	source string,
	context string,
	baseUrl string,
	urlDomains []UrlDomain,
	headers map[string]string,
	enableProxy bool,
	proxyLoaders []ProxyLoader,
	enableCaptcha bool,
	captchaSolvers []CaptchaSolver,
) (*HttpGatewayParams, error) {
	httpGatewayParams := &HttpGatewayParams{
		ID:             config.NewID(service, source),
		Service:        service,
		Source:         source,
		Context:        context,
		BaseUrl:        baseUrl,
		UrlDomains:     urlDomains,
		Headers:        headers,
		EnableProxy:    enableProxy,
		ProxyLoaders:   proxyLoaders,
		EnableCaptcha:  enableCaptcha,
		CaptchaSolvers: captchaSolvers,
		CreatedAt:      time.Now().Format(time.RFC3339),
		UpdatedAt:      time.Now().Format(time.RFC3339),
	}
	err := httpGatewayParams.IsHttpGatewayParamsValid()
	if err != nil {
		return nil, err
	}
	err = httpGatewayParams.SetJobParamsID()
	if err != nil {
		return nil, err
	}
	return httpGatewayParams, nil
}

func (h *HttpGatewayParams) SetJobParamsID() error {
	jobParamsDataForVersionID := getHttpGatewayParamsToGenerateVersionID(h)

	jobParamsID, err := configuuid.GenerateConfigID(jobParamsDataForVersionID)
	if err != nil {
		return err
	}
	h.JobParamsID = jobParamsID
	return nil

}

func ConvertUrlDomainsToMap(UrlDomains []UrlDomain) []map[string]interface{} {
	mapUrlDomains := make([]map[string]interface{}, len(UrlDomains))
	for i, urlDomain := range UrlDomains {
		mapUrlDomains[i] = map[string]interface{}{
			"url":  urlDomain.Url,
			"name": urlDomain.Name,
		}
	}
	return mapUrlDomains
}

func ConvertProxyLoadersToMap(ProxyLoaders []ProxyLoader) []map[string]interface{} {
	mapProxyLoaders := make([]map[string]interface{}, len(ProxyLoaders))
	for i, proxyLoader := range ProxyLoaders {
		mapProxyLoaders[i] = map[string]interface{}{
			"name":     proxyLoader.Name,
			"priority": proxyLoader.Priority,
		}
	}
	return mapProxyLoaders
}

func ConvertCaptchaSolversToMap(CaptchaSolvers []CaptchaSolver) []map[string]interface{} {
	mapCaptchaSolvers := make([]map[string]interface{}, len(CaptchaSolvers))
	for i, captchaSolver := range CaptchaSolvers {
		mapCaptchaSolvers[i] = map[string]interface{}{
			"name":     captchaSolver.Name,
			"priority": captchaSolver.Priority,
		}
	}
	return mapCaptchaSolvers
}

func getHttpGatewayParamsToGenerateVersionID(h *HttpGatewayParams) map[string]interface{} {
	return map[string]interface{}{
		"service":         h.Service,
		"source":          h.Source,
		"context":         h.Context,
		"base_url":        h.BaseUrl,
		"url_domains":     ConvertUrlDomainsToMap(h.UrlDomains),
		"headers":         h.Headers,
		"enable_proxy":    h.EnableProxy,
		"proxy_loaders":   ConvertProxyLoadersToMap(h.ProxyLoaders),
		"enable_captcha":  h.EnableCaptcha,
		"captcha_solvers": ConvertCaptchaSolversToMap(h.CaptchaSolvers),
	}
}

func (h *HttpGatewayParams) IsHttpGatewayParamsValid() error {
	if h.Service == "" {
		return ErrHttpGatewayParamsServiceEmpty
	}
	if h.Source == "" {
		return ErrHttpGatewayParamsSourceEmpty
	}
	if h.Context == "" {
		return ErrHttpGatewayParamsContextEmpty
	}
	if h.BaseUrl == "" {
		return ErrHttpGatewayParamsBaseUrlEmpty
	}
	if h.Headers == nil {
		return ErrHttpGatewayParamsHeadersEmpty
	}
	if h.EnableProxy {
		if len(h.ProxyLoaders) == 0 {
			return ErrHttpGatewayParamsProxyLoadersEmpty
		}
	}
	if h.EnableCaptcha {
		if len(h.CaptchaSolvers) == 0 {
			return ErrHttpGatewayParamsCaptchaSolversEmpty
		}
	}
	return nil
}
