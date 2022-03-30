/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.16.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type VersionedParameter struct {
	// The name of the parameter
	Name string `json:"name,omitempty"`
	// The description of the param
	Description string `json:"description,omitempty"`
	// Whether or not the parameter value is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
	// The value of the parameter
	Value string `json:"value,omitempty"`
}
