package httpStatusText

import "net/http"

const (
	SUCCESS = "success"
	ERROR   = "error"
	FAILED  = "failed"
)

func ParseRequestBody(request *http.Request) {
	body := request.Body
	defer body.Close()

}
