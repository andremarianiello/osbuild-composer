/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// PulpAnsibleApiAPIService PulpAnsibleApiAPI service
type PulpAnsibleApiAPIService service

type PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiAPIService
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest) Fields(fields []string) PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiGetExecute(r)
}

/*
PulpAnsibleGalaxyApiGet Method for PulpAnsibleGalaxyApiGet

Return a response to the "GET" action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest
*/
func (a *PulpAnsibleApiAPIService) PulpAnsibleGalaxyApiGet(ctx context.Context, path string) PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest {
	return PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
func (a *PulpAnsibleApiAPIService) PulpAnsibleGalaxyApiGetExecute(r PulpAnsibleApiAPIPulpAnsibleGalaxyApiGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiAPIService.PulpAnsibleGalaxyApiGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}