package bitmex

// 公告栏相关

import (
	"assets/bitmex/resource"
	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"
)

// AnnouncementAPI 对象
type AnnouncementAPI struct {
	Configuration resource.Configuration
}

// NewAnnouncementAPI 创建公告API对象
func NewAnnouncementAPI() *AnnouncementAPI {
	configuration := resource.NewConfiguration()
	return &AnnouncementAPI{
		Configuration: *configuration,
	}
}

// NewAnnouncementAPIWithBasePath 创建公告API对象，并设置服务器路径
func NewAnnouncementAPIWithBasePath(basePath string) *AnnouncementAPI {
	configuration := resource.NewConfiguration()
	configuration.BasePath = basePath

	return &AnnouncementAPI{
		Configuration: *configuration,
	}
}

// AnnouncementGet 获取公告列表
func (a AnnouncementAPI) announcementGet(columns string) ([]resource.Announcement, *resource.APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/announcement"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("columns", a.Configuration.APIClient.ParameterToString(columns, ""))

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		headerParams["Content-Type"] = localVarHTTPContentType
	}
	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHTTPHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		headerParams["Accept"] = localVarHTTPHeaderAccept
	}
	var successPayload = new([]resource.Announcement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
}

// AnnouncementGetUrgent 获取紧急公告
func (a AnnouncementAPI) announcementGetUrgent() ([]resource.Announcement, *resource.APIResponse, error) {
	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/announcement/urgent"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		headerParams["Content-Type"] = localVarHTTPContentType
	}
	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHTTPHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		headerParams["Accept"] = localVarHTTPHeaderAccept
	}
	var successPayload = new([]resource.Announcement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, resource.NewAPIResponse(httpResponse.RawResponse), err
}

// GetList gin的获取列表
func (a AnnouncementAPI) GetList(c *gin.Context) {
	columns := c.DefaultQuery("columns", "")
	result := new([]resource.Announcement)
	var err error
	*result, _, err = a.announcementGet(columns)
	if err != nil {
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
