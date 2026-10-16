---
title: "pg_vectorsearch"
linkTitle: "pg_vectorsearch"
description: "Approximate vector search with PRISM indexes and RaBitQ quantization"
weight: 1990
categories: ["RAG"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_vectorsearch**](https://github.com/timescale/pg_vectorsearch) : Approximate vector search with PRISM indexes and RaBitQ quantization


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1990** | {{< badge content="pg_vectorsearch" link="https://github.com/timescale/pg_vectorsearch" >}} | {{< ext "pg_vectorsearch" >}} | `0.1.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "vectorscale" >}} |

> [!Note] Alpha 0.1.0 for testing; packaged for PostgreSQL 18 only. Requires OpenBLAS; no pgvector dependency. No in-place upgrade from earlier release candidates.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_vectorsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_vectorsearch_18" "green" >}} {{< bg "17" "pg_vectorsearch_17" "red" >}} {{< bg "16" "pg_vectorsearch_16" "red" >}} {{< bg "15" "pg_vectorsearch_15" "red" >}} {{< bg "14" "pg_vectorsearch_14" "red" >}} | `pg_vectorsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-vectorsearch" "green" >}} {{< bg "17" "postgresql-17-vectorsearch" "red" >}} {{< bg "16" "postgresql-16-vectorsearch" "red" >}} {{< bg "15" "postgresql-15-vectorsearch" "red" >}} {{< bg "14" "postgresql-14-vectorsearch" "red" >}} | `postgresql-$v-vectorsearch` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "pg_vectorsearch_18 : MISS 0" "red" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-18-vectorsearch : MISS 0" "red" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-vectorsearch : N/A 0" "gray" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_vectorsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_vectorsearch" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vectorsearch-0.1.0.tar.gz meson-0.63.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_vectorsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_vectorsearch;		# install via package name, for the active PG version

pig install pg_vectorsearch -v 18;   # install for PG 18

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_vectorsearch;
```

## Usage

Sources:

- [README.md](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/README.md)
- [sql/pg_vectorsearch.sql](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/sql/pg_vectorsearch.sql)
- [docs/maintenance.md](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/docs/maintenance.md)
- [docs/tuning.md](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/docs/tuning.md)
- [CHANGELOG.md](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/CHANGELOG.md)
- [VERSION](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/VERSION)
- [meson.build](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/meson.build)
- [pg_vectorsearch.control.in](https://github.com/timescale/pg_vectorsearch/blob/36b75d1cb56c12a380d600ef27217d88c9f419ec/pg_vectorsearch.control.in)

`pg_vectorsearch` 0.1.0 is an alpha extension for approximate nearest-neighbor search with the PRISM index. It supports PostgreSQL 18 and 19, native `vec32` / `vec16` types and optional pgvector interoperability. It is intended for testing, with evolving behavior and maintenance requirements.

### Core Workflow

```sql
CREATE EXTENSION pg_vectorsearch;
CREATE TABLE vectorsearch_items (
    id serial PRIMARY KEY,
    embedding vec32(3) STORAGE PLAIN
);
INSERT INTO vectorsearch_items (embedding)
SELECT ARRAY[i::real, (i % 7)::real, (i % 5)::real]::vec32
FROM generate_series(1, 1000) AS i;
CREATE INDEX vectorsearch_items_prism ON vectorsearch_items
USING prism (embedding vec32_l2_ops);
SET prism.nprobe = 40;
SELECT id FROM vectorsearch_items
ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

### Operational Boundaries

Install as a superuser; no preload is required and the extension is not relocatable. The `prism` access method has `vec32_l2_ops`, `vec32_ip_ops`, `vec32_cosine_ops` and corresponding vec16 operator classes; `<->`, `<#>` and `<=>` express L2, negative inner product and cosine distance. `prism.nprobe` trades recall for query work; zero selects automatic tuning, not exact recall.

Load representative data before building the index. Indexed dimensions are limited to 1968; expression and unlogged indexes are unsupported. Keep `prism.leaf_refine_threshold` at zero. Filtering can return fewer rows than LIMIT, and candidate buffers bounded by `work_mem` can truncate large or unbounded results. Measure against an exact baseline.

Inspect `prism_index_settings` and follow the maintenance guide for churn. `prism_rebalance` supports only flat RaBitQ centroid-page indexes with `centroid_fastscan` off, not the default multilevel format. DELETE leaves tombstones, and retired pages are reclaimed only by REINDEX. There is no in-place upgrade from the earlier release candidates.
