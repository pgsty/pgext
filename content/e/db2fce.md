---
title: "db2fce"
linkTitle: "db2fce"
description: "DB2 compatibility functions, types, operators, and SYSIBM.SYSDUMMY1 for PostgreSQL."
weight: 9200
categories: ["SIM"]
width: full
---

[**db2fce**](https://github.com/credativ/db2fce) : DB2 compatibility functions, types, operators, and SYSIBM.SYSDUMMY1 for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9200** | {{< badge content="db2fce" link="https://github.com/credativ/db2fce" >}} | {{< ext "db2fce" >}} | `0.0.17` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `db2` `sysibm` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "orafce" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_job" >}} |

> [!Note] PGDG APT is complete for PG14-18; Pigsty RPM noarch spec fills the PGDG YUM gap for PG14-18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.0.17` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `db2fce` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.17` | {{< bg "18" "db2fce_18" "green" >}} {{< bg "17" "db2fce_17" "green" >}} {{< bg "16" "db2fce_16" "green" >}} {{< bg "15" "db2fce_15" "green" >}} {{< bg "14" "db2fce_14" "green" >}} | `db2fce_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.0.17` | {{< bg "18" "postgresql-18-db2fce" "green" >}} {{< bg "17" "postgresql-17-db2fce" "green" >}} {{< bg "16" "postgresql-16-db2fce" "green" >}} {{< bg "15" "postgresql-15-db2fce" "green" >}} {{< bg "14" "postgresql-14-db2fce" "green" >}} | `postgresql-$v-db2fce` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.17" "db2fce_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.0.17" "postgresql-18-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-17-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-16-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-15-db2fce : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.0.17" "postgresql-14-db2fce : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2fce_18` | `0.0.17` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.5 KiB | [db2fce_18-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db2fce_18-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_18` | `0.0.17` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.5 KiB | [db2fce_18-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db2fce_18-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_18` | `0.0.17` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.4 KiB | [db2fce_18-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db2fce_18-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_18` | `0.0.17` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [db2fce_18-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db2fce_18-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_18` | `0.0.17` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [db2fce_18-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db2fce_18-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `db2fce_18` | `0.0.17` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [db2fce_18-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db2fce_18-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-db2fce` | `0.0.17` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |
| `postgresql-18-db2fce` | `0.0.17` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.3 KiB | [postgresql-18-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-18-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2fce_17` | `0.0.17` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.5 KiB | [db2fce_17-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db2fce_17-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_17` | `0.0.17` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.5 KiB | [db2fce_17-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db2fce_17-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_17` | `0.0.17` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.4 KiB | [db2fce_17-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db2fce_17-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_17` | `0.0.17` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [db2fce_17-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db2fce_17-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_17` | `0.0.17` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [db2fce_17-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db2fce_17-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `db2fce_17` | `0.0.17` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [db2fce_17-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db2fce_17-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-db2fce` | `0.0.17` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |
| `postgresql-17-db2fce` | `0.0.17` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.3 KiB | [postgresql-17-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-17-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2fce_16` | `0.0.17` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.5 KiB | [db2fce_16-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db2fce_16-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_16` | `0.0.17` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.5 KiB | [db2fce_16-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db2fce_16-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_16` | `0.0.17` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.4 KiB | [db2fce_16-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db2fce_16-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_16` | `0.0.17` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [db2fce_16-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db2fce_16-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_16` | `0.0.17` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [db2fce_16-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db2fce_16-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `db2fce_16` | `0.0.17` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [db2fce_16-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db2fce_16-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-db2fce` | `0.0.17` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |
| `postgresql-16-db2fce` | `0.0.17` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.3 KiB | [postgresql-16-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-16-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2fce_15` | `0.0.17` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.5 KiB | [db2fce_15-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db2fce_15-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_15` | `0.0.17` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.5 KiB | [db2fce_15-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db2fce_15-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_15` | `0.0.17` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.4 KiB | [db2fce_15-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db2fce_15-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_15` | `0.0.17` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [db2fce_15-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db2fce_15-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_15` | `0.0.17` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [db2fce_15-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db2fce_15-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `db2fce_15` | `0.0.17` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [db2fce_15-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db2fce_15-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-db2fce` | `0.0.17` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |
| `postgresql-15-db2fce` | `0.0.17` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.3 KiB | [postgresql-15-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-15-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2fce_14` | `0.0.17` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.5 KiB | [db2fce_14-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db2fce_14-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_14` | `0.0.17` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.5 KiB | [db2fce_14-0.0.17-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db2fce_14-0.0.17-1PIGSTY.el8.noarch.rpm) |
| `db2fce_14` | `0.0.17` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.4 KiB | [db2fce_14-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db2fce_14-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_14` | `0.0.17` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.4 KiB | [db2fce_14-0.0.17-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db2fce_14-0.0.17-1PIGSTY.el9.noarch.rpm) |
| `db2fce_14` | `0.0.17` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [db2fce_14-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db2fce_14-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `db2fce_14` | `0.0.17` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [db2fce_14-0.0.17-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db2fce_14-0.0.17-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-db2fce` | `0.0.17` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg12+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.8 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.8 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg13+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg22.04+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg24.04+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |
| `postgresql-14-db2fce` | `0.0.17` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-14-db2fce_0.0.17-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/db2fce/postgresql-14-db2fce_0.0.17-1.pgdg26.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/db2fce" title="Repository" icon="github" subtitle="github.com/credativ/db2fce" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="db2fce-0.0.17.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg db2fce;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install db2fce;		# install via package name, for the active PG version

pig install db2fce -v 18;   # install for PG 18
pig install db2fce -v 17;   # install for PG 17
pig install db2fce -v 16;   # install for PG 16
pig install db2fce -v 15;   # install for PG 15
pig install db2fce -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION db2fce CASCADE; -- requires plpgsql
```




## Usage

Sources: [README](https://github.com/credativ/db2fce/blob/master/README.md), [SQL objects](https://github.com/credativ/db2fce/blob/master/db2fce.sql), [control file](https://github.com/credativ/db2fce/blob/master/db2fce.control)

`db2fce` provides a DB2 compatibility environment for PostgreSQL. It creates DB2-style functions, types, operators, and the `SYSIBM.SYSDUMMY1` compatibility view so SQL originally written for IBM Db2 can be adapted with fewer changes.

### Enable

```sql
CREATE EXTENSION db2fce;

SET search_path = db2, sysibm, public;
```

The extension creates most compatibility objects in the `db2` schema and creates `sysibm.sysdummy1` for DB2 queries that expect a dummy single-row table.

```sql
SELECT * FROM sysibm.sysdummy1;
```

### Compatibility Functions

The `db2` schema includes date/time helpers such as `microsecond`, `second`, `minute`, `hour`, `day`, `month`, `year`, `days`, `months_between`, `date`, `time`, and `timestamp_format`.

String and conversion helpers include `locate`, `translate`, `lcase`, `upper`, `lower`, `strip`, `char`, `integer`, `int`, `double`, `decimal`, `dec`, `hex`, `round`, `digits`, and `value`.

### Operators

The SQL layer also defines DB2-style operators such as `^=` for inequality and `!!` for concatenation across several data types.

```sql
SELECT db2.int('42');
SELECT db2.days(current_date);
SELECT 'db' !! '2';
```

### Notes

Adding `db2` to `search_path` lets many DB2 function calls work without schema qualification. Some names that conflict with PostgreSQL syntax or built-in behavior may still need explicit `db2.` qualification.
