package pin

type Response struct {
	Data    interface{}            `json:"data,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
	TraceId string                 `json:"trace_id,omitempty"`
	Error   *ResponseError         `json:"error,omitempty"`
}

type ResponseError struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
	Key     string `json:"key,omitempty"`
}
