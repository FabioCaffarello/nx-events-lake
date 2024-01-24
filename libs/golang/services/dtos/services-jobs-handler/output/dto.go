package output

import (
	sharedDTO "libs/golang/services/dtos/services-jobs-handler/shared"
)

type HttpGatewayParamsDTO struct {
	ID             string                       `json:"id"`
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
    ParamsID       string                       `json:"job_params_id"`
	CreatedAt      string                       `json:"created_at"`
	UpdatedAt      string                       `json:"updated_at"`
}

type HttpGatewayParamsVersionData struct {
    JobParamsID string                `json:"job_params_id"`
    Params      *HttpGatewayParamsDTO `json:"http_gateway_params"`
}

type HttpGatewayParamsVersionDTO struct {
    ID       string                       `json:"id"`
    Versions []HttpGatewayParamsVersionData `json:"versions"`
}
