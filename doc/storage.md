![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go) [![Go](https://github.com/cybergarage/puzzledb-go/puzzledb-go/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-go/puzzledb/actions/workflows/make.yml)
 [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-go.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-go)
![](img/logo.png)

# Concept

PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

![](img/storage.png)

PuzzleDB defines a plug-in interface to the Key-Value store, which allows importing small local in-memory databases like memdb or large in-memory databases like FoundationDB or TiKV.
