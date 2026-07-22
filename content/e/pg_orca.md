---
title: "pg_orca"
linkTitle: "pg_orca"
description: "ORCA query optimizer as a PostgreSQL extension"
weight: 2540
categories: ["OLAP"]
width: full
---

[**pg_orca**](https://github.com/quantumiodb/pgorca) : ORCA query optimizer as a PostgreSQL extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2540** | {{< badge content="pg_orca" link="https://github.com/quantumiodb/pgorca" >}} | {{< ext "pg_orca" >}} | `1.0.0` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} {{< ext "index_advisor" >}} |

> [!Note] PG18 only; use session_preload_libraries=pg_orca for automatic planner hook loading.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_orca` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_orca_18" "green" >}} {{< bg "17" "pg_orca_17" "red" >}} {{< bg "16" "pg_orca_16" "red" >}} {{< bg "15" "pg_orca_15" "red" >}} {{< bg "14" "pg_orca_14" "red" >}} | `pg_orca_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-orca" "green" >}} {{< bg "17" "postgresql-17-pg-orca" "red" >}} {{< bg "16" "postgresql-16-pg-orca" "red" >}} {{< bg "15" "postgresql-15-pg-orca" "red" >}} {{< bg "14" "postgresql-14-pg-orca" "red" >}} | `postgresql-$v-pg-orca` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_orca_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_orca_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_orca_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-orca : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-orca : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-orca : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orca_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [pg_orca_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orca_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orca_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_orca_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orca_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orca_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.2 MiB | [pg_orca_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orca_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orca_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_orca_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orca_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orca_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_orca_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orca_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orca_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_orca_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orca_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-orca` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.7 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.7 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.7 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.7 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.7 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-orca` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-orca_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-orca/postgresql-18-pg-orca_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/quantumiodb/pgorca" title="Repository" icon="github" subtitle="github.com/quantumiodb/pgorca" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_orca-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_orca;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_orca;		# install via package name, for the active PG version

pig install pg_orca -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_orca';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_orca;
```




## Usage

Sources: [pgorca README](https://github.com/quantumiodb/pgorca), [entrypoint/GUC source](https://github.com/quantumiodb/pgorca/blob/main/pg_orca.cpp), [control file](https://github.com/quantumiodb/pgorca/blob/main/pg_orca.control).

`pg_orca` plugs the ORCA cost-based optimizer from the Greenplum/Apache Cloudberry lineage into a standard PostgreSQL 18 server. The upstream README describes this project as PostgreSQL 18-only, and the local package metadata also marks it for PG18 only.

### Enable ORCA For A Session

`CREATE EXTENSION` loads the library in the current session, so the `pg_orca.*` GUCs and planner hook are available immediately:

```sql
CREATE EXTENSION pg_orca;

SET pg_orca.enable_orca = on;

EXPLAIN
SELECT *
FROM orders
WHERE customer_id = 42
  AND created_at >= now() - interval '30 days';
```

If ORCA cannot handle a query, the README says it falls back to the standard PostgreSQL planner automatically. Turn on fallback logging while validating a workload:

```sql
SET pg_orca.trace_fallback = on;
```

### Preload For New Connections

For automatic planner-hook loading in later sessions, upstream recommends `session_preload_libraries`, not `shared_preload_libraries`:

```sql
ALTER DATABASE mydb SET session_preload_libraries = 'pg_orca';
ALTER DATABASE mydb SET pg_orca.enable_orca = on;
```

Existing sessions are unaffected until they reconnect. If another session preload library is already configured, include both values explicitly:

```sql
ALTER DATABASE mydb
SET session_preload_libraries = 'pg_orca,pg_stat_statements';
```

Role-local and cluster-wide scopes are also valid:

```sql
ALTER ROLE bench SET session_preload_libraries = 'pg_orca';

ALTER SYSTEM SET session_preload_libraries = 'pg_orca';
SELECT pg_reload_conf();
```

### Tuning And Diagnostics

The README documents these main GUCs:

- `pg_orca.enable_orca`: enable ORCA; default `off`.
- `pg_orca.trace_fallback`: log fallback to the standard planner; default `off`.
- `optimizer_segments`: segment count for costing; default `1`.
- `optimizer_sort_factor`: sort cost scaling; default `1.0`.
- `optimizer_metadata_caching`: cache relation metadata; default `on`.
- `optimizer_mdcache_size`: metadata cache size in KB; default `16384`.
- `optimizer_search_strategy_path`: optional custom search strategy XML path.

The entrypoint source also defines additional ORCA tuning and debug GUCs such as `optimizer_join_order`, `pg_orca.join_order_dynamic_threshold`, `pg_orca.cost_model`, and `optimizer_print_*`. Treat those as workload/debug knobs and validate plans before keeping them in a persistent database setting.

### Caveats

- PostgreSQL 18 only.
- Use `session_preload_libraries = 'pg_orca'` for automatic loading in new sessions.
- ORCA is disabled by default after loading; set `pg_orca.enable_orca = on`.
- Fallback to the PostgreSQL planner is expected for unsupported queries; enable `pg_orca.trace_fallback` when checking coverage.
