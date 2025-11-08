---
title: "pg_bigm"
linkTitle: "pg_bigm"
description: "create 2-gram (bigram) index for faster full text search."
weight: 2120
categories: ["FTS"]
width: full
---

create 2-gram (bigram) index for faster full text search.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2120** | {{< badge content="pg_bigm" link="https://github.com/pgbigm/pg_bigm" >}} | {{< ext "pg_bigm" >}} | `1.2` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_tokenizer" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_bigm" >}} | `1.2` | {{< bg "18" "pg_bigm_18*" "green" >}} {{< bg "17" "pg_bigm_17*" "green" >}} {{< bg "16" "pg_bigm_16*" "green" >}} {{< bg "15" "pg_bigm_15*" "green" >}} {{< bg "14" "pg_bigm_14*" "green" >}} {{< bg "13" "pg_bigm_13*" "green" >}} | `pg_bigm_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_bigm" >}} | `1.2` | {{< bg "18" "postgresql-18-pg-bigm" "green" >}} {{< bg "17" "postgresql-17-pg-bigm" "green" >}} {{< bg "16" "postgresql-16-pg-bigm" "green" >}} {{< bg "15" "postgresql-15-pg-bigm" "green" >}} {{< bg "14" "postgresql-14-pg-bigm" "green" >}} {{< bg "13" "postgresql-13-pg-bigm" "green" >}} | `postgresql-$v-pg-bigm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_13 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-13-pg-bigm : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_18` | 1.2 | `el8.x86_64` | pgdg | 20.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_18` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_18` | 1.2 | `el9.x86_64` | pgdg | 21.0 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_18` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_18` | 1.2 | `el10.x86_64` | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_18` | 1.2 | `el10.aarch64` | pgdg | 21.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 28.4 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 28.1 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 28.6 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 28.3 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 29.8 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 29.8 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 29.6 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 29.5 KiB | [postgresql-18-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_17` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.x86_64` | pgdg | 18.5 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.aarch64` | pgdg | 20.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.x86_64` | pgdg | 18.9 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.x86_64` | pgdg | 20.8 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.aarch64` | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.x86_64` | pigsty | 19.5 KiB | [pg_bigm_17-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_17-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.x86_64` | pgdg | 21.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.x86_64` | pgdg | 21.2 KiB | [pg_bigm_17-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.aarch64` | pigsty | 19.7 KiB | [pg_bigm_17-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_17-1.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.aarch64` | pgdg | 21.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_17` | 1.2 | `el10.aarch64` | pgdg | 20.9 KiB | [pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 27.0 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 26.8 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 27.1 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 27.0 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 32.1 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.7 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.2 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.2 KiB | [postgresql-17-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.x86_64` | pgdg | 18.2 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 18.2 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 18.0 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el8.aarch64` | pgdg | 20.1 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 20.8 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.x86_64` | pgdg | 18.4 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.x86_64` | pigsty | 19.5 KiB | [pg_bigm_16-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_16-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.x86_64` | pgdg | 21.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.x86_64` | pgdg | 21.2 KiB | [pg_bigm_16-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.aarch64` | pigsty | 19.7 KiB | [pg_bigm_16-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_16-1.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.aarch64` | pgdg | 21.0 KiB | [pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_16` | 1.2 | `el10.aarch64` | pgdg | 21.2 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 26.9 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 26.7 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 27.1 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 27.0 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 32.0 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.7 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.2 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.2 KiB | [postgresql-16-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.x86_64` | pigsty | 19.8 KiB | [pg_bigm_15-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_15-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.x86_64` | pgdg | 21.7 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.x86_64` | pgdg | 21.4 KiB | [pg_bigm_15-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.aarch64` | pigsty | 19.8 KiB | [pg_bigm_15-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_15-1.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.aarch64` | pgdg | 21.1 KiB | [pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_15` | 1.2 | `el10.aarch64` | pgdg | 21.4 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 27.2 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 26.9 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 27.4 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 27.2 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 32.1 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.8 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.4 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.3 KiB | [postgresql-15-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el8.aarch64` | pgdg | 20.3 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 18.5 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.x86_64` | pigsty | 19.8 KiB | [pg_bigm_14-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_14-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.x86_64` | pgdg | 21.7 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.x86_64` | pgdg | 21.4 KiB | [pg_bigm_14-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bigm_14-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.aarch64` | pigsty | 19.8 KiB | [pg_bigm_14-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_14-1.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.aarch64` | pgdg | 21.1 KiB | [pg_bigm_14-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bigm_14-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_14` | 1.2 | `el10.aarch64` | pgdg | 21.4 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 27.2 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 26.8 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 27.4 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 27.1 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.9 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.7 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.3 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.3 KiB | [postgresql-14-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 18.4 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 18.2 KiB | [pg_bigm_13-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.x86_64` | pgdg | 20.4 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bigm_13-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 20.2 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el8.aarch64` | pgdg | 18.1 KiB | [pg_bigm_13-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bigm_13-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 18.5 KiB | [pg_bigm_13-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 18.8 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bigm_13-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_bigm_13-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 20.5 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el9.aarch64` | pgdg | 18.1 KiB | [pg_bigm_13-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bigm_13-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.x86_64` | pigsty | 19.8 KiB | [pg_bigm_13-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_13-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.x86_64` | pgdg | 21.4 KiB | [pg_bigm_13-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_bigm_13-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.x86_64` | pgdg | 21.7 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_bigm_13-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.aarch64` | pigsty | 19.8 KiB | [pg_bigm_13-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_13-1.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.aarch64` | pgdg | 21.1 KiB | [pg_bigm_13-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_bigm_13-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_13` | 1.2 | `el10.aarch64` | pgdg | 21.4 KiB | [pg_bigm_13-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_bigm_13-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-bigm` | 1.2 | `d12.x86_64` | pigsty | 26.8 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `d12.aarch64` | pigsty | 26.6 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `d13.x86_64` | pigsty | 27.1 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `d13.aarch64` | pigsty | 26.9 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u22.x86_64` | pigsty | 31.6 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u22.aarch64` | pigsty | 31.4 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u24.x86_64` | pigsty | 28.0 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-bigm` | 1.2 | `u24.aarch64` | pigsty | 28.0 KiB | [postgresql-13-pg-bigm_1.2-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-13-pg-bigm_1.2-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgbigm/pg_bigm" title="Repository" icon="github" subtitle="github.com/pgbigm/pg_bigm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bigm-1.2-20250903.tar.gz" >}}
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
pig ext install pg_bigm -v 18;   # install for PG 18
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

