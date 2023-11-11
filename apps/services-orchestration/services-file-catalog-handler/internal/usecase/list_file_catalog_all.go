package usecase

import (
	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	outputDTO "libs/golang/services/dtos/services-file-catalog-handler/output"
)

type ListAllFileCatalogsUseCase struct {
	FileCatalogRepository entity.FileCatalogInterface
}

func NewListAllFileCatalogsUseCase(
	repository entity.FileCatalogInterface,
) *ListAllFileCatalogsUseCase {
	return &ListAllFileCatalogsUseCase{
		FileCatalogRepository: repository,
	}
}

func (la *ListAllFileCatalogsUseCase) Execute() ([]outputDTO.FileCatalogDTO, error) {
	items, err := la.FileCatalogRepository.FindAll()
	if err != nil {
		return []outputDTO.FileCatalogDTO{}, err
	}
	var result []outputDTO.FileCatalogDTO
	for _, item := range items {
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
		result = append(result, dto)
	}
	return result, nil
}
