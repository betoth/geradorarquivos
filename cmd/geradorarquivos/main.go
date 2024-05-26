package main

import (
	"fmt"
	"log"
	"os"

	"github.com/betoth/geradorarquivos/internal/datafiles"
	"github.com/betoth/geradorarquivos/internal/file"
	"github.com/betoth/geradorarquivos/internal/util"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Token := os.Getenv("S3_TOKEN")

	fileName := uuid.New()
	filePath := "../../data/"
	data := datafiles.LoanData
	removeFiles := true

	fmt.Println(s3Token)
	fmt.Println("\n#Fila de Criação")
	fmt.Println("\n```json")
	fmt.Println("{'id':'fec8fa8f-14ff-4b7c-ba14-67cd29556889', 'hash': 'ed72ea0b10f3e44bfcc5eff9ec29a83c470399606e2b79302be1929c7849788b'}")
	fmt.Println("```")
	fmt.Println("\n#S3")
	fmt.Println("\n```json")
	util.PrintFormatedJSON(data)
	fmt.Println("```")

	hash, err := file.GenerateSha256(data)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Hash SHA-256: %s\n", hash)
	}

	file.CreateJson(fileName.String(), filePath, data, removeFiles)

}
