package requests

import (
	"io"
	"net/http"
)

func GetDataForTask(url string, writer http.ResponseWriter) []byte {
	response, error := http.Get(url)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	defer response.Body.Close()
	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	return responseBody
}
