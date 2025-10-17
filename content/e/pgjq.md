---
title: "pgjq"
linkTitle: "pgjq"
description: "Use jq in Postgres"
weight: 4150
categories: ["Util"]
width: full
---

Use jq in Postgres

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4150** | {{< badge content="pgjq" link="https://github.com/Florents-Tselai/pgJQ" >}} | {{< ext "pgjq" "pgjq" >}} | `0.1.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjwt" >}} {{< ext "pg_protobuf" >}} {{< ext "jsquery" >}} {{< ext "sparql" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgjq" >}} | `0.1.0` | {{< badge content="18" color="red" alt="pgjq_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgjq_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgjq" >}} | `0.1.0` | {{< badge content="18" color="red" alt="postgresql-18-pgjq" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgjq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgjq_18" >}}     | {{< pkg "pgjq_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_17-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgjq_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_16-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgjq_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_15-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgjq_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_14-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgjq_18" >}}     | {{< pkg "pgjq_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_17-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgjq_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_16-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgjq_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_15-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgjq_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_14-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgjq_18" >}}     | {{< pkg "pgjq_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_17-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgjq_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_16-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgjq_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_15-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgjq_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_14-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgjq_18" >}}     | {{< pkg "pgjq_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_17-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgjq_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_16-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgjq_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_15-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgjq_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_14-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgjq" >}}     | {{< pkg "postgresql-17-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgjq" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgjq_17` | 0.1.0 | `el8.aarch64` | pigsty | 17.6 KiB | [pgjq_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjq_17` | 0.1.0 | `el8.x86_64` | pigsty | 17.7 KiB | [pgjq_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjq_17` | 0.1.0 | `el9.aarch64` | pigsty | 17.7 KiB | [pgjq_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjq_17` | 0.1.0 | `el9.x86_64` | pigsty | 18.0 KiB | [pgjq_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgjq` | 0.1.0 | `d12.x86_64` | pigsty | 19.3 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgjq` | 0.1.0 | `d12.aarch64` | pigsty | 19.0 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgjq` | 0.1.0 | `u22.aarch64` | pigsty | 19.8 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgjq` | 0.1.0 | `u22.x86_64` | pigsty | 20.1 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgjq` | 0.1.0 | `u24.x86_64` | pigsty | 19.4 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgjq` | 0.1.0 | `u24.aarch64` | pigsty | 18.9 KiB | [postgresql-17-pgjq_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-17-pgjq_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgjq_16` | 0.1.0 | `el8.x86_64` | pigsty | 17.8 KiB | [pgjq_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjq_16` | 0.1.0 | `el8.aarch64` | pigsty | 17.6 KiB | [pgjq_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjq_16` | 0.1.0 | `el9.aarch64` | pigsty | 17.7 KiB | [pgjq_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjq_16` | 0.1.0 | `el9.x86_64` | pigsty | 18.0 KiB | [pgjq_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pgjq` | 0.1.0 | `d12.aarch64` | pigsty | 18.9 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgjq` | 0.1.0 | `d12.x86_64` | pigsty | 19.3 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgjq` | 0.1.0 | `u22.x86_64` | pigsty | 20.1 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgjq` | 0.1.0 | `u22.aarch64` | pigsty | 19.8 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgjq` | 0.1.0 | `u24.aarch64` | pigsty | 18.9 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgjq` | 0.1.0 | `u24.x86_64` | pigsty | 19.4 KiB | [postgresql-16-pgjq_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-16-pgjq_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgjq_15` | 0.1.0 | `el8.x86_64` | pigsty | 17.7 KiB | [pgjq_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjq_15` | 0.1.0 | `el8.aarch64` | pigsty | 17.6 KiB | [pgjq_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjq_15` | 0.1.0 | `el9.x86_64` | pigsty | 18.0 KiB | [pgjq_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjq_15` | 0.1.0 | `el9.aarch64` | pigsty | 17.7 KiB | [pgjq_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgjq` | 0.1.0 | `d12.x86_64` | pigsty | 19.3 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgjq` | 0.1.0 | `d12.aarch64` | pigsty | 19.0 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgjq` | 0.1.0 | `u22.aarch64` | pigsty | 19.8 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgjq` | 0.1.0 | `u22.x86_64` | pigsty | 20.1 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgjq` | 0.1.0 | `u24.aarch64` | pigsty | 18.9 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgjq` | 0.1.0 | `u24.x86_64` | pigsty | 19.4 KiB | [postgresql-15-pgjq_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-15-pgjq_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgjq_14` | 0.1.0 | `el8.aarch64` | pigsty | 17.6 KiB | [pgjq_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjq_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjq_14` | 0.1.0 | `el8.x86_64` | pigsty | 17.7 KiB | [pgjq_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjq_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjq_14` | 0.1.0 | `el9.x86_64` | pigsty | 18.0 KiB | [pgjq_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjq_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjq_14` | 0.1.0 | `el9.aarch64` | pigsty | 17.7 KiB | [pgjq_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjq_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgjq` | 0.1.0 | `d12.aarch64` | pigsty | 19.0 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgjq` | 0.1.0 | `d12.x86_64` | pigsty | 19.2 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgjq` | 0.1.0 | `u22.x86_64` | pigsty | 20.0 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgjq` | 0.1.0 | `u22.aarch64` | pigsty | 19.8 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgjq` | 0.1.0 | `u24.aarch64` | pigsty | 18.9 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgjq` | 0.1.0 | `u24.x86_64` | pigsty | 19.3 KiB | [postgresql-14-pgjq_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjq/postgresql-14-pgjq_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/pgJQ" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/pgJQ" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgjq-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgjq; # get pgjq source code
pig build dep pgjq; # install build dependencies
pig build pkg pgjq; # build extension rpm or deb
pig build ext pgjq; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgjq; # install by extension name, for the current active PG version
pig ext install pgjq; # install via package alias, for the active PG version
pig ext install pgjq -v 18;   # install for PG 18
pig ext install pgjq -v 17;   # install for PG 17
pig ext install pgjq -v 16;   # install for PG 16
pig ext install pgjq -v 15;   # install for PG 15
pig ext install pgjq -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgjq;
```

