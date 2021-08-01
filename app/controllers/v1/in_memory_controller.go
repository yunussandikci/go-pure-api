package v1

import (
	"encoding/json"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/database"
	"github.com/yunussandikci/go-pure-api/app/models/request"
	"github.com/yunussandikci/go-pure-api/app/models/response"
	"github.com/yunussandikci/go-pure-api/app/server"
	"github.com/yunussandikci/go-pure-api/app/services"
	"net/http"
)

type InMemoryController interface {
	Create(request *server.Request, response *server.Response)
	GetRecords(request *server.Request, response *server.Response)
}

type inMemoryController struct {
	recordsService services.InMemoryRecordsService
}

func NewInMemoryController(database database.InMemoryDatabase) InMemoryController {
	return &inMemoryController{
		recordsService: services.NewInMemoryRecordsService(database),
	}
}

func NewInMemoryControllerWith(recordsService services.InMemoryRecordsService) InMemoryController {
	return &inMemoryController{
		recordsService: recordsService,
	}
}

// Create
// @Summary Creates a a new key-value
// @Description This endpoints persists a new key-value in the in-memory database
// @Accept  json
// @Produce  json
// @Param Value body request.InMemoryCreateRecordRequest true "The key and value that will be persist."
// @Success 201 {object} response.InMemoryRecordResponse
// @Failure 400 {object} common.ApiError
// @tags values
// @Router /in-memory [post]
func (i inMemoryController) Create(req *server.Request, res *server.Response) {
	var requestBody request.InMemoryCreateRecordRequest
	unmarshallErr := json.Unmarshal(req.Body, &requestBody)
	if unmarshallErr != nil {
		res.Error = common.NewBadRequestError()
		return
	}
	validateErr := common.Validate.Struct(requestBody)
	if validateErr != nil {
		res.Error = common.NewBadRequestErrorWithMessage(common.TranslateValidationErrors(validateErr))
		return
	}
	createErr := i.recordsService.Create(requestBody.Key, requestBody.Value)
	if createErr != nil {
		res.Error = createErr
		return
	}
	res.StatusCode = http.StatusCreated
	res.Body = response.NewInMemoryRecordResponse(requestBody.Key, requestBody.Value)
}

// GetRecords
// @Summary Gets a value of the key provided
// @Description This endpoints returns value of the key provided
// @Accept  json
// @Produce  json
// @Param key query string true "Key"
// @Success 200 {object} response.InMemoryRecordResponse
// @Failure 400 {object} common.ApiError
// @tags values
// @Router /in-memory [get]
func (i inMemoryController) GetRecords(req *server.Request, res *server.Response) {
	var requestData request.InMemoryGetRecordRequest
	value, paramExist := req.Parameters["key"]
	if !paramExist {
		res.Error = common.NewBadRequestError()
		return
	}
	requestData.Key = value[0]
	validateErr := common.Validate.Struct(requestData)
	if validateErr != nil {
		res.Error = common.NewBadRequestErrorWithMessage(common.TranslateValidationErrors(validateErr))
		return
	}
	record, getErr := i.recordsService.Get(requestData.Key)
	if getErr != nil {
		res.Error = getErr
		return
	}
	res.Body = response.NewInMemoryRecordResponse(value[0], record)
}
