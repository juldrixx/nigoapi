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

type ParameterDto struct {
	// The name of the Parameter
	Name string `json:"name,omitempty"`
	// The description of the Parameter
	Description *string `json:"description,omitempty"`
	// Whether or not the Parameter is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
	// The value of the Parameter
	Value *string `json:"value,omitempty"`
	// Whether or not the value of the Parameter was removed. When a request is made to change a parameter, the value may be null. The absence of the value may be used either to indicate that the value is not to be changed, or that the value is to be set to null (i.e., removed). This denotes which of the two scenarios is being encountered.
	ValueRemoved bool `json:"valueRemoved,omitempty"`
	// The set of all components in the flow that are referencing this Parameter
	ReferencingComponents []AffectedComponentEntity `json:"referencingComponents,omitempty"`
	// A reference to the Parameter Context that contains this one
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	// Whether or not the Parameter is inherited from another context
	Inherited bool `json:"inherited,omitempty"`
}
