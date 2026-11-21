---
title: "pgaudit"
linkTitle: "pgaudit"
description: "provides auditing functionality"
weight: 7100
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pgaudit**](https://github.com/pgaudit/pgaudit) : provides auditing functionality


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7100** | {{< badge content="pgaudit" link="https://github.com/pgaudit/pgaudit" >}} | {{< ext "pgaudit" >}} | `19.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgauditlogtofile" >}} {{< ext "pg_auditor" >}} {{< ext "pg_roast" >}} {{< ext "pgmemento" >}} {{< ext "table_log" >}} {{< ext "pg_drop_events" >}} {{< ext "pgelog" >}} {{< ext "logerrors" >}} {{< ext "pg_stat_log" >}} {{< ext "pg_auth_mon" >}} {{< ext "data_historization" >}} {{< ext "table_version" >}} |

> [!Note] Preload; PG14/15 RPM names differ.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `19.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pgaudit` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `19.0` | {{< bg "19" "pgaudit_19" "green" >}} {{< bg "18" "pgaudit_18" "green" >}} {{< bg "17" "pgaudit_17" "green" >}} {{< bg "16" "pgaudit_16" "green" >}} {{< bg "15" "pgaudit_15" "green" >}} | `pgaudit_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `19.0` | {{< bg "19" "postgresql-19-pgaudit" "green" >}} {{< bg "18" "postgresql-18-pgaudit" "green" >}} {{< bg "17" "postgresql-17-pgaudit" "green" >}} {{< bg "16" "postgresql-16-pgaudit" "green" >}} {{< bg "15" "postgresql-15-pgaudit" "green" >}} | `postgresql-$v-pgaudit` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pgaudit_19 : N/A 0" "gray" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 4" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pgaudit_19 : N/A 0" "gray" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 4" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 19.0" "pgaudit_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 19.0" "pgaudit_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 19.0" "pgaudit_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 19.0" "pgaudit_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 18.1" "pgaudit_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 17.2" "pgaudit_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 16.2" "pgaudit_16 : AVAIL 6" "blue" >}} | {{< bg "PIGSTY 1.7.1" "pgaudit_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 19.0" "postgresql-19-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.1" "postgresql-18-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 17.2" "postgresql-17-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 16.2" "postgresql-16-pgaudit : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.2" "postgresql-15-pgaudit : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_19` | `19.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.8 KiB | [pgaudit_19-19.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pgaudit_19-19.0-1PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_19` | `19.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [pgaudit_19-19.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pgaudit_19-19.0-1PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_19` | `19.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.9 KiB | [pgaudit_19-19.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pgaudit_19-19.0-1PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_19` | `19.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.9 KiB | [pgaudit_19-19.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pgaudit_19-19.0-1PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pgaudit` | `19.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg12+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg12+2_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg12+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 47.1 KiB | [postgresql-19-pgaudit_19.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg12+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 47.2 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg12+2_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 47.1 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg12+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg13+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg13+2_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg13+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 47.4 KiB | [postgresql-19-pgaudit_19.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg13+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 47.4 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg13+2_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 47.3 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg13+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.1 KiB | [postgresql-19-pgaudit_19.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.1 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.0 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.5 KiB | [postgresql-19-pgaudit_19.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.5 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.4 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 48.0 KiB | [postgresql-19-pgaudit_19.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 48.0 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 47.9 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 47.2 KiB | [postgresql-19-pgaudit_19.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 47.2 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 47.2 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 47.4 KiB | [postgresql-19-pgaudit_19.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 47.4 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-pgaudit` | `19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 47.4 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-pgaudit` | `19.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.8 KiB | [postgresql-19-pgaudit_19.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-pgaudit` | `19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.8 KiB | [postgresql-19-pgaudit_19~rc1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-19/postgresql-19-pgaudit_19~rc1-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_18` | `18.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pgaudit_18-18.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgaudit_18-18.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_18` | `18.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.7 KiB | [pgaudit_18-18.0-9PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgaudit_18-18.0-9PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_18` | `18.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.5 KiB | [pgaudit_18-18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgaudit_18-18.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_18` | `18.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [pgaudit_18-18.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgaudit_18-18.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.2 KiB | [pgaudit_18-18.0-9PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgaudit_18-18.0-9PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.0 KiB | [pgaudit_18-18.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgaudit_18-18.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_18` | `18.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [pgaudit_18-18.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgaudit_18-18.1-1PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_18` | `18.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.1 KiB | [pgaudit_18-18.0-9PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgaudit_18-18.0-9PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_18` | `18.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [pgaudit_18-18.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgaudit_18-18.1-1PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.8 KiB | [pgaudit_18-18.0-9PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgaudit_18-18.0-9PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_18` | `18.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.2 KiB | [pgaudit_18-18.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgaudit_18-18.1-1PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_18` | `18.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.1 KiB | [pgaudit_18-18.0-9PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgaudit_18-18.0-9PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_18` | `18.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.1 KiB | [pgaudit_18-18.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.1-1PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_18` | `18.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.1 KiB | [pgaudit_18-18.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.2 KiB | [pgaudit_18-18.0-9PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.0-9PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.2 KiB | [pgaudit_18-18.0-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.0-9PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgaudit_18-18.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.0-4PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_18` | `18.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgaudit_18-18.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgaudit` | `18.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 48.2 KiB | [postgresql-18-pgaudit_18.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.2 KiB | [postgresql-18-pgaudit_18.0-3.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg12+2_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.2 KiB | [postgresql-18-pgaudit_18.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 47.4 KiB | [postgresql-18-pgaudit_18.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.5 KiB | [postgresql-18-pgaudit_18.0-3.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg12+2_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.4 KiB | [postgresql-18-pgaudit_18.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 48.2 KiB | [postgresql-18-pgaudit_18.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.2 KiB | [postgresql-18-pgaudit_18.0-3.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg13+2_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.1 KiB | [postgresql-18-pgaudit_18.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 47.6 KiB | [postgresql-18-pgaudit_18.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.7 KiB | [postgresql-18-pgaudit_18.0-3.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg13+2_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.6 KiB | [postgresql-18-pgaudit_18.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.6 KiB | [postgresql-18-pgaudit_18.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.8 KiB | [postgresql-18-pgaudit_18.0-3.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg22.04+2_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.7 KiB | [postgresql-18-pgaudit_18.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.0 KiB | [postgresql-18-pgaudit_18.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.0 KiB | [postgresql-18-pgaudit_18.0-3.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg22.04+2_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.0 KiB | [postgresql-18-pgaudit_18.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 48.3 KiB | [postgresql-18-pgaudit_18.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 47.3 KiB | [postgresql-18-pgaudit_18.0-3.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg24.04+2_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 47.3 KiB | [postgresql-18-pgaudit_18.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 47.6 KiB | [postgresql-18-pgaudit_18.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 46.5 KiB | [postgresql-18-pgaudit_18.0-3.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg24.04+2_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 46.5 KiB | [postgresql-18-pgaudit_18.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 47.7 KiB | [postgresql-18-pgaudit_18.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.6 KiB | [postgresql-18-pgaudit_18.0-3.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg26.04+2_amd64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.6 KiB | [postgresql-18-pgaudit_18.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | `18.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 47.1 KiB | [postgresql-18-pgaudit_18.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.4 KiB | [postgresql-18-pgaudit_18.0-3.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg26.04+2_arm64.deb) |
| `postgresql-18-pgaudit` | `18.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.3 KiB | [postgresql-18-pgaudit_18.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-3.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_17` | `17.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [pgaudit_17-17.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_17` | `17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pgaudit_17-17.1-9PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.1-9PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_17` | `17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.0 KiB | [pgaudit_17-17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.1-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_17` | `17.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.5 KiB | [pgaudit_17-17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_17` | `17.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.9 KiB | [pgaudit_17-17.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.8 KiB | [pgaudit_17-17.1-9PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.1-9PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.6 KiB | [pgaudit_17-17.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.1-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_17` | `17.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [pgaudit_17-17.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_17` | `17.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [pgaudit_17-17.2-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgaudit_17-17.2-1PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_17` | `17.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.2 KiB | [pgaudit_17-17.1-9PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgaudit_17-17.1-9PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_17` | `17.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [pgaudit_17-17.2-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgaudit_17-17.2-1PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.9 KiB | [pgaudit_17-17.1-9PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgaudit_17-17.1-9PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_17` | `17.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.4 KiB | [pgaudit_17-17.2-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgaudit_17-17.2-1PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_17` | `17.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.2 KiB | [pgaudit_17-17.1-9PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgaudit_17-17.1-9PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_17` | `17.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.4 KiB | [pgaudit_17-17.2-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.2-1PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_17` | `17.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.4 KiB | [pgaudit_17-17.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.4 KiB | [pgaudit_17-17.1-9PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.1-9PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.4 KiB | [pgaudit_17-17.1-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.1-9PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.7 KiB | [pgaudit_17-17.1-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.1-4PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_17` | `17.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.7 KiB | [pgaudit_17-17.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgaudit` | `17.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.6 KiB | [postgresql-17-pgaudit_17.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 46.1 KiB | [postgresql-17-pgaudit_17.1-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg12+2_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 46.1 KiB | [postgresql-17-pgaudit_17.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 47.1 KiB | [postgresql-17-pgaudit_17.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.8 KiB | [postgresql-17-pgaudit_17.1-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg12+2_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 45.7 KiB | [postgresql-17-pgaudit_17.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.6 KiB | [postgresql-17-pgaudit_17.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 46.1 KiB | [postgresql-17-pgaudit_17.1-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg13+2_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 46.0 KiB | [postgresql-17-pgaudit_17.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 47.3 KiB | [postgresql-17-pgaudit_17.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.9 KiB | [postgresql-17-pgaudit_17.1-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg13+2_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.9 KiB | [postgresql-17-pgaudit_17.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 53.8 KiB | [postgresql-17-pgaudit_17.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.6 KiB | [postgresql-17-pgaudit_17.1-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg22.04+2_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.6 KiB | [postgresql-17-pgaudit_17.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 53.2 KiB | [postgresql-17-pgaudit_17.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.1 KiB | [postgresql-17-pgaudit_17.1-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg22.04+2_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.0 KiB | [postgresql-17-pgaudit_17.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 47.7 KiB | [postgresql-17-pgaudit_17.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.2 KiB | [postgresql-17-pgaudit_17.1-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg24.04+2_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.2 KiB | [postgresql-17-pgaudit_17.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 47.2 KiB | [postgresql-17-pgaudit_17.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.8 KiB | [postgresql-17-pgaudit_17.1-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg24.04+2_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.9 KiB | [postgresql-17-pgaudit_17.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 47.2 KiB | [postgresql-17-pgaudit_17.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.8 KiB | [postgresql-17-pgaudit_17.1-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg26.04+2_amd64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.2 KiB | [postgresql-17-pgaudit_17.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | `17.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 46.7 KiB | [postgresql-17-pgaudit_17.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-17-pgaudit_17.1-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg26.04+2_arm64.deb) |
| `postgresql-17-pgaudit` | `17.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.9 KiB | [postgresql-17-pgaudit_17.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-2.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_16` | `16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [pgaudit_16-16.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_16` | `16.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 KiB | [pgaudit_16-16.1-9PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.1-9PGDG.rhel8.10.x86_64.rpm) |
| `pgaudit_16` | `16.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.4 KiB | [pgaudit_16-16.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.1-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_16` | `16.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.9 KiB | [pgaudit_16-16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_16` | `16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.4 KiB | [pgaudit_16-16.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.3 KiB | [pgaudit_16-16.1-9PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.1-9PGDG.rhel8.10.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.1 KiB | [pgaudit_16-16.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.1-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_16` | `16.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [pgaudit_16-16.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_16` | `16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [pgaudit_16-16.2-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgaudit_16-16.2-1PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_16` | `16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.9 KiB | [pgaudit_16-16.1-9PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgaudit_16-16.1-9PGDG.rhel9.9.x86_64.rpm) |
| `pgaudit_16` | `16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [pgaudit_16-16.2-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgaudit_16-16.2-1PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.7 KiB | [pgaudit_16-16.1-9PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgaudit_16-16.1-9PGDG.rhel9.9.aarch64.rpm) |
| `pgaudit_16` | `16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.1 KiB | [pgaudit_16-16.2-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgaudit_16-16.2-1PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_16` | `16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.0 KiB | [pgaudit_16-16.1-9PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgaudit_16-16.1-9PGDG.rhel10.3.x86_64.rpm) |
| `pgaudit_16` | `16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [pgaudit_16-16.2-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.2-1PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_16` | `16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [pgaudit_16-16.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.1 KiB | [pgaudit_16-16.1-9PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.1-9PGDG.rhel10.3.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.1 KiB | [pgaudit_16-16.1-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.1-9PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.2 KiB | [pgaudit_16-16.1-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.1-4PGDG.rhel10.2.aarch64.rpm) |
| `pgaudit_16` | `16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.3 KiB | [pgaudit_16-16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgaudit` | `16.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 46.7 KiB | [postgresql-16-pgaudit_16.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.1 KiB | [postgresql-16-pgaudit_16.1-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg12+2_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-16-pgaudit_16.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.2 KiB | [postgresql-16-pgaudit_16.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgaudit_16.1-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg12+2_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgaudit_16.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 46.6 KiB | [postgresql-16-pgaudit_16.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.3 KiB | [postgresql-16-pgaudit_16.1-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg13+2_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-16-pgaudit_16.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.4 KiB | [postgresql-16-pgaudit_16.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.1 KiB | [postgresql-16-pgaudit_16.1-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg13+2_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.2 KiB | [postgresql-16-pgaudit_16.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.3 KiB | [postgresql-16-pgaudit_16.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.2 KiB | [postgresql-16-pgaudit_16.1-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg22.04+2_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.2 KiB | [postgresql-16-pgaudit_16.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.8 KiB | [postgresql-16-pgaudit_16.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.7 KiB | [postgresql-16-pgaudit_16.1-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg22.04+2_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.7 KiB | [postgresql-16-pgaudit_16.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.7 KiB | [postgresql-16-pgaudit_16.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-16-pgaudit_16.1-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg24.04+2_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-16-pgaudit_16.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 46.3 KiB | [postgresql-16-pgaudit_16.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.1 KiB | [postgresql-16-pgaudit_16.1-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg24.04+2_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.0 KiB | [postgresql-16-pgaudit_16.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.2 KiB | [postgresql-16-pgaudit_16.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.8 KiB | [postgresql-16-pgaudit_16.1-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg26.04+2_amd64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.5 KiB | [postgresql-16-pgaudit_16.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | `16.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.8 KiB | [postgresql-16-pgaudit_16.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.5 KiB | [postgresql-16-pgaudit_16.1-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg26.04+2_arm64.deb) |
| `postgresql-16-pgaudit` | `16.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.0 KiB | [postgresql-16-pgaudit_16.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-2.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_15` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 50.7 KiB | [pgaudit_15-1.7.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgaudit_15-1.7.1-1PGSTY.el8.x86_64.rpm) |
| `pgaudit_15` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 50.4 KiB | [pgaudit_15-1.7.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgaudit_15-1.7.1-1PGSTY.el8.aarch64.rpm) |
| `pgaudit_15` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.6 KiB | [pgaudit_15-1.7.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgaudit_15-1.7.1-1PGSTY.el9.x86_64.rpm) |
| `pgaudit_15` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.4 KiB | [pgaudit_15-1.7.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgaudit_15-1.7.1-1PGSTY.el9.aarch64.rpm) |
| `pgaudit_15` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 51.7 KiB | [pgaudit_15-1.7.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgaudit_15-1.7.1-1PGSTY.el10.x86_64.rpm) |
| `pgaudit_15` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 51.9 KiB | [pgaudit_15-1.7.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgaudit_15-1.7.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgaudit` | `1.7.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.1 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg12+2_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.0 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.8 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg12+2_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.5 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.2 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg13+2_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.1 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 45.1 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.8 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg13+2_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 51.3 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.3 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg22.04+2_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.2 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 50.6 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.5 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg22.04+2_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.5 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.6 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.3 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg24.04+2_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.3 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.0 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.7 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg24.04+2_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.0 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 43.7 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg26.04+2_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.2 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | `1.7.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.8 KiB | [postgresql-15-pgaudit_1.7.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.2-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 43.6 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg26.04+2_arm64.deb) |
| `postgresql-15-pgaudit` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.0 KiB | [postgresql-15-pgaudit_1.7.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-2.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgaudit/pgaudit" title="Repository" icon="github" subtitle="github.com/pgaudit/pgaudit" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgaudit-1.6.3.tar.gz pgaudit-1.7.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgaudit;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgaudit;		# install via package name, for the active PG version

pig install pgaudit -v 19;   # install for PG 19
pig install pgaudit -v 18;   # install for PG 18
pig install pgaudit -v 17;   # install for PG 17
pig install pgaudit -v 16;   # install for PG 16
pig install pgaudit -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgaudit';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pgaudit;
```

## Usage

Sources:

- [pgAudit 19.0 README](https://github.com/pgaudit/pgaudit/blob/19.0/README.md)
- [PostgreSQL 19 release](https://github.com/pgaudit/pgaudit/releases/tag/19.0)
- [Release control file](https://github.com/pgaudit/pgaudit/blob/19.0/pgaudit.control)

`pgaudit` records session and object events through PostgreSQL logging. Use the branch matching the server: 19.0 for PostgreSQL 19, 18.1 for 18, 17.2 for 17, 16.2 for 16, 1.7.2 for 15 and 1.6.4 for 14. The branches are separate server builds; upgrading PostgreSQL does not make an older library ABI-compatible.

### Enable auditing

Add `pgaudit` to `shared_preload_libraries` and restart the server. As a superuser, create the extension before setting `pgaudit.log`:

```sql
CREATE EXTENSION pgaudit;
```

Only a superuser can change audit settings. Size and protect the log volume: SQL text and parameters can contain sensitive information, and broad logging can consume substantial space. The extension provides audit records; it does not itself establish compliance.

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgaudit.log` | `none` | Statement classes to log: `READ`, `WRITE`, `FUNCTION`, `ROLE`, `DDL`, `MISC`, `MISC_SET`, `ALL` |
| `pgaudit.log_catalog` | `on` | Log statements when all relations are in pg_catalog |
| `pgaudit.log_client` | `off` | Show audit log messages to client |
| `pgaudit.log_level` | `log` | Log level for audit entries |
| `pgaudit.log_parameter` | `off` | Include statement parameters in log |
| `pgaudit.log_parameter_max_size` | `0` | Max parameter size in bytes (0=unlimited) |
| `pgaudit.log_relation` | `off` | Separate log entry per relation in SELECT/DML |
| `pgaudit.log_rows` | `off` | Include row count in log |
| `pgaudit.log_statement` | `on` | Include statement text in log |
| `pgaudit.log_statement_once` | `off` | Log statement text only with first entry |
| `pgaudit.role` | (none) | Master role for object audit logging |
{.wide .expand}

### Session Audit Logging

Log all DML and DDL with per-relation detail:

```sql
SET pgaudit.log = 'write, ddl';
SET pgaudit.log_relation = on;
```

Log everything except miscellaneous commands:

```sql
SET pgaudit.log = 'all, -misc';
```

Example output:
```text
AUDIT: SESSION,1,1,DDL,CREATE TABLE,TABLE,public.account,create table account(...)
AUDIT: SESSION,2,1,READ,SELECT,,,select * from account
```

### Object Audit Logging

Grant permissions to an audit role to control which relations are logged:

```sql
SET pgaudit.role = 'auditor';

GRANT SELECT, DELETE
   ON public.account
   TO auditor;
```

Now any `SELECT` or `DELETE` on the `account` table will be audit logged.

### Log Format

Entries are CSV with fields: `AUDIT_TYPE`, `STATEMENT_ID`, `SUBSTATEMENT_ID`, `CLASS`, `COMMAND`, `OBJECT_TYPE`, `OBJECT_NAME`, `STATEMENT`, `PARAMETER`.

### Logging boundary

Audit output is best-effort PostgreSQL logging, outside transaction durability. A rolled-back statement can leave an audit record, and a committed transaction is not guaranteed to have a durable audit record. Protect and collect logs separately; do not treat them as a transactionally complete ledger.
