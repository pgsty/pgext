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
| **2660** | {{< badge content="pgmq" link="https://github.com/pgmq/pgmq" >}} | {{< ext "pgmq" >}} | `1.11.1` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmq` |
|    **Need By**    | {{< ext "pg_later" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "kafka_fdw" >}} {{< ext "pg_task" >}} {{< ext "pg_net" >}} {{< ext "pg_background" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmq` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "pgmq_18" "green" >}} {{< bg "17" "pgmq_17" "green" >}} {{< bg "16" "pgmq_16" "green" >}} {{< bg "15" "pgmq_15" "green" >}} {{< bg "14" "pgmq_14" "green" >}} | `pgmq_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "postgresql-18-pgmq" "green" >}} {{< bg "17" "postgresql-17-pgmq" "green" >}} {{< bg "16" "postgresql-16-pgmq" "green" >}} {{< bg "15" "postgresql-15-pgmq" "green" >}} {{< bg "14" "postgresql-14-pgmq" "green" >}} | `postgresql-$v-pgmq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "pgmq_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_18` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.7 KiB | [pgmq_18-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_18-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_18` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.6 KiB | [pgmq_18-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_18-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_18` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmq_18-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmq_18-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 40.8 KiB | [pgmq_18-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_18-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_18` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.7 KiB | [pgmq_18-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_18-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_18` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmq_18-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.9 KiB | [pgmq_18-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_18-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_18` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [pgmq_18-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_18-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_18` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_18` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmq_18-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `postgresql-18-pgmq` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.0 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.0 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmq` | `1.11.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.11.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmq` | `1.11.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.4 KiB | [postgresql-18-pgmq_1.11.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.11.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.9 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmq` | `1.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-18-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-18-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_17` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.6 KiB | [pgmq_17-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_17-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_17` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.6 KiB | [pgmq_17-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_17-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_17` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmq_17-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmq_17-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 40.8 KiB | [pgmq_17-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_17-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_17` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.7 KiB | [pgmq_17-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_17-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_17` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmq_17-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.9 KiB | [pgmq_17-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_17-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_17` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [pgmq_17-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_17-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_17` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_17` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmq_17-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `postgresql-17-pgmq` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.5 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.5 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.4 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.4 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.0 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.0 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmq` | `1.11.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-17-pgmq_1.11.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.11.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmq` | `1.11.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.4 KiB | [postgresql-17-pgmq_1.11.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.11.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.9 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmq` | `1.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-17-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-17-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_16` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.7 KiB | [pgmq_16-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_16-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_16` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.6 KiB | [pgmq_16-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_16-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_16` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmq_16-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmq_16-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 40.8 KiB | [pgmq_16-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_16-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_16` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.7 KiB | [pgmq_16-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_16-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_16` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmq_16-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.9 KiB | [pgmq_16-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_16-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_16` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [pgmq_16-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_16-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_16` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_16` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmq_16-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `postgresql-16-pgmq` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.9 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.9 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmq` | `1.11.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.11.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmq` | `1.11.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.4 KiB | [postgresql-16-pgmq_1.11.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.11.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.9 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmq` | `1.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-16-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-16-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_15` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.7 KiB | [pgmq_15-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_15-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_15` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.6 KiB | [pgmq_15-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_15-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_15` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmq_15-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmq_15-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 40.8 KiB | [pgmq_15-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_15-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_15` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.7 KiB | [pgmq_15-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_15-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_15` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmq_15-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.9 KiB | [pgmq_15-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_15-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_15` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [pgmq_15-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_15-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_15` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_15` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmq_15-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `postgresql-15-pgmq` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.9 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.9 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmq` | `1.11.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.11.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmq` | `1.11.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.4 KiB | [postgresql-15-pgmq_1.11.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.11.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.9 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmq` | `1.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-15-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-15-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_14` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.7 KiB | [pgmq_14-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_14-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_14` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.6 KiB | [pgmq_14-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_14-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_14` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmq_14-1.11.0-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmq_14-1.10.1-1PGDG.rhel8.10.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 40.8 KiB | [pgmq_14-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_14-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_14` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.7 KiB | [pgmq_14-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_14-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_14` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.11.0-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.6 KiB | [pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmq_14-1.10.1-1PGDG.rhel9.7.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.9 KiB | [pgmq_14-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_14-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_14` | `1.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.6 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 40.8 KiB | [pgmq_14-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_14-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `pgmq_14` | `1.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.11.0-1PGDG.rhel10.1.noarch.rpm) |
| `pgmq_14` | `1.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmq_14-1.10.1-1PGDG.rhel10.1.noarch.rpm) |
| `postgresql-14-pgmq` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.0 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.0 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmq` | `1.11.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.11.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmq` | `1.11.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.4 KiB | [postgresql-14-pgmq_1.11.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.11.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.9 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmq` | `1.11.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 KiB | [postgresql-14-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmq/postgresql-14-pgmq_1.11.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmq/pgmq" title="Repository" icon="github" subtitle="github.com/pgmq/pgmq" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmq-1.11.1.tar.gz" >}}
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

> [pgmq: A lightweight message queue for PostgreSQL](https://github.com/pgmq/pgmq)

PGMQ is a lightweight message queue built on PostgreSQL, providing guaranteed "exactly once" delivery within a visibility timeout, FIFO queues, topic-based routing, and message archival.

```sql
CREATE EXTENSION pgmq;
```

### Create a Queue

```sql
SELECT pgmq.create('my_queue');
```

### Send Messages

```sql
-- Send a single message (returns msg_id)
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}'
);

-- Send with delay (invisible for 5 seconds)
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}',
  delay      => 5
);

-- Send a batch of messages
SELECT pgmq.send_batch(
  queue_name => 'my_queue',
  msgs       => ARRAY['{"a":1}','{"b":2}','{"c":3}']::jsonb[]
);
```

### Read Messages

Read messages and make them invisible for a visibility timeout (in seconds):

```sql
SELECT * FROM pgmq.read(
  queue_name => 'my_queue',
  vt         => 30,    -- visibility timeout in seconds
  qty        => 2      -- number of messages to read
);
```

### Pop a Message

Read and immediately delete a message:

```sql
SELECT * FROM pgmq.pop('my_queue');
```

### Delete a Message

```sql
SELECT pgmq.delete('my_queue', 6);
```

### Archive a Message

Move a message from the queue to the archive table for long-term retention:

```sql
SELECT pgmq.archive(queue_name => 'my_queue', msg_id => 2);

-- Archive multiple messages
SELECT pgmq.archive(queue_name => 'my_queue', msg_ids => ARRAY[3, 4, 5]);
```

Inspect archived messages:

```sql
SELECT * FROM pgmq.a_my_queue;
```

### Drop a Queue

```sql
SELECT pgmq.drop_queue('my_queue');
```

### Visibility Timeout

Messages become invisible after being read for the duration of the visibility timeout (`vt`). If not deleted or archived within that time, they become visible again for other consumers. Set `vt` greater than the expected processing time.

### Key Functions

| Function | Description |
|----------|-------------|
| `pgmq.create(queue_name)` | Create a new queue |
| `pgmq.send(queue_name, msg, [delay])` | Send a message |
| `pgmq.send_batch(queue_name, msgs)` | Send multiple messages |
| `pgmq.read(queue_name, vt, qty)` | Read messages with visibility timeout |
| `pgmq.pop(queue_name)` | Read and delete a message atomically |
| `pgmq.delete(queue_name, msg_id)` | Delete a message |
| `pgmq.archive(queue_name, msg_id/msg_ids)` | Archive message(s) |
| `pgmq.drop_queue(queue_name)` | Delete a queue |
