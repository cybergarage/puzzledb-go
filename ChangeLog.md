# Changelog

## v1.x.x (2023-xx-xx)
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

## v1.1.0 (2023-xx-xx)
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
    - PostgreSQL plug-in
      - Added support for new queries
    - MongoDB plug-in
      - Added support for new queries
    - Redis plug-in
      - Added support for new queries
    - Storage plug-ins
      - Update cache store plug-in

## v1.0.3 (2023-xx-xx)
- Improvement
  - PostgreSQL plug-in
    - Support DATE and TIMESTAMP data types
    - Support pgbench workload
  - MySQL plug-in
    - Support DATE and TIMESTAMP data types

## v1.0.2 (2023-09-12)
- Improvement
  - PostgreSQL plug-in
    - SELECT
      - Supported basic aggregate functions
        - COUNT, SUM, AVG, MIN and MAX
      - Supported basic math functions
        - ABS, CEIL and FLOOR

## v1.0.1 (2023-09-06)
- Improvement
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
