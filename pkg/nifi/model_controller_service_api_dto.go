/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.11.1-SNAPSHOT
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ControllerServiceApiDto struct {
	// The fully qualified name of the service interface.
	Type_ string `json:"type,omitempty"`
	// The details of the artifact that bundled this service interface.
	Bundle *BundleDto `json:"bundle,omitempty"`
}
