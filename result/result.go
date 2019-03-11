package result

import (
	"encoding/json"
	"net/http"
)

func GetResult() int {
	return http.StatusOK
}

func GetResultCreated() int {
	return http.StatusCreated
}

func GetResultOK() int {
	return http.StatusOK
}

func WriteErrorResponse(responseWriter http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(responseWriter, msgbytes, http.StatusBadRequest)
	return
}

func WriteJsonResponse(responseWriter http.ResponseWriter, bytes []byte, statusCode int) {
	responseWriter.WriteHeader(statusCode)
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.Write(bytes)
}
