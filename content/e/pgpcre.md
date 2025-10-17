---
title: "pgpcre"
linkTitle: "pgpcre"
description: "Perl Compatible Regular Expression functions"
weight: 4230
categories: ["Util"]
width: full
---

Perl Compatible Regular Expression functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4230** | {{< badge content="pgpcre" link="https://github.com/petere/pgpcre" >}} | {{< ext "pgpcre" "pgpcre" >}} | `1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "icu_ext" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgpcre" >}} | `1` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgpcre_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgpcre" >}} | `1` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgpcre` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgpcre_18" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgpcre_18-0.20190509-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpcre_17" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgpcre_17-0.20190509-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpcre_16" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgpcre_16-0.20190509-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpcre_15" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgpcre_15-0.20190509-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgpcre_14" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgpcre_14-0.20190509-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgpcre_18" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgpcre_18-0.20190509-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpcre_17" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgpcre_17-0.20190509-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpcre_16" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgpcre_16-0.20190509-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpcre_15" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgpcre_15-0.20190509-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgpcre_14" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgpcre_14-0.20190509-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgpcre_18" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgpcre_18-0.20190509-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpcre_17" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgpcre_17-0.20190509-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpcre_16" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgpcre_16-0.20190509-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpcre_15" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgpcre_15-0.20190509-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgpcre_14" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgpcre_14-0.20190509-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgpcre_18" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgpcre_18-0.20190509-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpcre_17" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgpcre_17-0.20190509-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpcre_16" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgpcre_16-0.20190509-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpcre_15" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgpcre_15-0.20190509-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgpcre_14" "0.20190509" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgpcre_14-0.20190509-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgpcre" "0.20190509" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_18` | 0.20190509 | `el8.x86_64` | pgdg | 17.4 KiB | [pgpcre_18-0.20190509-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgpcre_18-0.20190509-3PGDG.rhel8.x86_64.rpm) |
| `pgpcre_18` | 0.20190509 | `el8.aarch64` | pgdg | 17.2 KiB | [pgpcre_18-0.20190509-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgpcre_18-0.20190509-3PGDG.rhel8.aarch64.rpm) |
| `pgpcre_18` | 0.20190509 | `el9.aarch64` | pgdg | 17.2 KiB | [pgpcre_18-0.20190509-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgpcre_18-0.20190509-3PGDG.rhel9.aarch64.rpm) |
| `pgpcre_18` | 0.20190509 | `el9.x86_64` | pgdg | 17.6 KiB | [pgpcre_18-0.20190509-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgpcre_18-0.20190509-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-18-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 18.1 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-18-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.0 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.3 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_17` | 1 | `el8.x86_64` | pigsty | 16.2 KiB | [pgpcre_17-1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_17-1-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_17` | 1 | `el8.aarch64` | pigsty | 16.1 KiB | [pgpcre_17-1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_17-1-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_17` | 0.20190509 | `el8.aarch64` | pgdg | 17.0 KiB | [pgpcre_17-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgpcre_17-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_17` | 0.20190509 | `el8.x86_64` | pgdg | 17.1 KiB | [pgpcre_17-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgpcre_17-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_17` | 1 | `el9.x86_64` | pigsty | 16.4 KiB | [pgpcre_17-1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_17-1-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_17` | 1 | `el9.aarch64` | pigsty | 16.1 KiB | [pgpcre_17-1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_17-1-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_17` | 0.20190509 | `el9.aarch64` | pgdg | 17.1 KiB | [pgpcre_17-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgpcre_17-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_17` | 0.20190509 | `el9.x86_64` | pgdg | 17.3 KiB | [pgpcre_17-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgpcre_17-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 18.0 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-17-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 18.1 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-17-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.6 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.9 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 18.2 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.1 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_16` | 1 | `el8.x86_64` | pigsty | 16.1 KiB | [pgpcre_16-1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_16-1-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_16` | 1 | `el8.aarch64` | pigsty | 16.1 KiB | [pgpcre_16-1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_16-1-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_16` | 0.20190509 | `el8.aarch64` | pgdg | 17.0 KiB | [pgpcre_16-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgpcre_16-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_16` | 0.20190509 | `el8.x86_64` | pgdg | 17.1 KiB | [pgpcre_16-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgpcre_16-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_16` | 1 | `el9.aarch64` | pigsty | 16.1 KiB | [pgpcre_16-1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_16-1-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_16` | 1 | `el9.x86_64` | pigsty | 16.4 KiB | [pgpcre_16-1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_16-1-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_16` | 0.20190509 | `el9.aarch64` | pgdg | 17.1 KiB | [pgpcre_16-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgpcre_16-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_16` | 0.20190509 | `el9.x86_64` | pgdg | 17.3 KiB | [pgpcre_16-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgpcre_16-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 18.0 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-16-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 18.1 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-16-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.9 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.6 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 18.2 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.1 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_15` | 1 | `el8.aarch64` | pigsty | 16.1 KiB | [pgpcre_15-1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_15-1-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_15` | 1 | `el8.x86_64` | pigsty | 16.2 KiB | [pgpcre_15-1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_15-1-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_15` | 0.20190509 | `el8.x86_64` | pgdg | 17.1 KiB | [pgpcre_15-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgpcre_15-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_15` | 0.20190509 | `el8.aarch64` | pgdg | 17.0 KiB | [pgpcre_15-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgpcre_15-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_15` | 1 | `el9.aarch64` | pigsty | 16.1 KiB | [pgpcre_15-1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_15-1-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_15` | 1 | `el9.x86_64` | pigsty | 16.4 KiB | [pgpcre_15-1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_15-1-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_15` | 0.20190509 | `el9.x86_64` | pgdg | 17.3 KiB | [pgpcre_15-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgpcre_15-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_15` | 0.20190509 | `el9.aarch64` | pgdg | 17.1 KiB | [pgpcre_15-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgpcre_15-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 18.0 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-15-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 18.1 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-15-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.9 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.6 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 18.2 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.1 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_14` | 1 | `el8.x86_64` | pigsty | 16.1 KiB | [pgpcre_14-1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_14-1-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_14` | 1 | `el8.aarch64` | pigsty | 16.1 KiB | [pgpcre_14-1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_14-1-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_14` | 0.20190509 | `el8.aarch64` | pgdg | 17.0 KiB | [pgpcre_14-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgpcre_14-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_14` | 0.20190509 | `el8.x86_64` | pgdg | 17.0 KiB | [pgpcre_14-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgpcre_14-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_14` | 1 | `el9.x86_64` | pigsty | 16.4 KiB | [pgpcre_14-1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_14-1-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_14` | 1 | `el9.aarch64` | pigsty | 16.1 KiB | [pgpcre_14-1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_14-1-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_14` | 0.20190509 | `el9.x86_64` | pgdg | 17.3 KiB | [pgpcre_14-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgpcre_14-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_14` | 0.20190509 | `el9.aarch64` | pgdg | 17.1 KiB | [pgpcre_14-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgpcre_14-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 18.0 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-14-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 18.0 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-14-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.8 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.5 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgpcre_13` | 1 | `el8.aarch64` | pigsty | 16.1 KiB | [pgpcre_13-1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_13-1-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_13` | 1 | `el8.x86_64` | pigsty | 16.1 KiB | [pgpcre_13-1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_13-1-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_13` | 0.20190509 | `el8.x86_64` | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgpcre_13-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_13` | 0.20190509 | `el8.aarch64` | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgpcre_13-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_13` | 1 | `el9.aarch64` | pigsty | 16.1 KiB | [pgpcre_13-1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_13-1-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_13` | 1 | `el9.x86_64` | pigsty | 16.4 KiB | [pgpcre_13-1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_13-1-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_13` | 0.20190509 | `el9.aarch64` | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgpcre_13-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_13` | 0.20190509 | `el9.x86_64` | pgdg | 17.3 KiB | [pgpcre_13-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgpcre_13-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-pgpcre` | 0.20190509 | `d12.aarch64` | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-13-pgpcre` | 0.20190509 | `d12.x86_64` | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-13-pgpcre` | 0.20190509 | `u22.x86_64` | pgdg | 18.4 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgpcre` | 0.20190509 | `u22.aarch64` | pgdg | 18.2 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgpcre` | 0.20190509 | `u24.x86_64` | pgdg | 18.0 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgpcre` | 0.20190509 | `u24.aarch64` | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pgpcre" title="Repository" icon="github" subtitle="github.com/petere/pgpcre" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgpcre-1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgpcre; # get pgpcre source code
pig build dep pgpcre; # install build dependencies
pig build pkg pgpcre; # build extension rpm or deb
pig build ext pgpcre; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgpcre; # install by extension name, for the current active PG version
pig ext install pgpcre; # install via package alias, for the active PG version
pig ext install pgpcre -v 18;   # install for PG 18
pig ext install pgpcre -v 17;   # install for PG 17
pig ext install pgpcre -v 16;   # install for PG 16
pig ext install pgpcre -v 15;   # install for PG 15
pig ext install pgpcre -v 14;   # install for PG 14
pig ext install pgpcre -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgpcre;
```

