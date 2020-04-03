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

type ProcessGroupStatusSnapshotDto struct {
	// The id of the process group.
	Id string `json:"id,omitempty"`
	// The name of this process group.
	Name string `json:"name,omitempty"`
	// The status of all connections in the process group.
	ConnectionStatusSnapshots []ConnectionStatusSnapshotEntity `json:"connectionStatusSnapshots,omitempty"`
	// The status of all processors in the process group.
	ProcessorStatusSnapshots []ProcessorStatusSnapshotEntity `json:"processorStatusSnapshots,omitempty"`
	// The status of all process groups in the process group.
	ProcessGroupStatusSnapshots []ProcessGroupStatusSnapshotEntity `json:"processGroupStatusSnapshots,omitempty"`
	// The status of all remote process groups in the process group.
	RemoteProcessGroupStatusSnapshots []RemoteProcessGroupStatusSnapshotEntity `json:"remoteProcessGroupStatusSnapshots,omitempty"`
	// The status of all input ports in the process group.
	InputPortStatusSnapshots []PortStatusSnapshotEntity `json:"inputPortStatusSnapshots,omitempty"`
	// The status of all output ports in the process group.
	OutputPortStatusSnapshots []PortStatusSnapshotEntity `json:"outputPortStatusSnapshots,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	VersionedFlowState string `json:"versionedFlowState,omitempty"`
	// The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`
	// The number of bytes that have come into this ProcessGroup in the last 5 minutes
	BytesIn int64 `json:"bytesIn,omitempty"`
	// The input count/size for the process group in the last 5 minutes (pretty printed).
	Input string `json:"input,omitempty"`
	// The number of FlowFiles that are queued up in this ProcessGroup right now
	FlowFilesQueued int32 `json:"flowFilesQueued,omitempty"`
	// The number of bytes that are queued up in this ProcessGroup right now
	BytesQueued int64 `json:"bytesQueued,omitempty"`
	// The count/size that is queued in the the process group.
	Queued string `json:"queued,omitempty"`
	// The count that is queued for the process group.
	QueuedCount string `json:"queuedCount,omitempty"`
	// The size that is queued for the process group.
	QueuedSize string `json:"queuedSize,omitempty"`
	// The number of bytes read by components in this ProcessGroup in the last 5 minutes
	BytesRead int64 `json:"bytesRead,omitempty"`
	// The number of bytes read in the last 5 minutes.
	Read string `json:"read,omitempty"`
	// The number of bytes written by components in this ProcessGroup in the last 5 minutes
	BytesWritten int64 `json:"bytesWritten,omitempty"`
	// The number of bytes written in the last 5 minutes.
	Written string `json:"written,omitempty"`
	// The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`
	// The number of bytes transferred out of this ProcessGroup in the last 5 minutes
	BytesOut int64 `json:"bytesOut,omitempty"`
	// The output count/size for the process group in the last 5 minutes.
	Output string `json:"output,omitempty"`
	// The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes
	FlowFilesTransferred int32 `json:"flowFilesTransferred,omitempty"`
	// The number of bytes transferred in this ProcessGroup in the last 5 minutes
	BytesTransferred int64 `json:"bytesTransferred,omitempty"`
	// The count/size transferred to/from queues in the process group in the last 5 minutes.
	Transferred string `json:"transferred,omitempty"`
	// The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes
	BytesReceived int64 `json:"bytesReceived,omitempty"`
	// The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes
	FlowFilesReceived int32 `json:"flowFilesReceived,omitempty"`
	// The count/size sent to the process group in the last 5 minutes.
	Received string `json:"received,omitempty"`
	// The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes
	BytesSent int64 `json:"bytesSent,omitempty"`
	// The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes
	FlowFilesSent int32 `json:"flowFilesSent,omitempty"`
	// The count/size sent from this process group in the last 5 minutes.
	Sent string `json:"sent,omitempty"`
	// The active thread count for this process group.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The number of threads currently terminated for the process group.
	TerminatedThreadCount int32 `json:"terminatedThreadCount,omitempty"`
}
