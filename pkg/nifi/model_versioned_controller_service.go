/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.15.3
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type VersionedControllerService struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// The type of the controller service.
	Type_ string `json:"type,omitempty"`
	// The details of the artifact that bundled this processor type.
	Bundle *Bundle `json:"bundle,omitempty"`
	// Lists the APIs this Controller Service implements.
	ControllerServiceApis []ControllerServiceApi `json:"controllerServiceApis,omitempty"`
	// The properties of the controller service.
	Properties map[string]string `json:"properties,omitempty"`
	// The property descriptors for the processor.
	PropertyDescriptors map[string]VersionedPropertyDescriptor `json:"propertyDescriptors,omitempty"`
	// The annotation for the controller service. This is how the custom UI relays configuration to the controller service.
	AnnotationData string `json:"annotationData,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
