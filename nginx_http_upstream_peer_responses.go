/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NginxHttpUpstreamPeerResponses struct {

	// The number of responses with “<code>1xx</code>” status codes.
	Var1xx int32 `json:"1xx,omitempty"`

	// The number of responses with “<code>2xx</code>” status codes.
	Var2xx int32 `json:"2xx,omitempty"`

	// The number of responses with “<code>3xx</code>” status codes.
	Var3xx int32 `json:"3xx,omitempty"`

	// The number of responses with “<code>4xx</code>” status codes.
	Var4xx int32 `json:"4xx,omitempty"`

	// The number of responses with “<code>5xx</code>” status codes.
	Var5xx int32 `json:"5xx,omitempty"`

	// The total number of responses obtained from this server.
	Total int32 `json:"total,omitempty"`
}
