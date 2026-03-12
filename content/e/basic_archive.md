---
title: "basic_archive"
linkTitle: "basic_archive"
description: "an example of an archive module"
weight: 5940
categories: ["ADMIN"]
width: full
---

[**basic_archive**](https://www.postgresql.org/docs/current/basic-archive.html) : an example of an archive module


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5940** | {{< badge content="basic_archive" link="https://www.postgresql.org/docs/current/basic-archive.html" >}} | {{< ext "basic_archive" >}} | `-` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "basebackup_to_shell" >}} {{< ext "pg_walinspect" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [basic_archive: an example of an archive module](https://www.postgresql.org/docs/current/basic-archive.html)

The `basic_archive` module is a WAL archive module that copies completed WAL segment files to a specified directory. It serves as a reference implementation for custom archive modules.

### Configuration

Add to `postgresql.conf`:

```ini
archive_mode = 'on'
archive_library = 'basic_archive'
basic_archive.archive_directory = '/path/to/archive/directory'
```

### Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `basic_archive.archive_directory` | string | Directory to copy WAL files to (must already exist) |

If `archive_mode` is enabled but `basic_archive.archive_directory` is empty (default), the server will accumulate WAL files until a directory path is configured.

### Notes

- The target directory must be created before use; the module will not create it
- After a server crash, temporary files with the `archtemp` prefix may be left in the archive directory and should be deleted before restarting
- These temporary files can also be safely removed while the server is running, provided they are not related to an ongoing archive operation
- This module is primarily intended as a simple example and starting point for developing custom archive modules
