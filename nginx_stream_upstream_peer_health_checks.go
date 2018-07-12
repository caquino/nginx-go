/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NginxStreamUpstreamPeerHealthChecks struct {

	// The total number of <a href=\"https://nginx.org/en/docs/stream/ngx_stream_upstream_hc_module.html#health_check\">health check</a> requests made.
	Checks int32 `json:"checks,omitempty"`

	// The number of failed health checks.
	Fails int32 `json:"fails,omitempty"`

	// How many times the server became unhealthy (state “<code>unhealthy</code>”).
	Unhealthy int32 `json:"unhealthy,omitempty"`

	// Boolean indicating whether the last health check request was successful and passed <a href=\"https://nginx.org/en/docs/stream/ngx_stream_upstream_hc_module.html#match\">tests</a>.
	LastPassed bool `json:"last_passed,omitempty"`
}
