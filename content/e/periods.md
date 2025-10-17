---
title: "periods"
linkTitle: "periods"
description: "Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING"
weight: 1030
categories: ["Time"]
width: full
---

Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1030** | {{< badge content="periods" link="https://github.com/xocolatl/periods" >}} | {{< ext "periods" "periods" >}} | `1.2.3` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/periods" >}} | `1.2.2` | {{< badge content="18" color="red" alt="periods_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `periods_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/periods" >}} | `1.2.3` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-periods` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "periods_18" >}}     | {{< pkg "periods_17" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/periods_17-1.2.2-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "periods_16" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/periods_16-1.2.2-1.rhel8.1.x86_64.rpm" >}} | {{< pkg "periods_15" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/periods_15-1.2.2-1.rhel8.x86_64.rpm" >}} | {{< pkg "periods_14" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/periods_14-1.2.2-1.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "periods_18" >}}     | {{< pkg "periods_17" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/periods_17-1.2.2-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "periods_16" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/periods_16-1.2.2-1.rhel8.1.aarch64.rpm" >}} | {{< pkg "periods_15" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/periods_15-1.2.2-1.rhel8.aarch64.rpm" >}} | {{< pkg "periods_14" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/periods_14-1.2.2-1.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "periods_18" >}}     | {{< pkg "periods_17" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/periods_17-1.2.2-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "periods_16" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/periods_16-1.2.2-1.rhel9.1.x86_64.rpm" >}} | {{< pkg "periods_15" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/periods_15-1.2.2-1.rhel9.x86_64.rpm" >}} | {{< pkg "periods_14" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/periods_14-1.2.2-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "periods_18" >}}     | {{< pkg "periods_17" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/periods_17-1.2.2-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "periods_16" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/periods_16-1.2.2-1.rhel9.1.aarch64.rpm" >}} | {{< pkg "periods_15" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/periods_15-1.2.2-1.rhel9.aarch64.rpm" >}} | {{< pkg "periods_14" "1.2.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/periods_14-1.2.2-1.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-periods" "1.2.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-18-periods` | 1.2.3 | `d12.x86_64` | pgdg | 47.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-periods` | 1.2.3 | `u22.aarch64` | pgdg | 45.5 KiB | [postgresql-18-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-periods` | 1.2.3 | `u22.x86_64` | pgdg | 46.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-periods` | 1.2.3 | `u24.x86_64` | pgdg | 46.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `periods_17` | 1.2.2 | `el8.aarch64` | pgdg | 44.1 KiB | [periods_17-1.2.2-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/periods_17-1.2.2-3PGDG.rhel8.aarch64.rpm) |
| `periods_17` | 1.2.2 | `el8.x86_64` | pgdg | 44.4 KiB | [periods_17-1.2.2-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/periods_17-1.2.2-3PGDG.rhel8.x86_64.rpm) |
| `periods_17` | 1.2.2 | `el9.aarch64` | pgdg | 42.1 KiB | [periods_17-1.2.2-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/periods_17-1.2.2-3PGDG.rhel9.aarch64.rpm) |
| `periods_17` | 1.2.2 | `el9.x86_64` | pgdg | 42.4 KiB | [periods_17-1.2.2-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/periods_17-1.2.2-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-periods` | 1.2.3 | `d12.x86_64` | pgdg | 47.0 KiB | [postgresql-17-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.3 KiB | [postgresql-17-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-periods` | 1.2.3 | `u22.x86_64` | pgdg | 50.1 KiB | [postgresql-17-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-periods` | 1.2.3 | `u22.aarch64` | pgdg | 49.6 KiB | [postgresql-17-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-periods` | 1.2.3 | `u24.x86_64` | pgdg | 46.1 KiB | [postgresql-17-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.3 KiB | [postgresql-17-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `periods_16` | 1.2.2 | `el8.x86_64` | pgdg | 44.2 KiB | [periods_16-1.2.2-1.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/periods_16-1.2.2-1.rhel8.1.x86_64.rpm) |
| `periods_16` | 1.2.2 | `el8.aarch64` | pgdg | 43.9 KiB | [periods_16-1.2.2-1.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/periods_16-1.2.2-1.rhel8.1.aarch64.rpm) |
| `periods_16` | 1.2.2 | `el9.aarch64` | pgdg | 41.6 KiB | [periods_16-1.2.2-1.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/periods_16-1.2.2-1.rhel9.1.aarch64.rpm) |
| `periods_16` | 1.2.2 | `el9.x86_64` | pgdg | 42.1 KiB | [periods_16-1.2.2-1.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/periods_16-1.2.2-1.rhel9.1.x86_64.rpm) |
| `postgresql-16-periods` | 1.2.3 | `d12.x86_64` | pgdg | 47.0 KiB | [postgresql-16-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.3 KiB | [postgresql-16-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-periods` | 1.2.3 | `u22.aarch64` | pgdg | 49.2 KiB | [postgresql-16-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-periods` | 1.2.3 | `u22.x86_64` | pgdg | 49.7 KiB | [postgresql-16-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-periods` | 1.2.3 | `u24.x86_64` | pgdg | 46.0 KiB | [postgresql-16-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.4 KiB | [postgresql-16-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `periods_15` | 1.2.2 | `el8.aarch64` | pgdg | 43.9 KiB | [periods_15-1.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/periods_15-1.2.2-1.rhel8.aarch64.rpm) |
| `periods_15` | 1.2.2 | `el8.x86_64` | pgdg | 44.2 KiB | [periods_15-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/periods_15-1.2.2-1.rhel8.x86_64.rpm) |
| `periods_15` | 1.2 | `el8.x86_64` | pgdg | 60.8 KiB | [periods_15-1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/periods_15-1.2-2.rhel8.x86_64.rpm) |
| `periods_15` | 1.2 | `el8.aarch64` | pgdg | 60.4 KiB | [periods_15-1.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/periods_15-1.2-2.rhel8.aarch64.rpm) |
| `periods_15` | 1.2.2 | `el9.x86_64` | pgdg | 42.1 KiB | [periods_15-1.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/periods_15-1.2.2-1.rhel9.x86_64.rpm) |
| `periods_15` | 1.2.2 | `el9.aarch64` | pgdg | 41.6 KiB | [periods_15-1.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/periods_15-1.2.2-1.rhel9.aarch64.rpm) |
| `periods_15` | 1.2 | `el9.aarch64` | pgdg | 59.2 KiB | [periods_15-1.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/periods_15-1.2-2.rhel9.aarch64.rpm) |
| `periods_15` | 1.2 | `el9.x86_64` | pgdg | 59.8 KiB | [periods_15-1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/periods_15-1.2-2.rhel9.x86_64.rpm) |
| `postgresql-15-periods` | 1.2.3 | `d12.x86_64` | pgdg | 46.9 KiB | [postgresql-15-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-periods` | 1.2.3 | `u22.x86_64` | pgdg | 49.7 KiB | [postgresql-15-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-periods` | 1.2.3 | `u22.aarch64` | pgdg | 49.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-periods` | 1.2.3 | `u24.x86_64` | pgdg | 46.0 KiB | [postgresql-15-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `periods_14` | 1.2.2 | `el8.aarch64` | pgdg | 43.9 KiB | [periods_14-1.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/periods_14-1.2.2-1.rhel8.aarch64.rpm) |
| `periods_14` | 1.2.2 | `el8.x86_64` | pgdg | 44.2 KiB | [periods_14-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/periods_14-1.2.2-1.rhel8.x86_64.rpm) |
| `periods_14` | 1.2 | `el8.aarch64` | pgdg | 60.2 KiB | [periods_14-1.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/periods_14-1.2-2.rhel8.aarch64.rpm) |
| `periods_14` | 1.2 | `el8.x86_64` | pgdg | 61.2 KiB | [periods_14-1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/periods_14-1.2-2.rhel8.x86_64.rpm) |
| `periods_14` | 1.2.2 | `el9.x86_64` | pgdg | 42.1 KiB | [periods_14-1.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/periods_14-1.2.2-1.rhel9.x86_64.rpm) |
| `periods_14` | 1.2.2 | `el9.aarch64` | pgdg | 41.6 KiB | [periods_14-1.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/periods_14-1.2.2-1.rhel9.aarch64.rpm) |
| `periods_14` | 1.2 | `el9.aarch64` | pgdg | 59.0 KiB | [periods_14-1.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/periods_14-1.2-2.rhel9.aarch64.rpm) |
| `postgresql-14-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.2 KiB | [postgresql-14-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-periods` | 1.2.3 | `d12.x86_64` | pgdg | 46.9 KiB | [postgresql-14-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-periods` | 1.2.3 | `u22.x86_64` | pgdg | 49.6 KiB | [postgresql-14-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-periods` | 1.2.3 | `u22.aarch64` | pgdg | 49.2 KiB | [postgresql-14-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.3 KiB | [postgresql-14-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-periods` | 1.2.3 | `u24.x86_64` | pgdg | 46.0 KiB | [postgresql-14-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `periods_13` | 1.2.2 | `el8.aarch64` | pgdg | 43.9 KiB | [periods_13-1.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/periods_13-1.2.2-1.rhel8.aarch64.rpm) |
| `periods_13` | 1.2.2 | `el8.x86_64` | pgdg | 44.1 KiB | [periods_13-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/periods_13-1.2.2-1.rhel8.x86_64.rpm) |
| `periods_13` | 1.2 | `el8.aarch64` | pgdg | 60.2 KiB | [periods_13-1.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/periods_13-1.2-2.rhel8.aarch64.rpm) |
| `periods_13` | 1.2 | `el8.x86_64` | pgdg | 60.5 KiB | [periods_13-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/periods_13-1.2-1.rhel8.x86_64.rpm) |
| `periods_13` | 1.2.2 | `el9.aarch64` | pgdg | 41.6 KiB | [periods_13-1.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/periods_13-1.2.2-1.rhel9.aarch64.rpm) |
| `periods_13` | 1.2.2 | `el9.x86_64` | pgdg | 42.1 KiB | [periods_13-1.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/periods_13-1.2.2-1.rhel9.x86_64.rpm) |
| `periods_13` | 1.2 | `el9.aarch64` | pgdg | 59.0 KiB | [periods_13-1.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/periods_13-1.2-2.rhel9.aarch64.rpm) |
| `postgresql-13-periods` | 1.2.3 | `d12.aarch64` | pgdg | 46.4 KiB | [postgresql-13-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-13-periods` | 1.2.3 | `d12.x86_64` | pgdg | 46.4 KiB | [postgresql-13-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-13-periods` | 1.2.3 | `u22.aarch64` | pgdg | 48.9 KiB | [postgresql-13-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-periods` | 1.2.3 | `u22.x86_64` | pgdg | 49.3 KiB | [postgresql-13-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-periods` | 1.2.3 | `u24.x86_64` | pgdg | 45.5 KiB | [postgresql-13-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-periods` | 1.2.3 | `u24.aarch64` | pgdg | 45.2 KiB | [postgresql-13-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-13-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/xocolatl/periods" title="Repository" icon="github" subtitle="github.com/xocolatl/periods" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install periods; # install by extension name, for the current active PG version
pig ext install periods; # install via package alias, for the active PG version
pig ext install periods -v 18;   # install for PG 18
pig ext install periods -v 17;   # install for PG 17
pig ext install periods -v 16;   # install for PG 16
pig ext install periods -v 15;   # install for PG 15
pig ext install periods -v 14;   # install for PG 14
pig ext install periods -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION periods;
```

