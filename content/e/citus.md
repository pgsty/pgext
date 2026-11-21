---
title: "citus"
linkTitle: "citus"
description: "Distributed PostgreSQL as an extension"
weight: 2400
categories: ["OLAP"]
languages: ["C"]
licenses: ["AGPL-3.0-only"]
repos: ["PIGSTY"]
page_width: full
---

[**citus**](https://github.com/citusdata/citus) : Distributed PostgreSQL as an extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2400** | {{< badge content="citus" link="https://github.com/citusdata/citus" >}} | {{< ext "citus" >}} | `14.3.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0-only" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "pgspider_ext" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_partman" >}} {{< ext "timescaledb" >}} {{< ext "pg_ttl_index" >}} {{< ext "pgactive" >}} {{< ext "spock" >}} {{< ext "pg_fkpart" >}} {{< ext "db_migrator" >}} |
|    **Siblings**   | {{< ext "citus_columnar" >}} |

> [!Note] SQL 14.3-1; preload; conflicts with Hydra.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.3.0` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.3.0` | {{< bg "19" "citus_19" "red" >}} {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.3.0` | {{< bg "19" "postgresql-19-citus" "red" >}} {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} | `postgresql-$v-citus` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 19" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 22" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 19" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 21" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "citus_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "citus_18 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 14.3.0" "citus_16 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 13.5.0" "citus_15 : AVAIL 5" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-citus : N/A 0" "gray" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.3.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.5.0" "postgresql-15-citus : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_18` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [citus_18-14.3.0-9PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_18-14.3.0-9PGSTY.el8.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 892.3 KiB | [citus_18-14.3.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.3.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 891.3 KiB | [citus_18-14.3.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.3.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 878.8 KiB | [citus_18-14.2.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.2.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 871.9 KiB | [citus_18-14.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 859.5 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [citus_18-14.3.0-9PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_18-14.3.0-9PGSTY.el8.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 841.3 KiB | [citus_18-14.3.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.3.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.4 KiB | [citus_18-14.3.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.3.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 829.0 KiB | [citus_18-14.2.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.2.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 822.0 KiB | [citus_18-14.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.2 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [citus_18-14.3.0-9PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_18-14.3.0-9PGSTY.el9.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 851.0 KiB | [citus_18-14.3.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.3.0-2PGDG.rhel9.9.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [citus_18-14.3.0-9PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_18-14.3.0-9PGSTY.el9.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 824.2 KiB | [citus_18-14.3.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.3.0-2PGDG.rhel9.9.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [citus_18-14.3.0-9PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_18-14.3.0-9PGSTY.el10.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 866.8 KiB | [citus_18-14.3.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.3.0-2PGDG.rhel10.3.x86_64.rpm) |
| `citus_18` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [citus_18-14.3.0-9PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_18-14.3.0-9PGSTY.el10.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 836.3 KiB | [citus_18-14.3.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.3.0-2PGDG.rhel10.3.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 836.4 KiB | [citus_18-14.3.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.3.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 835.3 KiB | [citus_18-14.3.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.3.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 824.8 KiB | [citus_18-14.2.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.2.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 818.3 KiB | [citus_18-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 805.9 KiB | [citus_18-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 806.0 KiB | [citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 806.1 KiB | [citus_18-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-citus` | `14.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.3.0-9PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~bookworm_amd64.deb) |
| `postgresql-18-citus` | `14.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.3.0-9PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~bookworm_arm64.deb) |
| `postgresql-18-citus` | `14.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.3.0-9PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~trixie_amd64.deb) |
| `postgresql-18-citus` | `14.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.3.0-9PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~trixie_arm64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-18-citus_14.3.0-9PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~jammy_amd64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-18-citus_14.3.0-9PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~jammy_arm64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.3.0-9PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~noble_amd64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.3.0-9PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~noble_arm64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.3.0-9PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~resolute_amd64.deb) |
| `postgresql-18-citus` | `14.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.3.0-9PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-18-citus_14.3.0-9PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_17` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [citus_17-14.3.0-9PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-14.3.0-9PGSTY.el8.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 889.8 KiB | [citus_17-14.3.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.3.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 889.1 KiB | [citus_17-14.3.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.3.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `14.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 876.4 KiB | [citus_17-14.2.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.2.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 869.7 KiB | [citus_17-14.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 857.0 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 850.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 827.1 KiB | [citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.7 KiB | [citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 803.6 KiB | [citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [citus_17-14.3.0-9PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-14.3.0-9PGSTY.el8.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 839.4 KiB | [citus_17-14.3.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.3.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 838.3 KiB | [citus_17-14.3.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.3.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `14.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 826.7 KiB | [citus_17-14.2.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.2.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 819.9 KiB | [citus_17-14.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 808.6 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 802.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 781.5 KiB | [citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.4 KiB | [citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.5 KiB | [citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.2 KiB | [citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 758.9 KiB | [citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [citus_17-14.3.0-9PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-14.3.0-9PGSTY.el9.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 848.5 KiB | [citus_17-14.3.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.3.0-2PGDG.rhel9.9.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [citus_17-14.3.0-9PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-14.3.0-9PGSTY.el9.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 822.0 KiB | [citus_17-14.3.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.3.0-2PGDG.rhel9.9.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [citus_17-14.3.0-9PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_17-14.3.0-9PGSTY.el10.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 864.6 KiB | [citus_17-14.3.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.3.0-2PGDG.rhel10.3.x86_64.rpm) |
| `citus_17` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [citus_17-14.3.0-9PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_17-14.3.0-9PGSTY.el10.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 834.5 KiB | [citus_17-14.3.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.3.0-2PGDG.rhel10.3.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 834.4 KiB | [citus_17-14.3.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.3.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 833.4 KiB | [citus_17-14.3.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.3.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 822.9 KiB | [citus_17-14.2.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.2.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 816.4 KiB | [citus_17-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 804.0 KiB | [citus_17-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 803.8 KiB | [citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 803.7 KiB | [citus_17-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 798.8 KiB | [citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 776.8 KiB | [citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 758.8 KiB | [citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 759.3 KiB | [citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-citus` | `14.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.3.0-9PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~bookworm_amd64.deb) |
| `postgresql-17-citus` | `14.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.3.0-9PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~bookworm_arm64.deb) |
| `postgresql-17-citus` | `14.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.3.0-9PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~trixie_amd64.deb) |
| `postgresql-17-citus` | `14.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.3.0-9PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~trixie_arm64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-17-citus_14.3.0-9PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~jammy_amd64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.4 MiB | [postgresql-17-citus_14.3.0-9PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~jammy_arm64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.3.0-9PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~noble_amd64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.3.0-9PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~noble_arm64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.3.0-9PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~resolute_amd64.deb) |
| `postgresql-17-citus` | `14.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.3.0-9PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-17-citus_14.3.0-9PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_16` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [citus_16-14.3.0-9PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-14.3.0-9PGSTY.el8.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 883.2 KiB | [citus_16-14.3.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-14.3.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 882.3 KiB | [citus_16-14.3.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-14.3.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_16` | `14.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 869.2 KiB | [citus_16-14.2.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-14.2.0-2PGDG.rhel8.10.x86_64.rpm) |
| `citus_16` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 862.3 KiB | [citus_16-14.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-14.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 849.8 KiB | [citus_16-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 843.6 KiB | [citus_16-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 819.5 KiB | [citus_16-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 800.7 KiB | [citus_16-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 800.6 KiB | [citus_16-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 800.4 KiB | [citus_16-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 798.7 KiB | [citus_16-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 797.3 KiB | [citus_16-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 796.2 KiB | [citus_16-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 796.4 KiB | [citus_16-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 796.1 KiB | [citus_16-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 795.2 KiB | [citus_16-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 795.0 KiB | [citus_16-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `12.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 794.9 KiB | [citus_16-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/citus_16-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [citus_16-14.3.0-9PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-14.3.0-9PGSTY.el8.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 833.0 KiB | [citus_16-14.3.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-14.3.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 832.0 KiB | [citus_16-14.3.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-14.3.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_16` | `14.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 820.4 KiB | [citus_16-14.2.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-14.2.0-2PGDG.rhel8.10.aarch64.rpm) |
| `citus_16` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 813.9 KiB | [citus_16-14.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-14.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 801.9 KiB | [citus_16-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 796.0 KiB | [citus_16-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 774.6 KiB | [citus_16-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 756.4 KiB | [citus_16-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 756.4 KiB | [citus_16-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 756.2 KiB | [citus_16-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 754.1 KiB | [citus_16-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 753.1 KiB | [citus_16-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 751.8 KiB | [citus_16-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 751.8 KiB | [citus_16-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 751.6 KiB | [citus_16-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 750.4 KiB | [citus_16-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 750.4 KiB | [citus_16-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `12.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 750.0 KiB | [citus_16-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/citus_16-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.8 MiB | [citus_16-14.3.0-9PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-14.3.0-9PGSTY.el9.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 842.7 KiB | [citus_16-14.3.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.3.0-2PGDG.rhel9.9.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [citus_16-14.3.0-9PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-14.3.0-9PGSTY.el9.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 816.3 KiB | [citus_16-14.3.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.3.0-2PGDG.rhel9.9.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [citus_16-14.3.0-9PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_16-14.3.0-9PGSTY.el10.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 857.5 KiB | [citus_16-14.3.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.3.0-2PGDG.rhel10.3.x86_64.rpm) |
| `citus_16` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [citus_16-14.3.0-9PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_16-14.3.0-9PGSTY.el10.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 828.4 KiB | [citus_16-14.3.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.3.0-2PGDG.rhel10.3.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 828.4 KiB | [citus_16-14.3.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.3.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 827.0 KiB | [citus_16-14.3.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.3.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 816.4 KiB | [citus_16-14.2.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.2.0-2PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 809.2 KiB | [citus_16-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.5 KiB | [citus_16-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.4 KiB | [citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.3 KiB | [citus_16-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 792.1 KiB | [citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 770.0 KiB | [citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.3 KiB | [citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.7 KiB | [citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-citus` | `14.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.3.0-9PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~bookworm_amd64.deb) |
| `postgresql-16-citus` | `14.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.3.0-9PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~bookworm_arm64.deb) |
| `postgresql-16-citus` | `14.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.3.0-9PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~trixie_amd64.deb) |
| `postgresql-16-citus` | `14.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.3.0-9PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~trixie_arm64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-16-citus_14.3.0-9PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~jammy_amd64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-16-citus_14.3.0-9PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~jammy_arm64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-16-citus_14.3.0-9PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~noble_amd64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-16-citus_14.3.0-9PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~noble_arm64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-16-citus_14.3.0-9PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~resolute_amd64.deb) |
| `postgresql-16-citus` | `14.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.3.0-9PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-16-citus_14.3.0-9PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_15` | `13.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [citus_15-13.5.0-9PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_15-13.5.0-9PGSTY.el8.x86_64.rpm) |
| `citus_15` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 867.4 KiB | [citus_15-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 844.2 KiB | [citus_15-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 824.2 KiB | [citus_15-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 824.0 KiB | [citus_15-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 823.9 KiB | [citus_15-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 821.8 KiB | [citus_15-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 821.0 KiB | [citus_15-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 819.8 KiB | [citus_15-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 819.4 KiB | [citus_15-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 819.2 KiB | [citus_15-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 818.0 KiB | [citus_15-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 818.2 KiB | [citus_15-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 818.0 KiB | [citus_15-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `12.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 820.5 KiB | [citus_15-12.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-12.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_15` | `11.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 801.6 KiB | [citus_15-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_15` | `11.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 779.8 KiB | [citus_15-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_15` | `11.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 778.2 KiB | [citus_15-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_15` | `11.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 766.8 KiB | [citus_15-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_15` | `11.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 766.6 KiB | [citus_15-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_15` | `11.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 766.3 KiB | [citus_15-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_15` | `11.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 765.8 KiB | [citus_15-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/citus_15-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_15` | `13.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [citus_15-13.5.0-9PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_15-13.5.0-9PGSTY.el8.aarch64.rpm) |
| `citus_15` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 815.5 KiB | [citus_15-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 793.3 KiB | [citus_15-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 774.8 KiB | [citus_15-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 774.9 KiB | [citus_15-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 774.7 KiB | [citus_15-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 772.5 KiB | [citus_15-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 771.3 KiB | [citus_15-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 770.2 KiB | [citus_15-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 769.9 KiB | [citus_15-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 769.8 KiB | [citus_15-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 768.9 KiB | [citus_15-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 768.9 KiB | [citus_15-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 768.7 KiB | [citus_15-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `12.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 770.5 KiB | [citus_15-12.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-12.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_15` | `11.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 752.7 KiB | [citus_15-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_15` | `11.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 732.0 KiB | [citus_15-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_15` | `11.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 730.9 KiB | [citus_15-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_15` | `11.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 719.3 KiB | [citus_15-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_15` | `11.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 719.0 KiB | [citus_15-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_15` | `11.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 718.7 KiB | [citus_15-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/citus_15-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_15` | `13.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.8 MiB | [citus_15-13.5.0-9PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_15-13.5.0-9PGSTY.el9.x86_64.rpm) |
| `citus_15` | `13.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [citus_15-13.5.0-9PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_15-13.5.0-9PGSTY.el9.aarch64.rpm) |
| `citus_15` | `13.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [citus_15-13.5.0-9PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_15-13.5.0-9PGSTY.el10.x86_64.rpm) |
| `citus_15` | `13.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [citus_15-13.5.0-9PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_15-13.5.0-9PGSTY.el10.aarch64.rpm) |
| `citus_15` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 826.0 KiB | [citus_15-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 805.6 KiB | [citus_15-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 787.5 KiB | [citus_15-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 787.6 KiB | [citus_15-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-citus` | `13.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.5.0-9PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~bookworm_amd64.deb) |
| `postgresql-15-citus` | `13.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-15-citus_13.5.0-9PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~bookworm_arm64.deb) |
| `postgresql-15-citus` | `13.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.5.0-9PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~trixie_amd64.deb) |
| `postgresql-15-citus` | `13.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-15-citus_13.5.0-9PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~trixie_arm64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-15-citus_13.5.0-9PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~jammy_amd64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-15-citus_13.5.0-9PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~jammy_arm64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-15-citus_13.5.0-9PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~noble_amd64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.5.0-9PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~noble_arm64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-15-citus_13.5.0-9PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~resolute_amd64.deb) |
| `postgresql-15-citus` | `13.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.5.0-9PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-15-citus_13.5.0-9PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install citus;		# install via package name, for the active PG version

pig install citus -v 18;   # install for PG 18
pig install citus -v 17;   # install for PG 17
pig install citus -v 16;   # install for PG 16

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'citus';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION citus;
```

## Usage

Sources:

- [Citus v14.3.0 release](https://github.com/citusdata/citus/releases/tag/v14.3.0)
- [Citus v14.3.0 CHANGELOG](https://github.com/citusdata/citus/blob/v14.3.0/CHANGELOG.md)
- [Citus v14.3.0 control file](https://github.com/citusdata/citus/blob/v14.3.0/src/backend/distributed/citus.control)
- [Citus v14.3.0 README](https://github.com/citusdata/citus/blob/v14.3.0/README.md)
- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Citus Utility Functions](https://docs.citusdata.com/en/stable/develop/api_udf.html)
- [Citus upgrade procedure](https://docs.citusdata.com/en/stable/admin_guide/upgrading_citus.html)

Citus turns PostgreSQL into a distributed database by sharding tables across worker nodes while keeping PostgreSQL SQL, indexes, extensions, transactions, and operational tooling as the user-facing surface. It is commonly used for multi-tenant SaaS databases, real-time analytics, time-series/event workloads, and distributed microservice schemas.

The extension name is `citus`. Release 14.3.0 uses SQL/control version `14.3-1` and supports PostgreSQL 16-18. Every node must preload the library before a superuser runs `CREATE EXTENSION`; the extension is not relocatable.

### Enable Citus

```ini
shared_preload_libraries = 'citus'
```

Preserve any other required preload libraries. Restart PostgreSQL on the coordinator and workers, then create the extension in each participating database:

```sql
CREATE EXTENSION IF NOT EXISTS citus;
SELECT citus_version();
```

On a multi-node cluster, register the coordinator and workers from the coordinator:

```sql
SELECT citus_set_coordinator_host('coord-1', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

SELECT * FROM citus_get_active_worker_nodes();
```

### Distributed Tables

Distribute a table by a shard key. Rows with the same shard-key value are colocated on the same shard, so tenant-scoped joins and point lookups stay local.

```sql
CREATE TABLE events (
  tenant_id  bigint,
  event_id   bigserial,
  event_at   timestamptz DEFAULT now(),
  kind       text,
  payload    jsonb,
  PRIMARY KEY (tenant_id, event_id)
);
```

Distribute the table and tune the shard count and colocation explicitly:

```sql
SELECT create_distributed_table(
  'events',
  'tenant_id',
  shard_count  := 64,
  colocate_with := 'none'
);
```

When choosing an explicit shard count, start a new colocation group with `colocate_with := 'none'`. To colocate with an existing distributed table, name that table and let its shard count determine the layout.

Queries that filter on the distribution column can route to a single shard:

```sql
SELECT *
FROM events
WHERE tenant_id = 42
ORDER BY event_at DESC
LIMIT 50;
```

Cross-shard queries are planned as distributed tasks and run in parallel on the workers:

```sql
SELECT kind, count(*)
FROM events
WHERE event_at >= now() - interval '1 hour'
GROUP BY kind
ORDER BY count DESC;
```

### Reference Tables

Reference tables are fully replicated to all workers. They are useful for small lookup tables that must join with many distributed tables.

```sql
CREATE TABLE countries (
  code text PRIMARY KEY,
  name text NOT NULL
);

SELECT create_reference_table('countries');
```

### Schema-Based Sharding

Schema-based sharding is useful when each tenant or service owns its own schema. Citus supports running schema-sharding DDLs from any node, including `CREATE SCHEMA`, `DROP SCHEMA`, `ALTER SCHEMA RENAME`, `ALTER SCHEMA OWNER`, and table-level DDL on distributed schemas.

```sql
CREATE SCHEMA tenant_42;
SELECT citus_schema_distribute('tenant_42');

CREATE TABLE tenant_42.orders (
  id bigserial PRIMARY KEY,
  amount numeric,
  created_at timestamptz DEFAULT now()
);
```

Use row-based distribution for shared tables and schema-based sharding for per-tenant schema layouts; do not mix the two models casually without checking colocation and SQL-support implications.

### Node and Shard Operations

```sql
-- Add or disable nodes.
SELECT * FROM citus_add_node('worker-3', 5432);
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- Drain and remove a node.
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- Rebalance shards.
SELECT citus_rebalance_start();
SELECT * FROM citus_rebalance_status();
SELECT rebalance_table_shards('events');

-- Inspect tables and shards.
SELECT * FROM citus_tables;
SELECT * FROM citus_shards;
```

### Backup Coordination

Citus v14.1.0 adds UDFs for blocking distributed 2PC commit decisions and schema/topology changes while taking coordinated disk snapshots. Use them only inside a controlled backup workflow, and always unblock the cluster after the snapshot step.

```sql
SELECT citus_cluster_changes_block();
SELECT * FROM citus_cluster_changes_block_status();

-- Take coordinated filesystem or volume snapshots here.

SELECT citus_cluster_changes_unblock();
```

Pair these functions with regular PostgreSQL backup discipline: consistent checkpoints, WAL archiving, snapshot ordering across nodes, and a tested restore procedure.

### Administrative Controls

Citus 14.2 adds the superuser-only `citus_internal.distribute_object()` repair helper for supported database objects whose metadata was not propagated correctly. Treat it as a targeted recovery operation, not as the normal distribution API.

The release also adds `citus.allow_unsafe_insert_select_pushdown` for explicitly opting into batched `INSERT ... SELECT` pushdown on colocated tables, and improves single-shard stored-procedure execution. Keep the unsafe pushdown setting off unless the workload has been checked against the release-note restrictions.

### Caveats

- Citus 14 supports PostgreSQL 16-18 and dropped PostgreSQL 15 support.
- `shared_preload_libraries = 'citus'` must be set before extension creation. A plain `CREATE EXTENSION citus` is not enough on a fresh server.
- Choose the distribution column carefully. Primary keys and unique constraints on distributed tables generally need to include the distribution column.
- Cross-shard joins, repartition joins, distributed DDL, and multi-shard writes are powerful but have different planning and locking behavior from single-node PostgreSQL.
- The distribution also provides `citus_columnar`; its storage and maintenance requirements are separate from distributing row tables.
- The cluster-change blocking functions are operational tools for backups. Do not leave a cluster blocked after a failed backup script.

### Release 14.3 and Upgrades

Release 14.3.0 merges sorted worker results for multi-shard ORDER BY queries and fetches rows in smaller batches where possible. It adds `citus.enable_or_clause_arm_pruning` to prune irrelevant OR-clause arms per shard. It also restricts internal metadata cleanup to superusers, uses an empty search path for internal SPI queries, and replays object definitions under their owner's restricted execution context.

Coordinate the upgrade across all nodes. After installing matching binaries and restarting PostgreSQL, run the SQL update in each database containing Citus:

```sql
ALTER EXTENSION citus UPDATE TO '14.3-1';
SELECT extversion FROM pg_extension WHERE extname = 'citus';
SELECT citus_version();
```

Finish the cluster upgrade on the coordinator:

```sql
CALL citus_finish_citus_upgrade();
```

Use the SQL/control version in ALTER EXTENSION, not the package version. If both Citus and PostgreSQL need upgrading, complete the Citus upgrade first and follow the separate PostgreSQL-major upgrade procedure.
