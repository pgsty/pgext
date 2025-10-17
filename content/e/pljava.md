---
title: "pljava"
linkTitle: "pljava"
description: "PL/Java procedural language"
weight: 3090
categories: ["Lang"]
width: full
---

PL/Java procedural language

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3090** | {{< badge content="pljava" link="https://github.com/tada/pljava" >}} | {{< ext "pljava" "pljava" >}} | `1.6.9` | {{< category "LANG" >}} | {{< license "BSD 3-Clause" >}} | {{< language "Java" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltclu" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pljava" >}} | `1.6.8` | {{< badge content="18" color="red" alt="pljava_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pljava_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pljava" >}} | `1.6.9` | {{< badge content="18" color="red" alt="postgresql-18-pljava" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pljava` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pljava_18" "1.6.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pljava_18-1.6.10-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pljava_17" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pljava_16" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pljava_15" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pljava_14" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pljava_18" "1.6.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pljava_18-1.6.10-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pljava_17" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pljava_16" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pljava_15" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pljava_14" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pljava_18" "1.6.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pljava_18-1.6.10-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pljava_17" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pljava_16" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pljava_15" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pljava_14" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pljava_18" "1.6.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pljava_18-1.6.10-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pljava_17" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pljava_16" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pljava_15" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pljava_14" "1.6.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pljava" >}}     | {{< pkg "postgresql-17-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pljava" "1.6.9" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_18` | 1.6.10 | `el8.x86_64` | pgdg | 927.7 KiB | [pljava_18-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pljava_18-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_18` | 1.6.10 | `el8.aarch64` | pgdg | 923.4 KiB | [pljava_18-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pljava_18-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_18` | 1.6.10 | `el9.aarch64` | pgdg | 914.2 KiB | [pljava_18-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pljava_18-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_18` | 1.6.10 | `el9.x86_64` | pgdg | 917.4 KiB | [pljava_18-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pljava_18-1.6.10-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_17` | 1.6.8 | `el8.aarch64` | pgdg | 910.0 KiB | [pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | 1.6.8 | `el8.x86_64` | pgdg | 914.1 KiB | [pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | 1.6.10 | `el8.aarch64` | pgdg | 923.4 KiB | [pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | 1.6.10 | `el8.x86_64` | pgdg | 927.5 KiB | [pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | 1.6.8 | `el9.aarch64` | pgdg | 892.1 KiB | [pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | 1.6.8 | `el9.x86_64` | pgdg | 895.2 KiB | [pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_17` | 1.6.10 | `el9.aarch64` | pgdg | 914.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | 1.6.10 | `el9.x86_64` | pgdg | 917.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pljava` | 1.6.9 | `d12.aarch64` | pgdg | 906.1 KiB | [postgresql-17-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-17-pljava` | 1.6.9 | `d12.x86_64` | pgdg | 911.3 KiB | [postgresql-17-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-17-pljava` | 1.6.9 | `u22.aarch64` | pgdg | 897.4 KiB | [postgresql-17-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pljava` | 1.6.9 | `u22.x86_64` | pgdg | 901.7 KiB | [postgresql-17-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pljava` | 1.6.9 | `u24.aarch64` | pgdg | 904.7 KiB | [postgresql-17-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pljava` | 1.6.9 | `u24.x86_64` | pgdg | 908.5 KiB | [postgresql-17-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_16` | 1.6.8 | `el8.x86_64` | pgdg | 913.9 KiB | [pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | 1.6.8 | `el8.aarch64` | pgdg | 910.1 KiB | [pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | 1.6.10 | `el8.aarch64` | pgdg | 923.7 KiB | [pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | 1.6.10 | `el8.x86_64` | pgdg | 927.8 KiB | [pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | 1.6.8 | `el9.aarch64` | pgdg | 892.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | 1.6.8 | `el9.x86_64` | pgdg | 895.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | 1.6.6 | `el9.x86_64` | pgdg | 891.8 KiB | [pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | 1.6.6 | `el9.aarch64` | pgdg | 888.2 KiB | [pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | 1.6.10 | `el9.x86_64` | pgdg | 917.4 KiB | [pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | 1.6.10 | `el9.aarch64` | pgdg | 914.3 KiB | [pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pljava` | 1.6.9 | `d12.aarch64` | pgdg | 906.3 KiB | [postgresql-16-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-16-pljava` | 1.6.9 | `d12.x86_64` | pgdg | 911.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-16-pljava` | 1.6.9 | `u22.aarch64` | pgdg | 897.7 KiB | [postgresql-16-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pljava` | 1.6.9 | `u22.x86_64` | pgdg | 901.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pljava` | 1.6.9 | `u24.x86_64` | pgdg | 908.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pljava` | 1.6.9 | `u24.aarch64` | pgdg | 904.8 KiB | [postgresql-16-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_15` | 1.6.8 | `el8.aarch64` | pgdg | 909.9 KiB | [pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | 1.6.8 | `el8.x86_64` | pgdg | 914.1 KiB | [pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | 1.6.10 | `el8.aarch64` | pgdg | 923.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | 1.6.10 | `el8.x86_64` | pgdg | 927.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | 1.6.8 | `el9.x86_64` | pgdg | 895.6 KiB | [pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | 1.6.8 | `el9.aarch64` | pgdg | 892.0 KiB | [pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | 1.6.6 | `el9.aarch64` | pgdg | 887.6 KiB | [pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | 1.6.6 | `el9.x86_64` | pgdg | 891.5 KiB | [pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | 1.6.10 | `el9.x86_64` | pgdg | 917.4 KiB | [pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | 1.6.10 | `el9.aarch64` | pgdg | 914.3 KiB | [pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pljava` | 1.6.9 | `d12.x86_64` | pgdg | 911.5 KiB | [postgresql-15-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-15-pljava` | 1.6.9 | `d12.aarch64` | pgdg | 906.0 KiB | [postgresql-15-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-15-pljava` | 1.6.9 | `u22.x86_64` | pgdg | 901.9 KiB | [postgresql-15-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pljava` | 1.6.9 | `u22.aarch64` | pgdg | 897.4 KiB | [postgresql-15-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pljava` | 1.6.9 | `u24.aarch64` | pgdg | 904.8 KiB | [postgresql-15-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pljava` | 1.6.9 | `u24.x86_64` | pgdg | 908.2 KiB | [postgresql-15-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_14` | 1.6.8 | `el8.x86_64` | pgdg | 914.2 KiB | [pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | 1.6.8 | `el8.aarch64` | pgdg | 910.0 KiB | [pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | 1.6.10 | `el8.x86_64` | pgdg | 927.6 KiB | [pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | 1.6.10 | `el8.aarch64` | pgdg | 923.4 KiB | [pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | 1.6.8 | `el9.x86_64` | pgdg | 895.1 KiB | [pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | 1.6.8 | `el9.aarch64` | pgdg | 892.3 KiB | [pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | 1.6.6 | `el9.x86_64` | pgdg | 891.4 KiB | [pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | 1.6.6 | `el9.aarch64` | pgdg | 887.8 KiB | [pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | 1.6.10 | `el9.aarch64` | pgdg | 914.2 KiB | [pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | 1.6.10 | `el9.x86_64` | pgdg | 917.3 KiB | [pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-pljava` | 1.6.9 | `d12.aarch64` | pgdg | 906.8 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-14-pljava` | 1.6.9 | `d12.x86_64` | pgdg | 910.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-14-pljava` | 1.6.9 | `u22.x86_64` | pgdg | 901.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljava` | 1.6.9 | `u22.aarch64` | pgdg | 897.2 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljava` | 1.6.9 | `u24.aarch64` | pgdg | 904.6 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pljava` | 1.6.9 | `u24.x86_64` | pgdg | 908.1 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pljava_13` | 1.6.8 | `el8.aarch64` | pgdg | 909.7 KiB | [pljava_13-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pljava_13-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_13` | 1.6.8 | `el8.x86_64` | pgdg | 912.8 KiB | [pljava_13-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pljava_13-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_13` | 1.6.10 | `el8.x86_64` | pgdg | 926.4 KiB | [pljava_13-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pljava_13-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_13` | 1.6.10 | `el8.aarch64` | pgdg | 923.3 KiB | [pljava_13-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pljava_13-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_13` | 1.6.8 | `el9.x86_64` | pgdg | 894.8 KiB | [pljava_13-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_13` | 1.6.8 | `el9.aarch64` | pgdg | 892.2 KiB | [pljava_13-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | 1.6.6 | `el9.x86_64` | pgdg | 891.4 KiB | [pljava_13-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_13` | 1.6.6 | `el9.aarch64` | pgdg | 888.0 KiB | [pljava_13-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | 1.6.4 | `el9.aarch64` | pgdg | 879.9 KiB | [pljava_13-1.6.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.4-1.rhel9.aarch64.rpm) |
| `pljava_13` | 1.6.4 | `el9.x86_64` | pgdg | 883.8 KiB | [pljava_13-1.6.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.4-1.rhel9.x86_64.rpm) |
| `pljava_13` | 1.6.10 | `el9.aarch64` | pgdg | 914.1 KiB | [pljava_13-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | 1.6.10 | `el9.x86_64` | pgdg | 917.0 KiB | [pljava_13-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-pljava` | 1.6.9 | `d12.x86_64` | pgdg | 910.9 KiB | [postgresql-13-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-13-pljava` | 1.6.9 | `d12.aarch64` | pgdg | 906.0 KiB | [postgresql-13-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-13-pljava` | 1.6.9 | `u22.x86_64` | pgdg | 901.3 KiB | [postgresql-13-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pljava` | 1.6.9 | `u22.aarch64` | pgdg | 897.2 KiB | [postgresql-13-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pljava` | 1.6.9 | `u24.x86_64` | pgdg | 908.1 KiB | [postgresql-13-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pljava` | 1.6.9 | `u24.aarch64` | pgdg | 904.8 KiB | [postgresql-13-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tada/pljava" title="Repository" icon="github" subtitle="github.com/tada/pljava" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pljava; # install by extension name, for the current active PG version
pig ext install pljava; # install via package alias, for the active PG version
pig ext install pljava -v 17;   # install for PG 17
pig ext install pljava -v 16;   # install for PG 16
pig ext install pljava -v 15;   # install for PG 15
pig ext install pljava -v 14;   # install for PG 14
pig ext install pljava -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pljava CASCADE SCHEMA sqlj;
```

