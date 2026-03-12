---
title: "test_decoding"
linkTitle: "test_decoding"
description: "SQL-based test/example module for WAL logical decoding"
weight: 9970
categories: ["ETL"]
width: full
---

[**test_decoding**](https://www.postgresql.org/docs/current/test-decoding.html) : SQL-based test/example module for WAL logical decoding


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9970** | {{< badge content="test_decoding" link="https://www.postgresql.org/docs/current/test-decoding.html" >}} | {{< ext "test_decoding" >}} | `-` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "pgoutput" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "kafka_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


This extension does not need `CREATE EXTENSION` DDL command





## Usage

> [test_decoding: SQL-based test/example module for WAL logical decoding](https://www.postgresql.org/docs/current/test-decoding.html)

A built-in PostgreSQL logical decoding output plugin that produces text representations of WAL changes. Primarily used for testing and as a reference implementation.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('test_slot', 'test_decoding');

-- Perform some changes
CREATE TABLE test_table (id serial PRIMARY KEY, data text);
INSERT INTO test_table (data) VALUES ('hello');
UPDATE test_table SET data = 'world' WHERE id = 1;
DELETE FROM test_table WHERE id = 1;

-- Peek at changes (without consuming)
SELECT * FROM pg_logical_slot_peek_changes('test_slot', NULL, NULL);

-- Get and consume changes
SELECT * FROM pg_logical_slot_get_changes('test_slot', NULL, NULL);
```

### Output Format

```
BEGIN 1234
table public.test_table: INSERT: id[integer]:1 data[text]:'hello'
table public.test_table: UPDATE: id[integer]:1 data[text]:'world'
table public.test_table: DELETE: id[integer]:1
COMMIT 1234
```

### Using with pg_recvlogical

```bash
# Create slot
pg_recvlogical -d postgres --slot test_slot --create-slot -P test_decoding

# Stream changes
pg_recvlogical -d postgres --slot test_slot --start -f -

# Drop slot
pg_recvlogical -d postgres --slot test_slot --drop-slot
```

### Options

Pass options as key-value pairs:

```sql
SELECT * FROM pg_logical_slot_get_changes('test_slot', NULL, NULL,
    'include-xids', '1',
    'skip-empty-xacts', '1',
    'include-timestamp', '1'
);
```

- `include-xids` - include transaction IDs in output
- `skip-empty-xacts` - skip transactions with no changes
- `include-timestamp` - include commit timestamps

### Notes

- Ships with PostgreSQL (contrib module, no separate installation needed)
- Intended for testing and debugging logical decoding
- For production CDC, use purpose-built plugins (wal2json, pgoutput, decoderbufs)
