---
title: "pgmemcache"
linkTitle: "pgmemcache"
description: "memcached interface"
weight: 9410
categories: ["Sim"]
width: full
---

memcached interface

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9410** | {{< badge content="pgmemcache" link="https://github.com/ohmu/pgmemcache" >}} | {{< ext "pgmemcache" "pgmemcache" >}} | `2.3.0` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "redis_fdw" >}} {{< ext "redis" >}} {{< ext "spat" >}} {{< ext "mongo_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] missing pg12-14 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgmemcache" >}} | `2.3.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="pgmemcache_14*" >}} | `pgmemcache_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgmemcache" >}} | `2.3.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgmemcache` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgmemcache_18" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgmemcache_17" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgmemcache_16" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmemcache_16-2.3.0-6.rhel8.1.x86_64.rpm" >}} | {{< pkg "pgmemcache_15" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmemcache_15-2.3.0-5.rhel8.x86_64.rpm" >}} | {{< pkg "pgmemcache_14" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmemcache_14-2.3.0-5.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgmemcache_18" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgmemcache_17" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgmemcache_16" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmemcache_16-2.3.0-6.rhel8.1.aarch64.rpm" >}} |    {{< pkg "pgmemcache_15" >}}     |    {{< pkg "pgmemcache_14" >}}     |
|    `el9.x86_64`    | {{< pkg "pgmemcache_18" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgmemcache_17" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgmemcache_16" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmemcache_16-2.3.0-6.rhel9.1.x86_64.rpm" >}} | {{< pkg "pgmemcache_15" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmemcache_15-2.3.0-5.rhel9.x86_64.rpm" >}} |    {{< pkg "pgmemcache_14" >}}     |
|    `el9.aarch64`    | {{< pkg "pgmemcache_18" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgmemcache_17" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgmemcache_16" "2.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmemcache_16-2.3.0-6.rhel9.1.aarch64.rpm" >}} |    {{< pkg "pgmemcache_15" >}}     |    {{< pkg "pgmemcache_14" >}}     |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgmemcache" "2.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgmemcache_18` | 2.3.0 | `el8.aarch64` | pgdg | 26.0 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel8.aarch64.rpm) |
| `pgmemcache_18` | 2.3.0 | `el8.x86_64` | pgdg | 26.5 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel8.x86_64.rpm) |
| `pgmemcache_18` | 2.3.0 | `el9.x86_64` | pgdg | 25.9 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel9.x86_64.rpm) |
| `pgmemcache_18` | 2.3.0 | `el9.aarch64` | pgdg | 25.4 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.4 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 45.0 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 46.3 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 46.5 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.1 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.7 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgmemcache_17` | 2.3.0 | `el8.x86_64` | pgdg | 26.4 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel8.x86_64.rpm) |
| `pgmemcache_17` | 2.3.0 | `el8.aarch64` | pgdg | 25.9 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel8.aarch64.rpm) |
| `pgmemcache_17` | 2.3.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel9.x86_64.rpm) |
| `pgmemcache_17` | 2.3.0 | `el9.aarch64` | pgdg | 25.6 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 45.0 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 51.9 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 52.3 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.7 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.2 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgmemcache_16` | 2.3.0 | `el8.x86_64` | pgdg | 26.2 KiB | [pgmemcache_16-2.3.0-6.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmemcache_16-2.3.0-6.rhel8.1.x86_64.rpm) |
| `pgmemcache_16` | 2.3.0 | `el8.aarch64` | pgdg | 25.7 KiB | [pgmemcache_16-2.3.0-6.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmemcache_16-2.3.0-6.rhel8.1.aarch64.rpm) |
| `pgmemcache_16` | 2.3.0 | `el9.aarch64` | pgdg | 25.1 KiB | [pgmemcache_16-2.3.0-6.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmemcache_16-2.3.0-6.rhel9.1.aarch64.rpm) |
| `pgmemcache_16` | 2.3.0 | `el9.x86_64` | pgdg | 25.7 KiB | [pgmemcache_16-2.3.0-6.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmemcache_16-2.3.0-6.rhel9.1.x86_64.rpm) |
| `postgresql-16-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 44.9 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.4 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 51.7 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 51.6 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.1 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.7 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgmemcache_15` | 2.3.0 | `el8.x86_64` | pgdg | 56.1 KiB | [pgmemcache_15-2.3.0-5.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmemcache_15-2.3.0-5.rhel8.x86_64.rpm) |
| `pgmemcache_15` | 2.3.0 | `el9.x86_64` | pgdg | 57.3 KiB | [pgmemcache_15-2.3.0-5.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmemcache_15-2.3.0-5.rhel9.x86_64.rpm) |
| `postgresql-15-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 45.1 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 51.8 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 51.6 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.1 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.7 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgmemcache_14` | 2.3.0 | `el8.x86_64` | pgdg | 56.6 KiB | [pgmemcache_14-2.3.0-5.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmemcache_14-2.3.0-5.rhel8.x86_64.rpm) |
| `postgresql-14-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 45.1 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 51.7 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 51.5 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.2 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.7 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-13-pgmemcache` | 2.3.0 | `d12.aarch64` | pgdg | 44.9 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-13-pgmemcache` | 2.3.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-13-pgmemcache` | 2.3.0 | `u22.x86_64` | pgdg | 51.4 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgmemcache` | 2.3.0 | `u22.aarch64` | pgdg | 51.0 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgmemcache` | 2.3.0 | `u24.x86_64` | pgdg | 45.4 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgmemcache` | 2.3.0 | `u24.aarch64` | pgdg | 45.1 KiB | [postgresql-13-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-13-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ohmu/pgmemcache" title="Repository" icon="github" subtitle="github.com/ohmu/pgmemcache" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgmemcache; # install by extension name, for the current active PG version
pig ext install pgmemcache; # install via package alias, for the active PG version
pig ext install pgmemcache -v 18;   # install for PG 18
pig ext install pgmemcache -v 17;   # install for PG 17
pig ext install pgmemcache -v 16;   # install for PG 16
pig ext install pgmemcache -v 15;   # install for PG 15
pig ext install pgmemcache -v 14;   # install for PG 14
pig ext install pgmemcache -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgmemcache;
```

