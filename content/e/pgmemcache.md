---
title: "pgmemcache"
linkTitle: "pgmemcache"
description: "memcached interface"
weight: 9410
categories: ["SIM"]
width: full
---

[**pgmemcache**](https://github.com/ohmu/pgmemcache) : memcached interface


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9410** | {{< badge content="pgmemcache" link="https://github.com/ohmu/pgmemcache" >}} | {{< ext "pgmemcache" >}} | `2.3.0` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "redis_fdw" >}} {{< ext "redis" >}} {{< ext "spat" >}} {{< ext "mongo_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] missing pg12-14 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmemcache` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "pgmemcache_18" "green" >}} {{< bg "17" "pgmemcache_17" "green" >}} {{< bg "16" "pgmemcache_16" "green" >}} {{< bg "15" "pgmemcache_15" "green" >}} {{< bg "14" "pgmemcache_14" "green" >}} | `pgmemcache_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3.0` | {{< bg "18" "postgresql-18-pgmemcache" "green" >}} {{< bg "17" "postgresql-17-pgmemcache" "green" >}} {{< bg "16" "postgresql-16-pgmemcache" "green" >}} {{< bg "15" "postgresql-15-pgmemcache" "green" >}} {{< bg "14" "postgresql-14-pgmemcache" "green" >}} | `postgresql-$v-pgmemcache` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgmemcache_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgmemcache_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgmemcache_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgmemcache_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgmemcache_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "pgmemcache_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.3.0" "postgresql-18-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-17-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-16-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-15-pgmemcache : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.0" "postgresql-14-pgmemcache : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemcache_18` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.5 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel8.x86_64.rpm) |
| `pgmemcache_18` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel8.aarch64.rpm) |
| `pgmemcache_18` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.9 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel9.x86_64.rpm) |
| `pgmemcache_18` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel9.aarch64.rpm) |
| `pgmemcache_18` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.3 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmemcache_18-2.3.0-9PGDG.rhel10.x86_64.rpm) |
| `pgmemcache_18` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.2 KiB | [pgmemcache_18-2.3.0-9PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmemcache_18-2.3.0-9PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgmemcache` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.0 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.0 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.5 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.3 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgmemcache` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.1 KiB | [postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-18-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemcache_17` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.4 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel8.x86_64.rpm) |
| `pgmemcache_17` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.9 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel8.aarch64.rpm) |
| `pgmemcache_17` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.0 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmemcache_17-2.3.0-8PGDG.rhel9.x86_64.rpm) |
| `pgmemcache_17` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.6 KiB | [pgmemcache_17-2.3.0-8PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmemcache_17-2.3.0-8PGDG.rhel9.aarch64.rpm) |
| `pgmemcache_17` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.3 KiB | [pgmemcache_17-2.3.0-9PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmemcache_17-2.3.0-9PGDG.rhel10.x86_64.rpm) |
| `pgmemcache_17` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.2 KiB | [pgmemcache_17-2.3.0-9PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmemcache_17-2.3.0-9PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgmemcache` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.0 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.7 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.1 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.3 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.9 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmemcache` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.2 KiB | [postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-17-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemcache_16` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.2 KiB | [pgmemcache_16-2.3.0-6.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmemcache_16-2.3.0-6.rhel8.1.x86_64.rpm) |
| `pgmemcache_16` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.7 KiB | [pgmemcache_16-2.3.0-6.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmemcache_16-2.3.0-6.rhel8.1.aarch64.rpm) |
| `pgmemcache_16` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pgmemcache_16-2.3.0-6.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmemcache_16-2.3.0-6.rhel9.1.x86_64.rpm) |
| `pgmemcache_16` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.1 KiB | [pgmemcache_16-2.3.0-6.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmemcache_16-2.3.0-6.rhel9.1.aarch64.rpm) |
| `pgmemcache_16` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.4 KiB | [pgmemcache_16-2.3.0-9PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmemcache_16-2.3.0-9PGDG.rhel10.x86_64.rpm) |
| `pgmemcache_16` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.2 KiB | [pgmemcache_16-2.3.0-9PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmemcache_16-2.3.0-9PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgmemcache` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.1 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.7 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.6 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgmemcache` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.1 KiB | [postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-16-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemcache_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.1 KiB | [pgmemcache_15-2.3.0-5.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmemcache_15-2.3.0-5.rhel8.x86_64.rpm) |
| `pgmemcache_15` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pgmemcache_15-2.3.0-5.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmemcache_15-2.3.0-5.rhel9.x86_64.rpm) |
| `pgmemcache_15` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.4 KiB | [pgmemcache_15-2.3.0-9PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmemcache_15-2.3.0-9PGDG.rhel10.x86_64.rpm) |
| `pgmemcache_15` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.2 KiB | [pgmemcache_15-2.3.0-9PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmemcache_15-2.3.0-9PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgmemcache` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.1 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.7 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.2 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.8 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.6 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgmemcache` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.1 KiB | [postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-15-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemcache_14` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.6 KiB | [pgmemcache_14-2.3.0-5.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmemcache_14-2.3.0-5.rhel8.x86_64.rpm) |
| `pgmemcache_14` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.4 KiB | [pgmemcache_14-2.3.0-9PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmemcache_14-2.3.0-9PGDG.rhel10.x86_64.rpm) |
| `pgmemcache_14` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.2 KiB | [pgmemcache_14-2.3.0-9PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmemcache_14-2.3.0-9PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgmemcache` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.1 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg13+1_amd64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.1 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg13+1_arm64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.7 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.5 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgmemcache` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.2 KiB | [postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgmemcache/postgresql-14-pgmemcache_2.3.0-15.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ohmu/pgmemcache" title="Repository" icon="github" subtitle="github.com/ohmu/pgmemcache" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmemcache;		# install via package name, for the active PG version

pig install pgmemcache -v 18;   # install for PG 18
pig install pgmemcache -v 17;   # install for PG 17
pig install pgmemcache -v 16;   # install for PG 16
pig install pgmemcache -v 15;   # install for PG 15
pig install pgmemcache -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmemcache;
```



## Usage

> [pgmemcache: memcached interface](https://github.com/ohmu/pgmemcache)

Provides PostgreSQL user-defined functions for interacting with memcached servers.

### Enabling

```sql
CREATE EXTENSION pgmemcache;
```

Configure default servers in `postgresql.conf`:

```ini
shared_preload_libraries = 'pgmemcache'
pgmemcache.default_servers = 'localhost:11211'
pgmemcache.default_behavior = 'DEAD_TIMEOUT:2'
```

### Server Management

```sql
SELECT memcache_server_add('localhost:11211');
SELECT memcache_server_add('cache-host');  -- uses default port 11211
```

### Setting and Getting Values

```sql
-- Set a key (overwrites if exists)
SELECT memcache_set('user:1:name', 'John Doe');
SELECT memcache_set('session:abc', 'data', now() + interval '1 hour');

-- Add a key (fails if exists)
SELECT memcache_add('user:2:name', 'Jane Doe');
SELECT memcache_add('temp_key', 'value', interval '5 minutes');

-- Replace (fails if key doesn't exist)
SELECT memcache_replace('user:1:name', 'John Smith');

-- Get a value
SELECT memcache_get('user:1:name');  -- returns text or NULL

-- Get multiple values
SELECT key, value FROM memcache_get_multi('{key1,key2,key3}'::text[]);
```

### Atomic Counters

```sql
SELECT memcache_incr('counter');        -- increment by 1
SELECT memcache_incr('counter', 5);     -- increment by 5
SELECT memcache_decr('counter');        -- decrement by 1
SELECT memcache_decr('counter', 3);     -- decrement by 3
```

### Deleting and Flushing

```sql
SELECT memcache_delete('user:1:name');
SELECT memcache_flush_all();  -- flush all servers
```

### Statistics

```sql
SELECT memcache_stats();  -- returns stats from all servers
```

### Trigger Example

Invalidate cache on table updates:

```sql
CREATE OR REPLACE FUNCTION auth_passwd_upd()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    IF OLD.passwd <> NEW.passwd THEN
        PERFORM memcache_delete('user_id_' || NEW.user_id || '_password');
    END IF;
    RETURN NEW;
END;
$$;

CREATE TRIGGER auth_passwd_upd_trg AFTER UPDATE ON passwd
    FOR EACH ROW EXECUTE PROCEDURE auth_passwd_upd();
```
