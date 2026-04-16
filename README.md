<!-- omit in toc -->
# check-memory

Go-based tooling used to monitor memory usage.

[![Latest Release](https://img.shields.io/github/release/atc0005/check-memory.svg?style=flat-square)](https://github.com/atc0005/check-memory/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/check-memory.svg)](https://pkg.go.dev/github.com/atc0005/check-memory)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/check-memory)](https://github.com/atc0005/check-memory)
[![Lint and Build](https://github.com/atc0005/check-memory/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/check-memory/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/check-memory/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/check-memory/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
  - [`check_memory`](#check_memory)
    - [Performance Data](#performance-data)
- [Features](#features)
  - [`check_memory` plugin](#check_memory-plugin)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
  - [Deployment](#deployment)
- [Configuration](#configuration)
  - [Command-line arguments](#command-line-arguments)
    - [`check_memory`](#check_memory-1)
- [Examples](#examples)
  - [`OK` result](#ok-result)
  - [`WARNING` result](#warning-result)
  - [`CRITICAL` result](#critical-result)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo is intended to provide various tools used to monitor memory usage.

| Tool Name      | Overall Status | Description                                 |
| -------------- | -------------- | ------------------------------------------- |
| `check_memory` | Alpha          | Nagios plugin used to monitor memory usage. |

### `check_memory`

#### Performance Data

Initial support has been added for emitting Performance Data / Metrics, but
refinement suggestions are welcome.

Consult the table below for the metrics implemented thus far.

Please add to an existing
[Discussion](https://github.com/atc0005/check-memory/discussions) thread
(if applicable) or [open a new
one](https://github.com/atc0005/check-memory/discussions/new) with any
feedback that you may have. Thanks in advance!

| Emitted Performance Data / Metric | Meaning                      |
| --------------------------------- | ---------------------------- |
| `time`                            | Runtime for plugin           |
| `memory_available`                | Total memory available in KB |
| `memory_total`                    | Total memory in KB           |
| `memory_usage`                    | Percentage of memory used    |

## Features

### `check_memory` plugin

Nagios plugin (`check_memory`) used to monitor for problematic process states
on Linux distros.

> [!NOTE]
>
> The intent is to support multiple operating systems, but as of this writing
> Linux is the only supported OS.

- Optional branding "signature"
  - used to indicate what Nagios plugin (and what version) is responsible for
    the service check result

- Optional, leveled logging using `rs/zerolog` package
  - [`logfmt`][logfmt] format output (to `stderr`)
  - choice of `disabled`, `panic`, `fatal`, `error`, `warn`, `info` (the
    default), `debug` or `trace`

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for the *target* version this project
    was developed against
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Red Hat Enterprise Linux 8
- Ubuntu 22.04

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/check-memory`
   1. `cd check-memory`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - manually, explicitly specifying target OS and architecture
     - `GOOS=linux GOARCH=amd64 go build -mod=vendor ./cmd/check_memory/`
       - most likely this is what you want (if building manually)
       - substitute `amd64` with the appropriate architecture if using
         different hardware (e.g., `arm64` or `386`)
   - using Makefile `linux` recipe
     - `make linux`
       - generates x86 and x64 binaries
   - using Makefile `release-build` recipe
     - `make release-build`
       - generates the same release assets as provided by this project's
         releases
1. Locate generated binaries
   - if using `Makefile`
     - look in `/tmp/check-memory/release_assets/check_memory/`
   - if using `go build`
     - look in `/tmp/check-memory/`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**: Depending on which `Makefile` recipe you use the generated binary
may be compressed and have an `xz` extension. If so, you should decompress the
binary first before deploying it (e.g., `xz -d check_memory-linux-amd64.xz`).

### Using release binaries

1. Download the [latest release][repo-url] binaries
1. Decompress binaries
   - e.g., `xz -d check_memory-linux-amd64.xz`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

### Deployment

1. Place `check_memory` in a location where it can be executed by the
   monitoring agent
   - Usually the same place as other Nagios plugins
   - For example, on a default Red Hat Enterprise Linux system using
   `check_nrpe` the `check_memory` plugin would be deployed to
   `/usr/lib64/nagios/plugins/check_memory` or
   `/usr/local/nagios/libexec/check_memory`

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

## Configuration

### Command-line arguments

- Use the `-h` or `--help` flag to display current usage information.
- Flags marked as **`required`** must be set via CLI flag.
- Flags *not* marked as required are for settings where a useful default is
  already defined, but may be overridden if desired.

#### `check_memory`

| Flag              | Required | Default | Repeat | Possible                                                                | Description                                                                                          |
| ----------------- | -------- | ------- | ------ | ----------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `branding`        | No       | `false` | No     | `branding`                                                              | Toggles emission of branding details with plugin status details. This output is disabled by default. |
| `h`, `help`       | No       | `false` | No     | `h`, `help`                                                             | Show Help text along with the list of supported flags.                                               |
| `version`         | No       | `false` | No     | `version`                                                               | Whether to display application version and then immediately exit application.                        |
| `ll`, `log-level` | No       | `info`  | No     | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | Log message priority filter. Log messages with a lower level are ignored.                            |
| `c`, `critical`   | No       | `3`     | No     | *positive whole number between 0-100, exclusive*                        | Specifies the available memory percentage (as a whole number) when a CRITICAL threshold is reached.  |
| `w`, `warning`    | No       | `6`     | No     | *positive whole number between 0-100, exclusive*                        | Specifies the available memory percentage (as a whole number) when a WARNING threshold is reached.   |

## Examples

### `OK` result

This output is emitted by the plugin when memory availability is greater than
specified thresholds.

```console
$ ./check_memory -w 6 -c 3
OK: 1052 MB (53%) available memory | 'memory_available'=1077532;;;; 'memory_total'=2014912;;;; 'memory_usage'=47%;94;97;; 'time'=1ms;;;;

$ ./check_memory
OK: 1049 MB (53%) available memory | 'memory_available'=1075004;;;; 'memory_total'=2014912;;;; 'memory_usage'=47%;94;97;; 'time'=1ms;;;;
```

Regarding the output:

- Similar output is emitted by specifying explicit arguments (which match
  default threshold values) and running the plugin without arguments (relying
  solely on default values).
- The values and the `|` symbol are performance data metrics emitted by the
  plugin. Depending on your monitoring system, these metrics may be collected
  and exposed as graphs/charts.
- This output was captured on an Ubuntu 22.04 WSL instance. The output is
  comparable to other Linux distros.

### `WARNING` result

This output is emitted by the plugin when available memory drops below the
specified WARNING threshold. Here we use fake problem thresholds to trigger a
WARNING state.

```console
$ ./check_memory -w 54 -c 53
WARNING: 1048 MB (53%) available memory

**ERRORS**

* available memory below specified threshold
 | 'memory_available'=1073188;;;; 'memory_total'=2014912;;;; 'memory_usage'=47%;46;47;; 'time'=3ms;;;;
```

### `CRITICAL` result

This output is emitted by the plugin when available memory drops below the
specified CRITICAL threshold. Here we use fake problem thresholds to trigger a
CRITICAL state.

```console
$ ./check_memory -w 60 -c 54
CRITICAL: 1048 MB (53%) available memory

**ERRORS**

* available memory below specified threshold
 | 'memory_available'=1074036;;;; 'memory_total'=2014912;;;; 'memory_usage'=47%;40;46;; 'time'=2ms;;;;
```

## License

See the [LICENSE](LICENSE) file for details.

## References

- `MemAvailable` field added to `/proc/meminfo` in Linux kernel 3.14+ and
  backported to similar age OSes
  - <http://access.redhat.com/solutions/776393>
  - <https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/7/html-single/7.6_release_notes/index#new_features_kernel>
  - <https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=34e431b0a>
  - <https://unix.stackexchange.com/questions/261247/how-can-i-get-the-amount-of-available-memory-portably-across-distributions>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/check-memory>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

[logfmt]: <https://brandur.org/logfmt>
