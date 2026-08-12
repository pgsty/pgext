---
title: "pg_squeeze"
linkTitle: "pg_squeeze"
description: "A tool to remove unused space from a relation."
weight: 5040
categories: ["ADMIN"]
width: full
---

[**pg_squeeze**](https://github.com/cybertec-postgresql/pg_squeeze) : A tool to remove unused space from a relation.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5040** | {{< badge content="pg_squeeze" link="https://github.com/cybertec-postgresql/pg_squeeze" >}} | {{< ext "pg_squeeze" >}} | `1.9.4` | {{< category "ADMIN" >}} | {{< license "BSD-2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `squeeze` |
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pgstattuple" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_column_tetris" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_squeeze` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.4` | {{< bg "18" "pg_squeeze_18" "green" >}} {{< bg "17" "pg_squeeze_17" "green" >}} {{< bg "16" "pg_squeeze_16" "green" >}} {{< bg "15" "pg_squeeze_15" "green" >}} {{< bg "14" "pg_squeeze_14" "green" >}} | `pg_squeeze_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.4` | {{< bg "18" "postgresql-18-squeeze" "green" >}} {{< bg "17" "postgresql-17-squeeze" "green" >}} {{< bg "16" "postgresql-16-squeeze" "green" >}} {{< bg "15" "postgresql-15-squeeze" "green" >}} {{< bg "14" "postgresql-14-squeeze" "green" >}} | `postgresql-$v-squeeze` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 9" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 12" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 11" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.4" "pg_squeeze_14 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.9.4" "postgresql-18-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-17-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-16-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-15-squeeze : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.9.4" "postgresql-14-squeeze : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_18` | `1.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.6 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_squeeze_18-1.9.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.2 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_squeeze_18-1.9.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.0 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_18-1.9.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.1-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.6 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_18-1.9.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.1-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.6 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.0 KiB | [pg_squeeze_18-1.9.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.1-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.0 KiB | [pg_squeeze_18-1.9.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_18-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_18-1.9.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-squeeze` | `1.9.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.3 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.2 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.2 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.9 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.8 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.7 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.9 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.6 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.4 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.0 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.9 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.8 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 119.6 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 119.0 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 118.9 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 114.2 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 113.9 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 113.9 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.1 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.2 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.1 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.7 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.7 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.4 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.9 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.6 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.7 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.2 KiB | [postgresql-18-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.3 KiB | [postgresql-18-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.3 KiB | [postgresql-18-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_17` | `1.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.7 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.9.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.3 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.9.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.5 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.7.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.4 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.0 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.9.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.1-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.2 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.9 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.7.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.0 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.5 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.0 KiB | [pg_squeeze_17-1.9.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.1-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.7.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_17-1.9.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.1-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_squeeze_17-1.9.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_17-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_17-1.9.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.1 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-squeeze` | `1.9.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.8 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.9 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.4 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.2 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.9 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.0 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 117.1 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.8 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.8 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.5 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.2 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.2 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.2 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.1 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 139.9 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 135.0 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 134.9 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 134.8 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.8 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.4 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.4 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.7 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.8 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.7 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.9 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.9 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.9 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.6 KiB | [postgresql-17-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.4 KiB | [postgresql-17-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.6 KiB | [postgresql-17-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_16` | `1.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.7 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.9.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.2 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.5 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.9.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.0 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.0 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.4 KiB | [pg_squeeze_16-1.9.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.1-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.5 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.9 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.6 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.0 KiB | [pg_squeeze_16-1.9.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.1-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.4 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.4 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_16-1.9.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.1-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.1 KiB | [pg_squeeze_16-1.9.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_16-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.3 KiB | [pg_squeeze_16-1.9.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-squeeze` | `1.9.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.7 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.4 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.4 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.0 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.8 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.8 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.9 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.6 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.6 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.3 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.1 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.0 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.5 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 137.8 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 137.8 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.5 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 132.7 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 132.7 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.4 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.3 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.1 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.7 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.6 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.6 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.7 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.8 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.0 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.3 KiB | [postgresql-16-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.5 KiB | [postgresql-16-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.1 KiB | [postgresql-16-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_15` | `1.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.8 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.9.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.5 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.9 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.8 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.1 KiB | [pg_squeeze_15-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.5.0-1.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.9.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.5 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.5 KiB | [pg_squeeze_15-1.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.5.0-1.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 58.0 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_15-1.9.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.1-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.4 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.3 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_squeeze_15-1.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.5.0-1.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.6 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.2 KiB | [pg_squeeze_15-1.9.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.1-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.5 KiB | [pg_squeeze_15-1.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.5.0-1.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.3 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.2 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_15-1.9.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.1-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.5 KiB | [pg_squeeze_15-1.9.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.2 KiB | [pg_squeeze_15-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_15-1.9.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.5 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-squeeze` | `1.9.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.7 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.5 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.4 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.9 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.7 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.6 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.7 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.5 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.4 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.3 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.9 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.0 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.8 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.7 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.7 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 134.0 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.7 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.7 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.5 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.3 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.2 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.8 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.6 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.6 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.0 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 114.9 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.1 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.2 KiB | [postgresql-15-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.3 KiB | [postgresql-15-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.3 KiB | [postgresql-15-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_14` | `1.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.1 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.9.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.6 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.1 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.0 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.0 KiB | [pg_squeeze_14-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.5.0-1.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.2 KiB | [pg_squeeze_14-1.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.4.1-2.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 56.1 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.9.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.0 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.0 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.3 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.5 KiB | [pg_squeeze_14-1.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.5.0-1.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.8 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_14-1.9.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.1-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.8 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.2 KiB | [pg_squeeze_14-1.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.5.0-1.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [pg_squeeze_14-1.4.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.4.1-2.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.0 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.8 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.2 KiB | [pg_squeeze_14-1.9.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.1-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.8 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.9 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.4 KiB | [pg_squeeze_14-1.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.5.0-1.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.7 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.2 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.2 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 58.6 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_14-1.9.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.1-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.7 KiB | [pg_squeeze_14-1.9.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.4 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.4 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.4 KiB | [pg_squeeze_14-1.9.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.0 KiB | [pg_squeeze_14-1.9.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.1-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.0 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-squeeze` | `1.9.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.7 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.7 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 116.6 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.3 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.2 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 112.0 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.9 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 116.8 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 117.0 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.5 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.2 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 112.4 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.8 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.6 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.5 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.9 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.8 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.7 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.8 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.6 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 116.6 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.8 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.9 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 111.8 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.2 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.0 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 115.1 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.5 KiB | [postgresql-14-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.7 KiB | [postgresql-14-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 110.5 KiB | [postgresql-14-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_squeeze" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_squeeze" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_squeeze;		# install via package name, for the active PG version

pig install pg_squeeze -v 18;   # install for PG 18
pig install pg_squeeze -v 17;   # install for PG 17
pig install pg_squeeze -v 16;   # install for PG 16
pig install pg_squeeze -v 15;   # install for PG 15
pig install pg_squeeze -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_squeeze';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_squeeze;
```

## Usage

Sources:

- [pg_squeeze REL1_9_4 release](https://github.com/cybertec-postgresql/pg_squeeze/releases/tag/REL1_9_4)
- [pg_squeeze REL1_9_4 README](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/README.md)
- [pg_squeeze release notes](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/NEWS)

`pg_squeeze` removes bloat from a table and its indexes while allowing concurrent reads and writes. It copies live tuples to new storage and applies concurrent changes through logical decoding, avoiding the long exclusive lock of `VACUUM FULL`. Use it only after sizing replication slots, disk space, and the table's replica identity.

### Configure and Install

```conf
max_replication_slots = 1  # or add one to the existing requirement
shared_preload_libraries = 'pg_squeeze'
wal_level = logical       # required on PostgreSQL versions before 19
```

Restart PostgreSQL, then create the extension:

```sql
CREATE EXTENSION pg_squeeze;
```

The table must have an identity index. A primary key works with the default replica identity; otherwise select a suitable unique index with `ALTER TABLE ... REPLICA IDENTITY USING INDEX`.

### Run an Ad-Hoc Squeeze

```sql
SELECT squeeze.squeeze_table('public', 'pgbench_accounts');

SELECT squeeze.squeeze_table(
  'public',
  'large_table',
  'large_table_cluster_idx',
  'target_tablespace'
);
```

The function starts background work and is not transactional in the ordinary SQL-function sense. Monitor the operation rather than assuming a surrounding `ROLLBACK` cancels it.

### Schedule Tables and Monitor Work

```sql
INSERT INTO squeeze.tables (tabschema, tabname, schedule)
VALUES ('public', 'events', ('{30}', '{22}', NULL, NULL, '{3,5}'));

SELECT * FROM squeeze.get_active_workers();
SELECT * FROM squeeze.log ORDER BY finished DESC;
SELECT * FROM squeeze.errors;
```

The schedule tuple contains minutes, hours, days of month, months, and days of week. Registration also supports thresholds and placement options such as `free_space_extra`, `min_size`, `vacuum_max_age`, `max_retry`, `clustering_index`, relation/index tablespaces, and `skip_analyze`.

For automatic startup:

```conf
squeeze.worker_autostart = 'my_database'
squeeze.worker_role = 'postgres'
```

### Version 1.9.4 and Operational Caveats

- Version 1.9.4 fixes unsafe quoting in dynamically constructed `ANALYZE`, log, and error statements, including a superuser SQL-injection path. Upgrade earlier 1.9 builds promptly.
- A full-table squeeze needs free disk space of roughly twice the combined size of the target table and its indexes.
- Disruptive DDL, `VACUUM FULL`, `CLUSTER`, or `TRUNCATE` can make an in-progress squeeze abort. Coordinate schema changes and use `max_retry` deliberately.
- Like other online rewrite tools, `pg_squeeze` changes row visibility and has documented MVCC caveats for concurrent sessions that retain old snapshots.
- Configure `pg_squeeze` in `shared_preload_libraries` on the new cluster before `pg_upgrade` or dump/restore of a database containing the extension.
- Current Pigsty packages cover PostgreSQL 14-18. For those versions, keep `wal_level = logical`; upstream's relaxed PostgreSQL 19 rule does not apply to this package matrix yet.
