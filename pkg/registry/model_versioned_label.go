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

type VersionedLabel struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// The text that appears in the label.
	Label string `json:"label,omitempty"`
	// The width of the label in pixels when at a 1:1 scale.
	Width float64 `json:"width,omitempty"`
	// The height of the label in pixels when at a 1:1 scale.
	Height float64 `json:"height,omitempty"`
	// The styles for this label (font-size : 12px, background-color : #eee, etc).
	Style map[string]string `json:"style,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
