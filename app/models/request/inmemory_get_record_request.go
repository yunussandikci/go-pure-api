package request

type InMemoryGetRecordRequest struct {
	Key string `json:"key" validate:"required"`
}
