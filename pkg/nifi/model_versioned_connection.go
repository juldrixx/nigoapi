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

type VersionedConnection struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// The source of the connection.
	Source *ConnectableComponent `json:"source,omitempty"`
	// The destination of the connection.
	Destination *ConnectableComponent `json:"destination,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex int32 `json:"labelIndex,omitempty"`
	// The z index of the connection.
	ZIndex int64 `json:"zIndex,omitempty"`
	// The selected relationship that comprise the connection.
	SelectedRelationships []string `json:"selectedRelationships,omitempty"`
	// The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureObjectThreshold int64 `json:"backPressureObjectThreshold,omitempty"`
	// The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureDataSizeThreshold string `json:"backPressureDataSizeThreshold,omitempty"`
	// The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
	FlowFileExpiration string `json:"flowFileExpiration,omitempty"`
	// The comparators used to prioritize the queue.
	Prioritizers []string `json:"prioritizers,omitempty"`
	// The bend points on the connection.
	Bends []Position `json:"bends,omitempty"`
	// The Strategy to use for load balancing data across the cluster, or null, if no Load Balance Strategy has been specified.
	LoadBalanceStrategy string `json:"loadBalanceStrategy,omitempty"`
	// The attribute to use for partitioning data as it is load balanced across the cluster. If the Load Balance Strategy is configured to use PARTITION_BY_ATTRIBUTE, the value returned by this method is the name of the FlowFile Attribute that will be used to determine which node in the cluster should receive a given FlowFile. If the Load Balance Strategy is unset or is set to any other value, the Partitioning Attribute has no effect.
	PartitioningAttribute string `json:"partitioningAttribute,omitempty"`
	// Whether or not compression should be used when transferring FlowFiles between nodes
	LoadBalanceCompression string `json:"loadBalanceCompression,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
