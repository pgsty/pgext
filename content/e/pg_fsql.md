---
title: "pg_fsql"
linkTitle: "pg_fsql"
description: "Recursive SQL template engine with JSONB-driven execution"
weight: 4110
categories: ["UTIL"]
width: full
---

[**pg_fsql**](https://github.com/yurc/pg_fsql) : Recursive SQL template engine with JSONB-driven execution


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4110** | {{< badge content="pg_fsql" link="https://github.com/yurc/pg_fsql" >}} | {{< ext "pg_fsql" >}} | `1.1.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fsql` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pg_readme" >}} {{< ext "schedoc" >}} |

> [!Note] Release tag 1.1.0 still ships extension SQL version 1.0; shared_preload_libraries is optional and only needed for session-start GUC availability.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_fsql` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_fsql_18" "green" >}} {{< bg "17" "pg_fsql_17" "green" >}} {{< bg "16" "pg_fsql_16" "green" >}} {{< bg "15" "pg_fsql_15" "green" >}} {{< bg "14" "pg_fsql_14" "green" >}} | `pg_fsql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pg-fsql" "green" >}} {{< bg "17" "postgresql-17-pg-fsql" "green" >}} {{< bg "16" "postgresql-16-pg-fsql" "green" >}} {{< bg "15" "postgresql-15-pg-fsql" "green" >}} {{< bg "14" "postgresql-14-pg-fsql" "green" >}} | `postgresql-$v-pg-fsql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_fsql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-fsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-fsql : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fsql_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.2 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fsql_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fsql_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fsql_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fsql_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.0 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fsql_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fsql_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fsql_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fsql_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.1 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fsql_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fsql_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pg_fsql_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fsql_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-fsql` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.9 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.9 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.1 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.0 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.8 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-fsql` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.1 KiB | [postgresql-18-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-18-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fsql_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.2 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fsql_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fsql_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fsql_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fsql_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fsql_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fsql_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.9 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fsql_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fsql_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.1 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fsql_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fsql_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pg_fsql_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fsql_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-fsql` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.6 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.8 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.9 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.4 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.3 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fsql` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.0 KiB | [postgresql-17-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-17-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fsql_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.2 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fsql_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fsql_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fsql_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fsql_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.0 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fsql_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fsql_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fsql_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fsql_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.1 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fsql_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fsql_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pg_fsql_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fsql_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-fsql` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.6 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.8 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.9 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.3 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.7 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-fsql` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.0 KiB | [postgresql-16-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-16-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fsql_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.2 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fsql_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fsql_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.5 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fsql_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fsql_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fsql_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fsql_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fsql_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fsql_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.1 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fsql_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fsql_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pg_fsql_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fsql_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-fsql` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.9 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.0 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.5 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.4 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.8 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-fsql` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.1 KiB | [postgresql-15-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-15-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fsql_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.2 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fsql_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fsql_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.5 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fsql_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fsql_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.1 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fsql_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fsql_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fsql_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fsql_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.1 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fsql_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fsql_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pg_fsql_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fsql_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-fsql` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.6 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.9 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.9 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.4 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.3 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.8 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-fsql` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.0 KiB | [postgresql-14-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fsql/postgresql-14-pg-fsql_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/yurc/pg_fsql" title="Repository" icon="github" subtitle="github.com/yurc/pg_fsql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_fsql-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_fsql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_fsql;		# install via package name, for the active PG version

pig install pg_fsql -v 18;   # install for PG 18
pig install pg_fsql -v 17;   # install for PG 17
pig install pg_fsql -v 16;   # install for PG 16
pig install pg_fsql -v 15;   # install for PG 15
pig install pg_fsql -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_fsql CASCADE; -- requires plpgsql
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_fsql;
> INSERT INTO fsql.templates (path, cmd, body)
> VALUES ('user_count', 'exec',
>         'SELECT jsonb_build_object(''total'', count(*)) FROM users WHERE status = {d[status]!r}');
> SELECT fsql.run('user_count', '{"status":"active"}');
> ```
>
> Source: [README](https://github.com/yurc/pg_fsql)

`pg_fsql` is a recursive SQL template engine for PostgreSQL. It combines a C-based placeholder renderer with PL/pgSQL template execution, hierarchical template composition, and optional SPI plan caching. The upstream project emphasizes that it does not require superuser privileges.

## Core Objects

The extension installs two main catalog tables:

```sql
fsql.templates (
    path varchar(500) primary key,
    cmd varchar(50),
    body text,
    defaults text,
    cached boolean default false
)

fsql.params (
    key_param varchar(255) primary key,
    type_param varchar(255) not null
)
```

`path` is dot-separated and defines the template hierarchy.

## Template Commands

The README documents six command types:

- `exec` to execute SQL and return `jsonb`
- `ref` to redirect to another template
- `if` to choose a child branch
- `exec_tpl` to execute SQL and re-render the result as a template
- `map` to collect children into a JSON object
- `NULL` for text fragments inserted into parents

## Placeholders

The renderer supports placeholders such as:

- `{d[key]}`
- `{d[key]!r}` for `quote_literal`
- `{d[key]!j}` for JSONB literals
- `{d[key]!i}` for `quote_identifier`

The special key `_self` injects the full input JSON object.

## Public API

The upstream public functions include:

- `fsql.run(path, data, debug)` to execute a template tree
- `fsql.render(path, data)` to preview rendered SQL
- `fsql.tree(path)` to inspect hierarchy
- `fsql.explain(path, data)` to trace expansion
- `fsql.validate()` to check templates
- `fsql.depends_on(path)` to inspect dependencies
- `fsql.clear_cache()` to free cached SPI plans

## Example

```sql
INSERT INTO fsql.templates (path, cmd, body) VALUES
    ('report', 'exec',
     'SELECT jsonb_build_object(''data'', array_agg(row_to_json(t)))
      FROM (SELECT {d[cols]} FROM {d[src]} {d[where]}) t'),
    ('report.cols',  NULL, 'id, name, email'),
    ('report.src',   NULL, 'customers'),
    ('report.where', NULL, 'WHERE city = {d[city]!r}');

SELECT fsql.run('report', '{"city":"Moscow"}');
SELECT fsql.render('report', '{"city":"Moscow"}');
```

## Requirements

The README lists PostgreSQL 14+, `plpgsql`, and standard build dependencies such as `gcc`, `make`, and PostgreSQL server development headers.
