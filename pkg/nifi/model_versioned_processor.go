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

type VersionedProcessor struct {
	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`
	// The component's name
	Name string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`
	// The component's position on the graph
	Position *Position `json:"position,omitempty"`
	// Information about the bundle from which the component came
	Bundle *Bundle `json:"bundle,omitempty"`
	// Stylistic data for rendering in a UI
	Style map[string]string `json:"style,omitempty"`
	// The type of Processor
	Type_ string `json:"type,omitempty"`
	// The properties for the processor. Properties whose value is not set will only contain the property name.
	Properties map[string]string `json:"properties,omitempty"`
	// The property descriptors for the processor.
	PropertyDescriptors map[string]VersionedPropertyDescriptor `json:"propertyDescriptors,omitempty"`
	// The annotation data for the processor used to relay configuration between a custom UI and the procesosr.
	AnnotationData string `json:"annotationData,omitempty"`
	// The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.
	SchedulingPeriod string `json:"schedulingPeriod,omitempty"`
	// Indcates whether the prcessor should be scheduled to run in event or timer driven mode.
	SchedulingStrategy string `json:"schedulingStrategy,omitempty"`
	// Indicates the node where the process will execute.
	ExecutionNode string `json:"executionNode,omitempty"`
	// The amout of time that is used when the process penalizes a flowfile.
	PenaltyDuration string `json:"penaltyDuration,omitempty"`
	// The amount of time that must elapse before this processor is scheduled again after yielding.
	YieldDuration string `json:"yieldDuration,omitempty"`
	// The level at which the processor will report bulletins.
	BulletinLevel string `json:"bulletinLevel,omitempty"`
	// The run duration for the processor in milliseconds.
	RunDurationMillis int64 `json:"runDurationMillis,omitempty"`
	// The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.
	AutoTerminatedRelationships []string `json:"autoTerminatedRelationships,omitempty"`
	// The scheduled state of the component
	ScheduledState string `json:"scheduledState,omitempty"`
	ComponentType string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`
}
