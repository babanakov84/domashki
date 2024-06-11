package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/babanakov84/domashki/hw06_testing/pkg/fixapp/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		// fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		// fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		// fmt.Printf("Error: %v", err)
		return nil, err
	}

	return data, nil
}
