package usecase

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-jobs-handler/output"
	sharedDTO "libs/golang/services/dtos/services-jobs-handler/shared"
	"log"
)

func ConvertUrlDomainDTOToEntity(urlDomains []sharedDTO.UrlDomainDTO) []entity.UrlDomain {
	var urlDomainsEntity []entity.UrlDomain
	for _, urlDomain := range urlDomains {
		urlDomainsEntity = append(urlDomainsEntity, entity.UrlDomain{
			Name: urlDomain.Name,
			Url:  urlDomain.Url,
		})
	}
	return urlDomainsEntity
}

func ConvertProxyLoaderDTOToEntity(proxyLoaders []sharedDTO.ProxyLoaderDTO) []entity.ProxyLoader {
	var proxyLoadersEntity []entity.ProxyLoader
	for _, proxyLoader := range proxyLoaders {
		proxyLoadersEntity = append(proxyLoadersEntity, entity.ProxyLoader{
			Name:     proxyLoader.Name,
			Priority: proxyLoader.Priority,
		})
	}
	return proxyLoadersEntity
}

func ConvertCaptchaSolverDTOToEntity(captchaSolvers []sharedDTO.CaptchaSolverDTO) []entity.CaptchaSolver {
	var captchaSolversEntity []entity.CaptchaSolver
	for _, captchaSolver := range captchaSolvers {
		captchaSolversEntity = append(captchaSolversEntity, entity.CaptchaSolver{
			Name:     captchaSolver.Name,
			Priority: captchaSolver.Priority,
		})
	}
	return captchaSolversEntity
}

func ConvertEntityToUrlDomainDTO(urlDomains []entity.UrlDomain) []sharedDTO.UrlDomainDTO {
	var urlDomainsDTO []sharedDTO.UrlDomainDTO
	for _, urlDomain := range urlDomains {
		urlDomainsDTO = append(urlDomainsDTO, sharedDTO.UrlDomainDTO{
			Name: urlDomain.Name,
			Url:  urlDomain.Url,
		})
	}
	return urlDomainsDTO
}

func ConvertEntityToProxyLoaderDTO(proxyLoaders []entity.ProxyLoader) []sharedDTO.ProxyLoaderDTO {
	var proxyLoadersDTO []sharedDTO.ProxyLoaderDTO
	for _, proxyLoader := range proxyLoaders {
		proxyLoadersDTO = append(proxyLoadersDTO, sharedDTO.ProxyLoaderDTO{
			Name:     proxyLoader.Name,
			Priority: proxyLoader.Priority,
		})
	}
	return proxyLoadersDTO
}

func ConvertEntityToCaptchaSolverDTO(captchaSolvers []entity.CaptchaSolver) []sharedDTO.CaptchaSolverDTO {
	var captchaSolversDTO []sharedDTO.CaptchaSolverDTO
	for _, captchaSolver := range captchaSolvers {
		captchaSolversDTO = append(captchaSolversDTO, sharedDTO.CaptchaSolverDTO{
			Name:     captchaSolver.Name,
			Priority: captchaSolver.Priority,
		})
	}
	return captchaSolversDTO
}

func ConvertEntityToUseCaseHttpGatewayParamsVersion(httpGatewayParamsVersion []entity.HttpGatewayParamsData) []outputDTO.HttpGatewayParamsVersionData {
	var httpGatewayParamsVersionDTO []outputDTO.HttpGatewayParamsVersionData
    log.Println("httpGatewayParamsVersion", httpGatewayParamsVersion)
	for _, httpGatewayParams := range httpGatewayParamsVersion {
		httpGatewayParamsVersionDTO = append(httpGatewayParamsVersionDTO, outputDTO.HttpGatewayParamsVersionData{
			JobParamsID: string(httpGatewayParams.JobParamsID),
			Params: &outputDTO.HttpGatewayParamsDTO{
				ID:             string(httpGatewayParams.Params.ID),
				Service:        httpGatewayParams.Params.Service,
				Source:         httpGatewayParams.Params.Source,
				Context:        httpGatewayParams.Params.Context,
				BaseUrl:        httpGatewayParams.Params.BaseUrl,
				UrlDomains:     ConvertEntityToUrlDomainDTO(httpGatewayParams.Params.UrlDomains),
				Headers:        httpGatewayParams.Params.Headers,
				EnableProxy:    httpGatewayParams.Params.EnableProxy,
				ProxyLoaders:   ConvertEntityToProxyLoaderDTO(httpGatewayParams.Params.ProxyLoaders),
				EnableCaptcha:  httpGatewayParams.Params.EnableCaptcha,
				CaptchaSolvers: ConvertEntityToCaptchaSolverDTO(httpGatewayParams.Params.CaptchaSolvers),
				ParamsID:       string(httpGatewayParams.JobParamsID),
				CreatedAt:      httpGatewayParams.Params.CreatedAt,
				UpdatedAt:      httpGatewayParams.Params.UpdatedAt,
			},
		})
	}
    log.Println("httpGatewayParamsVersionDTO", httpGatewayParamsVersionDTO)
    return httpGatewayParamsVersionDTO
}

func ConvertUseCaseToEntityHttpGatewayParamsVersion(httpGatewayParams outputDTO.HttpGatewayParamsDTO) *entity.HttpGatewayParams {
	return &entity.HttpGatewayParams{
		ID:             string(httpGatewayParams.ID),
		Service:        httpGatewayParams.Service,
		Source:         httpGatewayParams.Source,
		Context:        httpGatewayParams.Context,
		BaseUrl:        httpGatewayParams.BaseUrl,
		JobParamsID:    httpGatewayParams.ParamsID,
		UrlDomains:     ConvertUrlDomainDTOToEntity(httpGatewayParams.UrlDomains),
		Headers:        httpGatewayParams.Headers,
		EnableProxy:    httpGatewayParams.EnableProxy,
		ProxyLoaders:   ConvertProxyLoaderDTOToEntity(httpGatewayParams.ProxyLoaders),
		EnableCaptcha:  httpGatewayParams.EnableCaptcha,
		CaptchaSolvers: ConvertCaptchaSolverDTOToEntity(httpGatewayParams.CaptchaSolvers),
		CreatedAt:      httpGatewayParams.CreatedAt,
		UpdatedAt:      httpGatewayParams.UpdatedAt,
	}
}
