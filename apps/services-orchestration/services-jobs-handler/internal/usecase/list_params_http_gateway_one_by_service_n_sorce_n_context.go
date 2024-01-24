package usecase

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-jobs-handler/output"
	"log"
)

type ListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase struct {
	JobParamsHttpGatewayRepository entity.HttpGatewayParamsInterface
}

func NewListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase(
	repository entity.HttpGatewayParamsInterface,
) *ListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase {
	return &ListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase{
		JobParamsHttpGatewayRepository: repository,
	}
}

func (la *ListOneJobParamsHttpGatewayByServiceAndSourceAndContextUseCase) Execute(service string, source string, context string) (outputDTO.HttpGatewayParamsDTO, error) {
	item, err := la.JobParamsHttpGatewayRepository.FindOneByServiceAndSourceAndContext(service, source, context)
	if err != nil {
		return outputDTO.HttpGatewayParamsDTO{}, err
	}

    log.Println("item", item)

	dto := outputDTO.HttpGatewayParamsDTO{
		ID:             string(item.ID),
		Service:        item.Service,
		Source:         item.Source,
		Context:        item.Context,
		BaseUrl:        item.BaseUrl,
		UrlDomains:     ConvertEntityToUrlDomainDTO(item.UrlDomains),
		Headers:        item.Headers,
		EnableProxy:    item.EnableProxy,
		ProxyLoaders:   ConvertEntityToProxyLoaderDTO(item.ProxyLoaders),
		EnableCaptcha:  item.EnableCaptcha,
		CaptchaSolvers: ConvertEntityToCaptchaSolverDTO(item.CaptchaSolvers),
		ParamsID:       item.JobParamsID,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
	}

	return dto, nil

}
