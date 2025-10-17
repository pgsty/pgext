---
title: "semver"
linkTitle: "semver"
description: "Semantic version data type"
weight: 3510
categories: ["Type"]
width: full
---

Semantic version data type

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3510** | {{< badge content="semver" link="https://github.com/theory/pg-semver" >}} | {{< ext "semver" "pg_semver" >}} | `0.40.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "ltree" >}} {{< ext "citext" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/semver" >}} | `0.32.1` | {{< badge content="18" color="red" alt="semver_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `semver_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/semver" >}} | `0.40.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-semver` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "semver_18" "0.40.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "semver_17" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "semver_16" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "semver_15" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "semver_14" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "semver_18" "0.40.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "semver_17" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "semver_16" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "semver_15" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "semver_14" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "semver_18" "0.40.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "semver_17" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "semver_16" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "semver_15" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "semver_14" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "semver_18" "0.40.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "semver_17" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "semver_16" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "semver_15" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "semver_14" "0.32.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-semver" "0.40.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_18` | 0.40.0 | `el8.x86_64` | pgdg | 27.9 KiB | [semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm) |
| `semver_18` | 0.40.0 | `el8.aarch64` | pgdg | 27.5 KiB | [semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm) |
| `semver_18` | 0.40.0 | `el9.aarch64` | pgdg | 26.5 KiB | [semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm) |
| `semver_18` | 0.40.0 | `el9.x86_64` | pgdg | 26.9 KiB | [semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.6 KiB | [postgresql-18-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.0 KiB | [postgresql-18-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-semver` | 0.40.0 | `u22.aarch64` | pgdg | 37.8 KiB | [postgresql-18-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-semver` | 0.40.0 | `u22.x86_64` | pgdg | 38.2 KiB | [postgresql-18-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.3 KiB | [postgresql-18-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.1 KiB | [postgresql-18-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_17` | 0.32.1 | `el8.x86_64` | pgdg | 27.1 KiB | [semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_17` | 0.32.1 | `el8.aarch64` | pgdg | 26.8 KiB | [semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_17` | 0.32.1 | `el9.x86_64` | pgdg | 26.5 KiB | [semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_17` | 0.32.1 | `el9.aarch64` | pgdg | 26.1 KiB | [semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.0 KiB | [postgresql-17-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.6 KiB | [postgresql-17-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-semver` | 0.40.0 | `u22.x86_64` | pgdg | 38.9 KiB | [postgresql-17-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-semver` | 0.40.0 | `u22.aarch64` | pgdg | 38.4 KiB | [postgresql-17-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.1 KiB | [postgresql-17-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.2 KiB | [postgresql-17-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_16` | 0.32.1 | `el8.x86_64` | pgdg | 27.1 KiB | [semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_16` | 0.32.1 | `el8.aarch64` | pgdg | 26.8 KiB | [semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_16` | 0.32.1 | `el9.aarch64` | pgdg | 25.8 KiB | [semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_16` | 0.32.1 | `el9.x86_64` | pgdg | 26.3 KiB | [semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.6 KiB | [postgresql-16-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.0 KiB | [postgresql-16-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-semver` | 0.40.0 | `u22.aarch64` | pgdg | 38.5 KiB | [postgresql-16-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-semver` | 0.40.0 | `u22.x86_64` | pgdg | 38.9 KiB | [postgresql-16-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.2 KiB | [postgresql-16-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.1 KiB | [postgresql-16-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_15` | 0.32.1 | `el8.x86_64` | pgdg | 27.3 KiB | [semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_15` | 0.32.1 | `el8.aarch64` | pgdg | 26.9 KiB | [semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_15` | 0.32.0 | `el8.aarch64` | pgdg | 41.3 KiB | [semver_15-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_15` | 0.32.0 | `el8.x86_64` | pgdg | 41.9 KiB | [semver_15-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_15` | 0.31.2 | `el8.x86_64` | pgdg | 40.1 KiB | [semver_15-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_15` | 0.32.1 | `el9.aarch64` | pgdg | 26.1 KiB | [semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_15` | 0.32.1 | `el9.x86_64` | pgdg | 26.7 KiB | [semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_15` | 0.32.0 | `el9.aarch64` | pgdg | 41.6 KiB | [semver_15-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_15` | 0.32.0 | `el9.x86_64` | pgdg | 42.4 KiB | [semver_15-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_15` | 0.31.2 | `el9.x86_64` | pgdg | 40.9 KiB | [semver_15-0.31.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.31.2-1.rhel9.x86_64.rpm) |
| `postgresql-15-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.0 KiB | [postgresql-15-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-semver` | 0.40.0 | `u22.x86_64` | pgdg | 39.1 KiB | [postgresql-15-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-semver` | 0.40.0 | `u22.aarch64` | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.1 KiB | [postgresql-15-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_14` | 0.32.1 | `el8.x86_64` | pgdg | 27.3 KiB | [semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_14` | 0.32.1 | `el8.aarch64` | pgdg | 26.9 KiB | [semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_14` | 0.32.0 | `el8.aarch64` | pgdg | 41.4 KiB | [semver_14-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_14` | 0.32.0 | `el8.x86_64` | pgdg | 41.9 KiB | [semver_14-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_14` | 0.31.2 | `el8.x86_64` | pgdg | 40.4 KiB | [semver_14-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_14` | 0.31.1 | `el8.x86_64` | pgdg | 39.8 KiB | [semver_14-0.31.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.1-2.rhel8.x86_64.rpm) |
| `semver_14` | 0.32.1 | `el9.x86_64` | pgdg | 26.7 KiB | [semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_14` | 0.32.1 | `el9.aarch64` | pgdg | 26.0 KiB | [semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_14` | 0.32.0 | `el9.aarch64` | pgdg | 41.6 KiB | [semver_14-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_14` | 0.32.0 | `el9.x86_64` | pgdg | 42.4 KiB | [semver_14-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.0-1.rhel9.x86_64.rpm) |
| `postgresql-14-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.0 KiB | [postgresql-14-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-semver` | 0.40.0 | `u22.x86_64` | pgdg | 39.0 KiB | [postgresql-14-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-semver` | 0.40.0 | `u22.aarch64` | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.1 KiB | [postgresql-14-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `semver_13` | 0.32.1 | `el8.aarch64` | pgdg | 26.9 KiB | [semver_13-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/semver_13-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_13` | 0.32.1 | `el8.x86_64` | pgdg | 27.2 KiB | [semver_13-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_13` | 0.32.0 | `el8.x86_64` | pgdg | 41.7 KiB | [semver_13-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_13` | 0.32.0 | `el8.aarch64` | pgdg | 41.2 KiB | [semver_13-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/semver_13-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_13` | 0.31.2 | `el8.x86_64` | pgdg | 40.4 KiB | [semver_13-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_13` | 0.31.1 | `el8.x86_64` | pgdg | 39.3 KiB | [semver_13-0.31.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.31.1-1.rhel8.x86_64.rpm) |
| `semver_13` | 0.32.1 | `el9.aarch64` | pgdg | 26.1 KiB | [semver_13-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/semver_13-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_13` | 0.32.1 | `el9.x86_64` | pgdg | 26.7 KiB | [semver_13-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/semver_13-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_13` | 0.32.0 | `el9.aarch64` | pgdg | 41.5 KiB | [semver_13-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/semver_13-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_13` | 0.32.0 | `el9.x86_64` | pgdg | 42.2 KiB | [semver_13-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/semver_13-0.32.0-1.rhel9.x86_64.rpm) |
| `postgresql-13-semver` | 0.40.0 | `d12.x86_64` | pgdg | 38.5 KiB | [postgresql-13-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-semver` | 0.40.0 | `d12.aarch64` | pgdg | 38.3 KiB | [postgresql-13-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-semver` | 0.40.0 | `u22.aarch64` | pgdg | 38.6 KiB | [postgresql-13-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-semver` | 0.40.0 | `u22.x86_64` | pgdg | 39.0 KiB | [postgresql-13-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-semver` | 0.40.0 | `u24.x86_64` | pgdg | 38.2 KiB | [postgresql-13-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-semver` | 0.40.0 | `u24.aarch64` | pgdg | 38.2 KiB | [postgresql-13-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pg-semver" title="Repository" icon="github" subtitle="github.com/theory/pg-semver" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg-semver-0.40.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get semver; # get semver source code
pig build dep semver; # install build dependencies
pig build pkg semver; # build extension rpm or deb
pig build ext semver; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install semver; # install by extension name, for the current active PG version
pig ext install pg_semver; # install via package alias, for the active PG version
pig ext install semver -v 18;   # install for PG 18
pig ext install semver -v 17;   # install for PG 17
pig ext install semver -v 16;   # install for PG 16
pig ext install semver -v 15;   # install for PG 15
pig ext install semver -v 14;   # install for PG 14
pig ext install semver -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION semver;
```

