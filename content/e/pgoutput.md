---
title: "pgoutput"
linkTitle: "pgoutput"
description: "Logical Replication output plugin"
weight: 9980
categories: ["ETL"]
width: full
---

[**pgoutput**](https://www.postgresql.org/docs/current/protocol-logical-replication.html) : Logical Replication output plugin


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9980** | {{< badge content="pgoutput" link="https://www.postgresql.org/docs/current/protocol-logical-replication.html" >}} | {{< ext "pgoutput" >}} | `-` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "test_decoding" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "kafka_fdw" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


This extension does not need `CREATE EXTENSION` DDL command





## Usage

> [pgoutput: Logical Replication output plugin](https://www.postgresql.org/docs/current/protocol-logical-replication.html)

The built-in logical decoding output plugin used by PostgreSQL's native logical replication system. It is the default plugin for `CREATE PUBLICATION` / `CREATE SUBSCRIPTION`.

### Overview

`pgoutput` is PostgreSQL's native logical replication output plugin. Unlike third-party plugins, it does not need to be installed separately -- it is built into PostgreSQL core (version 10+).

### Using with Native Logical Replication

On the **publisher**:

```sql
-- Create a publication for specific tables
CREATE PUBLICATION my_pub FOR TABLE orders, customers;

-- Or publish all tables
CREATE PUBLICATION my_pub FOR ALL TABLES;

-- Publish with row filtering (PG 15+)
CREATE PUBLICATION filtered_pub FOR TABLE orders WHERE (status = 'active');

-- Publish with column filtering (PG 15+)
CREATE PUBLICATION col_pub FOR TABLE users (id, name, email);
```

On the **subscriber**:

```sql
-- Create a subscription
CREATE SUBSCRIPTION my_sub
    CONNECTION 'host=publisher dbname=mydb user=replicator'
    PUBLICATION my_pub;

-- Check subscription status
SELECT * FROM pg_stat_subscription;
```

### Using pgoutput Directly with Replication Slots

```sql
-- Create a logical replication slot using pgoutput
SELECT * FROM pg_create_logical_replication_slot('my_slot', 'pgoutput');

-- Consume changes (requires protocol-level parameters)
-- pgoutput is designed for the streaming replication protocol,
-- not the SQL slot functions (use test_decoding for SQL-based testing)
```

### Publication Management

```sql
-- Add tables to an existing publication
ALTER PUBLICATION my_pub ADD TABLE new_table;

-- Remove tables
ALTER PUBLICATION my_pub DROP TABLE old_table;

-- View publications
SELECT * FROM pg_publication;
SELECT * FROM pg_publication_tables;
```

### Subscription Management

```sql
-- Disable subscription
ALTER SUBSCRIPTION my_sub DISABLE;

-- Refresh subscription (pick up new tables)
ALTER SUBSCRIPTION my_sub REFRESH PUBLICATION;

-- Drop subscription
DROP SUBSCRIPTION my_sub;
```

### Key Features

- Built into PostgreSQL core (10+)
- Binary format for efficient data transfer
- Row and column filtering (PostgreSQL 15+)
- Supports initial table synchronization
- Handles schema changes via subscription refresh
- Supports multiple publications per subscription
