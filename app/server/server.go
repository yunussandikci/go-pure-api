package server

import (
	"encoding/json"
	"fmt"
	"github.com/yunussandikci/go-pure-api/app/common"
	_ "github.com/yunussandikci/go-pure-api/app/docs"
	"io/ioutil"
	"net/http"
)

type ApplicationServer struct {
	httpMux *http.ServeMux
}

func NewApplicationServer() *ApplicationServer {
	httpMux := http.NewServeMux()
	return &ApplicationServer{
		httpMux: httpMux,
	}
}

//handleRequest Handles base http request with the provided handler functions according to method of the request.
//It returns bool according to if request is handled by a function or not. Also it initializes
//request object for inner method use-cases.
func handleRequest(req *http.Request, handlerFunctions HandlerFunctions, response *Response) bool {
	readBytes, _ := ioutil.ReadAll(req.Body)
	var request = &Request{
		Body:       readBytes,
		Method:     req.Method,
		Headers:    req.Header,
		Parameters: req.URL.Query(),
	}
	switch req.Method {
	case http.MethodGet:
		if handlerFunctions.Get != nil {
			handlerFunctions.Get(request, response)
			return true
		}
	case http.MethodPost:
		if handlerFunctions.Post != nil {
			handlerFunctions.Post(request, response)
			return true
		}
	case http.MethodPut:
		if handlerFunctions.Put != nil {
			handlerFunctions.Put(request, response)
			return true
		}
	case http.MethodPatch:
		if handlerFunctions.Patch != nil {
			handlerFunctions.Patch(request, response)
			return true
		}
	case http.MethodDelete:
		if handlerFunctions.Delete != nil {
			handlerFunctions.Delete(request, response)
			return true
		}
	}
	return false
}

//handleResponse Handles responses in order to send them to the requester. It handles also unified errors if needed.
func handleResponse(writer http.ResponseWriter, req *http.Request, response *Response) {
	if response.Error != nil {
		apiError, isApiError := response.Error.(*common.ApiError)
		if isApiError {
			response = newErrorResponse(apiError.Message, apiError.StatusCode, apiError.Code)
		} else {
			internalServerError(writer, req)
			return
		}
	}
	responseJson, jsonMarshallErr := json.Marshal(response.Body)
	if jsonMarshallErr != nil {
		internalServerError(writer, req)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")
	writer.WriteHeader(response.StatusCode)
	_, _ = writer.Write(responseJson)
	common.Logger.Infof("Method:%s Path:%s Code:%d", req.Method, req.RequestURI, response.StatusCode)
}

//recoverError Recovers panics and returns proper error the to requesters.
func recoverError(writer http.ResponseWriter, req *http.Request) {
	r := recover()
	if r != nil {
		common.Logger.Errorf("Error:%s", r)
		internalServerError(writer, req)
	}
}

//HandleFunctions Commits handler functions based on their pattern to handle request and responses in further.
func (s *ApplicationServer) HandleFunctions(pattern string, handlerFunctions HandlerFunctions) {
	s.httpMux.HandleFunc(pattern, func(writer http.ResponseWriter, req *http.Request) {
		defer recoverError(writer, req)
		response := NewResponse()
		if handleRequest(req, handlerFunctions, response) {
			handleResponse(writer, req, response)
		} else {
			notFound(writer, req)
		}
	})
}


//HandleFunc Commits base httpHandler provided with the pattern.
func (s *ApplicationServer) HandleFunc(pattern string, handler http.HandlerFunc) {
	s.httpMux.HandleFunc(pattern, handler)
}

//Run Runs the application with the provided port, returns error if required.
func (s *ApplicationServer) Run(port string) error {
	s.httpMux.HandleFunc("/", notFound)
	listenErr := http.ListenAndServe(fmt.Sprintf(":%s", port), s.httpMux)
	if listenErr != nil {
		return listenErr
	}
	return nil
}
