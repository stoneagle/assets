package resource

import (
	"net/http"
)

// APIResponse 对象
type APIResponse struct {
	*http.Response
	Message string `json:"message,omitempty"`
}

// NewAPIResponse 创建APIResponse对象
func NewAPIResponse(r *http.Response) *APIResponse {
	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError 创建带有error的APIResponse对象
func NewAPIResponseWithError(errorMessage string) *APIResponse {
	response := &APIResponse{Message: errorMessage}
	return response
}
