package configuuid

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"

  "github.com/google/uuid"
)

type ID = string


func GenerateConfigID(properties map[string]interface{}) (string, error) {
  serializedConfigProperties, err := json.Marshal(properties)
  if err != nil {
      return "", fmt.Errorf("error marshaling config properties: %w", err)
  }

  configHash := hashConfig(serializedConfigProperties)
  configID := generateUUIDFromHash(configHash)

  return ID(configID), nil
}

func hashConfig(data []byte) []byte {
  hash := sha256.Sum256(data)
  return hash[:]
}

func generateUUIDFromHash(hash []byte) string {
  combinedHash := sha256.Sum256(hash)
  return uuid.NewSHA1(uuid.Nil, combinedHash[:]).String()
}
