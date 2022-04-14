package main

import (
	"net/http"

	"trainingProjectInGo/dataTransformers"
	"trainingProjectInGo/httpControllers/requests"
	"trainingProjectInGo/taskSolvers/arrayShift"
	"trainingProjectInGo/taskSolvers/searchMissingElement"
	"trainingProjectInGo/taskSolvers/sequenceCheck"
	"trainingProjectInGo/taskSolvers/weirdEntry"
	"trainingProjectInGo/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/tasks/", func(writer http.ResponseWriter, request *http.Request) {

	})

	router.Route("/task/{taskName}", func(router chi.Router) {
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			switch chi.URLParam(request, "taskName") {
			case "Циклическая ротация":

				var rawDataForTask []byte = requests.GetDataForTask("http://116.203.203.76:3000/tasks/Циклическая ротация", writer)
				var dataForTask []types.TaskData = dataTransformers.Parse(rawDataForTask, writer)

				var outputArray [][]int
				for _, value := range dataForTask {
					outputArray = append(outputArray, arrayShift.Solution(value.Data, value.Offset))
				}

				var dataForRequest = dataTransformers.Collect("Циклическая ротация", dataForTask, outputArray, writer)
				writer.Write(dataForRequest)
			case "Проверка последовательности":

				var rawDataForTask []byte = requests.GetDataForTask("http://116.203.203.76:3000/tasks/Проверка последовательности", writer)
				var dataForTask []types.TaskData = dataTransformers.Parse(rawDataForTask, writer)

				var outputArray []int
				for _, value := range dataForTask {
					outputArray = append(outputArray, sequenceCheck.Solution(value.Data))
				}

				var dataForRequest = dataTransformers.Collect("Проверка последовательности", dataForTask, outputArray, writer)
				writer.Write(dataForRequest)
			case "Чудные вхождения в массив":

				var rawDataForTask []byte = requests.GetDataForTask("http://116.203.203.76:3000/tasks/Чудные вхождения в массив", writer)
				var dataForTask []types.TaskData = dataTransformers.Parse(rawDataForTask, writer)

				var outputArray []int
				for _, value := range dataForTask {
					outputArray = append(outputArray, weirdEntry.Solution(value.Data))
				}

				var dataForRequest = dataTransformers.Collect("Чудные вхождения в массив", dataForTask, outputArray, writer)
				writer.Write(dataForRequest)
			case "Поиск отсутствующего элемента":

				var rawDataForTask []byte = requests.GetDataForTask("http://116.203.203.76:3000/tasks/Поиск отсутствующего элемента", writer)
				var dataForTask []types.TaskData = dataTransformers.Parse(rawDataForTask, writer)

				var outputArray []int
				for _, value := range dataForTask {
					outputArray = append(outputArray, searchMissingElement.Solution(value.Data))
				}

				var dataForRequest = dataTransformers.Collect("Поиск отсутствующего элемента", dataForTask, outputArray, writer)
				writer.Write(dataForRequest)
			}
		})
	})

	http.ListenAndServe(":80", router)
}
