// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"os"
)

// supportedValuesFlagHelpText is a flag package helper function that combines
// base help text with a list of supported values for the flag.
func supportedValuesFlagHelpText(baseHelpText string, supportedValues []string) string {
	return fmt.Sprintf(
		"%s Supported values: %v",
		baseHelpText,
		supportedValues,
	)
}

// handleFlagsConfig handles toggling the exposure of specific configuration
// flags to the user. This behavior is controlled via the specified
// application type as set by each cmd. Based on the application's specified
// type, a smaller subset of flags specific to each type are exposed along
// with a set common to all application types.
func (c *Config) handleFlagsConfig(appType AppType) error {
	if c == nil {
		return fmt.Errorf(
			"nil configuration, cannot process flags: %w",
			ErrConfigNotInitialized,
		)
	}

	// shared flags
	c.flagSet.BoolVar(&c.ShowHelp, HelpFlagShort, defaultHelp, helpFlagHelp+shorthandFlagSuffix)
	c.flagSet.BoolVar(&c.ShowHelp, HelpFlagLong, defaultHelp, helpFlagHelp)

	c.flagSet.BoolVar(&c.ShowVersion, VersionFlagLong, defaultDisplayVersionAndExit, versionFlagHelp)

	c.flagSet.StringVar(
		&c.LoggingLevel,
		LogLevelFlagShort,
		defaultLogLevel,
		supportedValuesFlagHelpText(logLevelFlagHelp, supportedLogLevels())+shorthandFlagSuffix,
	)
	c.flagSet.StringVar(
		&c.LoggingLevel,
		LogLevelFlagLong,
		defaultLogLevel,
		supportedValuesFlagHelpText(logLevelFlagHelp, supportedLogLevels()),
	)

	switch {
	case appType.Inspector:
		// placeholder for future sysadmin CLI app
	case appType.Plugin:
		c.flagSet.BoolVar(&c.EmitBranding, BrandingFlag, defaultEmitBranding, brandingFlagHelp)

		c.flagSet.IntVar(&c.CriticalThreshold, memAvailableCriticalFlagShort, defaultMemAvailableCriticalThreshold, memAvailableCriticalFlagHelp+shorthandFlagSuffix)
		c.flagSet.IntVar(&c.CriticalThreshold, memAvailableCriticalFlagLong, defaultMemAvailableCriticalThreshold, memAvailableCriticalFlagHelp)
		c.flagSet.IntVar(&c.WarningThreshold, memAvailableWarningFlagShort, defaultMemAvailableWarningThreshold, memAvailableWarningFlagHelp+shorthandFlagSuffix)
		c.flagSet.IntVar(&c.WarningThreshold, memAvailableWarningFlagLong, defaultMemAvailableWarningThreshold, memAvailableWarningFlagHelp)
	}

	// Allow our function to override the default Help output.
	//
	// Override default of stderr as destination for help output. This allows
	// Nagios XI and similar monitoring systems to call plugins with the
	// `--help` flag and have it display within the Admin web UI.
	c.flagSet.Usage = Usage(c.flagSet, os.Stdout)

	// parse flag definitions from the argument list
	return c.flagSet.Parse(os.Args[1:])

}
