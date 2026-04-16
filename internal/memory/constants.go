// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package memory

const (
	// ProcRootDir is the default mount point for the proc virtual filesystem.
	ProcRootDir string = "/proc"

	// ProcMemFilename is the name of the file in the proc virtual filesystem
	// containing human-readable memory usage information which resides in the
	// proc filesystem.
	ProcMemFilename string = "meminfo"

	// ProcZoneInfoFilename is the name of the file in the proc virtual
	// filesystem containing detailed statistics and information about memory
	// zones within the system.
	ProcZoneInfoFilename string = "zoneinfo"
)

const (
	// ProcMemAvailableFieldName is the modern, recommended metric. Added in
	// Linux kernel 3.14 and backported to RHEL 7 (based on older 3.10 kernel)
	// and other similar age OSes. This metric was notably also backported to
	// the 2.6.32 kernel in RHEL 6.
	//
	// See also:
	//
	//   - http://access.redhat.com/solutions/776393
	//   - https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/7/html-single/7.6_release_notes/index#new_features_kernel
	ProcMemAvailableFieldName string = "MemAvailable"

	// ProcMemTotalFieldName is the total memory exposed to the OS.
	ProcMemTotalFieldName string = "MemTotal"
)

// Fallback metric values used to calculate free memory.
const (
	ProcMemFreeFieldName         string = "MemFree"
	ProcMemBuffersFieldName      string = "Buffers"
	ProcMemCachedFieldName       string = "Cached"
	ProcMemSReclaimableFieldName string = "SReclaimable"
	ProcMemActiveFileFieldName   string = "Active(file)"
	ProcMemInactiveFileFieldName string = "Inactive(file)"
)
