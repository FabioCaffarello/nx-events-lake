package input

import (
	sharedDTO "libs/golang/services/dtos/services-jobs-handler/shared"
)

type HttpGatewayParamsDTO struct {
	Service        string                       `json:"service"`
	Source         string                       `json:"source"`
	Context        string                       `json:"context"`
	BaseUrl        string                       `json:"base_url"`
	UrlDomains     []sharedDTO.UrlDomainDTO     `json:"url_domains"`
	Headers        map[string]string            `json:"headers"`
	EnableProxy    bool                         `json:"enable_proxy"`
	ProxyLoaders   []sharedDTO.ProxyLoaderDTO   `json:"proxy_loaders"`
	EnableCaptcha  bool                         `json:"enable_captcha"`
	CaptchaSolvers []sharedDTO.CaptchaSolverDTO `json:"captcha_solvers"`
}
