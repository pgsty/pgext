---
title: "etcd_fdw"
linkTitle: "etcd_fdw"
description: "Foreign data wrapper for etcd"
weight: 8660
categories: ["FDW"]
width: full
---

[**etcd_fdw**](https://github.com/cybertec-postgresql/etcd_fdw) : Foreign data wrapper for etcd


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8660** | {{< badge content="etcd_fdw" link="https://github.com/cybertec-postgresql/etcd_fdw" >}} | {{< ext "etcd_fdw" >}} | `0.0.0` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "redis_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "mongo_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `etcd_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.0` | {{< bg "18" "etcd_fdw_18" "green" >}} {{< bg "17" "etcd_fdw_17" "green" >}} {{< bg "16" "etcd_fdw_16" "green" >}} {{< bg "15" "etcd_fdw_15" "green" >}} {{< bg "14" "etcd_fdw_14" "green" >}} {{< bg "13" "etcd_fdw_13" "green" >}} | `etcd_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.0` | {{< bg "18" "postgresql-18-etcd-fdw" "green" >}} {{< bg "17" "postgresql-17-etcd-fdw" "green" >}} {{< bg "16" "postgresql-16-etcd-fdw" "green" >}} {{< bg "15" "postgresql-15-etcd-fdw" "green" >}} {{< bg "14" "postgresql-14-etcd-fdw" "green" >}} {{< bg "13" "postgresql-13-etcd-fdw" "green" >}} | `postgresql-$v-etcd-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "etcd_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-18-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-17-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-etcd-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-etcd-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_18` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_18-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_18` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_18-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_18` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_18-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_18` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_18-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_18` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_18-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_18` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_18-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_18-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-18-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-18-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_17` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_17-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_17` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_17-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_17` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_17-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_17` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_17-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_17` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_17-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_17` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_17-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_17-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-17-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-17-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_16` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_16-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_16` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_16-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_16` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_16-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_16` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_16-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_16` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_16-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_16` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_16-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_16-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-16-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-16-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_15` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_15-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_15` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_15-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_15` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_15-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_15` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_15-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_15` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_15-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_15` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_15-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_15-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-15-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-15-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_14` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_14-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_14` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_14-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_14` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_14-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_14` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_14-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_14` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_14-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_14` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_14-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_14-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-14-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-14-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `etcd_fdw_13` | `0.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/etcd_fdw_13-0.0.0-1PIGSTY.el8.x86_64.rpm) |
| `etcd_fdw_13` | `0.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/etcd_fdw_13-0.0.0-1PIGSTY.el8.aarch64.rpm) |
| `etcd_fdw_13` | `0.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/etcd_fdw_13-0.0.0-1PIGSTY.el9.x86_64.rpm) |
| `etcd_fdw_13` | `0.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/etcd_fdw_13-0.0.0-1PIGSTY.el9.aarch64.rpm) |
| `etcd_fdw_13` | `0.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/etcd_fdw_13-0.0.0-1PIGSTY.el10.x86_64.rpm) |
| `etcd_fdw_13` | `0.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [etcd_fdw_13-0.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/etcd_fdw_13-0.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-etcd-fdw` | `0.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-13-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/etcd-fdw/postgresql-13-etcd-fdw_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/etcd_fdw" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/etcd_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="etcd_fdw-0.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg etcd_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install etcd_fdw;		# install via package name, for the active PG version

pig install etcd_fdw -v 18;   # install for PG 18
pig install etcd_fdw -v 17;   # install for PG 17
pig install etcd_fdw -v 16;   # install for PG 16
pig install etcd_fdw -v 15;   # install for PG 15
pig install etcd_fdw -v 14;   # install for PG 14
pig install etcd_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION etcd_fdw;
```



## Usage

