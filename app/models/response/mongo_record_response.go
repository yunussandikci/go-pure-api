package response

import "time"

type MongoRecordsResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"msg"`
	Records []MongoRecordResponse `json:"records"`
}

type MongoRecordResponse struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

func NewMongoRecordResponse(key string, createdAt time.Time, totalCount int) MongoRecordResponse {
	return MongoRecordResponse{key, createdAt, totalCount}
}

func NewMongoRecordsResponse(code int, message string, records []MongoRecordResponse) MongoRecordsResponse {
	return MongoRecordsResponse{code, message, records}
}
