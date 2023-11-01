package usecase

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-config-handler/output"
)

type ListOneConfigVersionByIdUseCase struct {
	ConfigVersionRepository entity.ConfigVersionInterface
}

func NewListOneConfigVersionByIdUseCase(
	repository entity.ConfigVersionInterface,
) *ListOneConfigVersionByIdUseCase {
	return &ListOneConfigVersionByIdUseCase{
		ConfigVersionRepository: repository,
	}
}

func (lcv *ListOneConfigVersionByIdUseCase) Execute(id string) (outputDTO.ConfigVersionDTO, error) {
	item, err := lcv.ConfigVersionRepository.FindOneById(id)
	if err != nil {
		return outputDTO.ConfigVersionDTO{}, err
	}

	dto := outputDTO.ConfigVersionDTO{
		ID:       string(item.ID),
		Versions: ConvertEntityToUseCaseConfigVersion(item.Versions),
	}

	return dto, nil
}
