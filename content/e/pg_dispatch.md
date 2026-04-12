---
title: "pg_dispatch"
linkTitle: "pg_dispatch"
description: "Asynchronous SQL dispatcher built on pg_cron"
weight: 1100
categories: ["TIME"]
width: full
---

[**pg_dispatch**](https://github.com/Snehil-Shah/pg_dispatch) : Asynchronous SQL dispatcher built on pg_cron


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1100** | {{< badge content="pg_dispatch" link="https://github.com/Snehil-Shah/pg_dispatch" >}} | {{< ext "pg_dispatch" >}} | `0.1.5` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcrypto" >}} {{< ext "pg_cron" >}} |
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pg_background" >}} |

> [!Note] Pure SQL extension; runtime also needs pgcrypto from contrib in addition to pg_cron.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dispatch` | `pgcrypto`, `pg_cron` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "pg_dispatch_18" "green" >}} {{< bg "17" "pg_dispatch_17" "green" >}} {{< bg "16" "pg_dispatch_16" "green" >}} {{< bg "15" "pg_dispatch_15" "green" >}} {{< bg "14" "pg_dispatch_14" "green" >}} | `pg_dispatch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "postgresql-18-pg-dispatch" "green" >}} {{< bg "17" "postgresql-17-pg-dispatch" "green" >}} {{< bg "16" "postgresql-16-pg-dispatch" "green" >}} {{< bg "15" "postgresql-15-pg-dispatch" "green" >}} {{< bg "14" "postgresql-14-pg-dispatch" "green" >}} | `postgresql-$v-pg-dispatch` | `postgresql-$v-cron` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "pg_dispatch_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-pg-dispatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-pg-dispatch : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dispatch_18` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dispatch_18-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_18` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dispatch_18-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_18` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dispatch_18-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_18` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dispatch_18-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_18` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dispatch_18-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `pg_dispatch_18` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pg_dispatch_18-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dispatch_18-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.9 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-dispatch` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 KiB | [postgresql-18-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-18-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dispatch_17` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dispatch_17-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_17` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dispatch_17-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_17` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dispatch_17-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_17` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dispatch_17-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_17` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dispatch_17-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `pg_dispatch_17` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pg_dispatch_17-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dispatch_17-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.9 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-dispatch` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 KiB | [postgresql-17-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-17-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dispatch_16` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dispatch_16-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_16` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dispatch_16-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_16` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dispatch_16-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_16` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dispatch_16-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_16` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dispatch_16-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `pg_dispatch_16` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pg_dispatch_16-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dispatch_16-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.9 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-dispatch` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 KiB | [postgresql-16-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-16-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dispatch_15` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dispatch_15-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_15` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dispatch_15-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_15` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dispatch_15-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_15` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dispatch_15-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_15` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dispatch_15-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `pg_dispatch_15` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pg_dispatch_15-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dispatch_15-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.9 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-dispatch` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 KiB | [postgresql-15-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-15-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dispatch_14` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dispatch_14-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_14` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dispatch_14-0.1.5-1PIGSTY.el8.noarch.rpm) |
| `pg_dispatch_14` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dispatch_14-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_14` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dispatch_14-0.1.5-1PIGSTY.el9.noarch.rpm) |
| `pg_dispatch_14` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dispatch_14-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `pg_dispatch_14` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pg_dispatch_14-0.1.5-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dispatch_14-0.1.5-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.9 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.9 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.9 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-dispatch` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.9 KiB | [postgresql-14-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dispatch/postgresql-14-pg-dispatch_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Snehil-Shah/pg_dispatch" title="Repository" icon="github" subtitle="github.com/Snehil-Shah/pg_dispatch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_dispatch-0.1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_dispatch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dispatch;		# install via package name, for the active PG version

pig install pg_dispatch -v 18;   # install for PG 18
pig install pg_dispatch -v 17;   # install for PG 17
pig install pg_dispatch -v 16;   # install for PG 16
pig install pg_dispatch -v 15;   # install for PG 15
pig install pg_dispatch -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dispatch CASCADE; -- requires pgcrypto, pg_cron
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION "Snehil_Shah@pg_dispatch";
> SELECT pgdispatch.fire('SELECT pg_sleep(40);');
> SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
> ```
>
> Sources: [README](https://github.com/Snehil-Shah/pg_dispatch), [database.dev page](https://database.dev/Snehil_Shah/pg_dispatch)

`pg_dispatch` is an asynchronous SQL dispatcher for PostgreSQL. It is designed as a TLE-compatible alternative to `pg_later`, built on top of `pg_cron`, so it can be used in environments such as Supabase and AWS RDS.

## Prerequisites

The upstream README lists:

- PostgreSQL 13 or newer
- `pg_cron` 1.5.0 or newer
- `pgcrypto`

## Installation

The documented TLE installation path is:

```sql
SELECT dbdev.install(Snehil_Shah@pg_dispatch);
CREATE EXTENSION "Snehil_Shah@pg_dispatch";
```

The README warns that the extension installs into the `pgdispatch` schema and can collide with an existing schema of the same name.

## Functions

### `pgdispatch.fire(command text)`

Dispatch an SQL command for asynchronous execution:

```sql
SELECT pgdispatch.fire('SELECT pg_sleep(40);');
```

### `pgdispatch.snooze(command text, delay interval)`

Dispatch a delayed SQL command:

```sql
SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
```

The README notes that the delay is scheduled asynchronously and does not block the caller's main transaction.

## Use Cases

The project positions itself for database-native async side effects, especially in PL/pgSQL or trigger-based workflows. Its example use case is moving expensive notification or analytics work out of an `AFTER INSERT` trigger so the main RPC or application request can return sooner.
