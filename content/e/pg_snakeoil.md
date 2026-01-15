---
title: "pg_snakeoil"
linkTitle: "pg_snakeoil"
description: "The PostgreSQL Antivirus"
weight: 7380
categories: ["SEC"]
width: full
---

[**pg_snakeoil**](https://github.com/credativ/pg_snakeoil) : The PostgreSQL Antivirus


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7380** | {{< badge content="pg_snakeoil" link="https://github.com/credativ/pg_snakeoil" >}} | {{< ext "pg_snakeoil" >}} | `1.4` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL--r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_crash" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_savior" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_catcheck" >}} {{< ext "amcheck" >}} |

> [!Note] require clamV libs


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_snakeoil` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4` | {{< bg "18" "pg_snakeoil_18" "green" >}} {{< bg "17" "pg_snakeoil_17" "green" >}} {{< bg "16" "pg_snakeoil_16" "green" >}} {{< bg "15" "pg_snakeoil_15" "green" >}} {{< bg "14" "pg_snakeoil_14" "green" >}} {{< bg "13" "pg_snakeoil_13" "green" >}} | `pg_snakeoil_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4` | {{< bg "18" "postgresql-18-snakeoil" "green" >}} {{< bg "17" "postgresql-17-snakeoil" "green" >}} {{< bg "16" "postgresql-16-snakeoil" "green" >}} {{< bg "15" "postgresql-15-snakeoil" "green" >}} {{< bg "14" "postgresql-14-snakeoil" "green" >}} {{< bg "13" "postgresql-13-snakeoil" "green" >}} | `postgresql-$v-snakeoil` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4" "pg_snakeoil_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-18-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-snakeoil : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-snakeoil : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_18` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_18-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.3 KiB | [pg_snakeoil_18-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_snakeoil_18-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_18-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.2 KiB | [pg_snakeoil_18-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_snakeoil_18-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_18-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.5 KiB | [pg_snakeoil_18-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_snakeoil_18-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.0 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_18-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.1 KiB | [pg_snakeoil_18-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_snakeoil_18-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_18-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.1 KiB | [pg_snakeoil_18-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_18-1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 16.9 KiB | [postgresql-18-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.4 KiB | [postgresql-18-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 16.8 KiB | [postgresql-18-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.4 KiB | [postgresql-18-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 16.6 KiB | [postgresql-18-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 16.4 KiB | [postgresql-18-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 16.9 KiB | [postgresql-18-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.4 KiB | [postgresql-18-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-18-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_17-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.4 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_snakeoil_17-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_17-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.3 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_snakeoil_17-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.2 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_17-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.6 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_snakeoil_17-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.0 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_17-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.1 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_snakeoil_17-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_17-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.9 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_snakeoil_17-1.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_17-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_17-1.4-1PIGSTY.el10.aarch64.rpm) |
| `pg_snakeoil_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.8 KiB | [pg_snakeoil_17-1.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_snakeoil_17-1.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 16.9 KiB | [postgresql-17-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.4 KiB | [postgresql-17-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 16.8 KiB | [postgresql-17-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.4 KiB | [postgresql-17-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 17.1 KiB | [postgresql-17-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 17.0 KiB | [postgresql-17-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 16.9 KiB | [postgresql-17-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.4 KiB | [postgresql-17-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-17-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_16-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.4 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_snakeoil_16-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_16-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.2 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_snakeoil_16-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.2 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_16-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.6 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_snakeoil_16-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.0 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_16-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.1 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_snakeoil_16-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_16-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.9 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_snakeoil_16-1.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_16-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_16-1.4-1PIGSTY.el10.aarch64.rpm) |
| `pg_snakeoil_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.8 KiB | [pg_snakeoil_16-1.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_snakeoil_16-1.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 16.9 KiB | [postgresql-16-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.4 KiB | [postgresql-16-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 16.8 KiB | [postgresql-16-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.4 KiB | [postgresql-16-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 17.1 KiB | [postgresql-16-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 17.0 KiB | [postgresql-16-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 16.9 KiB | [postgresql-16-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.4 KiB | [postgresql-16-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-16-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_15-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.5 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_snakeoil_15-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.3 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_15-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.4 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_snakeoil_15-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_15-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.8 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_snakeoil_15-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_15-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.3 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_snakeoil_15-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_15-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.1 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_snakeoil_15-1.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.4 KiB | [pg_snakeoil_15-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_15-1.4-1PIGSTY.el10.aarch64.rpm) |
| `pg_snakeoil_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [pg_snakeoil_15-1.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_snakeoil_15-1.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 17.0 KiB | [postgresql-15-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.5 KiB | [postgresql-15-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 17.0 KiB | [postgresql-15-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.6 KiB | [postgresql-15-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 17.2 KiB | [postgresql-15-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 17.2 KiB | [postgresql-15-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 17.0 KiB | [postgresql-15-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.5 KiB | [postgresql-15-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-15-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_14-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.5 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_snakeoil_14-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.3 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_14-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.4 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_snakeoil_14-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_14-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.8 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_snakeoil_14-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_14-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.3 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_snakeoil_14-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.3 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_14-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.1 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_snakeoil_14-1.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.4 KiB | [pg_snakeoil_14-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_14-1.4-1PIGSTY.el10.aarch64.rpm) |
| `pg_snakeoil_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [pg_snakeoil_14-1.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_snakeoil_14-1.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 17.0 KiB | [postgresql-14-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.5 KiB | [postgresql-14-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 17.0 KiB | [postgresql-14-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.5 KiB | [postgresql-14-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 17.2 KiB | [postgresql-14-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 17.1 KiB | [postgresql-14-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 17.0 KiB | [postgresql-14-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.5 KiB | [postgresql-14-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-14-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_snakeoil_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_snakeoil_13-1.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.4 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_snakeoil_13-1.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.3 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_snakeoil_13-1.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.3 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_snakeoil_13-1.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.2 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_snakeoil_13-1.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.7 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_snakeoil_13-1.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_snakeoil_13-1.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.3 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_snakeoil_13-1.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.2 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_snakeoil_13-1.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.0 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_snakeoil_13-1.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.3 KiB | [pg_snakeoil_13-1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_snakeoil_13-1.4-1PIGSTY.el10.aarch64.rpm) |
| `pg_snakeoil_13` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.9 KiB | [pg_snakeoil_13-1.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_snakeoil_13-1.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-snakeoil` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 16.6 KiB | [postgresql-13-snakeoil_1.4-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg12+2_amd64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 16.5 KiB | [postgresql-13-snakeoil_1.4-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg12+2_arm64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 16.6 KiB | [postgresql-13-snakeoil_1.4-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg13+1_amd64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 16.5 KiB | [postgresql-13-snakeoil_1.4-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg13+1_arm64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 17.1 KiB | [postgresql-13-snakeoil_1.4-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 17.0 KiB | [postgresql-13-snakeoil_1.4-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 16.7 KiB | [postgresql-13-snakeoil_1.4-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-snakeoil` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 16.5 KiB | [postgresql-13-snakeoil_1.4-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-snakeoil/postgresql-13-snakeoil_1.4-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/pg_snakeoil" title="Repository" icon="github" subtitle="github.com/credativ/pg_snakeoil" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_snakeoil-1.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_snakeoil;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_snakeoil;		# install via package name, for the active PG version

pig install pg_snakeoil -v 18;   # install for PG 18
pig install pg_snakeoil -v 17;   # install for PG 17
pig install pg_snakeoil -v 16;   # install for PG 16
pig install pg_snakeoil -v 15;   # install for PG 15
pig install pg_snakeoil -v 14;   # install for PG 14
pig install pg_snakeoil -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_snakeoil';
```


This extension does not need `CREATE EXTENSION` DDL command


