package pipelineuuid

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type ID = string

func GeneratePipelineID(properties map[string]interface{}) (string, error) {
	serializedPipelineProperties, err := json.Marshal(properties)
	if err != nil {
		return "", fmt.Errorf("error marshaling pipeline properties: %w", err)
	}

	pipelineHash := hashPipeline(serializedPipelineProperties)
	pipelineID := generateUUIDFromHash(pipelineHash)

	return ID(pipelineID), nil
}

func hashPipeline(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func generateUUIDFromHash(hash []byte) string {
	combinedHash := sha256.Sum256(hash)
	return uuid.NewSHA1(uuid.Nil, combinedHash[:]).String()
}
