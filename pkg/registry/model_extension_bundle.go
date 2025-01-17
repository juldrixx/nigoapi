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

type ExtensionBundle struct {
	// An WebLink to this entity.
	Link *JaxbLink `json:"link,omitempty"`
	// An ID to uniquely identify this object.
	Identifier string `json:"identifier,omitempty"`
	// The name of the item.
	Name string `json:"name"`
	// A description of the item.
	Description string `json:"description,omitempty"`
	// The identifier of the bucket this items belongs to. This cannot be changed after the item is created.
	BucketIdentifier string `json:"bucketIdentifier"`
	// The name of the bucket this items belongs to.
	BucketName string `json:"bucketName,omitempty"`
	// The timestamp of when the item was created, as milliseconds since epoch.
	CreatedTimestamp int64 `json:"createdTimestamp,omitempty"`
	// The timestamp of when the item was last modified, as milliseconds since epoch.
	ModifiedTimestamp int64 `json:"modifiedTimestamp,omitempty"`
	// The type of item.
	Type_ string `json:"type"`
	// The access that the current user has to the bucket containing this item.
	Permissions *Permissions `json:"permissions,omitempty"`
	// The type of the extension bundle
	BundleType string `json:"bundleType"`
	// The group id of the extension bundle
	GroupId string `json:"groupId,omitempty"`
	// The artifact id of the extension bundle
	ArtifactId string `json:"artifactId,omitempty"`
	// The number of versions of this extension bundle.
	VersionCount int64 `json:"versionCount,omitempty"`
}
