package usecase

import (
    "apps/services-orchestration/services-jobs-handler/internal/entity"
    inputDTO "libs/golang/services/dtos/services-jobs-handler/input"
    outputDTO "libs/golang/services/dtos/services-jobs-handler/output"
)

type UpdateJobParamsHttpGatewayUseCase struct {
    JobParamsHttpGatewayRepository entity.HttpGatewayParamsInterface
}

func NewUpdateJobParamsHttpGatewayUseCase(
    repository entity.HttpGatewayParamsInterface,
) *UpdateJobParamsHttpGatewayUseCase {
    return &UpdateJobParamsHttpGatewayUseCase{
        JobParamsHttpGatewayRepository: repository,
    }
}

func (ccu *UpdateJobParamsHttpGatewayUseCase) Execute(jobParamsHttpGateway inputDTO.HttpGatewayParamsDTO) (outputDTO.HttpGatewayParamsDTO, error) {
    jobParamsHttpGatewayEntity, err := entity.NewHttpGatewayParams(
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
    )
    if err != nil {
        return outputDTO.HttpGatewayParamsDTO{}, err
    }

    err = ccu.JobParamsHttpGatewayRepository.Save(jobParamsHttpGatewayEntity)
    if err != nil {
        return outputDTO.HttpGatewayParamsDTO{}, err
    }

    dto := outputDTO.HttpGatewayParamsDTO{
        ID:             string(jobParamsHttpGatewayEntity.ID),
        Service:        jobParamsHttpGatewayEntity.Service,
        Source:         jobParamsHttpGatewayEntity.Source,
        Context:        jobParamsHttpGatewayEntity.Context,
        BaseUrl:        jobParamsHttpGatewayEntity.BaseUrl,
        UrlDomains:     ConvertEntityToUrlDomainDTO(jobParamsHttpGatewayEntity.UrlDomains),
        Headers:        jobParamsHttpGatewayEntity.Headers,
        EnableProxy:    jobParamsHttpGatewayEntity.EnableProxy,
        ParamsID:       jobParamsHttpGatewayEntity.JobParamsID,
        ProxyLoaders:   ConvertEntityToProxyLoaderDTO(jobParamsHttpGatewayEntity.ProxyLoaders),
        EnableCaptcha:  jobParamsHttpGatewayEntity.EnableCaptcha,
        CaptchaSolvers: ConvertEntityToCaptchaSolverDTO(jobParamsHttpGatewayEntity.CaptchaSolvers),
        CreatedAt:      jobParamsHttpGatewayEntity.CreatedAt,
        UpdatedAt:      jobParamsHttpGatewayEntity.UpdatedAt,
    }

    return dto, nil
}
