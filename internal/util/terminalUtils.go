package util

import (
	"encoding/json"
	"fmt"
)

func PrintFormatedJSON(data interface{}) {
	formatedJson, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	// Escrevendo o JSON formatado no terminal
	fmt.Println(string(formatedJson))
}
