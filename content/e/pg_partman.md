---
title: "pg_partman"
linkTitle: "pg_partman"
description: "Extension to manage partitioned tables by time or ID"
weight: 2510
categories: ["OLAP"]
width: full
---

[**pg_partman**](https://github.com/pgpartman/pg_partman) : Extension to manage partitioned tables by time or ID


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2510** | {{< badge content="pg_partman" link="https://github.com/pgpartman/pg_partman" >}} | {{< ext "pg_partman" >}} | `5.3.1` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "timeseries" >}} |
|   **See Also**    | {{< ext "citus" >}} {{< ext "pg_fkpart" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "emaj" >}} {{< ext "pg_cron" >}} {{< ext "plproxy" >}} {{< ext "temporal_tables" >}} |

> [!Note] no pg13,12 on u24 pgdg repo


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_partman` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.3.1` | {{< bg "18" "pg_partman_18*" "green" >}} {{< bg "17" "pg_partman_17*" "green" >}} {{< bg "16" "pg_partman_16*" "green" >}} {{< bg "15" "pg_partman_15*" "green" >}} {{< bg "14" "pg_partman_14*" "green" >}} {{< bg "13" "pg_partman_13*" "green" >}} | `pg_partman_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.3.0` | {{< bg "18" "postgresql-18-partman" "green" >}} {{< bg "17" "postgresql-17-partman" "green" >}} {{< bg "16" "postgresql-16-partman" "green" >}} {{< bg "15" "postgresql-15-partman" "green" >}} {{< bg "14" "postgresql-14-partman" "green" >}} {{< bg "13" "postgresql-13-partman" "green" >}} | `postgresql-$v-partman` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.7.4" "pg_partman_13 : AVAIL 12" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.7.4" "pg_partman_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 18" "blue" >}} | {{< bg "PGDG 4.7.4" "pg_partman_13 : AVAIL 8" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.7.4" "pg_partman_13 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_partman_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 5.3.1" "pg_partman_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.3.1" "pg_partman_14 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_partman_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.4" "postgresql-13-partman : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-partman : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 5.3.1" "postgresql-18-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-17-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-16-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-15-partman : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.3.1" "postgresql-14-partman : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-partman : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_18` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.3 KiB | [pg_partman_18-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_18-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 262.2 KiB | [pg_partman_18-5.2.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.2.4-2PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_18-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 262.2 KiB | [pg_partman_18-5.2.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.2.4-2PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_18-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.0 KiB | [pg_partman_18-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 208.0 KiB | [pg_partman_18-5.2.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.2.4-2PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.1 KiB | [pg_partman_18-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.5 KiB | [pg_partman_18-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.6 KiB | [pg_partman_18-5.2.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.2.4-2PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.5 KiB | [pg_partman_18-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_18-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.4 KiB | [pg_partman_18-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_18-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-partman` | `5.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 232.8 KiB | [postgresql-18-partman_5.3.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg12+1_amd64.deb) |
| `postgresql-18-partman` | `5.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.7 KiB | [postgresql-18-partman_5.3.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg12+1_arm64.deb) |
| `postgresql-18-partman` | `5.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 232.8 KiB | [postgresql-18-partman_5.3.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg13+1_amd64.deb) |
| `postgresql-18-partman` | `5.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.7 KiB | [postgresql-18-partman_5.3.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg13+1_arm64.deb) |
| `postgresql-18-partman` | `5.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 226.7 KiB | [postgresql-18-partman_5.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 226.3 KiB | [postgresql-18-partman_5.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 225.9 KiB | [postgresql-18-partman_5.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.8 KiB | [postgresql-18-partman_5.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_17` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_17-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 261.4 KiB | [pg_partman_17-5.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_17-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.6 KiB | [pg_partman_17-5.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.3 KiB | [pg_partman_17-5.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 254.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.1.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_17-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 261.3 KiB | [pg_partman_17-5.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_17-5.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.6 KiB | [pg_partman_17-5.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.2 KiB | [pg_partman_17-5.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 254.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.1.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_17-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 212.8 KiB | [pg_partman_17-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 207.4 KiB | [pg_partman_17-5.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.2 KiB | [pg_partman_17-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.9 KiB | [pg_partman_17-5.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.5 KiB | [pg_partman_17-5.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.9 KiB | [pg_partman_17-5.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.1.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.4 KiB | [pg_partman_17-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.5 KiB | [pg_partman_17-5.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.0 KiB | [pg_partman_17-5.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.3 KiB | [pg_partman_17-5.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.8 KiB | [pg_partman_17-5.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.4 KiB | [pg_partman_17-5.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.1.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_17-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.4 KiB | [pg_partman_17-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_17-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_17-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-partman` | `5.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 232.6 KiB | [postgresql-17-partman_5.3.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg12+1_amd64.deb) |
| `postgresql-17-partman` | `5.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.6 KiB | [postgresql-17-partman_5.3.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg12+1_arm64.deb) |
| `postgresql-17-partman` | `5.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 232.7 KiB | [postgresql-17-partman_5.3.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg13+1_amd64.deb) |
| `postgresql-17-partman` | `5.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.6 KiB | [postgresql-17-partman_5.3.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg13+1_arm64.deb) |
| `postgresql-17-partman` | `5.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 231.3 KiB | [postgresql-17-partman_5.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.8 KiB | [postgresql-17-partman_5.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 225.9 KiB | [postgresql-17-partman_5.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.6 KiB | [postgresql-17-partman_5.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_16` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.2 KiB | [pg_partman_16-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_16-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 261.3 KiB | [pg_partman_16-5.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.8 KiB | [pg_partman_16-5.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_16-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.6 KiB | [pg_partman_16-5.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.3 KiB | [pg_partman_16-5.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 254.7 KiB | [pg_partman_16-5.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 249.3 KiB | [pg_partman_16-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 248.4 KiB | [pg_partman_16-5.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `4.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.9 KiB | [pg_partman_16-4.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-4.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_16` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.5 KiB | [pg_partman_16-4.7.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-4.7.3-3.rhel8.x86_64.rpm) |
| `pg_partman_16` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_16-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_16-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 261.3 KiB | [pg_partman_16-5.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.8 KiB | [pg_partman_16-5.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_16-5.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.5 KiB | [pg_partman_16-5.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.2 KiB | [pg_partman_16-5.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 254.7 KiB | [pg_partman_16-5.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 249.3 KiB | [pg_partman_16-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 248.3 KiB | [pg_partman_16-5.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `4.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.8 KiB | [pg_partman_16-4.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-4.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_16` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.4 KiB | [pg_partman_16-4.7.3-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-4.7.3-3.rhel8.aarch64.rpm) |
| `pg_partman_16` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_16-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 212.9 KiB | [pg_partman_16-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 207.3 KiB | [pg_partman_16-5.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.8 KiB | [pg_partman_16-5.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.3 KiB | [pg_partman_16-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.0 KiB | [pg_partman_16-5.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.6 KiB | [pg_partman_16-5.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.8 KiB | [pg_partman_16-5.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.9 KiB | [pg_partman_16-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.3 KiB | [pg_partman_16-5.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `4.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.9 KiB | [pg_partman_16-4.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-4.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_16` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 194.2 KiB | [pg_partman_16-4.7.3-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-4.7.3-3.rhel9.x86_64.rpm) |
| `pg_partman_16` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.1 KiB | [pg_partman_16-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.5 KiB | [pg_partman_16-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.5 KiB | [pg_partman_16-5.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.0 KiB | [pg_partman_16-5.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.2 KiB | [pg_partman_16-5.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.7 KiB | [pg_partman_16-5.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.4 KiB | [pg_partman_16-5.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.6 KiB | [pg_partman_16-5.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.9 KiB | [pg_partman_16-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.0 KiB | [pg_partman_16-5.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `4.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.4 KiB | [pg_partman_16-4.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-4.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_16` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 194.1 KiB | [pg_partman_16-4.7.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-4.7.3-3.rhel9.aarch64.rpm) |
| `pg_partman_16` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.3 KiB | [pg_partman_16-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_16-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_16-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_16-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_16` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_16-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_16` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.6 KiB | [pg_partman_16-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-partman` | `5.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 232.6 KiB | [postgresql-16-partman_5.3.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg12+1_amd64.deb) |
| `postgresql-16-partman` | `5.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.5 KiB | [postgresql-16-partman_5.3.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg12+1_arm64.deb) |
| `postgresql-16-partman` | `5.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 232.6 KiB | [postgresql-16-partman_5.3.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg13+1_amd64.deb) |
| `postgresql-16-partman` | `5.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.5 KiB | [postgresql-16-partman_5.3.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg13+1_arm64.deb) |
| `postgresql-16-partman` | `5.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 230.8 KiB | [postgresql-16-partman_5.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.4 KiB | [postgresql-16-partman_5.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 225.9 KiB | [postgresql-16-partman_5.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.6 KiB | [postgresql-16-partman_5.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_15` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.3 KiB | [pg_partman_15-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 261.4 KiB | [pg_partman_15-5.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.8 KiB | [pg_partman_15-5.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_15-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.6 KiB | [pg_partman_15-5.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.2 KiB | [pg_partman_15-5.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 254.7 KiB | [pg_partman_15-5.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 249.3 KiB | [pg_partman_15-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 248.4 KiB | [pg_partman_15-5.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.9 KiB | [pg_partman_15-4.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.5 KiB | [pg_partman_15-4.7.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.3-3.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.2 KiB | [pg_partman_15-4.7.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.3-1.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 245.7 KiB | [pg_partman_15-4.7.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.2-1.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.6 KiB | [pg_partman_15-4.7.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.1-1.rhel8.x86_64.rpm) |
| `pg_partman_15` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_15-4.7.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-4.7.0-2.rhel8.x86_64.rpm) |
| `pg_partman_15` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_15-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 261.3 KiB | [pg_partman_15-5.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.8 KiB | [pg_partman_15-5.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_15-5.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.5 KiB | [pg_partman_15-5.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.2 KiB | [pg_partman_15-5.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 254.7 KiB | [pg_partman_15-5.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 249.3 KiB | [pg_partman_15-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 248.4 KiB | [pg_partman_15-5.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `4.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.8 KiB | [pg_partman_15-4.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-4.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_15` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.4 KiB | [pg_partman_15-4.7.3-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-4.7.3-3.rhel8.aarch64.rpm) |
| `pg_partman_15` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.1 KiB | [pg_partman_15-4.7.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-4.7.3-1.rhel8.aarch64.rpm) |
| `pg_partman_15` | `4.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 245.6 KiB | [pg_partman_15-4.7.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-4.7.2-1.rhel8.aarch64.rpm) |
| `pg_partman_15` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_15-4.7.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-4.7.1-1.rhel8.aarch64.rpm) |
| `pg_partman_15` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_15-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.1 KiB | [pg_partman_15-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 207.3 KiB | [pg_partman_15-5.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.8 KiB | [pg_partman_15-5.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.2 KiB | [pg_partman_15-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.0 KiB | [pg_partman_15-5.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.6 KiB | [pg_partman_15-5.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.8 KiB | [pg_partman_15-5.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.9 KiB | [pg_partman_15-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.2 KiB | [pg_partman_15-5.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.9 KiB | [pg_partman_15-4.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.3 KiB | [pg_partman_15-4.7.3-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.3-3.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.5 KiB | [pg_partman_15-4.7.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.3-1.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.3 KiB | [pg_partman_15-4.7.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.2-1.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_15-4.7.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.1-1.rhel9.x86_64.rpm) |
| `pg_partman_15` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.1 KiB | [pg_partman_15-4.7.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-4.7.0-2.rhel9.x86_64.rpm) |
| `pg_partman_15` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.4 KiB | [pg_partman_15-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.5 KiB | [pg_partman_15-5.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.9 KiB | [pg_partman_15-5.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.1 KiB | [pg_partman_15-5.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.7 KiB | [pg_partman_15-5.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.4 KiB | [pg_partman_15-5.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.5 KiB | [pg_partman_15-5.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.8 KiB | [pg_partman_15-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.1 KiB | [pg_partman_15-5.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `4.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.4 KiB | [pg_partman_15-4.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-4.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_15` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.1 KiB | [pg_partman_15-4.7.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-4.7.3-3.rhel9.aarch64.rpm) |
| `pg_partman_15` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.1 KiB | [pg_partman_15-4.7.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-4.7.3-1.rhel9.aarch64.rpm) |
| `pg_partman_15` | `4.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.8 KiB | [pg_partman_15-4.7.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-4.7.2-1.rhel9.aarch64.rpm) |
| `pg_partman_15` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.8 KiB | [pg_partman_15-4.7.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-4.7.1-1.rhel9.aarch64.rpm) |
| `pg_partman_15` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_15-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_15-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_15` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_15-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_15` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_15-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-partman` | `5.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 232.6 KiB | [postgresql-15-partman_5.3.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg12+1_amd64.deb) |
| `postgresql-15-partman` | `5.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.6 KiB | [postgresql-15-partman_5.3.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg12+1_arm64.deb) |
| `postgresql-15-partman` | `5.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 232.7 KiB | [postgresql-15-partman_5.3.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg13+1_amd64.deb) |
| `postgresql-15-partman` | `5.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.6 KiB | [postgresql-15-partman_5.3.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg13+1_arm64.deb) |
| `postgresql-15-partman` | `5.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 230.9 KiB | [postgresql-15-partman_5.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.4 KiB | [postgresql-15-partman_5.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 225.8 KiB | [postgresql-15-partman_5.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.7 KiB | [postgresql-15-partman_5.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_14` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.3 KiB | [pg_partman_14-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_14-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 261.4 KiB | [pg_partman_14-5.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.9 KiB | [pg_partman_14-5.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.1 KiB | [pg_partman_14-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.6 KiB | [pg_partman_14-5.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.3 KiB | [pg_partman_14-5.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 254.7 KiB | [pg_partman_14-5.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 249.4 KiB | [pg_partman_14-5.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 248.4 KiB | [pg_partman_14-5.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.9 KiB | [pg_partman_14-4.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.5 KiB | [pg_partman_14-4.7.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.3-3.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.2 KiB | [pg_partman_14-4.7.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.3-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 245.7 KiB | [pg_partman_14-4.7.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.2-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.6 KiB | [pg_partman_14-4.7.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.1-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.9 KiB | [pg_partman_14-4.7.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.7.0-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 256.2 KiB | [pg_partman_14-4.6.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.6.2-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 255.7 KiB | [pg_partman_14-4.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.6.1-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 252.2 KiB | [pg_partman_14-4.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.6.0-1.rhel8.x86_64.rpm) |
| `pg_partman_14` | `4.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.5 KiB | [pg_partman_14-4.5.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-4.5.1-2.rhel8.x86_64.rpm) |
| `pg_partman_14` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_14-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_14-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 261.3 KiB | [pg_partman_14-5.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.8 KiB | [pg_partman_14-5.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_14-5.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.5 KiB | [pg_partman_14-5.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.2 KiB | [pg_partman_14-5.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 254.7 KiB | [pg_partman_14-5.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 249.3 KiB | [pg_partman_14-5.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 248.4 KiB | [pg_partman_14-5.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `4.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.8 KiB | [pg_partman_14-4.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-4.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_14` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.4 KiB | [pg_partman_14-4.7.3-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-4.7.3-3.rhel8.aarch64.rpm) |
| `pg_partman_14` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.1 KiB | [pg_partman_14-4.7.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-4.7.3-1.rhel8.aarch64.rpm) |
| `pg_partman_14` | `4.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 245.6 KiB | [pg_partman_14-4.7.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-4.7.2-1.rhel8.aarch64.rpm) |
| `pg_partman_14` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_14-4.7.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-4.7.1-1.rhel8.aarch64.rpm) |
| `pg_partman_14` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.7 KiB | [pg_partman_14-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 212.9 KiB | [pg_partman_14-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 207.3 KiB | [pg_partman_14-5.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.9 KiB | [pg_partman_14-5.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.2 KiB | [pg_partman_14-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.9 KiB | [pg_partman_14-5.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.6 KiB | [pg_partman_14-5.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_partman_14-5.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.9 KiB | [pg_partman_14-5.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.2 KiB | [pg_partman_14-5.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.7 KiB | [pg_partman_14-4.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.5 KiB | [pg_partman_14-4.7.3-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.3-3.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.4 KiB | [pg_partman_14-4.7.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.3-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.1 KiB | [pg_partman_14-4.7.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.2-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_14-4.7.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.1-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.1 KiB | [pg_partman_14-4.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.7.0-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 211.1 KiB | [pg_partman_14-4.6.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.6.2-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `4.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 210.6 KiB | [pg_partman_14-4.6.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-4.6.1-1.rhel9.x86_64.rpm) |
| `pg_partman_14` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.2 KiB | [pg_partman_14-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.6 KiB | [pg_partman_14-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.4 KiB | [pg_partman_14-5.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.8 KiB | [pg_partman_14-5.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.2 KiB | [pg_partman_14-5.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.7 KiB | [pg_partman_14-5.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.4 KiB | [pg_partman_14-5.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.5 KiB | [pg_partman_14-5.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.9 KiB | [pg_partman_14-5.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.2 KiB | [pg_partman_14-5.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `4.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.3 KiB | [pg_partman_14-4.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-4.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_14` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.1 KiB | [pg_partman_14-4.7.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-4.7.3-3.rhel9.aarch64.rpm) |
| `pg_partman_14` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.0 KiB | [pg_partman_14-4.7.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-4.7.3-1.rhel9.aarch64.rpm) |
| `pg_partman_14` | `4.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.7 KiB | [pg_partman_14-4.7.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-4.7.2-1.rhel9.aarch64.rpm) |
| `pg_partman_14` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.8 KiB | [pg_partman_14-4.7.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-4.7.1-1.rhel9.aarch64.rpm) |
| `pg_partman_14` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.3 KiB | [pg_partman_14-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.7 KiB | [pg_partman_14-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_14-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_14-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_14` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_14-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_14` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.8 KiB | [pg_partman_14-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-partman` | `5.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 232.6 KiB | [postgresql-14-partman_5.3.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg12+1_amd64.deb) |
| `postgresql-14-partman` | `5.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.5 KiB | [postgresql-14-partman_5.3.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg12+1_arm64.deb) |
| `postgresql-14-partman` | `5.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 232.6 KiB | [postgresql-14-partman_5.3.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg13+1_amd64.deb) |
| `postgresql-14-partman` | `5.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.5 KiB | [postgresql-14-partman_5.3.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg13+1_arm64.deb) |
| `postgresql-14-partman` | `5.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 229.5 KiB | [postgresql-14-partman_5.3.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 229.0 KiB | [postgresql-14-partman_5.3.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 225.9 KiB | [postgresql-14-partman_5.3.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.7 KiB | [postgresql-14-partman_5.3.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.3.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_13` | `4.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.8 KiB | [pg_partman_13-4.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.4 KiB | [pg_partman_13-4.7.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.3-3.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.1 KiB | [pg_partman_13-4.7.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.3-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 245.6 KiB | [pg_partman_13-4.7.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.2-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_13-4.7.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.1-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.3 KiB | [pg_partman_13-4.7.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.7.0-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 255.6 KiB | [pg_partman_13-4.6.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.6.2-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 255.5 KiB | [pg_partman_13-4.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.6.1-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 251.8 KiB | [pg_partman_13-4.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.6.0-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 246.0 KiB | [pg_partman_13-4.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.5.1-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 245.3 KiB | [pg_partman_13-4.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.5.0-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 242.1 KiB | [pg_partman_13-4.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_partman_13-4.4.1-1.rhel8.x86_64.rpm) |
| `pg_partman_13` | `4.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.7 KiB | [pg_partman_13-4.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_partman_13-4.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_13` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.4 KiB | [pg_partman_13-4.7.3-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_partman_13-4.7.3-3.rhel8.aarch64.rpm) |
| `pg_partman_13` | `4.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 246.1 KiB | [pg_partman_13-4.7.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_partman_13-4.7.3-1.rhel8.aarch64.rpm) |
| `pg_partman_13` | `4.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 245.6 KiB | [pg_partman_13-4.7.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_partman_13-4.7.2-1.rhel8.aarch64.rpm) |
| `pg_partman_13` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.8 KiB | [pg_partman_13-4.7.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_partman_13-4.7.1-1.rhel8.aarch64.rpm) |
| `pg_partman_13` | `4.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_partman_13-4.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.4 KiB | [pg_partman_13-4.7.3-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.3-3.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.4 KiB | [pg_partman_13-4.7.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.3-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.1 KiB | [pg_partman_13-4.7.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.2-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.0 KiB | [pg_partman_13-4.7.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.1-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 212.5 KiB | [pg_partman_13-4.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.7.0-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 210.8 KiB | [pg_partman_13-4.6.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.6.2-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 210.4 KiB | [pg_partman_13-4.6.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_partman_13-4.6.1-1.rhel9.x86_64.rpm) |
| `pg_partman_13` | `4.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.5 KiB | [pg_partman_13-4.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_partman_13-4.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_13` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.0 KiB | [pg_partman_13-4.7.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_partman_13-4.7.3-3.rhel9.aarch64.rpm) |
| `pg_partman_13` | `4.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.1 KiB | [pg_partman_13-4.7.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_partman_13-4.7.3-1.rhel9.aarch64.rpm) |
| `pg_partman_13` | `4.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.9 KiB | [pg_partman_13-4.7.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_partman_13-4.7.2-1.rhel9.aarch64.rpm) |
| `pg_partman_13` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.8 KiB | [pg_partman_13-4.7.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_partman_13-4.7.1-1.rhel9.aarch64.rpm) |
| `postgresql-13-partman` | `4.7.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 251.3 KiB | [postgresql-13-partman_4.7.4-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg120+1_amd64.deb) |
| `postgresql-13-partman` | `4.7.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 250.8 KiB | [postgresql-13-partman_4.7.4-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg120+1_arm64.deb) |
| `postgresql-13-partman` | `4.7.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 251.3 KiB | [postgresql-13-partman_4.7.4-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg130+1_amd64.deb) |
| `postgresql-13-partman` | `4.7.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 250.9 KiB | [postgresql-13-partman_4.7.4-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg130+1_arm64.deb) |
| `postgresql-13-partman` | `4.7.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 215.8 KiB | [postgresql-13-partman_4.7.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-partman` | `4.7.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 215.0 KiB | [postgresql-13-partman_4.7.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-13-partman_4.7.4-2.pgdg22.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpartman/pg_partman" title="Repository" icon="github" subtitle="github.com/pgpartman/pg_partman" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_partman-5.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_partman;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_partman;		# install via package name, for the active PG version

pig install pg_partman -v 18;   # install for PG 18
pig install pg_partman -v 17;   # install for PG 17
pig install pg_partman -v 16;   # install for PG 16
pig install pg_partman -v 15;   # install for PG 15
pig install pg_partman -v 14;   # install for PG 14
pig install pg_partman -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_partman;
```
