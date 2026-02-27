package utils

import (
	m "blog/pkg/models"
	"encoding/json"
	"fmt"
)

func ErrorMessageJson(err string, code string, details ...string) []byte {
	errorMessage := m.ErrorMessage{
		Error:   err,
		Code:    code,
		Details: details,
	}

	errorJson, err2 := json.Marshal(errorMessage)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}
	return errorJson
}
