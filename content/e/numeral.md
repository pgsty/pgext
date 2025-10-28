---
title: "numeral"
linkTitle: "numeral"
description: "numeral datatypes extension"
weight: 3710
categories: ["TYPE"]
width: full
---

numeral datatypes extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3710** | {{< badge content="numeral" link="https://github.com/df7cb/postgresql-numeral" >}} | {{< ext "numeral" >}} | `1.3` | {{< category "TYPE" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "currency" >}} {{< ext "pgmp" >}} {{< ext "unit" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/numeral" >}} | `1.3` | {{< bg "18" "numeral_18*" "green" >}} {{< bg "17" "numeral_17*" "green" >}} {{< bg "16" "numeral_16*" "green" >}} {{< bg "15" "numeral_15*" "green" >}} {{< bg "14" "numeral_14*" "green" >}} {{< bg "13" "numeral_13*" "green" >}} | `numeral_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/numeral" >}} | `1.3` | {{< bg "18" "postgresql-18-numeral" "green" >}} {{< bg "17" "postgresql-17-numeral" "green" >}} {{< bg "16" "postgresql-16-numeral" "green" >}} {{< bg "15" "postgresql-15-numeral" "green" >}} {{< bg "14" "postgresql-14-numeral" "green" >}} {{< bg "13" "postgresql-13-numeral" "green" >}} | `postgresql-$v-numeral` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.3" "numeral_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.3" "numeral_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.3" "numeral_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.3" "numeral_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "numeral_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "numeral_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "numeral_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "numeral_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-13-numeral : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_18` | 1.3 | `el8.x86_64` | pigsty | 30.5 KiB | [numeral_18-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_18-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_18` | 1.3 | `el8.aarch64` | pigsty | 29.2 KiB | [numeral_18-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_18-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_18` | 1.3 | `el9.x86_64` | pigsty | 30.6 KiB | [numeral_18-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_18-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_18` | 1.3 | `el9.aarch64` | pigsty | 30.9 KiB | [numeral_18-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_18-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-18-numeral` | 1.3 | `d12.x86_64` | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-18-numeral` | 1.3 | `d12.aarch64` | pgdg | 72.2 KiB | [postgresql-18-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-18-numeral` | 1.3 | `d13.x86_64` | pgdg | 75.0 KiB | [postgresql-18-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-18-numeral` | 1.3 | `d13.aarch64` | pgdg | 73.3 KiB | [postgresql-18-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-18-numeral` | 1.3 | `u22.x86_64` | pgdg | 74.5 KiB | [postgresql-18-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-18-numeral` | 1.3 | `u22.aarch64` | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-18-numeral` | 1.3 | `u24.x86_64` | pgdg | 73.8 KiB | [postgresql-18-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-18-numeral` | 1.3 | `u24.aarch64` | pgdg | 73.2 KiB | [postgresql-18-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_17` | 1.3 | `el8.x86_64` | pigsty | 30.5 KiB | [numeral_17-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_17-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_17` | 1.3 | `el8.aarch64` | pigsty | 29.2 KiB | [numeral_17-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_17-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_17` | 1.3 | `el9.x86_64` | pigsty | 30.6 KiB | [numeral_17-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_17-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_17` | 1.3 | `el9.aarch64` | pigsty | 31.2 KiB | [numeral_17-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_17-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-numeral` | 1.3 | `d12.x86_64` | pgdg | 74.1 KiB | [postgresql-17-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-17-numeral` | 1.3 | `d12.aarch64` | pgdg | 72.2 KiB | [postgresql-17-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-17-numeral` | 1.3 | `d13.x86_64` | pgdg | 75.0 KiB | [postgresql-17-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-17-numeral` | 1.3 | `d13.aarch64` | pgdg | 73.3 KiB | [postgresql-17-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-17-numeral` | 1.3 | `u22.x86_64` | pgdg | 77.4 KiB | [postgresql-17-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-17-numeral` | 1.3 | `u22.aarch64` | pgdg | 77.1 KiB | [postgresql-17-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-17-numeral` | 1.3 | `u24.x86_64` | pgdg | 73.8 KiB | [postgresql-17-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-17-numeral` | 1.3 | `u24.aarch64` | pgdg | 73.2 KiB | [postgresql-17-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_16` | 1.3 | `el8.x86_64` | pigsty | 30.5 KiB | [numeral_16-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_16-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_16` | 1.3 | `el8.aarch64` | pigsty | 29.2 KiB | [numeral_16-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_16-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_16` | 1.3 | `el9.x86_64` | pigsty | 30.6 KiB | [numeral_16-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_16-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_16` | 1.3 | `el9.aarch64` | pigsty | 31.2 KiB | [numeral_16-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_16-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-numeral` | 1.3 | `d12.x86_64` | pgdg | 74.1 KiB | [postgresql-16-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-16-numeral` | 1.3 | `d12.aarch64` | pgdg | 72.1 KiB | [postgresql-16-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-16-numeral` | 1.3 | `d13.x86_64` | pgdg | 75.0 KiB | [postgresql-16-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-16-numeral` | 1.3 | `d13.aarch64` | pgdg | 73.3 KiB | [postgresql-16-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-16-numeral` | 1.3 | `u22.x86_64` | pgdg | 77.3 KiB | [postgresql-16-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-16-numeral` | 1.3 | `u22.aarch64` | pgdg | 77.1 KiB | [postgresql-16-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-16-numeral` | 1.3 | `u24.x86_64` | pgdg | 73.8 KiB | [postgresql-16-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-16-numeral` | 1.3 | `u24.aarch64` | pgdg | 73.2 KiB | [postgresql-16-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_15` | 1.3 | `el8.x86_64` | pigsty | 32.5 KiB | [numeral_15-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_15-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_15` | 1.3 | `el8.aarch64` | pigsty | 30.9 KiB | [numeral_15-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_15-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_15` | 1.3 | `el9.x86_64` | pigsty | 34.2 KiB | [numeral_15-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_15-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_15` | 1.3 | `el9.aarch64` | pigsty | 34.0 KiB | [numeral_15-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_15-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-numeral` | 1.3 | `d12.x86_64` | pgdg | 76.0 KiB | [postgresql-15-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-15-numeral` | 1.3 | `d12.aarch64` | pgdg | 74.0 KiB | [postgresql-15-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-15-numeral` | 1.3 | `d13.x86_64` | pgdg | 77.4 KiB | [postgresql-15-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-15-numeral` | 1.3 | `d13.aarch64` | pgdg | 75.1 KiB | [postgresql-15-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-15-numeral` | 1.3 | `u22.x86_64` | pgdg | 79.9 KiB | [postgresql-15-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-15-numeral` | 1.3 | `u22.aarch64` | pgdg | 78.7 KiB | [postgresql-15-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-15-numeral` | 1.3 | `u24.x86_64` | pgdg | 75.8 KiB | [postgresql-15-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-15-numeral` | 1.3 | `u24.aarch64` | pgdg | 74.6 KiB | [postgresql-15-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_14` | 1.3 | `el8.x86_64` | pigsty | 32.5 KiB | [numeral_14-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_14-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_14` | 1.3 | `el8.aarch64` | pigsty | 30.9 KiB | [numeral_14-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_14-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_14` | 1.3 | `el9.x86_64` | pigsty | 34.2 KiB | [numeral_14-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_14-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_14` | 1.3 | `el9.aarch64` | pigsty | 34.0 KiB | [numeral_14-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_14-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-numeral` | 1.3 | `d12.x86_64` | pgdg | 76.0 KiB | [postgresql-14-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-14-numeral` | 1.3 | `d12.aarch64` | pgdg | 74.0 KiB | [postgresql-14-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-14-numeral` | 1.3 | `d13.x86_64` | pgdg | 77.3 KiB | [postgresql-14-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-14-numeral` | 1.3 | `d13.aarch64` | pgdg | 75.1 KiB | [postgresql-14-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-14-numeral` | 1.3 | `u22.x86_64` | pgdg | 79.8 KiB | [postgresql-14-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-14-numeral` | 1.3 | `u22.aarch64` | pgdg | 78.7 KiB | [postgresql-14-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-14-numeral` | 1.3 | `u24.x86_64` | pgdg | 75.7 KiB | [postgresql-14-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-14-numeral` | 1.3 | `u24.aarch64` | pgdg | 74.6 KiB | [postgresql-14-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `numeral_13` | 1.3 | `el8.x86_64` | pigsty | 32.4 KiB | [numeral_13-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/numeral_13-1.3-1PIGSTY.el8.x86_64.rpm) |
| `numeral_13` | 1.3 | `el8.aarch64` | pigsty | 30.9 KiB | [numeral_13-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/numeral_13-1.3-1PIGSTY.el8.aarch64.rpm) |
| `numeral_13` | 1.3 | `el9.x86_64` | pigsty | 33.8 KiB | [numeral_13-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/numeral_13-1.3-1PIGSTY.el9.x86_64.rpm) |
| `numeral_13` | 1.3 | `el9.aarch64` | pigsty | 34.0 KiB | [numeral_13-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/numeral_13-1.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-numeral` | 1.3 | `d12.x86_64` | pgdg | 76.0 KiB | [postgresql-13-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-13-numeral` | 1.3 | `d12.aarch64` | pgdg | 74.0 KiB | [postgresql-13-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-13-numeral` | 1.3 | `d13.x86_64` | pgdg | 77.2 KiB | [postgresql-13-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-13-numeral` | 1.3 | `d13.aarch64` | pgdg | 75.0 KiB | [postgresql-13-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-13-numeral` | 1.3 | `u22.x86_64` | pgdg | 79.8 KiB | [postgresql-13-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-13-numeral` | 1.3 | `u22.aarch64` | pgdg | 78.6 KiB | [postgresql-13-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-13-numeral` | 1.3 | `u24.x86_64` | pgdg | 75.7 KiB | [postgresql-13-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-13-numeral` | 1.3 | `u24.aarch64` | pgdg | 74.4 KiB | [postgresql-13-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-13-numeral_1.3-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/postgresql-numeral" title="Repository" icon="github" subtitle="github.com/df7cb/postgresql-numeral" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql-numeral-1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get numeral; # get numeral source code
pig build dep numeral; # install build dependencies
pig build pkg numeral; # build extension rpm or deb
pig build ext numeral; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install numeral; # install by extension name, for the current active PG version
pig ext install numeral; # install via package alias, for the active PG version
pig ext install numeral -v 18;   # install for PG 18
pig ext install numeral -v 17;   # install for PG 17
pig ext install numeral -v 16;   # install for PG 16
pig ext install numeral -v 15;   # install for PG 15
pig ext install numeral -v 14;   # install for PG 14
pig ext install numeral -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION numeral;
```

