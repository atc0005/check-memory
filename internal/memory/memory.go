// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package memory

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/atc0005/go-nagios"
)

// Statistic represents a memory statistic (in KB) such as total memory or
// memory available.
type Statistic struct {
	Value int64
}

// MemInfo represents a collection of memory statistics for an OS such as
// total memory or memory available.
type MemInfo struct {
	criticalThreshold int
	warningThreshold  int
	Total             Statistic
	Available         Statistic
}

// MemPercentage represents the percentage for a memory statistic such as the
// percentage of available memory.
type MemPercentage struct {
	rawValue     float64
	roundedValue int
}

// String implements the Stringer interface to emit a percentage memory
// available value rounded down (so as to not overestimate available memory).
func (mp MemPercentage) String() string {
	return fmt.Sprintf("%d", mp.roundedValue)
}

// Rounded returns the percentage memory available value rounded down (so as
// to not overestimate available memory).
func (mp MemPercentage) Rounded() int {
	return mp.roundedValue
}

// Raw returns the percentage memory available value without rounding applied.
func (mp MemPercentage) Raw() float64 {
	return mp.rawValue
}

// // String implements the Stringer interface to emit a percentage memory
// // available value rounded down (so as to not overestimate available memory).
// func (mp MemPercentage) PercentageAvailable() string {
// 	memRoundDown := math.Floor((mp.Value * 100) / 100)

// 	return fmt.Sprintf("%.0f", memRoundDown)
// }

// String implements the Stringer interface to convert the default value (in
// KB) to MB for easier reading.
func (ms Statistic) String() string {
	memStatMB := ms.Value / 1024
	return fmt.Sprintf("%d MB", memStatMB)
}

// PercentageAvailable emits the percentage of available memory rounded down
// so as to not overestimate available memory.
func (mi MemInfo) PercentageAvailable() MemPercentage {
	// memPercentage := MemPercentage{
	// 	Value: float64(mi.Available.Value) / float64(mi.Total.Value) * 100,
	// }

	// fmt.Println("DEBUG (string): ", memPercentage.String())
	// fmt.Println("DEBUG (val): ", memPercentage.Value)

	rawPercentage := (float64(mi.Available.Value) / float64(mi.Total.Value)) * 100

	// percentage memory available value rounded down (so as to not
	// overestimate available memory).
	roundedPercentage := math.Floor((rawPercentage * 100) / 100)

	memPercentage := MemPercentage{
		rawValue:     rawPercentage,
		roundedValue: int(roundedPercentage),
	}

	return memPercentage
}

// IsOKState indicates whether evaluated memory was found to be in an OK or
// non-problematic state.
func (mi MemInfo) IsOKState() bool {
	return !mi.IsCriticalState() && !mi.IsWarningState()
}

// IsCriticalState indicates whether the percentage of available memory is less
// than then specified CRITICAL threshold.
func (mi MemInfo) IsCriticalState() bool {
	return mi.PercentageAvailable().roundedValue < mi.criticalThreshold
}

// IsWarningState indicates whether the percentage of available memory is less
// than then specified WARNING threshold.
func (mi MemInfo) IsWarningState() bool {
	return mi.PercentageAvailable().roundedValue < mi.warningThreshold
}

// ServiceState returns the appropriate Service Check Status label and exit
// code for the evaluation results.
func (mi MemInfo) ServiceState() nagios.ServiceState {
	var stateLabel string
	var stateExitCode int

	switch {
	case mi.IsCriticalState():
		stateLabel = nagios.StateCRITICALLabel
		stateExitCode = nagios.StateCRITICALExitCode
	case mi.IsWarningState():
		stateLabel = nagios.StateWARNINGLabel
		stateExitCode = nagios.StateWARNINGExitCode
	case mi.IsOKState():
		stateLabel = nagios.StateOKLabel
		stateExitCode = nagios.StateOKExitCode
	default:
		stateLabel = nagios.StateUNKNOWNLabel
		stateExitCode = nagios.StateUNKNOWNExitCode
	}

	return nagios.ServiceState{
		Label:    stateLabel,
		ExitCode: stateExitCode,
	}
}

