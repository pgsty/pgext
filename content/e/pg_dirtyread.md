---
title: "pg_dirtyread"
linkTitle: "pg_dirtyread"
description: "Read dead but unvacuumed rows from table"
weight: 5050
categories: ["ADMIN"]
width: full
---

[**pg_dirtyread**](https://github.com/df7cb/pg_dirtyread) : Read dead but unvacuumed rows from table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5050** | {{< badge content="pg_dirtyread" link="https://github.com/df7cb/pg_dirtyread" >}} | {{< ext "pg_dirtyread" >}} | `2.8` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_orphaned" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "amcheck" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dirtyread` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.8` | {{< bg "18" "pg_dirtyread_18" "green" >}} {{< bg "17" "pg_dirtyread_17" "green" >}} {{< bg "16" "pg_dirtyread_16" "green" >}} {{< bg "15" "pg_dirtyread_15" "green" >}} {{< bg "14" "pg_dirtyread_14" "green" >}} | `pg_dirtyread_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8` | {{< bg "18" "postgresql-18-dirtyread" "green" >}} {{< bg "17" "postgresql-17-dirtyread" "green" >}} {{< bg "16" "postgresql-16-dirtyread" "green" >}} {{< bg "15" "postgresql-15-dirtyread" "green" >}} {{< bg "14" "postgresql-14-dirtyread" "green" >}} | `postgresql-$v-dirtyread` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.8" "pg_dirtyread_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.8" "postgresql-18-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-17-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-16-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-15-dirtyread : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8" "postgresql-14-dirtyread : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dirtyread_18` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dirtyread_18-2.8-1PIGSTY.el8.x86_64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.0 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dirtyread_18-2.7-4PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_18` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dirtyread_18-2.8-1PIGSTY.el8.aarch64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.9 KiB | [pg_dirtyread_18-2.7-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dirtyread_18-2.7-4PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_18` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dirtyread_18-2.8-1PIGSTY.el9.x86_64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dirtyread_18-2.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dirtyread_18-2.7-6PGDG.rhel9.8.x86_64.rpm) |
| `pg_dirtyread_18` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dirtyread_18-2.8-1PIGSTY.el9.aarch64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.9 KiB | [pg_dirtyread_18-2.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dirtyread_18-2.7-6PGDG.rhel9.8.aarch64.rpm) |
| `pg_dirtyread_18` | `2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.3 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dirtyread_18-2.8-1PIGSTY.el10.x86_64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pg_dirtyread_18-2.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dirtyread_18-2.7-6PGDG.rhel10.2.x86_64.rpm) |
| `pg_dirtyread_18` | `2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_18-2.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dirtyread_18-2.8-1PIGSTY.el10.aarch64.rpm) |
| `pg_dirtyread_18` | `2.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.3 KiB | [pg_dirtyread_18-2.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dirtyread_18-2.7-6PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-dirtyread` | `2.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.1 KiB | [postgresql-18-dirtyread_2.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg12+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.1 KiB | [postgresql-18-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg12+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-18-dirtyread_2.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg13+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-18-dirtyread_2.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg13+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 21.0 KiB | [postgresql-18-dirtyread_2.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg13+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg13+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 22.0 KiB | [postgresql-18-dirtyread_2.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 22.0 KiB | [postgresql-18-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 21.5 KiB | [postgresql-18-dirtyread_2.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 21.4 KiB | [postgresql-18-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.2 KiB | [postgresql-18-dirtyread_2.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.2 KiB | [postgresql-18-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 20.9 KiB | [postgresql-18-dirtyread_2.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 21.3 KiB | [postgresql-18-dirtyread_2.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-dirtyread` | `2.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.5 KiB | [postgresql-18-dirtyread_2.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.8-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-dirtyread` | `2.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.8 KiB | [postgresql-18-dirtyread_2.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-18-dirtyread_2.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dirtyread_17` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dirtyread_17-2.8-1PIGSTY.el8.x86_64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.8 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dirtyread_17-2.7-2PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_17` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dirtyread_17-2.8-1PIGSTY.el8.aarch64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.7 KiB | [pg_dirtyread_17-2.7-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dirtyread_17-2.7-2PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_17` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dirtyread_17-2.8-1PIGSTY.el9.x86_64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dirtyread_17-2.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dirtyread_17-2.7-6PGDG.rhel9.8.x86_64.rpm) |
| `pg_dirtyread_17` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dirtyread_17-2.8-1PIGSTY.el9.aarch64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.9 KiB | [pg_dirtyread_17-2.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dirtyread_17-2.7-6PGDG.rhel9.8.aarch64.rpm) |
| `pg_dirtyread_17` | `2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.3 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dirtyread_17-2.8-1PIGSTY.el10.x86_64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.5 KiB | [pg_dirtyread_17-2.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dirtyread_17-2.7-6PGDG.rhel10.2.x86_64.rpm) |
| `pg_dirtyread_17` | `2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_17-2.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dirtyread_17-2.8-1PIGSTY.el10.aarch64.rpm) |
| `pg_dirtyread_17` | `2.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.3 KiB | [pg_dirtyread_17-2.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dirtyread_17-2.7-6PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-dirtyread` | `2.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.0 KiB | [postgresql-17-dirtyread_2.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg12+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg12+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.0 KiB | [postgresql-17-dirtyread_2.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg13+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg13+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg13+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg13+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 26.0 KiB | [postgresql-17-dirtyread_2.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 26.0 KiB | [postgresql-17-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.4 KiB | [postgresql-17-dirtyread_2.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.3 KiB | [postgresql-17-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-17-dirtyread_2.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-17-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-17-dirtyread_2.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 20.8 KiB | [postgresql-17-dirtyread_2.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 21.1 KiB | [postgresql-17-dirtyread_2.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-dirtyread` | `2.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.4 KiB | [postgresql-17-dirtyread_2.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.8-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-dirtyread` | `2.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.7 KiB | [postgresql-17-dirtyread_2.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-17-dirtyread_2.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dirtyread_16` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dirtyread_16-2.8-1PIGSTY.el8.x86_64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.7 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dirtyread_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_16` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dirtyread_16-2.8-1PIGSTY.el8.aarch64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_dirtyread_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dirtyread_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_16` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dirtyread_16-2.8-1PIGSTY.el9.x86_64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dirtyread_16-2.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dirtyread_16-2.7-6PGDG.rhel9.8.x86_64.rpm) |
| `pg_dirtyread_16` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dirtyread_16-2.8-1PIGSTY.el9.aarch64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.9 KiB | [pg_dirtyread_16-2.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dirtyread_16-2.7-6PGDG.rhel9.8.aarch64.rpm) |
| `pg_dirtyread_16` | `2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.3 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dirtyread_16-2.8-1PIGSTY.el10.x86_64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.5 KiB | [pg_dirtyread_16-2.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dirtyread_16-2.7-6PGDG.rhel10.2.x86_64.rpm) |
| `pg_dirtyread_16` | `2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_16-2.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dirtyread_16-2.8-1PIGSTY.el10.aarch64.rpm) |
| `pg_dirtyread_16` | `2.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.3 KiB | [pg_dirtyread_16-2.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dirtyread_16-2.7-6PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-dirtyread` | `2.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.0 KiB | [postgresql-16-dirtyread_2.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg12+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 20.9 KiB | [postgresql-16-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg12+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.0 KiB | [postgresql-16-dirtyread_2.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg13+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.0 KiB | [postgresql-16-dirtyread_2.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg13+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-16-dirtyread_2.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg13+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-16-dirtyread_2.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg13+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.5 KiB | [postgresql-16-dirtyread_2.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.5 KiB | [postgresql-16-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 24.9 KiB | [postgresql-16-dirtyread_2.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 24.9 KiB | [postgresql-16-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-16-dirtyread_2.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-16-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.8 KiB | [postgresql-16-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 20.7 KiB | [postgresql-16-dirtyread_2.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 21.1 KiB | [postgresql-16-dirtyread_2.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-dirtyread` | `2.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.4 KiB | [postgresql-16-dirtyread_2.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.8-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-dirtyread` | `2.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.7 KiB | [postgresql-16-dirtyread_2.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-16-dirtyread_2.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dirtyread_15` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dirtyread_15-2.8-1PIGSTY.el8.x86_64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.8 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dirtyread_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_15` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dirtyread_15-2.8-1PIGSTY.el8.aarch64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_dirtyread_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dirtyread_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_15` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.3 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dirtyread_15-2.8-1PIGSTY.el9.x86_64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.4 KiB | [pg_dirtyread_15-2.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dirtyread_15-2.7-6PGDG.rhel9.8.x86_64.rpm) |
| `pg_dirtyread_15` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dirtyread_15-2.8-1PIGSTY.el9.aarch64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dirtyread_15-2.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dirtyread_15-2.7-6PGDG.rhel9.8.aarch64.rpm) |
| `pg_dirtyread_15` | `2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.5 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dirtyread_15-2.8-1PIGSTY.el10.x86_64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pg_dirtyread_15-2.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dirtyread_15-2.7-6PGDG.rhel10.2.x86_64.rpm) |
| `pg_dirtyread_15` | `2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [pg_dirtyread_15-2.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dirtyread_15-2.8-1PIGSTY.el10.aarch64.rpm) |
| `pg_dirtyread_15` | `2.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dirtyread_15-2.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dirtyread_15-2.7-6PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-dirtyread` | `2.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg12+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg12+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-15-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg13+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg13+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 21.0 KiB | [postgresql-15-dirtyread_2.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg13+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg13+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.6 KiB | [postgresql-15-dirtyread_2.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.6 KiB | [postgresql-15-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.1 KiB | [postgresql-15-dirtyread_2.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.0 KiB | [postgresql-15-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.2 KiB | [postgresql-15-dirtyread_2.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-15-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 20.8 KiB | [postgresql-15-dirtyread_2.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 21.2 KiB | [postgresql-15-dirtyread_2.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-dirtyread` | `2.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.5 KiB | [postgresql-15-dirtyread_2.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.8-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-dirtyread` | `2.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.9 KiB | [postgresql-15-dirtyread_2.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-15-dirtyread_2.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dirtyread_14` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_dirtyread_14-2.8-1PIGSTY.el8.x86_64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.8 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dirtyread_14-2.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_dirtyread_14` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dirtyread_14-2.8-1PIGSTY.el8.aarch64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.6 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_dirtyread_14` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.3 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_dirtyread_14-2.8-1PIGSTY.el9.x86_64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.4 KiB | [pg_dirtyread_14-2.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dirtyread_14-2.7-6PGDG.rhel9.8.x86_64.rpm) |
| `pg_dirtyread_14` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.2 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_dirtyread_14-2.8-1PIGSTY.el9.aarch64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dirtyread_14-2.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dirtyread_14-2.7-6PGDG.rhel9.8.aarch64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.8 KiB | [pg_dirtyread_14-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dirtyread_14-2.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_dirtyread_14` | `2.8` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.4 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_dirtyread_14-2.8-1PIGSTY.el10.x86_64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.6 KiB | [pg_dirtyread_14-2.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dirtyread_14-2.7-6PGDG.rhel10.2.x86_64.rpm) |
| `pg_dirtyread_14` | `2.8` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.5 KiB | [pg_dirtyread_14-2.8-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_dirtyread_14-2.8-1PIGSTY.el10.aarch64.rpm) |
| `pg_dirtyread_14` | `2.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dirtyread_14-2.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dirtyread_14-2.7-6PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-dirtyread` | `2.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.1 KiB | [postgresql-14-dirtyread_2.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg12+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 21.0 KiB | [postgresql-14-dirtyread_2.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.9 KiB | [postgresql-14-dirtyread_2.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg12+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 20.8 KiB | [postgresql-14-dirtyread_2.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg12+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-14-dirtyread_2.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg13+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 21.1 KiB | [postgresql-14-dirtyread_2.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg13+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-14-dirtyread_2.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg13+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 20.9 KiB | [postgresql-14-dirtyread_2.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg13+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.6 KiB | [postgresql-14-dirtyread_2.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 25.5 KiB | [postgresql-14-dirtyread_2.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.0 KiB | [postgresql-14-dirtyread_2.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 25.0 KiB | [postgresql-14-dirtyread_2.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.2 KiB | [postgresql-14-dirtyread_2.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 21.1 KiB | [postgresql-14-dirtyread_2.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.9 KiB | [postgresql-14-dirtyread_2.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 20.8 KiB | [postgresql-14-dirtyread_2.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 20.8 KiB | [postgresql-14-dirtyread_2.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 21.2 KiB | [postgresql-14-dirtyread_2.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-dirtyread` | `2.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.5 KiB | [postgresql-14-dirtyread_2.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.8-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-dirtyread` | `2.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 20.9 KiB | [postgresql-14-dirtyread_2.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-dirtyread/postgresql-14-dirtyread_2.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/pg_dirtyread" title="Repository" icon="github" subtitle="github.com/df7cb/pg_dirtyread" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_dirtyread-2.8.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_dirtyread;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dirtyread;		# install via package name, for the active PG version

pig install pg_dirtyread -v 18;   # install for PG 18
pig install pg_dirtyread -v 17;   # install for PG 17
pig install pg_dirtyread -v 16;   # install for PG 16
pig install pg_dirtyread -v 15;   # install for PG 15
pig install pg_dirtyread -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dirtyread;
```

## Usage

Sources: [upstream README](https://github.com/df7cb/pg_dirtyread), [Debian changelog](https://github.com/df7cb/pg_dirtyread/blob/master/debian/changelog), [tags](https://github.com/df7cb/pg_dirtyread/tags).

`pg_dirtyread` reads dead or updated heap rows that have not yet been vacuumed away. The function returns `record`, so every call needs a table alias that declares the columns you want back.

### Basic Usage

```sql
CREATE EXTENSION pg_dirtyread;

SELECT *
FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

Columns are matched by name, so the alias can omit columns or place them in a different order.

### Example

```sql
CREATE TABLE foo (bar bigint, baz text);
ALTER TABLE foo SET (autovacuum_enabled = false, toast.autovacuum_enabled = false);

INSERT INTO foo VALUES (1, 'Test'), (2, 'New Test');
DELETE FROM foo WHERE bar = 1;

SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

The deleted row can remain visible until vacuum removes it.

### Dropped Columns

Dropped column contents can be retrieved as long as the table has not been rewritten by operations such as `VACUUM FULL` or `CLUSTER`. Use `dropped_N`, where `N` is the original 1-based column position:

```sql
CREATE TABLE ab(a text, b text);
INSERT INTO ab VALUES ('Hello', 'World');
ALTER TABLE ab DROP COLUMN b;
DELETE FROM ab;

SELECT *
FROM pg_dirtyread('ab') AS ab(a text, dropped_2 text);
```

Only limited type checks are possible because PostgreSQL removes the dropped column's original type metadata.

### System Columns

Include system columns in the alias to retrieve them. A special `dead boolean` column reports rows that are surely dead:

```sql
SELECT *
FROM pg_dirtyread('foo') AS t(
  tableoid oid,
  ctid tid,
  xmin xid,
  xmax xid,
  cmin cid,
  cmax cid,
  dead boolean,
  bar bigint,
  baz text
);
```

The `dead` column is not usable during recovery, including on standby servers. The `oid` system column is only available on PostgreSQL 11 and earlier.

### Caveats

- Pigsty packages `pg_dirtyread` 2.8 as RPMs for PostgreSQL 14-18; DEB availability comes from PGDG as `postgresql-$v-dirtyread`.
- Upstream 2.8 is a PostgreSQL 19 compatibility release for the default TOAST compression change to `lz4`; no new user-facing SQL function is documented.
- `pg_dirtyread` is for forensic and recovery-style inspection. It bypasses normal MVCC visibility expectations and should not be used for application reads.
