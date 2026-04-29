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
| **2400** | {{< badge content="citus" link="https://github.com/citusdata/citus" >}} | {{< ext "citus" >}} | `14.0.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} {{< bg "14" "citus_14" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} {{< bg "14" "postgresql-14-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 29" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 21" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 26" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_18` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 969.6 KiB | [citus_18-14.0.0-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_18-14.0.0-4PIGSTY.el8.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 859.5 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/citus_18-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 930.4 KiB | [citus_18-14.0.0-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_18-14.0.0-4PIGSTY.el8.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.2 KiB | [citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/citus_18-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 850.7 KiB | [citus_18-14.0.0-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_18-14.0.0-4PIGSTY.el9.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 823.9 KiB | [citus_18-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/citus_18-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 823.4 KiB | [citus_18-14.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_18-14.0.0-4PIGSTY.el9.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 797.3 KiB | [citus_18-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/citus_18-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 861.1 KiB | [citus_18-14.0.0-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_18-14.0.0-4PIGSTY.el10.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 835.4 KiB | [citus_18-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/citus_18-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 831.6 KiB | [citus_18-14.0.0-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_18-14.0.0-4PIGSTY.el10.aarch64.rpm) |
| `citus_18` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 806.0 KiB | [citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/citus_18-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-citus` | `14.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-citus` | `14.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-citus` | `14.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-18-citus` | `14.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-18-citus` | `14.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-18-citus` | `14.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-18-citus` | `14.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-18-citus` | `14.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-18-citus_14.0.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-18-citus_14.0.0-4PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_17` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 967.0 KiB | [citus_17-14.0.0-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-14.0.0-4PIGSTY.el8.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 857.0 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-14.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 850.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 827.1 KiB | [citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.7 KiB | [citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 803.6 KiB | [citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 928.5 KiB | [citus_17-14.0.0-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-14.0.0-4PIGSTY.el8.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 808.6 KiB | [citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-14.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 802.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 781.5 KiB | [citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.4 KiB | [citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.5 KiB | [citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.2 KiB | [citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 758.9 KiB | [citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 847.8 KiB | [citus_17-14.0.0-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-14.0.0-4PIGSTY.el9.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 821.5 KiB | [citus_17-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.8 KiB | [citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 793.2 KiB | [citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.1 KiB | [citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.6 KiB | [citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.4 KiB | [citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 772.0 KiB | [citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 821.4 KiB | [citus_17-14.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-14.0.0-4PIGSTY.el9.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 795.4 KiB | [citus_17-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.2 KiB | [citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 768.0 KiB | [citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 750.0 KiB | [citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 746.8 KiB | [citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 858.1 KiB | [citus_17-14.0.0-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_17-14.0.0-4PIGSTY.el10.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 832.8 KiB | [citus_17-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 827.5 KiB | [citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 804.7 KiB | [citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.1 KiB | [citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.4 KiB | [citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 829.4 KiB | [citus_17-14.0.0-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_17-14.0.0-4PIGSTY.el10.aarch64.rpm) |
| `citus_17` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 803.8 KiB | [citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 798.8 KiB | [citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 776.8 KiB | [citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 758.8 KiB | [citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 759.3 KiB | [citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-citus` | `14.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-citus` | `14.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-citus` | `14.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-17-citus` | `14.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-17-citus` | `14.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-17-citus` | `14.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-17-citus` | `14.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-17-citus` | `14.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-17-citus_14.0.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_14.0.0-4PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_16` | `14.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 959.5 KiB | [citus_16-14.0.0-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-14.0.0-4PIGSTY.el8.x86_64.rpm) |
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
| `citus_16` | `14.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 921.6 KiB | [citus_16-14.0.0-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-14.0.0-4PIGSTY.el8.aarch64.rpm) |
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
| `citus_16` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 841.8 KiB | [citus_16-14.0.0-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-14.0.0-4PIGSTY.el9.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.4 KiB | [citus_16-14.0.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/citus_16-14.0.0-1PGDG.rhel9.7.x86_64.rpm) |
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
| `citus_16` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 814.9 KiB | [citus_16-14.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-14.0.0-4PIGSTY.el9.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 788.9 KiB | [citus_16-14.0.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/citus_16-14.0.0-1PGDG.rhel9.7.aarch64.rpm) |
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
| `citus_16` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 851.1 KiB | [citus_16-14.0.0-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_16-14.0.0-4PIGSTY.el10.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 826.1 KiB | [citus_16-14.0.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-14.0.0-1PGDG.rhel10.1.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 820.4 KiB | [citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 797.3 KiB | [citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.1 KiB | [citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.4 KiB | [citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 823.1 KiB | [citus_16-14.0.0-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_16-14.0.0-4PIGSTY.el10.aarch64.rpm) |
| `citus_16` | `14.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 797.4 KiB | [citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-14.0.0-1PGDG.rhel10.1.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 792.1 KiB | [citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 770.0 KiB | [citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.3 KiB | [citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.7 KiB | [citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-citus` | `14.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-citus` | `14.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-citus` | `14.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~trixie_amd64.deb) |
| `postgresql-16-citus` | `14.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~trixie_arm64.deb) |
| `postgresql-16-citus` | `14.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~jammy_amd64.deb) |
| `postgresql-16-citus` | `14.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~jammy_arm64.deb) |
| `postgresql-16-citus` | `14.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~noble_amd64.deb) |
| `postgresql-16-citus` | `14.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-citus_14.0.0-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_14.0.0-4PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_15` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 999.3 KiB | [citus_15-13.2.0-8PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_15-13.2.0-8PIGSTY.el8.x86_64.rpm) |
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
| `citus_15` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 957.8 KiB | [citus_15-13.2.0-8PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_15-13.2.0-8PIGSTY.el8.aarch64.rpm) |
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
| `citus_15` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 893.5 KiB | [citus_15-13.2.0-8PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_15-13.2.0-8PIGSTY.el9.x86_64.rpm) |
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
| `citus_15` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 856.2 KiB | [citus_15-13.2.0-8PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_15-13.2.0-8PIGSTY.el9.aarch64.rpm) |
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
| `citus_15` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 895.7 KiB | [citus_15-13.2.0-8PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_15-13.2.0-8PIGSTY.el10.x86_64.rpm) |
| `citus_15` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 858.2 KiB | [citus_15-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 837.8 KiB | [citus_15-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 818.2 KiB | [citus_15-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 818.2 KiB | [citus_15-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/citus_15-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_15` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 862.0 KiB | [citus_15-13.2.0-8PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_15-13.2.0-8PIGSTY.el10.aarch64.rpm) |
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
| `citus_14` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 807.0 KiB | [citus_14-13.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_14-13.0.0-1PIGSTY.el8.x86_64.rpm) |
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
| `citus_14` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 758.4 KiB | [citus_14-13.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_14-13.0.0-1PIGSTY.el8.aarch64.rpm) |
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
| `citus_14` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 801.9 KiB | [citus_14-13.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_14-13.0.0-1PIGSTY.el9.x86_64.rpm) |
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
| `citus_14` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 769.9 KiB | [citus_14-13.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_14-13.0.0-1PIGSTY.el9.aarch64.rpm) |
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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.0.0.tar.gz" >}}
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

> [citus: Distributed PostgreSQL for multi-tenant and real-time analytics](https://github.com/citusdata/citus)

Citus transforms PostgreSQL into a distributed database, enabling horizontal scaling by sharding tables across multiple nodes. It supports multi-tenant SaaS workloads, real-time analytics, and high-throughput transactional use cases while preserving the full PostgreSQL feature set.

**Key Documentation:**

- [What is Citus?](https://docs.citusdata.com/en/stable/get_started/what_is_citus.html)
- [Distributed Tables](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-distributed-table)
- [Reference Tables](https://docs.citusdata.com/en/stable/develop/api_udf.html#create-reference-table)
- [Sharding Models](https://docs.citusdata.com/en/stable/sharding/data_modeling.html)
- [Colocation](https://docs.citusdata.com/en/stable/sharding/data_modeling.html#colocation)
- [Distributed Queries](https://docs.citusdata.com/en/stable/develop/reference_sql.html)
- [Cluster Management](https://docs.citusdata.com/en/stable/admin_guide/cluster_management.html)
- [Configuration Reference](https://docs.citusdata.com/en/stable/develop/api_guc.html)
- [Columnar Storage](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)

### Getting Started

Enable the extension and add worker nodes:

```sql
CREATE EXTENSION citus;

-- Add worker nodes to the cluster
SELECT citus_set_coordinator_host('coord-host', 5432);
SELECT * FROM citus_add_node('worker-1', 5432);
SELECT * FROM citus_add_node('worker-2', 5432);

-- Verify the cluster
SELECT * FROM citus_get_active_worker_nodes();
```

### Creating Distributed Tables

Distribute a table by a chosen distribution column (shard key). Rows with the same key value are colocated on the same shard.

```sql
CREATE TABLE events (
    tenant_id   INT,
    event_id    BIGSERIAL,
    event_time  TIMESTAMPTZ DEFAULT now(),
    event_type  TEXT,
    payload     JSONB,
    PRIMARY KEY (tenant_id, event_id)
);

-- Hash-distribute by tenant_id (default: 32 shards)
SELECT create_distributed_table('events', 'tenant_id');
```

You can control the shard count:

```sql
SELECT create_distributed_table('events', 'tenant_id', shard_count := 64);
```

### Reference Tables

Small lookup tables that need to be joined with distributed tables should be created as reference tables. They are replicated in full to every node.

```sql
CREATE TABLE countries (
    code CHAR(2) PRIMARY KEY,
    name TEXT NOT NULL
);

SELECT create_reference_table('countries');
```

Reference tables can be joined with any distributed table without restrictions.

### Colocation

Tables distributed on the same column type and shard count are automatically colocated, meaning rows with matching distribution keys are stored on the same node. This enables efficient local joins.

```sql
CREATE TABLE tenants (
    id   INT PRIMARY KEY,
    name TEXT
);
SELECT create_distributed_table('tenants', 'id');

CREATE TABLE orders (
    tenant_id  INT REFERENCES tenants(id),
    order_id   BIGSERIAL,
    amount     NUMERIC,
    PRIMARY KEY (tenant_id, order_id)
);
SELECT create_distributed_table('orders', 'tenant_id');

-- This join is pushed down to each node (no cross-shard traffic)
SELECT t.name, sum(o.amount)
FROM tenants t JOIN orders o ON t.id = o.tenant_id
GROUP BY t.name;
```

You can also explicitly specify colocation groups:

```sql
SELECT create_distributed_table('orders', 'tenant_id',
    colocate_with := 'tenants');
```

### Distributed Queries

Citus pushes queries down to individual shards when possible. Queries that filter on the distribution column are routed to a single shard:

```sql
-- Single-shard query (fast, touches one node)
SELECT * FROM events WHERE tenant_id = 42;
```

Cross-shard queries are parallelized across all workers:

```sql
-- Parallel aggregation across all shards
SELECT event_type, count(*), avg(payload->>'duration')::numeric
FROM events
WHERE event_time > now() - INTERVAL '1 hour'
GROUP BY event_type
ORDER BY count DESC
LIMIT 10;
```

### Distributed Joins

Joins between colocated tables on the distribution column are executed locally on each shard:

```sql
-- Colocated join: efficient, no data movement
SELECT e.*, o.amount
FROM events e JOIN orders o
    ON e.tenant_id = o.tenant_id
WHERE e.tenant_id = 42;
```

Joins with reference tables work from any distributed table:

```sql
SELECT e.*, c.name AS country_name
FROM events e JOIN countries c ON e.payload->>'country' = c.code;
```

### Node Management

```sql
-- Add a new node
SELECT * FROM citus_add_node('worker-3', 5432);

-- Remove a node (moves shards to other nodes first)
SELECT * FROM citus_drain_node('worker-1', 5432);
SELECT * FROM citus_remove_node('worker-1', 5432);

-- Disable a node temporarily
SELECT * FROM citus_disable_node('worker-2', 5432);
SELECT * FROM citus_activate_node('worker-2', 5432);

-- View current cluster state
SELECT * FROM citus_get_active_worker_nodes();
```

### Shard Rebalancing

After adding or removing nodes, rebalance shards to distribute data evenly:

```sql
-- Rebalance all distributed tables
SELECT citus_rebalance_start();

-- Monitor rebalance progress
SELECT * FROM citus_rebalance_status();

-- Rebalance a specific table
SELECT rebalance_table_shards('events');
```

### Shard Management

```sql
-- View shard placements
SELECT * FROM citus_shards;

-- View shard sizes
SELECT table_name, shard_count, citus_total_relation_size(table_name::text)
FROM citus_tables;

-- Move a specific shard to another node
SELECT citus_move_shard_placement(shard_id, 'source-host', 5432, 'dest-host', 5432);
```

### Configuration Parameters

Key GUC parameters for tuning Citus:

```sql
-- Number of parallel connections per node for multi-shard queries
SET citus.max_adaptive_executor_pool_size = 4;

-- Shard replication factor (default 1; set to 2 for HA without streaming replication)
SET citus.shard_replication_factor = 1;

-- Control executor behavior
SET citus.multi_shard_modify_mode = 'parallel';   -- or 'sequential'
SET citus.enable_repartition_joins = on;           -- enable repartitioned joins

-- Task assignment policy
SET citus.task_assignment_policy = 'round-robin';  -- or 'greedy', 'first-replica'

-- Log distributed queries
SET citus.log_multi_join_order = on;
SET citus.explain_all_tasks = on;
```

### Example: Multi-Tenant SaaS Schema

A typical multi-tenant schema distributes all tenant-scoped tables by `tenant_id`:

```sql
CREATE TABLE tenants (
    tenant_id   INT PRIMARY KEY,
    name        TEXT,
    plan        TEXT DEFAULT 'free',
    created_at  TIMESTAMPTZ DEFAULT now()
);
SELECT create_distributed_table('tenants', 'tenant_id');

CREATE TABLE users (
    tenant_id   INT,
    user_id     BIGSERIAL,
    email       TEXT,
    PRIMARY KEY (tenant_id, user_id)
);
SELECT create_distributed_table('users', 'tenant_id');

CREATE TABLE projects (
    tenant_id   INT,
    project_id  BIGSERIAL,
    name        TEXT,
    owner_id    BIGINT,
    PRIMARY KEY (tenant_id, project_id)
);
SELECT create_distributed_table('projects', 'tenant_id');

-- Shared lookup: replicated to every node
CREATE TABLE plans (
    name TEXT PRIMARY KEY,
    max_users INT,
    max_projects INT
);
SELECT create_reference_table('plans');

-- All joins scoped to a tenant are colocated and fast
SELECT u.email, p.name AS project
FROM users u
JOIN projects p ON u.tenant_id = p.tenant_id AND u.user_id = p.owner_id
WHERE u.tenant_id = 7;
```

### Example: Real-Time Analytics

Distributed aggregation for dashboards and analytics:

```sql
CREATE TABLE page_views (
    site_id      INT,
    url          TEXT,
    view_time    TIMESTAMPTZ DEFAULT now(),
    user_agent   TEXT,
    country      CHAR(2)
);
SELECT create_distributed_table('page_views', 'site_id');

-- Real-time rollup: parallelized across shards
SELECT
    date_trunc('minute', view_time) AS minute,
    count(*)                        AS views,
    count(DISTINCT country)         AS countries
FROM page_views
WHERE site_id = 1
  AND view_time > now() - INTERVAL '1 hour'
GROUP BY minute
ORDER BY minute DESC;

-- Top pages across all sites (cross-shard parallel query)
SELECT url, count(*) AS views
FROM page_views
WHERE view_time > now() - INTERVAL '24 hours'
GROUP BY url
ORDER BY views DESC
LIMIT 20;
```
