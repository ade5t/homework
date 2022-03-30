package main

import (
	"encoding/json"
	"homework/arrayShift"
	"homework/searchMissingElement"
	"homework/sequenceCheck"
	"homework/weirdEntry"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type taskData struct {
	data   []int
	offset int
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/tasks/", func(writer http.ResponseWriter, request *http.Request) {

	})

	router.Route("/task/{taskName}", func(router chi.Router) {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			switch chi.URLParam(request, "taskName") {
			case "Циклическая ротация":
				var inputArray []taskData
				var outputArray [][]int
				inputArray, _ = getDataForTask("Циклическая ротация", writer)

				for _, value := range inputArray {
					outputArray = append(outputArray, arrayShift.Solution(value.data, value.offset))
				}
			case "Проверка последовательности":
				var inputArray []taskData
				var outputArray []int
				inputArray, _ = getDataForTask("Проверка последовательности", writer)

				for _, value := range inputArray {
					outputArray = append(outputArray, sequenceCheck.Solution(value.data))
				}
			case "Чудные вхождения в массив":
				var inputArray []taskData
				var outputArray []int
				inputArray, _ = getDataForTask("Чудные вхождения в массив", writer)

				for _, value := range inputArray {
					outputArray = append(outputArray, weirdEntry.Solution(value.data))
				}
			case "Поиск отсутствующего элемента":
				var inputArray []taskData
				var outputArray []int
				inputArray, _ = getDataForTask("Поиск отсутствующего элемента", writer)

				for _, value := range inputArray {
					outputArray = append(outputArray, searchMissingElement.Solution(value.data))
				}
			}
		})
	})

	http.ListenAndServe(":80", router)
}

func getDataForTask(taskName string, writer http.ResponseWriter) ([]taskData, [][]interface{}) {
	response, error := http.Get("http://116.203.203.76:3000/tasks/" + taskName)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	responseBody, error := io.ReadAll(response.Body)
	defer response.Body.Close()

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	return parseBody(responseBody)
}

func parseBody(responseBody []byte) ([]taskData, [][]interface{}) {
	var bodyData [][]interface{}

	_ = json.Unmarshal(responseBody, &bodyData)

	var elements []taskData

	for _, structure := range bodyData {
		var tmp taskData
		for _, value := range structure[0].([]interface{}) {
			tmp.data = append(tmp.data, int(value.(float64)))
		}

		if len(structure) > 1 {
			tmp.offset = int(structure[1].(float64))
		}

		elements = append(elements, tmp)
	}

	return elements, bodyData
}
