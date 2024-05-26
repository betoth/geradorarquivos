package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateJson(fileName string, filePath string, data interface{}, clearDir bool) error {

	// Verificar e criar o diretório se necessário.
	if err := CreateDirIfNotExist(filePath); err != nil {
		return fmt.Errorf("falha ao verificar/criar o diretório: %v", err)
	}
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
	// Verificar e criar o diretório se necessário.
	if err := CreateDirIfNotExist(path); err != nil {
		return fmt.Errorf("falha ao verificar/criar o diretório: %v", err)
	}
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

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("falha ao criar a pasta: %v", err)
		}
	}
	return nil
}
