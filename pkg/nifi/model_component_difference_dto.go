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

type ComponentDifferenceDto struct {
	// The type of component
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the component
	ComponentId string `json:"componentId,omitempty"`
	// The name of the component
	ComponentName string `json:"componentName,omitempty"`
	// The ID of the Process Group that the component belongs to
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// The differences in the component between the two flows
	Differences []DifferenceDto `json:"differences,omitempty"`
}
