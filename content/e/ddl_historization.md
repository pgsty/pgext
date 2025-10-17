---
title: "ddl_historization"
linkTitle: "ddl_historization"
description: "Historize the ddl changes inside PostgreSQL database"
weight: 4310
categories: ["Util"]
width: full
---

Historize the ddl changes inside PostgreSQL database

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4310** | {{< badge content="ddl_historization" link="https://github.com/rodo/pg_ddl_historization" >}} | {{< ext "ddl_historization" "ddl_historization" >}} | `0.0.7` | {{< category "UTIL" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "schedoc" >}} |
|   **See Also**    | {{< ext "pg_readme" >}} {{< ext "data_historization" >}} {{< ext "table_version" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/ddl_historization" >}} | `0.0.7` | {{< badge content="18" color="red" alt="ddl_historization_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `ddl_historization_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/ddl_historization" >}} | `0.0.7` | {{< badge content="18" color="red" alt="postgresql-18-ddl-historization" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-ddl-historization` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "ddl_historization_18" >}}     | {{< pkg "ddl_historization_17" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddl_historization_16" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddl_historization_15" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "ddl_historization_14" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "ddl_historization_18" >}}     | {{< pkg "ddl_historization_17" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddl_historization_16" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddl_historization_15" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "ddl_historization_14" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "ddl_historization_18" >}}     | {{< pkg "ddl_historization_17" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddl_historization_16" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddl_historization_15" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "ddl_historization_14" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "ddl_historization_18" >}}     | {{< pkg "ddl_historization_17" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddl_historization_16" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddl_historization_15" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "ddl_historization_14" "0.0.7" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-ddl-historization" >}}     | {{< pkg "postgresql-17-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-ddl-historization" "0.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddl_historization_17` | 0.0.7 | `el8.x86_64` | pigsty | 16.0 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_17` | 0.0.7 | `el8.aarch64` | pigsty | 16.0 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_17` | 0.0.7 | `el9.aarch64` | pigsty | 15.8 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_17` | 0.0.7 | `el9.x86_64` | pigsty | 15.9 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-ddl-historization` | 0.0.7 | `d12.x86_64` | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddl-historization` | 0.0.7 | `d12.aarch64` | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddl-historization` | 0.0.7 | `u22.x86_64` | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddl-historization` | 0.0.7 | `u22.aarch64` | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddl-historization` | 0.0.7 | `u24.x86_64` | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddl-historization` | 0.0.7 | `u24.aarch64` | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddl_historization_16` | 0.0.7 | `el8.x86_64` | pigsty | 16.0 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_16` | 0.0.7 | `el8.aarch64` | pigsty | 16.0 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_16` | 0.0.7 | `el9.x86_64` | pigsty | 15.9 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_16` | 0.0.7 | `el9.aarch64` | pigsty | 15.8 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-ddl-historization` | 0.0.7 | `d12.x86_64` | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddl-historization` | 0.0.7 | `d12.aarch64` | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddl-historization` | 0.0.7 | `u22.aarch64` | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddl-historization` | 0.0.7 | `u22.x86_64` | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddl-historization` | 0.0.7 | `u24.x86_64` | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddl-historization` | 0.0.7 | `u24.aarch64` | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddl_historization_15` | 0.0.7 | `el8.x86_64` | pigsty | 16.0 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_15` | 0.0.7 | `el8.aarch64` | pigsty | 16.0 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_15` | 0.0.7 | `el9.x86_64` | pigsty | 15.9 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_15` | 0.0.7 | `el9.aarch64` | pigsty | 15.8 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-ddl-historization` | 0.0.7 | `d12.aarch64` | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddl-historization` | 0.0.7 | `d12.x86_64` | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddl-historization` | 0.0.7 | `u22.aarch64` | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddl-historization` | 0.0.7 | `u22.x86_64` | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddl-historization` | 0.0.7 | `u24.x86_64` | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddl-historization` | 0.0.7 | `u24.aarch64` | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddl_historization_14` | 0.0.7 | `el8.x86_64` | pigsty | 16.0 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_14` | 0.0.7 | `el8.aarch64` | pigsty | 16.0 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_14` | 0.0.7 | `el9.x86_64` | pigsty | 15.9 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_14` | 0.0.7 | `el9.aarch64` | pigsty | 15.8 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-ddl-historization` | 0.0.7 | `d12.x86_64` | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddl-historization` | 0.0.7 | `d12.aarch64` | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddl-historization` | 0.0.7 | `u22.x86_64` | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddl-historization` | 0.0.7 | `u22.aarch64` | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddl-historization` | 0.0.7 | `u24.x86_64` | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddl-historization` | 0.0.7 | `u24.aarch64` | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `ddl_historization_13` | 0.0.7 | `el8.aarch64` | pigsty | 16.0 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_13-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_13` | 0.0.7 | `el8.x86_64` | pigsty | 16.0 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_13-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_13` | 0.0.7 | `el9.aarch64` | pigsty | 15.8 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_13-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_13` | 0.0.7 | `el9.x86_64` | pigsty | 15.9 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_13-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-ddl-historization` | 0.0.7 | `d12.aarch64` | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-ddl-historization` | 0.0.7 | `d12.x86_64` | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-ddl-historization` | 0.0.7 | `u22.aarch64` | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-ddl-historization` | 0.0.7 | `u22.x86_64` | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-ddl-historization` | 0.0.7 | `u24.aarch64` | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-ddl-historization` | 0.0.7 | `u24.x86_64` | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rodo/pg_ddl_historization" title="Repository" icon="github" subtitle="github.com/rodo/pg_ddl_historization" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_ddl_historization-0.0.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get ddl_historization; # get ddl_historization source code
pig build dep ddl_historization; # install build dependencies
pig build pkg ddl_historization; # build extension rpm or deb
pig build ext ddl_historization; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install ddl_historization; # install by extension name, for the current active PG version
pig ext install ddl_historization; # install via package alias, for the active PG version
pig ext install ddl_historization -v 18;   # install for PG 18
pig ext install ddl_historization -v 17;   # install for PG 17
pig ext install ddl_historization -v 16;   # install for PG 16
pig ext install ddl_historization -v 15;   # install for PG 15
pig ext install ddl_historization -v 14;   # install for PG 14
pig ext install ddl_historization -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION ddl_historization;
```

