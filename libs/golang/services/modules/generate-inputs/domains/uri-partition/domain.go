package domain

import (
	inputDTO "libs/golang/services/dtos/services-input-handler/input"
)

type GenerateInput struct {
}

func NewGenerateInputUsingBucketUriAndPartition() *GenerateInput {
	return &GenerateInput{}
}

func (gi *GenerateInput) Execute(data map[string]interface{}) (inputDTO.InputDTO, error) {
	input := inputDTO.InputDTO{
		Data: map[string]interface{}{
			"documentUri": data["documentUri"],
			"partition":   data["partition"],
		},
	}
	return input, nil
}
