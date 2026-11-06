---
title: "pg_lion_vector"
linkTitle: "pg_lion_vector"
description: "Filtered nearest-neighbor indexes for pgvector vector and halfvec"
weight: 2717
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_lion**](https://github.com/cgkanchi/pg_lion) : Filtered nearest-neighbor indexes for pgvector vector and halfvec


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2717** | {{< badge content="pg_lion_vector" link="https://github.com/cgkanchi/pg_lion" >}} | {{< ext "pg_lion_vector" "pg_lion" >}} | `0.3.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_lion" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "pg_lion" >}} {{< ext "pg_lion_citext" >}} |

> [!Note] pgvector >= 0.7.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lion` | `pg_lion`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_lion_18" "green" >}} {{< bg "17" "pg_lion_17" "green" >}} {{< bg "16" "pg_lion_16" "green" >}} {{< bg "15" "pg_lion_15" "red" >}} {{< bg "14" "pg_lion_14" "red" >}} | `pg_lion_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-lion" "green" >}} {{< bg "17" "postgresql-17-pg-lion" "green" >}} {{< bg "16" "postgresql-16-pg-lion" "green" >}} {{< bg "15" "postgresql-15-pg-lion" "red" >}} {{< bg "14" "postgresql-14-pg-lion" "red" >}} | `postgresql-$v-pg-lion` | `postgresql-$v-pgvector` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
{.matrix .wide .expand}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cgkanchi/pg_lion" title="Repository" icon="github" subtitle="github.com/cgkanchi/pg_lion" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lion-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_lion;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_lion;		# install via package name, for the active PG version
pig install pg_lion_vector;		# install by extension name, for the current active PG version

pig install pg_lion_vector -v 18;   # install for PG 18
pig install pg_lion_vector -v 17;   # install for PG 17
pig install pg_lion_vector -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_lion_vector CASCADE; -- requires pg_lion, vector
```

## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector.control)
- [Extension SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector--0.3.0.sql)

`pg_lion_vector` 0.3.0 adds the `lionvec` nearest-neighbor access method for pgvector `vector` and `halfvec` values. It uses the parent `pg_lion` library and requires pgvector 0.7 or later.

### Core Workflow

Create the parent and `vector` extensions first. A separate `lion` index can filter candidates before distance ranking.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_vector;
CREATE TABLE lion_items (id integer, category integer, embedding vector(3));
INSERT INTO lion_items VALUES (1, 3, '[1,0,0]'), (2, 3, '[0,1,0]'), (3, 4, '[0,0,1]');
CREATE INDEX lion_embedding ON lion_items USING lionvec (embedding vector_cosine_ops);
CREATE INDEX lion_category ON lion_items USING lion (category);
ALTER FUNCTION cosine_distance(vector, vector) COST 100;
VACUUM ANALYZE lion_items;
SELECT id FROM lion_items
WHERE category = 3
ORDER BY embedding <=> '[1,0,0]'::vector
LIMIT 10;
```

Upstream recommends `COST 100` for the distance function used by the chosen operator class so the planner can compare scan and sort costs. This setting is not preserved by `pg_dump` and may be reset by an update of `vector`.

### Operator Classes and Tuning

| Type | L2 (`<->`) | Negative inner product (`<#>`) | Cosine (`<=>`) |
| --- | --- | --- | --- |
| `vector` | `vector_l2_ops` (default) | `vector_ip_ops` | `vector_cosine_ops` |
| `halfvec` | `halfvec_l2_ops` (default) | `halfvec_ip_ops` | `halfvec_cosine_ops` |

`partitions` controls the index build's clustering; `maintenance_work_mem` must hold an adequate sample. `pg_lion.vector_margin` and `pg_lion.vector_patience` control approximate search. Small candidate sets under a supported lion filter can use exact scoring. `pg_lion.vector_postfilter` controls predicates the lion indexes cannot answer; review the reference before changing it.

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Only one `vector` or `halfvec` column, up to 8128 dimensions, is supported; this is generally approximate search. A filter does not turn every query into exact nearest-neighbor search. Heavy churn may require occasional `REINDEX`. Parent custom-WAL indexes retain their preload requirement. Keep parent and companion files at the same version.
