---
title: "pg_bigm"
linkTitle: "pg_bigm"
description: "create 2-gram (bigram) index for faster full text search."
weight: 2120
categories: ["FTS"]
width: full
---

[**pg_bigm**](https://github.com/pgbigm/pg_bigm) : create 2-gram (bigram) index for faster full text search.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2120** | {{< badge content="pg_bigm" link="https://github.com/pgbigm/pg_bigm" >}} | {{< ext "pg_bigm" >}} | `1.2` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_tokenizer" >}} {{< ext "fuzzystrmatch" >}} {{< ext "rum" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_bigm` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "pg_bigm_18" "green" >}} {{< bg "17" "pg_bigm_17" "green" >}} {{< bg "16" "pg_bigm_16" "green" >}} {{< bg "15" "pg_bigm_15" "green" >}} {{< bg "14" "pg_bigm_14" "green" >}} | `pg_bigm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "postgresql-18-pg-bigm" "green" >}} {{< bg "17" "postgresql-17-pg-bigm" "green" >}} {{< bg "16" "postgresql-16-pg-bigm" "green" >}} {{< bg "15" "postgresql-15-pg-bigm" "green" >}} {{< bg "14" "postgresql-14-pg-bigm" "green" >}} | `postgresql-$v-pg-bigm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2" "pg_bigm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "pg_bigm_14 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-pg-bigm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-pg-bigm : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.9 KiB | [pg_bigm_18-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_18-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_bigm_18-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_18-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.7 KiB | [pg_bigm_18-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_18-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.7 KiB | [pg_bigm_18-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_18-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bigm_18-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.0 KiB | [pg_bigm_18-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_18-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.3 KiB | [pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bigm_18-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.0 KiB | [pg_bigm_18-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_18-1.2-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.4 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.1 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.6 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.3 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.8 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 29.8 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 29.6 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 29.6 KiB | [postgresql-18-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-18-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.5 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_bigm_17-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_17-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pg_bigm_17-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_17-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.9 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_bigm_17-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_17-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.6 KiB | [pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bigm_17-1.2_20240606-2PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [pg_bigm_17-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_17-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.4 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.2 KiB | [pg_bigm_17-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bigm_17-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_17-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_17-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pg_bigm_17-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_17-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.1 KiB | [pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.9 KiB | [pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bigm_17-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_17-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_17-1.2-2PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.7 KiB | [pg_bigm_17-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_17-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.0 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.8 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.1 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.8 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.2 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.2 KiB | [postgresql-17-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-17-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.4 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_bigm_16-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_16-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bigm_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.2 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pg_bigm_16-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_16-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.0 KiB | [pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bigm_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_bigm_16-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_16-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.4 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bigm_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.5 KiB | [pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [pg_bigm_16-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_16-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bigm_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.4 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.2 KiB | [pg_bigm_16-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bigm_16-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_16-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_16-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pg_bigm_16-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_16-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.2 KiB | [pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.0 KiB | [pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bigm_16-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_16-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_16-1.2-2PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.7 KiB | [pg_bigm_16-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_16-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.0 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.8 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.0 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.2 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.2 KiB | [postgresql-16-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-16-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_15-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.4 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bigm_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_15-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bigm_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.6 KiB | [pg_bigm_15-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_15-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.5 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bigm_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.5 KiB | [pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.6 KiB | [pg_bigm_15-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_15-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bigm_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.4 KiB | [pg_bigm_15-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bigm_15-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.0 KiB | [pg_bigm_15-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_15-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_15-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.4 KiB | [pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.1 KiB | [pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bigm_15-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.1 KiB | [pg_bigm_15-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_15-1.2-2PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_15-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_15-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.9 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.5 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.2 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.1 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.8 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.4 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.4 KiB | [postgresql-15-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-15-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bigm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.8 KiB | [pg_bigm_14-1.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bigm_14-1.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.4 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bigm_14-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_bigm_14-1.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bigm_14-1.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bigm_14-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_bigm_14-1.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bigm_14-1.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.5 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bigm_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.5 KiB | [pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2_20240606-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.7 KiB | [pg_bigm_14-1.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bigm_14-1.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bigm_14-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bigm_14-1.2_20250903-1PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.4 KiB | [pg_bigm_14-1.2_20240606-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bigm_14-1.2_20240606-3PGDG.rhel10.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.0 KiB | [pg_bigm_14-1.2-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_14-1.2-2PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.8 KiB | [pg_bigm_14-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bigm_14-1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.4 KiB | [pg_bigm_14-1.2_20250903-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bigm_14-1.2_20250903-1PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.1 KiB | [pg_bigm_14-1.2_20240606-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bigm_14-1.2_20240606-3PGDG.rhel10.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.1 KiB | [pg_bigm_14-1.2-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_14-1.2-2PIGSTY.el10.aarch64.rpm) |
| `pg_bigm_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.8 KiB | [pg_bigm_14-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bigm_14-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-bigm` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.8 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.4 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.0 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.4 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bigm` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.3 KiB | [postgresql-14-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bigm/postgresql-14-pg-bigm_1.2-20250903PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgbigm/pg_bigm" title="Repository" icon="github" subtitle="github.com/pgbigm/pg_bigm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bigm-1.2-20250903.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_bigm;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_bigm;		# install via package name, for the active PG version

pig install pg_bigm -v 18;   # install for PG 18
pig install pg_bigm -v 17;   # install for PG 17
pig install pg_bigm -v 16;   # install for PG 16
pig install pg_bigm -v 15;   # install for PG 15
pig install pg_bigm -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_bigm;
```



## Usage

> [pg_bigm Documentation](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html) | [GitHub: pgbigm/pg_bigm](https://github.com/pgbigm/pg_bigm)

The `pg_bigm` module provides full text search capability in PostgreSQL. This module allows a user to create **2-gram** (bigram) index for faster full text search.

pg_bigm is released under the PostgreSQL License, a liberal Open Source license, similar to the BSD or MIT licenses.

## Features

- **Bigram indexing**: Creates 2-gram (bigram) GIN indexes for text columns
- **Faster LIKE searches**: Accelerates `LIKE` queries including prefix, suffix, and substring searches
- **All language support**: Works with any language including CJK (Chinese, Japanese, Korean) without additional configuration
- **Simple API**: Provides similarity search functions and operators

## Functions and Operators

### Functions

| Function | Return Type | Description |
|----------|-------------|-------------|
| `likequery(text)` | `text` | Generates a search query for full text search from a keyword |
| `show_bigm(text)` | `text[]` | Shows all 2-grams in the given string |
| `pg_gin_pending_stats(regclass)` | `record` | Returns the number of pages and tuples in the pending list of a GIN index |

### Operators

| Operator | Description |
|----------|-------------|
| `text =% text` | Returns true if the similarity between the left and right operands is greater than or equal to `pg_bigm.similarity_limit` |

## GUC Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `pg_bigm.last_update` | `text` | - | Shows the last update date of the module (read-only) |
| `pg_bigm.enable_recheck` | `bool` | `on` | Controls whether recheck is performed during index scan |
| `pg_bigm.gin_key_limit` | `int` | `0` | Limits the maximum number of bigrams used for full text search. 0 means no limit |
| `pg_bigm.similarity_limit` | `real` | `0.3` | Sets the minimum similarity threshold for the `=%` operator |

## Examples

### Basic Full Text Search

```sql
-- Create extension
CREATE EXTENSION pg_bigm;

-- Create a table with text data
CREATE TABLE documents (id serial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
  ('PostgreSQL is a powerful database'),
  ('Full text search with bigram indexing'),
  ('Japanese text: 日本語テキスト検索');

-- Create a bigram index
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);

-- Search using LIKE
SELECT * FROM documents WHERE content LIKE '%search%';

-- Search using likequery function
SELECT * FROM documents WHERE content LIKE likequery('database');
```

### Similarity Search

```sql
-- Show bigrams for a string
SELECT show_bigm('PostgreSQL');

-- Similarity search
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

### Comparison with pg_trgm

pg_bigm has the following advantages over the built-in `pg_trgm`:

| Feature | pg_bigm | pg_trgm |
|---------|---------|---------|
| N-gram type | 2-gram (bigram) | 3-gram (trigram) |
| Minimum search string | 1 character | 3 characters |
| Non-alphabetic languages | Full support | Limited support |
| LIKE search types | Prefix, suffix, and substring | Prefix, suffix, and substring |

For detailed documentation including advanced usage and performance tuning, see the [official pg_bigm documentation](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html).
