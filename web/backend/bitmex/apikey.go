package bitmex

// apiKey相关操作

import (
	"assets/web/backend/bitmex/resource"
	"encoding/json"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
)

// APIKeyAPI 对象
type APIKeyAPI struct {
	Configuration resource.Configuration
}

// NewAPIKeyAPI 创建APIKeyAPI对象
func NewAPIKeyAPI() *APIKeyAPI {
	configuration := resource.NewConfiguration()
	return &APIKeyAPI{
		Configuration: *configuration,
	}
}

// NewAPIKeyAPIWithBasePath 创建APIKeyAPI对象，并设置服务器路径
func NewAPIKeyAPIWithBasePath(basePath string) *APIKeyAPI {
	configuration := resource.NewConfiguration()
	configuration.BasePath = basePath
	return &APIKeyAPI{
		Configuration: *configuration,
	}
}

// apiKeyGet 获取apiKey
func (a APIKeyAPI) apiKeyGet(reverse string) ([]resource.APIKey, *resource.APIResponse, error) {
	var httpMethod = "Get"
	// create path and map variables
	pathURI := "/apiKey"
	path := a.Configuration.BasePath + pathURI

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	queryParams.Add("reverse", a.Configuration.APIClient.ParameterToString(reverse, ""))

	// 不支持APIKey的校验模式，需要token

	// 添加APIKey验证
	// a.Configuration.AddAPIKey(pathURI, httpMethod, queryParams, formParams)

	for key := range a.Configuration.DefaultHeader {
		log.Printf("key--------------:%s", key)
		log.Printf("value--------------:%s", a.Configuration.DefaultHeader[key])
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	localVarHTTPContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		headerParams["Content-Type"] = localVarHTTPContentType
	}
	localVarHTTPHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	localVarHTTPHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		headerParams["Accept"] = localVarHTTPHeaderAccept
	}
	var successPayload = new([]resource.APIKey)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
}

// GetList 获取APIKey列表
func (a APIKeyAPI) GetList(c *gin.Context) {
	reverse := c.DefaultQuery("reverse", "false")
	var result = new([]resource.APIKey)
	var response = new(resource.APIResponse)
	var err error
	*result, response, err = a.apiKeyGet(reverse)
	if err != nil {
		log.Printf("response--------------%v", response.Response)
		c.JSON(200, gin.H{
			"code": "0001",
			"data": "",
			"msg":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": "0000",
			"data": result,
			"msg":  "success",
		})
	}
}
