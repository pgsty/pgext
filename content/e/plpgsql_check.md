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
| **3060** | {{< badge content="plpgsql_check" link="https://github.com/okbob/plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | `2.8.11` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.11` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plpgsql_check` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.10` | {{< bg "18" "plpgsql_check_18" "green" >}} {{< bg "17" "plpgsql_check_17" "green" >}} {{< bg "16" "plpgsql_check_16" "green" >}} {{< bg "15" "plpgsql_check_15" "green" >}} {{< bg "14" "plpgsql_check_14" "green" >}} | `plpgsql_check_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.11` | {{< bg "18" "postgresql-18-plpgsql-check" "green" >}} {{< bg "17" "postgresql-17-plpgsql-check" "green" >}} {{< bg "16" "postgresql-16-plpgsql-check" "green" >}} {{< bg "15" "postgresql-15-plpgsql-check" "green" >}} {{< bg "14" "postgresql-14-plpgsql-check" "green" >}} | `postgresql-$v-plpgsql-check` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 25" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 33" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 43" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 25" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 32" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 32" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 25" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 33" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 40" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 25" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 32" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 32" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.10" "plpgsql_check_14 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.8.11" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.11" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_18` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.7 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.8 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.0 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.5 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.4 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.4 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.4 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.6 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.6 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.7 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.7 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.6 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.7 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.1 KiB | [plpgsql_check_18-2.8.5-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.5-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.2 KiB | [plpgsql_check_18-2.8.4-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.4-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.5 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.2 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.7 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.1 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.6 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.5 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 293.1 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.5 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 301.7 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 291.1 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.9 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.7 KiB | [postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `plpgsql_check_17` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.6 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.8 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.8 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.2 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.1 KiB | [plpgsql_check_17-2.8.5-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.5-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.2 KiB | [plpgsql_check_17-2.8.4-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.4-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.6 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 110.4 KiB | [plpgsql_check_17-2.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.8.1-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.3 KiB | [plpgsql_check_17-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 292.9 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.1 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 372.8 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 361.6 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.7 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.4 KiB | [postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `plpgsql_check_16` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
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
| `plpgsql_check_16` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.7 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.7 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.7 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.2 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.8 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.5 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.4 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.4 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 100.2 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.8 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.7 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.0 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.9 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.1 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.7 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.3 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.1 KiB | [plpgsql_check_16-2.8.5-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.5-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.1 KiB | [plpgsql_check_16-2.8.4-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.4-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.6 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.6 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 110.4 KiB | [plpgsql_check_16-2.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.8.1-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.3 KiB | [plpgsql_check_16-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 292.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 281.2 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 292.8 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 282.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 367.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 356.1 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 291.8 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 280.5 KiB | [postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `plpgsql_check_15` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.5 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.3 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
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
| `plpgsql_check_15` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.6 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.8 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.8 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.9 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.9 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.0 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.9 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 101.6 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 101.3 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.9 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.5 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.6 KiB | [plpgsql_check_15-2.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.5 KiB | [plpgsql_check_15-2.3.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.2-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [plpgsql_check_15-2.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.1-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [plpgsql_check_15-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.0-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.3 KiB | [plpgsql_check_15-2.2.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.6-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.2 KiB | [plpgsql_check_15-2.2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.5-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.1 KiB | [plpgsql_check_15-2.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 91.0 KiB | [plpgsql_check_15-2.2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.3-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.3 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.3 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.2 KiB | [plpgsql_check_15-2.8.5-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.5-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.1 KiB | [plpgsql_check_15-2.8.4-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.4-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.3 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.8 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_15-2.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.8.1-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.4 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_15-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 295.3 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 283.7 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 295.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 284.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 370.3 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 359.5 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 295.1 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 283.9 KiB | [postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `plpgsql_check_14` | `2.8.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.3 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.10-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.8-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.5-1PGDG.rhel9.7.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.7 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.4-1PGDG.rhel9.7.x86_64.rpm) |
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
| `plpgsql_check_14` | `2.8.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.10-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.8-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.5-1PGDG.rhel9.7.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.4-1PGDG.rhel9.7.aarch64.rpm) |
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
| `plpgsql_check_14` | `2.8.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.3 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.10-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.3 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.8-1PGDG.rhel10.1.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.0 KiB | [plpgsql_check_14-2.8.5-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.5-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.5 KiB | [plpgsql_check_14-2.8.4-1PGDGrhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.4-1PGDGrhel10.1.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.8 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.3-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.4 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.2-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 114.5 KiB | [plpgsql_check_14-2.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plpgsql_check_14-2.8.1-1PGDG.rhel10.x86_64.rpm) |
| `plpgsql_check_14` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.7 KiB | [plpgsql_check_14-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.4 KiB | [plpgsql_check_14-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_14-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_14-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_14` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.7 KiB | [plpgsql_check_14-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plpgsql_check_14-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d12.x86_64](/os/d12.x86_64) | pgdg | 295.5 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d12.aarch64](/os/d12.aarch64) | pgdg | 283.5 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d13.x86_64](/os/d13.x86_64) | pgdg | 296.0 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [d13.aarch64](/os/d13.aarch64) | pgdg | 284.8 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u22.x86_64](/os/u22.x86_64) | pgdg | 353.9 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u22.aarch64](/os/u22.aarch64) | pgdg | 343.0 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u24.x86_64](/os/u24.x86_64) | pgdg | 295.2 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | `2.8.11` | [u24.aarch64](/os/u24.aarch64) | pgdg | 283.4 KiB | [postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.11-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/plpgsql_check" title="Repository" icon="github" subtitle="github.com/okbob/plpgsql_check" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plpgsql_check-2.8.11.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
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

> [plpgsql_check: extended check for plpgsql functions](https://github.com/okbob/plpgsql_check)

`plpgsql_check` is a linter and checker for PL/pgSQL functions that detects errors at development time rather than runtime.

```sql
CREATE EXTENSION plpgsql_check;
```

### Check a Function

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

### Output Formats

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Check Trigger Functions

```sql
-- Trigger functions need the associated table
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

-- With transition tables
SELECT * FROM plpgsql_check_function(
  'my_trigger_func()', 'my_table',
  newtable := 'newtab', oldtable := 'oldtab'
);
```

### Warning Categories

```sql
SELECT * FROM plpgsql_check_function('fx()',
  extra_warnings := true,         -- dead code, unused parameters
  performance_warnings := true,   -- index and casting issues
  security_warnings := true,      -- SQL injection checks
  compatibility_warnings := true  -- obsolete patterns
);
```

### Batch Check All Functions

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql' AND p.prorettype <> 2279;
```

### Passive Mode (Check on Execution)

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';  -- check before each execution
```

Or in `postgresql.conf`:

```
shared_preload_libraries = 'plpgsql,plpgsql_check'
plpgsql_check.mode = 'every_start'
```

### Profiler

```sql
-- Enable profiling
SELECT plpgsql_check_profiler(true);

-- Execute functions to collect data
SELECT my_function();

-- View per-line execution times
SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

-- Per-statement profile
SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

-- All function statistics
SELECT * FROM plpgsql_profiler_functions_all();

-- Reset profiling data
SELECT plpgsql_profiler_reset_all();
```

### Dependency Tracking

```sql
SELECT * FROM plpgsql_show_dependency_tb('my_function(int)');
```

### Coverage Metrics

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma Directives

Embed checking options in function comments:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  /* @plpgsql_check_options: anyelementtype = text */
  RETURN $1;
END;
$$ LANGUAGE plpgsql;
```
