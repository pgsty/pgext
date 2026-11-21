---
title: "pgsql_tweaks"
linkTitle: "pgsql_tweaks"
description: "Some functions and views for daily usage"
weight: 4210
categories: ["UTIL"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgsql_tweaks**](https://codeberg.org/pgsql_tweaks/pgsql_tweaks) : Some functions and views for daily usage


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4210** | {{< badge content="pgsql_tweaks" link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" >}} | `1.1.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgsql_tweaks` |
|   **See Also**    | {{< ext "pgcalendar" >}} {{< ext "pg_math" >}} {{< ext "pg_extra_time" >}} {{< ext "pg_rrule" >}} |

> [!Note] SQL-only.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pgsql_tweaks` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "19" "pgsql_tweaks_19" "green" >}} {{< bg "18" "pgsql_tweaks_18" "green" >}} {{< bg "17" "pgsql_tweaks_17" "green" >}} {{< bg "16" "pgsql_tweaks_16" "green" >}} {{< bg "15" "pgsql_tweaks_15" "green" >}} | `pgsql_tweaks_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "19" "postgresql-19-pgsql-tweaks" "green" >}} {{< bg "18" "postgresql-18-pgsql-tweaks" "green" >}} {{< bg "17" "postgresql-17-pgsql-tweaks" "green" >}} {{< bg "16" "postgresql-16-pgsql-tweaks" "green" >}} {{< bg "15" "postgresql-15-pgsql-tweaks" "green" >}} | `postgresql-$v-pgsql-tweaks` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pgsql_tweaks_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 14" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 14" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pgsql_tweaks_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 14" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 14" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_18 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 1.1.0" "pgsql_tweaks_15 : AVAIL 13" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-pgsql-tweaks : MISS 0" "red" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_19` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.0 KiB | [pgsql_tweaks_19-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pgsql_tweaks_19-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_19` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pgsql_tweaks_19-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pgsql_tweaks_19-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_19` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.2 KiB | [pgsql_tweaks_19-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pgsql_tweaks_19-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_19` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_19-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pgsql_tweaks_19-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_tweaks_18-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_18-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [pgsql_tweaks_18-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_tweaks_18-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_18-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_tweaks_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pgsql_tweaks_18-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_tweaks_18-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_tweaks_18-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.7 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_18-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_18-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_tweaks_18-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.5 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_18-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_tweaks_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_18-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_tweaks_18-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.1 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_tweaks_18-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.9 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_18-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.7 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_18-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_tweaks_18-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_18-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_18-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.2 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsql_tweaks_18-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.0 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_18-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_18-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_18-1.1.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.1.0-2PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [pgsql_tweaks_18-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_18-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_18-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pgsql_tweaks_18-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_18-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_tweaks_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_18-1.0.3-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.3-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_18-1.0.3-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.3-1PGDG.rhel10.1.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_18-1.0.3-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.3-1PGDG.rhel10.0.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.9 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.6 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.6 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.6 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 21.0 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-18-pgsql-tweaks` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_17-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [pgsql_tweaks_17-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_17-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_tweaks_17` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pgsql_tweaks_17-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.7 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_17-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_17-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.5 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_17-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_tweaks_17` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.7 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `0.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.0 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.9 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_17-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.7 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_17-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_17-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_17-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_tweaks_17-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.0 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_17-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_17-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_17-1.1.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.1.0-2PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [pgsql_tweaks_17-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_17-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_17-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pgsql_tweaks_17-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_17-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_tweaks_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.3-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.3-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.3-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.3-1PGDG.rhel10.1.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.3-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.3-1PGDG.rhel10.0.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.9 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | `0.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.6 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.6 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.6 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 21.0 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-17-pgsql-tweaks` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_16-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [pgsql_tweaks_16-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_16-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_tweaks_16` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pgsql_tweaks_16-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.7 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_16-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_16-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.5 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_16-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_tweaks_16` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.7 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | `0.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.0 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.9 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_16-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.7 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_16-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_16-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.7 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_16-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_tweaks_16-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.0 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_16-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_16-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_16-1.1.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.1.0-2PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [pgsql_tweaks_16-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_16-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_16-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pgsql_tweaks_16-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_16-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_tweaks_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.3-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.3-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.3-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.3-1PGDG.rhel10.1.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.3-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.3-1PGDG.rhel10.0.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.9 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | `0.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.6 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.6 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.6 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 21.0 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-16-pgsql-tweaks` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_15-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [pgsql_tweaks_15-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_tweaks_15-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_tweaks_15` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pgsql_tweaks_15-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.0 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.4 KiB | [pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.7 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_15-1.1.0-1PGSTY.el8.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [pgsql_tweaks_15-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.5 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_tweaks_15-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_tweaks_15` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.3-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.0.3-1PGDG.rhel8.10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.0 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.7 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | `0.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.1 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.9 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_15-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.7 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_tweaks_15-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.0 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-1.1.0-2PGDG.rhel9.9.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_15-1.1.0-1PGSTY.el9.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_tweaks_15-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_tweaks_15-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.0 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_15-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_tweaks_15-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.1.0-2PGDG.rhel10.3.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_15-1.1.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.1.0-2PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [pgsql_tweaks_15-1.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_15-1.1.0-1PGSTY.el10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.1 KiB | [pgsql_tweaks_15-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pgsql_tweaks_15-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_tweaks_15-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_tweaks_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.3-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.3-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.3-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.3-1PGDG.rhel10.1.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.3-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.3-1PGDG.rhel10.0.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.9 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel10.2.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | `0.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.6 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.6 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.6 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.9 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.0 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~noble_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 21.0 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-15-pgsql-tweaks` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.0 KiB | [postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" title="Repository" icon="link" subtitle="codeberg.org/pgsql_tweaks/pgsql_tweaks" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql_tweaks-1.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgsql_tweaks;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgsql_tweaks;		# install via package name, for the active PG version

pig install pgsql_tweaks -v 19;   # install for PG 19
pig install pgsql_tweaks -v 18;   # install for PG 18
pig install pgsql_tweaks -v 17;   # install for PG 17
pig install pgsql_tweaks -v 16;   # install for PG 16
pig install pgsql_tweaks -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pgsql_tweaks;
```

## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/README.md)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/pgsql_tweaks.control)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/makefile)
- [Official source](https://api.pgxn.org/src/pgsql_tweaks/pgsql_tweaks-1.1.0/sql/out/versions/pgsql_tweaks--1.1.0.sql)

`pgsql_tweaks` 1.1.0 provides DBA-oriented helper functions and views in the `pgsql_tweaks` schema. The upstream test matrix covers PostgreSQL 14 through PostgreSQL 19rc1; Pigsty packages it for PostgreSQL 14-18.

### Enablement

```sql
CREATE EXTENSION pgsql_tweaks;
```

### Data Type Check Functions

```sql
SELECT pgsql_tweaks.is_date('2024-01-15');
SELECT pgsql_tweaks.is_time('10:30:00');
SELECT pgsql_tweaks.is_timestamp('2024-01-15 10:30:00');
SELECT pgsql_tweaks.is_real('3.14');
SELECT pgsql_tweaks.is_double_precision('3.14');
SELECT pgsql_tweaks.is_numeric('3.14');
SELECT pgsql_tweaks.is_bigint('9223372036854775807');
SELECT pgsql_tweaks.is_integer('42');
SELECT pgsql_tweaks.is_smallint('42');
SELECT pgsql_tweaks.is_boolean('true');
SELECT pgsql_tweaks.is_json('{"a":1}');
SELECT pgsql_tweaks.is_jsonb('{"a":1}');
SELECT pgsql_tweaks.is_hex('FF');
```

### System Information Functions

```sql
SELECT pgsql_tweaks.pg_schema_size('public');
SELECT * FROM pgsql_tweaks.role_inheritance();
```

### Encoding Functions

```sql
SELECT pgsql_tweaks.is_encoding('UTF8');
SELECT pgsql_tweaks.is_latin1('abc');
SELECT pgsql_tweaks.return_not_part_of_latin1('abc');
SELECT pgsql_tweaks.replace_latin1('abc', '?');
SELECT pgsql_tweaks.return_not_part_of_encoding('abc', 'UTF8');
SELECT pgsql_tweaks.replace_encoding('abc', 'UTF8', '?');
```

### Aggregate Functions

- `gap_fill`, for filling gaps in time series.
- `array_min`, `array_max`, `array_avg`, and `array_sum`.

### Format And Conversion Functions

```sql
SELECT pgsql_tweaks.date_de(current_date);
SELECT pgsql_tweaks.datetime_de(now());
SELECT pgsql_tweaks.to_unix_timestamp(now());
SELECT pgsql_tweaks.hex2bigint('FF');
```

### Utility Functions

```sql
SELECT pgsql_tweaks.is_empty('');
SELECT pgsql_tweaks.array_trim(ARRAY['a', '', 'b']);
SELECT pgsql_tweaks.get_markdown_doku_by_schema('pgsql_tweaks');
```

### System Information Views

- `pg_db_views`, `pg_foreign_keys`, `pg_functions`, `pg_active_locks`.
- `pg_table_matview_infos`, `pg_object_ownership`, `pg_partitioned_tables_infos`.
- `pg_unused_indexes`, `pg_bloat_info`, `pg_table_bloat`, `pg_missing_indexes`.
- `pg_role_permissions`, `pg_role_infos`.

### Statistic And Monitoring Views

- `statistics_top_ten_query_times`, `top_ten_query_average_time_in_seconds`.
- `statistics_top_ten_time_consuming_queries`, `statistics_top_ten_memory_usage_queries`.
- `statistics_top_ten_called_queries`, `statistics_top_ten_rows_returned_queries`.
- `statistics_top_ten_shared_block_hits_queries`, `statistics_top_ten_wal_records_generated_queries`.
- `statistics_query_activity`.
- `monitoring_wal`, `wal_archiving`, `monitoring_active_locks`, `monitoring_replication`.
- `monitoring_database_conflicts`, `monitoring_blocked_and_blocking_activity`.
- `monitoring_follower_wal_status`, `monitoring_vacuum`.

### Caveats

- Some views/functions depend on additional PostgreSQL supplied modules; upstream lists `pg_stat_statements` and `pgstattuple`.
- This distribution is pure SQL and needs no preloaded shared library. After installing matching files, use `ALTER EXTENSION pgsql_tweaks UPDATE TO '1.1.0'` for an existing installation. Helpers still require the relevant object or statistics privileges.
