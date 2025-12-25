---
title: "hll"
linkTitle: "hll"
description: "type for storing hyperloglog data"
weight: 2740
categories: ["FEAT"]
width: full
---

[**hll**](https://github.com/citusdata/postgresql-hll) : type for storing hyperloglog data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2740** | {{< badge content="hll" link="https://github.com/citusdata/postgresql-hll" >}} | {{< ext "hll" >}} | `2.19` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "bloom" >}} {{< ext "roaringbitmap" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "citus" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.19` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `hll` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.19` | {{< bg "18" "hll_18" "green" >}} {{< bg "17" "hll_17" "green" >}} {{< bg "16" "hll_16" "green" >}} {{< bg "15" "hll_15" "green" >}} {{< bg "14" "hll_14" "green" >}} {{< bg "13" "hll_13" "green" >}} | `hll_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.19` | {{< bg "18" "postgresql-18-hll" "green" >}} {{< bg "17" "postgresql-17-hll" "green" >}} {{< bg "16" "postgresql-16-hll" "green" >}} {{< bg "15" "postgresql-15-hll" "green" >}} {{< bg "14" "postgresql-14-hll" "green" >}} {{< bg "13" "postgresql-13-hll" "green" >}} | `postgresql-$v-hll` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.19" "hll_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.18" "hll_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.19" "postgresql-18-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-17-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-16-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-15-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-14-hll : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.19" "postgresql-13-hll : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_18` | `2.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [hll_18-2.19-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/hll_18-2.19-1PGDG.rhel8.x86_64.rpm) |
| `hll_18` | `2.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [hll_18-2.19-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/hll_18-2.19-1PGDG.rhel8.aarch64.rpm) |
| `hll_18` | `2.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.1 KiB | [hll_18-2.19-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/hll_18-2.19-1PGDG.rhel9.x86_64.rpm) |
| `hll_18` | `2.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [hll_18-2.19-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/hll_18-2.19-1PGDG.rhel9.aarch64.rpm) |
| `hll_18` | `2.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.5 KiB | [hll_18-2.19-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/hll_18-2.19-1PGDG.rhel10.x86_64.rpm) |
| `hll_18` | `2.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.9 KiB | [hll_18-2.19-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/hll_18-2.19-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.5 KiB | [postgresql-18-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-18-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 75.6 KiB | [postgresql-18-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-18-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.4 KiB | [postgresql-18-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-18-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.6 KiB | [postgresql-18-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-18-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.3 KiB | [postgresql-18-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 75.2 KiB | [postgresql-18-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.0 KiB | [postgresql-18-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.0 KiB | [postgresql-18-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-18-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_17` | `2.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [hll_17-2.18-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hll_17-2.18-2PGDG.rhel8.x86_64.rpm) |
| `hll_17` | `2.18` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [hll_17-2.18-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hll_17-2.18-2PGDG.rhel8.aarch64.rpm) |
| `hll_17` | `2.18` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.9 KiB | [hll_17-2.18-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hll_17-2.18-2PGDG.rhel9.x86_64.rpm) |
| `hll_17` | `2.18` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [hll_17-2.18-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hll_17-2.18-2PGDG.rhel9.aarch64.rpm) |
| `hll_17` | `2.18` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [hll_17-2.18-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/hll_17-2.18-3PGDG.rhel10.x86_64.rpm) |
| `hll_17` | `2.18` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.6 KiB | [hll_17-2.18-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/hll_17-2.18-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.6 KiB | [postgresql-17-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-17-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 75.5 KiB | [postgresql-17-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-17-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.5 KiB | [postgresql-17-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-17-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.6 KiB | [postgresql-17-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-17-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 82.7 KiB | [postgresql-17-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 81.6 KiB | [postgresql-17-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.1 KiB | [postgresql-17-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.0 KiB | [postgresql-17-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-17-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_16` | `2.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [hll_16-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hll_16-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_16` | `2.18` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [hll_16-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hll_16-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_16` | `2.18` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [hll_16-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hll_16-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_16` | `2.18` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [hll_16-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hll_16-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_16` | `2.18` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [hll_16-2.18-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/hll_16-2.18-3PGDG.rhel10.x86_64.rpm) |
| `hll_16` | `2.18` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.6 KiB | [hll_16-2.18-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/hll_16-2.18-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.6 KiB | [postgresql-16-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-16-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 75.5 KiB | [postgresql-16-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-16-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.5 KiB | [postgresql-16-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-16-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.5 KiB | [postgresql-16-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-16-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 82.3 KiB | [postgresql-16-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 81.1 KiB | [postgresql-16-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.1 KiB | [postgresql-16-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.0 KiB | [postgresql-16-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-16-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_15` | `2.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.0 KiB | [hll_15-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hll_15-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_15` | `2.17` | [el8.x86_64](/os/el8.x86_64) | pgdg | 89.5 KiB | [hll_15-2.17-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hll_15-2.17-1.rhel8.x86_64.rpm) |
| `hll_15` | `2.18` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [hll_15-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hll_15-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_15` | `2.17` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.2 KiB | [hll_15-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hll_15-2.17-1.rhel8.aarch64.rpm) |
| `hll_15` | `2.18` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.9 KiB | [hll_15-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hll_15-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_15` | `2.17` | [el9.x86_64](/os/el9.x86_64) | pgdg | 90.6 KiB | [hll_15-2.17-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hll_15-2.17-1.rhel9.x86_64.rpm) |
| `hll_15` | `2.18` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [hll_15-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hll_15-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_15` | `2.17` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.1 KiB | [hll_15-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hll_15-2.17-1.rhel9.aarch64.rpm) |
| `hll_15` | `2.18` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.1 KiB | [hll_15-2.18-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/hll_15-2.18-3PGDG.rhel10.x86_64.rpm) |
| `hll_15` | `2.18` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [hll_15-2.18-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/hll_15-2.18-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.0 KiB | [postgresql-15-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-15-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.0 KiB | [postgresql-15-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-15-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.9 KiB | [postgresql-15-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-15-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 76.0 KiB | [postgresql-15-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-15-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 83.1 KiB | [postgresql-15-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 82.2 KiB | [postgresql-15-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.1 KiB | [postgresql-15-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.5 KiB | [postgresql-15-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-15-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_14` | `2.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.0 KiB | [hll_14-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hll_14-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_14` | `2.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 90.0 KiB | [hll_14-2.16-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hll_14-2.16-1.rhel8.x86_64.rpm) |
| `hll_14` | `2.18` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [hll_14-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hll_14-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_14` | `2.17` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.0 KiB | [hll_14-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hll_14-2.17-1.rhel8.aarch64.rpm) |
| `hll_14` | `2.18` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.9 KiB | [hll_14-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hll_14-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_14` | `2.18` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [hll_14-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hll_14-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_14` | `2.17` | [el9.aarch64](/os/el9.aarch64) | pgdg | 90.9 KiB | [hll_14-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hll_14-2.17-1.rhel9.aarch64.rpm) |
| `hll_14` | `2.18` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.1 KiB | [hll_14-2.18-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/hll_14-2.18-3PGDG.rhel10.x86_64.rpm) |
| `hll_14` | `2.18` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [hll_14-2.18-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/hll_14-2.18-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.9 KiB | [postgresql-14-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-14-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 75.8 KiB | [postgresql-14-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-14-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.8 KiB | [postgresql-14-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-14-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.9 KiB | [postgresql-14-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-14-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 83.1 KiB | [postgresql-14-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 82.2 KiB | [postgresql-14-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.1 KiB | [postgresql-14-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.4 KiB | [postgresql-14-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-14-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hll_13` | `2.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [hll_13-2.18-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/hll_13-2.18-1PGDG.rhel8.x86_64.rpm) |
| `hll_13` | `2.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.4 KiB | [hll_13-2.15-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/hll_13-2.15-1.rhel8.x86_64.rpm) |
| `hll_13` | `2.18` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [hll_13-2.18-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/hll_13-2.18-1PGDG.rhel8.aarch64.rpm) |
| `hll_13` | `2.17` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.7 KiB | [hll_13-2.17-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/hll_13-2.17-1.rhel8.aarch64.rpm) |
| `hll_13` | `2.18` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.8 KiB | [hll_13-2.18-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/hll_13-2.18-1PGDG.rhel9.x86_64.rpm) |
| `hll_13` | `2.18` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [hll_13-2.18-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/hll_13-2.18-1PGDG.rhel9.aarch64.rpm) |
| `hll_13` | `2.17` | [el9.aarch64](/os/el9.aarch64) | pgdg | 90.7 KiB | [hll_13-2.17-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/hll_13-2.17-1.rhel9.aarch64.rpm) |
| `hll_13` | `2.18` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.1 KiB | [hll_13-2.18-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/hll_13-2.18-3PGDG.rhel10.x86_64.rpm) |
| `hll_13` | `2.18` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [hll_13-2.18-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/hll_13-2.18-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-hll` | `2.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.5 KiB | [postgresql-13-hll_2.19-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg12+1_amd64.deb) |
| `postgresql-13-hll` | `2.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 75.7 KiB | [postgresql-13-hll_2.19-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg12+1_arm64.deb) |
| `postgresql-13-hll` | `2.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.8 KiB | [postgresql-13-hll_2.19-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg13+1_amd64.deb) |
| `postgresql-13-hll` | `2.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.7 KiB | [postgresql-13-hll_2.19-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg13+1_arm64.deb) |
| `postgresql-13-hll` | `2.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 82.8 KiB | [postgresql-13-hll_2.19-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-hll` | `2.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 81.8 KiB | [postgresql-13-hll_2.19-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-hll` | `2.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.2 KiB | [postgresql-13-hll_2.19-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-hll` | `2.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.0 KiB | [postgresql-13-hll_2.19-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-hll/postgresql-13-hll_2.19-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/postgresql-hll" title="Repository" icon="github" subtitle="github.com/citusdata/postgresql-hll" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install hll;		# install via package name, for the active PG version

pig install hll -v 18;   # install for PG 18
pig install hll -v 17;   # install for PG 17
pig install hll -v 16;   # install for PG 16
pig install hll -v 15;   # install for PG 15
pig install hll -v 14;   # install for PG 14
pig install hll -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hll;
```
