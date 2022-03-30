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

type RepositoryUsageDto struct {
	// The name of the repository
	Name string `json:"name,omitempty"`
	// A SHA-256 hash of the File Store name/path that is used to store the repository's data. This information is exposed as a hash in order to avoid exposing potentially sensitive information that is not generally relevant. What is typically relevant is whether or not multiple repositories on the same node are using the same File Store, as this indicates that the repositories are competing for the resources of the backing disk/storage mechanism.
	FileStoreHash string `json:"fileStoreHash,omitempty"`
	// Amount of free space.
	FreeSpace string `json:"freeSpace,omitempty"`
	// Amount of total space.
	TotalSpace string `json:"totalSpace,omitempty"`
	// The number of bytes of free space.
	FreeSpaceBytes int64 `json:"freeSpaceBytes,omitempty"`
	// The number of bytes of total space.
	TotalSpaceBytes int64 `json:"totalSpaceBytes,omitempty"`
	// Utilization of this storage location.
	Utilization string `json:"utilization,omitempty"`
}
