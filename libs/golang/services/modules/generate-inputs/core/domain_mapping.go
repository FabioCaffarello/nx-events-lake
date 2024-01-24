package core

import (
	inputDTO "libs/golang/services/dtos/services-input-handler/input"
	domainUriPartition "libs/golang/services/modules/generate-inputs/domains/uri-partition"
)

type GenerateInputGenerator interface {
	Execute(data map[string]interface{}) (inputDTO.InputDTO, error)
}

func GetGenerateInputsMapping() map[string]GenerateInputGenerator {
	return map[string]GenerateInputGenerator{
		"GenerateInputUsingBucketUriAndPartition": domainUriPartition.NewGenerateInputUsingBucketUriAndPartition(),
	}
}
