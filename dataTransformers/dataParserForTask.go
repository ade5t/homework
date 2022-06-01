package dataTransformers

import (
	"encoding/json"
	"net/http"
	"trainingProjectInGo/types"
)

// Распарсить исходные данные для задачи, полученные с сервера
func Parse(responseBody []byte, writer http.ResponseWriter) []types.TaskData {
	var bodyData [][]interface{}

	error := json.Unmarshal(responseBody, &bodyData)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	var elements []types.TaskData

	for _, structure := range bodyData {
		var element types.TaskData

		for _, value := range structure[0].([]interface{}) {
			element.Data = append(element.Data, int(value.(float64)))
		}

		if len(structure) > 1 {
			element.Offset = int(structure[1].(float64))
		}

		elements = append(elements, element)
	}

	return elements
}
