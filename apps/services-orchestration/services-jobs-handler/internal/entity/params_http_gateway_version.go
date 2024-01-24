package entity

import (
	"errors"
	"libs/golang/shared/go-id/config"
	configuuid "libs/golang/shared/go-id/config/uuid"
)

var (
	ErrHttpGatewayParamsVersionIDEmpty       = errors.New("http gateway params version id is empty")
	ErrHttpGatewayParamsVersionVersionsEmpty = errors.New("http gateway params version versions is empty")
)

type HttpGatewayParamsData struct {
	JobParamsID configuuid.ID      `bson:"job_params_id"`
	Params      *HttpGatewayParams `bson:"http_gateway_params"`
}

type HttpGatewayParamsVersion struct {
	ID       config.ID               `bson:"id"`
	Versions []HttpGatewayParamsData `bson:"versions"`
}

func NewHttpGatewayParamsVersion(
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
	createdAt string,
	updatedAt string,
	jobParamsId configuuid.ID,
) (*HttpGatewayParamsVersion, error) {
	httpGatewayParamsData := HttpGatewayParamsData{
		JobParamsID: jobParamsId,
		Params: &HttpGatewayParams{
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
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		},
	}
	httpGatewayParamsVersion := &HttpGatewayParamsVersion{
		ID:       httpGatewayParamsData.Params.ID,
		Versions: []HttpGatewayParamsData{httpGatewayParamsData},
	}
	err := httpGatewayParamsVersion.IsHttpGatewayParamsVersionValid()
	if err != nil {
		return nil, err
	}

	return httpGatewayParamsVersion, nil
}

func (h *HttpGatewayParamsVersion) IsHttpGatewayParamsVersionValid() error {
	if h.ID == "" {
		return ErrHttpGatewayParamsVersionIDEmpty
	}
	if len(h.Versions) == 0 {
		return ErrHttpGatewayParamsVersionVersionsEmpty
	}
	return nil
}
