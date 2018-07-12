/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type MethodPATCHApiService service


/* MethodPATCHApiService Modify a key-value or delete a key
 Changes the value of the selected key in the key-value pair or deletes a key by setting the key value to &lt;literal&gt;null&lt;/literal&gt;.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param httpKeyvalZoneName The name of an HTTP keyval shared memory zone.
 @param httpKeyvalZoneKeyValue A new value for the key is specified in the JSON format.
 @return */
func (a *MethodPATCHApiService) PatchHttpKeyvalZoneKeyValue(ctx context.Context, httpKeyvalZoneName string, httpKeyvalZoneKeyValue NginxHttpKeyvalZone) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/http/keyvals/{httpKeyvalZoneName}"
	localVarPath = strings.Replace(localVarPath, "{"+"httpKeyvalZoneName"+"}", fmt.Sprintf("%v", httpKeyvalZoneName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &httpKeyvalZoneKeyValue
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* MethodPATCHApiService Modify a server in an HTTP upstream server group
 Modifies settings of a particular server in an HTTP upstream server group. Server parameters are specified in the JSON format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param httpUpstreamName The name of the upstream server group.
 @param httpUpstreamServerId The ID of the server.
 @param patchHttpUpstreamServer Server parameters, specified in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed.
 @return NginxHttpUpstreamConfServer*/
func (a *MethodPATCHApiService) PatchHttpUpstreamPeer(ctx context.Context, httpUpstreamName string, httpUpstreamServerId string, patchHttpUpstreamServer NginxHttpUpstreamConfServer) (NginxHttpUpstreamConfServer,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NginxHttpUpstreamConfServer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/http/upstreams/{httpUpstreamName}/servers/{httpUpstreamServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"httpUpstreamName"+"}", fmt.Sprintf("%v", httpUpstreamName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"httpUpstreamServerId"+"}", fmt.Sprintf("%v", httpUpstreamServerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &patchHttpUpstreamServer
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MethodPATCHApiService Modify a key-value or delete a key
 Changes the value of the selected key in the key-value pair or deletes a key by setting the key value to &lt;literal&gt;null&lt;/literal&gt;.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param streamKeyvalZoneName The name of a stream keyval shared memory zone.
 @param streamKeyvalZoneKeyValue A new value for the key is specified in the JSON format.
 @return */
func (a *MethodPATCHApiService) PatchStreamKeyvalZoneKeyValue(ctx context.Context, streamKeyvalZoneName string, streamKeyvalZoneKeyValue NginxStreamKeyvalZone) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/stream/keyvals/{streamKeyvalZoneName}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamKeyvalZoneName"+"}", fmt.Sprintf("%v", streamKeyvalZoneName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &streamKeyvalZoneKeyValue
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* MethodPATCHApiService Modify a server in a stream upstream server group
 Modifies settings of a particular server in a stream upstream server group. Server parameters are specified in the JSON format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param streamUpstreamName The name of the upstream server group.
 @param streamUpstreamServerId The ID of the server.
 @param patchStreamUpstreamServer Server parameters, specified in the JSON format. The “*ID*”, “*backup*”, and “*service*” parameters cannot be changed.
 @return NginxStreamUpstreamConfServer*/
func (a *MethodPATCHApiService) PatchStreamUpstreamServer(ctx context.Context, streamUpstreamName string, streamUpstreamServerId string, patchStreamUpstreamServer NginxStreamUpstreamConfServer) (NginxStreamUpstreamConfServer,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NginxStreamUpstreamConfServer
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/stream/upstreams/{streamUpstreamName}/servers/{streamUpstreamServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamUpstreamName"+"}", fmt.Sprintf("%v", streamUpstreamName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"streamUpstreamServerId"+"}", fmt.Sprintf("%v", streamUpstreamServerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &patchStreamUpstreamServer
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}
