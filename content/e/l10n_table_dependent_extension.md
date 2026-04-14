---
title: "l10n_table_dependent_extension"
linkTitle: "l10n_table_dependent_extension"
description: "PostgreSQL l10n toolbox"
weight: 3671
categories: ["TYPE"]
width: full
---

[**pg_xenophile**](https://github.com/bigsmoke/pg_xenophile) : PostgreSQL l10n toolbox


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3671** | {{< badge content="l10n_table_dependent_extension" link="https://github.com/bigsmoke/pg_xenophile" >}} | {{< ext "l10n_table_dependent_extension" "pg_xenophile" >}} | `0.8.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_xenophile" >}} |
|   **See Also**    | {{< ext "country" >}} {{< ext "currency" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |
|    **Siblings**   | {{< ext "pg_xenophile" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_xenophile` | `pg_xenophile` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "pg_xenophile_18" "green" >}} {{< bg "17" "pg_xenophile_17" "green" >}} {{< bg "16" "pg_xenophile_16" "green" >}} {{< bg "15" "pg_xenophile_15" "green" >}} {{< bg "14" "pg_xenophile_14" "green" >}} | `pg_xenophile_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "postgresql-18-pg-xenophile" "green" >}} {{< bg "17" "postgresql-17-pg-xenophile" "green" >}} {{< bg "16" "postgresql-16-pg-xenophile" "green" >}} {{< bg "15" "postgresql-15-pg-xenophile" "green" >}} {{< bg "14" "postgresql-14-pg-xenophile" "green" >}} | `postgresql-$v-pg-xenophile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_xenophile" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_xenophile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_xenophile-0.8.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_xenophile;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_xenophile;		# install via package name, for the active PG version
pig install l10n_table_dependent_extension;		# install by extension name, for the current active PG version

pig install l10n_table_dependent_extension -v 18;   # install for PG 18
pig install l10n_table_dependent_extension -v 17;   # install for PG 17
pig install l10n_table_dependent_extension -v 16;   # install for PG 16
pig install l10n_table_dependent_extension -v 15;   # install for PG 15
pig install l10n_table_dependent_extension -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION l10n_table_dependent_extension CASCADE; -- requires pg_xenophile
```



## Usage

> [l10n_table_dependent_extension: l10n table dependent extension for pg_xenophile](https://github.com/bigsmoke/pg_xenophile)

The `l10n_table_dependent_extension` is a companion extension to `pg_xenophile` that provides infrastructure for other extensions that depend on its localization (l10n) table mechanism.

```sql
CREATE EXTENSION l10n_table_dependent_extension;
```

This extension works in conjunction with `pg_xenophile`'s `l10n_table` system, allowing dependent extensions to register their tables with the localization framework. When a base table is registered in `xeno.l10n_table`, the system automatically creates companion translation tables and language-specific views.

See the [`pg_xenophile`](https://github.com/bigsmoke/pg_xenophile) documentation for the full l10n table management API.
