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
| **2510** | {{< badge content="pg_partman" link="https://github.com/pgpartman/pg_partman" >}} | {{< ext "pg_partman" >}} | `5.4.3` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "timeseries" >}} |
|   **See Also**    | {{< ext "citus" >}} {{< ext "pg_fkpart" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "emaj" >}} {{< ext "pg_cron" >}} {{< ext "plproxy" >}} {{< ext "temporal_tables" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.4.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_partman` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.4.3` | {{< bg "18" "pg_partman_18" "green" >}} {{< bg "17" "pg_partman_17" "green" >}} {{< bg "16" "pg_partman_16" "green" >}} {{< bg "15" "pg_partman_15" "green" >}} {{< bg "14" "pg_partman_14" "green" >}} | `pg_partman_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.4.3` | {{< bg "18" "postgresql-18-partman" "green" >}} {{< bg "17" "postgresql-17-partman" "green" >}} {{< bg "16" "postgresql-16-partman" "green" >}} {{< bg "15" "postgresql-15-partman" "green" >}} {{< bg "14" "postgresql-14-partman" "green" >}} | `postgresql-$v-partman` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 24" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 19" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 19" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 22" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 16" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 19" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 19" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 5.4.3" "pg_partman_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.4.3" "pg_partman_14 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 5.4.3" "postgresql-18-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-17-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-16-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-15-partman : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.4.3" "postgresql-14-partman : AVAIL 2" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_18` | `5.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.6 KiB | [pg_partman_18-5.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_18` | `5.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.0 KiB | [pg_partman_18-5.4.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.4.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_18` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 278.6 KiB | [pg_partman_18-5.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_18` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 275.3 KiB | [pg_partman_18-5.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.3 KiB | [pg_partman_18-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_18-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 262.2 KiB | [pg_partman_18-5.2.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_partman_18-5.2.4-2PGDG.rhel8.x86_64.rpm) |
| `pg_partman_18` | `5.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.5 KiB | [pg_partman_18-5.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_18` | `5.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.0 KiB | [pg_partman_18-5.4.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.4.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_18` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.6 KiB | [pg_partman_18-5.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_18` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.3 KiB | [pg_partman_18-5.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_18` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_18-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 262.2 KiB | [pg_partman_18-5.2.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_partman_18-5.2.4-2PGDG.rhel8.aarch64.rpm) |
| `pg_partman_18` | `5.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.6 KiB | [pg_partman_18-5.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_18` | `5.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.1 KiB | [pg_partman_18-5.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_18` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 217.8 KiB | [pg_partman_18-5.4.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.4.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_18` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.1 KiB | [pg_partman_18-5.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_18-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.0 KiB | [pg_partman_18-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 208.0 KiB | [pg_partman_18-5.2.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_partman_18-5.2.4-2PGDG.rhel9.x86_64.rpm) |
| `pg_partman_18` | `5.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [pg_partman_18-5.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_18` | `5.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.0 KiB | [pg_partman_18-5.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_18` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 217.7 KiB | [pg_partman_18-5.4.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.4.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_18` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.2 KiB | [pg_partman_18-5.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_18` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.1 KiB | [pg_partman_18-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.5 KiB | [pg_partman_18-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.6 KiB | [pg_partman_18-5.2.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_partman_18-5.2.4-2PGDG.rhel9.aarch64.rpm) |
| `pg_partman_18` | `5.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.8 KiB | [pg_partman_18-5.4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_18` | `5.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.4 KiB | [pg_partman_18-5.4.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.4.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_18` | `5.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.1 KiB | [pg_partman_18-5.4.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.4.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_18` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 218.5 KiB | [pg_partman_18-5.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_18` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.5 KiB | [pg_partman_18-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_18-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_partman_18-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_18` | `5.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.8 KiB | [pg_partman_18-5.4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_18` | `5.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.3 KiB | [pg_partman_18-5.4.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.4.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_18` | `5.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.0 KiB | [pg_partman_18-5.4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_18` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.4 KiB | [pg_partman_18-5.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_18` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_18-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_18` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.4 KiB | [pg_partman_18-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_18` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_18-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_partman_18-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-partman` | `5.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 238.4 KiB | [postgresql-18-partman_5.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 237.9 KiB | [postgresql-18-partman_5.4.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 238.2 KiB | [postgresql-18-partman_5.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.8 KiB | [postgresql-18-partman_5.4.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 238.3 KiB | [postgresql-18-partman_5.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 237.9 KiB | [postgresql-18-partman_5.4.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 238.2 KiB | [postgresql-18-partman_5.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.9 KiB | [postgresql-18-partman_5.4.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 231.4 KiB | [postgresql-18-partman_5.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 231.2 KiB | [postgresql-18-partman_5.4.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.9 KiB | [postgresql-18-partman_5.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.7 KiB | [postgresql-18-partman_5.4.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.6 KiB | [postgresql-18-partman_5.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.5 KiB | [postgresql-18-partman_5.4.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.4 KiB | [postgresql-18-partman_5.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.2 KiB | [postgresql-18-partman_5.4.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.3 KiB | [postgresql-18-partman_5.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.7 KiB | [postgresql-18-partman_5.4.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-partman` | `5.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.0 KiB | [postgresql-18-partman_5.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-partman` | `5.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.4 KiB | [postgresql-18-partman_5.4.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-18-partman_5.4.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_17` | `5.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.5 KiB | [pg_partman_17-5.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_17` | `5.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.0 KiB | [pg_partman_17-5.4.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.4.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_17` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 278.6 KiB | [pg_partman_17-5.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_17` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 275.3 KiB | [pg_partman_17-5.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 271.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 270.4 KiB | [pg_partman_17-5.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 261.4 KiB | [pg_partman_17-5.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.0 KiB | [pg_partman_17-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.6 KiB | [pg_partman_17-5.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 259.3 KiB | [pg_partman_17-5.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 254.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_partman_17-5.1.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_partman_17` | `5.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.5 KiB | [pg_partman_17-5.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_17` | `5.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.0 KiB | [pg_partman_17-5.4.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.4.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_17` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.6 KiB | [pg_partman_17-5.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_17` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.3 KiB | [pg_partman_17-5.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_17` | `5.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 271.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 270.3 KiB | [pg_partman_17-5.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 261.3 KiB | [pg_partman_17-5.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 260.0 KiB | [pg_partman_17-5.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.6 KiB | [pg_partman_17-5.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 259.2 KiB | [pg_partman_17-5.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 254.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_partman_17-5.1.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_partman_17` | `5.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.5 KiB | [pg_partman_17-5.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_17` | `5.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.0 KiB | [pg_partman_17-5.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_17` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 217.8 KiB | [pg_partman_17-5.4.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.4.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_17` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.1 KiB | [pg_partman_17-5.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 213.6 KiB | [pg_partman_17-5.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 212.8 KiB | [pg_partman_17-5.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 207.4 KiB | [pg_partman_17-5.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.8 KiB | [pg_partman_17-5.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.2 KiB | [pg_partman_17-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.9 KiB | [pg_partman_17-5.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 205.5 KiB | [pg_partman_17-5.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.9 KiB | [pg_partman_17-5.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_partman_17-5.1.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_partman_17` | `5.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [pg_partman_17-5.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_17` | `5.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.0 KiB | [pg_partman_17-5.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_17` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 217.7 KiB | [pg_partman_17-5.4.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.4.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_17` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.1 KiB | [pg_partman_17-5.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_17` | `5.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.4 KiB | [pg_partman_17-5.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.5 KiB | [pg_partman_17-5.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 207.0 KiB | [pg_partman_17-5.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 206.3 KiB | [pg_partman_17-5.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.8 KiB | [pg_partman_17-5.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 205.4 KiB | [pg_partman_17-5.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.8 KiB | [pg_partman_17-5.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_partman_17-5.1.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_partman_17` | `5.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.8 KiB | [pg_partman_17-5.4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_17` | `5.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.2 KiB | [pg_partman_17-5.4.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.4.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_17` | `5.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.1 KiB | [pg_partman_17-5.4.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.4.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_17` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 218.5 KiB | [pg_partman_17-5.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_17` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_17-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.4 KiB | [pg_partman_17-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_partman_17-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_17` | `5.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.8 KiB | [pg_partman_17-5.4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_17` | `5.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.3 KiB | [pg_partman_17-5.4.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.4.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_17` | `5.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.0 KiB | [pg_partman_17-5.4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_17` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.3 KiB | [pg_partman_17-5.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_17` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_17-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_17` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_17-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_17` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_17-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_partman_17-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-partman` | `5.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 238.2 KiB | [postgresql-17-partman_5.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 237.8 KiB | [postgresql-17-partman_5.4.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 238.2 KiB | [postgresql-17-partman_5.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.7 KiB | [postgresql-17-partman_5.4.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 238.3 KiB | [postgresql-17-partman_5.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 237.8 KiB | [postgresql-17-partman_5.4.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 238.1 KiB | [postgresql-17-partman_5.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.7 KiB | [postgresql-17-partman_5.4.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.9 KiB | [postgresql-17-partman_5.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.8 KiB | [postgresql-17-partman_5.4.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 235.4 KiB | [postgresql-17-partman_5.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 235.3 KiB | [postgresql-17-partman_5.4.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.5 KiB | [postgresql-17-partman_5.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.4 KiB | [postgresql-17-partman_5.4.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.3 KiB | [postgresql-17-partman_5.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.1 KiB | [postgresql-17-partman_5.4.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.2 KiB | [postgresql-17-partman_5.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.6 KiB | [postgresql-17-partman_5.4.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-partman` | `5.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.9 KiB | [postgresql-17-partman_5.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-partman` | `5.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.3 KiB | [postgresql-17-partman_5.4.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-17-partman_5.4.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_16` | `5.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.5 KiB | [pg_partman_16-5.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_16` | `5.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.0 KiB | [pg_partman_16-5.4.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.4.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_16` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 278.6 KiB | [pg_partman_16-5.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_16` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 275.3 KiB | [pg_partman_16-5.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_partman_16-5.4.0-1PGDG.rhel8.10.x86_64.rpm) |
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
| `pg_partman_16` | `5.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.5 KiB | [pg_partman_16-5.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_16` | `5.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.9 KiB | [pg_partman_16-5.4.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.4.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_16` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.6 KiB | [pg_partman_16-5.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_16` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.2 KiB | [pg_partman_16-5.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_partman_16-5.4.0-1PGDG.rhel8.10.aarch64.rpm) |
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
| `pg_partman_16` | `5.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.5 KiB | [pg_partman_16-5.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_16` | `5.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.0 KiB | [pg_partman_16-5.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_16` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 217.6 KiB | [pg_partman_16-5.4.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.4.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_16` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.1 KiB | [pg_partman_16-5.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_partman_16-5.4.0-1PGDG.rhel9.7.x86_64.rpm) |
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
| `pg_partman_16` | `5.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.5 KiB | [pg_partman_16-5.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_16` | `5.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.0 KiB | [pg_partman_16-5.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_16` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 217.7 KiB | [pg_partman_16-5.4.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.4.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_16` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.1 KiB | [pg_partman_16-5.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_partman_16-5.4.0-1PGDG.rhel9.7.aarch64.rpm) |
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
| `pg_partman_16` | `5.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.9 KiB | [pg_partman_16-5.4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_16` | `5.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.4 KiB | [pg_partman_16-5.4.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.4.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_16` | `5.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.1 KiB | [pg_partman_16-5.4.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.4.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_16` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 218.4 KiB | [pg_partman_16-5.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_16` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.3 KiB | [pg_partman_16-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_16-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_16-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_partman_16-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_16` | `5.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.8 KiB | [pg_partman_16-5.4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_16` | `5.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.2 KiB | [pg_partman_16-5.4.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.4.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_16` | `5.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.0 KiB | [pg_partman_16-5.4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_16` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.4 KiB | [pg_partman_16-5.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_16` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_16-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_16` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_16-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_16` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.6 KiB | [pg_partman_16-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_partman_16-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-partman` | `5.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 238.1 KiB | [postgresql-16-partman_5.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 237.8 KiB | [postgresql-16-partman_5.4.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 238.1 KiB | [postgresql-16-partman_5.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.7 KiB | [postgresql-16-partman_5.4.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 238.2 KiB | [postgresql-16-partman_5.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 237.9 KiB | [postgresql-16-partman_5.4.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 238.1 KiB | [postgresql-16-partman_5.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.8 KiB | [postgresql-16-partman_5.4.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.5 KiB | [postgresql-16-partman_5.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.2 KiB | [postgresql-16-partman_5.4.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 235.0 KiB | [postgresql-16-partman_5.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 234.8 KiB | [postgresql-16-partman_5.4.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.5 KiB | [postgresql-16-partman_5.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.3 KiB | [postgresql-16-partman_5.4.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.3 KiB | [postgresql-16-partman_5.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.1 KiB | [postgresql-16-partman_5.4.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.2 KiB | [postgresql-16-partman_5.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.6 KiB | [postgresql-16-partman_5.4.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-partman` | `5.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.9 KiB | [postgresql-16-partman_5.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-partman` | `5.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.3 KiB | [postgresql-16-partman_5.4.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-16-partman_5.4.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_15` | `5.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.5 KiB | [pg_partman_15-5.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_15` | `5.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.0 KiB | [pg_partman_15-5.4.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.4.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_15` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 278.6 KiB | [pg_partman_15-5.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_15` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 275.3 KiB | [pg_partman_15-5.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_partman_15-5.4.0-1PGDG.rhel8.10.x86_64.rpm) |
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
| `pg_partman_15` | `5.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.5 KiB | [pg_partman_15-5.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_15` | `5.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.9 KiB | [pg_partman_15-5.4.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.4.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_15` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.6 KiB | [pg_partman_15-5.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_15` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.2 KiB | [pg_partman_15-5.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_partman_15-5.4.0-1PGDG.rhel8.10.aarch64.rpm) |
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
| `pg_partman_15` | `5.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.5 KiB | [pg_partman_15-5.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_15` | `5.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.0 KiB | [pg_partman_15-5.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_15` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 217.7 KiB | [pg_partman_15-5.4.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.4.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_15` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.1 KiB | [pg_partman_15-5.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_partman_15-5.4.0-1PGDG.rhel9.7.x86_64.rpm) |
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
| `pg_partman_15` | `5.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [pg_partman_15-5.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_15` | `5.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.0 KiB | [pg_partman_15-5.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_15` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 217.7 KiB | [pg_partman_15-5.4.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.4.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_15` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.2 KiB | [pg_partman_15-5.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_partman_15-5.4.0-1PGDG.rhel9.7.aarch64.rpm) |
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
| `pg_partman_15` | `5.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.8 KiB | [pg_partman_15-5.4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_15` | `5.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.4 KiB | [pg_partman_15-5.4.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.4.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_15` | `5.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.0 KiB | [pg_partman_15-5.4.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.4.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_15` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 218.5 KiB | [pg_partman_15-5.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_15` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.6 KiB | [pg_partman_15-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_15-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_partman_15-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_15` | `5.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.8 KiB | [pg_partman_15-5.4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_15` | `5.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.3 KiB | [pg_partman_15-5.4.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.4.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_15` | `5.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 219.9 KiB | [pg_partman_15-5.4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_15` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.4 KiB | [pg_partman_15-5.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_15` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_15-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_15` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_15-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_15` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.7 KiB | [pg_partman_15-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_partman_15-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-partman` | `5.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 238.2 KiB | [postgresql-15-partman_5.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 237.8 KiB | [postgresql-15-partman_5.4.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.9 KiB | [postgresql-15-partman_5.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.7 KiB | [postgresql-15-partman_5.4.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 238.2 KiB | [postgresql-15-partman_5.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 237.8 KiB | [postgresql-15-partman_5.4.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 238.1 KiB | [postgresql-15-partman_5.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.7 KiB | [postgresql-15-partman_5.4.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.5 KiB | [postgresql-15-partman_5.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.3 KiB | [postgresql-15-partman_5.4.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 235.0 KiB | [postgresql-15-partman_5.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 234.9 KiB | [postgresql-15-partman_5.4.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.5 KiB | [postgresql-15-partman_5.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.2 KiB | [postgresql-15-partman_5.4.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.3 KiB | [postgresql-15-partman_5.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.0 KiB | [postgresql-15-partman_5.4.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.2 KiB | [postgresql-15-partman_5.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.5 KiB | [postgresql-15-partman_5.4.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-partman` | `5.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.9 KiB | [postgresql-15-partman_5.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-partman` | `5.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.3 KiB | [postgresql-15-partman_5.4.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-15-partman_5.4.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_partman_14` | `5.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.6 KiB | [pg_partman_14-5.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_14` | `5.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 279.0 KiB | [pg_partman_14-5.4.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.4.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_14` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 278.6 KiB | [pg_partman_14-5.4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_partman_14` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 275.3 KiB | [pg_partman_14-5.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_partman_14-5.4.0-1PGDG.rhel8.10.x86_64.rpm) |
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
| `pg_partman_14` | `5.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.5 KiB | [pg_partman_14-5.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_14` | `5.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.0 KiB | [pg_partman_14-5.4.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.4.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_14` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.6 KiB | [pg_partman_14-5.4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_partman_14` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.3 KiB | [pg_partman_14-5.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_partman_14-5.4.0-1PGDG.rhel8.10.aarch64.rpm) |
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
| `pg_partman_14` | `5.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.5 KiB | [pg_partman_14-5.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_14` | `5.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 218.0 KiB | [pg_partman_14-5.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_14` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 217.8 KiB | [pg_partman_14-5.4.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.4.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_partman_14` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.1 KiB | [pg_partman_14-5.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_partman_14-5.4.0-1PGDG.rhel9.7.x86_64.rpm) |
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
| `pg_partman_14` | `5.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [pg_partman_14-5.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_14` | `5.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.0 KiB | [pg_partman_14-5.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_14` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 217.6 KiB | [pg_partman_14-5.4.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.4.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_partman_14` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.1 KiB | [pg_partman_14-5.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_partman_14-5.4.0-1PGDG.rhel9.7.aarch64.rpm) |
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
| `pg_partman_14` | `5.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.9 KiB | [pg_partman_14-5.4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_14` | `5.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.4 KiB | [pg_partman_14-5.4.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.4.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_14` | `5.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 220.1 KiB | [pg_partman_14-5.4.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.4.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_14` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 218.5 KiB | [pg_partman_14-5.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_partman_14` | `5.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.3 KiB | [pg_partman_14-5.3.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.3.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 215.7 KiB | [pg_partman_14-5.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 210.5 KiB | [pg_partman_14-5.2.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_partman_14-5.2.4-2PGDG.rhel10.x86_64.rpm) |
| `pg_partman_14` | `5.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.8 KiB | [pg_partman_14-5.4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_14` | `5.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.4 KiB | [pg_partman_14-5.4.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.4.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_14` | `5.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 220.0 KiB | [pg_partman_14-5.4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_14` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.5 KiB | [pg_partman_14-5.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_partman_14` | `5.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 216.2 KiB | [pg_partman_14-5.3.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.3.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_14` | `5.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.5 KiB | [pg_partman_14-5.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.3.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_partman_14` | `5.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 210.8 KiB | [pg_partman_14-5.2.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_partman_14-5.2.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-partman` | `5.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 238.1 KiB | [postgresql-14-partman_5.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 237.8 KiB | [postgresql-14-partman_5.4.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 238.1 KiB | [postgresql-14-partman_5.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 237.7 KiB | [postgresql-14-partman_5.4.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 238.2 KiB | [postgresql-14-partman_5.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 237.8 KiB | [postgresql-14-partman_5.4.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 238.1 KiB | [postgresql-14-partman_5.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.7 KiB | [postgresql-14-partman_5.4.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 234.2 KiB | [postgresql-14-partman_5.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 233.9 KiB | [postgresql-14-partman_5.4.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 233.7 KiB | [postgresql-14-partman_5.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 233.4 KiB | [postgresql-14-partman_5.4.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.5 KiB | [postgresql-14-partman_5.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.3 KiB | [postgresql-14-partman_5.4.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.3 KiB | [postgresql-14-partman_5.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.0 KiB | [postgresql-14-partman_5.4.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.1 KiB | [postgresql-14-partman_5.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 230.6 KiB | [postgresql-14-partman_5.4.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-partman` | `5.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.9 KiB | [postgresql-14-partman_5.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-partman` | `5.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 230.3 KiB | [postgresql-14-partman_5.4.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-partman/postgresql-14-partman_5.4.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgpartman/pg_partman" title="Repository" icon="github" subtitle="github.com/pgpartman/pg_partman" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_partman-5.4.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_partman;		# build rpm/deb
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_partman CASCADE; -- requires plpgsql
```



## Usage

> [pg_partman: Extension to manage partitioned tables by time or ID](https://github.com/pgpartman/pg_partman)

`pg_partman` automates creation and management of both time-based and number-based partition sets
using PostgreSQL's native declarative partitioning (v5.0+). It handles adding new partitions and
removing old ones per retention policies, with an optional background worker for automatic maintenance.

### Create the Extension

```sql
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman SCHEMA partman;
```

### Create a Time-Based Partition Set

```sql
CREATE TABLE public.measurements (
    id          bigserial,
    created_at  timestamptz NOT NULL DEFAULT now(),
    value       numeric
) PARTITION BY RANGE (created_at);

SELECT partman.create_parent(
    p_parent_table  := 'public.measurements',
    p_control       := 'created_at',
    p_interval      := '1 day'
);
```

### Create a Serial/ID-Based Partition Set

```sql
CREATE TABLE public.events (
    id      bigserial,
    data    text
) PARTITION BY RANGE (id);

SELECT partman.create_parent(
    p_parent_table  := 'public.events',
    p_control       := 'id',
    p_interval      := '100000'
);
```

### Run Maintenance

Manually trigger partition maintenance (create new partitions, drop expired ones):

```sql
SELECT partman.run_maintenance();
```

Or for a specific table:

```sql
SELECT partman.run_maintenance(p_parent_table := 'public.measurements');
```

### Configure Retention

Update the configuration to set retention policy:

```sql
UPDATE partman.part_config
SET    retention = '30 days',
       retention_keep_table = false
WHERE  parent_table = 'public.measurements';
```

### Background Worker

Enable automatic maintenance in `postgresql.conf`:

```
shared_preload_libraries = 'pg_partman_bgw'
pg_partman_bgw.interval = 3600          -- run every hour (seconds)
pg_partman_bgw.dbname = 'mydb'
```

### Migrate Existing Data into Partitions

```sql
CALL partman.partition_data_proc('public.measurements');
```

### Show Partitions

```sql
SELECT * FROM partman.show_partitions('public.measurements');
```

### Undo Partitioning

```sql
CALL partman.undo_partition_proc('public.measurements');
```
