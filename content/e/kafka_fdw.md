---
title: "kafka_fdw"
linkTitle: "kafka_fdw"
description: "kafka Foreign Data Wrapper for CSV formatted messages"
weight: 8730
categories: ["FDW"]
width: full
---

kafka Foreign Data Wrapper for CSV formatted messages


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8730** | {{< badge content="kafka_fdw" link="https://github.com/adjust/kafka_fdw" >}} | {{< ext "kafka_fdw" >}} | `0.0.3` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "mongo_fdw" >}} {{< ext "redis_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "redis" >}} {{< ext "hdfs_fdw" >}} {{< ext "wal2json" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/kafka_fdw" >}} | `0.0.3` | {{< bg "18" "kafka_fdw_18" "green" >}} {{< bg "17" "kafka_fdw_17" "green" >}} {{< bg "16" "kafka_fdw_16" "green" >}} {{< bg "15" "kafka_fdw_15" "green" >}} {{< bg "14" "kafka_fdw_14" "green" >}} {{< bg "13" "kafka_fdw_13" "green" >}} | `kafka_fdw_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/kafka_fdw" >}} | `0.0.3` | {{< bg "18" "postgresql-18-kafka-fdw" "green" >}} {{< bg "17" "postgresql-17-kafka-fdw" "green" >}} {{< bg "16" "postgresql-16-kafka-fdw" "green" >}} {{< bg "15" "postgresql-15-kafka-fdw" "green" >}} {{< bg "14" "postgresql-14-kafka-fdw" "green" >}} {{< bg "13" "postgresql-13-kafka-fdw" "green" >}} | `postgresql-$v-kafka-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "kafka_fdw_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-kafka-fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-kafka-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-kafka-fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-kafka-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-kafka-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_18` | 0.0.3 | `el8.x86_64` | pigsty | 34.6 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_18` | 0.0.3 | `el8.aarch64` | pigsty | 33.5 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_18` | 0.0.3 | `el9.x86_64` | pigsty | 33.5 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_18` | 0.0.3 | `el9.aarch64` | pigsty | 33.0 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_18` | 0.0.3 | `el10.x86_64` | pigsty | 34.6 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_18-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_18` | 0.0.3 | `el10.aarch64` | pigsty | 33.8 KiB | [kafka_fdw_18-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_18-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 79.0 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 76.8 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 84.9 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 83.5 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 82.4 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 81.2 KiB | [postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-18-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_17` | 0.0.3 | `el8.x86_64` | pigsty | 34.6 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_17` | 0.0.3 | `el8.aarch64` | pigsty | 33.4 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_17` | 0.0.3 | `el9.x86_64` | pigsty | 33.5 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_17` | 0.0.3 | `el9.aarch64` | pigsty | 33.0 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_17` | 0.0.3 | `el10.x86_64` | pigsty | 34.5 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_17-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_17` | 0.0.3 | `el10.aarch64` | pigsty | 33.9 KiB | [kafka_fdw_17-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_17-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 78.8 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 76.6 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 106.0 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 104.5 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 82.2 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 80.9 KiB | [postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-17-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_16` | 0.0.3 | `el8.x86_64` | pigsty | 35.1 KiB | [kafka_fdw_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el8.x86_64` | pigsty | 37.0 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el8.aarch64` | pigsty | 33.2 KiB | [kafka_fdw_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el8.aarch64` | pigsty | 35.4 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el9.x86_64` | pigsty | 36.0 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el9.x86_64` | pigsty | 35.3 KiB | [kafka_fdw_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el9.aarch64` | pigsty | 34.4 KiB | [kafka_fdw_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el9.aarch64` | pigsty | 35.2 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el10.x86_64` | pigsty | 36.8 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_16-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_16` | 0.0.3 | `el10.aarch64` | pigsty | 36.1 KiB | [kafka_fdw_16-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_16-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 84.4 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 81.9 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 112.0 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 110.1 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 88.1 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 86.9 KiB | [postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-16-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_15` | 0.0.3 | `el8.x86_64` | pigsty | 35.4 KiB | [kafka_fdw_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el8.x86_64` | pigsty | 37.3 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el8.aarch64` | pigsty | 33.4 KiB | [kafka_fdw_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el8.aarch64` | pigsty | 35.6 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el9.x86_64` | pigsty | 36.1 KiB | [kafka_fdw_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el9.x86_64` | pigsty | 36.6 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el9.aarch64` | pigsty | 35.9 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el9.aarch64` | pigsty | 35.1 KiB | [kafka_fdw_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el10.x86_64` | pigsty | 37.4 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_15-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_15` | 0.0.3 | `el10.aarch64` | pigsty | 36.8 KiB | [kafka_fdw_15-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_15-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 84.3 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 81.8 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 111.7 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 109.8 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 88.2 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 87.0 KiB | [postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-15-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_14` | 0.0.3 | `el8.x86_64` | pigsty | 37.3 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el8.x86_64` | pigsty | 35.4 KiB | [kafka_fdw_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el8.aarch64` | pigsty | 35.6 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el8.aarch64` | pigsty | 33.4 KiB | [kafka_fdw_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el9.x86_64` | pigsty | 36.6 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el9.x86_64` | pigsty | 36.1 KiB | [kafka_fdw_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el9.aarch64` | pigsty | 35.0 KiB | [kafka_fdw_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el9.aarch64` | pigsty | 35.9 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el10.x86_64` | pigsty | 37.4 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_14-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_14` | 0.0.3 | `el10.aarch64` | pigsty | 36.9 KiB | [kafka_fdw_14-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_14-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 84.3 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 81.6 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 111.7 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 109.9 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 88.1 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 86.9 KiB | [postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-14-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `kafka_fdw_13` | 0.0.3 | `el8.x86_64` | pigsty | 35.0 KiB | [kafka_fdw_13-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_13-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el8.x86_64` | pigsty | 36.8 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/kafka_fdw_13-0.0.3-2PIGSTY.el8.x86_64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el8.aarch64` | pigsty | 35.6 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_13-0.0.3-2PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el8.aarch64` | pigsty | 33.3 KiB | [kafka_fdw_13-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/kafka_fdw_13-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el9.x86_64` | pigsty | 36.0 KiB | [kafka_fdw_13-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_13-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el9.x86_64` | pigsty | 36.6 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/kafka_fdw_13-0.0.3-2PIGSTY.el9.x86_64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el9.aarch64` | pigsty | 35.0 KiB | [kafka_fdw_13-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_13-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el9.aarch64` | pigsty | 35.8 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/kafka_fdw_13-0.0.3-2PIGSTY.el9.aarch64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el10.x86_64` | pigsty | 37.2 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/kafka_fdw_13-0.0.3-2PIGSTY.el10.x86_64.rpm) |
| `kafka_fdw_13` | 0.0.3 | `el10.aarch64` | pigsty | 36.7 KiB | [kafka_fdw_13-0.0.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/kafka_fdw_13-0.0.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `d12.x86_64` | pigsty | 84.1 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `d12.aarch64` | pigsty | 81.5 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `u22.x86_64` | pigsty | 110.4 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `u22.aarch64` | pigsty | 109.1 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `u24.x86_64` | pigsty | 87.9 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-kafka-fdw` | 0.0.3 | `u24.aarch64` | pigsty | 86.7 KiB | [postgresql-13-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/k/kafka-fdw/postgresql-13-kafka-fdw_0.0.3-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/kafka_fdw" title="Repository" icon="github" subtitle="github.com/adjust/kafka_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="kafka_fdw-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get kafka_fdw; # get kafka_fdw source code
pig build dep kafka_fdw; # install build dependencies
pig build pkg kafka_fdw; # build extension rpm or deb
pig build ext kafka_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install kafka_fdw; # install by extension name, for the current active PG version
pig ext install kafka_fdw; # install via package alias, for the active PG version
pig ext install kafka_fdw -v 18;   # install for PG 18
pig ext install kafka_fdw -v 17;   # install for PG 17
pig ext install kafka_fdw -v 16;   # install for PG 16
pig ext install kafka_fdw -v 15;   # install for PG 15
pig ext install kafka_fdw -v 14;   # install for PG 14
pig ext install kafka_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION kafka_fdw;
```

