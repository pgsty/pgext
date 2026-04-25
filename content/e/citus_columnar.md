---
title: "citus_columnar"
linkTitle: "citus_columnar"
description: "Citus columnar storage engine"
weight: 2401
categories: ["OLAP"]
width: full
---

[**citus**](https://github.com/citusdata/citus) : Citus columnar storage engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2401** | {{< badge content="citus_columnar" link="https://github.com/citusdata/citus" >}} | {{< ext "citus_columnar" "citus" >}} | `14.0.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_parquet" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "orioledb" >}} |
|    **Siblings**   | {{< ext "citus" >}} |

> [!Note] conflict with hydra columnar, no pg18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} {{< bg "14" "citus_14" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} {{< bg "14" "postgresql-14-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 29" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 21" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 26" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install citus;		# install via package name, for the active PG version
pig install citus_columnar;		# install by extension name, for the current active PG version

pig install citus_columnar -v 18;   # install for PG 18
pig install citus_columnar -v 17;   # install for PG 17
pig install citus_columnar -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citus_columnar;
```



## Usage

> [citus_columnar: Columnar storage access method for PostgreSQL](https://github.com/citusdata/citus)

Citus Columnar provides a columnar storage engine for PostgreSQL. It stores data in a column-oriented format with automatic compression, making it ideal for analytical workloads on append-only data where queries typically read a subset of columns.

**Key Documentation:**

- [Columnar Storage](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)
- [Columnar Compression](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#compression)

### Creating Columnar Tables

Use the `USING columnar` clause when creating a table:

```sql
CREATE TABLE events (
    event_id    BIGINT,
    event_time  TIMESTAMPTZ,
    event_type  TEXT,
    user_id     INT,
    payload     JSONB
) USING columnar;
```

### Compression Options

Configure compression per table. Supported methods: `zstd` (default), `lz4`, `pglz`, `none`.

```sql
ALTER TABLE events SET (
    columnar.compression = zstd,
    columnar.compression_level = 3
);
```

### Chunk Group and Stripe Settings

Columnar stores data in stripes, each containing chunk groups. Tuning these affects both compression ratio and query performance.

```sql
ALTER TABLE events SET (
    columnar.stripe_row_limit = 150000,    -- max rows per stripe (default 150000)
    columnar.chunk_group_row_limit = 10000 -- rows per chunk group (default 10000)
);
```

### When to Use Columnar

Columnar storage works best for:

- **Analytics and reporting** on wide tables where queries read few columns
- **Append-only workloads** (e.g., logs, events, time-series archives)
- **Large fact tables** scanned in bulk with aggregations
- **Cold data archival** where high compression is valuable

### Limitations

- **No UPDATE or DELETE**: columnar tables are append-only
- **No indexes**: sequential/columnar scans only
- **No TOAST**: large values stored inline
- **No logical replication** as a publisher
- **No tid scans**

### Column Projection and Chunk Group Skipping

Columnar automatically reads only the columns referenced in a query (column projection) and skips chunk groups whose min/max metadata does not match the query predicates:

```sql
-- Only reads event_type and event_time columns; skips irrelevant chunks
SELECT event_type, count(*)
FROM events
WHERE event_time > '2025-01-01'
GROUP BY event_type;
```

### Monitoring Columnar Storage

Inspect stripe and chunk metadata:

```sql
-- View stripes for a columnar table
SELECT * FROM columnar.stripe WHERE relation = 'events'::regclass;

-- View chunk group details
SELECT * FROM columnar.chunk_group WHERE relation = 'events'::regclass;

-- Check storage size and compression ratio
SELECT pg_size_pretty(pg_total_relation_size('events')) AS total_size;
```

### Converting Between Heap and Columnar

Convert an existing heap table to columnar:

```sql
-- Create a columnar copy
CREATE TABLE events_columnar (LIKE events) USING columnar;
INSERT INTO events_columnar SELECT * FROM events;

-- Or use ALTER TABLE (Citus 11+)
SELECT alter_table_set_access_method('events', 'columnar');
```

Convert columnar back to heap:

```sql
SELECT alter_table_set_access_method('events', 'heap');
```

### Using with Partitioning

Combine columnar with partitioning to keep recent data in heap (for updates/indexes) and archive older partitions as columnar:

```sql
CREATE TABLE events (
    event_time TIMESTAMPTZ,
    data       JSONB
) PARTITION BY RANGE (event_time);

-- Recent data: heap (supports indexes and updates)
CREATE TABLE events_current PARTITION OF events
    FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');

-- Archived data: columnar (compressed, read-optimized)
CREATE TABLE events_archive PARTITION OF events
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01')
    USING columnar;
```
