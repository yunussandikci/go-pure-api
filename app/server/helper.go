package server

import "net/http"

//internalServerError Responds Internal Server Error when called
func internalServerError(writer http.ResponseWriter, request *http.Request) {
	handleResponse(writer, request, newErrorResponse(http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError, http.StatusInternalServerError))
}

//notFound Responds Not Found when called
func notFound(writer http.ResponseWriter, request *http.Request) {
	handleResponse(writer, request, newErrorResponse(http.StatusText(http.StatusNotFound),
		http.StatusNotFound, http.StatusNotFound))
}
