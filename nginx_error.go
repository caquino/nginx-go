/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// nginx error object. 
type NginxError struct {

	// API path.
	Path string `json:"path,omitempty"`

	// HTTP method.
	Method string `json:"method,omitempty"`

	Error_ *NginxErrorError `json:"error,omitempty"`

	// The ID of the request, equals the value of the <a href=\"https://nginx.org/en/docs/http/ngx_http_core_module.html#var_request_id\">$request_id</a> variable.
	RequestId string `json:"request_id,omitempty"`

	// Link to reference documentation.
	Href string `json:"href,omitempty"`
}
