/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The number of free and used memory pages.
type NginxSlabZonePages struct {

	// The current number of used memory pages.
	Used int32 `json:"used,omitempty"`

	// The current number of free memory pages.
	Free int32 `json:"free,omitempty"`
}