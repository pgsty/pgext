---
title: "citus"
linkTitle: "citus"
description: "Distributed PostgreSQL as an extension"
weight: 2400
categories: ["OLAP"]
width: full
---

[**citus**](https://github.com/citusdata/citus) : Distributed PostgreSQL as an extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2400** | {{< badge content="citus" link="https://github.com/citusdata/citus" >}} | {{< ext "citus" >}} | `14.1.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "documentdb_distributed" >}} |
|   **See Also**    | {{< ext "pg_partman" >}} {{< ext "plproxy" >}} {{< ext "columnar" >}} {{< ext "pg_fkpart" >}} {{< ext "timescaledb" >}} {{< ext "pg_duckdb" >}} {{< ext "tablefunc" >}} {{< ext "hll" >}} |
|    **Siblings**   | {{< ext "citus_columnar" >}} |

> [!Note] conflict with hydra


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.1.0` | {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} {{< bg "14" "citus_14" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.1.0` | {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} {{< bg "14" "postgresql-14-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 16" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 28" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 16" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 18" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 25" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 18" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 15" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 9" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 4" "blue" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.1.0" "citus_16 : AVAIL 9" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 4" "blue" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.1.0" "postgresql-16-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_18` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 983.9 KiB | [citus_18-14.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_18-14.1.0-1PIGSTY.el8.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 871.9 KiB | [citus_18-14.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 859.5 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 943.8 KiB | [citus_18-14.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_18-14.1.0-1PIGSTY.el8.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 822.0 KiB | [citus_18-14.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.2 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 881.8 KiB | [citus_18-14.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_18-14.1.0-1PIGSTY.el9.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 837.2 KiB | [citus_18-14.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 823.8 KiB | [citus_18-14.0.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.0.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 823.9 KiB | [citus_18-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 824.3 KiB | [citus_18-14.0.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.0.0-1PGDG.rhel9.6.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 835.7 KiB | [citus_18-14.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_18-14.1.0-1PIGSTY.el9.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 808.3 KiB | [citus_18-14.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 797.4 KiB | [citus_18-14.0.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.0.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 797.3 KiB | [citus_18-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 797.8 KiB | [citus_18-14.0.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.0.0-1PGDG.rhel9.6.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 875.7 KiB | [citus_18-14.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_18-14.1.0-1PIGSTY.el10.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 849.0 KiB | [citus_18-14.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 836.1 KiB | [citus_18-14.0.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.0.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 835.4 KiB | [citus_18-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 835.7 KiB | [citus_18-14.0.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.0.0-1PGDG.rhel10.0.x86_64.rpm) |
| `citus_18` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 843.7 KiB | [citus_18-14.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_18-14.1.0-1PIGSTY.el10.aarch64.rpm) |
| `citus_18` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 818.3 KiB | [citus_18-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 805.9 KiB | [citus_18-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 806.0 KiB | [citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 806.1 KiB | [citus_18-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-citus` | `14.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-citus` | `14.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-citus` | `14.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-citus` | `14.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-citus` | `14.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-18-citus_14.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_17` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 981.6 KiB | [citus_17-14.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-14.1.0-1PIGSTY.el8.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 869.7 KiB | [citus_17-14.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 857.0 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 850.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 827.1 KiB | [citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.7 KiB | [citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 803.6 KiB | [citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 941.1 KiB | [citus_17-14.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-14.1.0-1PIGSTY.el8.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 819.9 KiB | [citus_17-14.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 808.6 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 802.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 781.5 KiB | [citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.4 KiB | [citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.5 KiB | [citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.2 KiB | [citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 758.9 KiB | [citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 879.9 KiB | [citus_17-14.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-14.1.0-1PIGSTY.el9.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 834.4 KiB | [citus_17-14.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 821.6 KiB | [citus_17-14.0.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.0.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 821.5 KiB | [citus_17-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 822.0 KiB | [citus_17-14.0.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.0.0-1PGDG.rhel9.6.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.8 KiB | [citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 793.2 KiB | [citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.1 KiB | [citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.6 KiB | [citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.4 KiB | [citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 772.0 KiB | [citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 833.6 KiB | [citus_17-14.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-14.1.0-1PIGSTY.el9.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 807.1 KiB | [citus_17-14.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 795.4 KiB | [citus_17-14.0.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.0.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 795.4 KiB | [citus_17-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 795.9 KiB | [citus_17-14.0.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.0.0-1PGDG.rhel9.6.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.2 KiB | [citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 768.0 KiB | [citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 750.0 KiB | [citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 746.8 KiB | [citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 872.1 KiB | [citus_17-14.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_17-14.1.0-1PIGSTY.el10.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 845.9 KiB | [citus_17-14.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 833.0 KiB | [citus_17-14.0.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.0.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 832.8 KiB | [citus_17-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 833.0 KiB | [citus_17-14.0.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.0.0-1PGDG.rhel10.0.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 827.5 KiB | [citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 804.7 KiB | [citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.1 KiB | [citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.4 KiB | [citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 841.8 KiB | [citus_17-14.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_17-14.1.0-1PIGSTY.el10.aarch64.rpm) |
| `citus_17` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 816.4 KiB | [citus_17-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 804.0 KiB | [citus_17-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 803.8 KiB | [citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 803.7 KiB | [citus_17-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 798.8 KiB | [citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 776.8 KiB | [citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 758.8 KiB | [citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 759.3 KiB | [citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-citus` | `14.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-citus` | `14.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-citus` | `14.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-citus` | `14.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.4 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-citus` | `14.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-17-citus_14.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_16` | `14.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 973.4 KiB | [citus_16-14.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-14.1.0-1PIGSTY.el8.x86_64.rpm) |
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
| `citus_16` | `14.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 934.9 KiB | [citus_16-14.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-14.1.0-1PIGSTY.el8.aarch64.rpm) |
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
| `citus_16` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 874.2 KiB | [citus_16-14.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-14.1.0-1PIGSTY.el9.x86_64.rpm) |
| `citus_16` | `14.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 828.2 KiB | [citus_16-14.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.4 KiB | [citus_16-14.0.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.0.0-1PGDG.rhel9.8.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.4 KiB | [citus_16-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.8 KiB | [citus_16-14.0.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.0.0-1PGDG.rhel9.6.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 809.6 KiB | [citus_16-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `13.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 786.8 KiB | [citus_16-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `13.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 768.5 KiB | [citus_16-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `13.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 769.2 KiB | [citus_16-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `13.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 768.9 KiB | [citus_16-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 766.9 KiB | [citus_16-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 767.5 KiB | [citus_16-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 766.6 KiB | [citus_16-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 766.4 KiB | [citus_16-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 766.2 KiB | [citus_16-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 765.2 KiB | [citus_16-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 765.2 KiB | [citus_16-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `12.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 765.4 KiB | [citus_16-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_16` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 826.7 KiB | [citus_16-14.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-14.1.0-1PIGSTY.el9.aarch64.rpm) |
| `citus_16` | `14.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 800.3 KiB | [citus_16-14.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 788.9 KiB | [citus_16-14.0.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.0.0-1PGDG.rhel9.8.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 788.9 KiB | [citus_16-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.6 KiB | [citus_16-14.0.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.0.0-1PGDG.rhel9.6.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 782.8 KiB | [citus_16-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 761.8 KiB | [citus_16-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 744.1 KiB | [citus_16-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 744.1 KiB | [citus_16-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `13.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 743.8 KiB | [citus_16-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 741.7 KiB | [citus_16-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 742.2 KiB | [citus_16-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 741.4 KiB | [citus_16-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 741.6 KiB | [citus_16-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 740.4 KiB | [citus_16-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 739.1 KiB | [citus_16-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 738.8 KiB | [citus_16-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `12.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 738.9 KiB | [citus_16-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `citus_16` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 865.1 KiB | [citus_16-14.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_16-14.1.0-1PIGSTY.el10.x86_64.rpm) |
| `citus_16` | `14.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 838.2 KiB | [citus_16-14.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 826.0 KiB | [citus_16-14.0.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.0.0-1PGDG.rhel10.2.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 826.1 KiB | [citus_16-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 826.4 KiB | [citus_16-14.0.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.0.0-1PGDG.rhel10.0.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 820.4 KiB | [citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 797.3 KiB | [citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.1 KiB | [citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.4 KiB | [citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 834.8 KiB | [citus_16-14.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_16-14.1.0-1PIGSTY.el10.aarch64.rpm) |
| `citus_16` | `14.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 809.2 KiB | [citus_16-14.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.5 KiB | [citus_16-14.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.4 KiB | [citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.3 KiB | [citus_16-14.0.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.0.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 792.1 KiB | [citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 770.0 KiB | [citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.3 KiB | [citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.7 KiB | [citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-citus` | `14.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-citus` | `14.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-citus` | `14.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-citus` | `14.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-citus` | `14.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/citus/postgresql-16-citus_14.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `citus_15` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 855.4 KiB | [citus_15-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `13.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 831.8 KiB | [citus_15-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `13.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 813.3 KiB | [citus_15-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `13.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 813.5 KiB | [citus_15-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `13.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 813.3 KiB | [citus_15-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 811.4 KiB | [citus_15-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 809.2 KiB | [citus_15-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 809.5 KiB | [citus_15-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 809.6 KiB | [citus_15-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 809.1 KiB | [citus_15-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 805.9 KiB | [citus_15-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 806.0 KiB | [citus_15-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 807.2 KiB | [citus_15-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `12.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 810.5 KiB | [citus_15-12.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-12.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_15` | `11.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 791.4 KiB | [citus_15-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_15` | `11.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 769.9 KiB | [citus_15-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_15` | `11.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 769.1 KiB | [citus_15-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_15` | `11.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 755.7 KiB | [citus_15-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_15` | `11.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 757.4 KiB | [citus_15-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_15` | `11.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 757.2 KiB | [citus_15-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_15` | `11.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 756.6 KiB | [citus_15-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/citus_15-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_15` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 819.5 KiB | [citus_15-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `13.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 798.4 KiB | [citus_15-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `13.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 780.6 KiB | [citus_15-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `13.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 780.8 KiB | [citus_15-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `13.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 780.7 KiB | [citus_15-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 778.6 KiB | [citus_15-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 778.1 KiB | [citus_15-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 777.2 KiB | [citus_15-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 777.4 KiB | [citus_15-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 774.8 KiB | [citus_15-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 773.9 KiB | [citus_15-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 773.8 KiB | [citus_15-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 773.7 KiB | [citus_15-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `12.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 775.5 KiB | [citus_15-12.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-12.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_15` | `11.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 759.0 KiB | [citus_15-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_15` | `11.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 738.2 KiB | [citus_15-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_15` | `11.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 737.5 KiB | [citus_15-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_15` | `11.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 726.0 KiB | [citus_15-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_15` | `11.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 724.9 KiB | [citus_15-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_15` | `11.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 724.7 KiB | [citus_15-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.3-1.rhel9.aarch64.rpm) |
| `citus_15` | `11.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 724.3 KiB | [citus_15-11.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/citus_15-11.1.2-1.rhel9.aarch64.rpm) |
| `citus_15` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 858.2 KiB | [citus_15-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 837.8 KiB | [citus_15-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 818.2 KiB | [citus_15-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 818.2 KiB | [citus_15-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 826.0 KiB | [citus_15-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 805.6 KiB | [citus_15-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 787.5 KiB | [citus_15-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_15` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 787.6 KiB | [citus_15-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/citus_15-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-citus` | `13.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-citus` | `13.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-citus` | `13.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~trixie_amd64.deb) |
| `postgresql-15-citus` | `13.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~trixie_arm64.deb) |
| `postgresql-15-citus` | `13.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~jammy_amd64.deb) |
| `postgresql-15-citus` | `13.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~jammy_arm64.deb) |
| `postgresql-15-citus` | `13.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~noble_amd64.deb) |
| `postgresql-15-citus` | `13.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-15-citus_13.2.0-8PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-15-citus_13.2.0-8PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_14` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 814.2 KiB | [citus_14-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 813.6 KiB | [citus_14-12.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.6-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.5 KiB | [citus_14-12.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.5-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.5 KiB | [citus_14-12.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.3 KiB | [citus_14-12.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 811.1 KiB | [citus_14-12.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 811.0 KiB | [citus_14-12.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.1-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 810.5 KiB | [citus_14-12.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.1.0-2PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `12.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 813.4 KiB | [citus_14-12.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-12.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_14` | `11.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 796.7 KiB | [citus_14-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_14` | `11.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 776.9 KiB | [citus_14-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 776.0 KiB | [citus_14-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 765.6 KiB | [citus_14-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 765.3 KiB | [citus_14-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 765.1 KiB | [citus_14-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 764.4 KiB | [citus_14-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 762.6 KiB | [citus_14-11.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.1.1-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 701.4 KiB | [citus_14-11.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.6-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 700.3 KiB | [citus_14-11.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.5-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 699.6 KiB | [citus_14-11.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.4-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 699.5 KiB | [citus_14-11.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.3-1.rhel8.x86_64.rpm) |
| `citus_14` | `11.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 698.6 KiB | [citus_14-11.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-11.0.2-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 618.5 KiB | [citus_14-10.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.5-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 618.6 KiB | [citus_14-10.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.4-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 618.5 KiB | [citus_14-10.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.3-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 615.3 KiB | [citus_14-10.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.2-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 614.8 KiB | [citus_14-10.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.1-1.rhel8.x86_64.rpm) |
| `citus_14` | `10.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 614.0 KiB | [citus_14-10.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/citus_14-10.2.0-1.rhel8.x86_64.rpm) |
| `citus_14` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 765.7 KiB | [citus_14-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 764.6 KiB | [citus_14-12.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.6-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 763.5 KiB | [citus_14-12.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.5-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 763.4 KiB | [citus_14-12.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 763.4 KiB | [citus_14-12.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 762.3 KiB | [citus_14-12.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 762.2 KiB | [citus_14-12.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.1-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.9 KiB | [citus_14-12.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.1.0-2PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `12.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 763.9 KiB | [citus_14-12.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-12.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_14` | `11.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 748.9 KiB | [citus_14-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_14` | `11.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 729.5 KiB | [citus_14-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_14` | `11.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 729.4 KiB | [citus_14-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_14` | `11.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 718.0 KiB | [citus_14-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_14` | `11.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 717.7 KiB | [citus_14-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_14` | `11.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 717.4 KiB | [citus_14-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/citus_14-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_14` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 803.4 KiB | [citus_14-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 802.9 KiB | [citus_14-12.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.6-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 800.4 KiB | [citus_14-12.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.5-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 800.6 KiB | [citus_14-12.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 800.1 KiB | [citus_14-12.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 798.9 KiB | [citus_14-12.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 798.9 KiB | [citus_14-12.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.1-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 798.9 KiB | [citus_14-12.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.1.0-2PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `12.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 802.0 KiB | [citus_14-12.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-12.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_14` | `11.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 787.6 KiB | [citus_14-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_14` | `11.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 767.2 KiB | [citus_14-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 766.5 KiB | [citus_14-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 756.8 KiB | [citus_14-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 755.9 KiB | [citus_14-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 755.8 KiB | [citus_14-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 755.4 KiB | [citus_14-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 754.1 KiB | [citus_14-11.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.1.1-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 691.3 KiB | [citus_14-11.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.6-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 690.5 KiB | [citus_14-11.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.5-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 690.0 KiB | [citus_14-11.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.4-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 689.8 KiB | [citus_14-11.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.3-1.rhel9.x86_64.rpm) |
| `citus_14` | `11.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 689.0 KiB | [citus_14-11.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-11.0.2-1.rhel9.x86_64.rpm) |
| `citus_14` | `10.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 612.1 KiB | [citus_14-10.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.5-1.rhel9.x86_64.rpm) |
| `citus_14` | `10.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 613.7 KiB | [citus_14-10.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.4-1.rhel9.x86_64.rpm) |
| `citus_14` | `10.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 613.7 KiB | [citus_14-10.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/citus_14-10.2.3-1.rhel9.x86_64.rpm) |
| `citus_14` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 771.1 KiB | [citus_14-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 770.5 KiB | [citus_14-12.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.6-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 770.0 KiB | [citus_14-12.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.5-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 770.2 KiB | [citus_14-12.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 768.2 KiB | [citus_14-12.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 766.8 KiB | [citus_14-12.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 766.5 KiB | [citus_14-12.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.1-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 766.8 KiB | [citus_14-12.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.1.0-2PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `12.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 769.0 KiB | [citus_14-12.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-12.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_14` | `11.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 754.5 KiB | [citus_14-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_14` | `11.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 736.1 KiB | [citus_14-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_14` | `11.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 735.2 KiB | [citus_14-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_14` | `11.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 724.2 KiB | [citus_14-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_14` | `11.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 723.3 KiB | [citus_14-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_14` | `11.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 723.1 KiB | [citus_14-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/citus_14-11.1.3-1.rhel9.aarch64.rpm) |
| `postgresql-14-citus` | `13.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-citus` | `13.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.9 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-citus` | `13.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-citus` | `13.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-citus` | `13.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-citus` | `13.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-14-citus_13.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-14-citus_13.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install citus;		# install via package name, for the active PG version

pig install citus -v 18;   # install for PG 18
pig install citus -v 17;   # install for PG 17
pig install citus -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'citus';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citus;
```




## Usage

Sources:

- [Citus v14.1.0 release](https://github.com/citusdata/citus/releases/tag/v14.1.0)
- [Citus v14.1.0 CHANGELOG](https://github.com/citusdata/citus/blob/v14.1.0/CHANGELOG.md)
- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Citus Utility Functions](https://docs.citusdata.com/en/stable/develop/api_udf.html)
- [Local package metadata](../db/extension.csv)

Citus turns PostgreSQL into a distributed database by sharding tables across worker nodes while keeping PostgreSQL SQL, indexes, extensions, transactions, and operational tooling as the user-facing surface. It is commonly used for multi-tenant SaaS databases, real-time analytics, time-series/event workloads, and distributed microservice schemas.

The local Pigsty catalog packages Citus as `citus` and exposes the lead extension `citus`; the same package also contains `citus_columnar`. Citus is a preload extension, so every node must load the library before `CREATE EXTENSION`.

### Enable Citus

```conf
shared_preload_libraries = 'citus'
```

Restart PostgreSQL on the coordinator and workers, then create the extension in the database:

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

SELECT create_distributed_table('events', 'tenant_id');
```

You can tune the shard count and colocation explicitly:

```sql
SELECT create_distributed_table(
  'events',
  'tenant_id',
  shard_count  := 64,
  colocate_with := 'default'
);
```

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

Schema-based sharding is useful when each tenant or service owns its own schema. In v14.1.0, Citus adds support for running several schema-sharding DDLs from any node, including `CREATE SCHEMA`, `DROP SCHEMA`, `ALTER SCHEMA RENAME`, `ALTER SCHEMA OWNER`, and table-level DDL on distributed schemas.

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

### Caveats

- Pigsty local metadata currently tracks Citus 14.x for PostgreSQL 16-18; Citus 14 dropped PostgreSQL 15 support.
- `shared_preload_libraries = 'citus'` must be set before extension creation. A plain `CREATE EXTENSION citus` is not enough on a fresh server.
- Choose the distribution column carefully. Primary keys and unique constraints on distributed tables generally need to include the distribution column.
- Cross-shard joins, repartition joins, distributed DDL, and multi-shard writes are powerful but have different planning and locking behavior from single-node PostgreSQL.
- Citus includes its own columnar storage surface through `citus_columnar`; Pigsty metadata marks it as conflicting with Hydra `columnar`.
- The cluster-change blocking functions are operational tools for backups. Do not leave a cluster blocked after a failed backup script.
