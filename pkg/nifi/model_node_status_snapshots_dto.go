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

type NodeStatusSnapshotsDto struct {
	// The id of the node.
	NodeId string `json:"nodeId,omitempty"`
	// The node's host/ip address.
	Address string `json:"address,omitempty"`
	// The port the node is listening for API requests.
	ApiPort int32 `json:"apiPort,omitempty"`
	// A list of StatusSnapshotDTO objects that provide the actual metric values for the component for this node.
	StatusSnapshots []StatusSnapshotDto `json:"statusSnapshots,omitempty"`
}
