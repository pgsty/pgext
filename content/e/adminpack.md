---
title: "adminpack"
linkTitle: "adminpack"
description: "administrative functions for PostgreSQL"
weight: 5970
categories: ["ADMIN"]
width: full
---

[**adminpack**](https://www.postgresql.org/docs/16/adminpack.html) : administrative functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5970** | {{< badge content="adminpack" link="https://www.postgresql.org/docs/16/adminpack.html" >}} | {{< ext "adminpack" >}} | `2.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "fio" >}} {{< ext "lo" >}} {{< ext "file_fdw" >}} {{< ext "ddlx" >}} {{< ext "pgdd" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pg_repack" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "N/A" "PostgreSQL 18: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 17: not available" "red" >}} | {{< bg "2.1" "PostgreSQL 16: version 2.1" "green" >}} | {{< bg "2.1" "PostgreSQL 15: version 2.1" "green" >}} | {{< bg "2.1" "PostgreSQL 14: version 2.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION adminpack;
```




## Usage

> [adminpack: administrative functions for PostgreSQL](https://www.postgresql.org/docs/16/adminpack.html)

The `adminpack` extension provides server-side file management and log access functions, primarily used by pgAdmin and other administration tools. All functions are restricted to superusers by default.

### File Operations

```sql
-- Write text to a file (append=false: file must not exist; append=true: append)
SELECT pg_file_write('/tmp/test.txt', 'Hello World', false);   -- returns bytes written

-- Append to existing file
SELECT pg_file_write('/tmp/test.txt', E'\nMore data', true);

-- Sync file to disk
SELECT pg_file_sync('/tmp/test.txt');

-- Rename a file
SELECT pg_file_rename('/tmp/old.txt', '/tmp/new.txt');

-- Rename with archiving (moves existing newname to archive first)
SELECT pg_file_rename('/tmp/old.txt', '/tmp/new.txt', '/tmp/archive.txt');

-- Delete a file
SELECT pg_file_unlink('/tmp/test.txt');   -- returns true on success
```

### Log File Access

```sql
-- List server log files (requires default log_filename format)
SELECT * FROM pg_logdir_ls();
```

Returns timestamps and paths of log files from the `log_directory`.

### Function Reference

| Function | Returns | Description |
|----------|---------|-------------|
| `pg_file_write(filename, data, append)` | bigint | Write text to file; returns bytes written |
| `pg_file_sync(filename)` | void | Flush file or directory to disk |
| `pg_file_rename(old, new [, archive])` | boolean | Rename file, optionally archiving existing target |
| `pg_file_unlink(filename)` | boolean | Delete a file |
| `pg_logdir_ls()` | setof record | List log files with timestamps |

### Access Control

- All functions default to superuser-only access
- Permissions can be granted via `GRANT` to non-superusers
- File access is restricted to the database cluster directory unless the user has `pg_read_server_files` or `pg_write_server_files` roles

Note: `adminpack` was removed in PostgreSQL 17. For PostgreSQL 16 and earlier only.
