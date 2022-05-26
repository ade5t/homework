package httpControllers

import (
	"bytes"
	"io"
	"net/http"
)

func GetResultForTask(url string, body []byte, writer http.ResponseWriter) []byte {
	response, error := http.Post(url, "application/json", bytes.NewBuffer(body))

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
