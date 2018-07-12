/*
 * NGINX Plus REST API
 *
 * NGINX Plus REST [API](https://nginx.org/en/docs/http/ngx_http_api_module.html) provides access to NGINX Plus status information, on-the-fly configuration of upstream servers and key-value pairs management for [http](https://nginx.org/en/docs/http/ngx_http_keyval_module.html) and [stream](https://nginx.org/en/docs/stream/ngx_stream_keyval_module.html).
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Synchronization status of a shared memory zone.
type NginxStreamZoneSyncZone struct {

	// The number of records that need to be sent to the cluster.
	RecordsPending int32 `json:"records_pending,omitempty"`

	// The total number of records stored in the shared memory zone.
	RecordsTotal int32 `json:"records_total,omitempty"`
}