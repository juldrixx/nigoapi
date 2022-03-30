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

type FlowConfigurationDto struct {
	// Whether this NiFi supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI.
	SupportsManagedAuthorizer bool `json:"supportsManagedAuthorizer,omitempty"`
	// Whether this NiFi supports a configurable authorizer.
	SupportsConfigurableAuthorizer bool `json:"supportsConfigurableAuthorizer,omitempty"`
	// Whether this NiFi supports configurable users and groups.
	SupportsConfigurableUsersAndGroups bool `json:"supportsConfigurableUsersAndGroups,omitempty"`
	// The interval in seconds between the automatic NiFi refresh requests.
	AutoRefreshIntervalSeconds int64 `json:"autoRefreshIntervalSeconds,omitempty"`
	// The current time on the system.
	CurrentTime string `json:"currentTime,omitempty"`
	// The time offset of the system.
	TimeOffset int32 `json:"timeOffset,omitempty"`
	// The default back pressure object threshold.
	DefaultBackPressureObjectThreshold int64 `json:"defaultBackPressureObjectThreshold,omitempty"`
	// The default back pressure data size threshold.
	DefaultBackPressureDataSizeThreshold string `json:"defaultBackPressureDataSizeThreshold,omitempty"`
}
