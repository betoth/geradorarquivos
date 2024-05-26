package file

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func GenerateSha256(data interface{}) {
	datab, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error serializing data: %v\n", err)
		return
	}

	hasher := sha256.New()
	hasher.Write(datab)
	fmt.Printf("Generate in function %x\n", hasher.Sum(nil))
}
