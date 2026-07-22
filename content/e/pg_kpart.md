---
title: "pg_kpart"
linkTitle: "pg_kpart"
description: "Reject full partition scans that omit the partition key"
weight: 7450
categories: ["SEC"]
width: full
---

[**pg_kpart**](https://github.com/hexacluster/pg_kpart) : Reject full partition scans that omit the partition key


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7450** | {{< badge content="pg_kpart" link="https://github.com/hexacluster/pg_kpart" >}} | {{< ext "pg_kpart" >}} | `1.0` | {{< category "SEC" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL--r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_partman" >}} {{< ext "pg_fkpart" >}} {{< ext "plan_filter" >}} {{< ext "pg_hint_plan" >}} {{< ext "citus" >}} {{< ext "timescaledb" >}} |

> [!Note] Planner hook must be loaded through shared_preload_libraries or session_preload_libraries; CREATE EXTENSION is optional.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_kpart` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_kpart_18" "green" >}} {{< bg "17" "pg_kpart_17" "green" >}} {{< bg "16" "pg_kpart_16" "green" >}} {{< bg "15" "pg_kpart_15" "green" >}} {{< bg "14" "pg_kpart_14" "green" >}} | `pg_kpart_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-kpart" "green" >}} {{< bg "17" "postgresql-17-pg-kpart" "green" >}} {{< bg "16" "postgresql-16-pg-kpart" "green" >}} {{< bg "15" "postgresql-15-pg-kpart" "green" >}} {{< bg "14" "postgresql-14-pg-kpart" "green" >}} | `postgresql-$v-pg-kpart` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_kpart_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-kpart : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-kpart : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kpart_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.9 KiB | [pg_kpart_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kpart_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kpart_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_kpart_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kpart_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kpart_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.9 KiB | [pg_kpart_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kpart_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kpart_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_kpart_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kpart_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kpart_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_kpart_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kpart_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kpart_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [pg_kpart_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kpart_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-kpart` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.8 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.9 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.7 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.8 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.5 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-kpart` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-18-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-18-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kpart_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.9 KiB | [pg_kpart_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kpart_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kpart_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_kpart_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kpart_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kpart_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.9 KiB | [pg_kpart_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kpart_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kpart_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [pg_kpart_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kpart_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kpart_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_kpart_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kpart_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kpart_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [pg_kpart_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kpart_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-kpart` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.7 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.9 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.7 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.2 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.2 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.5 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-kpart` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-17-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-17-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kpart_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.9 KiB | [pg_kpart_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kpart_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kpart_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_kpart_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kpart_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kpart_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.9 KiB | [pg_kpart_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kpart_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kpart_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [pg_kpart_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kpart_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kpart_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_kpart_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kpart_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kpart_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_kpart_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kpart_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-kpart` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.7 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.8 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.7 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.6 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.5 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-kpart` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-16-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-16-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kpart_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.9 KiB | [pg_kpart_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kpart_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kpart_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_kpart_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kpart_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kpart_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.9 KiB | [pg_kpart_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kpart_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kpart_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [pg_kpart_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kpart_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kpart_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_kpart_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kpart_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kpart_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_kpart_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kpart_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-kpart` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.7 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.8 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.7 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.6 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.5 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-kpart` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.5 KiB | [postgresql-15-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-15-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kpart_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.9 KiB | [pg_kpart_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kpart_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kpart_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_kpart_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kpart_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kpart_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.9 KiB | [pg_kpart_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kpart_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kpart_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [pg_kpart_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kpart_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kpart_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_kpart_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kpart_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kpart_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [pg_kpart_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kpart_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-kpart` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.6 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.8 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.6 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.3 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.3 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.6 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 19.4 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-kpart` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.4 KiB | [postgresql-14-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kpart/postgresql-14-pg-kpart_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hexacluster/pg_kpart" title="Repository" icon="github" subtitle="github.com/hexacluster/pg_kpart" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_kpart-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_kpart;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_kpart;		# install via package name, for the active PG version

pig install pg_kpart -v 18;   # install for PG 18
pig install pg_kpart -v 17;   # install for PG 17
pig install pg_kpart -v 16;   # install for PG 16
pig install pg_kpart -v 15;   # install for PG 15
pig install pg_kpart -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_kpart';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [Official v1.0 README](https://github.com/HexaCluster/pg_kpart/blob/v1.0/README.md)
- [v1.0 control file](https://github.com/HexaCluster/pg_kpart/blob/v1.0/pg_kpart.control)

`pg_kpart` prevents accidental queries that would scan every leaf partition of a partitioned table without effective partition pruning. Its planner hook can raise, warn, or log before execution. The functional unit is the preloaded library; there are no SQL objects to create, and upstream describes `CREATE EXTENSION` only as optional catalog registration.

### Enable and Roll Out

For cluster-wide enforcement, preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_kpart'
```

It can also be loaded for selected sessions or databases without a server restart:

```conf
session_preload_libraries = 'pg_kpart'
```

Start in audit mode before enforcing errors:

```sql
ALTER SYSTEM SET pg_kpart.message_level = 'warning';
SELECT pg_reload_conf();
```

Once the observed queries are understood, set `pg_kpart.message_level = 'error'`.

### Scope and Behavior

```sql
-- Check only these tables and their sub-partitions.
ALTER SYSTEM SET pg_kpart.blacklisted =
    'public.measurement, public.orders';

-- Or check all partitioned tables except selected hierarchies.
ALTER SYSTEM SET pg_kpart.whitelisted = 'public.audit_log';
SELECT pg_reload_conf();
```

```sql
-- Partition key is logdate.
SELECT * FROM measurement WHERE city_id = 5;              -- violation
SELECT * FROM measurement WHERE logdate = DATE '2026-07-01'; -- pruned, allowed
SELECT * FROM measurement WHERE logdate = $1;             -- runtime pruning, allowed
```

Violations use SQLSTATE `FS001`, which applications can trap when `message_level` is `error`.

### Configuration Index and Caveats

- `pg_kpart.enabled`: master switch; default `on`.
- `pg_kpart.message_level`: `error`, `warning`, `notice`, `log`, and other PostgreSQL message levels.
- `pg_kpart.min_partitions`: minimum leaf-partition count to check; default `2`.
- `pg_kpart.check_superuser`: superusers bypass checks by default.
- `pg_kpart.blacklisted`: when nonempty, only named hierarchies are checked and `whitelisted` is ignored.
- `pg_kpart.whitelisted`: hierarchies exempt from checking when no blacklist is set.
- A predicate whose range still includes every partition is treated as a full scan and rejected, even if it mentions the partition key.
- The hook also applies to `UPDATE`, `DELETE`, and `EXPLAIN` without `ANALYZE`. It relies on PostgreSQL's planned pruning result, not textual inspection of `WHERE` clauses.
- Upstream v1.0 is tested on PostgreSQL 14 and newer.

