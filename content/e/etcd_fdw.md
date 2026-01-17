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
