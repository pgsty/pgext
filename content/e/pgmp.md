---
title: "pgmp"
linkTitle: "pgmp"
description: "Multiple Precision Arithmetic extension"
weight: 3700
categories: ["TYPE"]
width: full
---

[**pgmp**](https://github.com/dvarrazzo/pgmp/) : Multiple Precision Arithmetic extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3700** | {{< badge content="pgmp" link="https://github.com/dvarrazzo/pgmp/" >}} | {{< ext "pgmp" >}} | `1.0.6` | {{< category "TYPE" >}} | {{< license "LGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "unit" >}} {{< ext "numeral" >}} {{< ext "pg_rational" >}} {{< ext "uint" >}} {{< ext "uint128" >}} {{< ext "seg" >}} {{< ext "cube" >}} |

> [!Note] PIGSTY RPM and DEB packages are aligned at 1.0.6 for PostgreSQL 14 through 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmp` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "pgmp_18" "green" >}} {{< bg "17" "pgmp_17" "green" >}} {{< bg "16" "pgmp_16" "green" >}} {{< bg "15" "pgmp_15" "green" >}} {{< bg "14" "pgmp_14" "green" >}} | `pgmp_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.6` | {{< bg "18" "postgresql-18-pgmp" "green" >}} {{< bg "17" "postgresql-17-pgmp" "green" >}} {{< bg "16" "postgresql-16-pgmp" "green" >}} {{< bg "15" "postgresql-15-pgmp" "green" >}} {{< bg "14" "postgresql-14-pgmp" "green" >}} | `postgresql-$v-pgmp` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 4" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0.6" "pgmp_14 : AVAIL 4" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.0.6" "postgresql-18-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-17-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-16-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-15-pgmp : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.6" "postgresql-14-pgmp : AVAIL 4" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_18` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.1 KiB | [pgmp_18-1.0.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmp_18-1.0.6-1PIGSTY.el8.x86_64.rpm) |
| `pgmp_18` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pgmp_18-1.0.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmp_18-1.0.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgmp_18-1.0.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmp_18-1.0.5-4PGDG.rhel8.x86_64.rpm) |
| `pgmp_18` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.8 KiB | [pgmp_18-1.0.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmp_18-1.0.6-1PIGSTY.el8.aarch64.rpm) |
| `pgmp_18` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pgmp_18-1.0.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmp_18-1.0.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmp_18-1.0.5-4PGDG.rhel8.aarch64.rpm) |
| `pgmp_18` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.6 KiB | [pgmp_18-1.0.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmp_18-1.0.6-1PIGSTY.el9.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pgmp_18-1.0.5-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmp_18-1.0.5-6PGDG.rhel9.8.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmp_18-1.0.5-4PGDG.rhel9.x86_64.rpm) |
| `pgmp_18` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.3 KiB | [pgmp_18-1.0.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmp_18-1.0.6-1PIGSTY.el9.aarch64.rpm) |
| `pgmp_18` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pgmp_18-1.0.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmp_18-1.0.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pgmp_18-1.0.5-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmp_18-1.0.5-6PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pgmp_18-1.0.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmp_18-1.0.5-4PGDG.rhel9.aarch64.rpm) |
| `pgmp_18` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.6 KiB | [pgmp_18-1.0.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmp_18-1.0.6-1PIGSTY.el10.x86_64.rpm) |
| `pgmp_18` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.9 KiB | [pgmp_18-1.0.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmp_18-1.0.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pgmp_18-1.0.5-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmp_18-1.0.5-6PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pgmp_18-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmp_18-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_18` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.5 KiB | [pgmp_18-1.0.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmp_18-1.0.6-1PIGSTY.el10.aarch64.rpm) |
| `pgmp_18` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_18-1.0.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmp_18-1.0.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.6 KiB | [pgmp_18-1.0.5-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmp_18-1.0.5-6PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_18-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmp_18-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.2 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.7 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.5 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.7 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.4 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.2 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.5 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.5 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg13+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.1 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.6 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.6 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg13+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.5 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 104.6 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.2 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.8 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.7 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.6 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 91.3 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 100.9 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 100.9 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.5 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.0 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.1 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.1 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.4 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.8 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.8 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.6 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 90.0 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.3 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.3 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.8 KiB | [postgresql-18-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 88.6 KiB | [postgresql-18-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-18-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.7 KiB | [postgresql-18-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.9 KiB | [postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-18-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_17` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.0 KiB | [pgmp_17-1.0.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmp_17-1.0.6-1PIGSTY.el8.x86_64.rpm) |
| `pgmp_17` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pgmp_17-1.0.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmp_17-1.0.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [pgmp_17-1.0.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmp_17-1.0.5-3PGDG.rhel8.x86_64.rpm) |
| `pgmp_17` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.8 KiB | [pgmp_17-1.0.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmp_17-1.0.6-1PIGSTY.el8.aarch64.rpm) |
| `pgmp_17` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pgmp_17-1.0.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmp_17-1.0.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [pgmp_17-1.0.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmp_17-1.0.5-3PGDG.rhel8.aarch64.rpm) |
| `pgmp_17` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.8 KiB | [pgmp_17-1.0.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmp_17-1.0.6-1PIGSTY.el9.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pgmp_17-1.0.5-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmp_17-1.0.5-6PGDG.rhel9.8.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pgmp_17-1.0.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmp_17-1.0.5-3PGDG.rhel9.x86_64.rpm) |
| `pgmp_17` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [pgmp_17-1.0.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmp_17-1.0.6-1PIGSTY.el9.aarch64.rpm) |
| `pgmp_17` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pgmp_17-1.0.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmp_17-1.0.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.5 KiB | [pgmp_17-1.0.5-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmp_17-1.0.5-6PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pgmp_17-1.0.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmp_17-1.0.5-3PGDG.rhel9.aarch64.rpm) |
| `pgmp_17` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.6 KiB | [pgmp_17-1.0.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmp_17-1.0.6-1PIGSTY.el10.x86_64.rpm) |
| `pgmp_17` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pgmp_17-1.0.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmp_17-1.0.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pgmp_17-1.0.5-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmp_17-1.0.5-6PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pgmp_17-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmp_17-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_17` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.6 KiB | [pgmp_17-1.0.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmp_17-1.0.6-1PIGSTY.el10.aarch64.rpm) |
| `pgmp_17` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_17-1.0.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmp_17-1.0.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [pgmp_17-1.0.5-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmp_17-1.0.5-6PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_17-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmp_17-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.3 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.6 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.8 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.1 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.4 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.2 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.7 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg13+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.5 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.0 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.6 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.6 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg13+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.7 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 111.6 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.5 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.3 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.2 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 109.5 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.5 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.3 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.4 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.8 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.0 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.2 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.1 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.4 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.0 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.0 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.4 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 90.0 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.1 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.3 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 100.0 KiB | [postgresql-17-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 88.7 KiB | [postgresql-17-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-17-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.6 KiB | [postgresql-17-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.9 KiB | [postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-17-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_16` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.1 KiB | [pgmp_16-1.0.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmp_16-1.0.6-1PIGSTY.el8.x86_64.rpm) |
| `pgmp_16` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pgmp_16-1.0.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmp_16-1.0.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pgmp_16-1.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmp_16-1.0.5-1PGDG.rhel8.x86_64.rpm) |
| `pgmp_16` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.8 KiB | [pgmp_16-1.0.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmp_16-1.0.6-1PIGSTY.el8.aarch64.rpm) |
| `pgmp_16` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pgmp_16-1.0.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmp_16-1.0.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pgmp_16-1.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmp_16-1.0.5-1PGDG.rhel8.aarch64.rpm) |
| `pgmp_16` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.8 KiB | [pgmp_16-1.0.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmp_16-1.0.6-1PIGSTY.el9.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pgmp_16-1.0.5-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmp_16-1.0.5-6PGDG.rhel9.8.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmp_16-1.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmp_16-1.0.5-1PGDG.rhel9.x86_64.rpm) |
| `pgmp_16` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.5 KiB | [pgmp_16-1.0.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmp_16-1.0.6-1PIGSTY.el9.aarch64.rpm) |
| `pgmp_16` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [pgmp_16-1.0.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmp_16-1.0.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.5 KiB | [pgmp_16-1.0.5-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmp_16-1.0.5-6PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.0 KiB | [pgmp_16-1.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmp_16-1.0.5-1PGDG.rhel9.aarch64.rpm) |
| `pgmp_16` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.6 KiB | [pgmp_16-1.0.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmp_16-1.0.6-1PIGSTY.el10.x86_64.rpm) |
| `pgmp_16` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.9 KiB | [pgmp_16-1.0.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmp_16-1.0.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pgmp_16-1.0.5-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmp_16-1.0.5-6PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pgmp_16-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmp_16-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_16` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.6 KiB | [pgmp_16-1.0.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmp_16-1.0.6-1PIGSTY.el10.aarch64.rpm) |
| `pgmp_16` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_16-1.0.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmp_16-1.0.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [pgmp_16-1.0.5-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmp_16-1.0.5-6PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.9 KiB | [pgmp_16-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmp_16-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.1 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.6 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.7 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.1 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.2 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.7 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg13+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.6 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.0 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.5 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.7 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg13+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.5 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 111.4 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.4 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.0 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.1 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 109.4 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.4 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.3 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.4 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.8 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.0 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.2 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.3 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.1 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.4 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 100.0 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.8 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.6 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 90.0 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.9 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 101.2 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.9 KiB | [postgresql-16-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 88.7 KiB | [postgresql-16-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-16-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.7 KiB | [postgresql-16-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.8 KiB | [postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-16-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_15` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.2 KiB | [pgmp_15-1.0.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmp_15-1.0.6-1PIGSTY.el8.x86_64.rpm) |
| `pgmp_15` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pgmp_15-1.0.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmp_15-1.0.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pgmp_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.9 KiB | [pgmp_15-1.0.4-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmp_15-1.0.4-4.rhel8.x86_64.rpm) |
| `pgmp_15` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.9 KiB | [pgmp_15-1.0.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmp_15-1.0.6-1PIGSTY.el8.aarch64.rpm) |
| `pgmp_15` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pgmp_15-1.0.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmp_15-1.0.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pgmp_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [pgmp_15-1.0.4-4.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmp_15-1.0.4-4.rhel8.aarch64.rpm) |
| `pgmp_15` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.5 KiB | [pgmp_15-1.0.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmp_15-1.0.6-1PIGSTY.el9.x86_64.rpm) |
| `pgmp_15` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.1 KiB | [pgmp_15-1.0.5-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmp_15-1.0.5-6PGDG.rhel9.8.x86_64.rpm) |
| `pgmp_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.4 KiB | [pgmp_15-1.0.4-4.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmp_15-1.0.4-4.rhel9.x86_64.rpm) |
| `pgmp_15` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmp_15-1.0.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmp_15-1.0.6-1PIGSTY.el9.aarch64.rpm) |
| `pgmp_15` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [pgmp_15-1.0.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmp_15-1.0.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_15` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pgmp_15-1.0.5-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmp_15-1.0.5-6PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.1 KiB | [pgmp_15-1.0.4-4.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmp_15-1.0.4-4.rhel9.aarch64.rpm) |
| `pgmp_15` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 45.5 KiB | [pgmp_15-1.0.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmp_15-1.0.6-1PIGSTY.el10.x86_64.rpm) |
| `pgmp_15` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.7 KiB | [pgmp_15-1.0.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmp_15-1.0.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.1 KiB | [pgmp_15-1.0.5-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmp_15-1.0.5-6PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.4 KiB | [pgmp_15-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmp_15-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_15` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.2 KiB | [pgmp_15-1.0.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmp_15-1.0.6-1PIGSTY.el10.aarch64.rpm) |
| `pgmp_15` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pgmp_15-1.0.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmp_15-1.0.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pgmp_15-1.0.5-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmp_15-1.0.5-6PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pgmp_15-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmp_15-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.2 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.7 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.7 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.3 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.0 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.9 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg13+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.0 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.5 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.9 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg13+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 111.3 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.2 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.9 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 109.4 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.4 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.4 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.3 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.2 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.7 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 100.8 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.9 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.1 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.7 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.7 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.5 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 88.7 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.1 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.5 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.2 KiB | [postgresql-15-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 88.0 KiB | [postgresql-15-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-15-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.3 KiB | [postgresql-15-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.5 KiB | [postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-15-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmp_14` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.2 KiB | [pgmp_14-1.0.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmp_14-1.0.6-1PIGSTY.el8.x86_64.rpm) |
| `pgmp_14` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pgmp_14-1.0.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmp_14-1.0.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pgmp_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.6 KiB | [pgmp_14-1.0.4-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmp_14-1.0.4-4.rhel8.x86_64.rpm) |
| `pgmp_14` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.9 KiB | [pgmp_14-1.0.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmp_14-1.0.6-1PIGSTY.el8.aarch64.rpm) |
| `pgmp_14` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pgmp_14-1.0.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmp_14-1.0.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pgmp_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.7 KiB | [pgmp_14-1.0.4-4.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmp_14-1.0.4-4.rhel8.aarch64.rpm) |
| `pgmp_14` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.4 KiB | [pgmp_14-1.0.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmp_14-1.0.6-1PIGSTY.el9.x86_64.rpm) |
| `pgmp_14` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.1 KiB | [pgmp_14-1.0.5-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmp_14-1.0.5-6PGDG.rhel9.8.x86_64.rpm) |
| `pgmp_14` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.6 KiB | [pgmp_14-1.0.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmp_14-1.0.6-1PIGSTY.el9.aarch64.rpm) |
| `pgmp_14` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [pgmp_14-1.0.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmp_14-1.0.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_14` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pgmp_14-1.0.5-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmp_14-1.0.5-6PGDG.rhel9.8.aarch64.rpm) |
| `pgmp_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.1 KiB | [pgmp_14-1.0.4-4.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmp_14-1.0.4-4.rhel9.aarch64.rpm) |
| `pgmp_14` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 45.2 KiB | [pgmp_14-1.0.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmp_14-1.0.6-1PIGSTY.el10.x86_64.rpm) |
| `pgmp_14` | `1.0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.7 KiB | [pgmp_14-1.0.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmp_14-1.0.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.1 KiB | [pgmp_14-1.0.5-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmp_14-1.0.5-6PGDG.rhel10.2.x86_64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.4 KiB | [pgmp_14-1.0.5-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmp_14-1.0.5-4PGDG.rhel10.x86_64.rpm) |
| `pgmp_14` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.0 KiB | [pgmp_14-1.0.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmp_14-1.0.6-1PIGSTY.el10.aarch64.rpm) |
| `pgmp_14` | `1.0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pgmp_14-1.0.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmp_14-1.0.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.1 KiB | [pgmp_14-1.0.5-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmp_14-1.0.5-6PGDG.rhel10.2.aarch64.rpm) |
| `pgmp_14` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pgmp_14-1.0.5-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmp_14-1.0.5-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.1 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.7 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.8 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg12+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.2 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg12+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.3 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.8 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.9 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg13+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.8 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg13+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.0 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.5 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg13+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg13+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 111.4 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 100.2 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.7 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 108.7 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 109.4 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.3 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.3 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 107.2 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.2 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.6 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 101.0 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 100.8 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.1 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.7 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 99.9 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.4 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 88.7 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.1 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 100.3 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.2 KiB | [postgresql-14-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 87.9 KiB | [postgresql-14-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmp/postgresql-14-pgmp_1.0.6-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.2 KiB | [postgresql-14-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-5.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgmp` | `1.0.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 99.3 KiB | [postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pgmp/postgresql-14-pgmp_1.0.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dvarrazzo/pgmp/" title="Repository" icon="github" subtitle="github.com/dvarrazzo/pgmp/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmp-1.0.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmp;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmp;		# install via package name, for the active PG version

pig install pgmp -v 18;   # install for PG 18
pig install pgmp -v 17;   # install for PG 17
pig install pgmp -v 16;   # install for PG 16
pig install pgmp -v 15;   # install for PG 15
pig install pgmp -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmp;
```

## Usage

Sources:

- [pgmp 1.0.6 README](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/README.rst)
- [pgmp 1.0.6 release notes](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/NEWS.rst)
- [pgmp 1.0.6 metadata](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/META.json)
- [pgmp control file](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/pgmp.control)
- [Official pgmp documentation](https://dvarrazzo.github.io/pgmp/)

`pgmp` exposes GNU MP arithmetic inside PostgreSQL. It adds arbitrary-size integer values through `mpz` and exact rational values through `mpq`, together with casts, arithmetic, comparison, aggregate, number-theory, bit, and random-number functions.

### Core Workflow

```sql
CREATE EXTENSION pgmp;

SELECT '123456789012345678901234567890'::mpz * 2;
SELECT mpq(1::mpz, 3::mpz) + mpq(1::mpz, 6::mpz);
SELECT gcd(48::mpz, 18::mpz);
SELECT nextprime(100000000000000000000::mpz);
```

`mpz` is an arbitrary-size integer type, subject to PostgreSQL's value-size limits. `mpq` stores a canonical numerator and denominator so fractional arithmetic remains exact until explicitly converted to an approximate type.

### Important Objects

- `mpz(text)` and casts construct integers in decimal or supported base-prefixed forms.
- `mpq(text)` and `mpq(mpz, mpz)` construct rational values.
- Both types support ordinary comparisons and btree or hash indexes.
- Integer helpers include division with explicit rounding modes, powers, roots, primality tests, `gcd`, `lcm`, factorials, Fibonacci and Lucas numbers, bit operations, and random-state functions.
- Rational helpers include numerator and denominator access, inversion, denominator limiting, arithmetic, comparison, and aggregates.
- `gmp_version()` and `gmp_max_bitcnt()` expose library information.

Do not use floating-point input when exact decimal or rational meaning matters; construct values from text, integers, or explicit numerator and denominator values.

### Version 1.0.6 Notes

The 1.0.6 distribution adds PostgreSQL 19 build compatibility, sets PostgreSQL 14 as the supported runtime floor in its metadata, and adds missing unsigned-long range checks for the power, Fibonacci, and Lucas-number paths.

The upstream distribution version is 1.0.6, while its tagged `pgmp.control` currently declares SQL extension version `1.1`. Create the extension without forcing a version and inspect the database-reported value before designing an upgrade:

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmp';
```

pgmp requires the GMP shared library. GMP 4.1 lacks a few functions documented by upstream, including some root, bit, and random-state helpers; use a current GMP release when those objects are required. Large operands can consume substantial backend memory and CPU, so apply statement timeouts and input limits to untrusted arithmetic workloads.
