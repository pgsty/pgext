---
title: "pgmq"
linkTitle: "pgmq"
description: "A lightweight message queue. Like AWS SQS and RSMQ but on Postgres."
weight: 2660
categories: ["FEAT"]
width: full
---

[**pgmq**](https://github.com/pgmq/pgmq) : A lightweight message queue. Like AWS SQS and RSMQ but on Postgres.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2660** | {{< badge content="pgmq" link="https://github.com/pgmq/pgmq" >}} | {{< ext "pgmq" >}} | `1.12.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmq` |
|    **Need By**    | {{< ext "fsm_core" >}} {{< ext "pg_later" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "kafka_fdw" >}} {{< ext "pg_task" >}} {{< ext "pg_net" >}} {{< ext "pg_background" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.12.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmq` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.12.0` | {{< bg "18" "pgmq_18" "green" >}} {{< bg "17" "pgmq_17" "green" >}} {{< bg "16" "pgmq_16" "green" >}} {{< bg "15" "pgmq_15" "green" >}} {{< bg "14" "pgmq_14" "green" >}} | `pgmq_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.12.0` | {{< bg "18" "postgresql-18-pgmq" "green" >}} {{< bg "17" "postgresql-17-pgmq" "green" >}} {{< bg "16" "postgresql-16-pgmq" "green" >}} {{< bg "15" "postgresql-15-pgmq" "green" >}} {{< bg "14" "postgresql-14-pgmq" "green" >}} | `postgresql-$v-pgmq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 7" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 7" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 7" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_18 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.12.0" "pgmq_14 : AVAIL 7" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_18` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.0 KiB | [pgmq_18-1.12.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_18-1.12.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_18` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [pgmq_18-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmq_18-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pgmq_18-1.12.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_18-1.12.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_18` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pgmq_18-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmq_18-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.0 KiB | [pgmq_18-1.12.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_18-1.12.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_18` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [pgmq_18-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pgmq_18-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_18` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.0 KiB | [pgmq_18-1.12.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_18-1.12.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_18` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgmq_18-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pgmq_18-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_18` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.1 KiB | [pgmq_18-1.12.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_18-1.12.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_18` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pgmq_18-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.8 KiB | [pgmq_18-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_18` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.1 KiB | [pgmq_18-1.12.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_18-1.12.0-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_18` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pgmq_18-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.7 KiB | [pgmq_18-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-18-pgmq` | `1.12.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.6 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.6 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.6 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.6 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.0 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.1 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmq` | `1.12.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.1 KiB | [postgresql-18-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-18-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_17` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.0 KiB | [pgmq_17-1.12.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_17-1.12.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_17` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [pgmq_17-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmq_17-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pgmq_17-1.12.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_17-1.12.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_17` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pgmq_17-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmq_17-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.0 KiB | [pgmq_17-1.12.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_17-1.12.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_17` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [pgmq_17-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pgmq_17-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_17` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.0 KiB | [pgmq_17-1.12.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_17-1.12.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_17` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgmq_17-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pgmq_17-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_17` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.1 KiB | [pgmq_17-1.12.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_17-1.12.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_17` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pgmq_17-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.8 KiB | [pgmq_17-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_17` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.1 KiB | [pgmq_17-1.12.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_17-1.12.0-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_17` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pgmq_17-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.7 KiB | [pgmq_17-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-17-pgmq` | `1.12.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.6 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.6 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.6 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.6 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.0 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.1 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmq` | `1.12.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.1 KiB | [postgresql-17-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-17-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_16` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.0 KiB | [pgmq_16-1.12.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_16-1.12.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_16` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [pgmq_16-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmq_16-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pgmq_16-1.12.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_16-1.12.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_16` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pgmq_16-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmq_16-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.0 KiB | [pgmq_16-1.12.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_16-1.12.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_16` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [pgmq_16-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pgmq_16-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_16` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.0 KiB | [pgmq_16-1.12.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_16-1.12.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_16` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgmq_16-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pgmq_16-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_16` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.1 KiB | [pgmq_16-1.12.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_16-1.12.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_16` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pgmq_16-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.8 KiB | [pgmq_16-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_16` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.1 KiB | [pgmq_16-1.12.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_16-1.12.0-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_16` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pgmq_16-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.7 KiB | [pgmq_16-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-16-pgmq` | `1.12.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.6 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.6 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.6 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.6 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.0 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.1 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmq` | `1.12.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.1 KiB | [postgresql-16-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-16-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_15` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.0 KiB | [pgmq_15-1.12.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_15-1.12.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_15` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [pgmq_15-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmq_15-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pgmq_15-1.12.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_15-1.12.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_15` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pgmq_15-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmq_15-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.0 KiB | [pgmq_15-1.12.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_15-1.12.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_15` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [pgmq_15-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pgmq_15-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_15` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.0 KiB | [pgmq_15-1.12.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_15-1.12.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_15` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgmq_15-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pgmq_15-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_15` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.1 KiB | [pgmq_15-1.12.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_15-1.12.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_15` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pgmq_15-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.8 KiB | [pgmq_15-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_15` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.1 KiB | [pgmq_15-1.12.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_15-1.12.0-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_15` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pgmq_15-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.7 KiB | [pgmq_15-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-15-pgmq` | `1.12.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.6 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.6 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.6 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.6 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.0 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.1 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmq` | `1.12.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.1 KiB | [postgresql-15-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-15-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_14` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.0 KiB | [pgmq_14-1.12.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_14-1.12.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_14` | `1.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.8 KiB | [pgmq_14-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmq_14-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pgmq_14-1.12.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_14-1.12.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_14` | `1.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.8 KiB | [pgmq_14-1.12.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmq_14-1.12.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.0 KiB | [pgmq_14-1.12.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_14-1.12.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_14` | `1.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [pgmq_14-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pgmq_14-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_14` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.0 KiB | [pgmq_14-1.12.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_14-1.12.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_14` | `1.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgmq_14-1.12.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.12.0-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pgmq_14-1.11.1-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.11.1-1PGDG.rhel9.8.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.11.0-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.10.1-1PGDG.rhel9.6.noarch.rpm) |
| `pgmq_14` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 41.1 KiB | [pgmq_14-1.12.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_14-1.12.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_14` | `1.12.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pgmq_14-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.8 KiB | [pgmq_14-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_14` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.1 KiB | [pgmq_14-1.12.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_14-1.12.0-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_14` | `1.12.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pgmq_14-1.12.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.12.0-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.7 KiB | [pgmq_14-1.11.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.11.1-1PGDG.rhel10.2.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.11.0-1PGDG.rhel10.0.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.10.1-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-14-pgmq` | `1.12.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.6 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.6 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.6 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.6 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmq` | `1.12.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.1 KiB | [postgresql-14-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-14-pgmq_1.12.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmq/pgmq" title="Repository" icon="github" subtitle="github.com/pgmq/pgmq" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmq-1.12.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmq;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmq;		# install via package name, for the active PG version

pig install pgmq -v 18;   # install for PG 18
pig install pgmq -v 17;   # install for PG 17
pig install pgmq -v 16;   # install for PG 16
pig install pgmq -v 15;   # install for PG 15
pig install pgmq -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmq;
```

## Usage

Sources:

- [PGMQ v1.12.0 README](https://github.com/pgmq/pgmq/blob/v1.12.0/README.md)
- [PGMQ v1.12.0 release notes](https://github.com/pgmq/pgmq/releases/tag/v1.12.0)
- [Version 1.12.0 migration SQL](https://github.com/pgmq/pgmq/blob/v1.12.0/pgmq-extension/sql/pgmq--1.11.1--1.12.0.sql)
- [PGMQ documentation](https://pgmq.github.io/pgmq/)

pgmq implements durable message queues as PostgreSQL tables and SQL functions. It supports delayed delivery, visibility timeouts, FIFO groups, message headers, polling, topics, and archival. Use it when queue transactions should be coordinated with relational changes in the same database.

### Create a Queue and Send

    CREATE EXTENSION pgmq;
    SELECT pgmq.create('jobs');

    SELECT *
    FROM pgmq.send(
      queue_name => 'jobs',
      msg        => '{"task":"refresh"}'::jsonb,
      delay      => 0
    );

send returns the message identifier. send_batch inserts multiple JSONB messages. Headers can carry routing or tracing metadata separately from the body where the selected overload supports them.

### Read with a Visibility Timeout

    SELECT *
    FROM pgmq.read(
      queue_name => 'jobs',
      vt         => 30,
      qty        => 10
    );

Reading hides each message for vt seconds. On success, delete or archive it:

    SELECT pgmq.delete('jobs', 42);
    SELECT pgmq.archive('jobs', 43);

If processing fails or the consumer disappears, an unacknowledged message becomes visible again. Consumers must therefore be idempotent; pgmq does not make arbitrary application side effects globally exactly once.

pop reads and deletes immediately and is appropriate only when losing a message after the call is acceptable.

### FIFO Group Head Polling

Version 1.12.0 adds a polling read for the head message of multiple FIFO groups:

    SELECT *
    FROM pgmq.read_grouped_head_with_poll(
      queue_name          => 'jobs',
      vt                  => 30,
      qty                 => 10,
      max_poll_seconds    => 5,
      poll_interval_ms    => 100
    );

It selects head-of-group messages while polling for up to max_poll_seconds. This preserves per-group order while allowing different groups to be processed concurrently.

### Queue Administration Index

- pgmq.create(queue_name): create queue and archive structures.
- pgmq.send and pgmq.send_batch: enqueue JSONB messages, optionally delayed.
- pgmq.read: claim visible messages for a visibility timeout.
- pgmq.read_grouped_head_with_poll: poll FIFO group heads.
- pgmq.pop: read and immediately delete.
- pgmq.delete: acknowledge by removing a message.
- pgmq.archive: move a message to the queue archive table.
- pgmq.drop_queue: remove queue objects.
- pgmq.metrics and related helpers: inspect queue depth and age where available.

For queue jobs, archive rows are stored in pgmq.a_<queue_name>. Treat those tables as extension-managed objects.

### Operational Notes

- Set vt longer than normal processing time and design for redelivery after timeouts.
- Queue and archive tables consume ordinary PostgreSQL WAL, storage, vacuum, and backup capacity.
- Archive or delete completed messages and enforce an archive-retention policy.
- Long polling holds a database connection. Size connection pools and polling intervals for the number of consumers.
- Keep queue names within pgmq's identifier rules; call the API rather than constructing table names from untrusted input.
