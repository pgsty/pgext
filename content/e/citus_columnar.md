---
title: "citus_columnar"
linkTitle: "citus_columnar"
description: "Citus columnar storage engine"
weight: 2401
categories: ["Olap"]
width: full
---

Citus columnar storage engine

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2401** | {{< badge content="citus_columnar" link="https://github.com/citusdata/citus" >}} | {{< ext "citus_columnar" "citus" >}} | `13.2.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_parquet" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "orioledb" >}} |
|    **Siblings**   | {{< ext "citus" >}} |

> [!Note] conflict with hydra columnar


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/citus" >}} | `13.2.0` | {{< badge content="18" color="red" alt="citus_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `citus_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/citus" >}} | `13.2.0` | {{< badge content="18" color="red" alt="postgresql-18-citus" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "citus_18" >}}     | {{< pkg "citus_17" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "citus_16" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "citus_15" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "citus_14" "13.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-13.0.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "citus_18" >}}     | {{< pkg "citus_17" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "citus_16" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "citus_15" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "citus_14" "13.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-13.0.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "citus_18" >}}     | {{< pkg "citus_17" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "citus_16" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "citus_15" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "citus_14" "13.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-13.0.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "citus_18" >}}     | {{< pkg "citus_17" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "citus_16" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "citus_15" "13.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "citus_14" "13.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-13.0.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-citus" >}}     | {{< pkg "postgresql-17-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-citus" "13.2.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-citus" "13.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `citus_17` | 13.2.0 | `el8.x86_64` | pigsty | 960.9 KiB | [citus_17-13.2.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-13.2.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_17` | 13.2.0 | `el8.x86_64` | pgdg | 850.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.2.0 | `el8.aarch64` | pgdg | 802.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.2.0 | `el8.aarch64` | pigsty | 922.8 KiB | [citus_17-13.2.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-13.2.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el8.x86_64` | pgdg | 827.1 KiB | [citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.1.0 | `el8.aarch64` | pigsty | 898.9 KiB | [citus_17-13.1.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-13.1.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el8.aarch64` | pgdg | 781.5 KiB | [citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el8.x86_64` | pigsty | 935.4 KiB | [citus_17-13.1.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-13.1.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_17` | 13.0.4 | `el8.x86_64` | pgdg | 805.9 KiB | [citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.0.4 | `el8.aarch64` | pgdg | 761.4 KiB | [citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.0.3 | `el8.aarch64` | pgdg | 761.5 KiB | [citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.0.3 | `el8.x86_64` | pgdg | 805.9 KiB | [citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.0.2 | `el8.aarch64` | pgdg | 761.2 KiB | [citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.0.2 | `el8.x86_64` | pgdg | 805.7 KiB | [citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.0.0 | `el8.aarch64` | pgdg | 758.9 KiB | [citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | 13.0.0 | `el8.x86_64` | pgdg | 803.6 KiB | [citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | 13.2.0 | `el9.aarch64` | pgdg | 789.2 KiB | [citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.2.0 | `el9.x86_64` | pgdg | 815.8 KiB | [citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | 13.2.0 | `el9.x86_64` | pigsty | 842.0 KiB | [citus_17-13.2.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-13.2.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_17` | 13.2.0 | `el9.aarch64` | pigsty | 815.3 KiB | [citus_17-13.2.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-13.2.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el9.aarch64` | pgdg | 768.0 KiB | [citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el9.aarch64` | pigsty | 791.7 KiB | [citus_17-13.1.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-13.1.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_17` | 13.1.0 | `el9.x86_64` | pigsty | 818.3 KiB | [citus_17-13.1.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-13.1.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_17` | 13.1.0 | `el9.x86_64` | pgdg | 793.2 KiB | [citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | 13.0.4 | `el9.aarch64` | pgdg | 749.8 KiB | [citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.0.4 | `el9.x86_64` | pgdg | 774.1 KiB | [citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | 13.0.3 | `el9.x86_64` | pgdg | 774.6 KiB | [citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | 13.0.3 | `el9.aarch64` | pgdg | 750.0 KiB | [citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.0.2 | `el9.x86_64` | pgdg | 774.4 KiB | [citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | 13.0.2 | `el9.aarch64` | pgdg | 749.8 KiB | [citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.0.0 | `el9.aarch64` | pgdg | 746.8 KiB | [citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | 13.0.0 | `el9.x86_64` | pgdg | 772.0 KiB | [citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-citus` | 13.2.0 | `d12.x86_64` | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-citus` | 13.2.0 | `d12.aarch64` | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-citus` | 13.2.0 | `u22.x86_64` | pigsty | 3.4 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-17-citus` | 13.2.0 | `u22.aarch64` | pigsty | 3.3 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-17-citus` | 13.2.0 | `u24.aarch64` | pigsty | 2.7 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~noble_arm64.deb) |
| `postgresql-17-citus` | 13.2.0 | `u24.x86_64` | pigsty | 2.8 MiB | [postgresql-17-citus_13.2.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-9PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `citus_16` | 13.2.0 | `el8.x86_64` | pgdg | 843.6 KiB | [citus_16-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.2.0 | `el8.aarch64` | pigsty | 915.0 KiB | [citus_16-13.2.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-13.2.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_16` | 13.2.0 | `el8.x86_64` | pigsty | 953.5 KiB | [citus_16-13.2.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-13.2.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_16` | 13.2.0 | `el8.aarch64` | pgdg | 796.0 KiB | [citus_16-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.1.0 | `el8.aarch64` | pgdg | 774.6 KiB | [citus_16-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.1.0 | `el8.x86_64` | pgdg | 819.5 KiB | [citus_16-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.1.0 | `el8.x86_64` | pigsty | 927.4 KiB | [citus_16-13.1.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-13.1.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_16` | 13.1.0 | `el8.aarch64` | pigsty | 891.9 KiB | [citus_16-13.1.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-13.1.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_16` | 13.0.4 | `el8.x86_64` | pgdg | 800.7 KiB | [citus_16-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.0.4 | `el8.aarch64` | pgdg | 756.4 KiB | [citus_16-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.0.3 | `el8.aarch64` | pgdg | 756.4 KiB | [citus_16-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.0.3 | `el8.x86_64` | pgdg | 800.6 KiB | [citus_16-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.0.2 | `el8.aarch64` | pgdg | 756.2 KiB | [citus_16-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.0.2 | `el8.x86_64` | pgdg | 800.4 KiB | [citus_16-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.0.0 | `el8.x86_64` | pgdg | 798.7 KiB | [citus_16-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 13.0.0 | `el8.aarch64` | pgdg | 754.1 KiB | [citus_16-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.6 | `el8.x86_64` | pgdg | 797.3 KiB | [citus_16-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.6 | `el8.aarch64` | pgdg | 753.1 KiB | [citus_16-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.5 | `el8.aarch64` | pgdg | 751.8 KiB | [citus_16-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.5 | `el8.x86_64` | pgdg | 796.2 KiB | [citus_16-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.4 | `el8.aarch64` | pgdg | 751.8 KiB | [citus_16-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.4 | `el8.x86_64` | pgdg | 796.4 KiB | [citus_16-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.3 | `el8.aarch64` | pgdg | 751.6 KiB | [citus_16-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.3 | `el8.x86_64` | pgdg | 796.1 KiB | [citus_16-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.2 | `el8.x86_64` | pgdg | 795.2 KiB | [citus_16-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.2 | `el8.aarch64` | pgdg | 750.4 KiB | [citus_16-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.1 | `el8.aarch64` | pgdg | 750.4 KiB | [citus_16-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 12.1.1 | `el8.x86_64` | pgdg | 795.0 KiB | [citus_16-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.0 | `el8.x86_64` | pgdg | 794.9 KiB | [citus_16-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_16` | 12.1.0 | `el8.aarch64` | pgdg | 750.0 KiB | [citus_16-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_16` | 13.2.0 | `el9.aarch64` | pigsty | 808.1 KiB | [citus_16-13.2.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-13.2.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_16` | 13.2.0 | `el9.aarch64` | pgdg | 782.8 KiB | [citus_16-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.2.0 | `el9.x86_64` | pgdg | 809.6 KiB | [citus_16-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 13.2.0 | `el9.x86_64` | pigsty | 835.6 KiB | [citus_16-13.2.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-13.2.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_16` | 13.1.0 | `el9.aarch64` | pgdg | 761.8 KiB | [citus_16-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.1.0 | `el9.x86_64` | pgdg | 786.8 KiB | [citus_16-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 13.1.0 | `el9.aarch64` | pigsty | 785.6 KiB | [citus_16-13.1.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-13.1.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_16` | 13.1.0 | `el9.x86_64` | pigsty | 812.9 KiB | [citus_16-13.1.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-13.1.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_16` | 13.0.4 | `el9.aarch64` | pgdg | 744.1 KiB | [citus_16-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.0.4 | `el9.x86_64` | pgdg | 768.5 KiB | [citus_16-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 13.0.3 | `el9.x86_64` | pgdg | 769.2 KiB | [citus_16-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 13.0.3 | `el9.aarch64` | pgdg | 744.1 KiB | [citus_16-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.0.2 | `el9.x86_64` | pgdg | 768.9 KiB | [citus_16-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 13.0.2 | `el9.aarch64` | pgdg | 743.8 KiB | [citus_16-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.0.0 | `el9.aarch64` | pgdg | 741.7 KiB | [citus_16-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 13.0.0 | `el9.x86_64` | pgdg | 766.9 KiB | [citus_16-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.6 | `el9.aarch64` | pgdg | 742.2 KiB | [citus_16-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.6 | `el9.x86_64` | pgdg | 767.5 KiB | [citus_16-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.5 | `el9.x86_64` | pgdg | 766.6 KiB | [citus_16-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.5 | `el9.aarch64` | pgdg | 741.4 KiB | [citus_16-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.4 | `el9.x86_64` | pgdg | 766.4 KiB | [citus_16-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.4 | `el9.aarch64` | pgdg | 741.6 KiB | [citus_16-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.3 | `el9.x86_64` | pgdg | 766.2 KiB | [citus_16-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.3 | `el9.aarch64` | pgdg | 740.4 KiB | [citus_16-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.2 | `el9.aarch64` | pgdg | 739.1 KiB | [citus_16-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.2 | `el9.x86_64` | pgdg | 765.2 KiB | [citus_16-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.1 | `el9.x86_64` | pgdg | 765.2 KiB | [citus_16-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.1 | `el9.aarch64` | pgdg | 738.8 KiB | [citus_16-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | 12.1.0 | `el9.x86_64` | pgdg | 765.4 KiB | [citus_16-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_16` | 12.1.0 | `el9.aarch64` | pgdg | 738.9 KiB | [citus_16-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-citus` | 13.2.0 | `d12.x86_64` | pigsty | 2.6 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-citus` | 13.2.0 | `d12.aarch64` | pigsty | 2.5 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-citus` | 13.2.0 | `u22.x86_64` | pigsty | 3.3 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-16-citus` | 13.2.0 | `u22.aarch64` | pigsty | 3.2 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-16-citus` | 13.2.0 | `u24.aarch64` | pigsty | 2.7 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~noble_arm64.deb) |
| `postgresql-16-citus` | 13.2.0 | `u24.x86_64` | pigsty | 2.7 MiB | [postgresql-16-citus_13.2.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-9PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `citus_15` | 13.2.0 | `el8.aarch64` | pgdg | 815.5 KiB | [citus_15-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 13.2.0 | `el8.x86_64` | pgdg | 867.4 KiB | [citus_15-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.2.0 | `el8.aarch64` | pigsty | 957.9 KiB | [citus_15-13.2.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_15-13.2.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_15` | 13.2.0 | `el8.x86_64` | pigsty | 999.4 KiB | [citus_15-13.2.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_15-13.2.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_15` | 13.1.0 | `el8.aarch64` | pigsty | 933.9 KiB | [citus_15-13.1.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_15-13.1.0-9PIGSTY.el8.aarch64.rpm) |
| `citus_15` | 13.1.0 | `el8.x86_64` | pigsty | 973.6 KiB | [citus_15-13.1.0-9PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_15-13.1.0-9PIGSTY.el8.x86_64.rpm) |
| `citus_15` | 13.1.0 | `el8.x86_64` | pgdg | 844.2 KiB | [citus_15-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.1.0 | `el8.aarch64` | pgdg | 793.3 KiB | [citus_15-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 13.0.4 | `el8.x86_64` | pgdg | 824.2 KiB | [citus_15-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.0.4 | `el8.aarch64` | pgdg | 774.8 KiB | [citus_15-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 13.0.3 | `el8.x86_64` | pgdg | 824.0 KiB | [citus_15-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.0.3 | `el8.aarch64` | pgdg | 774.9 KiB | [citus_15-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 13.0.2 | `el8.aarch64` | pgdg | 774.7 KiB | [citus_15-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 13.0.2 | `el8.x86_64` | pgdg | 823.9 KiB | [citus_15-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.0.0 | `el8.x86_64` | pgdg | 821.8 KiB | [citus_15-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 13.0.0 | `el8.aarch64` | pgdg | 772.5 KiB | [citus_15-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.6 | `el8.aarch64` | pgdg | 771.3 KiB | [citus_15-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.6 | `el8.x86_64` | pgdg | 821.0 KiB | [citus_15-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.5 | `el8.x86_64` | pgdg | 819.8 KiB | [citus_15-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.5 | `el8.aarch64` | pgdg | 770.2 KiB | [citus_15-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.4 | `el8.aarch64` | pgdg | 769.9 KiB | [citus_15-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.4 | `el8.x86_64` | pgdg | 819.4 KiB | [citus_15-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.3 | `el8.x86_64` | pgdg | 819.2 KiB | [citus_15-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.3 | `el8.aarch64` | pgdg | 769.8 KiB | [citus_15-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.2 | `el8.aarch64` | pgdg | 768.9 KiB | [citus_15-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.2 | `el8.x86_64` | pgdg | 818.0 KiB | [citus_15-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.1 | `el8.aarch64` | pgdg | 768.9 KiB | [citus_15-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.1.1 | `el8.x86_64` | pgdg | 818.2 KiB | [citus_15-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.0 | `el8.x86_64` | pgdg | 818.0 KiB | [citus_15-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 12.1.0 | `el8.aarch64` | pgdg | 768.7 KiB | [citus_15-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.0.0 | `el8.aarch64` | pgdg | 770.5 KiB | [citus_15-12.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | 12.0.0 | `el8.x86_64` | pgdg | 820.5 KiB | [citus_15-12.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | 11.3.0 | `el8.x86_64` | pgdg | 801.6 KiB | [citus_15-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_15` | 11.3.0 | `el8.aarch64` | pgdg | 752.7 KiB | [citus_15-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_15` | 11.2.1 | `el8.aarch64` | pgdg | 732.0 KiB | [citus_15-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_15` | 11.2.1 | `el8.x86_64` | pgdg | 779.8 KiB | [citus_15-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_15` | 11.2.0 | `el8.x86_64` | pgdg | 778.2 KiB | [citus_15-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_15` | 11.2.0 | `el8.aarch64` | pgdg | 730.9 KiB | [citus_15-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_15` | 11.1.5 | `el8.x86_64` | pgdg | 766.8 KiB | [citus_15-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_15` | 11.1.5 | `el8.aarch64` | pgdg | 719.3 KiB | [citus_15-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_15` | 11.1.4 | `el8.aarch64` | pgdg | 719.0 KiB | [citus_15-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_15` | 11.1.4 | `el8.x86_64` | pgdg | 766.6 KiB | [citus_15-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_15` | 11.1.3 | `el8.aarch64` | pgdg | 718.7 KiB | [citus_15-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_15` | 11.1.3 | `el8.x86_64` | pgdg | 766.3 KiB | [citus_15-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_15` | 11.1.2 | `el8.x86_64` | pgdg | 765.8 KiB | [citus_15-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_15` | 13.2.0 | `el9.x86_64` | pgdg | 855.4 KiB | [citus_15-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 13.2.0 | `el9.x86_64` | pigsty | 893.4 KiB | [citus_15-13.2.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_15-13.2.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_15` | 13.2.0 | `el9.aarch64` | pgdg | 819.5 KiB | [citus_15-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.2.0 | `el9.aarch64` | pigsty | 856.2 KiB | [citus_15-13.2.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_15-13.2.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_15` | 13.1.0 | `el9.aarch64` | pgdg | 798.4 KiB | [citus_15-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.1.0 | `el9.x86_64` | pgdg | 831.8 KiB | [citus_15-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 13.1.0 | `el9.aarch64` | pigsty | 833.2 KiB | [citus_15-13.1.0-9PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_15-13.1.0-9PIGSTY.el9.aarch64.rpm) |
| `citus_15` | 13.1.0 | `el9.x86_64` | pigsty | 869.1 KiB | [citus_15-13.1.0-9PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_15-13.1.0-9PIGSTY.el9.x86_64.rpm) |
| `citus_15` | 13.0.4 | `el9.x86_64` | pgdg | 813.3 KiB | [citus_15-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 13.0.4 | `el9.aarch64` | pgdg | 780.6 KiB | [citus_15-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.0.3 | `el9.x86_64` | pgdg | 813.5 KiB | [citus_15-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 13.0.3 | `el9.aarch64` | pgdg | 780.8 KiB | [citus_15-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.0.2 | `el9.x86_64` | pgdg | 813.3 KiB | [citus_15-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 13.0.2 | `el9.aarch64` | pgdg | 780.7 KiB | [citus_15-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.0.0 | `el9.aarch64` | pgdg | 778.6 KiB | [citus_15-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 13.0.0 | `el9.x86_64` | pgdg | 811.4 KiB | [citus_15-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.6 | `el9.aarch64` | pgdg | 778.1 KiB | [citus_15-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.6 | `el9.x86_64` | pgdg | 809.2 KiB | [citus_15-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.5 | `el9.aarch64` | pgdg | 777.2 KiB | [citus_15-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.5 | `el9.x86_64` | pgdg | 809.5 KiB | [citus_15-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.4 | `el9.x86_64` | pgdg | 809.6 KiB | [citus_15-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.4 | `el9.aarch64` | pgdg | 777.4 KiB | [citus_15-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.3 | `el9.x86_64` | pgdg | 809.1 KiB | [citus_15-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.3 | `el9.aarch64` | pgdg | 774.8 KiB | [citus_15-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.2 | `el9.aarch64` | pgdg | 773.9 KiB | [citus_15-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.2 | `el9.x86_64` | pgdg | 805.9 KiB | [citus_15-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.1 | `el9.x86_64` | pgdg | 806.0 KiB | [citus_15-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.1.1 | `el9.aarch64` | pgdg | 773.8 KiB | [citus_15-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.0 | `el9.aarch64` | pgdg | 773.7 KiB | [citus_15-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 12.1.0 | `el9.x86_64` | pgdg | 807.2 KiB | [citus_15-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.0.0 | `el9.x86_64` | pgdg | 810.5 KiB | [citus_15-12.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | 12.0.0 | `el9.aarch64` | pgdg | 775.5 KiB | [citus_15-12.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | 11.3.0 | `el9.x86_64` | pgdg | 791.4 KiB | [citus_15-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_15` | 11.3.0 | `el9.aarch64` | pgdg | 759.0 KiB | [citus_15-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_15` | 11.2.1 | `el9.aarch64` | pgdg | 738.2 KiB | [citus_15-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_15` | 11.2.1 | `el9.x86_64` | pgdg | 769.9 KiB | [citus_15-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.2.0 | `el9.aarch64` | pgdg | 737.5 KiB | [citus_15-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_15` | 11.2.0 | `el9.x86_64` | pgdg | 769.1 KiB | [citus_15-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.1.5 | `el9.aarch64` | pgdg | 726.0 KiB | [citus_15-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_15` | 11.1.5 | `el9.x86_64` | pgdg | 755.7 KiB | [citus_15-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.1.4 | `el9.x86_64` | pgdg | 757.4 KiB | [citus_15-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.1.4 | `el9.aarch64` | pgdg | 724.9 KiB | [citus_15-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_15` | 11.1.3 | `el9.x86_64` | pgdg | 757.2 KiB | [citus_15-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.1.3 | `el9.aarch64` | pgdg | 724.7 KiB | [citus_15-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.3-1.rhel9.aarch64.rpm) |
| `citus_15` | 11.1.2 | `el9.x86_64` | pgdg | 756.6 KiB | [citus_15-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_15` | 11.1.2 | `el9.aarch64` | pgdg | 724.3 KiB | [citus_15-11.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.2-1.rhel9.aarch64.rpm) |
| `postgresql-15-citus` | 13.2.0 | `d12.aarch64` | pigsty | 2.6 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-citus` | 13.2.0 | `d12.x86_64` | pigsty | 2.6 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-citus` | 13.2.0 | `u22.x86_64` | pigsty | 3.3 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~jammy_amd64.deb) |
| `postgresql-15-citus` | 13.2.0 | `u22.aarch64` | pigsty | 3.3 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~jammy_arm64.deb) |
| `postgresql-15-citus` | 13.2.0 | `u24.aarch64` | pigsty | 2.7 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~noble_arm64.deb) |
| `postgresql-15-citus` | 13.2.0 | `u24.x86_64` | pigsty | 2.8 MiB | [postgresql-15-citus_13.2.0-9PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-9PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `citus_14` | 13.0.0 | `el8.x86_64` | pigsty | 807.0 KiB | [citus_14-13.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_14-13.0.0-1PIGSTY.el8.x86_64.rpm) |
| `citus_14` | 13.0.0 | `el8.aarch64` | pgdg | 765.7 KiB | [citus_14-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 13.0.0 | `el8.aarch64` | pigsty | 758.4 KiB | [citus_14-13.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_14-13.0.0-1PIGSTY.el8.aarch64.rpm) |
| `citus_14` | 13.0.0 | `el8.x86_64` | pgdg | 814.2 KiB | [citus_14-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.6 | `el8.x86_64` | pgdg | 813.6 KiB | [citus_14-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.6 | `el8.aarch64` | pgdg | 764.6 KiB | [citus_14-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.5 | `el8.aarch64` | pgdg | 763.5 KiB | [citus_14-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.5 | `el8.x86_64` | pgdg | 812.5 KiB | [citus_14-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.4 | `el8.x86_64` | pgdg | 812.5 KiB | [citus_14-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.4 | `el8.aarch64` | pgdg | 763.4 KiB | [citus_14-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.3 | `el8.aarch64` | pgdg | 763.4 KiB | [citus_14-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.3 | `el8.x86_64` | pgdg | 812.3 KiB | [citus_14-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.2 | `el8.x86_64` | pgdg | 811.1 KiB | [citus_14-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.2 | `el8.aarch64` | pgdg | 762.3 KiB | [citus_14-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.1 | `el8.x86_64` | pgdg | 811.0 KiB | [citus_14-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.1 | `el8.aarch64` | pgdg | 762.2 KiB | [citus_14-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.1.0 | `el8.x86_64` | pgdg | 810.5 KiB | [citus_14-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.1.0 | `el8.aarch64` | pgdg | 761.9 KiB | [citus_14-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 12.0.0 | `el8.x86_64` | pgdg | 813.4 KiB | [citus_14-12.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | 12.0.0 | `el8.aarch64` | pgdg | 763.9 KiB | [citus_14-12.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | 11.3.0 | `el8.aarch64` | pgdg | 748.9 KiB | [citus_14-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_14` | 11.3.0 | `el8.x86_64` | pgdg | 796.7 KiB | [citus_14-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_14` | 11.2.1 | `el8.x86_64` | pgdg | 776.9 KiB | [citus_14-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.2.1 | `el8.aarch64` | pgdg | 729.5 KiB | [citus_14-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_14` | 11.2.0 | `el8.x86_64` | pgdg | 776.0 KiB | [citus_14-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.2.0 | `el8.aarch64` | pgdg | 729.4 KiB | [citus_14-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_14` | 11.1.5 | `el8.aarch64` | pgdg | 718.0 KiB | [citus_14-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_14` | 11.1.5 | `el8.x86_64` | pgdg | 765.6 KiB | [citus_14-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.1.4 | `el8.aarch64` | pgdg | 717.7 KiB | [citus_14-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_14` | 11.1.4 | `el8.x86_64` | pgdg | 765.3 KiB | [citus_14-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.1.3 | `el8.aarch64` | pgdg | 717.4 KiB | [citus_14-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_14` | 11.1.3 | `el8.x86_64` | pgdg | 765.1 KiB | [citus_14-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.1.2 | `el8.x86_64` | pgdg | 764.4 KiB | [citus_14-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.1.1 | `el8.x86_64` | pgdg | 762.6 KiB | [citus_14-11.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.1-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.0.6 | `el8.x86_64` | pgdg | 701.4 KiB | [citus_14-11.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.6-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.0.5 | `el8.x86_64` | pgdg | 700.3 KiB | [citus_14-11.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.5-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.0.4 | `el8.x86_64` | pgdg | 699.6 KiB | [citus_14-11.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.4-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.0.3 | `el8.x86_64` | pgdg | 699.5 KiB | [citus_14-11.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.3-1.rhel8.x86_64.rpm) |
| `citus_14` | 11.0.2 | `el8.x86_64` | pgdg | 698.6 KiB | [citus_14-11.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.2-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.5 | `el8.x86_64` | pgdg | 618.5 KiB | [citus_14-10.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.5-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.4 | `el8.x86_64` | pgdg | 618.6 KiB | [citus_14-10.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.4-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.3 | `el8.x86_64` | pgdg | 618.5 KiB | [citus_14-10.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.3-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.2 | `el8.x86_64` | pgdg | 615.3 KiB | [citus_14-10.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.2-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.1 | `el8.x86_64` | pgdg | 614.8 KiB | [citus_14-10.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.1-1.rhel8.x86_64.rpm) |
| `citus_14` | 10.2.0 | `el8.x86_64` | pgdg | 614.0 KiB | [citus_14-10.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.0-1.rhel8.x86_64.rpm) |
| `citus_14` | 13.0.0 | `el9.aarch64` | pigsty | 769.9 KiB | [citus_14-13.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_14-13.0.0-1PIGSTY.el9.aarch64.rpm) |
| `citus_14` | 13.0.0 | `el9.aarch64` | pgdg | 771.1 KiB | [citus_14-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 13.0.0 | `el9.x86_64` | pigsty | 801.9 KiB | [citus_14-13.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_14-13.0.0-1PIGSTY.el9.x86_64.rpm) |
| `citus_14` | 13.0.0 | `el9.x86_64` | pgdg | 803.4 KiB | [citus_14-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.6 | `el9.aarch64` | pgdg | 770.5 KiB | [citus_14-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.6 | `el9.x86_64` | pgdg | 802.9 KiB | [citus_14-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.5 | `el9.aarch64` | pgdg | 770.0 KiB | [citus_14-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.5 | `el9.x86_64` | pgdg | 800.4 KiB | [citus_14-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.4 | `el9.x86_64` | pgdg | 800.6 KiB | [citus_14-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.4 | `el9.aarch64` | pgdg | 770.2 KiB | [citus_14-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.3 | `el9.aarch64` | pgdg | 768.2 KiB | [citus_14-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.3 | `el9.x86_64` | pgdg | 800.1 KiB | [citus_14-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.2 | `el9.aarch64` | pgdg | 766.8 KiB | [citus_14-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.2 | `el9.x86_64` | pgdg | 798.9 KiB | [citus_14-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.1 | `el9.aarch64` | pgdg | 766.5 KiB | [citus_14-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.1.1 | `el9.x86_64` | pgdg | 798.9 KiB | [citus_14-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.0 | `el9.x86_64` | pgdg | 798.9 KiB | [citus_14-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 12.1.0 | `el9.aarch64` | pgdg | 766.8 KiB | [citus_14-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.0.0 | `el9.aarch64` | pgdg | 769.0 KiB | [citus_14-12.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | 12.0.0 | `el9.x86_64` | pgdg | 802.0 KiB | [citus_14-12.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | 11.3.0 | `el9.x86_64` | pgdg | 787.6 KiB | [citus_14-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_14` | 11.3.0 | `el9.aarch64` | pgdg | 754.5 KiB | [citus_14-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_14` | 11.2.1 | `el9.x86_64` | pgdg | 767.2 KiB | [citus_14-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.2.1 | `el9.aarch64` | pgdg | 736.1 KiB | [citus_14-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_14` | 11.2.0 | `el9.aarch64` | pgdg | 735.2 KiB | [citus_14-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_14` | 11.2.0 | `el9.x86_64` | pgdg | 766.5 KiB | [citus_14-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.1.5 | `el9.aarch64` | pgdg | 724.2 KiB | [citus_14-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_14` | 11.1.5 | `el9.x86_64` | pgdg | 756.8 KiB | [citus_14-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.1.4 | `el9.x86_64` | pgdg | 755.9 KiB | [citus_14-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.1.4 | `el9.aarch64` | pgdg | 723.3 KiB | [citus_14-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_14` | 11.1.3 | `el9.aarch64` | pgdg | 723.1 KiB | [citus_14-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.3-1.rhel9.aarch64.rpm) |
| `citus_14` | 11.1.3 | `el9.x86_64` | pgdg | 755.8 KiB | [citus_14-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.1.2 | `el9.x86_64` | pgdg | 755.4 KiB | [citus_14-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.1.1 | `el9.x86_64` | pgdg | 754.1 KiB | [citus_14-11.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.1-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.0.6 | `el9.x86_64` | pgdg | 691.3 KiB | [citus_14-11.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.6-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.0.5 | `el9.x86_64` | pgdg | 690.5 KiB | [citus_14-11.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.5-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.0.4 | `el9.x86_64` | pgdg | 690.0 KiB | [citus_14-11.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.4-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.0.3 | `el9.x86_64` | pgdg | 689.8 KiB | [citus_14-11.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.3-1.rhel9.x86_64.rpm) |
| `citus_14` | 11.0.2 | `el9.x86_64` | pgdg | 689.0 KiB | [citus_14-11.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.2-1.rhel9.x86_64.rpm) |
| `citus_14` | 10.2.5 | `el9.x86_64` | pgdg | 612.1 KiB | [citus_14-10.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.5-1.rhel9.x86_64.rpm) |
| `citus_14` | 10.2.4 | `el9.x86_64` | pgdg | 613.7 KiB | [citus_14-10.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.4-1.rhel9.x86_64.rpm) |
| `citus_14` | 10.2.3 | `el9.x86_64` | pgdg | 613.7 KiB | [citus_14-10.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.3-1.rhel9.x86_64.rpm) |
| `postgresql-14-citus` | 13.0.0 | `d12.aarch64` | pigsty | 2.9 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-citus` | 13.0.0 | `d12.x86_64` | pigsty | 3.0 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-citus` | 13.0.0 | `u22.x86_64` | pigsty | 3.1 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-citus` | 13.0.0 | `u22.aarch64` | pigsty | 3.1 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-citus` | 13.0.0 | `u24.aarch64` | pigsty | 2.6 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-citus` | 13.0.0 | `u24.x86_64` | pigsty | 2.7 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `citus_13` | 9.5.4 | `el8.x86_64` | pgdg | 1.9 MiB | [citus_13-9.5.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.4-1.rhel8.x86_64.rpm) |
| `citus_13` | 9.5.2 | `el8.x86_64` | pgdg | 1.9 MiB | [citus_13-9.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 9.5.1 | `el8.x86_64` | pgdg | 1.9 MiB | [citus_13-9.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 9.5.0 | `el8.x86_64` | pgdg | 1.9 MiB | [citus_13-9.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.0-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.3.0 | `el8.aarch64` | pgdg | 747.9 KiB | [citus_13-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_13` | 11.3.0 | `el8.x86_64` | pgdg | 787.4 KiB | [citus_13-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_13` | 11.2.1 | `el8.aarch64` | pgdg | 728.8 KiB | [citus_13-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_13` | 11.2.1 | `el8.x86_64` | pgdg | 768.4 KiB | [citus_13-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.2.0 | `el8.x86_64` | pgdg | 767.1 KiB | [citus_13-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.2.0 | `el8.aarch64` | pgdg | 727.8 KiB | [citus_13-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_13` | 11.1.5 | `el8.aarch64` | pgdg | 716.8 KiB | [citus_13-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_13` | 11.1.5 | `el8.x86_64` | pgdg | 756.4 KiB | [citus_13-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.1.4 | `el8.x86_64` | pgdg | 756.2 KiB | [citus_13-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.1.4 | `el8.aarch64` | pgdg | 716.5 KiB | [citus_13-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_13` | 11.1.3 | `el8.aarch64` | pgdg | 716.2 KiB | [citus_13-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_13` | 11.1.3 | `el8.x86_64` | pgdg | 755.8 KiB | [citus_13-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.1.2 | `el8.x86_64` | pgdg | 755.4 KiB | [citus_13-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.1.1 | `el8.x86_64` | pgdg | 753.7 KiB | [citus_13-11.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.0.6 | `el8.x86_64` | pgdg | 692.9 KiB | [citus_13-11.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.6-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.0.5 | `el8.x86_64` | pgdg | 692.3 KiB | [citus_13-11.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.5-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.0.4 | `el8.x86_64` | pgdg | 691.4 KiB | [citus_13-11.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.4-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.0.3 | `el8.x86_64` | pgdg | 691.3 KiB | [citus_13-11.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.3-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.0.2 | `el8.x86_64` | pgdg | 690.5 KiB | [citus_13-11.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.5 | `el8.x86_64` | pgdg | 611.3 KiB | [citus_13-10.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.5-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.4 | `el8.x86_64` | pgdg | 611.3 KiB | [citus_13-10.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.4-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.3 | `el8.x86_64` | pgdg | 611.1 KiB | [citus_13-10.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.3-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.2 | `el8.x86_64` | pgdg | 608.3 KiB | [citus_13-10.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.1 | `el8.x86_64` | pgdg | 607.7 KiB | [citus_13-10.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.2.0 | `el8.x86_64` | pgdg | 607.2 KiB | [citus_13-10.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.0-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.1.2 | `el8.x86_64` | pgdg | 582.2 KiB | [citus_13-10.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.1.1 | `el8.x86_64` | pgdg | 581.0 KiB | [citus_13-10.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.1.0 | `el8.x86_64` | pgdg | 580.1 KiB | [citus_13-10.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.0-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.0.3 | `el8.x86_64` | pgdg | 556.1 KiB | [citus_13-10.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.3-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.0.2 | `el8.x86_64` | pgdg | 555.3 KiB | [citus_13-10.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.2-1.rhel8.x86_64.rpm) |
| `citus_13` | 10.0.1 | `el8.x86_64` | pgdg | 2.2 MiB | [citus_13-10.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.1-1.rhel8.x86_64.rpm) |
| `citus_13` | 11.3.0 | `el9.aarch64` | pgdg | 753.8 KiB | [citus_13-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_13` | 11.3.0 | `el9.x86_64` | pgdg | 788.2 KiB | [citus_13-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_13` | 11.2.1 | `el9.aarch64` | pgdg | 735.4 KiB | [citus_13-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_13` | 11.2.1 | `el9.x86_64` | pgdg | 768.0 KiB | [citus_13-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.2.0 | `el9.x86_64` | pgdg | 767.5 KiB | [citus_13-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.2.0 | `el9.aarch64` | pgdg | 734.0 KiB | [citus_13-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_13` | 11.1.5 | `el9.x86_64` | pgdg | 757.2 KiB | [citus_13-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.1.5 | `el9.aarch64` | pgdg | 723.5 KiB | [citus_13-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_13` | 11.1.4 | `el9.aarch64` | pgdg | 722.5 KiB | [citus_13-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_13` | 11.1.4 | `el9.x86_64` | pgdg | 754.6 KiB | [citus_13-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.1.3 | `el9.aarch64` | pgdg | 722.2 KiB | [citus_13-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.3-1.rhel9.aarch64.rpm) |
| `citus_13` | 11.1.3 | `el9.x86_64` | pgdg | 754.5 KiB | [citus_13-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.1.2 | `el9.x86_64` | pgdg | 753.8 KiB | [citus_13-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.1.1 | `el9.x86_64` | pgdg | 752.1 KiB | [citus_13-11.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.1-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.0.6 | `el9.x86_64` | pgdg | 691.1 KiB | [citus_13-11.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.6-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.0.5 | `el9.x86_64` | pgdg | 690.2 KiB | [citus_13-11.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.5-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.0.4 | `el9.x86_64` | pgdg | 690.0 KiB | [citus_13-11.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.4-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.0.3 | `el9.x86_64` | pgdg | 690.0 KiB | [citus_13-11.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.3-1.rhel9.x86_64.rpm) |
| `citus_13` | 11.0.2 | `el9.x86_64` | pgdg | 689.2 KiB | [citus_13-11.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.2-1.rhel9.x86_64.rpm) |
| `citus_13` | 10.2.5 | `el9.x86_64` | pgdg | 612.5 KiB | [citus_13-10.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.5-1.rhel9.x86_64.rpm) |
| `citus_13` | 10.2.4 | `el9.x86_64` | pgdg | 613.9 KiB | [citus_13-10.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.4-1.rhel9.x86_64.rpm) |
| `citus_13` | 10.2.3 | `el9.x86_64` | pgdg | 613.9 KiB | [citus_13-10.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.3-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="citus-13.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get citus_columnar; # get citus_columnar source code
pig build dep citus_columnar; # install build dependencies
pig build pkg citus_columnar; # build extension rpm or deb
pig build ext citus_columnar; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install citus_columnar; # install by extension name, for the current active PG version
pig ext install citus; # install via package alias, for the active PG version
pig ext install citus_columnar -v 17;   # install for PG 17
pig ext install citus_columnar -v 16;   # install for PG 16
pig ext install citus_columnar -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION citus_columnar CASCADE SCHEMA pg_catalog;
```

