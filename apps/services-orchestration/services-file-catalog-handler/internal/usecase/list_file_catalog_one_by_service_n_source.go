package usecase

import (
	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-file-catalog-handler/output"
)

type ListOneFileCatalogByServiceAndSourceUseCase struct {
	FileCatalogRepository entity.FileCatalogInterface
}

func NewListOneFileCatalogByServiceAndSourceUseCase(
	repository entity.FileCatalogInterface,
) *ListOneFileCatalogByServiceAndSourceUseCase {
	return &ListOneFileCatalogByServiceAndSourceUseCase{
		FileCatalogRepository: repository,
	}
}

func (la *ListOneFileCatalogByServiceAndSourceUseCase) Execute(service string, source string) (outputDTO.FileCatalogDTO, error) {
	item, err := la.FileCatalogRepository.FindOneByServiceAndSource(service, source)
	if err != nil {
		return outputDTO.FileCatalogDTO{}, err
	}

	dto := outputDTO.FileCatalogDTO{
		ID:         string(item.ID),
		Service:    item.Service,
		Source:     item.Source,
		Context:    item.Context,
		LakeLayer:  item.LakeLayer,
		SchemaType: item.SchemaType,
		CatalogID:  string(item.CatalogID),
		Catalog:    item.Catalog,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}

	return dto, nil
}
