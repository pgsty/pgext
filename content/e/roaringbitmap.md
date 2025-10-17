---
title: "roaringbitmap"
linkTitle: "roaringbitmap"
description: "support for Roaring Bitmaps"
weight: 3570
categories: ["Type"]
width: full
---

support for Roaring Bitmaps

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3570** | {{< badge content="roaringbitmap" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< ext "roaringbitmap" "roaringbitmap" >}} | `0.5.4` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pgfaceting" >}} |
|   **See Also**    | {{< ext "rum" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/roaringbitmap" >}} | `0.5.4` | {{< badge content="18" color="red" alt="pg_roaringbitmap_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_roaringbitmap_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/roaringbitmap" >}} | `0.5.4` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-roaringbitmap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_roaringbitmap_18" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_17" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_16" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_15" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_14" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_roaringbitmap_18" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_17" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_16" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_15" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_14" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_roaringbitmap_18" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_17" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_16" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_15" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_roaringbitmap_14" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_roaringbitmap_18" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_17" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_16" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_15" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_roaringbitmap_14" "0.5.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-roaringbitmap" "0.5.5" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_18` | 0.5.5 | `el8.x86_64` | pgdg | 159.1 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | 0.5.5 | `el8.aarch64` | pgdg | 133.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | 0.5.4 | `el8.x86_64` | pigsty | 103.9 KiB | [pg_roaringbitmap_18-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_18-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_18` | 0.5.4 | `el8.aarch64` | pigsty | 94.1 KiB | [pg_roaringbitmap_18-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_18-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_18` | 0.5.5 | `el9.x86_64` | pgdg | 84.8 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | 0.5.5 | `el9.aarch64` | pgdg | 72.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | 0.5.4 | `el9.x86_64` | pigsty | 68.0 KiB | [pg_roaringbitmap_18-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_18-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_18` | 0.5.4 | `el9.aarch64` | pigsty | 65.8 KiB | [pg_roaringbitmap_18-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_18-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 389.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 429.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 368.1 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 338.1 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 360.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 331.3 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_17` | 0.5.5 | `el8.aarch64` | pgdg | 133.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | 0.5.5 | `el8.x86_64` | pgdg | 159.1 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el8.x86_64` | pgdg | 107.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el8.x86_64` | pigsty | 103.9 KiB | [pg_roaringbitmap_17-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_17-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el8.aarch64` | pigsty | 94.1 KiB | [pg_roaringbitmap_17-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_17-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el8.aarch64` | pgdg | 98.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | 0.5.5 | `el9.aarch64` | pgdg | 72.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | 0.5.5 | `el9.x86_64` | pgdg | 84.8 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el9.aarch64` | pigsty | 66.0 KiB | [pg_roaringbitmap_17-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_17-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el9.x86_64` | pgdg | 72.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el9.x86_64` | pigsty | 68.0 KiB | [pg_roaringbitmap_17-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_17-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_17` | 0.5.4 | `el9.aarch64` | pgdg | 69.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 429.1 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 389.5 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 395.8 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 363.6 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 360.3 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 331.1 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_16` | 0.5.5 | `el8.x86_64` | pgdg | 159.1 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | 0.5.5 | `el8.aarch64` | pgdg | 133.7 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el8.aarch64` | pigsty | 94.0 KiB | [pg_roaringbitmap_16-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_16-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el8.aarch64` | pgdg | 98.1 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el8.x86_64` | pgdg | 107.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el8.x86_64` | pigsty | 103.9 KiB | [pg_roaringbitmap_16-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_16-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_16` | 0.5.5 | `el9.aarch64` | pgdg | 72.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.5 | `el9.x86_64` | pgdg | 84.8 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el9.aarch64` | pgdg | 69.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el9.aarch64` | pigsty | 65.9 KiB | [pg_roaringbitmap_16-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_16-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el9.x86_64` | pigsty | 68.0 KiB | [pg_roaringbitmap_16-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_16-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_16` | 0.5.4 | `el9.x86_64` | pgdg | 72.0 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 389.7 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 429.6 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 395.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 363.7 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 360.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 331.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_15` | 0.5.5 | `el8.x86_64` | pgdg | 162.2 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | 0.5.5 | `el8.aarch64` | pgdg | 135.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el8.x86_64` | pgdg | 109.1 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el8.aarch64` | pigsty | 95.8 KiB | [pg_roaringbitmap_15-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_15-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el8.aarch64` | pgdg | 99.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el8.x86_64` | pigsty | 105.0 KiB | [pg_roaringbitmap_15-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_15-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_15` | 0.5.5 | `el9.x86_64` | pgdg | 161.6 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | 0.5.5 | `el9.aarch64` | pgdg | 142.7 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el9.aarch64` | pigsty | 101.3 KiB | [pg_roaringbitmap_15-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_15-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el9.aarch64` | pgdg | 105.3 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el9.x86_64` | pgdg | 114.6 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | 0.5.4 | `el9.x86_64` | pigsty | 110.3 KiB | [pg_roaringbitmap_15-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_15-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 391.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 431.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 463.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 420.8 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 428.7 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 390.4 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_14` | 0.5.5 | `el8.aarch64` | pgdg | 135.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | 0.5.5 | `el8.x86_64` | pgdg | 162.2 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el8.aarch64` | pgdg | 99.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el8.x86_64` | pgdg | 109.1 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el8.x86_64` | pigsty | 105.0 KiB | [pg_roaringbitmap_14-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_14-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el8.aarch64` | pigsty | 95.8 KiB | [pg_roaringbitmap_14-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_14-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_14` | 0.5.5 | `el9.aarch64` | pgdg | 142.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | 0.5.5 | `el9.x86_64` | pgdg | 161.6 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el9.x86_64` | pigsty | 110.3 KiB | [pg_roaringbitmap_14-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_14-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el9.x86_64` | pgdg | 114.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el9.aarch64` | pgdg | 105.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | 0.5.4 | `el9.aarch64` | pigsty | 101.7 KiB | [pg_roaringbitmap_14-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_14-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 391.6 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 430.7 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 463.1 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 420.7 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 390.5 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 428.1 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_roaringbitmap_13` | 0.5.5 | `el8.x86_64` | pgdg | 161.8 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | 0.5.5 | `el8.aarch64` | pgdg | 135.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el8.aarch64` | pigsty | 95.8 KiB | [pg_roaringbitmap_13-0.5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_13-0.5.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el8.aarch64` | pgdg | 99.8 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el8.x86_64` | pigsty | 104.6 KiB | [pg_roaringbitmap_13-0.5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_13-0.5.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el8.x86_64` | pgdg | 108.7 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | 0.5.5 | `el9.aarch64` | pgdg | 142.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.5 | `el9.x86_64` | pgdg | 161.5 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el9.aarch64` | pigsty | 101.7 KiB | [pg_roaringbitmap_13-0.5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_13-0.5.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el9.x86_64` | pgdg | 114.3 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el9.aarch64` | pgdg | 105.2 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | 0.5.4 | `el9.x86_64` | pigsty | 110.3 KiB | [pg_roaringbitmap_13-0.5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_13-0.5.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `d12.x86_64` | pgdg | 430.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `d12.aarch64` | pgdg | 391.5 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `u22.x86_64` | pgdg | 463.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `u22.aarch64` | pgdg | 420.8 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `u24.x86_64` | pgdg | 427.9 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | 0.5.5 | `u24.aarch64` | pgdg | 390.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChenHuajun/pg_roaringbitmap" title="Repository" icon="github" subtitle="github.com/ChenHuajun/pg_roaringbitmap" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_roaringbitmap-0.5.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get roaringbitmap; # get roaringbitmap source code
pig build dep roaringbitmap; # install build dependencies
pig build pkg roaringbitmap; # build extension rpm or deb
pig build ext roaringbitmap; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install roaringbitmap; # install by extension name, for the current active PG version
pig ext install roaringbitmap; # install via package alias, for the active PG version
pig ext install roaringbitmap -v 18;   # install for PG 18
pig ext install roaringbitmap -v 17;   # install for PG 17
pig ext install roaringbitmap -v 16;   # install for PG 16
pig ext install roaringbitmap -v 15;   # install for PG 15
pig ext install roaringbitmap -v 14;   # install for PG 14
pig ext install roaringbitmap -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION roaringbitmap;
```

