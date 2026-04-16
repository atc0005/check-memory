// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"

	"github.com/atc0005/check-memory/internal/memory"
	"github.com/atc0005/go-nagios"
)

// getPerfData gathers performance data metrics that we wish to report. This
// function relies on the caller to ensure specified memory available
// threshold values are acceptable.
func getPerfData(memInfo memory.MemInfo, memoryPercentAvailableCritical int, memoryPercentAvailableWarning int) []nagios.PerformanceData {
	memoryPercentUsed := func() int {
		memoryPercentAvailable := memInfo.PercentageAvailable().Rounded()
		used := 100 - memoryPercentAvailable
		if used < 0 {
			used = 0
		}

		return used
	}()

	memUsageCritical := func() int {
		return 100 - memoryPercentAvailableCritical
	}()

	memUsageWarning := func() int {
		return 100 - memoryPercentAvailableWarning
	}()

	return []nagios.PerformanceData{
		// The `time` (runtime) metric is appended at plugin exit, so do not
		// duplicate it here.
		{
			Label: "memory_total",
			Value: fmt.Sprintf("%d", memInfo.Total.Value),
		},
		{
			Label: "memory_available",
			Value: fmt.Sprintf("%d", memInfo.Available.Value),
		},
		{
			// We report the percentage of memory used/unavailable instead of
			// what is available. This acts as the inverse of the provided
			// thresholds based on the assumption that this be more intuitive
			// to reason about when graphed.
			Label:             "memory_usage",
			UnitOfMeasurement: "%",
			Value:             fmt.Sprintf("%d", memoryPercentUsed),
			Crit:              fmt.Sprintf("%d", memUsageCritical),
			Warn:              fmt.Sprintf("%d", memUsageWarning),
		},
	}
}
