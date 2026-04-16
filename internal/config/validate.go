// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"

	"github.com/atc0005/check-memory/internal/textutils"
)

// validate verifies all Config struct fields have been provided acceptable
// values.
func (c Config) validate(appType AppType) error {
	// Validate the specified logging level
	supportedLogLevels := supportedLogLevels()
	if !textutils.InList(c.LoggingLevel, supportedLogLevels, true) {
		return fmt.Errorf(
			"%w: invalid logging level;"+
				" got %v, expected one of %v",
			ErrUnsupportedOption,
			c.LoggingLevel,
			supportedLogLevels,
		)
	}

	switch {
	case appType.Inspector:

	case appType.Plugin:
		if c.WarningThreshold <= 0 || c.WarningThreshold > 100 {
			return fmt.Errorf(
				"invalid memory available percentage WARNING threshold number: %d",
				c.WarningThreshold,
			)
		}

		if c.CriticalThreshold <= 0 || c.CriticalThreshold > 100 {
			return fmt.Errorf(
				"invalid memory available percentage CRITICAL threshold number: %d",
				c.CriticalThreshold,
			)
		}

		if c.CriticalThreshold >= c.WarningThreshold {
			return fmt.Errorf(
				"critical threshold set higher than or equal to warning threshold",
			)
		}
	}

	// Optimist
	return nil

}
