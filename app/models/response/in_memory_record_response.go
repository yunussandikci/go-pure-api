package response

type InMemoryRecordResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewInMemoryRecordResponse(key, value string) InMemoryRecordResponse {
	return InMemoryRecordResponse{Key:key, Value: value}
}