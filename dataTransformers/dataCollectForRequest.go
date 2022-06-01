package dataTransformers

import (
	"encoding/json"
	"net/http"
	"trainingProjectInGo/types"
)

const userName = "ade5t"

// Упаковать результаты решения задачи для отправки на сервер
func Collect(taskName string, inputTaskData []types.TaskData, taskSolution interface{}, writer http.ResponseWriter) []byte {
	var taskResult types.TaskResult
	var requestData types.RequestData

	for index, element := range inputTaskData {
		var taskPayload []interface{}
		taskPayload = append(taskPayload, element.Data)

		if element.Offset > 0 {
			taskPayload = append(taskPayload, element.Offset)
		}

		taskResult.Payload = append(taskResult.Payload, taskPayload)

		switch value := taskSolution.(type) {
		case []int:
			taskResult.Results = append(taskResult.Results, value[index])
		case [][]int:
			taskResult.Results = append(taskResult.Results, value[index])
		}
	}

	requestData.TaskName = taskName
	requestData.UserName = userName
	requestData.Result = taskResult

	result, error := json.Marshal(requestData)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	return result
}
