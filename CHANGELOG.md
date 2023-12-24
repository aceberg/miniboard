
# Change Log
All notable changes to this project will be documented in this file.

## [1.2.2] - 2023-12-24
### Added
- Horizontal layout for tabs

## [1.2.1] - 2023-12-04
### Fixed
- Memory leak fix [Issue #11](https://github.com/aceberg/miniboard/issues/11)

## [1.2.0] - 2023-11-26
### Added
- Session cookie auth [Issue #10](https://github.com/aceberg/miniboard/issues/10)
- Binaries with Goreleaser

## [1.1.0] - 2023-11-03
### Added
- Local selfhosted icons ([README](https://github.com/aceberg/miniboard#local-network-only))
- Uptime history stored in SQLite DB

## [1.0.9] - 2023-09-13
### Fixed
- Docker image changed from `scratch` to `alpine`, to fix `shoutrrr` notifications problem

## [1.0.8] - 2023-08-23
### Changed
- Monospace font in edit board file [Issue #3](https://github.com/aceberg/miniboard/issues/3)
- Edit scan and timeout on panel page

## [1.0.7] - 2023-07-23
### Fixed
- Uptime sort issue

## [1.0.6] - 2023-07-20
### Fixed
- Uptime sorted by Date and Time

### Change
- Reload function refactoring

## [1.0.5] - 2023-07-19
### Added
- Uptime: split date and time
- Uptime: filter by date, address, port, notification
- Edit board file online

### Changed
- Connection timeout to 3 seconds

### Fixed
- Error: concurrent map read and map write

## [1.0.4] - 2023-07-14
### Added
- Uptime: color Date
- Uptime: list services to notify
- Tabs: separate refresh config for each tab

## [1.0.3] - 2023-07-10
### Added
- Scan goroutine reload (`Reload` in the menu)
- Version on config page

### Fixed
- Auto add new panels to scan
- Auto remove deleted panels from scan

## [1.0.2] - 2023-07-10
### Added
- Binary user guide
### Fixed
- Send test notification bug

## [1.0.1] - 2023-07-09
### Added
- Button to automatically create panel from Docker containers (On `Edit panels` page)

## [1.0.0] - 2023-07-08
### Added
- Uptime monitoring
- Timeout for each panel
- GUI config page for Uptime
- Web page refresh interval
- Notifications to services, supported by [Shoutrrr](https://containrrr.dev/shoutrrr/0.7/services/overview/)

### Fixed
- Error: concurrent map writes

### Changed
- Code refactoring

## [0.2.1] - 2023-06-28
### Added
- Binary and Deb release

## [0.2.0] - 2023-06-27
### Added
- Edit everything in GUI
- Local node modules

## [0.1.0] - 2023-06-20
### Added
- Board config with web GUI
- Scan if host is online

