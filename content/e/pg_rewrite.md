---
title: "pg_rewrite"
linkTitle: "pg_rewrite"
description: "Tool allows read write to the table during the rewriting"
weight: 5020
categories: ["ADMIN"]
width: full
---

[**pg_rewrite**](https://github.com/cybertec-postgresql/pg_rewrite) : Tool allows read write to the table during the rewriting


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5020** | {{< badge content="pg_rewrite" link="https://github.com/cybertec-postgresql/pg_rewrite" >}} | {{< ext "pg_rewrite" >}} | `2.1.0` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "pgfincore" >}} {{< ext "pg_prewarm" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_rewrite` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "pg_rewrite_18" "green" >}} {{< bg "17" "pg_rewrite_17" "green" >}} {{< bg "16" "pg_rewrite_16" "green" >}} {{< bg "15" "pg_rewrite_15" "green" >}} {{< bg "14" "pg_rewrite_14" "green" >}} | `pg_rewrite_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "postgresql-18-pg-rewrite" "green" >}} {{< bg "17" "postgresql-17-pg-rewrite" "green" >}} {{< bg "16" "postgresql-16-pg-rewrite" "green" >}} {{< bg "15" "postgresql-15-pg-rewrite" "green" >}} {{< bg "14" "postgresql-14-pg-rewrite" "green" >}} | `postgresql-$v-pg-rewrite` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_rewrite_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-15-pg-rewrite : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-14-pg-rewrite : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rewrite_18` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.3 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_rewrite_18-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.6 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_rewrite_18-2.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_18` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_rewrite_18-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.4 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_rewrite_18-2.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.8 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_rewrite_18-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.1 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_rewrite_18-2.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.2 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_rewrite_18-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.4 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_rewrite_18-2.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.3 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_rewrite_18-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.0 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_rewrite_18-2.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.5 KiB | [pg_rewrite_18-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_rewrite_18-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_rewrite_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.1 KiB | [pg_rewrite_18-2.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_rewrite_18-2.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.6 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.8 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 76.3 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.7 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.0 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.0 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.3 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-rewrite` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 71.6 KiB | [postgresql-18-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-18-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rewrite_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.9 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_rewrite_17-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.0 KiB | [pg_rewrite_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rewrite_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.2 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_rewrite_17-2.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.8 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_rewrite_17-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [pg_rewrite_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rewrite_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_rewrite_17-2.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.5 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_rewrite_17-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 37.7 KiB | [pg_rewrite_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rewrite_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.9 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_rewrite_17-2.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_rewrite_17-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 36.1 KiB | [pg_rewrite_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rewrite_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.2 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_rewrite_17-2.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.0 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_rewrite_17-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.7 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_rewrite_17-2.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.5 KiB | [pg_rewrite_17-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_rewrite_17-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_rewrite_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.2 KiB | [pg_rewrite_17-2.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_rewrite_17-2.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.6 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.2 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.5 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.2 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 90.4 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 85.8 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.7 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-rewrite` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.9 KiB | [postgresql-17-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-17-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rewrite_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.9 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_rewrite_16-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.3 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_rewrite_16-2.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_rewrite_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.8 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_rewrite_16-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_rewrite_16-2.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.1 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_rewrite_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.4 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_rewrite_16-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.9 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_rewrite_16-2.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.4 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_rewrite_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_rewrite_16-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_rewrite_16-2.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.9 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_rewrite_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.0 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_rewrite_16-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.7 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_rewrite_16-2.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.3 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_rewrite_16-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.5 KiB | [pg_rewrite_16-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_rewrite_16-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_rewrite_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.2 KiB | [pg_rewrite_16-2.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_rewrite_16-2.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_rewrite_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.9 KiB | [pg_rewrite_16-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_rewrite_16-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.4 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.4 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.0 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.7 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 85.0 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.4 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-rewrite` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.8 KiB | [postgresql-16-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-16-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rewrite_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.0 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_rewrite_15-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.3 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_rewrite_15-2.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_rewrite_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.8 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_rewrite_15-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_rewrite_15-2.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.0 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_rewrite_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.6 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_rewrite_15-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.0 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_rewrite_15-2.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.5 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_rewrite_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_rewrite_15-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.2 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_rewrite_15-2.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.9 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_rewrite_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_15` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.9 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_rewrite_15-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.7 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_rewrite_15-2.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.5 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_rewrite_15-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_15` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.3 KiB | [pg_rewrite_15-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_rewrite_15-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_rewrite_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [pg_rewrite_15-2.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_rewrite_15-2.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_rewrite_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.1 KiB | [pg_rewrite_15-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_rewrite_15-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.4 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 70.8 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.2 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 70.6 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.5 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 84.5 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.2 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-rewrite` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.6 KiB | [postgresql-15-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-15-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rewrite_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.1 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_rewrite_14-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.5 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_rewrite_14-2.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_rewrite_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_rewrite_14` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.8 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_rewrite_14-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.2 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_rewrite_14-2.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.0 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_rewrite_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_rewrite_14` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.7 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_rewrite_14-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.1 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_rewrite_14-2.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.8 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_rewrite_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_rewrite_14` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.1 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_rewrite_14-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.5 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_rewrite_14-2.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.0 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_rewrite_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_rewrite_14` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.2 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_rewrite_14-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.0 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_rewrite_14-2.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.6 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_rewrite_14-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_rewrite_14` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.4 KiB | [pg_rewrite_14-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_rewrite_14-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_rewrite_14` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.2 KiB | [pg_rewrite_14-2.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_rewrite_14-2.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_rewrite_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.3 KiB | [pg_rewrite_14-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_rewrite_14-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.6 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.3 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.2 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 70.9 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.6 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 84.9 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.4 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-rewrite` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 71.1 KiB | [postgresql-14-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rewrite/postgresql-14-pg-rewrite_2.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_rewrite" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_rewrite" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_rewrite;		# install via package name, for the active PG version

pig install pg_rewrite -v 18;   # install for PG 18
pig install pg_rewrite -v 17;   # install for PG 17
pig install pg_rewrite -v 16;   # install for PG 16
pig install pg_rewrite -v 15;   # install for PG 15
pig install pg_rewrite -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_rewrite';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_rewrite;
```




## Usage

> [pg_rewrite: Tool allows read write to the table during the rewriting](https://github.com/cybertec-postgresql/pg_rewrite)

`pg_rewrite` requires `wal_level = logical` and must be added to `shared_preload_libraries`. Common use cases include changing column data types, partitioning tables, reordering columns, and moving tables to different tablespaces -- all while allowing concurrent reads and writes.

### Rewrite a Table

Create the target table with the desired schema, then call `rewrite_table()`:

```sql
-- Source table
CREATE TABLE measurement (id int, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate));

-- Target table with new schema (e.g., bigint id + partitioning)
CREATE TABLE measurement_aux (id bigint, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate))
  PARTITION BY RANGE (logdate);

CREATE TABLE measurement_y2006m02 PARTITION OF measurement_aux FOR VALUES FROM ('2006-02-01') TO ('2006-03-01');

-- Perform the rewrite (copies data, applies concurrent changes, then renames)
SELECT rewrite_table('measurement', 'measurement_aux', 'measurement_old');
```

Both source and target tables must have an identity index (typically a primary key). The function copies all rows, replays concurrent changes via logical decoding, then atomically renames the tables.

### Progress Monitoring

```sql
SELECT * FROM pg_rewrite_progress;
```

Shows `ins_initial` (initial rows copied), `ins`, `upd`, `del` (concurrent changes applied).

### Configuration

- **`rewrite.max_xlock_time`** -- Maximum time (ms) the exclusive lock is held during the final rename stage. Default `0` (unlimited). Set to e.g. `100` to limit lock duration to 0.1s; the function will retry if exceeded.

```sql
SET rewrite.max_xlock_time TO 100;
```

### Constraints Handling

- PRIMARY KEY, UNIQUE, EXCLUDE: add to target table before calling `rewrite_table()`
- CHECK, NOT NULL (PG18+), FOREIGN KEY: created automatically as NOT VALID; validate with `ALTER TABLE ... VALIDATE CONSTRAINT ...`
