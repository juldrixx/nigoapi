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

type ProcessGroupStatusDto struct {
	// The ID of the Process Group
	Id string `json:"id,omitempty"`
	// The name of the Process Group
	Name string `json:"name,omitempty"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	// The aggregate status of all nodes in the cluster
	AggregateSnapshot *ProcessGroupStatusSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null.
	NodeSnapshots []NodeProcessGroupStatusSnapshotDto `json:"nodeSnapshots,omitempty"`
}
