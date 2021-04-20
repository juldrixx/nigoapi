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

type RegistryConfiguration struct {
	// Whether this NiFi Registry supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI.
	SupportsManagedAuthorizer bool `json:"supportsManagedAuthorizer,omitempty"`
	// Whether this NiFi Registry supports a configurable authorizer.
	SupportsConfigurableAuthorizer bool `json:"supportsConfigurableAuthorizer,omitempty"`
	// Whether this NiFi Registry supports configurable users and groups.
	SupportsConfigurableUsersAndGroups bool `json:"supportsConfigurableUsersAndGroups,omitempty"`
}
