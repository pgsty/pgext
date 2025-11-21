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
| **2400** | {{< badge content="citus" link="https://github.com/citusdata/citus" >}} | {{< ext "citus" >}} | `13.2.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_partman" >}} {{< ext "plproxy" >}} {{< ext "columnar" >}} {{< ext "pg_fkpart" >}} {{< ext "timescaledb" >}} {{< ext "pg_duckdb" >}} {{< ext "tablefunc" >}} {{< ext "hll" >}} |
|    **Siblings**   | {{< ext "citus_columnar" >}} |

> [!Note] conflict with hydra


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `13.2.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `13.2.0` | {{< bg "18" "citus_18*" "red" >}} {{< bg "17" "citus_17*" "green" >}} {{< bg "16" "citus_16*" "green" >}} {{< bg "15" "citus_15*" "green" >}} {{< bg "14" "citus_14*" "green" >}} {{< bg "13" "citus_13*" "red" >}} | `citus_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `13.2.0` | {{< bg "18" "postgresql-18-citus" "red" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "green" >}} {{< bg "14" "postgresql-14-citus" "green" >}} {{< bg "13" "postgresql-13-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 29" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 29" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 21" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 26" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 16" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "citus_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "citus_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_17` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 960.5 KiB | [citus_17-13.2.0-8PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_17-13.2.0-8PIGSTY.el8.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 850.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.2.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 827.1 KiB | [citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.1.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.4-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.9 KiB | [citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.3-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 805.7 KiB | [citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.2-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 803.6 KiB | [citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/citus_17-13.0.0-1PGDG.rhel8.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 922.8 KiB | [citus_17-13.2.0-8PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_17-13.2.0-8PIGSTY.el8.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 802.6 KiB | [citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.2.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 781.5 KiB | [citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.1.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.4 KiB | [citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.4-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.5 KiB | [citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.3-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 761.2 KiB | [citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.2-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 758.9 KiB | [citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/citus_17-13.0.0-1PGDG.rhel8.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 842.1 KiB | [citus_17-13.2.0-8PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_17-13.2.0-8PIGSTY.el9.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 815.8 KiB | [citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.2.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 793.2 KiB | [citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.1.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.1 KiB | [citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.4-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.6 KiB | [citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.3-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 774.4 KiB | [citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.2-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 772.0 KiB | [citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/citus_17-13.0.0-1PGDG.rhel9.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 815.4 KiB | [citus_17-13.2.0-8PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_17-13.2.0-8PIGSTY.el9.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.2 KiB | [citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.2.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 768.0 KiB | [citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.1.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.4-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 750.0 KiB | [citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.3-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 749.8 KiB | [citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.2-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 746.8 KiB | [citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/citus_17-13.0.0-1PGDG.rhel9.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 853.5 KiB | [citus_17-13.2.0-8PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_17-13.2.0-8PIGSTY.el10.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 827.5 KiB | [citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 804.7 KiB | [citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.1 KiB | [citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 786.4 KiB | [citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/citus_17-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_17` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 823.9 KiB | [citus_17-13.2.0-8PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_17-13.2.0-8PIGSTY.el10.aarch64.rpm) |
| `citus_17` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 798.8 KiB | [citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 776.8 KiB | [citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 758.8 KiB | [citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_17` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 759.3 KiB | [citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/citus_17-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-citus` | `13.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-citus` | `13.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-citus` | `13.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~trixie_amd64.deb) |
| `postgresql-17-citus` | `13.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~trixie_arm64.deb) |
| `postgresql-17-citus` | `13.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.4 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~jammy_amd64.deb) |
| `postgresql-17-citus` | `13.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.3 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~jammy_arm64.deb) |
| `postgresql-17-citus` | `13.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~noble_amd64.deb) |
| `postgresql-17-citus` | `13.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-17-citus_13.2.0-8PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-17-citus_13.2.0-8PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_16` | `13.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 953.4 KiB | [citus_16-13.2.0-8PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/citus_16-13.2.0-8PIGSTY.el8.x86_64.rpm) |
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
| `citus_16` | `13.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 915.0 KiB | [citus_16-13.2.0-8PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/citus_16-13.2.0-8PIGSTY.el8.aarch64.rpm) |
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
| `citus_16` | `13.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 836.0 KiB | [citus_16-13.2.0-8PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/citus_16-13.2.0-8PIGSTY.el9.x86_64.rpm) |
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
| `citus_16` | `13.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 808.4 KiB | [citus_16-13.2.0-8PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/citus_16-13.2.0-8PIGSTY.el9.aarch64.rpm) |
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
| `citus_16` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 846.1 KiB | [citus_16-13.2.0-8PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/citus_16-13.2.0-8PIGSTY.el10.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 820.4 KiB | [citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.2.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 797.3 KiB | [citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.1.0-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.1 KiB | [citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.4-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 779.4 KiB | [citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/citus_16-13.0.3-1PGDG.rhel10.x86_64.rpm) |
| `citus_16` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 817.1 KiB | [citus_16-13.2.0-8PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/citus_16-13.2.0-8PIGSTY.el10.aarch64.rpm) |
| `citus_16` | `13.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 792.1 KiB | [citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.2.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 770.0 KiB | [citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.1.0-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.3 KiB | [citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.4-1PGDG.rhel10.aarch64.rpm) |
| `citus_16` | `13.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 752.7 KiB | [citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/citus_16-13.0.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-citus` | `13.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-citus` | `13.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-citus` | `13.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~trixie_amd64.deb) |
| `postgresql-16-citus` | `13.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.6 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~trixie_arm64.deb) |
| `postgresql-16-citus` | `13.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~jammy_amd64.deb) |
| `postgresql-16-citus` | `13.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.2 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~jammy_arm64.deb) |
| `postgresql-16-citus` | `13.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~noble_amd64.deb) |
| `postgresql-16-citus` | `13.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-citus_13.2.0-8PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/citus/postgresql-16-citus_13.2.0-8PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

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
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `citus_13` | `9.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.9 MiB | [citus_13-9.5.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.4-1.rhel8.x86_64.rpm) |
| `citus_13` | `9.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.9 MiB | [citus_13-9.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `9.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.9 MiB | [citus_13-9.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `9.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.9 MiB | [citus_13-9.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-9.5.0-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 787.4 KiB | [citus_13-11.3.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.3.0-2.rhel8.x86_64.rpm) |
| `citus_13` | `11.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 768.4 KiB | [citus_13-11.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.2.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 767.1 KiB | [citus_13-11.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.2.0-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 756.4 KiB | [citus_13-11.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.5-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 756.2 KiB | [citus_13-11.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.4-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 755.8 KiB | [citus_13-11.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.3-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 755.4 KiB | [citus_13-11.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 753.7 KiB | [citus_13-11.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.1.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 692.9 KiB | [citus_13-11.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.6-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 692.3 KiB | [citus_13-11.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.5-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 691.4 KiB | [citus_13-11.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.4-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 691.3 KiB | [citus_13-11.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.3-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 690.5 KiB | [citus_13-11.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-11.0.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 611.3 KiB | [citus_13-10.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.5-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 611.3 KiB | [citus_13-10.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.4-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 611.1 KiB | [citus_13-10.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.3-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 608.3 KiB | [citus_13-10.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 607.7 KiB | [citus_13-10.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 607.2 KiB | [citus_13-10.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.2.0-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 582.2 KiB | [citus_13-10.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 581.0 KiB | [citus_13-10.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 580.1 KiB | [citus_13-10.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.1.0-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 556.1 KiB | [citus_13-10.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.3-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 555.3 KiB | [citus_13-10.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.2-1.rhel8.x86_64.rpm) |
| `citus_13` | `10.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 2.2 MiB | [citus_13-10.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/citus_13-10.0.1-1.rhel8.x86_64.rpm) |
| `citus_13` | `11.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 747.9 KiB | [citus_13-11.3.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.3.0-2.rhel8.aarch64.rpm) |
| `citus_13` | `11.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 728.8 KiB | [citus_13-11.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.2.1-1.rhel8.aarch64.rpm) |
| `citus_13` | `11.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 727.8 KiB | [citus_13-11.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.2.0-1.rhel8.aarch64.rpm) |
| `citus_13` | `11.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 716.8 KiB | [citus_13-11.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.5-1.rhel8.aarch64.rpm) |
| `citus_13` | `11.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 716.5 KiB | [citus_13-11.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.4-1.rhel8.aarch64.rpm) |
| `citus_13` | `11.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 716.2 KiB | [citus_13-11.1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/citus_13-11.1.3-1.rhel8.aarch64.rpm) |
| `citus_13` | `11.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 788.2 KiB | [citus_13-11.3.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.3.0-2.rhel9.x86_64.rpm) |
| `citus_13` | `11.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 768.0 KiB | [citus_13-11.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.2.1-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 767.5 KiB | [citus_13-11.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.2.0-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 757.2 KiB | [citus_13-11.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.5-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 754.6 KiB | [citus_13-11.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.4-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 754.5 KiB | [citus_13-11.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.3-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 753.8 KiB | [citus_13-11.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.2-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 752.1 KiB | [citus_13-11.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.1.1-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 691.1 KiB | [citus_13-11.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.6-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 690.2 KiB | [citus_13-11.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.5-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 690.0 KiB | [citus_13-11.0.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.4-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 690.0 KiB | [citus_13-11.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.3-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 689.2 KiB | [citus_13-11.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-11.0.2-1.rhel9.x86_64.rpm) |
| `citus_13` | `10.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 612.5 KiB | [citus_13-10.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.5-1.rhel9.x86_64.rpm) |
| `citus_13` | `10.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 613.9 KiB | [citus_13-10.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.4-1.rhel9.x86_64.rpm) |
| `citus_13` | `10.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 613.9 KiB | [citus_13-10.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/citus_13-10.2.3-1.rhel9.x86_64.rpm) |
| `citus_13` | `11.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 753.8 KiB | [citus_13-11.3.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.3.0-2.rhel9.aarch64.rpm) |
| `citus_13` | `11.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 735.4 KiB | [citus_13-11.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.2.1-1.rhel9.aarch64.rpm) |
| `citus_13` | `11.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 734.0 KiB | [citus_13-11.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.2.0-1.rhel9.aarch64.rpm) |
| `citus_13` | `11.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 723.5 KiB | [citus_13-11.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.5-1.rhel9.aarch64.rpm) |
| `citus_13` | `11.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 722.5 KiB | [citus_13-11.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.4-1.rhel9.aarch64.rpm) |
| `citus_13` | `11.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 722.2 KiB | [citus_13-11.1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/citus_13-11.1.3-1.rhel9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-13.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install citus;		# install via package name, for the active PG version

pig install citus -v 17;   # install for PG 17
pig install citus -v 16;   # install for PG 16
pig install citus -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'citus';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citus;
```
