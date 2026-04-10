---
title: "pg_trickle"
linkTitle: "pg_trickle"
description: "Streaming tables and differential view maintenance for PostgreSQL 18"
weight: 2860
categories: ["FEAT"]
width: full
---

[**pg_trickle**](https://github.com/grove/pg-trickle) : Streaming tables and differential view maintenance for PostgreSQL 18


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2860** | {{< badge content="pg_trickle" link="https://github.com/grove/pg-trickle" >}} | {{< ext "pg_trickle" >}} | `0.16.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_trickle` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.0` | {{< bg "18" "pg_trickle_18" "green" >}} {{< bg "17" "pg_trickle_17" "red" >}} {{< bg "16" "pg_trickle_16" "red" >}} {{< bg "15" "pg_trickle_15" "red" >}} {{< bg "14" "pg_trickle_14" "red" >}} | `pg_trickle_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.0` | {{< bg "18" "postgresql-18-pg-trickle" "green" >}} {{< bg "17" "postgresql-17-pg-trickle" "red" >}} {{< bg "16" "postgresql-16-pg-trickle" "red" >}} {{< bg "15" "postgresql-15-pg-trickle" "red" >}} {{< bg "14" "postgresql-14-pg-trickle" "red" >}} | `postgresql-$v-pg-trickle` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "pg_trickle_18 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_trickle_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_trickle_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.16.0" "postgresql-18-pg-trickle : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-trickle : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-trickle : MISS 0" "red" >}}      |


{{< tabs items="PG18" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_trickle_18` | `0.16.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_trickle_18-0.16.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_trickle_18` | `0.16.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.8 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_trickle_18-0.16.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_trickle_18` | `0.16.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.2 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_trickle_18-0.16.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_trickle_18` | `0.16.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.9 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_trickle_18-0.16.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_trickle_18` | `0.16.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.2 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_trickle_18-0.16.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_trickle_18` | `0.16.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [pg_trickle_18-0.16.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_trickle_18-0.16.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-trickle` | `0.16.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.7 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-trickle` | `0.16.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.7 MiB | [postgresql-18-pg-trickle_0.16.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-trickle/postgresql-18-pg-trickle_0.16.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/grove/pg-trickle" title="Repository" icon="github" subtitle="github.com/grove/pg-trickle" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_trickle-0.16.0-complete.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_trickle;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_trickle;		# install via package name, for the active PG version

pig install pg_trickle -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_trickle';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_trickle;
```
