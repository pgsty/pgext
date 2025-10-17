---
title: "hll"
linkTitle: "hll"
description: "type for storing hyperloglog data"
weight: 2770
categories: ["Feat"]
width: full
---

type for storing hyperloglog data

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2770** | {{< badge content="hll" link="https://github.com/citusdata/postgresql-hll" >}} | {{< ext "hll" "hll" >}} | `2.18` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "bloom" >}} {{< ext "roaringbitmap" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "citus" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/hll" >}} | `2.18` | {{< badge content="18" color="red" alt="hll_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `hll_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/hll" >}} | `2.18` | {{< badge content="18" color="red" alt="postgresql-18-hll" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-hll` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "hll_18" >}}     | {{< pkg "hll_17" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hll_17-2.18-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "hll_16" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hll_16-2.18-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "hll_15" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hll_15-2.18-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "hll_14" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hll_14-2.18-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "hll_18" >}}     | {{< pkg "hll_17" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hll_17-2.18-2PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "hll_16" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hll_16-2.18-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "hll_15" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hll_15-2.18-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "hll_14" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hll_14-2.18-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "hll_18" >}}     | {{< pkg "hll_17" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hll_17-2.18-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "hll_16" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hll_16-2.18-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "hll_15" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hll_15-2.18-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "hll_14" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hll_14-2.18-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "hll_18" >}}     | {{< pkg "hll_17" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hll_17-2.18-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "hll_16" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hll_16-2.18-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "hll_15" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hll_15-2.18-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "hll_14" "2.18" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hll_14-2.18-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-hll" "2.18" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-18-hll` | 2.18 | `d12.x86_64` | pgdg | 76.5 KiB | [postgresql-18-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-18-hll` | 2.18 | `d12.aarch64` | pgdg | 75.5 KiB | [postgresql-18-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-18-hll` | 2.18 | `u22.aarch64` | pgdg | 75.1 KiB | [postgresql-18-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-hll` | 2.18 | `u22.x86_64` | pgdg | 76.2 KiB | [postgresql-18-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-hll` | 2.18 | `u24.aarch64` | pgdg | 73.8 KiB | [postgresql-18-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-hll` | 2.18 | `u24.x86_64` | pgdg | 75.0 KiB | [postgresql-18-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.18-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hll_17` | 2.18 | `el8.aarch64` | pgdg | 41.4 KiB | [hll_17-2.18-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hll_17-2.18-2PGDG.rhel8.aarch64.rpm) |
| `hll_17` | 2.18 | `el8.x86_64` | pgdg | 41.8 KiB | [hll_17-2.18-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hll_17-2.18-2PGDG.rhel8.x86_64.rpm) |
| `hll_17` | 2.18 | `el9.aarch64` | pgdg | 41.8 KiB | [hll_17-2.18-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hll_17-2.18-2PGDG.rhel9.aarch64.rpm) |
| `hll_17` | 2.18 | `el9.x86_64` | pgdg | 41.9 KiB | [hll_17-2.18-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hll_17-2.18-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-hll` | 2.18 | `d12.x86_64` | pgdg | 76.5 KiB | [postgresql-17-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-17-hll` | 2.18 | `d12.aarch64` | pgdg | 75.4 KiB | [postgresql-17-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-17-hll` | 2.18 | `u22.x86_64` | pgdg | 82.7 KiB | [postgresql-17-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-hll` | 2.18 | `u22.aarch64` | pgdg | 81.5 KiB | [postgresql-17-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-hll` | 2.18 | `u24.x86_64` | pgdg | 75.0 KiB | [postgresql-17-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-hll` | 2.18 | `u24.aarch64` | pgdg | 73.9 KiB | [postgresql-17-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.18-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hll_16` | 2.18 | `el8.x86_64` | pgdg | 41.7 KiB | [hll_16-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hll_16-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_16` | 2.18 | `el8.aarch64` | pgdg | 41.3 KiB | [hll_16-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hll_16-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_16` | 2.18 | `el9.aarch64` | pgdg | 41.8 KiB | [hll_16-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hll_16-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_16` | 2.18 | `el9.x86_64` | pgdg | 41.6 KiB | [hll_16-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hll_16-2.18-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-hll` | 2.18 | `d12.x86_64` | pgdg | 76.5 KiB | [postgresql-16-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-16-hll` | 2.18 | `d12.aarch64` | pgdg | 75.5 KiB | [postgresql-16-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-16-hll` | 2.18 | `u22.aarch64` | pgdg | 80.9 KiB | [postgresql-16-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-hll` | 2.18 | `u22.x86_64` | pgdg | 82.2 KiB | [postgresql-16-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-hll` | 2.18 | `u24.x86_64` | pgdg | 74.9 KiB | [postgresql-16-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-hll` | 2.18 | `u24.aarch64` | pgdg | 73.9 KiB | [postgresql-16-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.18-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hll_15` | 2.18 | `el8.aarch64` | pgdg | 41.6 KiB | [hll_15-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hll_15-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_15` | 2.18 | `el8.x86_64` | pgdg | 42.0 KiB | [hll_15-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hll_15-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_15` | 2.17 | `el8.x86_64` | pgdg | 89.5 KiB | [hll_15-2.17-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hll_15-2.17-1.rhel8.x86_64.rpm) |
| `hll_15` | 2.17 | `el8.aarch64` | pgdg | 89.2 KiB | [hll_15-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hll_15-2.17-1.rhel8.aarch64.rpm) |
| `hll_15` | 2.18 | `el9.x86_64` | pgdg | 41.9 KiB | [hll_15-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hll_15-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_15` | 2.18 | `el9.aarch64` | pgdg | 41.9 KiB | [hll_15-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hll_15-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_15` | 2.17 | `el9.aarch64` | pgdg | 91.1 KiB | [hll_15-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hll_15-2.17-1.rhel9.aarch64.rpm) |
| `hll_15` | 2.17 | `el9.x86_64` | pgdg | 90.6 KiB | [hll_15-2.17-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hll_15-2.17-1.rhel9.x86_64.rpm) |
| `postgresql-15-hll` | 2.18 | `d12.x86_64` | pgdg | 77.0 KiB | [postgresql-15-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-15-hll` | 2.18 | `d12.aarch64` | pgdg | 75.9 KiB | [postgresql-15-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-15-hll` | 2.18 | `u22.x86_64` | pgdg | 83.1 KiB | [postgresql-15-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-hll` | 2.18 | `u22.aarch64` | pgdg | 82.2 KiB | [postgresql-15-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-hll` | 2.18 | `u24.x86_64` | pgdg | 76.2 KiB | [postgresql-15-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-hll` | 2.18 | `u24.aarch64` | pgdg | 75.5 KiB | [postgresql-15-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.18-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hll_14` | 2.18 | `el8.aarch64` | pgdg | 41.6 KiB | [hll_14-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hll_14-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_14` | 2.18 | `el8.x86_64` | pgdg | 42.0 KiB | [hll_14-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hll_14-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_14` | 2.17 | `el8.aarch64` | pgdg | 89.0 KiB | [hll_14-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hll_14-2.17-1.rhel8.aarch64.rpm) |
| `hll_14` | 2.16 | `el8.x86_64` | pgdg | 90.0 KiB | [hll_14-2.16-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hll_14-2.16-1.rhel8.x86_64.rpm) |
| `hll_14` | 2.18 | `el9.x86_64` | pgdg | 41.9 KiB | [hll_14-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hll_14-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_14` | 2.18 | `el9.aarch64` | pgdg | 41.9 KiB | [hll_14-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hll_14-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_14` | 2.17 | `el9.aarch64` | pgdg | 90.9 KiB | [hll_14-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hll_14-2.17-1.rhel9.aarch64.rpm) |
| `postgresql-14-hll` | 2.18 | `d12.aarch64` | pgdg | 75.8 KiB | [postgresql-14-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-14-hll` | 2.18 | `d12.x86_64` | pgdg | 76.8 KiB | [postgresql-14-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-14-hll` | 2.18 | `u22.x86_64` | pgdg | 83.0 KiB | [postgresql-14-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-hll` | 2.18 | `u22.aarch64` | pgdg | 82.1 KiB | [postgresql-14-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-hll` | 2.18 | `u24.aarch64` | pgdg | 75.3 KiB | [postgresql-14-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-hll` | 2.18 | `u24.x86_64` | pgdg | 76.0 KiB | [postgresql-14-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.18-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hll_13` | 2.18 | `el8.aarch64` | pgdg | 41.5 KiB | [hll_13-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/hll_13-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_13` | 2.18 | `el8.x86_64` | pgdg | 41.5 KiB | [hll_13-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/hll_13-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_13` | 2.17 | `el8.aarch64` | pgdg | 88.7 KiB | [hll_13-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/hll_13-2.17-1.rhel8.aarch64.rpm) |
| `hll_13` | 2.15 | `el8.x86_64` | pgdg | 88.4 KiB | [hll_13-2.15-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/hll_13-2.15-1.rhel8.x86_64.rpm) |
| `hll_13` | 2.18 | `el9.aarch64` | pgdg | 41.9 KiB | [hll_13-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/hll_13-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_13` | 2.18 | `el9.x86_64` | pgdg | 41.8 KiB | [hll_13-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/hll_13-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_13` | 2.17 | `el9.aarch64` | pgdg | 90.7 KiB | [hll_13-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/hll_13-2.17-1.rhel9.aarch64.rpm) |
| `postgresql-13-hll` | 2.18 | `d12.aarch64` | pgdg | 75.5 KiB | [postgresql-13-hll_2.18-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg12+1_arm64.deb) |
| `postgresql-13-hll` | 2.18 | `d12.x86_64` | pgdg | 76.5 KiB | [postgresql-13-hll_2.18-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg12+1_amd64.deb) |
| `postgresql-13-hll` | 2.18 | `u22.aarch64` | pgdg | 81.8 KiB | [postgresql-13-hll_2.18-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-hll` | 2.18 | `u22.x86_64` | pgdg | 82.8 KiB | [postgresql-13-hll_2.18-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-hll` | 2.18 | `u24.x86_64` | pgdg | 76.1 KiB | [postgresql-13-hll_2.18-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-hll` | 2.18 | `u24.aarch64` | pgdg | 74.9 KiB | [postgresql-13-hll_2.18-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.18-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/postgresql-hll" title="Repository" icon="github" subtitle="github.com/citusdata/postgresql-hll" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hll; # install by extension name, for the current active PG version
pig ext install hll; # install via package alias, for the active PG version
pig ext install hll -v 17;   # install for PG 17
pig ext install hll -v 16;   # install for PG 16
pig ext install hll -v 15;   # install for PG 15
pig ext install hll -v 14;   # install for PG 14
pig ext install hll -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hll;
```