- [Intro Blog](https://www.cybertec-postgresql.com/en/bringing-etcd-to-the-database-with-rust-and-pgrx/)
- [Github Repo](https://github.com/cybertec-postgresql/etcd_fdw)


-----------

## Quick Start

### 1. Enable Extension

```sql
CREATE EXTENSION etcd_fdw;
```

### 2. Create Foreign Data Wrapper

```sql
CREATE FOREIGN DATA WRAPPER etcd_fdw
  HANDLER etcd_fdw_handler
  VALIDATOR etcd_fdw_validator;
```

### 3. Create Server

```sql
-- Basic connection
CREATE SERVER etcd_plain
  FOREIGN DATA WRAPPER etcd_fdw
  OPTIONS (connstr '127.0.0.1:2379');

-- Production etcd with authentication and SSL
CREATE SERVER etcd FOREIGN DATA WRAPPER etcd_fdw OPTIONS (
    connstr '127.0.0.1:2379',
    username 'root',
    password 'Etcd.Root',
    ssl_ca '/pg/cert/ca.crt',
    ssl_cert '/pg/cert/server.crt',
    ssl_key '/pg/cert/server.key'
);
```

### 4. Create Foreign Table

```sql
-- Basic table mapping all keys
CREATE FOREIGN TABLE etcd_kv (key TEXT, value TEXT) SERVER etcd OPTIONS (rowid_column 'key');

-- Table with prefix filter (only keys starting with '/config/')
CREATE FOREIGN TABLE etcd_config (key TEXT, value TEXT)
  SERVER etcd OPTIONS (rowid_column 'key', prefix '/config/');
```

### 5. Query Data

```sql
-- Read all keys
SELECT * FROM etcd_kv;

-- Filter by key pattern (pushdown supported)
SELECT * FROM etcd_kv WHERE key LIKE '/app/%';

-- Range query
SELECT * FROM etcd_kv WHERE key >= '/a' AND key < '/b';

-- Insert new key
INSERT INTO etcd_kv (key, value) VALUES ('/mykey', 'myvalue');

-- Delete key
DELETE FROM etcd_kv WHERE key = '/mykey';
```

### 6. Real-time Sync with etcd

Changes made outside PostgreSQL are immediately visible:

```bash
# Insert via etcdctl
etcdctl put '/config/db_pool_size' '20'
```

```sql
-- Instantly visible in PostgreSQL
SELECT * FROM etcd_config;
     key               | value
-----------------------+-------
 /config/db_pool_size  | 20
(1 row)
```

-----------

## Reference

### Server Options

| Option            | Required | Description                            |
|-------------------|----------|----------------------------------------|
| `connstr`         | Yes      | etcd endpoint (e.g., `127.0.0.1:2379`) |
| `username`        | No       | Authentication username                |
| `password`        | No       | Authentication password                |
| `ssl_ca`          | No       | CA certificate file path               |
| `ssl_cert`        | No       | Client certificate file path           |
| `ssl_key`         | No       | Client private key file path           |
| `ssl_servername`  | No       | Domain name for TLS verification       |
| `connect_timeout` | No       | Connection timeout (default: `10s`)    |
| `request_timeout` | No       | Request timeout (default: `30s`)       |

### Foreign Table Options

| Option         | Default  | Description                              |
|----------------|----------|------------------------------------------|
| `rowid_column` | Required | Column used as unique row identifier     |
| `prefix`       | None     | Restrict to keys with this prefix        |
| `keys_only`    | `false`  | Fetch only keys, skip values             |
| `revision`     | `0`      | Read at specific etcd revision           |
| `key`          | `\0`     | Starting key for range scan              |
| `range_end`    | None     | Exclusive end key for range scan         |
| `consistency`  | `l`      | `l` (linearizable) or `s` (serializable) |

### Query Pushdown

The following operations are pushed down to etcd for better performance:

- **WHERE**: `=`, `>=`, `>`, `<=`, `<`, `BETWEEN`, `LIKE 'prefix%'`
- **ORDER BY**: Remote sorting
- **LIMIT/OFFSET**: Remote pagination

### Limitations

- `UPDATE` on key column is not supported. Workaround: `INSERT` new key, then `DELETE` old key.
- Requires etcd v3 API.

