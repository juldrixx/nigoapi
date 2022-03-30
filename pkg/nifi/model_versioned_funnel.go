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

type VersionedFunnel struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component
	InstanceIdentifier string `json:"instanceIdentifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
