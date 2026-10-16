---
title: "semver"
linkTitle: "semver"
description: "Semantic version data type"
weight: 3510
categories: ["TYPE"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_semver**](https://github.com/theory/pg-semver) : Semantic version data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3510** | {{< badge content="semver" link="https://github.com/theory/pg-semver" >}} | {{< ext "semver" "pg_semver" >}} | `0.42.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_text_semver" >}} {{< ext "debversion" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.42.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_semver` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.42.0` | {{< bg "18" "semver_18" "green" >}} {{< bg "17" "semver_17" "green" >}} {{< bg "16" "semver_16" "green" >}} {{< bg "15" "semver_15" "green" >}} {{< bg "14" "semver_14" "green" >}} | `semver_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.42.0` | {{< bg "18" "postgresql-18-semver" "green" >}} {{< bg "17" "postgresql-17-semver" "green" >}} {{< bg "16" "postgresql-16-semver" "green" >}} {{< bg "15" "postgresql-15-semver" "green" >}} {{< bg "14" "postgresql-14-semver" "green" >}} | `postgresql-$v-semver` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 7" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 5" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 7" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 7" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 6" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.42.0" "semver_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.42.0" "semver_14 : AVAIL 5" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.42.0" "postgresql-18-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-17-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-16-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-15-semver : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.42.0" "postgresql-14-semver : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_18` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.0 KiB | [semver_18-0.42.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_18-0.42.0-1PGSTY.el8.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [semver_18-0.42.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.42.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 KiB | [semver_18-0.41.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.41.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.9 KiB | [semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 44.0 KiB | [semver_18-0.42.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_18-0.42.0-1PGSTY.el8.aarch64.rpm) |
| `semver_18` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [semver_18-0.42.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.42.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.0 KiB | [semver_18-0.41.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.41.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.5 KiB | [semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm) |
| `semver_18` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.4 KiB | [semver_18-0.42.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_18-0.42.0-1PGSTY.el9.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.4 KiB | [semver_18-0.42.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.42.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.2 KiB | [semver_18-0.41.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.41.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.2 KiB | [semver_18-0.41.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.41.0-1PGDG.rhel9.7.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.41.0-1PGDG.rhel9.6.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.9 KiB | [semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 44.4 KiB | [semver_18-0.42.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_18-0.42.0-1PGSTY.el9.aarch64.rpm) |
| `semver_18` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.0 KiB | [semver_18-0.42.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.42.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_18-0.41.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.41.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_18-0.41.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.41.0-1PGDG.rhel9.7.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.9 KiB | [semver_18-0.41.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.41.0-1PGDG.rhel9.6.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.5 KiB | [semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm) |
| `semver_18` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 45.5 KiB | [semver_18-0.42.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_18-0.42.0-1PGSTY.el10.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [semver_18-0.42.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.42.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.41.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.41.0-1PGDG.rhel10.1.x86_64.rpm) |
| `semver_18` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [semver_18-0.41.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.41.0-1PGDG.rhel10.0.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_18-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_18` | `0.42.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 45.1 KiB | [semver_18-0.42.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_18-0.42.0-1PGSTY.el10.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/semver_18-0.41.0-1PGDG.rhel10.2.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/semver_18-0.41.0-1PGDG.rhel10.1.aarch64.rpm) |
| `semver_18` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_18-0.41.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/semver_18-0.41.0-1PGDG.rhel10.0.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_18-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/semver_18-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.7 KiB | [postgresql-18-semver_0.42.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.9 KiB | [postgresql-18-semver_0.42.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-semver` | `0.41.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.1 KiB | [postgresql-18-semver_0.41.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.3 KiB | [postgresql-18-semver_0.42.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.6 KiB | [postgresql-18-semver_0.42.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-semver` | `0.41.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.5 KiB | [postgresql-18-semver_0.41.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.6 KiB | [postgresql-18-semver_0.42.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.8 KiB | [postgresql-18-semver_0.42.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-semver` | `0.41.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.6 KiB | [postgresql-18-semver_0.41.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.6 KiB | [postgresql-18-semver_0.42.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.8 KiB | [postgresql-18-semver_0.42.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-semver` | `0.41.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.8 KiB | [postgresql-18-semver_0.41.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 42.6 KiB | [postgresql-18-semver_0.42.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.6 KiB | [postgresql-18-semver_0.42.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.7 KiB | [postgresql-18-semver_0.41.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.0 KiB | [postgresql-18-semver_0.42.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.2 KiB | [postgresql-18-semver_0.42.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.2 KiB | [postgresql-18-semver_0.41.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.5 KiB | [postgresql-18-semver_0.42.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 36.0 KiB | [postgresql-18-semver_0.42.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.7 KiB | [postgresql-18-semver_0.41.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.3 KiB | [postgresql-18-semver_0.42.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 36.1 KiB | [postgresql-18-semver_0.42.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.6 KiB | [postgresql-18-semver_0.41.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 41.4 KiB | [postgresql-18-semver_0.42.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.6 KiB | [postgresql-18-semver_0.42.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 38.5 KiB | [postgresql-18-semver_0.41.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 41.3 KiB | [postgresql-18-semver_0.42.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.8 KiB | [postgresql-18-semver_0.42.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-18-semver_0.42.0-1PGSTY~resolute_arm64.deb) |
| `postgresql-18-semver` | `0.41.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 38.4 KiB | [postgresql-18-semver_0.41.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.41.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_17` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.8 KiB | [semver_17-0.42.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_17-0.42.0-1PGSTY.el8.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [semver_17-0.42.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.42.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 KiB | [semver_17-0.41.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.41.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_17` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [semver_17-0.42.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_17-0.42.0-1PGSTY.el8.aarch64.rpm) |
| `semver_17` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [semver_17-0.42.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.42.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.0 KiB | [semver_17-0.41.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.41.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_17` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.8 KiB | [semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_17` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.2 KiB | [semver_17-0.42.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_17-0.42.0-1PGSTY.el9.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.4 KiB | [semver_17-0.42.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.42.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.1 KiB | [semver_17-0.41.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.41.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.1 KiB | [semver_17-0.41.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.41.0-1PGDG.rhel9.7.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.2 KiB | [semver_17-0.41.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.41.0-1PGDG.rhel9.6.x86_64.rpm) |
| `semver_17` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 44.2 KiB | [semver_17-0.42.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_17-0.42.0-1PGSTY.el9.aarch64.rpm) |
| `semver_17` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.0 KiB | [semver_17-0.42.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.42.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_17-0.41.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.41.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_17-0.41.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.41.0-1PGDG.rhel9.7.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.9 KiB | [semver_17-0.41.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.41.0-1PGDG.rhel9.6.aarch64.rpm) |
| `semver_17` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_17` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 45.4 KiB | [semver_17-0.42.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_17-0.42.0-1PGSTY.el10.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [semver_17-0.42.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.42.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_17-0.41.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.41.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_17-0.41.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.41.0-1PGDG.rhel10.1.x86_64.rpm) |
| `semver_17` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [semver_17-0.41.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.41.0-1PGDG.rhel10.0.x86_64.rpm) |
| `semver_17` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_17-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_17` | `0.42.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.9 KiB | [semver_17-0.42.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_17-0.42.0-1PGSTY.el10.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_17-0.41.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/semver_17-0.41.0-1PGDG.rhel10.2.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_17-0.41.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/semver_17-0.41.0-1PGDG.rhel10.1.aarch64.rpm) |
| `semver_17` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_17-0.41.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/semver_17-0.41.0-1PGDG.rhel10.0.aarch64.rpm) |
| `semver_17` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_17-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/semver_17-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.5 KiB | [postgresql-17-semver_0.42.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.7 KiB | [postgresql-17-semver_0.42.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-semver` | `0.41.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.9 KiB | [postgresql-17-semver_0.41.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.2 KiB | [postgresql-17-semver_0.42.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.4 KiB | [postgresql-17-semver_0.42.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-semver` | `0.41.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.3 KiB | [postgresql-17-semver_0.41.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.4 KiB | [postgresql-17-semver_0.42.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-17-semver_0.42.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-semver` | `0.41.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.5 KiB | [postgresql-17-semver_0.41.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.5 KiB | [postgresql-17-semver_0.42.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.7 KiB | [postgresql-17-semver_0.42.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-semver` | `0.41.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.6 KiB | [postgresql-17-semver_0.41.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-17-semver_0.42.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.1 KiB | [postgresql-17-semver_0.42.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-17-semver_0.41.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.5 KiB | [postgresql-17-semver_0.42.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.7 KiB | [postgresql-17-semver_0.42.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.7 KiB | [postgresql-17-semver_0.41.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.4 KiB | [postgresql-17-semver_0.42.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.9 KiB | [postgresql-17-semver_0.42.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.5 KiB | [postgresql-17-semver_0.41.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.2 KiB | [postgresql-17-semver_0.42.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.9 KiB | [postgresql-17-semver_0.42.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.4 KiB | [postgresql-17-semver_0.41.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 41.3 KiB | [postgresql-17-semver_0.42.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.5 KiB | [postgresql-17-semver_0.42.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 38.4 KiB | [postgresql-17-semver_0.41.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 41.2 KiB | [postgresql-17-semver_0.42.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.6 KiB | [postgresql-17-semver_0.42.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-17-semver_0.42.0-1PGSTY~resolute_arm64.deb) |
| `postgresql-17-semver` | `0.41.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 38.2 KiB | [postgresql-17-semver_0.41.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.41.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_16` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.8 KiB | [semver_16-0.42.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_16-0.42.0-1PGSTY.el8.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [semver_16-0.42.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.42.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 KiB | [semver_16-0.41.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.41.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_16` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [semver_16-0.42.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_16-0.42.0-1PGSTY.el8.aarch64.rpm) |
| `semver_16` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [semver_16-0.42.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.42.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.0 KiB | [semver_16-0.41.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.41.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_16` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.8 KiB | [semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_16` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.2 KiB | [semver_16-0.42.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_16-0.42.0-1PGSTY.el9.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.4 KiB | [semver_16-0.42.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.42.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.1 KiB | [semver_16-0.41.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.41.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.1 KiB | [semver_16-0.41.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.41.0-1PGDG.rhel9.7.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.2 KiB | [semver_16-0.41.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.41.0-1PGDG.rhel9.6.x86_64.rpm) |
| `semver_16` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 44.2 KiB | [semver_16-0.42.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_16-0.42.0-1PGSTY.el9.aarch64.rpm) |
| `semver_16` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.0 KiB | [semver_16-0.42.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.42.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_16-0.41.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.41.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [semver_16-0.41.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.41.0-1PGDG.rhel9.7.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.9 KiB | [semver_16-0.41.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.41.0-1PGDG.rhel9.6.aarch64.rpm) |
| `semver_16` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_16` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 45.4 KiB | [semver_16-0.42.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_16-0.42.0-1PGSTY.el10.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.5 KiB | [semver_16-0.42.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.42.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_16-0.41.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.41.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [semver_16-0.41.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.41.0-1PGDG.rhel10.1.x86_64.rpm) |
| `semver_16` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [semver_16-0.41.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.41.0-1PGDG.rhel10.0.x86_64.rpm) |
| `semver_16` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_16-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_16` | `0.42.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 44.9 KiB | [semver_16-0.42.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_16-0.42.0-1PGSTY.el10.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_16-0.41.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/semver_16-0.41.0-1PGDG.rhel10.2.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_16-0.41.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/semver_16-0.41.0-1PGDG.rhel10.1.aarch64.rpm) |
| `semver_16` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.3 KiB | [semver_16-0.41.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/semver_16-0.41.0-1PGDG.rhel10.0.aarch64.rpm) |
| `semver_16` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_16-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/semver_16-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.5 KiB | [postgresql-16-semver_0.42.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.7 KiB | [postgresql-16-semver_0.42.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-semver` | `0.41.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.9 KiB | [postgresql-16-semver_0.41.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.2 KiB | [postgresql-16-semver_0.42.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.5 KiB | [postgresql-16-semver_0.42.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-semver` | `0.41.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.3 KiB | [postgresql-16-semver_0.41.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.4 KiB | [postgresql-16-semver_0.42.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-16-semver_0.42.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-semver` | `0.41.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.5 KiB | [postgresql-16-semver_0.41.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.4 KiB | [postgresql-16-semver_0.42.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.6 KiB | [postgresql-16-semver_0.42.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-semver` | `0.41.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.6 KiB | [postgresql-16-semver_0.41.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-16-semver_0.42.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.1 KiB | [postgresql-16-semver_0.42.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.2 KiB | [postgresql-16-semver_0.41.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.5 KiB | [postgresql-16-semver_0.42.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.7 KiB | [postgresql-16-semver_0.42.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.8 KiB | [postgresql-16-semver_0.41.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.3 KiB | [postgresql-16-semver_0.42.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.9 KiB | [postgresql-16-semver_0.42.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.5 KiB | [postgresql-16-semver_0.41.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.2 KiB | [postgresql-16-semver_0.42.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.7 KiB | [postgresql-16-semver_0.42.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.4 KiB | [postgresql-16-semver_0.41.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 41.3 KiB | [postgresql-16-semver_0.42.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.5 KiB | [postgresql-16-semver_0.42.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 38.4 KiB | [postgresql-16-semver_0.41.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 41.2 KiB | [postgresql-16-semver_0.42.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.6 KiB | [postgresql-16-semver_0.42.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-16-semver_0.42.0-1PGSTY~resolute_arm64.deb) |
| `postgresql-16-semver` | `0.41.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 38.2 KiB | [postgresql-16-semver_0.41.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.41.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_15` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.0 KiB | [semver_15-0.42.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_15-0.42.0-1PGSTY.el8.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.8 KiB | [semver_15-0.42.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.42.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.5 KiB | [semver_15-0.41.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.41.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_15` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_15` | `0.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [semver_15-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_15` | `0.31.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [semver_15-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 44.0 KiB | [semver_15-0.42.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_15-0.42.0-1PGSTY.el8.aarch64.rpm) |
| `semver_15` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.1 KiB | [semver_15-0.42.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.42.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [semver_15-0.41.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.41.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_15` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_15` | `0.32.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [semver_15-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_15` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.8 KiB | [semver_15-0.42.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_15-0.42.0-1PGSTY.el9.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.8 KiB | [semver_15-0.42.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.42.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [semver_15-0.41.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.41.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [semver_15-0.41.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.41.0-1PGDG.rhel9.7.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.6 KiB | [semver_15-0.41.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.41.0-1PGDG.rhel9.6.x86_64.rpm) |
| `semver_15` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_15` | `0.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [semver_15-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_15` | `0.31.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [semver_15-0.31.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.31.2-1.rhel9.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 44.5 KiB | [semver_15-0.42.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_15-0.42.0-1PGSTY.el9.aarch64.rpm) |
| `semver_15` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.2 KiB | [semver_15-0.42.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.42.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.1 KiB | [semver_15-0.41.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.41.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.1 KiB | [semver_15-0.41.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.41.0-1PGDG.rhel9.7.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.2 KiB | [semver_15-0.41.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.41.0-1PGDG.rhel9.6.aarch64.rpm) |
| `semver_15` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_15` | `0.32.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [semver_15-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_15` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.0 KiB | [semver_15-0.42.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_15-0.42.0-1PGSTY.el10.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [semver_15-0.42.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.42.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_15-0.41.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.41.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_15-0.41.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.41.0-1PGDG.rhel10.1.x86_64.rpm) |
| `semver_15` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.1 KiB | [semver_15-0.41.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.41.0-1PGDG.rhel10.0.x86_64.rpm) |
| `semver_15` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_15-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_15` | `0.42.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 45.3 KiB | [semver_15-0.42.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_15-0.42.0-1PGSTY.el10.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_15-0.41.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/semver_15-0.41.0-1PGDG.rhel10.2.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_15-0.41.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/semver_15-0.41.0-1PGDG.rhel10.1.aarch64.rpm) |
| `semver_15` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [semver_15-0.41.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/semver_15-0.41.0-1PGDG.rhel10.0.aarch64.rpm) |
| `semver_15` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_15-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/semver_15-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.5 KiB | [postgresql-15-semver_0.42.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.8 KiB | [postgresql-15-semver_0.42.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-semver` | `0.41.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.9 KiB | [postgresql-15-semver_0.41.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.4 KiB | [postgresql-15-semver_0.42.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 41.2 KiB | [postgresql-15-semver_0.42.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-semver` | `0.41.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.3 KiB | [postgresql-15-semver_0.41.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.5 KiB | [postgresql-15-semver_0.42.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.7 KiB | [postgresql-15-semver_0.42.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-semver` | `0.41.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.5 KiB | [postgresql-15-semver_0.41.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.6 KiB | [postgresql-15-semver_0.42.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.8 KiB | [postgresql-15-semver_0.42.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-semver` | `0.41.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.6 KiB | [postgresql-15-semver_0.41.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.3 KiB | [postgresql-15-semver_0.42.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-15-semver_0.42.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.4 KiB | [postgresql-15-semver_0.41.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-15-semver_0.42.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.9 KiB | [postgresql-15-semver_0.42.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.9 KiB | [postgresql-15-semver_0.41.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.5 KiB | [postgresql-15-semver_0.42.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 36.1 KiB | [postgresql-15-semver_0.42.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.9 KiB | [postgresql-15-semver_0.41.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.5 KiB | [postgresql-15-semver_0.42.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 36.3 KiB | [postgresql-15-semver_0.42.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.4 KiB | [postgresql-15-semver_0.41.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 41.7 KiB | [postgresql-15-semver_0.42.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.8 KiB | [postgresql-15-semver_0.42.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 38.6 KiB | [postgresql-15-semver_0.41.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 41.6 KiB | [postgresql-15-semver_0.42.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.9 KiB | [postgresql-15-semver_0.42.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-15-semver_0.42.0-1PGSTY~resolute_arm64.deb) |
| `postgresql-15-semver` | `0.41.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 38.3 KiB | [postgresql-15-semver_0.41.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.41.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_14` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.1 KiB | [semver_14-0.42.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_14-0.42.0-1PGSTY.el8.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.8 KiB | [semver_14-0.42.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.42.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.5 KiB | [semver_14-0.41.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.41.0-1PGDG.rhel8.10.x86_64.rpm) |
| `semver_14` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_14` | `0.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [semver_14-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_14` | `0.31.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.4 KiB | [semver_14-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_14` | `0.31.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.8 KiB | [semver_14-0.31.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.1-2.rhel8.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 44.0 KiB | [semver_14-0.42.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_14-0.42.0-1PGSTY.el8.aarch64.rpm) |
| `semver_14` | `0.42.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.1 KiB | [semver_14-0.42.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.42.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [semver_14-0.41.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.41.0-1PGDG.rhel8.10.aarch64.rpm) |
| `semver_14` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_14` | `0.32.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [semver_14-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_14` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.8 KiB | [semver_14-0.42.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_14-0.42.0-1PGSTY.el9.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.8 KiB | [semver_14-0.42.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.42.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [semver_14-0.41.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.41.0-1PGDG.rhel9.8.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [semver_14-0.41.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.41.0-1PGDG.rhel9.7.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.6 KiB | [semver_14-0.41.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.41.0-1PGDG.rhel9.6.x86_64.rpm) |
| `semver_14` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_14` | `0.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [semver_14-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 44.6 KiB | [semver_14-0.42.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_14-0.42.0-1PGSTY.el9.aarch64.rpm) |
| `semver_14` | `0.42.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.2 KiB | [semver_14-0.42.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.42.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.1 KiB | [semver_14-0.41.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.41.0-1PGDG.rhel9.8.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.1 KiB | [semver_14-0.41.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.41.0-1PGDG.rhel9.7.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.2 KiB | [semver_14-0.41.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.41.0-1PGDG.rhel9.6.aarch64.rpm) |
| `semver_14` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.0 KiB | [semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_14` | `0.32.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [semver_14-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_14` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.0 KiB | [semver_14-0.42.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_14-0.42.0-1PGSTY.el10.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [semver_14-0.42.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.42.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [semver_14-0.41.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.41.0-1PGDG.rhel10.2.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.7 KiB | [semver_14-0.41.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.41.0-1PGDG.rhel10.1.x86_64.rpm) |
| `semver_14` | `0.41.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 28.1 KiB | [semver_14-0.41.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.41.0-1PGDG.rhel10.0.x86_64.rpm) |
| `semver_14` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_14-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_14` | `0.42.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 45.3 KiB | [semver_14-0.42.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_14-0.42.0-1PGSTY.el10.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [semver_14-0.41.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/semver_14-0.41.0-1PGDG.rhel10.2.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [semver_14-0.41.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/semver_14-0.41.0-1PGDG.rhel10.1.aarch64.rpm) |
| `semver_14` | `0.41.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.6 KiB | [semver_14-0.41.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/semver_14-0.41.0-1PGDG.rhel10.0.aarch64.rpm) |
| `semver_14` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_14-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/semver_14-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.6 KiB | [postgresql-14-semver_0.42.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.8 KiB | [postgresql-14-semver_0.42.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-semver` | `0.41.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.9 KiB | [postgresql-14-semver_0.41.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.4 KiB | [postgresql-14-semver_0.42.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.6 KiB | [postgresql-14-semver_0.42.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-semver` | `0.41.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.3 KiB | [postgresql-14-semver_0.41.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.6 KiB | [postgresql-14-semver_0.42.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.8 KiB | [postgresql-14-semver_0.42.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-semver` | `0.41.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.6 KiB | [postgresql-14-semver_0.41.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.7 KiB | [postgresql-14-semver_0.42.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.9 KiB | [postgresql-14-semver_0.42.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-semver` | `0.41.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.6 KiB | [postgresql-14-semver_0.41.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.2 KiB | [postgresql-14-semver_0.42.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.3 KiB | [postgresql-14-semver_0.42.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.4 KiB | [postgresql-14-semver_0.41.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-14-semver_0.42.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.8 KiB | [postgresql-14-semver_0.42.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.9 KiB | [postgresql-14-semver_0.41.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.6 KiB | [postgresql-14-semver_0.42.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 36.1 KiB | [postgresql-14-semver_0.42.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.9 KiB | [postgresql-14-semver_0.41.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.6 KiB | [postgresql-14-semver_0.42.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 36.3 KiB | [postgresql-14-semver_0.42.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.4 KiB | [postgresql-14-semver_0.41.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 41.6 KiB | [postgresql-14-semver_0.42.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.8 KiB | [postgresql-14-semver_0.42.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 38.5 KiB | [postgresql-14-semver_0.41.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 41.5 KiB | [postgresql-14-semver_0.42.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.42.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.9 KiB | [postgresql-14-semver_0.42.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresql-semver/postgresql-14-semver_0.42.0-1PGSTY~resolute_arm64.deb) |
| `postgresql-14-semver` | `0.41.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 38.3 KiB | [postgresql-14-semver_0.41.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.41.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pg-semver" title="Repository" icon="github" subtitle="github.com/theory/pg-semver" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-semver-0.42.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_semver;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_semver;		# install via package name, for the active PG version
pig install semver;		# install by extension name, for the current active PG version

pig install semver -v 18;   # install for PG 18
pig install semver -v 17;   # install for PG 17
pig install semver -v 16;   # install for PG 16
pig install semver -v 15;   # install for PG 15
pig install semver -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION semver;
```

## Usage

Sources:

- [doc/semver.md](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/doc/semver.md)
- [Changes](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/Changes)
- [semver.control](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/semver.control)
- [sql/semver--0.41.0--0.42.0.sql](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/sql/semver--0.41.0--0.42.0.sql)

`semver` provides a Semantic Versioning 2.0.0 data type, comparison operators and the `semverrange` range type. Version 0.42.0 changes hashing: after upgrading, rebuild every hash index on semver values.

### Core Workflow

```sql
CREATE EXTENSION semver;
SELECT '1.2.0'::semver > '1.2.0-b1'::semver;
SELECT '1.2.0'::semver = '1.2.0+build1'::semver;
SELECT to_semver('1.0'), is_semver('1.2.0');
SELECT get_semver_major('4.2.0'::semver);
SELECT '1.0.5'::semver <@ '[1.0.0,2.0.0)'::semverrange;
```

### Operational Boundaries

`semver(text)` and casts from `text` require valid semantic versions; `to_semver(text)` performs more permissive conversion, and `is_semver(text)` checks validity. `get_semver_major`, `get_semver_minor` and `get_semver_patch` return integer components; `get_semver_prerelease` returns the prerelease text. Casts also accept `numeric`, `real`, `double precision`, `integer`, `bigint` and `smallint`. Version 0.42.0 rejects Infinity and NaN instead of producing malformed versions.

The comparison operators `=`, `<>`, `<`, `<=`, `>` and `>=`, plus `MIN(semver)` and `MAX(semver)`, use semantic ordering. A prerelease sorts before its final release; build metadata does not affect precedence. Btree and hash indexes are available. `semverrange` supports standard range operations.

Install this native, untrusted extension with administrator privileges. After installing 0.42.0, run `ALTER EXTENSION semver UPDATE` and use REINDEX on every existing hash index containing semver keys. The hash change is not repaired by updating the extension’s SQL definitions alone. The release also fixes ordering of alphanumeric prerelease identifiers beginning with a digit and validation of strings with multiple plus signs. When upgrading from old releases, inspect the upstream compatibility notes for previously accepted invalid values.
