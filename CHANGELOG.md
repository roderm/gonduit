# Changelog

All notable changes to this project will be documented in this file.

## [0.6.0] - 2019-11-26
### Changed
- `DifferentialRevision.Reviewers` now points not to map but
  `DifferentialRevisionReviewers` struct (which is map also).

### Fixed
- `differential.query` method does not fail anymore if revision has no
  reviewers.

## [0.5.0] - 2019-10-14
### Added
- Support for differential.getcommitmessage.

## [0.4.1] - 2019-10-08
### Added
- Support for differential.getcommitpaths.

## [0.4.0] - 2019-07-12
### Added
- Support for differential.querydiffs.
- Timeout field to code.ClientOptions.
- DifferentialStatusLegacy with int representations of statuses.
- Client interface to pass own http.Client instance.
- Introduced basic context.Context compatability.

### Changed
- Changed fields on entities.DifferentialRevision to match actual response
  returned from Phabricator. This is breaking change.

### Removed
- DifferentialStatus struct as it is not used anymore.

## [0.3.3] - 2019-06-07
### Added
- Added support for `maniphest.search` endpoint.

## [0.3.2] - 2019-01-31
### Fixed
- Return `ConduitError` with proper status code when Phabricator fails with
  HTML output and client can not parse JSON.

## [0.3.1] - 2019-01-08
### Added
- Added `Email` value to `entities.User` struct for response to `user.query`
  endpoint.

## [0.3.0] - 2018-11-19
### Changed
- Changed import paths from `etcinit` to `uber`.
- Updated vesions of dependencies.
