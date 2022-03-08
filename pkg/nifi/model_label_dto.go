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

type LabelDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	// The position of this component in the UI if applicable.
	Position *PositionDto `json:"position,omitempty"`
	// The text that appears in the label.
	Label string `json:"label,omitempty"`
	// The width of the label in pixels when at a 1:1 scale.
	Width float64 `json:"width,omitempty"`
	// The height of the label in pixels when at a 1:1 scale.
	Height float64 `json:"height,omitempty"`
	// The styles for this label (font-size : 12px, background-color : #eee, etc).
	Style map[string]string `json:"style,omitempty"`
}
