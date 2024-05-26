package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateJson(fileName string, filePath string, data interface{}, clearDir bool) error {
	// Serializando os dados para JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error serializing data: %v", err)
	}

	// Limpar o diretório se necessário
	if clearDir {
		removeFilesInDir(filePath)
	}

	// Escrever o JSON no arquivo
	err = os.WriteFile(filePath+"/"+fileName+".json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func removeFilesInDir(path string) error {
	// Abre o diretório
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()

	// Lê o conteúdo do diretório
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	// Itera sobre os arquivos no diretório
	for _, fileInfo := range fileInfos {
		// Verifica se o item é um arquivo
		if fileInfo.Mode().IsRegular() {
			// Constrói o caminho completo do arquivo
			filePath := filepath.Join(path, fileInfo.Name())
			// Remove o arquivo
			err := os.Remove(filePath)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
