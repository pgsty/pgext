---
title: "kafka_fdw"
linkTitle: "kafka_fdw"
description: "kafka Foreign Data Wrapper for CSV formatted messages"
weight: 8730
categories: ["FDW"]
width: full
---

[**kafka_fdw**](https://github.com/adjust/kafka_fdw) : kafka Foreign Data Wrapper for CSV formatted messages


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8730** | {{< badge content="kafka_fdw" link="https://github.com/adjust/kafka_fdw" >}} | {{< ext "kafka_fdw" >}} | `0.0.3` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "mongo_fdw" >}} {{< ext "redis_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "redis" >}} {{< ext "hdfs_fdw" >}} {{< ext "wal2json" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `kafka_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "kafka_fdw_18" "green" >}} {{< bg "17" "kafka_fdw_17" "green" >}} {{< bg "16" "kafka_fdw_16" "green" >}} {{< bg "15" "kafka_fdw_15" "green" >}} {{< bg "14" "kafka_fdw_14" "green" >}} | `kafka_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "postgresql-18-kafka-fdw" "green" >}} {{< bg "17" "postgresql-17-kafka-fdw" "green" >}} {{< bg "16" "postgresql-16-kafka-fdw" "green" >}} {{< bg "15" "postgresql-15-kafka-fdw" "green" >}} {{< bg "14" "postgresql-14-kafka-fdw" "green" >}} | `postgresql-$v-kafka-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-kafka-fdw : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-kafka-fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_18` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.6 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_18` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.5 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_18` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.5 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_18` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.0 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_18` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.6 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_18` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 79.0 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 76.8 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 79.0 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 77.3 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 84.9 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.5 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.4 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-kafka-fdw` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.2 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_17` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.6 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_17` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.4 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_17` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.5 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_17` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.0 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_17` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.5 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_17` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.9 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 78.8 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 76.6 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 78.8 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 77.0 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 106.0 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 104.5 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.2 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-kafka-fdw` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.9 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_16` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.0 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_16` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.4 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_16` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 36.0 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_16` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.2 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_16` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.8 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_16` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 36.1 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.4 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.9 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.4 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.3 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 112.0 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 110.1 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.1 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-kafka-fdw` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.9 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_15` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.3 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_15` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_15` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 36.6 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_15` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.9 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_15` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 37.4 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_15` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 36.8 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.4 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 111.7 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 109.8 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.2 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-kafka-fdw` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.0 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_14` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.3 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_14` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.6 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_14` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 36.6 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_14` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.9 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_14` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 37.4 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_14` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 36.9 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.6 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.3 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 111.7 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 109.9 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.1 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-kafka-fdw` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.9 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/kafka_fdw" title="Repository" icon="github" subtitle="github.com/adjust/kafka_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="kafka_fdw-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg kafka_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install kafka_fdw;		# install via package name, for the active PG version

pig install kafka_fdw -v 18;   # install for PG 18
pig install kafka_fdw -v 17;   # install for PG 17
pig install kafka_fdw -v 16;   # install for PG 16
pig install kafka_fdw -v 15;   # install for PG 15
pig install kafka_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION kafka_fdw;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION kafka_fdw;
> CREATE SERVER kafka_server FOREIGN DATA WRAPPER kafka_fdw
>   OPTIONS (brokers 'localhost:9092');
> ```
>
> Source: [README](https://github.com/adjust/kafka_fdw)

`kafka_fdw` is a foreign data wrapper that exposes Kafka messages as PostgreSQL foreign tables. The upstream README explicitly warns that the project is not yet production ready.

### Server and Mapping

Define a foreign server with the Kafka broker list, then add a user mapping:

```sql
CREATE EXTENSION kafka_fdw;

CREATE SERVER kafka_server
FOREIGN DATA WRAPPER kafka_fdw
OPTIONS (brokers 'localhost:9092');

CREATE USER MAPPING FOR PUBLIC SERVER kafka_server;
```

## Foreign Tables

Kafka foreign tables must declare two metadata columns, one marked with `partition 'true'` and one marked with `offset 'true'`. The remaining columns describe the message payload.

### CSV Messages

```sql
CREATE FOREIGN TABLE kafka_test (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int,
    some_text text,
    some_date date,
    some_time timestamp
)
SERVER kafka_server
OPTIONS (
    format 'csv',
    topic 'contrib_regress',
    batch_size '30',
    buffer_delay '100'
);
```

For CSV, columns are mapped by position. Upstream notes that schema enforcement depends on the message writer, so strict parsing and junk-handling options matter when input quality is uncertain.

### JSON Messages

```sql
CREATE FOREIGN TABLE kafka_test_json (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int OPTIONS (json 'int_val'),
    some_text text OPTIONS (json 'text_val'),
    some_date date OPTIONS (json 'date_val'),
    some_time timestamp OPTIONS (json 'time_val')
)
SERVER kafka_server
OPTIONS (
    format 'json',
    topic 'contrib_regress_json',
    batch_size '30',
    buffer_delay '100'
);
```

For JSON, each column can map to an object key with the `json` option. The current implementation supports JSON objects, not top-level JSON arrays.

## Querying and Producing

The offset and partition columns are special, and the upstream README recommends specifying them in queries whenever possible:

```sql
SELECT * FROM kafka_test WHERE part = 0 AND offs > 1000 LIMIT 60;

SELECT *
FROM kafka_test
WHERE (part = 0 AND offs > 100)
   OR (part = 1 AND offs > 300)
   OR (part = 3 AND offs > 700);
```

Messages can also be produced with `INSERT` statements. If a partition value is supplied, it is used; otherwise Kafka's builtin partitioner chooses one:

```sql
INSERT INTO kafka_test(part, some_int, some_text)
VALUES
    (0, 5464565, 'some text goes into partition 0'),
    (NULL, 5464565, 'some text goes into partition selected by kafka');
```

## Error Handling

The default behavior is permissive:

- missing trailing columns are treated as `NULL`
- extra fields are ignored
- unparsable values still raise errors by default

Relevant table options and helper columns include:

- `strict 'true'` to reject column count mismatches
- `ignore_junk 'true'` to set malformed values to `NULL`
- columns marked `junk 'true'` to capture the original payload
- columns marked `junk_error 'true'` to capture parsing errors

## Build Notes

The extension uses `librdkafka` and the upstream build instructions are the standard:

```bash
make && make install
make installcheck
```

The test setup assumes Kafka on `localhost:9092` and ZooKeeper on `localhost:2181`.
