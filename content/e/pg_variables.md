---
title: "pg_variables"
linkTitle: "pg_variables"
description: "Session-scoped variables with scalar, array, and record types"
weight: 2820
categories: ["FEAT"]
width: full
---

[**pg_variables**](https://github.com/postgrespro/pg_variables) : Session-scoped variables with scalar, array, and record types


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2820** | {{< badge content="pg_variables" link="https://github.com/postgrespro/pg_variables" >}} | {{< ext "pg_variables" >}} | `1.2.5` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "plisql" >}} |

> [!Note] Release tag 1.2.5 still ships extension SQL version 1.2.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_variables` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "pg_variables_18" "green" >}} {{< bg "17" "pg_variables_17" "green" >}} {{< bg "16" "pg_variables_16" "green" >}} {{< bg "15" "pg_variables_15" "green" >}} {{< bg "14" "pg_variables_14" "green" >}} | `pg_variables_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "postgresql-18-pg-variables" "green" >}} {{< bg "17" "postgresql-17-pg-variables" "green" >}} {{< bg "16" "postgresql-16-pg-variables" "green" >}} {{< bg "15" "postgresql-15-pg-variables" "green" >}} {{< bg "14" "postgresql-14-pg-variables" "green" >}} | `postgresql-$v-pg-variables` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "pg_variables_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-pg-variables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-pg-variables : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_variables_18` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.5 KiB | [pg_variables_18-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_variables_18-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_variables_18` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.0 KiB | [pg_variables_18-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_variables_18-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_variables_18` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.7 KiB | [pg_variables_18-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_variables_18-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_variables_18` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [pg_variables_18-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_variables_18-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_variables_18` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.4 KiB | [pg_variables_18-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_variables_18-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_variables_18` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.4 KiB | [pg_variables_18-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_variables_18-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-variables` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.0 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.7 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.2 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.8 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 65.5 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.4 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.1 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-variables` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.4 KiB | [postgresql-18-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-18-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_variables_17` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.4 KiB | [pg_variables_17-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_variables_17-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_variables_17` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.9 KiB | [pg_variables_17-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_variables_17-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_variables_17` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.7 KiB | [pg_variables_17-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_variables_17-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_variables_17` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.4 KiB | [pg_variables_17-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_variables_17-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_variables_17` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.4 KiB | [pg_variables_17-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_variables_17-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_variables_17` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.3 KiB | [pg_variables_17-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_variables_17-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-variables` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.8 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.8 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.0 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.9 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.1 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 71.0 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.9 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-variables` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.6 KiB | [postgresql-17-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-17-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_variables_16` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.4 KiB | [pg_variables_16-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_variables_16-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_variables_16` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.0 KiB | [pg_variables_16-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_variables_16-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_variables_16` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.7 KiB | [pg_variables_16-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_variables_16-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_variables_16` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.4 KiB | [pg_variables_16-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_variables_16-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_variables_16` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.4 KiB | [pg_variables_16-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_variables_16-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_variables_16` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.3 KiB | [pg_variables_16-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_variables_16-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-variables` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.7 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.8 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.9 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.8 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.8 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.6 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.8 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-variables` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.6 KiB | [postgresql-16-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-16-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_variables_15` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.5 KiB | [pg_variables_15-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_variables_15-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_variables_15` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.0 KiB | [pg_variables_15-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_variables_15-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_variables_15` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.0 KiB | [pg_variables_15-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_variables_15-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_variables_15` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.0 KiB | [pg_variables_15-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_variables_15-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_variables_15` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.5 KiB | [pg_variables_15-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_variables_15-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_variables_15` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.4 KiB | [pg_variables_15-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_variables_15-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-variables` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.6 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.5 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.7 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.5 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.7 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 71.0 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.6 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-variables` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.2 KiB | [postgresql-15-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-15-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_variables_14` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.5 KiB | [pg_variables_14-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_variables_14-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_variables_14` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.1 KiB | [pg_variables_14-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_variables_14-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_variables_14` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.1 KiB | [pg_variables_14-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_variables_14-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_variables_14` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.0 KiB | [pg_variables_14-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_variables_14-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_variables_14` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 35.5 KiB | [pg_variables_14-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_variables_14-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_variables_14` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.5 KiB | [pg_variables_14-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_variables_14-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-variables` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.8 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.6 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.9 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.6 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.0 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 71.1 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.8 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-variables` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 63.4 KiB | [postgresql-14-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-variables/postgresql-14-pg-variables_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/pg_variables" title="Repository" icon="github" subtitle="github.com/postgrespro/pg_variables" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_variables-1.2.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_variables;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_variables;		# install via package name, for the active PG version

pig install pg_variables -v 18;   # install for PG 18
pig install pg_variables -v 17;   # install for PG 17
pig install pg_variables -v 16;   # install for PG 16
pig install pg_variables -v 15;   # install for PG 15
pig install pg_variables -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_variables;
```

## Usage

- Sources: [README](https://github.com/postgrespro/pg_variables/blob/master/README.md), [repository tags](https://github.com/postgrespro/pg_variables/tags), [control file](https://github.com/postgrespro/pg_variables/blob/master/pg_variables.control)

`pg_variables` provides session-scoped variables grouped into named packages. Variables live only in the current session and are non-transactional by default unless created with `is_transactional := true`.

### Basic Set And Get

```sql
CREATE EXTENSION pg_variables;

SELECT pgv_set('vars', 'int1', 101);
SELECT pgv_get('vars', 'int1', NULL::int);
```

Transactional variables participate in savepoints and rollbacks:

```sql
BEGIN;
SELECT pgv_set('vars', 'trans_int', 101, true);
SAVEPOINT sp1;
SELECT pgv_set('vars', 'trans_int', 102, true);
ROLLBACK TO sp1;
COMMIT;
```

### Core APIs

The README documents generic scalar and array APIs:

- `pgv_set(package, name, value, is_transactional default false)`
- `pgv_get(package, name, NULL::type, strict default true)`

It also documents record-oriented APIs:

- `pgv_insert()`
- `pgv_update()`
- `pgv_delete()`
- `pgv_select()`

Useful administration helpers include `pgv_exists()`, `pgv_remove()`, `pgv_free()`, `pgv_list()`, and `pgv_stats()`.

### Error And Strictness Behavior

`pgv_get()` checks both existence and type. The README shows that missing packages, missing variables, or mismatched types raise errors unless `strict := false`, in which case `NULL` is returned for missing values.

### Deprecated Helpers And Version Note

Upstream still ships deprecated type-specific helpers such as `pgv_set_int()` / `pgv_get_int()` and `pgv_set_jsonb()` / `pgv_get_jsonb()`, but recommends the generic `pgv_set()` / `pgv_get()` API.

The repository tag is `v1.2.5`, while the current `pg_variables.control` file still declares `default_version = '1.2'`. That matches the packaging note that the release tag advanced without changing the SQL extension version string.
