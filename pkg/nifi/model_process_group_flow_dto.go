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

type ProcessGroupFlowDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	// The Parameter Context, or null if no Parameter Context has been bound to the Process Group
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	// The breadcrumb of the process group.
	Breadcrumb *FlowBreadcrumbEntity `json:"breadcrumb,omitempty"`
	// The flow structure starting at this Process Group.
	Flow *FlowDto `json:"flow,omitempty"`
	// The time the flow for the process group was last refreshed.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
}
