/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.16.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ProcessorDefinition struct {
	// The group name of the bundle that provides the referenced type.
	Group string `json:"group,omitempty"`
	// The artifact name of the bundle that provides the referenced type.
	Artifact string `json:"artifact,omitempty"`
	// The version of the bundle that provides the referenced type.
	Version string `json:"version,omitempty"`
	// The fully-qualified class type
	Type_ string `json:"type"`
	// The description of the type.
	TypeDescription string `json:"typeDescription,omitempty"`
	// The build metadata for this component
	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`
	// If this type represents a provider for an interface, this lists the APIs it implements
	ProvidedApiImplementations []DefinedType `json:"providedApiImplementations,omitempty"`
	// The tags associated with this type
	Tags []string `json:"tags,omitempty"`
	// Whether or not the component has been deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// If this component has been deprecated, this optional field can be used to provide an explanation
	DeprecationReason string `json:"deprecationReason,omitempty"`
	// Descriptions of configuration properties applicable to this component.
	PropertyDescriptors map[string]PropertyDescriptor `json:"propertyDescriptors,omitempty"`
	// Whether or not this component makes use of dynamic (user-set) properties.
	SupportsDynamicProperties bool `json:"supportsDynamicProperties,omitempty"`
	// Any input requirements this processor has.
	InputRequirement string `json:"inputRequirement,omitempty"`
	// The supported relationships for this processor.
	SupportedRelationships []Relationship `json:"supportedRelationships,omitempty"`
	// Whether or not this processor supports dynamic relationships.
	SupportsDynamicRelationships bool `json:"supportsDynamicRelationships,omitempty"`
	// Whether or not this processor should be triggered serially (i.e. no concurrent execution).
	TriggerSerially bool `json:"triggerSerially,omitempty"`
	// Whether or not this processor should be triggered when incoming queues are empty.
	TriggerWhenEmpty bool `json:"triggerWhenEmpty,omitempty"`
	// Whether or not this processor should be triggered when any destination queue has room.
	TriggerWhenAnyDestinationAvailable bool `json:"triggerWhenAnyDestinationAvailable,omitempty"`
	// Whether or not this processor supports batching. If a Processor uses this annotation, it allows the Framework to batch calls to session commits, as well as allowing the Framework to return the same session multiple times.
	SupportsBatching bool `json:"supportsBatching,omitempty"`
	// Whether or not this processor supports event driven scheduling. Indicates to the framework that the Processor is eligible to be scheduled to run based on the occurrence of an \"Event\" (e.g., when a FlowFile is enqueued in an incoming Connection), rather than being triggered periodically.
	SupportsEventDriven bool `json:"supportsEventDriven,omitempty"`
	// Whether or not this processor should be scheduled only on the primary node in a cluster.
	PrimaryNodeOnly bool `json:"primaryNodeOnly,omitempty"`
	// Whether or not this processor is considered side-effect free. Side-effect free indicate that the processor's operations on FlowFiles can be safely repeated across process sessions.
	SideEffectFree bool `json:"sideEffectFree,omitempty"`
	// The supported scheduling strategies, such as TIME_DRIVER, CRON, or EVENT_DRIVEN.
	SupportedSchedulingStrategies []string `json:"supportedSchedulingStrategies,omitempty"`
	// The default scheduling strategy for the processor.
	DefaultSchedulingStrategy string `json:"defaultSchedulingStrategy,omitempty"`
	// The default concurrent tasks for each scheduling strategy.
	DefaultConcurrentTasksBySchedulingStrategy map[string]int32 `json:"defaultConcurrentTasksBySchedulingStrategy,omitempty"`
	// The default scheduling period for each scheduling strategy. The scheduling period is expected to be a time period, such as \"30 sec\".
	DefaultSchedulingPeriodBySchedulingStrategy map[string]string `json:"defaultSchedulingPeriodBySchedulingStrategy,omitempty"`
	// The default penalty duration as a time period, such as \"30 sec\".
	DefaultPenaltyDuration string `json:"defaultPenaltyDuration,omitempty"`
	// The default yield duration as a time period, such as \"1 sec\".
	DefaultYieldDuration string `json:"defaultYieldDuration,omitempty"`
	// The default bulletin level, such as WARN, INFO, DEBUG, etc.
	DefaultBulletinLevel string `json:"defaultBulletinLevel,omitempty"`
}