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

import (
	"time"
)

type GcDiagnosticsSnapshotDto struct {
	// The timestamp of when the Snapshot was taken
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The number of times that Garbage Collection has occurred
	CollectionCount int64 `json:"collectionCount,omitempty"`
	// The number of milliseconds that the Garbage Collector spent performing Garbage Collection duties
	CollectionMillis int64 `json:"collectionMillis,omitempty"`
}
