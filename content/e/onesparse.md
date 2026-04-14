---
title: "onesparse"
linkTitle: "onesparse"
description: "Sparse linear algebra and graph extension for PostgreSQL 18"
weight: 2980
categories: ["FEAT"]
width: full
---

[**one_sparse**](https://github.com/OneSparse/OneSparse) : Sparse linear algebra and graph extension for PostgreSQL 18


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2980** | {{< badge content="onesparse" link="https://github.com/OneSparse/OneSparse" >}} | {{< ext "onesparse" "one_sparse" >}} | `1.0.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `onesparse` |
|   **See Also**    | {{< ext "age" >}} {{< ext "pgrouting" >}} {{< ext "postgis" >}} |

> [!Note] PG18 only; upstream release v1.0.0 ships extension SQL version 0.1.0


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `one_sparse` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "onesparse_18" "green" >}} {{< bg "17" "onesparse_17" "red" >}} {{< bg "16" "onesparse_16" "red" >}} {{< bg "15" "onesparse_15" "red" >}} {{< bg "14" "onesparse_14" "red" >}} | `onesparse_$v` | `graphblas`, `lagraph` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-onesparse" "green" >}} {{< bg "17" "postgresql-17-onesparse" "red" >}} {{< bg "16" "postgresql-16-onesparse" "red" >}} {{< bg "15" "postgresql-15-onesparse" "red" >}} {{< bg "14" "postgresql-14-onesparse" "red" >}} | `postgresql-$v-onesparse` | `libgraphblas10`, `liblagraph1`, `liblagraphx1` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "onesparse_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "onesparse_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "onesparse_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-onesparse : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-onesparse : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-onesparse : MISS 0" "red" >}}      |


{{< tabs items="PG18" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `onesparse_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 222.5 KiB | [onesparse_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/onesparse_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `onesparse_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 201.9 KiB | [onesparse_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/onesparse_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `onesparse_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 195.3 KiB | [onesparse_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/onesparse_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `onesparse_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 182.7 KiB | [onesparse_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/onesparse_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `onesparse_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 200.2 KiB | [onesparse_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/onesparse_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `onesparse_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 185.5 KiB | [onesparse_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/onesparse_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-onesparse` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 598.3 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 578.4 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 592.1 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 574.7 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 693.6 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 681.1 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 645.2 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-onesparse` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 634.9 KiB | [postgresql-18-onesparse_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/onesparse/postgresql-18-onesparse_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/OneSparse/OneSparse" title="Repository" icon="github" subtitle="github.com/OneSparse/OneSparse" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="onesparse-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg one_sparse;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install one_sparse;		# install via package name, for the active PG version
pig install onesparse;		# install by extension name, for the current active PG version

pig install onesparse -v 18;   # install for PG 18

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION onesparse;
```

## Usage
> Sources: [homepage](https://onesparse.github.io/OneSparse/), [intro docs](https://onesparse.github.io/OneSparse/docs.html), [Matrix](https://onesparse.github.io/OneSparse/test_matrix_header.html), [Vector](https://onesparse.github.io/OneSparse/test_vector_header.html), and [Algorithms](https://onesparse.github.io/OneSparse/test_examples_header.html).

OneSparse is a PostgreSQL extension that binds SuiteSparse:GraphBLAS into Postgres and exposes sparse linear algebra and graph algorithms as new types, functions, and operators.
The docs treat `matrix` as the core type, with `vector` and `scalar` built on top of the same model.

### Core Setup

```sql
CREATE EXTENSION onesparse;
SET search_path TO public,onesparse;

SELECT 'int32'::matrix;
SELECT 'int32'::vector;
SELECT 'int32:42'::scalar;
```

The intro docs note that OneSparse keeps its API in the `onesparse` schema, and the matrix/vector pages show the same `search_path` pattern for interactive use.

### Matrix and Vector

The matrix page shows common operations such as constructing, printing, drawing, resizing, casting, and aggregating matrices.
The vector page shows the matching vector API, including `nvals`, `size`, `set_element`, `get_element`, `eadd`, `emult`, `reduce_scalar`, `choose`, and `apply`.

```sql
SELECT print('int32(4:4)'::matrix);
SELECT draw('int32(4:4)[1:2:1 2:3:2 3:1:3]'::matrix);
SELECT eadd('int32[0:1 1:2 2:3]'::vector, 'int32[0:1 1:2 2:3]'::vector, 'plus_int32');
SELECT reduce_scalar('int32[0:1 1:2 2:3]'::vector, 'plus_monoid_int32');
```

### Graph Algorithms

The getting-started docs use graph examples built from Matrix Market files and random graphs.
They highlight these algorithms:

- `bfs(graph, 1)` for level and parent BFS
- `sssp(cast_to(graph, 'int32'), 1::bigint, 1)` for single-source shortest path
- `pagerank(graph)` for ranking vertices by link structure
- `triangle_centrality(graph)` for triangle-based centrality
- `betweenness(graph, ARRAY[...])` and `square_clustering(graph)` for additional graph analysis

Representative example from the docs:

```sql
SELECT draw(triu(graph), (SELECT level FROM bfs(graph, 1)), false, false, true, 0.5)
FROM karate;
```

The same guide also shows graph loading with `mmread(...)` and graph visualization with `draw(...)`.

### Scope

The documentation set is broad. This stub captures the core interface and the main examples that repeat across the intro, matrix, vector, and algorithms pages, without reproducing the full GraphBLAS catalog.
