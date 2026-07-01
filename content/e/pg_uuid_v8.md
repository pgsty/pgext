---
title: "pg_uuid_v8"
linkTitle: "pg_uuid_v8"
description: "UUID v8 generator with embedded timestamps for PostgreSQL"
weight: 4530
categories: ["FUNC"]
width: full
---

[**pg_uuid_v8**](https://github.com/ineron/pg_uuid_v8) : UUID v8 generator with embedded timestamps for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4530** | {{< badge content="pg_uuid_v8" link="https://github.com/ineron/pg_uuid_v8" >}} | {{< ext "pg_uuid_v8" >}} | `1.0.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **See Also**    | {{< ext "uuid-ossp" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} |

> [!Note] Pinned to public so uuid operator commutators resolve on PostgreSQL 17 and 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_uuid_v8` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_uuid_v8_18" "green" >}} {{< bg "17" "pg_uuid_v8_17" "green" >}} {{< bg "16" "pg_uuid_v8_16" "green" >}} {{< bg "15" "pg_uuid_v8_15" "green" >}} {{< bg "14" "pg_uuid_v8_14" "green" >}} | `pg_uuid_v8_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-uuid-v8" "green" >}} {{< bg "17" "postgresql-17-pg-uuid-v8" "green" >}} {{< bg "16" "postgresql-16-pg-uuid-v8" "green" >}} {{< bg "15" "postgresql-15-pg-uuid-v8" "green" >}} {{< bg "14" "postgresql-14-pg-uuid-v8" "green" >}} | `postgresql-$v-pg-uuid-v8` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_uuid_v8_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-uuid-v8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-uuid-v8 : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.7 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_uuid_v8_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.6 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.4 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.4 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.1 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.6 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-uuid-v8` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.6 KiB | [postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-18-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_uuid_v8_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.5 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.1 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.1 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.1 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.6 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-uuid-v8` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.6 KiB | [postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-17-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_uuid_v8_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.5 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.1 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.1 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.1 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.6 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-uuid-v8` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-16-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_uuid_v8_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.5 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.1 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.1 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.1 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.6 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-uuid-v8` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.6 KiB | [postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-15-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuid_v8_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uuid_v8_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uuid_v8_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uuid_v8_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uuid_v8_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uuid_v8_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uuid_v8_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uuid_v8_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uuid_v8_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_uuid_v8_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uuid_v8_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.5 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.0 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.6 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-uuid-v8` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uuid-v8/postgresql-14-pg-uuid-v8_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ineron/pg_uuid_v8" title="Repository" icon="github" subtitle="github.com/ineron/pg_uuid_v8" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_uuid_v8-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_uuid_v8;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

Sources: [pg_uuid_v8 README](https://github.com/ineron/pg_uuid_v8), [SQL definitions](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8--1.0.sql), [control file](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8.control).

`pg_uuid_v8` generates UUIDs that look like UUID v4 values while embedding encrypted microsecond timestamps for extraction, sorting, and range predicates. The SQL file exposes both `uuid_stego_*` names and `uuid_v8_*` convenience aliases.

### Generate UUIDs

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_generate();
```

The equivalent lower-level generator is:

```sql
SELECT uuid_stego_generate();
```

Use a default expression when inserting events:

```sql
CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz DEFAULT now()
);
```

### Extract And Query Hidden Timestamps

Extract the embedded timestamp as microseconds since the Unix epoch:

```sql
SELECT uuid_v8_extract_timestamp(id)
FROM events
ORDER BY uuid_v8_extract_timestamp(id)
LIMIT 10;
```

The README recommends functional indexes for time-based lookups:

```sql
CREATE INDEX events_uuid_v8_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

Helper functions convert between timestamps and the integer timestamp format:

```sql
SELECT timestamp_to_stego_time(now());
SELECT stego_time_to_timestamp(uuid_v8_extract_timestamp(id))
FROM events;
```

### Range Helpers And Operators

The SQL definition includes direct range helpers:

```sql
SELECT *
FROM events
WHERE uuid_stego_in_range(
  id,
  now() - interval '24 hours',
  now()
);
```

It also defines timestamp-aware comparison functions and operators for `uuid`:

- `uuid_stego_compare(uuid, uuid)` and `uuid_v8_compare(uuid, uuid)`.
- `uuid_stego_lt`, `uuid_stego_le`, `uuid_stego_gt`, `uuid_stego_ge`.
- Operators `<`, `<=`, `>`, and `>=` compare UUIDs by hidden timestamp.

### Seed And Encryption Mode

Set and inspect the seed:

```sql
SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_get_seed();
```

Available encryption modes are `XOR`, `AES128`, and `AES256`:

```sql
SELECT uuid_v8_get_encryption_mode();
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('XOR');
```

For a persistent default, the README documents the `uuid_v8.encryption_mode` GUC:

```sql
ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

### Caveats

- Keep the seed secret; it is required to interpret hidden timestamps.
- UUIDs generated with one seed and encryption mode must be decoded with the same settings.
- Functional indexes on extracted timestamps add storage and update overhead, but are the intended path for efficient time-range predicates.
- Local Pigsty metadata pins this extension to the `public` schema so the UUID comparison operator commutators resolve on PostgreSQL 17 and 18; test operators explicitly if using a different schema in a non-Pigsty build.
