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

type RemoteProcessGroupEntity struct {
	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDto `json:"revision,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	// The position of this component in the UI if applicable.
	Position *PositionDto `json:"position,omitempty"`
	// The permissions for this component.
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *RemoteProcessGroupDto `json:"component,omitempty"`
	// The status of the remote process group.
	Status *RemoteProcessGroupStatusDto `json:"status,omitempty"`
	// The number of remote input ports currently available on the target.
	InputPortCount int32 `json:"inputPortCount,omitempty"`
	// The number of remote output ports currently available on the target.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`
	// The permissions for this component operations.
	OperatePermissions *PermissionsDto `json:"operatePermissions,omitempty"`
}
