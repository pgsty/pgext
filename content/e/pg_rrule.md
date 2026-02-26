---
title: "pg_rrule"
linkTitle: "pg_rrule"
description: "RRULE field type for PostgreSQL"
weight: 3880
categories: ["TYPE"]
width: full
---

[**pg_rrule**](https://github.com/petropavel13/pg_rrule) : RRULE field type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3880** | {{< badge content="pg_rrule" link="https://github.com/petropavel13/pg_rrule" >}} | {{< ext "pg_rrule" >}} | `0.3.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] require libical


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_rrule` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_rrule_18" "green" >}} {{< bg "17" "pg_rrule_17" "green" >}} {{< bg "16" "pg_rrule_16" "green" >}} {{< bg "15" "pg_rrule_15" "green" >}} {{< bg "14" "pg_rrule_14" "green" >}} {{< bg "13" "pg_rrule_13" "green" >}} | `pg_rrule_$v` | `libical` |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-rrule" "green" >}} {{< bg "17" "postgresql-17-pg-rrule" "green" >}} {{< bg "16" "postgresql-16-pg-rrule" "green" >}} {{< bg "15" "postgresql-15-pg-rrule" "green" >}} {{< bg "14" "postgresql-14-pg-rrule" "green" >}} {{< bg "13" "postgresql-13-pg-rrule" "green" >}} | `postgresql-$v-pg-rrule` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_rrule_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.3.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.6 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [pg_rrule_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.4 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.2 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.3 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.3 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 24.2 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 23.5 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.4 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.1 KiB | [postgresql-18-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.6 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [pg_rrule_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.5 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.2 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.4 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.3 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.0 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.1 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.4 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.1 KiB | [postgresql-17-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.6 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [pg_rrule_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.4 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.2 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.4 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.3 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.0 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.1 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.4 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.1 KiB | [postgresql-16-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.1 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.9 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.7 KiB | [pg_rrule_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.4 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.2 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.3 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.3 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.1 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 24.6 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.5 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.3 KiB | [postgresql-15-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.1 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.9 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.7 KiB | [pg_rrule_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.4 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.1 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.4 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.2 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.1 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 24.6 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.5 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.2 KiB | [postgresql-14-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrule_13` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrule_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrule_13` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrule_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrule_13` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrule_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrule_13` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.3 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrule_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrule_13` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.8 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrule_13-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrule_13` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.6 KiB | [pg_rrule_13-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrule_13-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-rrule` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 24.1 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 24.3 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 24.0 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 24.3 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.0 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 24.3 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 24.2 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-rrule` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 24.3 KiB | [postgresql-13-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petropavel13/pg_rrule" title="Repository" icon="github" subtitle="github.com/petropavel13/pg_rrule" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_rrule-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_rrule;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_rrule;		# install via package name, for the active PG version

pig install pg_rrule -v 18;   # install for PG 18
pig install pg_rrule -v 17;   # install for PG 17
pig install pg_rrule -v 16;   # install for PG 16
pig install pg_rrule -v 15;   # install for PG 15
pig install pg_rrule -v 14;   # install for PG 14
pig install pg_rrule -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_rrule;
```
