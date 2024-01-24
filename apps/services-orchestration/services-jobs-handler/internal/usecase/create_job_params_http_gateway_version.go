package usecase

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-jobs-handler/output"
)

type CreateJobParamsHttpGatewayVersionUseCase struct {
	JobParamsHttpGatewayVersionRepository entity.HttpGatewayParamsVersionInterface
}

func NewCreateJobParamsHttpGatewayVersionUseCase(
	repository entity.HttpGatewayParamsVersionInterface,
) *CreateJobParamsHttpGatewayVersionUseCase {
	return &CreateJobParamsHttpGatewayVersionUseCase{
		JobParamsHttpGatewayVersionRepository: repository,
	}
}

func (ccu *CreateJobParamsHttpGatewayVersionUseCase) Execute(jobParamsHttpGateway outputDTO.HttpGatewayParamsDTO) (outputDTO.HttpGatewayParamsVersionDTO, error) {
	jobParamsHttpGatewayVersionEntity, err := entity.NewHttpGatewayParamsVersion(
		jobParamsHttpGateway.Service,
		jobParamsHttpGateway.Source,
		jobParamsHttpGateway.Context,
		jobParamsHttpGateway.BaseUrl,
		ConvertUrlDomainDTOToEntity(jobParamsHttpGateway.UrlDomains),
		jobParamsHttpGateway.Headers,
		jobParamsHttpGateway.EnableProxy,
		ConvertProxyLoaderDTOToEntity(jobParamsHttpGateway.ProxyLoaders),
		jobParamsHttpGateway.EnableCaptcha,
		ConvertCaptchaSolverDTOToEntity(jobParamsHttpGateway.CaptchaSolvers),
		jobParamsHttpGateway.ParamsID,
		jobParamsHttpGateway.CreatedAt,
		jobParamsHttpGateway.UpdatedAt,
	)
	if err != nil {
		return outputDTO.HttpGatewayParamsVersionDTO{}, err
	}

	err = ccu.JobParamsHttpGatewayVersionRepository.Save(jobParamsHttpGatewayVersionEntity)
	if err != nil {
		return outputDTO.HttpGatewayParamsVersionDTO{}, err
	}

	dto := outputDTO.HttpGatewayParamsVersionDTO{
		ID:       string(jobParamsHttpGatewayVersionEntity.ID),
		Versions: ConvertEntityToUseCaseHttpGatewayParamsVersion(jobParamsHttpGatewayVersionEntity.Versions),
	}

	return dto, nil
}
