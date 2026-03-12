---
title: "file_fdw"
linkTitle: "file_fdw"
description: "foreign-data wrapper for flat file access"
weight: 8980
categories: ["FDW"]
width: full
---

[**file_fdw**](https://www.postgresql.org/docs/current/file-fdw.html) : foreign-data wrapper for flat file access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8980** | {{< badge content="file_fdw" link="https://www.postgresql.org/docs/current/file-fdw.html" >}} | {{< ext "file_fdw" >}} | `1.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_sqlog" >}} |
|   **See Also**    | {{< ext "log_fdw" >}} {{< ext "wrappers" >}} {{< ext "sqlite_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_bulkload" >}} {{< ext "multicorn" >}} {{< ext "hdfs_fdw" >}} {{< ext "postgres_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION file_fdw;
```




## Usage

> [file_fdw: Foreign data wrapper for flat file access](https://www.postgresql.org/docs/current/file-fdw.html)

### Create Server

```sql
CREATE EXTENSION file_fdw;

CREATE SERVER file_server FOREIGN DATA WRAPPER file_fdw;
```

### Read a CSV File

```sql
CREATE FOREIGN TABLE csv_data (
  id integer,
  name text,
  value numeric
)
SERVER file_server
OPTIONS (filename '/path/to/data.csv', format 'csv', header 'true');

SELECT * FROM csv_data;
```

### Read PostgreSQL CSV Logs

```sql
CREATE FOREIGN TABLE pglog (
  log_time timestamp(3) with time zone,
  user_name text,
  database_name text,
  process_id integer,
  connection_from text,
  session_id text,
  session_line_num bigint,
  command_tag text,
  session_start_time timestamp with time zone,
  virtual_transaction_id text,
  transaction_id bigint,
  error_severity text,
  sql_state_code text,
  message text,
  detail text,
  hint text,
  internal_query text,
  internal_query_pos integer,
  context text,
  query text,
  query_pos integer,
  location text,
  application_name text,
  backend_type text,
  leader_pid integer,
  query_id bigint
)
SERVER file_server
OPTIONS (filename 'log/pglog.csv', format 'csv');
```

### Read Program Output

```sql
CREATE FOREIGN TABLE process_list (
  pid text,
  command text
)
SERVER file_server
OPTIONS (program 'ps aux | tail -n +2', format 'text', delimiter ' ');
```

### Table Options

| Option | Description |
|--------|-------------|
| `filename` | File path (relative to data directory). Required unless `program` is used |
| `program` | Shell command whose stdout is read. Required unless `filename` is used |
| `format` | Data format: `csv`, `text`, or `binary` (same as COPY) |
| `header` | `true` if file has a header row |
| `delimiter` | Column delimiter character |
| `quote` | Quote character |
| `escape` | Escape character |
| `null` | String representing NULL values |
| `encoding` | Data encoding |
| `on_error` | Error handling during type conversion |
| `reject_limit` | Maximum tolerated errors |

### Column Options

| Option | Description |
|--------|-------------|
| `force_not_null` | Do not match column values against the null string |
| `force_null` | Match quoted values against the null string and return NULL |

```sql
CREATE FOREIGN TABLE films (
  code char(5) NOT NULL,
  title text NOT NULL,
  rating text OPTIONS (force_null 'true')
)
SERVER file_server
OPTIONS (filename '/data/films.csv', format 'csv');
```

file_fdw is read-only. Changing table-level options requires superuser privileges or the `pg_read_server_files` / `pg_execute_server_program` role.
