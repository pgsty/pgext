---
title: "pgmp"
linkTitle: "pgmp"
description: "Multiple Precision Arithmetic extension"
weight: 3700
categories: ["TYPE"]
width: full
---

[**pgmp**](https://github.com/dvarrazzo/pgmp/) : Multiple Precision Arithmetic extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3700** | {{< badge content="pgmp" link="https://github.com/dvarrazzo/pgmp/" >}} | {{< ext "pgmp" >}} | `1.0.5` | {{< category "TYPE" >}} | {{< license "LGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "numeral" >}} {{< ext "unit" >}} {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |

> [!Note] missing pg14 on el pgdg repo


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmp` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.5` | {{< bg "18" "pgmp_18" "green" >}} {{< bg "17" "pgmp_17" "green" >}} {{< bg "16" "pgmp_16" "green" >}} {{< bg "15" "pgmp_15" "green" >}} {{< bg "14" "pgmp_14" "green" >}} | `pgmp_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.5` | {{< bg "18" "postgresql-18-pgmp" "green" >}} {{< bg "17" "postgresql-17-pgmp" "green" >}} {{< bg "16" "postgresql-16-pgmp" "green" >}} {{< bg "15" "postgresql-15-pgmp" "green" >}} {{< bg "14" "postgresql-14-pgmp" "green" >}} | `postgresql-$v-pgmp` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgmp_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pgmp_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.5" "pgmp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "pgmp_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pgmp : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pgmp : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_18` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgmp_18-1.0.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmp_18-1.0.5-4PGDG.rhel8.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmp_18-1.0.5-4PGDG.rhel8.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmp_18-1.0.5-4PGDG.rhel9.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pgmp_18-1.0.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmp_18-1.0.5-4PGDG.rhel9.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pgmp_18-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmp_18-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmp_18-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.5 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.4 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.5 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.7 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 100.9 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.8 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.3 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.9 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_17` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [pgmp_17-1.0.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmp_17-1.0.5-3PGDG.rhel8.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [pgmp_17-1.0.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmp_17-1.0.5-3PGDG.rhel8.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pgmp_17-1.0.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmp_17-1.0.5-3PGDG.rhel9.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pgmp_17-1.0.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmp_17-1.0.5-3PGDG.rhel9.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pgmp_17-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmp_17-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_17-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmp_17-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.4 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.5 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.7 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.2 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.4 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.0 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.3 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.9 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_16` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pgmp_16-1.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmp_16-1.0.5-1PGDG.rhel8.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pgmp_16-1.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmp_16-1.0.5-1PGDG.rhel8.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmp_16-1.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmp_16-1.0.5-1PGDG.rhel9.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.0 KiB | [pgmp_16-1.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmp_16-1.0.5-1PGDG.rhel9.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pgmp_16-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmp_16-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_16-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmp_16-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.5 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.1 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.4 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.3 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.8 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.2 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.8 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.9 KiB | [pgmp_15-1.0.4-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmp_15-1.0.4-4.rhel8.x86_64.rpm) |
| `pgmp_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [pgmp_15-1.0.4-4.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmp_15-1.0.4-4.rhel8.aarch64.rpm) |
| `pgmp_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.4 KiB | [pgmp_15-1.0.4-4.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmp_15-1.0.4-4.rhel9.x86_64.rpm) |
| `pgmp_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.1 KiB | [pgmp_15-1.0.4-4.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmp_15-1.0.4-4.rhel9.aarch64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.4 KiB | [pgmp_15-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmp_15-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pgmp_15-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmp_15-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.3 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.7 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.5 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.5 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.6 KiB | [pgmp_14-1.0.4-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmp_14-1.0.4-4.rhel8.x86_64.rpm) |
| `pgmp_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.7 KiB | [pgmp_14-1.0.4-4.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmp_14-1.0.4-4.rhel8.aarch64.rpm) |
| `pgmp_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.1 KiB | [pgmp_14-1.0.4-4.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmp_14-1.0.4-4.rhel9.aarch64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.4 KiB | [pgmp_14-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmp_14-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pgmp_14-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmp_14-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.8 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.7 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.2 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 100.8 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.3 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.3 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dvarrazzo/pgmp/" title="Repository" icon="github" subtitle="github.com/dvarrazzo/pgmp/" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmp;		# install via package name, for the active PG version

pig install pgmp -v 18;   # install for PG 18
pig install pgmp -v 17;   # install for PG 17
pig install pgmp -v 16;   # install for PG 16
pig install pgmp -v 15;   # install for PG 15
pig install pgmp -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmp;
```



## Usage

> [pgmp: multiple precision arithmetic (GMP) for PostgreSQL](https://github.com/dvarrazzo/pgmp/)

The `pgmp` extension integrates the GNU Multiple Precision (GMP) library into PostgreSQL, providing arbitrary-precision integer and rational number types.

```sql
CREATE EXTENSION pgmp;
```

### Data Types

- **`mpz`**: Arbitrary-size integers, limited only by PostgreSQL's 1GB varlena maximum
- **`mpq`**: Arbitrary-precision rational numbers for exact fractional arithmetic

### Basic Usage

```sql
-- Arbitrary precision integers
SELECT '123456789012345678901234567890'::mpz * 2;

-- Exact rational arithmetic (no rounding)
SELECT '1'::mpq / '3'::mpq;  -- 1/3

-- Mixed arithmetic with native PostgreSQL types
SELECT 42::mpz + '100'::mpz;
```

### Operators

Standard arithmetic operators (`+`, `-`, `*`, `/`, `%`) and comparison operators (`=`, `<>`, `<`, `>`, `<=`, `>=`) are supported for both `mpz` and `mpq`.

### Functions

The extension exposes all GMP library functions for these types, including:

- Prime number functions
- Random number generation
- Factorization
- GCD, LCM, and other number theory functions

### Index Support

Both `mpz` and `mpq` support btree and hash indexes for efficient queries and sorting.

Full documentation: https://www.varrazzo.com/pgmp/
