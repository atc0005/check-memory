// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/check-memory
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

const myAppName string = "check-memory"
const myAppURL string = "https://github.com/atc0005/check-memory"

// Shared flag help text
const (
	helpFlagHelp                 string = "Emit this help text"
	versionFlagHelp              string = "Whether to display application version and then immediately exit application."
	logLevelFlagHelp             string = "Sets log level."
	brandingFlagHelp             string = "Toggles emission of branding details with plugin status details. This output is disabled by default."
	memAvailableCriticalFlagHelp string = "Specifies the available memory percentage (as a whole number) when a CRITICAL threshold is reached."
	memAvailableWarningFlagHelp  string = "Specifies the available memory percentage (as a whole number) when a WARNING threshold is reached."
)

// shorthandFlagSuffix is appended to short flag help text to emphasize that
// the flag is a shorthand version of a longer flag.
const shorthandFlagSuffix = " (shorthand)"

// Flag names for consistent references. Exported so that they're available
// from tests.
const (
	HelpFlagLong                  string = "help"
	HelpFlagShort                 string = "h"
	VersionFlagLong               string = "version"
	BrandingFlag                  string = "branding"
	TimeoutFlagLong               string = "timeout"
	TimeoutFlagShort              string = "t"
	LogLevelFlagLong              string = "log-level"
	LogLevelFlagShort             string = "ll"
	memAvailableCriticalFlagShort string = "c"
	memAvailableCriticalFlagLong  string = "critical"
	memAvailableWarningFlagShort  string = "w"
	memAvailableWarningFlagLong   string = "warning"
)

// Default flag settings if not overridden by user input
const (
	defaultHelp                          bool   = false
	defaultLogLevel                      string = "info"
	defaultEmitBranding                  bool   = false
	defaultDisplayVersionAndExit         bool   = false
	defaultMemAvailableWarningThreshold  int    = 6 // borrow existing default value used by check_mem.sh and check_mem_new.sh scripts
	defaultMemAvailableCriticalThreshold int    = 3 // borrow existing default value used by check_mem.sh and check_mem_new.sh scripts
)

const (
	appTypePlugin    string = "plugin"
	appTypeInspector string = "Inspector"
)
