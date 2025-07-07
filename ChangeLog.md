# Changelog

## v2.x.x (2024-xx-xx)
- New features
  - Distributed plug-ins
    - Storage plug-ins
      - Add cache store plug-in
    - Coordinator plug-ins
      - Added etcd plug-in
- Improvements
  - Query plug-ins
    - MySQL plug-in
      - Added support for new queries
    - MongoDB plug-in
      - Added support for new queries
    - Redis plug-in
      - Added support for new queries
    - PostgreSQL plug-in
      - Added support for new queries

## v1.3.7 (2025-06-XX)
### Improved
- Update math and aggregation functions
  - PostgreSQL plug-in
  - MySQL plug-in

## v1.3.6 (2025-05-31)
### New features
- Enhanced `sql.Service.Insert()` to support inserting multiple objects at once.
- Added `sql.ResultSetOption` type for flexible result set configuration.
### Improved
- Updated FoundationDB from 7.3.57 to 7.3.67.
- Updated dependencies (`go-sqlparser`, `go-sqltest`, etc.)
- Updated test functions (`TestPostgreSQLServer`, `TestPostgreSQLTestSuite`, `TestMySQLTestSuite`) for compatibility with the latest `go-sqltest`.
- Various minor bug fixes and refactoring.
### Fixed
- Fixed warnings related to non-constant format strings.

## v1.3.5 (2025-01-01)
- Improvements
  - Updated authenticator interface
- Updated
  - PostgreSQL plug-in
  - MySQL plug-in
  - MongoDB plug-in
  - Redis plug-in

## v1.3.4 (2024-12-11)
- New Features
  - Query Plug-ins
    - MySQL Plug-in
      - Supported `CREATE INDEX` and `DROP INDEX` by converting them to `ALTER TABLE`
      - Added TLS support
    - PostgreSQL Plug-in
      - Supported `CREATE INDEX` and `DROP INDEX` by converting them to `ALTER TABLE`
- Improvements
  - Query Plug-ins
    - MySQL Plug-in
      - Standardized the PostgreSQL plugin and implementation during the upgrade to go-mysql v1.1
  - Store Plug-ins
    - Document Store Plug-ins
      - Updated the format and implementation of secondary indexes
      - KV Store Plug-ins
        - Enhanced the `memdb` plug-in to use a custom indexer
    - Coordinator Plug-ins
      - Enhanced the `memdb` plug-in to use a custom indexer

## v1.3.3 (2024-08-22)
- Improvements
  - Query Plug-ins
  　- TLS support settings

## v1.3.2 (2024-05-22)
- New features
  - TLS Support
    - Supported MongoDB plug-in

## v1.3.1 (2023-05-18)
- New features
  - TLS Support
    - Supported PostgreSQL plug-in
    - Supported Redis plug-in

## v1.3.0 (2023-12-30)
- New features
  - Authenticator plug-ins
    - Added password authenticator interface
      - Supported PostgreSQL plug-in
      - Supported Redis plug-in
- Improvements
  - Query plug-ins
    - Redis plug-in
      - Supported queries
        - HASH commands

## v1.2.0 (2023-11-15)
- New features
  - Distributed plug-ins
    - Storage plug-ins
      - Enabled cache store plug-in (ristretto) as default
- Updates
  - Redis plug-in
    - Support new commands
      - DEL and EXISTS
- Improvements
  - Updated to set service metrics to prometheus
    - Query plug-ins
      - PostgreSQL, MySQL, Redis and Mongo
    - Storage plug-ins
      - Cache store (ristretto) 

## v1.1.1 (2023-11-02)
- Fixed
  - PostgreSQL plug-in
    - Fixed transaction hangup using copy commands
    - Fixed to run pgbench on Ubuntu platforms
      - Upgraded go-postgresql to update the protocol message reader
- Changed
  - Docker image
    - Change to FoundationDB plug-in as default store
## v1.1.0 (2023-10-20)
- New features
  - Query plug-ins
    - MySQL plug-in
      - Support transaction control statements
        - BEGIN, COMMIT and ROLLBACK
    - PostgreSQL plug-in
      - Support transaction control statements
        - BEGIN, COMMIT and ROLLBACK

## v1.0.3 (2023-09-30)
- New features
  - Enable pprof
- Improvements
  - PostgreSQL plug-in
    - Supported new data types
      - TIMESTAMP
    - Supported new statements
      - TRUNCATE, VACCUM and COPY
    - Supported pgbench workload
  - MySQL plug-in
    - Supported new data types
      - DATETIME and TIMESTAMP

## v1.0.2 (2023-09-12)
- Improvements
  - PostgreSQL plug-in
    - SELECT
      - Supported basic aggregate functions
        - COUNT, SUM, AVG, MIN and MAX
      - Supported basic math functions
        - ABS, CEIL and FLOOR

## v1.0.1 (2023-09-06)
- Improvements
  - PostgreSQL plug-in
    - Improved schema validation for INSERT, SELECT, UPDATE, and DELETE queries
    - Enabled PICT based scenario tests of go-sqltest
  - MySQL plug-in
    - Improved schema validation for INSERT, SELECT, UPDATE, and DELETE queries
    - Enabled PICT based scenario tests of go-sqltest

## v1.0.0 (2023-08-30)
- Initial public technology preview release
- New features
  - Query plug-ins
    - Add PostgreSQL plug-in
  - Distributed plug-ins
    - Storage plug-ins
      - Add FoundationDB plug-in
      - Add cache store plug-in
    - Coordinator plug-ins
      - Add FoundationDB plug-in
      - go-memdb plug-in (standalone)

## v0.9.0 (2023-05-07)
- New features
  - CLI Utilities
    - Added [puzzledb-cli](doc/cmd/cli/puzzledb-cli.md)
  - Operator APIs
    - Added gRPC services for operator APIs and CLI utilities
    - Added Prometheus metrics expoter
  - Coordinator plug-ins
    - Added memdb plug-in (standalone)
  - Distributed tracer plug-ins
    - Added OpenTelemetry plug-in
- Improvements
  - Configuration
    - Added support for configuration with environment variables
    - Added support for configuration with puzzledb.yaml
  - Query plug-ins
    - MySQL plug-in
      - Added support for new queries
        - DROP DATABASE
        - DROP TABLE
- Bug Fixes
  - Coder plug-ins
    - Key coder plug-ins
      - Tuple plug-in
        - Fix encoder not to panic on Ubuntu 20.04

## v0.8.0 (2023-04-10)
- Initial release
- New Features
  - Coderr plug-ins
    - Key coder plug-ins
      - Tuple plug-in
    - Document coder plug-ins
      - CBOR plug-in
  - Store plug-ins
    - go-memdb plug-in (standalone)
  - Query plug-ins
    - MySQL plug-in
    - MongoDB plug-in
    - Redis plug-in