// GetMemInfo returns a collection of memory statistics.
func GetMemInfo(criticalThreshold int, warningThreshold int) (MemInfo, error) {
	memAvailableVal, err := getMemAvailableFieldVal()
	if err != nil {
		return MemInfo{}, err
	}

	memTotalVal, err := getMemTotalFieldVal()
	if err != nil {
		return MemInfo{}, err
	}

	memAvailable := Statistic{
		Value: memAvailableVal,
	}

	memTotal := Statistic{
		Value: memTotalVal,
	}

	return MemInfo{
		criticalThreshold: criticalThreshold,
		warningThreshold:  warningThreshold,
		Total:             memTotal,
		Available:         memAvailable,
	}, nil
}

// getMemAvailableFieldVal returns the memory available in KB from the
// MemAvailable field within the /proc/meminfo file.
func getMemAvailableFieldVal() (int64, error) {
	memInfoFile := filepath.Join(ProcRootDir, ProcMemFilename)

	// Read the content of /proc/meminfo
	content, err := os.ReadFile(filepath.Clean(memInfoFile))
	if err != nil {
		return 0, fmt.Errorf("failed to read memory statistics file %s: %w", memInfoFile, err)
	}

	var memAvailableKB int64
	var found bool

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, ProcMemAvailableFieldName+":") {
			// Split the line into fields (e.g., "MemAvailable:", "10021900", "kB")
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// The second field is the value in kilobytes
				found = true
				memAvailableKB, err = strconv.ParseInt(fields[1], 10, 64)
				if err != nil {
					return 0, fmt.Errorf("failed to parse %s field: %w", ProcMemAvailableFieldName, err)
				}
				break // Found the value, exit the loop
			}
		}
	}

	if !found {
		return 0, fmt.Errorf("failed to find %s field in memory statistics file %s: %w", ProcMemTotalFieldName, memInfoFile, err)
	}

	return memAvailableKB, nil
}

func getMemTotalFieldVal() (int64, error) {
	memInfoFile := filepath.Join(ProcRootDir, ProcMemFilename)

	// Read the content of /proc/meminfo
	content, err := os.ReadFile(filepath.Clean(memInfoFile))
	if err != nil {
		return 0, fmt.Errorf("failed to read memory statistics file %s: %w", memInfoFile, err)
	}

	var memTotalKB int64
	var found bool

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, ProcMemTotalFieldName+":") {
			// Split the line into fields (e.g., "MemTotal:", "14073456", "kB")
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// The second field is the value in kilobytes
				found = true
				memTotalKB, err = strconv.ParseInt(fields[1], 10, 64)
				if err != nil {
					return 0, fmt.Errorf("failed to parse %s field: %w", ProcMemTotalFieldName, err)
				}
				break // Found the value, exit the loop
			}
		}
	}

	if !found {
		return 0, fmt.Errorf("failed to find %s field in memory statistics file %s: %w", ProcMemTotalFieldName, memInfoFile, err)
	}

	return memTotalKB, nil
}

// FallbackMemAvailableCalculation performs a calculation that provides an
// approximation of the `MemAvailable` field included in Linux kernel 3.14+ and
// backported to similar age OSes.
//
// See also:
//   - https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=34e431b0a
//   - https://unix.stackexchange.com/questions/261247/how-can-i-get-the-amount-of-available-memory-portably-across-distributions
func FallbackMemAvailableCalculation() int {
	// lowWaterMarkTotal := getLowWatermarkTotal()

	// Calculation as provided by Stephen Kitt using `awk`:
	//
	//	awk -v low=$(grep low /proc/zoneinfo | awk '{k+=$2}END{print k}') \
	//	 '{a[$1]=$2}
	//	  END{
	//	   print a["MemFree:"]+a["Active(file):"]+a["Inactive(file):"]+a["SReclaimable:"]-(12*low);
	//	  }' /proc/meminfo

	return 9999 // FIXME
}

// func getLowWatermarkTotal() int {
// 	// open ProcRootDir/ProcZoneInfoFilename
// 	// loop over every line which doesn't have `low` substring
// 	// for lines which have it, split into two columns: `low` and a numeric value
// 	// accumulate all numeric values
// 	// return

// 	return 9999 // FIXME
// }

/*
	TODO:

	gather all metrics from /proc/meminfo that are available; fetch via
	separate helper functions.

	Finish FallbackMemAvailableCalculation
*/
