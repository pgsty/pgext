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
| **9590** | {{< badge content="pgclone" link="https://github.com/valehdba/pgclone" >}} | {{< ext "pgclone" >}} | `4.4.2` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "db_migrator" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} {{< ext "pgactive" >}} |

> [!Note] preload for async/progress; RPM LLVM_BINPATH build fix retained in the 4.4.2 package.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.4.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgclone` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.4.2` | {{< bg "18" "pgclone_18" "green" >}} {{< bg "17" "pgclone_17" "green" >}} {{< bg "16" "pgclone_16" "green" >}} {{< bg "15" "pgclone_15" "green" >}} {{< bg "14" "pgclone_14" "green" >}} | `pgclone_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.4.2` | {{< bg "18" "postgresql-18-pgclone" "green" >}} {{< bg "17" "postgresql-17-pgclone" "green" >}} {{< bg "16" "postgresql-16-pgclone" "green" >}} {{< bg "15" "postgresql-15-pgclone" "green" >}} {{< bg "14" "postgresql-14-pgclone" "green" >}} | `postgresql-$v-pgclone` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.4.2" "postgresql-14-pgclone : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_18` | `4.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.0 KiB | [pgclone_18-4.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_18-4.4.2-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_18` | `4.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.5 KiB | [pgclone_18-4.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_18-4.4.2-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_18` | `4.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.3 KiB | [pgclone_18-4.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_18-4.4.2-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_18` | `4.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 96.1 KiB | [pgclone_18-4.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_18-4.4.2-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_18` | `4.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.7 KiB | [pgclone_18-4.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_18-4.4.2-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_18` | `4.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 96.2 KiB | [pgclone_18-4.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_18-4.4.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgclone` | `4.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 222.0 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.8 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 222.3 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 216.5 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 220.8 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 220.0 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 216.1 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 214.2 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 213.8 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgclone` | `4.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 212.6 KiB | [postgresql-18-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-18-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_17` | `4.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.0 KiB | [pgclone_17-4.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_17-4.4.2-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_17` | `4.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.5 KiB | [pgclone_17-4.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_17-4.4.2-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_17` | `4.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.4 KiB | [pgclone_17-4.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_17-4.4.2-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_17` | `4.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 96.1 KiB | [pgclone_17-4.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_17-4.4.2-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_17` | `4.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.7 KiB | [pgclone_17-4.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_17-4.4.2-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_17` | `4.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 96.2 KiB | [pgclone_17-4.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_17-4.4.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgclone` | `4.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 221.7 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.7 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 222.5 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 216.4 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 236.6 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 235.7 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 215.9 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 213.9 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 213.7 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgclone` | `4.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 212.4 KiB | [postgresql-17-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-17-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_16` | `4.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.0 KiB | [pgclone_16-4.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_16-4.4.2-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_16` | `4.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.4 KiB | [pgclone_16-4.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_16-4.4.2-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_16` | `4.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.3 KiB | [pgclone_16-4.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_16-4.4.2-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_16` | `4.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 96.0 KiB | [pgclone_16-4.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_16-4.4.2-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_16` | `4.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.6 KiB | [pgclone_16-4.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_16-4.4.2-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_16` | `4.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 96.1 KiB | [pgclone_16-4.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_16-4.4.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgclone` | `4.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 221.8 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.8 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 222.1 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 216.5 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 236.0 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 235.2 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 215.8 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 213.9 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 213.7 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgclone` | `4.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 212.4 KiB | [postgresql-16-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-16-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_15` | `4.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 97.9 KiB | [pgclone_15-4.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_15-4.4.2-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_15` | `4.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.5 KiB | [pgclone_15-4.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_15-4.4.2-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_15` | `4.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.3 KiB | [pgclone_15-4.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_15-4.4.2-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_15` | `4.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 96.2 KiB | [pgclone_15-4.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_15-4.4.2-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_15` | `4.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.8 KiB | [pgclone_15-4.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_15-4.4.2-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_15` | `4.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 96.2 KiB | [pgclone_15-4.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_15-4.4.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgclone` | `4.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 221.6 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.8 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 222.1 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 216.6 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 235.9 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 235.5 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 215.8 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 214.0 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 213.8 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgclone` | `4.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 212.2 KiB | [postgresql-15-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-15-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_14` | `4.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 97.9 KiB | [pgclone_14-4.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_14-4.4.2-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_14` | `4.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.4 KiB | [pgclone_14-4.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_14-4.4.2-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_14` | `4.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.1 KiB | [pgclone_14-4.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_14-4.4.2-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_14` | `4.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 96.1 KiB | [pgclone_14-4.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_14-4.4.2-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_14` | `4.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.5 KiB | [pgclone_14-4.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_14-4.4.2-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_14` | `4.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 96.0 KiB | [pgclone_14-4.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_14-4.4.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgclone` | `4.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 221.3 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.2 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 221.7 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 215.9 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 234.1 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 233.5 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 215.1 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 213.5 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 213.2 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgclone` | `4.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 211.8 KiB | [postgresql-14-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgclone/postgresql-14-pgclone_4.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/valehdba/pgclone" title="Repository" icon="github" subtitle="github.com/valehdba/pgclone" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgclone-4.4.2.tar.gz" >}}
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

## Usage

Sources:

- [pgclone v4.4.2 README](https://github.com/valehdba/pgclone/blob/v4.4.2/README.md)
- [pgclone v4.4.2 usage guide](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/USAGE.md)
- [Async cloning guide](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/ASYNC.md)
- [pgclone v4.4.2 release notes](https://github.com/valehdba/pgclone/releases/tag/v4.4.2)

pgclone clones tables, schemas, functions, roles, or whole databases over a PostgreSQL connection. It also provides preflight checks, structural diffs, masking, consistent snapshots, and optional background jobs. Use it for controlled database copies, not as an unattended substitute for backup and recovery.

### Create and Run a Clone

    CREATE EXTENSION pgclone;
    SELECT pgclone.version();

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'customers',
      true
    );

Schema and database entry points follow the same connection-first pattern:

    SELECT pgclone.schema(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true
    );

    SELECT pgclone.database(
      'host=source.example dbname=app user=clone_user',
      true
    );

The main API includes pgclone.table, pgclone.schema, pgclone.functions, pgclone.database, and pgclone.database_create. The _ex variants expose explicit choices for indexes, constraints, and triggers.

### Filter and Mask Data

JSON options can restrict columns and rows:

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'users',
      true,
      'users_lite',
      '{"columns":["id","name","email"],"where":"active"}'
    );

Version 4.4 adds schema- and database-level masks, table inclusion patterns, and exclude_tables. Mask expressions run in the source-side COPY query, so values that are successfully masked do not reach the target unmasked.

The 4.4.2 mask validator skips unsafe or incompatible masks: constant values that cannot cast to the column, NULL for NOT NULL columns, non-hash masks on unique or primary-key columns, and masks on foreign-key columns. A skipped mask leaves that column unmasked. Treat warnings as a failed privacy gate and inspect the result before distributing a clone.

### Preflight, Diff, and Consistency

    SELECT pgclone.preflight(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

    SELECT pgclone.diff(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

preflight checks connectivity, versions, privileges, capacity, names, roles, extensions, and tablespaces. diff reports DDL differences without applying changes.

Schema and database clones use a shared exported snapshot by default so related tables are copied consistently. A long snapshot can delay source vacuum cleanup and WAL recycling. Set the consistent option to false only when accepting cross-table inconsistency is an explicit tradeoff.

### Async Jobs

Async execution requires preload and a restart:

    shared_preload_libraries = 'pgclone'

    SELECT pgclone.schema_async(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true,
      '{"parallel":4}'
    );

    SELECT * FROM pgclone.jobs_view;
    SELECT pgclone.progress(1);
    SELECT pgclone.cancel(1);

pgclone also exposes progress_detail, resume, and clear_jobs for job administration. Size max_worker_processes for the requested parallelism.

### Important Boundaries

- The upstream usage guide requires superuser privileges to install and use pgclone.
- Async schema/database/parallel paths do not honor masks, tables, or exclude_tables in v4.4.2. Use the documented synchronous path when those controls are a security requirement.
- Keep passwords out of stored SQL and logs; prefer libpq service files, passfiles, or another controlled credential mechanism.
- Version 4.4.2 improves sequence-state copying and protects PostgreSQL 17 source sessions from transaction_timeout, but callers must still validate object ownership, extensions, roles, large objects, and post-clone application behavior.
