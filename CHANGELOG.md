# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Push to `giantswarm-operations-platform` catalog.

## [0.3.2] - 2024-01-29

### Fixed

- Move pss values under the global property

### Added

- Add team label in resource.

## [0.3.1] - 2023-12-05

### Fixed

- Disable PSP if PSS is in enforced mode.

## [0.3.0] - 2023-12-04

### Changed

- Address PSS concerns.

## [0.2.0] - 2023-11-10

### Changed

- Add a switch for PSP CR installation.

## [0.1.0] - 2023-01-19

### Fixed

- Add PSP and securityContext so pod can start in restricted namespaces.

### Changed

- Make kubeconfig flag configurable.

## [0.0.2] - 2020-09-02

### Changed

- Move namespace to `test-infra`.

## [0.0.1] - 2020-08-18

### Added

- Initial implementation.
- CI/CD automation.


[Unreleased]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.3.2...HEAD
[0.3.2]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/giantswarm/prow-log-aggregator/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/prow-log-aggregator/releases/tag/v0.0.1
