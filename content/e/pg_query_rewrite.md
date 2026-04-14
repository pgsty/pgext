---
title: "pg_query_rewrite"
linkTitle: "pg_query_rewrite"
description: "Rewrite SQL statements with a PostgreSQL ProcessUtility hook"
weight: 5030
categories: ["ADMIN"]
width: full
---

[**pg_query_rewrite**](https://github.com/pierreforstmann/pg_query_rewrite) : Rewrite SQL statements with a PostgreSQL ProcessUtility hook


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5030** | {{< badge content="pg_query_rewrite" link="https://github.com/pierreforstmann/pg_query_rewrite" >}} | {{< ext "pg_query_rewrite" >}} | `0.0.5` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Requires shared_preload_libraries=pg_query_rewrite.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_query_rewrite` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "pg_query_rewrite_18" "green" >}} {{< bg "17" "pg_query_rewrite_17" "green" >}} {{< bg "16" "pg_query_rewrite_16" "green" >}} {{< bg "15" "pg_query_rewrite_15" "green" >}} {{< bg "14" "pg_query_rewrite_14" "green" >}} | `pg_query_rewrite_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "postgresql-18-pg-query-rewrite" "green" >}} {{< bg "17" "postgresql-17-pg-query-rewrite" "green" >}} {{< bg "16" "postgresql-16-pg-query-rewrite" "green" >}} {{< bg "15" "postgresql-15-pg-query-rewrite" "green" >}} {{< bg "14" "postgresql-14-pg-query-rewrite" "green" >}} | `postgresql-$v-pg-query-rewrite` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_query_rewrite_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pg-query-rewrite : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pg-query-rewrite : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_query_rewrite_18` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_query_rewrite_18-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_query_rewrite_18` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_query_rewrite_18-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_query_rewrite_18` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_query_rewrite_18-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_query_rewrite_18` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_query_rewrite_18-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_query_rewrite_18` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_query_rewrite_18-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_query_rewrite_18` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.1 KiB | [pg_query_rewrite_18-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_query_rewrite_18-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.1 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.0 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.1 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.1 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.5 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.6 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.1 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-query-rewrite` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-18-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_query_rewrite_17` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_query_rewrite_17-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_query_rewrite_17` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_query_rewrite_17-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_query_rewrite_17` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.9 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_query_rewrite_17-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_query_rewrite_17` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_query_rewrite_17-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_query_rewrite_17` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_query_rewrite_17-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_query_rewrite_17` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_query_rewrite_17-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_query_rewrite_17-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.1 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.1 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.1 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.2 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.1 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-query-rewrite` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-17-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_query_rewrite_16` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_query_rewrite_16-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_query_rewrite_16` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_query_rewrite_16-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_query_rewrite_16` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.9 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_query_rewrite_16-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_query_rewrite_16` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_query_rewrite_16-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_query_rewrite_16` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_query_rewrite_16-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_query_rewrite_16` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [pg_query_rewrite_16-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_query_rewrite_16-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.1 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.1 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.2 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.3 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.1 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-query-rewrite` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-16-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_query_rewrite_15` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_query_rewrite_15-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_query_rewrite_15` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.1 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_query_rewrite_15-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_query_rewrite_15` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.0 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_query_rewrite_15-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_query_rewrite_15` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_query_rewrite_15-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_query_rewrite_15` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_query_rewrite_15-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_query_rewrite_15` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [pg_query_rewrite_15-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_query_rewrite_15-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.1 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.1 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.3 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.3 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.2 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-query-rewrite` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-15-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_query_rewrite_14` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_query_rewrite_14-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_query_rewrite_14` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_query_rewrite_14-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_query_rewrite_14` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.9 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_query_rewrite_14-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_query_rewrite_14` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_query_rewrite_14-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_query_rewrite_14` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_query_rewrite_14-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_query_rewrite_14` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [pg_query_rewrite_14-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_query_rewrite_14-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.8 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.7 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.9 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.8 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.4 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.5 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.9 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-query-rewrite` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.9 KiB | [postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-query-rewrite/postgresql-14-pg-query-rewrite_0.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pierreforstmann/pg_query_rewrite" title="Repository" icon="github" subtitle="github.com/pierreforstmann/pg_query_rewrite" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_query_rewrite-0.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_query_rewrite;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_query_rewrite;		# install via package name, for the active PG version

pig install pg_query_rewrite -v 18;   # install for PG 18
pig install pg_query_rewrite -v 17;   # install for PG 17
pig install pg_query_rewrite -v 16;   # install for PG 16
pig install pg_query_rewrite -v 15;   # install for PG 15
pig install pg_query_rewrite -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_query_rewrite';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_query_rewrite;
```

## Usage
- GitHub Repo: [`pierreforstmann/pg_query_rewrite`](https://github.com/pierreforstmann/pg_query_rewrite)
- README: [pierreforstmann/pg_query_rewrite/blob/master/README.md](https://github.com/pierreforstmann/pg_query_rewrite/blob/master/README.md)

`pg_query_rewrite` is a PostgreSQL extension that translates an exact source SQL statement into another predefined SQL statement. It must be loaded at server level with `shared_preload_libraries`, then installed in each database.

The README shows the extension as a simple statement rewrite engine backed by shared-memory rules.

### Setup

```conf
shared_preload_libraries = 'pg_query_rewrite'
pg_query_rewrite.max_rules = 10
```

```sql
CREATE EXTENSION pg_query_rewrite;
```

The README says the extension has been successfully tested with PostgreSQL 9.5 through 18.

### Managing Rules

Use the helper functions to manage translations:

```sql
SELECT pgqr_add_rule('select 10;', 'select 11;');
SELECT pgqr_remove_rule('select 10;');
SELECT pgqr_truncate();
SELECT pgqr_rules();
```

The example in the README rewrites `select 10;` to `select 11;` and then shows the rule list after insertion.

### Behavior

- Matching is exact and sensitive to case, spaces, and semicolons.
- Parameterized SQL statements are not supported.
- The maximum SQL statement length is hard-coded at 32K.
- Rules live only in shared memory; the extension does not persist settings on its own.
- `pg_query_rewrite.max_rules` caps the number of SQL statements that can be translated and defaults to `10` when unset.

### Scope

The upstream README is sufficient here: it covers purpose, server-level loading, rule management, a concrete rewrite example, and the main limitations. No separate docs page or homepage was needed.
