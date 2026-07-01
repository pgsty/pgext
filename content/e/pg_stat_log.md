---
title: "pg_stat_log"
linkTitle: "pg_stat_log"
description: "Track cumulative PostgreSQL log message statistics by backend, database, user, level, and SQLSTATE."
weight: 6040
categories: ["STAT"]
width: full
---

[**pg_stat_log**](https://github.com/fabriziomello/pg_stat_log) : Track cumulative PostgreSQL log message statistics by backend, database, user, level, and SQLSTATE.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6040** | {{< badge content="pg_stat_log" link="https://github.com/fabriziomello/pg_stat_log" >}} | {{< ext "pg_stat_log" >}} | `0.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_statements" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_plans" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_log` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.1` | {{< bg "18" "pg_stat_log_18" "green" >}} {{< bg "17" "pg_stat_log_17" "red" >}} {{< bg "16" "pg_stat_log_16" "red" >}} {{< bg "15" "pg_stat_log_15" "red" >}} {{< bg "14" "pg_stat_log_14" "red" >}} | `pg_stat_log_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.1` | {{< bg "18" "postgresql-18-stat-log" "green" >}} {{< bg "17" "postgresql-17-stat-log" "red" >}} {{< bg "16" "postgresql-16-stat-log" "red" >}} {{< bg "15" "postgresql-15-stat-log" "red" >}} {{< bg "14" "postgresql-14-stat-log" "red" >}} | `postgresql-$v-stat-log` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.1" "pg_stat_log_18 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_log_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_log_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.1" "postgresql-18-stat-log : AVAIL 2" "blue" >}} |      {{< bg "MISS" "postgresql-17-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-stat-log : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-stat-log : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_log_18` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pg_stat_log_18-0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_log_18-0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [pg_stat_log_18-0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_log_18-0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_log_18-0.1-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_log_18-0.1-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.5 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_log_18-0.1-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_log_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pg_stat_log_18-0.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_log_18-0.1-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-stat-log` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg12+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 42.2 KiB | [postgresql-18-stat-log_0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg12+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 42.2 KiB | [postgresql-18-stat-log_0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg13+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 42.1 KiB | [postgresql-18-stat-log_0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg13+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 42.7 KiB | [postgresql-18-stat-log_0.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 42.6 KiB | [postgresql-18-stat-log_0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.4 KiB | [postgresql-18-stat-log_0.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.4 KiB | [postgresql-18-stat-log_0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.5 KiB | [postgresql-18-stat-log_0.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 42.4 KiB | [postgresql-18-stat-log_0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.4 KiB | [postgresql-18-stat-log_0.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 42.2 KiB | [postgresql-18-stat-log_0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 42.2 KiB | [postgresql-18-stat-log_0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 42.3 KiB | [postgresql-18-stat-log_0.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-stat-log` | `0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 42.2 KiB | [postgresql-18-stat-log_0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-log/postgresql-18-stat-log_0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fabriziomello/pg_stat_log" title="Repository" icon="github" subtitle="github.com/fabriziomello/pg_stat_log" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_log;		# install via package name, for the active PG version

pig install pg_stat_log -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/pg_stat_log';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_log;
```




## Usage

Sources: [README](https://github.com/fabriziomello/pg_stat_log/blob/main/README.md), [SQL objects](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log--0.1.sql), [control file](https://github.com/fabriziomello/pg_stat_log/blob/main/pg_stat_log.control)

`pg_stat_log` collects cumulative statistics about PostgreSQL log messages. It hooks into `emit_log_hook` and counts messages by backend type, database, user, severity, SQLSTATE, and SQLSTATE condition name.

### Enable

`pg_stat_log` requires PostgreSQL 18 or newer and must be preloaded:

```conf
shared_preload_libraries = 'pg_stat_log'
```

Restart PostgreSQL, then create the extension:

```sql
CREATE EXTENSION pg_stat_log;
```

### View Statistics

```sql
SELECT *
FROM pg_stat_log
ORDER BY count DESC;
```

The view exposes `backend_type`, `database_oid`, `database_name`, `user_oid`, `user_name`, `elevel`, `sqlerrcode`, `sqlerrcode_name`, and `count`.

### Common Queries

```sql
SELECT elevel, sqlerrcode, sqlerrcode_name, sum(count) AS total
FROM pg_stat_log
GROUP BY elevel, sqlerrcode, sqlerrcode_name
ORDER BY total DESC
LIMIT 10;

SELECT backend_type, elevel, sqlerrcode_name, count
FROM pg_stat_log
WHERE backend_type <> 'client backend'
ORDER BY count DESC;
```

### Reset And Capacity

```sql
SELECT pg_stat_log_reset();
SELECT * FROM pg_stat_log_info();
```

`pg_stat_log_info()` reports `max_entries`, `num_entries`, `n_dropped`, and `stats_reset`. Increase `pg_stat_log.max_entries` if `n_dropped` grows.

### Configuration

Settings include `pg_stat_log.enabled`, `pg_stat_log.min_error_level`, and `pg_stat_log.max_entries`.

`emit_log_hook` only sees messages that reach the server log. `log_min_messages` therefore acts as a floor for what can be counted.
