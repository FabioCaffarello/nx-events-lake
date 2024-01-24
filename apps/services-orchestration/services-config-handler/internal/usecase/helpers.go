package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
	sharedDTO "libs/golang/services/dtos/services-config-handler/shared"
)

func ConvertEntityToUseCaseDependencies(entityDeps []entity.JobDependencies) []sharedDTO.JobDependencies {
	usecaseDeps := make([]sharedDTO.JobDependencies, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = sharedDTO.JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return usecaseDeps
}

func ConvertDependsOnDTOToEntity(dependsOn []sharedDTO.JobDependencies) []entity.JobDependencies {
	entityDeps := make([]entity.JobDependencies, len(dependsOn))
	for i, dep := range dependsOn {
		entityDeps[i] = entity.JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return entityDeps
}

func ConvertEntityToUseCaseConfigVersion(entityDeps []entity.ConfigData) []outputDTO.ConfigVersionData {
	usecaseDeps := make([]outputDTO.ConfigVersionData, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = outputDTO.ConfigVersionData{
			ConfigID: dep.ConfigID,
			Config: &outputDTO.ConfigDTO{
				ID:                string(dep.Config.ID),
				Name:              dep.Config.Name,
				Active:            dep.Config.Active,
				Frequency:         dep.Config.Frequency,
				Service:           dep.Config.Service,
				Source:            dep.Config.Source,
				Context:           dep.Config.Context,
                InputMethod:       dep.Config.InputMethod,
                OutputMethod:      dep.Config.OutputMethod,
				DependsOn:         ConvertEntityToUseCaseDependencies(dep.Config.DependsOn),
				ConfigID:          dep.Config.ConfigID,
				ServiceParameters: dep.Config.ServiceParameters,
				JobParameters:     dep.Config.JobParameters,
				CreatedAt:         dep.Config.CreatedAt,
				UpdatedAt:         dep.Config.UpdatedAt,
			},
		}
	}
	return usecaseDeps
}

func ConvertConfigDTOToEntity(config outputDTO.ConfigDTO) *entity.Config {
	return &entity.Config{
		ID:                string(config.ID),
		Name:              config.Name,
		Active:            config.Active,
		Frequency:         config.Frequency,
		Service:           config.Service,
		Source:            config.Source,
		Context:           config.Context,
        InputMethod:       config.InputMethod,
        OutputMethod:      config.OutputMethod,
		DependsOn:         ConvertDependsOnDTOToEntity(config.DependsOn),
		ConfigID:          config.ConfigID,
		ServiceParameters: config.ServiceParameters,
		JobParameters:     config.JobParameters,
		CreatedAt:         config.CreatedAt,
		UpdatedAt:         config.UpdatedAt,
	}
}
