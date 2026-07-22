---
title: "pg_extension_updater"
linkTitle: "pg_extension_updater"
description: "Automatic extension updater"
weight: 2562
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_updater) : Automatic extension updater


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2562** | {{< badge content="pg_extension_updater" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_updater" >}} | {{< ext "pg_extension_updater" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `extension_updater` `pg_catalog` |
|   **Requires**    | {{< ext "pg_extension_base" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_base" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] Optional packaged component. It runs installed extension updates at database start through pg_extension_base; it is not part of CREATE EXTENSION pg_lake CASCADE dependency closure.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `pg_extension_base` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "pg_lake_18" "green" >}} {{< bg "17" "pg_lake_17" "green" >}} {{< bg "16" "pg_lake_16" "green" >}} {{< bg "15" "pg_lake_15" "red" >}} {{< bg "14" "pg_lake_14" "red" >}} | `pg_lake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "postgresql-18-pg-lake" "green" >}} {{< bg "17" "postgresql-17-pg-lake" "green" >}} {{< bg "16" "postgresql-16-pg-lake" "green" >}} {{< bg "15" "postgresql-15-pg-lake" "red" >}} {{< bg "14" "postgresql-14-pg-lake" "red" >}} | `postgresql-$v-pg-lake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_updater" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_updater" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lake-3.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_lake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_lake;		# install via package name, for the active PG version
pig install pg_extension_updater;		# install by extension name, for the current active PG version

pig install pg_extension_updater -v 18;   # install for PG 18
pig install pg_extension_updater -v 17;   # install for PG 17
pig install pg_extension_updater -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_extension_updater CASCADE; -- requires pg_extension_base
```

## Usage

Sources:

- [Official pg_extension_updater README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/pg_extension_updater.control)
- [Worker registration SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/pg_extension_updater--1.1.sql)
- [Updater implementation and configuration](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/src/pg_extension_updater.c)
- [pg_extension_base preload documentation](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/README.md)

`pg_extension_updater` runs `ALTER EXTENSION ... UPDATE` for each extension whose installed version differs from its available default version when the database lifecycle worker starts. It is intended to reduce SQL/binary version mismatches after new extension files are deployed; it does not install missing extensions or replace release testing.

### Enable Automatic Updates

The updater depends on `pg_extension_base`, which must be preloaded cluster-wide:

```conf
shared_preload_libraries = 'pg_extension_base'
```

After restarting PostgreSQL, create the updater in every database where automatic updates are desired:

```sql
CREATE EXTENSION pg_extension_updater CASCADE;
```

Creating it in `template1` causes new databases cloned from that template to contain the updater. The worker does not update extensions inside the template database itself, but starts in a newly cloned database.

### Runtime Behavior

- The `#!shared_preload_libraries` directive in the updater's control file lets `pg_extension_base` load its library.
- Installation registers the internal `extension_updater.main(internal)` lifecycle worker.
- At worker startup, it reads `pg_available_extensions` and updates entries whose `installed_version` differs from `default_version`.
- A failed `ALTER EXTENSION` is logged as a warning and is attempted only once during that worker run.
- `pg_extension_updater.enable` is a postmaster setting, defaults to `on`, and disables the worker's updates when set to `off`; changing it requires a restart.

### Caveats

- Automatic migration can execute arbitrary upgrade SQL supplied by every installed extension. Validate packages and upgrade paths before enabling it on production databases.
- Review extension dependency changes and take application-specific backups independently; a warning does not roll back unrelated successful extension updates.
- There is no user-facing force-update function or per-extension allowlist in version `3.4`.
- Version `3.4` changes no updater SQL API relative to `3.3`.
