// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package memory

import "errors"

var (
	// ErrAvailableMemoryBelowThreshold indicates that evaluated memory is
	// below a given set of thresholds.
	ErrAvailableMemoryBelowThreshold = errors.New("available memory below specified threshold")
)
