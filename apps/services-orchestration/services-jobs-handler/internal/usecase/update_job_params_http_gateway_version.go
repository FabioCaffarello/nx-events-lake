package usecase

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-jobs-handler/output"
)

type UpdateJobParamsHttpGatewayVersionUseCase struct {
	JobParamsHttpGatewayVersionRepository entity.HttpGatewayParamsVersionInterface
}

func NewUpdateJobParamsHttpGatewayVersionUseCase(
	repository entity.HttpGatewayParamsVersionInterface,
) *UpdateJobParamsHttpGatewayVersionUseCase {
	return &UpdateJobParamsHttpGatewayVersionUseCase{
		JobParamsHttpGatewayVersionRepository: repository,
	}
}

func (uc *UpdateJobParamsHttpGatewayVersionUseCase) Execute(jobParamsHttpGateway outputDTO.HttpGatewayParamsDTO) (outputDTO.HttpGatewayParamsDTO, error) {
	jobParamsHttpGatewayVersionEntity := ConvertUseCaseToEntityHttpGatewayParamsVersion(jobParamsHttpGateway)
	err := uc.JobParamsHttpGatewayVersionRepository.Update(jobParamsHttpGatewayVersionEntity)
	if err != nil {
		return outputDTO.HttpGatewayParamsDTO{}, err
	}

	dto := outputDTO.HttpGatewayParamsDTO{
		ID:             string(jobParamsHttpGatewayVersionEntity.ID),
		Service:        jobParamsHttpGatewayVersionEntity.Service,
		Source:         jobParamsHttpGatewayVersionEntity.Source,
		Context:        jobParamsHttpGatewayVersionEntity.Context,
		BaseUrl:        jobParamsHttpGatewayVersionEntity.BaseUrl,
		UrlDomains:     ConvertEntityToUrlDomainDTO(jobParamsHttpGatewayVersionEntity.UrlDomains),
		Headers:        jobParamsHttpGatewayVersionEntity.Headers,
		EnableProxy:    jobParamsHttpGatewayVersionEntity.EnableProxy,
		ProxyLoaders:   ConvertEntityToProxyLoaderDTO(jobParamsHttpGatewayVersionEntity.ProxyLoaders),
		EnableCaptcha:  jobParamsHttpGatewayVersionEntity.EnableCaptcha,
		CaptchaSolvers: ConvertEntityToCaptchaSolverDTO(jobParamsHttpGatewayVersionEntity.CaptchaSolvers),
		ParamsID:       jobParamsHttpGatewayVersionEntity.JobParamsID,
		CreatedAt:      jobParamsHttpGatewayVersionEntity.CreatedAt,
		UpdatedAt:      jobParamsHttpGatewayVersionEntity.UpdatedAt,
	}

	return dto, nil
}
