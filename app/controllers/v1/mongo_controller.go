package v1

import (
	"encoding/json"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/database"
	"github.com/yunussandikci/go-pure-api/app/models/request"
	"github.com/yunussandikci/go-pure-api/app/models/response"
	"github.com/yunussandikci/go-pure-api/app/server"
	"github.com/yunussandikci/go-pure-api/app/services"
	"time"
)

type MongoController interface {
	GetRecords(request *server.Request, response *server.Response)
}

type mongoController struct {
	recordsService services.MongoRecordsService
}

func NewMongoController(database database.MongoDatabase) MongoController {
	return &mongoController{
		recordsService: services.NewMongoRecordsService(database),
	}
}

func NewMongoControllerWith(recordsService services.MongoRecordsService) MongoController {
	return &mongoController{
		recordsService: recordsService,
	}
}

// GetRecords
// @Summary Gets a records from database
// @Description This endpoints returns records from the mongo database with the provided filter in request
// @Accept  json
// @Produce  json
// @Param Request body request.MongoGetRecordsRequest true "Filter for the request"
// @Success 200 {object} response.MongoRecordsResponse
// @Failure 400 {object} common.ApiError
// @tags values
// @Router /mongo [post]
func (m mongoController) GetRecords(req *server.Request, res *server.Response) {
	var requestBody request.MongoGetRecordsRequest
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
	startDate, _ := time.Parse(common.DefaultTimeFormat, requestBody.StartDate)
	endDate, _ := time.Parse(common.DefaultTimeFormat, requestBody.EndDate)
	records, recordsErr := m.recordsService.GetRecords(requestBody.MinCount, requestBody.MaxCount, startDate, endDate)
	if recordsErr != nil {
		res.Error = recordsErr
		return
	}
	res.Body = response.NewMongoRecordsResponse(0, "Success", records)
}
