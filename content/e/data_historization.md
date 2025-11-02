---
title: "data_historization"
linkTitle: "data_historization"
description: "PLPGSQL Script to historize data in partitionned table"
weight: 4320
categories: ["UTIL"]
width: full
---

PLPGSQL Script to historize data in partitionned table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4320** | {{< badge content="data_historization" link="https://github.com/rodo/postgresql-data-historization" >}} | {{< ext "data_historization" >}} | `1.1.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "ddl_historization" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/data_historization" >}} | `1.1.0` | {{< bg "18" "data_historization_18" "green" >}} {{< bg "17" "data_historization_17" "green" >}} {{< bg "16" "data_historization_16" "green" >}} {{< bg "15" "data_historization_15" "green" >}} {{< bg "14" "data_historization_14" "green" >}} {{< bg "13" "data_historization_13" "green" >}} | `data_historization_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/data_historization" >}} | `1.1.0` | {{< bg "18" "postgresql-18-data-historization" "green" >}} {{< bg "17" "postgresql-17-data-historization" "green" >}} {{< bg "16" "postgresql-16-data-historization" "green" >}} {{< bg "15" "postgresql-15-data-historization" "green" >}} {{< bg "14" "postgresql-14-data-historization" "green" >}} {{< bg "13" "postgresql-13-data-historization" "green" >}} | `postgresql-$v-data-historization` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "data_historization_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-data-historization : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-data-historization : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-data-historization : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-data-historization : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.0" "postgresql-17-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-data-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-13-data-historization : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_18` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_18` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_18` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_18` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_18` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_18` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_17` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_17` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_17` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_17` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_17` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_17` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-data-historization` | 1.1.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-data-historization` | 1.1.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-data-historization` | 1.1.0 | `u22.x86_64` | pigsty | 5.8 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-data-historization` | 1.1.0 | `u22.aarch64` | pigsty | 5.8 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-data-historization` | 1.1.0 | `u24.x86_64` | pigsty | 5.8 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-data-historization` | 1.1.0 | `u24.aarch64` | pigsty | 5.8 KiB | [postgresql-17-data-historization_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-17-data-historization_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_16` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_16` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_16` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_16` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_16` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_16` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-data-historization` | 1.1.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-data-historization` | 1.1.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-data-historization` | 1.1.0 | `u22.x86_64` | pigsty | 5.8 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-data-historization` | 1.1.0 | `u22.aarch64` | pigsty | 5.8 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-data-historization` | 1.1.0 | `u24.x86_64` | pigsty | 5.8 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-data-historization` | 1.1.0 | `u24.aarch64` | pigsty | 5.8 KiB | [postgresql-16-data-historization_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-16-data-historization_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_15` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_15` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_15` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_15` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_15` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_15` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-data-historization` | 1.1.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-data-historization` | 1.1.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-data-historization` | 1.1.0 | `u22.x86_64` | pigsty | 5.8 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-data-historization` | 1.1.0 | `u22.aarch64` | pigsty | 5.8 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-data-historization` | 1.1.0 | `u24.x86_64` | pigsty | 5.8 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-data-historization` | 1.1.0 | `u24.aarch64` | pigsty | 5.8 KiB | [postgresql-15-data-historization_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-15-data-historization_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_14` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_14` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_14` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_14` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_14` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_14` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-data-historization` | 1.1.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-data-historization` | 1.1.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-data-historization` | 1.1.0 | `u22.x86_64` | pigsty | 5.8 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-data-historization` | 1.1.0 | `u22.aarch64` | pigsty | 5.8 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-data-historization` | 1.1.0 | `u24.x86_64` | pigsty | 5.8 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-data-historization` | 1.1.0 | `u24.aarch64` | pigsty | 5.8 KiB | [postgresql-14-data-historization_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-14-data-historization_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `data_historization_13` | 1.1.0 | `el8.x86_64` | pigsty | 15.1 KiB | [data_historization_13-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/data_historization_13-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `data_historization_13` | 1.1.0 | `el8.aarch64` | pigsty | 15.1 KiB | [data_historization_13-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/data_historization_13-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `data_historization_13` | 1.1.0 | `el9.x86_64` | pigsty | 14.9 KiB | [data_historization_13-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/data_historization_13-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `data_historization_13` | 1.1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [data_historization_13-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/data_historization_13-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `data_historization_13` | 1.1.0 | `el10.x86_64` | pigsty | 15.0 KiB | [data_historization_13-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/data_historization_13-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `data_historization_13` | 1.1.0 | `el10.aarch64` | pigsty | 14.9 KiB | [data_historization_13-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/data_historization_13-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-data-historization` | 1.1.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-data-historization` | 1.1.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-data-historization` | 1.1.0 | `u22.x86_64` | pigsty | 5.8 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-data-historization` | 1.1.0 | `u22.aarch64` | pigsty | 5.8 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-data-historization` | 1.1.0 | `u24.x86_64` | pigsty | 5.8 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-data-historization` | 1.1.0 | `u24.aarch64` | pigsty | 5.8 KiB | [postgresql-13-data-historization_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/data-historization/postgresql-13-data-historization_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rodo/postgresql-data-historization" title="Repository" icon="github" subtitle="github.com/rodo/postgresql-data-historization" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql-data-historization-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get data_historization; # get data_historization source code
pig build dep data_historization; # install build dependencies
pig build pkg data_historization; # build extension rpm or deb
pig build ext data_historization; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install data_historization; # install by extension name, for the current active PG version
pig ext install data_historization; # install via package alias, for the active PG version
pig ext install data_historization -v 18;   # install for PG 18
pig ext install data_historization -v 17;   # install for PG 17
pig ext install data_historization -v 16;   # install for PG 16
pig ext install data_historization -v 15;   # install for PG 15
pig ext install data_historization -v 14;   # install for PG 14
pig ext install data_historization -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION data_historization;
```

