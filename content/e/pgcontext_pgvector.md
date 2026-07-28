---
title: "pgcontext_pgvector"
linkTitle: "pgcontext_pgvector"
description: "Optional pgvector compatibility bridge for pgcontext HNSW indexes."
weight: 1970
categories: ["RAG"]
width: full
---

[**pgcontext**](https://pgxn.org/dist/pgContext/0.2.0/) : Optional pgvector compatibility bridge for pgcontext HNSW indexes.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1970** | {{< badge content="pgcontext_pgvector" link="https://pgxn.org/dist/pgContext/0.2.0/" >}} | {{< ext "pgcontext_pgvector" "pgcontext" >}} | `0.2.0` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcontext" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "pgcontext" >}} |

> [!Note] Optional control shipped by pgcontext 0.2.0; requires pgcontext and vector.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pgcontext` | `pgcontext`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pgcontext_18" "green" >}} {{< bg "17" "pgcontext_17" "green" >}} {{< bg "16" "pgcontext_16" "red" >}} {{< bg "15" "pgcontext_15" "red" >}} {{< bg "14" "pgcontext_14" "red" >}} | `pgcontext_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pgcontext" "green" >}} {{< bg "17" "postgresql-17-pgcontext" "green" >}} {{< bg "16" "postgresql-16-pgcontext" "red" >}} {{< bg "15" "postgresql-15-pgcontext" "red" >}} {{< bg "14" "postgresql-14-pgcontext" "red" >}} | `postgresql-$v-pgcontext` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgcontext_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgcontext_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgcontext_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgcontext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgcontext : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgcontext : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgcontext : N/A 0" "gray" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/pgContext/0.2.0/" title="Repository" icon="link" subtitle="pgxn.org/dist/pgContext/0.2.0/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcontext-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgcontext;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgcontext;		# install via package name, for the active PG version
pig install pgcontext_pgvector;		# install by extension name, for the current active PG version

pig install pgcontext_pgvector -v 18;   # install for PG 18
pig install pgcontext_pgvector -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcontext_pgvector CASCADE; -- requires pgcontext, vector
```

## Usage

Sources:

- [pgContext 0.2.0 pgvector coexistence guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_coexist.md)
- [pgContext 0.2.0 pgvector migration guide](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_migration.md)
- [pgcontext_pgvector control file](https://github.com/evokoa/pgcontext/blob/v0.2.0/pgcontext_pgvector.control)
- [pgcontext_pgvector extension SQL](https://github.com/evokoa/pgcontext/blob/v0.2.0/sql/pgcontext_pgvector--0.2.0.sql)
- [pgContext 0.2.0 release notes](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/release_notes.md)

`pgcontext_pgvector` is the optional pgContext companion bridge for serving pgContext HNSW indexes over columns owned by the pgvector extension. It does not merge the two type systems or copy application data; it adds certified casts, support functions, and operator classes while exact distance semantics remain bound to pgvector operators.

### Certified Profile and Installation

Version 0.2.0 fails closed unless the database uses PostgreSQL 17, pgContext 0.2.0, and pgvector 0.8.x installed in `public`. Install the prerequisites and bridge explicitly:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgcontext;
CREATE EXTENSION pgcontext_pgvector;
```

The reverse order of the two prerequisite extensions is also valid, but `pgcontext_pgvector` must come after both. Installation requires superuser privileges.

### Index an Existing pgvector Column

```sql
CREATE INDEX items_embedding_pgc
    ON items USING pgcontext_hnsw
       (embedding pgcontext.vector_hnsw_pgvector_cosine_ops);

SELECT id
FROM items
ORDER BY embedding <=> $1::public.vector
LIMIT 10;
```

Existing pgvector-spelled SQL can use the pgContext access method. ANN candidates are resolved to live heap rows and reranked with the pgvector operator, preserving its `double precision` distance result semantics.

### Important Objects

- `pgcontext.vector_hnsw_pgvector_l2_ops`, `pgcontext.vector_hnsw_pgvector_ip_ops`, `pgcontext.vector_hnsw_pgvector_cosine_ops`, and `pgcontext.vector_hnsw_pgvector_l1_ops` serve existing `public.vector` columns.
- `pgcontext.sparsevec_hnsw_pgvector_cosine_ops` serves certified `public.sparsevec` columns, subject to the documented 16,000-dimension and page-envelope limits.
- `pgcontext.migration_report()` inventories pgvector columns, dependencies, HNSW, and IVFFlat without requiring the bridge.
- Ownership-conversion functions provide reviewed fast or restricted-online workflows; IVFFlat is rebuilt as HNSW rather than converted in place.

### Dependency and Removal Boundaries

The main `pgcontext` extension remains independent of pgvector. Bridge indexes depend on `pgcontext_pgvector`, and the bridge depends on both parent extensions, so PostgreSQL blocks removal under `RESTRICT` until those indexes are removed or converted.

Do not use `DROP EXTENSION vector CASCADE` as a migration method. Inventory arrays, views, functions, prepared sessions, expression indexes, and other application dependencies first. The bridge does not provide every pgvector helper, IVFFlat, iterative-scan GUC, parallel-build, subvector, or progress-reporting behavior.

No preload or restart is required. The bridge is a privileged compatibility surface, not a promise that all future pgContext, pgvector, PostgreSQL-major, or on-disk index combinations are compatible; rerun the certified preflight and rebuild validation when any component changes.
