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

type VersionedRemoteProcessGroup struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// [DEPRECATED] The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first uri in the uris. If neither target uri nor uris are set, then returns null.
	TargetUri string `json:"targetUri,omitempty"`
	// The target URIs of the remote process group. If target uris is not set but target uri is set, then returns the single target uri. If neither target uris nor target uri is set, then returns null.
	TargetUris string `json:"targetUris,omitempty"`
	// The time period used for the timeout when communicating with the target.
	CommunicationsTimeout string `json:"communicationsTimeout,omitempty"`
	// When yielding, this amount of time must elapse before the remote process group is scheduled again.
	YieldDuration string `json:"yieldDuration,omitempty"`
	// The Transport Protocol that is used for Site-to-Site communications
	TransportProtocol string `json:"transportProtocol,omitempty"`
	// The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier.
	LocalNetworkInterface string `json:"localNetworkInterface,omitempty"`
	ProxyHost string `json:"proxyHost,omitempty"`
	ProxyPort int32 `json:"proxyPort,omitempty"`
	ProxyUser string `json:"proxyUser,omitempty"`
	// A Set of Input Ports that can be connected to, in order to send data to the remote NiFi instance
	InputPorts []VersionedRemoteGroupPort `json:"inputPorts,omitempty"`
	// A Set of Output Ports that can be connected to, in order to pull data from the remote NiFi instance
	OutputPorts []VersionedRemoteGroupPort `json:"outputPorts,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
