---
title: "powa"
linkTitle: "powa"
description: "PostgreSQL Workload Analyser-core"
weight: 6810
categories: ["Stat"]
width: full
---

PostgreSQL Workload Analyser-core

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6810** | {{< badge content="powa" link="https://github.com/powa-team/powa" >}} | {{< ext "powa" "powa" >}} | `5.0.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Python" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "pg_stat_statements" >}} {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "pg_stat_kcache" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_wait_sampling" >}} {{< ext "hypopg" >}} {{< ext "plprofiler" >}} {{< ext "pg_profile" >}} {{< ext "pg_track_settings" >}} {{< ext "btree_gin" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/powa" >}} | `5.0.1` | {{< badge content="18" color="red" alt="powa_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `powa_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/powa" >}} | `5.0.1` | {{< badge content="18" color="red" alt="postgresql-18-powa" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-powa` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "powa_18" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/powa_18-5.0.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "powa_17" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/powa_17-5.0.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "powa_16" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/powa_16-5.0.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "powa_15" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-5.0.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "powa_14" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-5.0.1-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "powa_18" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/powa_18-5.0.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "powa_17" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/powa_17-5.0.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "powa_16" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/powa_16-5.0.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "powa_15" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-5.0.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "powa_14" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-5.0.1-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "powa_18" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/powa_18-5.0.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "powa_17" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/powa_17-5.0.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "powa_16" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/powa_16-5.0.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "powa_15" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-5.0.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "powa_14" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-5.0.1-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "powa_18" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/powa_18-5.0.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "powa_17" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/powa_17-5.0.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "powa_16" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/powa_16-5.0.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "powa_15" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-5.0.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "powa_14" "5.0.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-5.0.1-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg120+1_amd64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg120+1_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg120+1_arm64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg120+1_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-powa" >}}     | {{< pkg "postgresql-17-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-powa" "5.0.3" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_18` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_18-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/powa_18-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_18` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_18-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/powa_18-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_18` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_18-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/powa_18-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_18` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_18-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/powa_18-5.0.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_17` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_17-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/powa_17-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_17` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_17-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/powa_17-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_17` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_17-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/powa_17-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_17` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_17-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/powa_17-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-powa` | 5.0.3 | `d12.aarch64` | pgdg | 61.2 KiB | [postgresql-17-powa_5.0.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg120+1_arm64.deb) |
| `postgresql-17-powa` | 5.0.3 | `d12.x86_64` | pgdg | 61.5 KiB | [postgresql-17-powa_5.0.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg120+1_amd64.deb) |
| `postgresql-17-powa` | 5.0.3 | `u22.x86_64` | pgdg | 61.6 KiB | [postgresql-17-powa_5.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-powa` | 5.0.3 | `u22.aarch64` | pgdg | 61.4 KiB | [postgresql-17-powa_5.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-powa` | 5.0.3 | `u24.x86_64` | pgdg | 56.8 KiB | [postgresql-17-powa_5.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-powa` | 5.0.3 | `u24.aarch64` | pgdg | 56.4 KiB | [postgresql-17-powa_5.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-17-powa_5.0.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_16` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_16-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/powa_16-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_16` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_16-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/powa_16-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_16` | 4.2.2 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_16-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/powa_16-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `powa_16` | 4.2.2 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_16-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/powa_16-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `powa_16` | 4.2.1 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_16-4.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/powa_16-4.2.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_16` | 4.2.1 | `el8.x86_64` | pgdg | 6.5 KiB | [powa_16-4.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/powa_16-4.2.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_16` | 4.2.0 | `el8.aarch64` | pgdg | 6.4 KiB | [powa_16-4.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/powa_16-4.2.0-1PGDG.rhel8.aarch64.rpm) |
| `powa_16` | 4.2.0 | `el8.x86_64` | pgdg | 6.4 KiB | [powa_16-4.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/powa_16-4.2.0-1PGDG.rhel8.x86_64.rpm) |
| `powa_16` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_16-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/powa_16-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_16` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_16-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/powa_16-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_16` | 4.2.2 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_16-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/powa_16-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `powa_16` | 4.2.2 | `el9.aarch64` | pgdg | 6.5 KiB | [powa_16-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/powa_16-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `powa_16` | 4.2.1 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_16-4.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/powa_16-4.2.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_16` | 4.2.1 | `el9.aarch64` | pgdg | 6.4 KiB | [powa_16-4.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/powa_16-4.2.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_16` | 4.2.0 | `el9.x86_64` | pgdg | 6.5 KiB | [powa_16-4.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/powa_16-4.2.0-1PGDG.rhel9.x86_64.rpm) |
| `powa_16` | 4.2.0 | `el9.aarch64` | pgdg | 6.3 KiB | [powa_16-4.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/powa_16-4.2.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-powa` | 5.0.3 | `d12.x86_64` | pgdg | 61.6 KiB | [postgresql-16-powa_5.0.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg120+1_amd64.deb) |
| `postgresql-16-powa` | 5.0.3 | `d12.aarch64` | pgdg | 61.2 KiB | [postgresql-16-powa_5.0.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg120+1_arm64.deb) |
| `postgresql-16-powa` | 5.0.3 | `u22.x86_64` | pgdg | 61.1 KiB | [postgresql-16-powa_5.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-powa` | 5.0.3 | `u22.aarch64` | pgdg | 60.9 KiB | [postgresql-16-powa_5.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-powa` | 5.0.3 | `u24.x86_64` | pgdg | 56.8 KiB | [postgresql-16-powa_5.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-powa` | 5.0.3 | `u24.aarch64` | pgdg | 56.4 KiB | [postgresql-16-powa_5.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-16-powa_5.0.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_15` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_15-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_15` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_15-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_15` | 4.2.2 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_15-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `powa_15` | 4.2.2 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_15-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `powa_15` | 4.2.1 | `el8.x86_64` | pgdg | 6.5 KiB | [powa_15-4.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-4.2.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_15` | 4.2.1 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_15-4.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-4.2.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_15` | 4.2.0 | `el8.x86_64` | pgdg | 6.4 KiB | [powa_15-4.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-4.2.0-1PGDG.rhel8.x86_64.rpm) |
| `powa_15` | 4.2.0 | `el8.aarch64` | pgdg | 6.4 KiB | [powa_15-4.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-4.2.0-1PGDG.rhel8.aarch64.rpm) |
| `powa_15` | 4.1.4 | `el8.aarch64` | pgdg | 66.7 KiB | [powa_15-4.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/powa_15-4.1.4-1.rhel8.aarch64.rpm) |
| `powa_15` | 4.1.4 | `el8.x86_64` | pgdg | 66.9 KiB | [powa_15-4.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/powa_15-4.1.4-1.rhel8.x86_64.rpm) |
| `powa_15` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_15-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_15` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_15-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_15` | 4.2.2 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_15-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `powa_15` | 4.2.2 | `el9.aarch64` | pgdg | 6.5 KiB | [powa_15-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `powa_15` | 4.2.1 | `el9.aarch64` | pgdg | 6.4 KiB | [powa_15-4.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-4.2.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_15` | 4.2.1 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_15-4.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-4.2.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_15` | 4.2.0 | `el9.x86_64` | pgdg | 6.5 KiB | [powa_15-4.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-4.2.0-1PGDG.rhel9.x86_64.rpm) |
| `powa_15` | 4.2.0 | `el9.aarch64` | pgdg | 6.3 KiB | [powa_15-4.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-4.2.0-1PGDG.rhel9.aarch64.rpm) |
| `powa_15` | 4.1.4 | `el9.aarch64` | pgdg | 61.0 KiB | [powa_15-4.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/powa_15-4.1.4-1.rhel9.aarch64.rpm) |
| `powa_15` | 4.1.4 | `el9.x86_64` | pgdg | 61.6 KiB | [powa_15-4.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/powa_15-4.1.4-1.rhel9.x86_64.rpm) |
| `postgresql-15-powa` | 5.0.3 | `d12.aarch64` | pgdg | 61.7 KiB | [postgresql-15-powa_5.0.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg120+1_arm64.deb) |
| `postgresql-15-powa` | 5.0.3 | `d12.x86_64` | pgdg | 61.7 KiB | [postgresql-15-powa_5.0.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg120+1_amd64.deb) |
| `postgresql-15-powa` | 5.0.3 | `u22.aarch64` | pgdg | 61.2 KiB | [postgresql-15-powa_5.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-powa` | 5.0.3 | `u22.x86_64` | pgdg | 61.6 KiB | [postgresql-15-powa_5.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-powa` | 5.0.3 | `u24.x86_64` | pgdg | 57.1 KiB | [postgresql-15-powa_5.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-powa` | 5.0.3 | `u24.aarch64` | pgdg | 57.0 KiB | [postgresql-15-powa_5.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-15-powa_5.0.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_14` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_14-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_14` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_14-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_14` | 4.2.2 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_14-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `powa_14` | 4.2.2 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_14-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `powa_14` | 4.2.1 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_14-4.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-4.2.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_14` | 4.2.1 | `el8.x86_64` | pgdg | 6.5 KiB | [powa_14-4.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.2.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_14` | 4.2.0 | `el8.aarch64` | pgdg | 6.4 KiB | [powa_14-4.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-4.2.0-1PGDG.rhel8.aarch64.rpm) |
| `powa_14` | 4.2.0 | `el8.x86_64` | pgdg | 6.4 KiB | [powa_14-4.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.2.0-1PGDG.rhel8.x86_64.rpm) |
| `powa_14` | 4.1.4 | `el8.x86_64` | pgdg | 68.0 KiB | [powa_14-4.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.1.4-1.rhel8.x86_64.rpm) |
| `powa_14` | 4.1.4 | `el8.aarch64` | pgdg | 67.8 KiB | [powa_14-4.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/powa_14-4.1.4-1.rhel8.aarch64.rpm) |
| `powa_14` | 4.1.3 | `el8.x86_64` | pgdg | 68.1 KiB | [powa_14-4.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.1.3-1.rhel8.x86_64.rpm) |
| `powa_14` | 4.1.2 | `el8.x86_64` | pgdg | 66.1 KiB | [powa_14-4.1.2-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/powa_14-4.1.2-4.rhel8.x86_64.rpm) |
| `powa_14` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_14-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_14` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_14-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_14` | 4.2.2 | `el9.aarch64` | pgdg | 6.5 KiB | [powa_14-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `powa_14` | 4.2.2 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_14-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `powa_14` | 4.2.1 | `el9.aarch64` | pgdg | 6.4 KiB | [powa_14-4.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-4.2.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_14` | 4.2.1 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_14-4.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-4.2.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_14` | 4.2.0 | `el9.x86_64` | pgdg | 6.5 KiB | [powa_14-4.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-4.2.0-1PGDG.rhel9.x86_64.rpm) |
| `powa_14` | 4.2.0 | `el9.aarch64` | pgdg | 6.3 KiB | [powa_14-4.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-4.2.0-1PGDG.rhel9.aarch64.rpm) |
| `powa_14` | 4.1.4 | `el9.x86_64` | pgdg | 62.7 KiB | [powa_14-4.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-4.1.4-1.rhel9.x86_64.rpm) |
| `powa_14` | 4.1.4 | `el9.aarch64` | pgdg | 62.2 KiB | [powa_14-4.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/powa_14-4.1.4-1.rhel9.aarch64.rpm) |
| `powa_14` | 4.1.3 | `el9.x86_64` | pgdg | 62.2 KiB | [powa_14-4.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/powa_14-4.1.3-1.rhel9.x86_64.rpm) |
| `postgresql-14-powa` | 5.0.3 | `d12.x86_64` | pgdg | 63.6 KiB | [postgresql-14-powa_5.0.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg120+1_amd64.deb) |
| `postgresql-14-powa` | 5.0.3 | `d12.aarch64` | pgdg | 63.5 KiB | [postgresql-14-powa_5.0.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg120+1_arm64.deb) |
| `postgresql-14-powa` | 5.0.3 | `u22.aarch64` | pgdg | 62.0 KiB | [postgresql-14-powa_5.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-powa` | 5.0.3 | `u22.x86_64` | pgdg | 62.4 KiB | [postgresql-14-powa_5.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-powa` | 5.0.3 | `u24.aarch64` | pgdg | 58.7 KiB | [postgresql-14-powa_5.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-powa` | 5.0.3 | `u24.x86_64` | pgdg | 58.9 KiB | [postgresql-14-powa_5.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-14-powa_5.0.3-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `powa_13` | 5.0.1 | `el8.aarch64` | pgdg | 6.6 KiB | [powa_13-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/powa_13-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_13` | 5.0.1 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_13-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_13` | 4.2.2 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_13-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/powa_13-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `powa_13` | 4.2.2 | `el8.x86_64` | pgdg | 6.6 KiB | [powa_13-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `powa_13` | 4.2.1 | `el8.x86_64` | pgdg | 6.5 KiB | [powa_13-4.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.2.1-1PGDG.rhel8.x86_64.rpm) |
| `powa_13` | 4.2.1 | `el8.aarch64` | pgdg | 6.5 KiB | [powa_13-4.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/powa_13-4.2.1-1PGDG.rhel8.aarch64.rpm) |
| `powa_13` | 4.2.0 | `el8.aarch64` | pgdg | 6.4 KiB | [powa_13-4.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/powa_13-4.2.0-1PGDG.rhel8.aarch64.rpm) |
| `powa_13` | 4.2.0 | `el8.x86_64` | pgdg | 6.4 KiB | [powa_13-4.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.2.0-1PGDG.rhel8.x86_64.rpm) |
| `powa_13` | 4.1.4 | `el8.x86_64` | pgdg | 67.7 KiB | [powa_13-4.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.1.4-1.rhel8.x86_64.rpm) |
| `powa_13` | 4.1.4 | `el8.aarch64` | pgdg | 67.5 KiB | [powa_13-4.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/powa_13-4.1.4-1.rhel8.aarch64.rpm) |
| `powa_13` | 4.1.3 | `el8.x86_64` | pgdg | 67.9 KiB | [powa_13-4.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.1.3-1.rhel8.x86_64.rpm) |
| `powa_13` | 4.1.2 | `el8.x86_64` | pgdg | 65.6 KiB | [powa_13-4.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.1.2-1.rhel8.x86_64.rpm) |
| `powa_13` | 4.1.2 | `el8.x86_64` | pgdg | 65.9 KiB | [powa_13-4.1.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.1.2-3.rhel8.x86_64.rpm) |
| `powa_13` | 4.1.1 | `el8.x86_64` | pgdg | 64.6 KiB | [powa_13-4.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/powa_13-4.1.1-1.rhel8.x86_64.rpm) |
| `powa_13` | 5.0.1 | `el9.aarch64` | pgdg | 6.6 KiB | [powa_13-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/powa_13-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_13` | 5.0.1 | `el9.x86_64` | pgdg | 6.7 KiB | [powa_13-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_13` | 4.2.2 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_13-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `powa_13` | 4.2.2 | `el9.aarch64` | pgdg | 6.5 KiB | [powa_13-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/powa_13-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `powa_13` | 4.2.1 | `el9.x86_64` | pgdg | 6.6 KiB | [powa_13-4.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-4.2.1-1PGDG.rhel9.x86_64.rpm) |
| `powa_13` | 4.2.1 | `el9.aarch64` | pgdg | 6.4 KiB | [powa_13-4.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/powa_13-4.2.1-1PGDG.rhel9.aarch64.rpm) |
| `powa_13` | 4.2.0 | `el9.x86_64` | pgdg | 6.5 KiB | [powa_13-4.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-4.2.0-1PGDG.rhel9.x86_64.rpm) |
| `powa_13` | 4.2.0 | `el9.aarch64` | pgdg | 6.3 KiB | [powa_13-4.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/powa_13-4.2.0-1PGDG.rhel9.aarch64.rpm) |
| `powa_13` | 4.1.4 | `el9.x86_64` | pgdg | 62.4 KiB | [powa_13-4.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-4.1.4-1.rhel9.x86_64.rpm) |
| `powa_13` | 4.1.4 | `el9.aarch64` | pgdg | 61.9 KiB | [powa_13-4.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/powa_13-4.1.4-1.rhel9.aarch64.rpm) |
| `powa_13` | 4.1.3 | `el9.x86_64` | pgdg | 62.0 KiB | [powa_13-4.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/powa_13-4.1.3-1.rhel9.x86_64.rpm) |
| `postgresql-13-powa` | 5.0.3 | `d12.aarch64` | pgdg | 63.4 KiB | [postgresql-13-powa_5.0.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg120+1_arm64.deb) |
| `postgresql-13-powa` | 5.0.3 | `d12.x86_64` | pgdg | 63.6 KiB | [postgresql-13-powa_5.0.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg120+1_amd64.deb) |
| `postgresql-13-powa` | 5.0.3 | `u22.x86_64` | pgdg | 62.2 KiB | [postgresql-13-powa_5.0.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-powa` | 5.0.3 | `u22.aarch64` | pgdg | 61.6 KiB | [postgresql-13-powa_5.0.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-powa` | 5.0.3 | `u24.aarch64` | pgdg | 58.6 KiB | [postgresql-13-powa_5.0.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-13-powa` | 5.0.3 | `u24.x86_64` | pgdg | 59.0 KiB | [postgresql-13-powa_5.0.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/powa-archivist/postgresql-13-powa_5.0.3-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/powa-team/powa" title="Repository" icon="github" subtitle="github.com/powa-team/powa" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install powa; # install by extension name, for the current active PG version
pig ext install powa; # install via package alias, for the active PG version
pig ext install powa -v 17;   # install for PG 17
pig ext install powa -v 16;   # install for PG 16
pig ext install powa -v 15;   # install for PG 15
pig ext install powa -v 14;   # install for PG 14
pig ext install powa -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION powa CASCADE SCHEMA public;
```

