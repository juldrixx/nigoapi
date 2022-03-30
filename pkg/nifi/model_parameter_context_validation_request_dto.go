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

type ParameterContextValidationRequestDto struct {
	// The ID of the request
	RequestId string `json:"requestId,omitempty"`
	// The URI for the request
	Uri string `json:"uri,omitempty"`
	// The timestamp of when the request was submitted
	SubmissionTime string `json:"submissionTime,omitempty"`
	// The timestamp of when the request was last updated
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Whether or not the request is completed
	Complete bool `json:"complete,omitempty"`
	// The reason for the request failing, or null if the request has not failed
	FailureReason string `json:"failureReason,omitempty"`
	// A value between 0 and 100 (inclusive) indicating how close the request is to completion
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// A description of the current state of the request
	State string `json:"state,omitempty"`
	// The steps that are required in order to complete the request, along with the status of each
	UpdateSteps []ParameterContextValidationStepDto `json:"updateSteps,omitempty"`
	// The Parameter Context that is being operated on.
	ParameterContext *ParameterContextDto `json:"parameterContext,omitempty"`
	// The Validation Results that were calculated for each component. This value may not be set until the request completes.
	ComponentValidationResults *ComponentValidationResultsEntity `json:"componentValidationResults,omitempty"`
}
