package pin

type Response struct {
	Data    interface{}            `json:"data"`
	Meta    map[string]interface{} `json:"meta"`
	TraceId string                 `json:"trace_id"`
	Error   *ResponseError         `json:"error"`
}

type ResponseError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Key     string `json:"key"`
}
