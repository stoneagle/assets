package swagger

import (
	"encoding/json"
	"net/url"
)

type AnnouncementApi struct {
	Configuration Configuration
}

func NewAnnouncementApi() *AnnouncementApi {
	configuration := NewConfiguration()
	return &AnnouncementApi{
		Configuration: *configuration,
	}
}

func NewAnnouncementApiWithBasePath(basePath string) *AnnouncementApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AnnouncementApi{
		Configuration: *configuration,
	}
}

/**
 * Get site announcements.
 *
 * @param columns Array of column names to fetch. If omitted, will return all columns.
 * @return []Announcement
 */
func (a AnnouncementApi) AnnouncementGet(columns string) ([]Announcement, *APIResponse, error) {

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
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Announcement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get urgent (banner) announcements.
 *
 * @return []Announcement
 */
func (a AnnouncementApi) AnnouncementGetUrgent() ([]Announcement, *APIResponse, error) {

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
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]Announcement)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
