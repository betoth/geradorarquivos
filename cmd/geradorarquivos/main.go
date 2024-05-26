package main

import (
	"fmt"

	"github.com/betoth/geradorarquivos/internal/datafiles"
	"github.com/betoth/geradorarquivos/internal/file"
	"github.com/betoth/geradorarquivos/internal/util"

	"github.com/google/uuid"
)

func main() {

	fileName := uuid.New()
	filePath := "../../data/"
	data := datafiles.LoanData
	removeFiles := true

	fmt.Println("#Fila de Criação")
	fmt.Println("\n```json")
	fmt.Println("{'id':'fec8fa8f-14ff-4b7c-ba14-67cd29556889', 'hash': 'ed72ea0b10f3e44bfcc5eff9ec29a83c470399606e2b79302be1929c7849788b'}")
	fmt.Println("```")
	fmt.Println("\n#S3")
	fmt.Println("\n```json")
	util.PrintFormatedJSON(data)
	fmt.Println("```")

	file.GenerateSha256(data)
	file.CreateJson(fileName.String(), filePath, data, removeFiles)

}
