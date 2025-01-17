# ProcessorDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | The group name of the bundle that provides the referenced type. | [optional] [default to null]
**Artifact** | **string** | The artifact name of the bundle that provides the referenced type. | [optional] [default to null]
**Version** | **string** | The version of the bundle that provides the referenced type. | [optional] [default to null]
**Type_** | **string** | The fully-qualified class type | [default to null]
**TypeDescription** | **string** | The description of the type. | [optional] [default to null]
**BuildInfo** | [***BuildInfo**](BuildInfo.md) | The build metadata for this component | [optional] [default to null]
**ProvidedApiImplementations** | [**[]DefinedType**](DefinedType.md) | If this type represents a provider for an interface, this lists the APIs it implements | [optional] [default to null]
**Tags** | **[]string** | The tags associated with this type | [optional] [default to null]
**Deprecated** | **bool** | Whether or not the component has been deprecated | [optional] [default to null]
**DeprecationReason** | **string** | If this component has been deprecated, this optional field can be used to provide an explanation | [optional] [default to null]
**PropertyDescriptors** | [**map[string]PropertyDescriptor**](PropertyDescriptor.md) | Descriptions of configuration properties applicable to this component. | [optional] [default to null]
**SupportsDynamicProperties** | **bool** | Whether or not this component makes use of dynamic (user-set) properties. | [optional] [default to null]
**InputRequirement** | **string** | Any input requirements this processor has. | [optional] [default to null]
**SupportedRelationships** | [**[]Relationship**](Relationship.md) | The supported relationships for this processor. | [optional] [default to null]
**SupportsDynamicRelationships** | **bool** | Whether or not this processor supports dynamic relationships. | [optional] [default to null]
**TriggerSerially** | **bool** | Whether or not this processor should be triggered serially (i.e. no concurrent execution). | [optional] [default to null]
**TriggerWhenEmpty** | **bool** | Whether or not this processor should be triggered when incoming queues are empty. | [optional] [default to null]
**TriggerWhenAnyDestinationAvailable** | **bool** | Whether or not this processor should be triggered when any destination queue has room. | [optional] [default to null]
**SupportsBatching** | **bool** | Whether or not this processor supports batching. If a Processor uses this annotation, it allows the Framework to batch calls to session commits, as well as allowing the Framework to return the same session multiple times. | [optional] [default to null]
**SupportsEventDriven** | **bool** | Whether or not this processor supports event driven scheduling. Indicates to the framework that the Processor is eligible to be scheduled to run based on the occurrence of an \&quot;Event\&quot; (e.g., when a FlowFile is enqueued in an incoming Connection), rather than being triggered periodically. | [optional] [default to null]
**PrimaryNodeOnly** | **bool** | Whether or not this processor should be scheduled only on the primary node in a cluster. | [optional] [default to null]
**SideEffectFree** | **bool** | Whether or not this processor is considered side-effect free. Side-effect free indicate that the processor&#39;s operations on FlowFiles can be safely repeated across process sessions. | [optional] [default to null]
**SupportedSchedulingStrategies** | **[]string** | The supported scheduling strategies, such as TIME_DRIVER, CRON, or EVENT_DRIVEN. | [optional] [default to null]
**DefaultSchedulingStrategy** | **string** | The default scheduling strategy for the processor. | [optional] [default to null]
**DefaultConcurrentTasksBySchedulingStrategy** | **map[string]int32** | The default concurrent tasks for each scheduling strategy. | [optional] [default to null]
**DefaultSchedulingPeriodBySchedulingStrategy** | **map[string]string** | The default scheduling period for each scheduling strategy. The scheduling period is expected to be a time period, such as \&quot;30 sec\&quot;. | [optional] [default to null]
**DefaultPenaltyDuration** | **string** | The default penalty duration as a time period, such as \&quot;30 sec\&quot;. | [optional] [default to null]
**DefaultYieldDuration** | **string** | The default yield duration as a time period, such as \&quot;1 sec\&quot;. | [optional] [default to null]
**DefaultBulletinLevel** | **string** | The default bulletin level, such as WARN, INFO, DEBUG, etc. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


