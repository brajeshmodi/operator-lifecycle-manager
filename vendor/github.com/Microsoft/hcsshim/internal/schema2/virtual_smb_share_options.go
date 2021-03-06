/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VirtualSmbShareOptions struct {
	ReadOnly bool `json:"ReadOnly,omitempty"`

	//  convert exclusive access to shared read access
	ShareRead bool `json:"ShareRead,omitempty"`

	//  all opens will use cached I/O
	CacheIo bool `json:"CacheIo,omitempty"`

	//  disable oplock support
	NoOplocks bool `json:"NoOplocks,omitempty"`

	//  Acquire the backup privilege when attempting to open
	TakeBackupPrivilege bool `json:"TakeBackupPrivilege,omitempty"`

	//  Use the identity of the share root when opening
	UseShareRootIdentity bool `json:"UseShareRootIdentity,omitempty"`

	//  disable Direct Mapping
	NoDirectmap bool `json:"NoDirectmap,omitempty"`

	//  disable Byterange locks
	NoLocks bool `json:"NoLocks,omitempty"`

	//  disable Directory CHange Notifications
	NoDirnotify bool `json:"NoDirnotify,omitempty"`

	//  share is use for VM shared memory
	VmSharedMemory bool `json:"VmSharedMemory,omitempty"`

	//  allow access only to the files specified in AllowedFiles
	RestrictFileAccess bool `json:"RestrictFileAccess,omitempty"`

	//  disable all oplocks except Level II
	ForceLevelIIOplocks bool `json:"ForceLevelIIOplocks,omitempty"`

	//  Allow the host to reparse this base layer
	ReparseBaseLayer bool `json:"ReparseBaseLayer,omitempty"`

	//  Enable pseudo-oplocks
	PseudoOplocks bool `json:"PseudoOplocks,omitempty"`

	//  All opens will use non-cached IO
	NonCacheIo bool `json:"NonCacheIo,omitempty"`

	//  Enable pseudo directory change notifications
	PseudoDirnotify bool `json:"PseudoDirnotify,omitempty"`

	//  Block directory enumeration, renames, and deletes.
	SingleFileMapping bool `json:"SingleFileMapping,omitempty"`
}
