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
| **9590** | {{< badge content="pgclone" link="https://github.com/valehdba/pgclone" >}} | {{< ext "pgclone" >}} | `4.0.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "db_migrator" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} {{< ext "pgactive" >}} |

> [!Note] preload for async/progress


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgclone` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.0` | {{< bg "18" "pgclone_18" "green" >}} {{< bg "17" "pgclone_17" "green" >}} {{< bg "16" "pgclone_16" "green" >}} {{< bg "15" "pgclone_15" "green" >}} {{< bg "14" "pgclone_14" "green" >}} | `pgclone_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.0` | {{< bg "18" "postgresql-18-pgclone" "green" >}} {{< bg "17" "postgresql-17-pgclone" "green" >}} {{< bg "16" "postgresql-16-pgclone" "green" >}} {{< bg "15" "postgresql-15-pgclone" "green" >}} {{< bg "14" "postgresql-14-pgclone" "green" >}} | `postgresql-$v-pgclone` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "pgclone_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-18-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-17-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-16-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-15-pgclone : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.0" "postgresql-14-pgclone : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgclone : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgclone : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgclone : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_18` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.6 KiB | [pgclone_18-4.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_18-4.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_18` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.3 KiB | [pgclone_18-4.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_18-4.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_18` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.4 KiB | [pgclone_18-4.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_18-4.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_18` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [pgclone_18-4.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_18-4.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_18` | `4.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.2 KiB | [pgclone_18-4.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_18-4.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_18` | `4.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.0 KiB | [pgclone_18-4.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_18-4.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgclone` | `4.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.2 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.0 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 131.0 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.5 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 133.7 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 133.2 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.1 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgclone` | `4.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.1 KiB | [postgresql-18-pgclone_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-18-pgclone_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_17` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.6 KiB | [pgclone_17-4.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_17-4.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_17` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.3 KiB | [pgclone_17-4.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_17-4.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_17` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.3 KiB | [pgclone_17-4.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_17-4.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_17` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [pgclone_17-4.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_17-4.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_17` | `4.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.2 KiB | [pgclone_17-4.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_17-4.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_17` | `4.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.0 KiB | [pgclone_17-4.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_17-4.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgclone` | `4.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.1 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.9 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.8 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.5 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.7 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 144.4 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.0 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgclone` | `4.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.1 KiB | [postgresql-17-pgclone_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-17-pgclone_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_16` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.6 KiB | [pgclone_16-4.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_16-4.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_16` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.3 KiB | [pgclone_16-4.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_16-4.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_16` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.5 KiB | [pgclone_16-4.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_16-4.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_16` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [pgclone_16-4.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_16-4.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_16` | `4.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.2 KiB | [pgclone_16-4.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_16-4.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_16` | `4.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.9 KiB | [pgclone_16-4.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_16-4.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgclone` | `4.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 130.9 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.0 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.6 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.6 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.2 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 143.9 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.0 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgclone` | `4.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.1 KiB | [postgresql-16-pgclone_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-16-pgclone_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_15` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.5 KiB | [pgclone_15-4.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_15-4.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_15` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.3 KiB | [pgclone_15-4.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_15-4.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_15` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.2 KiB | [pgclone_15-4.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_15-4.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_15` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.0 KiB | [pgclone_15-4.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_15-4.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_15` | `4.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.1 KiB | [pgclone_15-4.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_15-4.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_15` | `4.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.0 KiB | [pgclone_15-4.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_15-4.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgclone` | `4.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 130.9 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.8 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.8 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.5 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.2 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 144.0 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.1 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgclone` | `4.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.1 KiB | [postgresql-15-pgclone_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-15-pgclone_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgclone_14` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.4 KiB | [pgclone_14-4.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgclone_14-4.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgclone_14` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.2 KiB | [pgclone_14-4.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgclone_14-4.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgclone_14` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.1 KiB | [pgclone_14-4.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgclone_14-4.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgclone_14` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [pgclone_14-4.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgclone_14-4.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgclone_14` | `4.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.1 KiB | [pgclone_14-4.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgclone_14-4.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgclone_14` | `4.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.8 KiB | [pgclone_14-4.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgclone_14-4.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgclone` | `4.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 130.4 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.5 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.3 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.1 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 143.0 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 142.4 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 129.7 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgclone` | `4.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 128.6 KiB | [postgresql-14-pgclone_4.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgclone/postgresql-14-pgclone_4.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/valehdba/pgclone" title="Repository" icon="github" subtitle="github.com/valehdba/pgclone" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgclone-4.0.0.tar.gz" >}}
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

Source: [README](https://github.com/valehdba/pgclone/blob/main/README.md), [Usage guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md), [Async guide](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md), [Release v4.0.0](https://github.com/valehdba/pgclone/releases/tag/v4.0.0), [SQL install script](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.0.0.sql)

`pgclone` clones tables, schemas, functions, roles, and whole databases directly from SQL. In v4.0.0 the public API is namespaced under the `pgclone` schema.

### Core clone functions

```sql
CREATE EXTENSION pgclone;

SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres password=secret',
  'public',
  'customers',
  true
);

SELECT pgclone.schema(
  'host=source-server dbname=mydb user=postgres password=secret',
  'sales',
  true
);

SELECT pgclone.database(
  'host=source-server dbname=mydb user=postgres password=secret',
  true
);
```

- `pgclone.table(...)`, `pgclone.schema(...)`, `pgclone.functions(...)`, `pgclone.database(...)`
- `pgclone.database_create(...)` creates a local target database and clones into it.
- `_ex` variants expose explicit booleans for indexes, constraints, and triggers.

### Options and masking

- JSON options support `columns`, `where`, `conflict`, and object toggles such as `indexes`, `constraints`, and `triggers`.
- Upstream documents masking, auto-discovery of sensitive columns, static masking, dynamic masking, clone verification, and GDPR/compliance reporting in the usage guide.

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

### Async and progress

```sql
-- postgresql.conf
shared_preload_libraries = 'pgclone'

SELECT pgclone.schema_async(
  'host=source-server dbname=mydb user=postgres',
  'sales', true, '{"parallel":4}'
);

SELECT * FROM pgclone.jobs_view;
SELECT pgclone.progress(1);
SELECT pgclone.cancel(1);
```

- `pgclone.table_async(...)` and `pgclone.schema_async(...)` run in background workers.
- `pgclone.jobs_view`, `pgclone.progress_detail()`, `pgclone.resume()`, and `pgclone.clear_jobs()` provide job tracking and recovery.

### Caveats

- Upstream requires PostgreSQL 14+.
- The usage guide states the extension requires superuser privileges to install and use.
- Async features need `shared_preload_libraries = 'pgclone'`; worker-pool parallelism also depends on `max_worker_processes`.
