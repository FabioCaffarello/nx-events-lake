package usecase

import (
	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	inputDTO "libs/golang/services/dtos/services-file-catalog-handler/input"
	outputDTO "libs/golang/services/dtos/services-file-catalog-handler/output"
)

type CreateFileCatalogUseCase struct {
	FileCatalogRepository entity.FileCatalogInterface
}

func NewCreateFileCatalogUseCase(
	repository entity.FileCatalogInterface,
) *CreateFileCatalogUseCase {
	return &CreateFileCatalogUseCase{
		FileCatalogRepository: repository,
	}
}

func (ccu *CreateFileCatalogUseCase) Execute(schemaCatalog inputDTO.FileCatalogDTO) (outputDTO.FileCatalogDTO, error) {
	schemaCatalogEntity, err := entity.NewFileCatalog(
		schemaCatalog.Service,
		schemaCatalog.Source,
		schemaCatalog.Context,
		schemaCatalog.LakeLayer,
		schemaCatalog.SchemaType,
		schemaCatalog.Catalog,
	)
	if err != nil {
		return outputDTO.FileCatalogDTO{}, err
	}

	err = ccu.FileCatalogRepository.Save(schemaCatalogEntity)
	if err != nil {
		return outputDTO.FileCatalogDTO{}, err
	}

	dto := outputDTO.FileCatalogDTO{
		ID:         string(schemaCatalogEntity.ID),
		Service:    schemaCatalogEntity.Service,
		Source:     schemaCatalogEntity.Source,
		Context:    schemaCatalogEntity.Context,
		LakeLayer:  schemaCatalogEntity.LakeLayer,
		SchemaType: schemaCatalogEntity.SchemaType,
		CatalogID:  string(schemaCatalogEntity.CatalogID),
		Catalog:    schemaCatalogEntity.Catalog,
		CreatedAt:  schemaCatalogEntity.CreatedAt,
		UpdatedAt:  schemaCatalogEntity.UpdatedAt,
	}

	return dto, nil
}
