---
title: "pgclone"
linkTitle: "pgclone"
description: "Clone PostgreSQL databases, schemas, tables, and functions across environments"
weight: 9590
categories: ["ETL"]
width: full
---

[**pgclone**](https://github.com/valehdba/pgclone) : Clone PostgreSQL databases, schemas, tables, and functions across environments


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9590** | {{< badge content="pgclone" link="https://github.com/valehdba/pgclone" >}} | {{< ext "pgclone" >}} | `2.2.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "db_migrator" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} {{< ext "pgactive" >}} |

> [!Note] shared_preload_libraries = pgclone is required for async/progress features; synchronous clone functions work without it.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgclone` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "pgclone_18" "green" >}} {{< bg "17" "pgclone_17" "green" >}} {{< bg "16" "pgclone_16" "green" >}} {{< bg "15" "pgclone_15" "green" >}} {{< bg "14" "pgclone_14" "green" >}} | `pgclone_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "postgresql-18-pgclone" "green" >}} {{< bg "17" "postgresql-17-pgclone" "green" >}} {{< bg "16" "postgresql-16-pgclone" "green" >}} {{< bg "15" "postgresql-15-pgclone" "green" >}} {{< bg "14" "postgresql-14-pgclone" "green" >}} | `postgresql-$v-pgclone` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [pgclone_18-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_18-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.5 KiB | [pgclone_18-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_18-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.6 KiB | [pgclone_18-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_18-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.1 KiB | [pgclone_18-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_18-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.7 KiB | [pgclone_18-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_18-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.9 KiB | [pgclone_18-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_18-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgclone` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.2 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 89.0 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 90.7 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.2 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.5 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.7 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.5 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgclone` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.3 KiB | [postgresql-18-pgclone_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [pgclone_17-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_17-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.5 KiB | [pgclone_17-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_17-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.5 KiB | [pgclone_17-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_17-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.0 KiB | [pgclone_17-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_17-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.7 KiB | [pgclone_17-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_17-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.8 KiB | [pgclone_17-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_17-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgclone` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.1 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 89.0 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 91.0 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.5 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 103.7 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 103.0 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.6 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgclone` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.3 KiB | [postgresql-17-pgclone_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.8 KiB | [pgclone_16-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_16-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.5 KiB | [pgclone_16-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_16-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.5 KiB | [pgclone_16-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_16-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.1 KiB | [pgclone_16-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_16-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.7 KiB | [pgclone_16-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_16-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.8 KiB | [pgclone_16-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_16-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgclone` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.2 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 89.0 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 90.9 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.4 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 103.2 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 102.5 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.5 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgclone` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.3 KiB | [postgresql-16-pgclone_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [pgclone_15-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_15-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.5 KiB | [pgclone_15-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_15-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.5 KiB | [pgclone_15-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_15-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.1 KiB | [pgclone_15-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_15-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.7 KiB | [pgclone_15-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_15-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.8 KiB | [pgclone_15-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_15-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgclone` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.2 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 89.0 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 90.9 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.5 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 103.2 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 102.6 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.4 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgclone` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 91.3 KiB | [postgresql-15-pgclone_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.6 KiB | [pgclone_14-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_14-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.3 KiB | [pgclone_14-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_14-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.4 KiB | [pgclone_14-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_14-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.1 KiB | [pgclone_14-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_14-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_14` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.7 KiB | [pgclone_14-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_14-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_14` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.8 KiB | [pgclone_14-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_14-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgclone` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 90.8 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 88.5 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 90.4 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.1 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 101.6 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 100.9 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 92.1 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgclone` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 90.8 KiB | [postgresql-14-pgclone_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_2.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/valehdba/pgclone" title="Repository" icon="github" subtitle="github.com/valehdba/pgclone" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgclone-2.2.0.zip" >}}
{{< /cards >}}


```bash
pig build pkg pgclone;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgclone;		# install via package name, for the active PG version

pig install pgclone -v 18;   # install for PG 18
pig install pgclone -v 17;   # install for PG 17
pig install pgclone -v 16;   # install for PG 16
pig install pgclone -v 15;   # install for PG 15
pig install pgclone -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgclone';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgclone;
```
