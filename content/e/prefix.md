---
title: "prefix"
linkTitle: "prefix"
description: "Prefix Range module for PostgreSQL"
weight: 3500
categories: ["TYPE"]
width: full
---

[**pg_prefix**](https://github.com/dimitri/prefix)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3500** | {{< badge content="prefix" link="https://github.com/dimitri/prefix" >}} | {{< ext "prefix" "pg_prefix" >}} | `1.2.10` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "semver" >}} {{< ext "ltree" >}} {{< ext "citext" >}} {{< ext "pg_trgm" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/prefix" >}} | `1.2.10` | {{< bg "18" "prefix_18*" "green" >}} {{< bg "17" "prefix_17*" "green" >}} {{< bg "16" "prefix_16*" "green" >}} {{< bg "15" "prefix_15*" "green" >}} {{< bg "14" "prefix_14*" "green" >}} {{< bg "13" "prefix_13*" "green" >}} | `prefix_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/prefix" >}} | `1.2.10` | {{< bg "18" "postgresql-18-prefix" "green" >}} {{< bg "17" "postgresql-17-prefix" "green" >}} {{< bg "16" "postgresql-16-prefix" "green" >}} {{< bg "15" "postgresql-15-prefix" "green" >}} {{< bg "14" "postgresql-14-prefix" "green" >}} {{< bg "13" "postgresql-13-prefix" "green" >}} | `postgresql-$v-prefix` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.9" "prefix_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.10" "prefix_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.10" "prefix_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "prefix_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.2.10" "postgresql-18-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-17-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-16-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-15-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-14-prefix : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.10" "postgresql-13-prefix : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_18` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [prefix_18-1.2.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/prefix_18-1.2.10-1PIGSTY.el8.x86_64.rpm) |
| `prefix_18` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.9 KiB | [prefix_18-1.2.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/prefix_18-1.2.10-1PIGSTY.el8.aarch64.rpm) |
| `prefix_18` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.6 KiB | [prefix_18-1.2.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/prefix_18-1.2.10-1PIGSTY.el9.x86_64.rpm) |
| `prefix_18` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.9 KiB | [prefix_18-1.2.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/prefix_18-1.2.10-1PIGSTY.el9.aarch64.rpm) |
| `prefix_18` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [prefix_18-1.2.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/prefix_18-1.2.10-1PIGSTY.el10.x86_64.rpm) |
| `prefix_18` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.1 KiB | [prefix_18-1.2.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/prefix_18-1.2.10-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.2 KiB | [postgresql-18-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.2 KiB | [postgresql-18-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.3 KiB | [postgresql-18-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.3 KiB | [postgresql-18-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 42.9 KiB | [postgresql-18-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 41.7 KiB | [postgresql-18-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.2 KiB | [postgresql-18-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.2 KiB | [postgresql-18-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-18-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_17` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [prefix_17-1.2.10-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/prefix_17-1.2.10-2PGDG.rhel8.x86_64.rpm) |
| `prefix_17` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.3 KiB | [prefix_17-1.2.10-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/prefix_17-1.2.10-2PGDG.rhel8.aarch64.rpm) |
| `prefix_17` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.6 KiB | [prefix_17-1.2.10-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/prefix_17-1.2.10-2PGDG.rhel9.x86_64.rpm) |
| `prefix_17` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.7 KiB | [prefix_17-1.2.10-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/prefix_17-1.2.10-2PGDG.rhel9.aarch64.rpm) |
| `prefix_17` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [prefix_17-1.2.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/prefix_17-1.2.10-3PGDG.rhel10.x86_64.rpm) |
| `prefix_17` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [prefix_17-1.2.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/prefix_17-1.2.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.2 KiB | [postgresql-17-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.2 KiB | [postgresql-17-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.3 KiB | [postgresql-17-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.3 KiB | [postgresql-17-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.9 KiB | [postgresql-17-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-17-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.2 KiB | [postgresql-17-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.3 KiB | [postgresql-17-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-17-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_16` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.0 KiB | [prefix_16-1.2.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/prefix_16-1.2.10-1PGDG.rhel8.x86_64.rpm) |
| `prefix_16` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.2 KiB | [prefix_16-1.2.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/prefix_16-1.2.10-1PGDG.rhel8.aarch64.rpm) |
| `prefix_16` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [prefix_16-1.2.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/prefix_16-1.2.10-1PGDG.rhel9.x86_64.rpm) |
| `prefix_16` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.4 KiB | [prefix_16-1.2.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/prefix_16-1.2.10-1PGDG.rhel9.aarch64.rpm) |
| `prefix_16` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [prefix_16-1.2.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/prefix_16-1.2.10-3PGDG.rhel10.x86_64.rpm) |
| `prefix_16` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [prefix_16-1.2.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/prefix_16-1.2.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.3 KiB | [postgresql-16-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.3 KiB | [postgresql-16-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.3 KiB | [postgresql-16-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.3 KiB | [postgresql-16-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 44.0 KiB | [postgresql-16-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-16-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.2 KiB | [postgresql-16-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.3 KiB | [postgresql-16-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-16-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_15` | `1.2.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [prefix_15-1.2.9-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/prefix_15-1.2.9-3.rhel8.x86_64.rpm) |
| `prefix_15` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.0 KiB | [prefix_15-1.2.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/prefix_15-1.2.10-1PGDG.rhel8.x86_64.rpm) |
| `prefix_15` | `1.2.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.9 KiB | [prefix_15-1.2.9-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/prefix_15-1.2.9-3.rhel8.aarch64.rpm) |
| `prefix_15` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.2 KiB | [prefix_15-1.2.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/prefix_15-1.2.10-1PGDG.rhel8.aarch64.rpm) |
| `prefix_15` | `1.2.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.5 KiB | [prefix_15-1.2.9-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/prefix_15-1.2.9-3.rhel9.x86_64.rpm) |
| `prefix_15` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [prefix_15-1.2.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/prefix_15-1.2.10-1PGDG.rhel9.x86_64.rpm) |
| `prefix_15` | `1.2.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.3 KiB | [prefix_15-1.2.9-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/prefix_15-1.2.9-3.rhel9.aarch64.rpm) |
| `prefix_15` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.4 KiB | [prefix_15-1.2.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/prefix_15-1.2.10-1PGDG.rhel9.aarch64.rpm) |
| `prefix_15` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [prefix_15-1.2.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/prefix_15-1.2.10-3PGDG.rhel10.x86_64.rpm) |
| `prefix_15` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [prefix_15-1.2.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/prefix_15-1.2.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.3 KiB | [postgresql-15-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.2 KiB | [postgresql-15-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.3 KiB | [postgresql-15-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.4 KiB | [postgresql-15-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 44.0 KiB | [postgresql-15-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.9 KiB | [postgresql-15-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.2 KiB | [postgresql-15-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.3 KiB | [postgresql-15-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-15-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_14` | `1.2.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.7 KiB | [prefix_14-1.2.9-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/prefix_14-1.2.9-3.rhel8.x86_64.rpm) |
| `prefix_14` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.0 KiB | [prefix_14-1.2.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/prefix_14-1.2.10-1PGDG.rhel8.x86_64.rpm) |
| `prefix_14` | `1.2.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.7 KiB | [prefix_14-1.2.9-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/prefix_14-1.2.9-3.rhel8.aarch64.rpm) |
| `prefix_14` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.2 KiB | [prefix_14-1.2.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/prefix_14-1.2.10-1PGDG.rhel8.aarch64.rpm) |
| `prefix_14` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [prefix_14-1.2.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/prefix_14-1.2.10-1PGDG.rhel9.x86_64.rpm) |
| `prefix_14` | `1.2.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.1 KiB | [prefix_14-1.2.9-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/prefix_14-1.2.9-3.rhel9.aarch64.rpm) |
| `prefix_14` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.4 KiB | [prefix_14-1.2.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/prefix_14-1.2.10-1PGDG.rhel9.aarch64.rpm) |
| `prefix_14` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [prefix_14-1.2.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/prefix_14-1.2.10-3PGDG.rhel10.x86_64.rpm) |
| `prefix_14` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [prefix_14-1.2.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/prefix_14-1.2.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.2 KiB | [postgresql-14-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.2 KiB | [postgresql-14-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.3 KiB | [postgresql-14-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.3 KiB | [postgresql-14-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.8 KiB | [postgresql-14-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-14-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.2 KiB | [postgresql-14-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.2 KiB | [postgresql-14-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-14-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `prefix_13` | `1.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [prefix_13-1.2.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/prefix_13-1.2.10-1PGDG.rhel8.x86_64.rpm) |
| `prefix_13` | `1.2.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.5 KiB | [prefix_13-1.2.9-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/prefix_13-1.2.9-3.rhel8.aarch64.rpm) |
| `prefix_13` | `1.2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.2 KiB | [prefix_13-1.2.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/prefix_13-1.2.10-1PGDG.rhel8.aarch64.rpm) |
| `prefix_13` | `1.2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [prefix_13-1.2.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/prefix_13-1.2.10-1PGDG.rhel9.x86_64.rpm) |
| `prefix_13` | `1.2.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.9 KiB | [prefix_13-1.2.9-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/prefix_13-1.2.9-3.rhel9.aarch64.rpm) |
| `prefix_13` | `1.2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.4 KiB | [prefix_13-1.2.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/prefix_13-1.2.10-1PGDG.rhel9.aarch64.rpm) |
| `prefix_13` | `1.2.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.3 KiB | [prefix_13-1.2.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/prefix_13-1.2.10-3PGDG.rhel10.x86_64.rpm) |
| `prefix_13` | `1.2.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [prefix_13-1.2.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/prefix_13-1.2.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-prefix` | `1.2.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.3 KiB | [postgresql-13-prefix_1.2.10-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg12+1_amd64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.1 KiB | [postgresql-13-prefix_1.2.10-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg12+1_arm64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.4 KiB | [postgresql-13-prefix_1.2.10-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg13+1_amd64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.2 KiB | [postgresql-13-prefix_1.2.10-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg13+1_arm64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.7 KiB | [postgresql-13-prefix_1.2.10-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.9 KiB | [postgresql-13-prefix_1.2.10-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.3 KiB | [postgresql-13-prefix_1.2.10-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-prefix` | `1.2.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 39.2 KiB | [postgresql-13-prefix_1.2.10-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/prefix/postgresql-13-prefix_1.2.10-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dimitri/prefix" title="Repository" icon="github" subtitle="github.com/dimitri/prefix" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="prefix-1.2.10.tar.gz" >}}
{{< /cards >}}


```bash
pig build get prefix; # get prefix source code
pig build dep prefix; # install build dependencies
pig build pkg prefix; # build extension rpm or deb
pig build ext prefix; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install prefix; # install by extension name, for the current active PG version
pig ext install pg_prefix; # install via package alias, for the active PG version
pig ext install prefix -v 18;   # install for PG 18
pig ext install prefix -v 17;   # install for PG 17
pig ext install prefix -v 16;   # install for PG 16
pig ext install prefix -v 15;   # install for PG 15
pig ext install prefix -v 14;   # install for PG 14
pig ext install prefix -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION prefix;
```

