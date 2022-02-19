package main

import (
	"encoding/json"
	"fmt"
	"homework/arrayShift"
	"homework/searchMissingElement"
	"homework/sequenceCheck"
	"homework/weirdEntry"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// var testArray = []int{1, 7, 3, 2, 6, 5}
type taskData struct {
	data   []int
	offset int
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/tasks/", func(writer http.ResponseWriter, request *http.Request) {
		// writer.Write([]byte("all tasks"))
		getDataForTask("Циклическая ротация", writer)
	})

	router.Route("/task/{taskName}", func(router chi.Router) {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			switch chi.URLParam(request, "taskName") {
			case "Циклическая ротация":
				var inputArray []taskData
				inputArray = getDataForTask("Циклическая ротация", writer)

				for _, value := range inputArray {
					fmt.Println("1", arrayShift.Solution(value.data, value.offset))
				}
				// v, _ := json.Marshal(inputArray[0].data)
				// writer.Write([]byte("ok"))
			case "Проверка последовательности":
				var inputArray []taskData
				inputArray = getDataForTask("Проверка последовательности", writer)

				for _, value := range inputArray {
					fmt.Println("1", sequenceCheck.Solution(value.data))
				}
			case "Чудные вхождения в массив":
				var inputArray []taskData
				inputArray = getDataForTask("Чудные вхождения в массив", writer)

				for _, value := range inputArray {
					fmt.Println("1", weirdEntry.Solution(value.data))
				}
			case "Поиск отсутствующего элемента":
				var inputArray []taskData
				inputArray = getDataForTask("Поиск отсутствующего элемента", writer)

				for _, value := range inputArray {
					fmt.Println("1", searchMissingElement.Solution(value.data))
				}
			}
		})
	})

	http.ListenAndServe(":80", router)
}

func getDataForTask(taskName string, writer http.ResponseWriter) []taskData {
	response, error := http.Get("http://116.203.203.76:3000/tasks/" + taskName)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	responseBody, error2 := io.ReadAll(response.Body)
	defer response.Body.Close()

	if error2 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(responseBody) //вот это вот выпилить

	return parseBody(responseBody)
}

func parseBody(responseBody []byte) []taskData {
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

	return elements
}
