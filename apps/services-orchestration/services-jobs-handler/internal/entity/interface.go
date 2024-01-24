package entity

type HttpGatewayParamsInterface interface {
	Save(httpGatewayParams *HttpGatewayParams) error
	FindAll() ([]*HttpGatewayParams, error)
	FindAllBySource(source string) ([]*HttpGatewayParams, error)
	FindAllByService(service string) ([]*HttpGatewayParams, error)
    FindOneByServiceAndSourceAndContext(service string, source string, contextEnv string) (*HttpGatewayParams, error)
	FindOneById(id string) (*HttpGatewayParams, error)
    Delete(id string) error
}

type HttpGatewayParamsVersionInterface interface {
	Save(httpGatewayParamsVersion *HttpGatewayParamsVersion) error
	Update(httpGatewayParams *HttpGatewayParams) error
}
