/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ProcessGroupReplaceRequestEntity struct {
	// The revision for the Process Group being updated.
	ProcessGroupRevision *RevisionDto `json:"processGroupRevision,omitempty"`
	// The Process Group Change Request
	Request *ProcessGroupReplaceRequestDto `json:"request,omitempty"`
	// Returns the Versioned Flow to replace with
	VersionedFlowSnapshot *VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
}
