package file

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// GenerateSha256 gera o hash SHA-256 dos dados fornecidos e retorna o hash como uma string hexadecimal e um possível erro.
func GenerateSha256(data interface{}) (string, error) {
	// Serializar os dados para JSON.
	datab, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error serializing data: %v", err)
	}

	// Criar um novo hasher SHA-256 e escrever os dados serializados.
	hasher := sha256.New()
	hasher.Write(datab)

	// Obter o hash e converter para uma string hexadecimal.
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)

	return hashString, nil
}
