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

type VersionedProcessGroup struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// The child Process Groups
	ProcessGroups []VersionedProcessGroup `json:"processGroups,omitempty"`
	// The Remote Process Groups
	RemoteProcessGroups []VersionedRemoteProcessGroup `json:"remoteProcessGroups,omitempty"`
	// The Processors
	Processors []VersionedProcessor `json:"processors,omitempty"`
	// The Input Ports
	InputPorts []VersionedPort `json:"inputPorts,omitempty"`
	// The Output Ports
	OutputPorts []VersionedPort `json:"outputPorts,omitempty"`
	// The Connections
	Connections []VersionedConnection `json:"connections,omitempty"`
	// The Labels
	Labels []VersionedLabel `json:"labels,omitempty"`
	// The Funnels
	Funnels []VersionedFunnel `json:"funnels,omitempty"`
	// The Controller Services
	ControllerServices []VersionedControllerService `json:"controllerServices,omitempty"`
	// The coordinates where the remote flow is stored, or null if the Process Group is not directly under Version Control
	VersionedFlowCoordinates *VersionedFlowCoordinates `json:"versionedFlowCoordinates,omitempty"`
	// The Variables in the Variable Registry for this Process Group (not including any ancestor or descendant Process Groups)
	Variables map[string]string `json:"variables,omitempty"`
	// The name of the parameter context used by this process group
	ParameterContextName string `json:"parameterContextName,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The configured FlowFile Concurrency for the Process Group
	FlowFileConcurrency string `json:"flowFileConcurrency,omitempty"`
	// The FlowFile Outbound Policy for the Process Group
	FlowFileOutboundPolicy string `json:"flowFileOutboundPolicy,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
