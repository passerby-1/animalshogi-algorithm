package tools

import (
	"encoding/json"
	"fmt"
)

func UnmarshalJSON(data []byte) (map[string]string, error) {

	var Raw map[string]string
	fmt.Printf("data ([]byte): %v\n", data)

	err := json.Unmarshal(data, &Raw)
	fmt.Printf("raw (unmarshaled): %v\n", Raw)

	if err != nil {
		return nil, err
	}

	return Raw, nil
}
