package core

import (
	"errors"
	inputDTO "libs/golang/services/dtos/services-input-handler/input"
)

type DomainFactory struct {
	Domains         map[string]GenerateInputGenerator
	SchemaInputType string
}

var (
	ErrDomainDoesNotExist = errors.New("domain does not exist")
)

func NewDomainFactory() *DomainFactory {
	return &DomainFactory{
		Domains: GetGenerateInputsMapping(),
	}
}

func (df *DomainFactory) GetDomain(domain string) (GenerateInputGenerator, error) {
	domainMethod, ok := df.Domains[domain]
	if !ok {
		return nil, ErrDomainDoesNotExist
	}
	return domainMethod, nil
}

func (df *DomainFactory) GenerateInputs(
	domain string,
	data map[string]interface{},
) ([]inputDTO.InputDTO, error) {
	generateInputDomain, err := df.GetDomain(domain)
	if err != nil {
		return nil, err
	}
	inputDTOs := make([]inputDTO.InputDTO, 0)

	result, errEvent := data["result"].(map[string]interface{})
	if errEvent {
		newInput, err := generateInputDomain.Execute(result)
		if err != nil {
			return nil, err
		}
		inputDTOs = append(inputDTOs, newInput)
	}

	resultStream, errStream := data["result"].([]map[string]interface{})
	if errStream {
		for _, stream := range resultStream {
			newInput, err := generateInputDomain.Execute(stream)
			if err != nil {
				return nil, err
			}
			inputDTOs = append(inputDTOs, newInput)
		}
	}
	if !errEvent && !errStream {
		return nil, err
	}
	return inputDTOs, nil
}
