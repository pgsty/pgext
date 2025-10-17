---
title: "pgfaceting"
linkTitle: "pgfaceting"
description: "fast faceting queries using an inverted index"
weight: 3580
categories: ["Type"]
width: full
---

fast faceting queries using an inverted index

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3580** | {{< badge content="pgfaceting" link="https://github.com/cybertec-postgresql/pgfaceting" >}} | {{< ext "pgfaceting" "pgfaceting" >}} | `0.2.0` | {{< category "TYPE" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "roaringbitmap" >}} |
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgfaceting" >}} | `0.2.0` | {{< badge content="18" color="red" alt="pgfaceting_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgfaceting_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgfaceting" >}} | `0.2.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgfaceting` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgfaceting_18" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_17" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_16" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_15" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_14" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgfaceting_18" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_17" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_16" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_15" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pgfaceting_14" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgfaceting_18" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_17" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_16" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_15" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_14" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgfaceting_18" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_17" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_16" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_15" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pgfaceting_14" "0.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-17-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-16-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-15-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} | {{< pkg "postgresql-14-pgfaceting" "0.2.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_18` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_18` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgfaceting_18-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_18` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_18` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_18` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgfaceting_18` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_18` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgfaceting_18-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_18` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-18-pgfaceting` | 0.2.0 | `d12.x86_64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-18-pgfaceting` | 0.2.0 | `d12.aarch64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-18-pgfaceting` | 0.2.0 | `u22.aarch64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-18-pgfaceting` | 0.2.0 | `u22.x86_64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-18-pgfaceting` | 0.2.0 | `u24.aarch64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |
| `postgresql-18-pgfaceting` | 0.2.0 | `u24.x86_64` | pgdg | 9.7 KiB | [postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-18-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_17` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_17` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_17` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgfaceting_17-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_17` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_17` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgfaceting_17` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgfaceting_17` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_17` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgfaceting_17-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pgfaceting` | 0.2.0 | `d12.aarch64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-17-pgfaceting` | 0.2.0 | `d12.x86_64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-17-pgfaceting` | 0.2.0 | `u22.aarch64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-17-pgfaceting` | 0.2.0 | `u22.x86_64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-17-pgfaceting` | 0.2.0 | `u24.x86_64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |
| `postgresql-17-pgfaceting` | 0.2.0 | `u24.aarch64` | pgdg | 9.7 KiB | [postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-17-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_16` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_16` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_16` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgfaceting_16-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_16` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_16` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_16` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgfaceting_16-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_16` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgfaceting_16` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pgfaceting` | 0.2.0 | `d12.x86_64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-16-pgfaceting` | 0.2.0 | `d12.aarch64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-16-pgfaceting` | 0.2.0 | `u22.x86_64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-16-pgfaceting` | 0.2.0 | `u22.aarch64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-16-pgfaceting` | 0.2.0 | `u24.x86_64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |
| `postgresql-16-pgfaceting` | 0.2.0 | `u24.aarch64` | pgdg | 9.7 KiB | [postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-16-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_15` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_15` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_15` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfaceting_15-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_15` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_15` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_15` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgfaceting_15` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgfaceting_15-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_15` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pgfaceting` | 0.2.0 | `d12.aarch64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-15-pgfaceting` | 0.2.0 | `d12.x86_64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-15-pgfaceting` | 0.2.0 | `u22.aarch64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-15-pgfaceting` | 0.2.0 | `u22.x86_64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-15-pgfaceting` | 0.2.0 | `u24.aarch64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |
| `postgresql-15-pgfaceting` | 0.2.0 | `u24.x86_64` | pgdg | 9.7 KiB | [postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-15-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_14` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_14` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_14` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfaceting_14-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_14` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_14` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_14` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgfaceting_14-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_14` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgfaceting_14` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-pgfaceting` | 0.2.0 | `d12.aarch64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-14-pgfaceting` | 0.2.0 | `d12.x86_64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg12+1_all.deb) |
| `postgresql-14-pgfaceting` | 0.2.0 | `u22.aarch64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-14-pgfaceting` | 0.2.0 | `u22.x86_64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg22.04+1_all.deb) |
| `postgresql-14-pgfaceting` | 0.2.0 | `u24.x86_64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |
| `postgresql-14-pgfaceting` | 0.2.0 | `u24.aarch64` | pgdg | 9.7 KiB | [postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfaceting/postgresql-14-pgfaceting_0.2.0-5.pgdg24.04+1_all.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgfaceting_13` | 0.2.0 | `el8.aarch64` | pigsty | 14.6 KiB | [pgfaceting_13-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfaceting_13-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgfaceting_13` | 0.2.0 | `el8.x86_64` | pgdg | 15.5 KiB | [pgfaceting_13-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgfaceting_13-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_13` | 0.2.0 | `el8.x86_64` | pigsty | 14.6 KiB | [pgfaceting_13-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfaceting_13-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgfaceting_13` | 0.2.0 | `el8.aarch64` | pgdg | 15.5 KiB | [pgfaceting_13-0.2.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgfaceting_13-0.2.0-1PGDG.rhel8.noarch.rpm) |
| `pgfaceting_13` | 0.2.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgfaceting_13-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgfaceting_13-0.2.0-1PGDG.rhel9.noarch.rpm) |
| `pgfaceting_13` | 0.2.0 | `el9.x86_64` | pigsty | 14.4 KiB | [pgfaceting_13-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfaceting_13-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgfaceting_13` | 0.2.0 | `el9.aarch64` | pigsty | 14.4 KiB | [pgfaceting_13-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfaceting_13-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgfaceting_13` | 0.2.0 | `el9.aarch64` | pgdg | 15.3 KiB | [pgfaceting_13-0.2.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgfaceting_13-0.2.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pgfaceting" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pgfaceting" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgfaceting-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgfaceting; # get pgfaceting source code
pig build dep pgfaceting; # install build dependencies
pig build pkg pgfaceting; # build extension rpm or deb
pig build ext pgfaceting; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgfaceting; # install by extension name, for the current active PG version
pig ext install pgfaceting; # install via package alias, for the active PG version
pig ext install pgfaceting -v 18;   # install for PG 18
pig ext install pgfaceting -v 17;   # install for PG 17
pig ext install pgfaceting -v 16;   # install for PG 16
pig ext install pgfaceting -v 15;   # install for PG 15
pig ext install pgfaceting -v 14;   # install for PG 14
pig ext install pgfaceting -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgfaceting CASCADE SCHEMA faceting;
```

