package httpd

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []byte `json:"data,omitempty"`
}
