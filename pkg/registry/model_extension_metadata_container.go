/*
 * Apache NiFi Registry REST API
 *
 * The REST API provides an interface to a registry with operations for saving, versioning, reading NiFi flows and components.
 *
 * API version: 0.8.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package registry

type ExtensionMetadataContainer struct {
	// The number of extensions in the response
	NumResults int32 `json:"numResults,omitempty"`
	// The filter parameters submitted for the request
	FilterParams *ExtensionFilterParams `json:"filterParams,omitempty"`
	// The metadata for the extensions
	Extensions []ExtensionMetadata `json:"extensions,omitempty"`
}
