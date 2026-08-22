---
title: "pg_uuid_v8"
linkTitle: "pg_uuid_v8"
description: "UUID v8 generator with embedded timestamps for PostgreSQL"
weight: 4530
categories: ["FUNC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_uuid_v8**](https://github.com/ineron/pg_uuid_v8) : UUID v8 generator with embedded timestamps for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4530** | {{< badge content="pg_uuid_v8" link="https://github.com/ineron/pg_uuid_v8" >}} | {{< ext "pg_uuid_v8" >}} | `1.1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **See Also**    | {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "snowflake" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "uuid-ossp" >}} {{< ext "typeid" >}} {{< ext "permuteseq" >}} |

> [!Note] Upstream 1.1.0 ships on PGXN only; pinned to public so uuid operator commutators resolve on PostgreSQL 17 and 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_uuid_v8` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_uuid_v8_18" "green" >}} {{< bg "17" "pg_uuid_v8_17" "green" >}} {{< bg "16" "pg_uuid_v8_16" "green" >}} {{< bg "15" "pg_uuid_v8_15" "green" >}} {{< bg "14" "pg_uuid_v8_14" "green" >}} | `pg_uuid_v8_$v` | `openssl` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pg-uuid-v8" "green" >}} {{< bg "17" "postgresql-17-pg-uuid-v8" "green" >}} {{< bg "16" "postgresql-16-pg-uuid-v8" "green" >}} {{< bg "15" "postgresql-15-pg-uuid-v8" "green" >}} {{< bg "14" "postgresql-14-pg-uuid-v8" "green" >}} | `postgresql-$v-pg-uuid-v8` | `libssl3 | libssl3t64` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.1 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.2 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.0 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.8 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.8 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.9 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.2 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.1 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.1 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.9 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.5 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.9 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.2 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.2 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.9 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.5 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.9 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.2 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.2 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.9 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.5 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.9 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.6 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.2 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.5 KiB | [pg_uuid_v8_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.1 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.1 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.9 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.4 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.9 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ineron/pg_uuid_v8" title="Repository" icon="github" subtitle="github.com/ineron/pg_uuid_v8" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_uuid_v8-1.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_uuid_v8;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_uuid_v8;		# install via package name, for the active PG version

pig install pg_uuid_v8 -v 18;   # install for PG 18
pig install pg_uuid_v8 -v 17;   # install for PG 17
pig install pg_uuid_v8 -v 16;   # install for PG 16
pig install pg_uuid_v8 -v 15;   # install for PG 15
pig install pg_uuid_v8 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_uuid_v8;
```

## Usage

Sources:

- [pg_uuid_v8 1.1.0 on PGXN](https://pgxn.org/dist/pg_uuid_v8/1.1.0/)
- [pg_uuid_v8 1.1.0 README](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/README.md)
- [pg_uuid_v8 1.1.0 control file](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8.control)
- [pg_uuid_v8 1.0 base SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0.sql)
- [pg_uuid_v8 1.0 to 1.1 upgrade SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0--1.1.sql)
- [Pigsty pg_uuid_v8 package matrix](https://pgext.cloud/ext/pg_uuid_v8)

`pg_uuid_v8` 1.1.0 generates UUID values with UUID-v4 version and variant bits while embedding an obfuscated creation time in the random payload. Its `uuid_v8_*` convenience functions mirror the lower-level `uuid_stego_*` API. Use it when hidden time extraction and time-range indexing are useful, but do not treat the embedded value as an authentication token or a substitute for a separate trusted creation timestamp.

### Generate Values

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_set_encryption_mode('AES128');

CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO events(data) VALUES ('{"type":"login"}');
```

The upstream implementation defaults to a published built-in seed and `XOR` mode. Set a deployment-specific secret before generating values. `AES128` and `AES256` are also available, but the same seed and mode must be selected when extracting a value.

### Extract and Index the Hidden Time

```sql
SELECT
  uuid_v8_extract_timestamp(id) AS epoch_microseconds,
  stego_time_to_timestamp(uuid_v8_extract_timestamp(id)) AS created_time
FROM events;

CREATE INDEX events_uuid_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

`uuid_v8_extract_timestamp(uuid)` returns a microsecond-scaled `bigint` so it remains compatible with `timestamp_to_stego_time()` and `stego_time_to_timestamp()`. In version 1.1 the internal 48-bit field stores milliseconds, so the returned value has millisecond resolution and its last three decimal digits are zero.

`uuid_stego_in_range()` offers a boolean timestamp-range helper. A functional B-tree index on the extraction function is the explicit and predictable path for indexed time predicates.

### Compare Hidden Times

`uuid_v8_compare(uuid, uuid)` and `uuid_stego_compare(uuid, uuid)` return ordering by extracted hidden time. The extension also defines `<`, `<=`, `>`, and `>=` operators for UUID arguments.

Pigsty packages install these added operators in `public` and qualify their commutator and negator references for PostgreSQL 17 and 18 compatibility. PostgreSQL already has built-in UUID ordering operators, so use the comparison functions or a schema-qualified `OPERATOR(public.<)` expression when hidden-time semantics must be unambiguous.

### Seed and Mode Controls

```sql
SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_get_seed();

SELECT uuid_v8_set_encryption_mode('XOR');
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('AES256');
SELECT uuid_v8_get_encryption_mode();

ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

The seed is exposed as `uuid_v8.stego_seed` and the mode as `uuid_v8.encryption_mode`. Setter functions change the current session; configuration settings can establish defaults for later sessions. `uuid_v8_get_seed()` returns the active seed, so restrict database access accordingly and never log its result.

### Upgrade and Compatibility Boundaries

```sql
ALTER EXTENSION pg_uuid_v8 UPDATE TO '1.1';
```

Version 1.1 changes timestamp storage from microseconds to milliseconds. The old 48-bit microsecond field rolled over about every 8.9 years and could not reliably recover current absolute dates; the 48-bit millisecond field lasts about 8,925 years. Relative ordering of pre-1.1 values was unaffected, but absolute time extraction and range predicates for those existing values remain unreliable after the upgrade because their encoded representation is not rewritten.

The PGXN metadata targets PostgreSQL 12 or later; current Pigsty packages cover PostgreSQL 14–18. Pigsty packages pin the extension to `public` and make it non-relocatable so the added operators resolve consistently. Keep an ordinary `created_at` column when provenance, auditability, sub-millisecond precision, or migrations across seeds and modes matter.
