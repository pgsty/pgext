---
title: "pg_bigm"
linkTitle: "pg_bigm"
description: "create 2-gram (bigram) index for faster full text search."
weight: 2120
categories: ["Fts"]
width: full
---

create 2-gram (bigram) index for faster full text search.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2120** | {{< badge content="pg_bigm" link="https://github.com/pgbigm/pg_bigm" >}} | {{< ext "pg_bigm" "pg_bigm" >}} | `1.2` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_tokenizer" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_bigm" >}} | `1.2` | {{< badge content="18" color="red" alt="pg_bigm_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_bigm_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_bigm" >}} | `1.2` | {{< badge content="18" color="red" alt="postgresql-18-pg-bigm" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-bigm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_bigm_18" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bigm_17" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bigm_16" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bigm_15" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bigm_14" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_bigm_18" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bigm_17" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bigm_16" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bigm_15" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bigm_14" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_bigm_18" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bigm_17" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bigm_16" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bigm_15" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bigm_14" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_bigm_18" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bigm_17" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bigm_16" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bigm_15" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bigm_14" "1.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-bigm" >}}     | {{< pkg "postgresql-17-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bigm" "1.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_18` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_18` | 1.2 | `el8.x86_64` | pgdg | 20.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_18` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_18` | 1.2 | `el9.x86_64` | pgdg | 21.0 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_17` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.x86_64` | pgdg | 18.5 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.aarch64` | pgdg | 20.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.aarch64` | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.x86_64` | pgdg | 18.9 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.x86_64` | pgdg | 20.8 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 30.6 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 30.1 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.8 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.4 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.1 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.3 KiB | [postgresql-17-pg-bigm_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 18.2 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 20.1 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 18.2 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 18.0 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 18.4 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 20.8 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 30.5 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 30.0 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.7 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.3 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.3 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.1 KiB | [postgresql-16-pg-bigm_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 30.6 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 30.2 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.4 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.7 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.2 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.4 KiB | [postgresql-15-pg-bigm_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 30.1 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 30.5 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.6 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.3 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.4 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.2 KiB | [postgresql-14-pg-bigm_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 18.2 KiB | [pg_bigm_13-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_13-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 20.2 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_13-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_13-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 29.8 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 30.2 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.0 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.2 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.1 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.0 KiB | [postgresql-13-pg-bigm_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgbigm/pg_bigm" title="Repository" icon="github" subtitle="github.com/pgbigm/pg_bigm" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_bigm-1.2-20240606.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_bigm; # get pg_bigm source code
pig build dep pg_bigm; # install build dependencies
pig build pkg pg_bigm; # build extension rpm or deb
pig build ext pg_bigm; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_bigm; # install by extension name, for the current active PG version
pig ext install pg_bigm; # install via package alias, for the active PG version
pig ext install pg_bigm -v 17;   # install for PG 17
pig ext install pg_bigm -v 16;   # install for PG 16
pig ext install pg_bigm -v 15;   # install for PG 15
pig ext install pg_bigm -v 14;   # install for PG 14
pig ext install pg_bigm -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_bigm;
```

