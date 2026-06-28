---
title: "plpgsql_check"
linkTitle: "plpgsql_check"
description: "extended check for plpgsql functions"
weight: 3060
categories: ["LANG"]
width: full
---

[**plpgsql_check**](https://github.com/okbob/plpgsql_check) : extended check for plpgsql functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3060** | {{< badge content="plpgsql_check" link="https://github.com/okbob/plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | `2.9.1` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pldbgapi" >}} {{< ext "plprofiler" >}} {{< ext "pg_hint_plan" >}} {{< ext "pgtap" >}} {{< ext "auto_explain" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.9.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plpgsql_check` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.9.1` | {{< bg "18" "plpgsql_check_18" "green" >}} {{< bg "17" "plpgsql_check_17" "green" >}} {{< bg "16" "plpgsql_check_16" "green" >}} {{< bg "15" "plpgsql_check_15" "green" >}} {{< bg "14" "plpgsql_check_14" "green" >}} | `plpgsql_check_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.9.1` | {{< bg "18" "postgresql-18-plpgsql-check" "green" >}} {{< bg "17" "postgresql-17-plpgsql-check" "green" >}} {{< bg "16" "postgresql-16-plpgsql-check" "green" >}} {{< bg "15" "postgresql-15-plpgsql-check" "green" >}} {{< bg "14" "postgresql-14-plpgsql-check" "green" >}} | `postgresql-$v-plpgsql-check` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 27" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 35" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 45" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 27" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 34" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 34" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 20" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 34" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 42" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 49" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 41" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.9.1" "plpgsql_check_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.9.1" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.1" "postgresql-14-plpgsql-check : AVAIL 4" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_18` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.5 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_18-2.9.1-1PIGSTY.el8.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.7 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.8 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.0 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 108.9 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_18-2.9.1-1PIGSTY.el8.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.2 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.5 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.4 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.4 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 111.1 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_18-2.9.1-1PIGSTY.el9.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.6 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.6 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.1 KiB | [plpgsql_check_18-2.8.11-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.11-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.4 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.7 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.6 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.7 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_18-2.9.1-1PIGSTY.el9.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_18-2.8.11-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.11-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 113.2 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_18-2.9.1-1PIGSTY.el10.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.5 KiB | [plpgsql_check_18-2.8.11-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.11-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 107.9 KiB | [plpgsql_check_18-2.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_18-2.9.1-1PIGSTY.el10.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.1 KiB | [plpgsql_check_18-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 303.6 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 301.6 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 299.1 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.6 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 293.0 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 290.6 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 288.8 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.5 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 304.0 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 302.0 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 299.8 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 293.1 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 293.7 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.8 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 289.2 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.5 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 313.5 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 328.1 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 308.4 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 301.7 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 302.7 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 321.6 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 298.0 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 291.1 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 303.0 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 314.4 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 298.7 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.9 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 291.9 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 308.1 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 287.5 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.7 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 300.4 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 312.3 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 295.6 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u26.x86_64](/os/u26.x86_64) | pgdg | 290.6 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 288.8 KiB | [postgresql-18-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 305.5 KiB | [postgresql-18-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.9.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 284.0 KiB | [postgresql-18-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u26.aarch64](/os/u26.aarch64) | pgdg | 278.6 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_17` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.3 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_17-2.9.1-1PIGSTY.el8.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.3 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.1 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.0 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.8 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.3 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.4 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 108.9 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_17-2.9.1-1PIGSTY.el8.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.3 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.4 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.4 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 111.2 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_17-2.9.1-1PIGSTY.el9.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.4 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.6 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.5 KiB | [plpgsql_check_17-2.8.11-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.11-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.5 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.2 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.5-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.5 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_17-2.9.1-1PIGSTY.el9.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 110.6 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [plpgsql_check_17-2.8.11-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.11-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 113.1 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_17-2.9.1-1PIGSTY.el10.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.1 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.4 KiB | [plpgsql_check_17-2.8.11-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.11-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 107.8 KiB | [plpgsql_check_17-2.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_17-2.9.1-1PIGSTY.el10.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_17-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 303.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 301.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 299.0 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 292.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 290.6 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 288.1 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 303.8 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 301.9 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 299.3 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 292.9 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 293.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 288.8 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 385.9 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 402.2 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 380.6 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 372.8 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 374.8 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 395.4 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 369.5 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 361.6 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 303.0 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 314.5 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 298.5 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.7 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 291.8 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 308.1 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 287.2 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.4 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 300.4 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 312.2 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 295.4 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u26.x86_64](/os/u26.x86_64) | pgdg | 290.4 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 288.1 KiB | [postgresql-17-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 305.4 KiB | [postgresql-17-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.9.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 284.1 KiB | [postgresql-17-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u26.aarch64](/os/u26.aarch64) | pgdg | 278.3 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_16` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.3 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_16-2.9.1-1PIGSTY.el8.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.9 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.4 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.1 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.0 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.7 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.2 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.9 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.4 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.0 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.3 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.0 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.8 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.7 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.8 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.6 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.0 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.6 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.3 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.2 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 108.8 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_16-2.9.1-1PIGSTY.el8.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.3 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.0 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.5 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.3 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.1 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.9 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.7 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.4 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.8 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.3 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.7 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.5 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.4 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.1 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.0 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 111.0 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_16-2.9.1-1PIGSTY.el9.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.4 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.4 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.5 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.0 KiB | [plpgsql_check_16-2.8.11-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.11-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.5 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.5-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.0 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.7 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.8 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.0 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.4 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.2 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.0 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.2 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.0 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.5 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.9 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.9 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.8 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.5 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.4 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 106.6 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_16-2.9.1-1PIGSTY.el9.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 110.7 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_16-2.8.11-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.11-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 113.2 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_16-2.9.1-1PIGSTY.el10.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.2 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.5 KiB | [plpgsql_check_16-2.8.11-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.11-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 107.8 KiB | [plpgsql_check_16-2.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_16-2.9.1-1PIGSTY.el10.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.9 KiB | [plpgsql_check_16-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 303.5 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 301.3 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 298.8 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 292.7 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 290.5 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 288.2 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.2 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 303.9 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 301.7 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 299.3 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 292.8 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 293.7 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 291.9 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 289.1 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 380.1 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 396.3 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 374.7 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 367.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 369.1 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 389.7 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 363.8 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 356.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 302.8 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 314.4 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 298.4 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.8 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 291.9 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 308.0 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 287.5 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.5 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 300.1 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 312.2 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 295.5 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u26.x86_64](/os/u26.x86_64) | pgdg | 290.4 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 288.7 KiB | [postgresql-16-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 305.4 KiB | [postgresql-16-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.9.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 283.6 KiB | [postgresql-16-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u26.aarch64](/os/u26.aarch64) | pgdg | 278.7 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_15` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.4 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_15-2.9.1-1PIGSTY.el8.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.2 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.3 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.1 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.8 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.2 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.1 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.9 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.5 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.2 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.4 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.2 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.9 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 98.6 KiB | [plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.4 KiB | [plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.3 KiB | [plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 96.7 KiB | [plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.6 KiB | [plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.8 KiB | [plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 111.4 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_15-2.9.1-1PIGSTY.el8.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.9 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.8 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.6 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.1 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.0 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.4 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.7 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.6 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.0 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.8 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.7 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.5 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.8 KiB | [plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.7 KiB | [plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.6 KiB | [plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.1 KiB | [plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.0 KiB | [plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.7 KiB | [plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.7 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_15-2.9.1-1PIGSTY.el9.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.6 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.5 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.9 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.2 KiB | [plpgsql_check_15-2.8.11-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.11-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.1 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.7 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.5 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.5 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.5-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.3 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.5 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.5 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.2 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.3 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.0 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.3 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.2 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.8 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.2 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.2 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.9 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.7 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.9 KiB | [plpgsql_check_15-2.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.8 KiB | [plpgsql_check_15-2.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.8 KiB | [plpgsql_check_15-2.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.4 KiB | [plpgsql_check_15-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.0-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.7 KiB | [plpgsql_check_15-2.2.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.6-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.4 KiB | [plpgsql_check_15-2.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.4 KiB | [plpgsql_check_15-2.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.5 KiB | [plpgsql_check_15-2.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 94.9 KiB | [plpgsql_check_15-2.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 111.0 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_15-2.9.1-1PIGSTY.el9.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 110.9 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_15-2.8.11-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.11-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.8 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_15-2.9.1-1PIGSTY.el10.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.1 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.6 KiB | [plpgsql_check_15-2.8.11-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.11-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.2 KiB | [plpgsql_check_15-2.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_15-2.9.1-1PIGSTY.el10.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.9 KiB | [plpgsql_check_15-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 306.8 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 304.8 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 302.4 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 295.3 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 295.5 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 293.4 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 290.9 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 283.7 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 307.6 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 305.5 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 303.2 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 295.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 296.5 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 294.3 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 291.8 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 284.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 384.1 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 400.5 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 379.1 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 370.3 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 373.1 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 394.1 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 368.0 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 359.5 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 306.7 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 318.4 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 302.4 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 295.1 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 296.0 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 312.5 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 291.3 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 283.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 303.9 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 315.7 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 299.3 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u26.x86_64](/os/u26.x86_64) | pgdg | 293.5 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 292.4 KiB | [postgresql-15-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 309.8 KiB | [postgresql-15-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.9.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 287.8 KiB | [postgresql-15-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u26.aarch64](/os/u26.aarch64) | pgdg | 281.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_14` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.4 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_14-2.9.1-1PIGSTY.el8.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.4 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.2 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.0 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.8 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.3 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.9 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.5 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.9 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.4 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.2 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.2 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.0 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.8 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.2 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.0 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.6 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 98.4 KiB | [plpgsql_check_14-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.3 KiB | [plpgsql_check_14-2.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.2 KiB | [plpgsql_check_14-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 96.7 KiB | [plpgsql_check_14-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.0-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.6 KiB | [plpgsql_check_14-2.2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.6-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.4 KiB | [plpgsql_check_14-2.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_14-2.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.2 KiB | [plpgsql_check_14-2.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.7 KiB | [plpgsql_check_14-2.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.2 KiB | [plpgsql_check_14-2.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 90.4 KiB | [plpgsql_check_14-2.1.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.10-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 89.5 KiB | [plpgsql_check_14-2.1.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.8-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 89.4 KiB | [plpgsql_check_14-2.1.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.7-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 89.3 KiB | [plpgsql_check_14-2.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.7 KiB | [plpgsql_check_14-2.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.4 KiB | [plpgsql_check_14-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.6 KiB | [plpgsql_check_14-2.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.0.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.0 KiB | [plpgsql_check_14-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.0.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `1.17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 83.0 KiB | [plpgsql_check_14-1.17.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-1.17.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 111.4 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_14-2.9.1-1PIGSTY.el8.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.0 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.9 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.8 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.5 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.1 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.9 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.5 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.8 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.4 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.1 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.6 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.4 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.0 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.4 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.7 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.6 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.5 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.3 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.6 KiB | [plpgsql_check_14-2.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.6 KiB | [plpgsql_check_14-2.3.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.2-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.5 KiB | [plpgsql_check_14-2.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.1-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.0 KiB | [plpgsql_check_14-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.0-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.9 KiB | [plpgsql_check_14-2.2.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.6-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.8 KiB | [plpgsql_check_14-2.2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.5-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.7 KiB | [plpgsql_check_14-2.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.6 KiB | [plpgsql_check_14-2.2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.3-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.3 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_14-2.9.1-1PIGSTY.el9.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.5 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.9.1-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.9.1-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.6 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.9.1-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_14-2.8.11-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.11-1PGDG.rhel9.8.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.5 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.10-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.2 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.8-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.0 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.5-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.7 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.8 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.9 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.8 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.3 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.1 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 105.9 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.7 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.8 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.8 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.8 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 100.6 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.8 KiB | [plpgsql_check_14-2.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.8 KiB | [plpgsql_check_14-2.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.7 KiB | [plpgsql_check_14-2.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.4 KiB | [plpgsql_check_14-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.0-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.5 KiB | [plpgsql_check_14-2.2.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.6-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.4 KiB | [plpgsql_check_14-2.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_14-2.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 95.2 KiB | [plpgsql_check_14-2.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 94.7 KiB | [plpgsql_check_14-2.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 94.3 KiB | [plpgsql_check_14-2.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 90.4 KiB | [plpgsql_check_14-2.1.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.10-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 89.8 KiB | [plpgsql_check_14-2.1.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.8-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 89.6 KiB | [plpgsql_check_14-2.1.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.7-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 89.6 KiB | [plpgsql_check_14-2.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.8 KiB | [plpgsql_check_14-2.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.7 KiB | [plpgsql_check_14-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 110.9 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_14-2.9.1-1PIGSTY.el9.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 110.8 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.9.1-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 110.9 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.9.1-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.9.1-1PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_14-2.8.11-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.11-1PGDG.rhel9.8.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.10-1PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.6 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.8-1PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.5-1PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.7 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.5 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.7 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.6 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.2 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.8 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.8 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 101.1 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 100.9 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.0 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.0 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.3 KiB | [plpgsql_check_14-2.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.5 KiB | [plpgsql_check_14-2.3.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.2-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [plpgsql_check_14-2.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.1-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.3 KiB | [plpgsql_check_14-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.0-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.3 KiB | [plpgsql_check_14-2.2.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.6-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.1 KiB | [plpgsql_check_14-2.2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.5-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.0 KiB | [plpgsql_check_14-2.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 90.9 KiB | [plpgsql_check_14-2.2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.3-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.8 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_14-2.9.1-1PIGSTY.el10.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.6 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.9.1-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.6 KiB | [plpgsql_check_14-2.8.11-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.11-1PGDG.rhel10.2.x86_64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.2 KiB | [plpgsql_check_14-2.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_14-2.9.1-1PIGSTY.el10.aarch64.rpm) |
| `plpgsql_check_14` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_14-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_14-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 306.6 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 304.5 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 302.7 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 295.5 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 295.1 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 293.0 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 290.9 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 283.5 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 307.4 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 305.2 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 302.9 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 296.0 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 296.4 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 294.0 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 291.8 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 284.8 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 368.0 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 385.0 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 362.5 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 353.9 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 357.1 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 378.0 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 351.8 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 343.0 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 306.8 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 318.3 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 302.5 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 295.2 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 295.5 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 312.5 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 291.1 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 283.4 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 303.7 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 315.5 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 299.4 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u26.x86_64](/os/u26.x86_64) | pgdg | 293.4 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 292.3 KiB | [postgresql-14-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 309.5 KiB | [postgresql-14-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.9.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 287.7 KiB | [postgresql-14-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.9.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u26.aarch64](/os/u26.aarch64) | pgdg | 281.6 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/plpgsql_check" title="Repository" icon="github" subtitle="github.com/okbob/plpgsql_check" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plpgsql_check-2.9.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plpgsql_check;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plpgsql_check;		# install via package name, for the active PG version

pig install plpgsql_check -v 18;   # install for PG 18
pig install plpgsql_check -v 17;   # install for PG 17
pig install plpgsql_check -v 16;   # install for PG 16
pig install plpgsql_check -v 15;   # install for PG 15
pig install plpgsql_check -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'plpgsql_check';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plpgsql_check CASCADE; -- requires plpgsql
```

## Usage

Sources: [official README](https://github.com/okbob/plpgsql_check), [v2.9.1 release notes](https://github.com/okbob/plpgsql_check/releases/tag/v2.9.1), [local package metadata](../db/extension.csv).

`plpgsql_check` is a PL/pgSQL checker, linter, profiler, tracer, and coverage tool. It detects many errors during development instead of waiting for runtime failures.

```sql
CREATE EXTENSION plpgsql_check;
```

### Check A Function

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

The table-returning variant is useful for structured reports:

```sql
SELECT *
FROM plpgsql_check_function_tb('my_function()');
```

### Output Formats

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Check Trigger Functions

```sql
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

SELECT * FROM plpgsql_check_function(
  'my_trigger_func()',
  'my_table',
  newtable := 'newtab',
  oldtable := 'oldtab'
);
```

### Warning Categories

```sql
SELECT * FROM plpgsql_check_function(
  'fx()',
  extra_warnings := true,
  performance_warnings := true,
  security_warnings := true,
  compatibility_warnings := true
);
```

- `extra_warnings` covers issues such as missing returns, dead code, and unused arguments.
- `performance_warnings` covers performance-related checks.
- `security_warnings` includes checks such as SQL injection risk.
- `compatibility_warnings` reports obsolete or version-sensitive PL/pgSQL patterns.

### Batch Check All Functions

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql'
  AND p.prorettype <> 'pg_catalog.trigger'::pg_catalog.regtype;
```

### Passive Mode

Passive mode checks functions when they start. It is intended for development or preproduction because it adds overhead.

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';
```

Common settings:

```text
plpgsql_check.mode = [ disabled | by_function | fresh_start | every_start ]
plpgsql_check.fatal_errors = [ yes | no ]
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### Profiler

```sql
SELECT plpgsql_check_profiler(true);

SELECT my_function();

SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

For shared profiler statistics, preload `plpgsql` before `plpgsql_check`:

```text
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

Without shared preload, profiler data is limited to the active session.

### Tracer

Tracing emits notices for function and statement entry/exit and can expose variable values. It is disabled by default and should be enabled only with superuser-controlled settings.

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true);
SET plpgsql_check.tracer_verbosity = terse;
```

### Dependency Tracking

```sql
SELECT *
FROM plpgsql_show_dependency_tb('my_function(int)');
```

### Coverage Metrics

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma Directives

Use pragma calls inside functions to tell `plpgsql_check` about dynamic SQL, temporary tables, inferred record types, or local check settings:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

### Caveats

- Pigsty packages `plpgsql_check` 2.9.1 for PostgreSQL 14-18 as RPMs; DEB packages come from PGDG.
- The extension requires `plpgsql`. Preloading is optional for active checking, but needed for shared profiler storage and reliable early tracer/profiler initialization.
- v2.9.1 fixes a possible crash when a traced inline block fails; no new user-facing SQL API was documented for this patch release.
- The tracer can expose function arguments or variable values. Use it carefully around security-definer functions or sensitive data.
