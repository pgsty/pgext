---
title: "pg_jobmon"
linkTitle: "pg_jobmon"
description: "Extension for logging and monitoring functions in PostgreSQL"
weight: 7130
categories: ["SEC"]
width: full
---

Extension for logging and monitoring functions in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7130** | {{< badge content="pg_jobmon" link="https://github.com/omniti-labs/pg_jobmon" >}} | {{< ext "pg_jobmon" >}} | `1.4.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} |
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pgagent" >}} {{< ext "pg_background" >}} {{< ext "logerrors" >}} {{< ext "bgw_replstatus" >}} {{< ext "pgauditlogtofile" >}} {{< ext "pg_auth_mon" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_jobmon" >}} | `1.4.1` | {{< bg "18" "pg_jobmon_18" "green" >}} {{< bg "17" "pg_jobmon_17" "green" >}} {{< bg "16" "pg_jobmon_16" "green" >}} {{< bg "15" "pg_jobmon_15" "green" >}} {{< bg "14" "pg_jobmon_14" "green" >}} {{< bg "13" "pg_jobmon_13" "green" >}} | `pg_jobmon_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_jobmon" >}} | `1.4.1` | {{< bg "18" "postgresql-18-pg-jobmon" "red" >}} {{< bg "17" "postgresql-17-pg-jobmon" "green" >}} {{< bg "16" "postgresql-16-pg-jobmon" "green" >}} {{< bg "15" "postgresql-15-pg-jobmon" "green" >}} {{< bg "14" "postgresql-14-pg-jobmon" "green" >}} {{< bg "13" "postgresql-13-pg-jobmon" "green" >}} | `postgresql-$v-pg-jobmon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_jobmon_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-jobmon : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-jobmon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-jobmon : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-jobmon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-13-pg-jobmon : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_18` | 1.4.1 | `el8.x86_64` | pgdg | 31.6 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_jobmon_18-1.4.1-5PGDG.rhel8.noarch.rpm) |
| `pg_jobmon_18` | 1.4.1 | `el8.aarch64` | pgdg | 31.6 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_jobmon_18-1.4.1-5PGDG.rhel8.noarch.rpm) |
| `pg_jobmon_18` | 1.4.1 | `el9.x86_64` | pgdg | 29.5 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_jobmon_18-1.4.1-5PGDG.rhel9.noarch.rpm) |
| `pg_jobmon_18` | 1.4.1 | `el9.aarch64` | pgdg | 29.4 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_jobmon_18-1.4.1-5PGDG.rhel9.noarch.rpm) |
| `pg_jobmon_18` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_jobmon_18-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_18` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_18-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_jobmon_18-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-18-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-18-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_17` | 1.4.1 | `el8.x86_64` | pgdg | 31.5 KiB | [pg_jobmon_17-1.4.1-4PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_jobmon_17-1.4.1-4PGDG.rhel8.noarch.rpm) |
| `pg_jobmon_17` | 1.4.1 | `el8.aarch64` | pgdg | 31.5 KiB | [pg_jobmon_17-1.4.1-4PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_jobmon_17-1.4.1-4PGDG.rhel8.noarch.rpm) |
| `pg_jobmon_17` | 1.4.1 | `el9.x86_64` | pgdg | 29.5 KiB | [pg_jobmon_17-1.4.1-4PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_jobmon_17-1.4.1-4PGDG.rhel9.noarch.rpm) |
| `pg_jobmon_17` | 1.4.1 | `el9.aarch64` | pgdg | 29.4 KiB | [pg_jobmon_17-1.4.1-4PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_jobmon_17-1.4.1-4PGDG.rhel9.noarch.rpm) |
| `pg_jobmon_17` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_17-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_jobmon_17-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_17` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_17-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_jobmon_17-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-17-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-17-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_16` | 1.4.1 | `el8.x86_64` | pgdg | 31.3 KiB | [pg_jobmon_16-1.4.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_jobmon_16-1.4.1-2.rhel8.noarch.rpm) |
| `pg_jobmon_16` | 1.4.1 | `el8.aarch64` | pgdg | 31.3 KiB | [pg_jobmon_16-1.4.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_jobmon_16-1.4.1-2.rhel8.noarch.rpm) |
| `pg_jobmon_16` | 1.4.1 | `el9.x86_64` | pgdg | 29.5 KiB | [pg_jobmon_16-1.4.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_jobmon_16-1.4.1-2.rhel9.noarch.rpm) |
| `pg_jobmon_16` | 1.4.1 | `el9.aarch64` | pgdg | 29.3 KiB | [pg_jobmon_16-1.4.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_jobmon_16-1.4.1-2.rhel9.noarch.rpm) |
| `pg_jobmon_16` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_16-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_jobmon_16-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_16` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_16-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_jobmon_16-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-16-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-16-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_15` | 1.4.1 | `el8.x86_64` | pgdg | 31.2 KiB | [pg_jobmon_15-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_jobmon_15-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_15` | 1.4.1 | `el8.aarch64` | pgdg | 31.2 KiB | [pg_jobmon_15-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_jobmon_15-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_15` | 1.4.1 | `el9.x86_64` | pgdg | 29.7 KiB | [pg_jobmon_15-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_jobmon_15-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_15` | 1.4.1 | `el9.aarch64` | pgdg | 29.6 KiB | [pg_jobmon_15-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_jobmon_15-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_15` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_15-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_jobmon_15-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_15` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_15-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_jobmon_15-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-15-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-15-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_14` | 1.4.1 | `el8.x86_64` | pgdg | 31.2 KiB | [pg_jobmon_14-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_jobmon_14-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_14` | 1.4.1 | `el8.aarch64` | pgdg | 31.2 KiB | [pg_jobmon_14-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_jobmon_14-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_14` | 1.4.1 | `el9.x86_64` | pgdg | 29.7 KiB | [pg_jobmon_14-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_jobmon_14-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_14` | 1.4.1 | `el9.aarch64` | pgdg | 29.6 KiB | [pg_jobmon_14-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_jobmon_14-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_14` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_14-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_jobmon_14-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_14` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_14-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_jobmon_14-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-14-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-14-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jobmon_13` | 1.4.1 | `el8.x86_64` | pgdg | 31.2 KiB | [pg_jobmon_13-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_jobmon_13-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_13` | 1.4.1 | `el8.aarch64` | pgdg | 31.2 KiB | [pg_jobmon_13-1.4.1-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_jobmon_13-1.4.1-1.rhel8.noarch.rpm) |
| `pg_jobmon_13` | 1.4.1 | `el9.x86_64` | pgdg | 29.7 KiB | [pg_jobmon_13-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_jobmon_13-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_13` | 1.4.1 | `el9.aarch64` | pgdg | 29.6 KiB | [pg_jobmon_13-1.4.1-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_jobmon_13-1.4.1-1.rhel9.noarch.rpm) |
| `pg_jobmon_13` | 1.4.1 | `el10.x86_64` | pgdg | 30.2 KiB | [pg_jobmon_13-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_jobmon_13-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `pg_jobmon_13` | 1.4.1 | `el10.aarch64` | pgdg | 30.2 KiB | [pg_jobmon_13-1.4.1-5PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_jobmon_13-1.4.1-5PGDG.rhel10.noarch.rpm) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `d12.x86_64` | pigsty | 26.1 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `d12.aarch64` | pigsty | 26.1 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `u22.x86_64` | pigsty | 23.3 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `u22.aarch64` | pigsty | 23.3 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `u24.x86_64` | pigsty | 23.2 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-jobmon` | 1.4.1 | `u24.aarch64` | pigsty | 23.2 KiB | [postgresql-13-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jobmon/postgresql-13-pg-jobmon_1.4.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/omniti-labs/pg_jobmon" title="Repository" icon="github" subtitle="github.com/omniti-labs/pg_jobmon" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_jobmon-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_jobmon; # get pg_jobmon source code
pig build dep pg_jobmon; # install build dependencies
pig build pkg pg_jobmon; # build extension rpm or deb
pig build ext pg_jobmon; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_jobmon; # install by extension name, for the current active PG version
pig ext install pg_jobmon; # install via package alias, for the active PG version
pig ext install pg_jobmon -v 18;   # install for PG 18
pig ext install pg_jobmon -v 17;   # install for PG 17
pig ext install pg_jobmon -v 16;   # install for PG 16
pig ext install pg_jobmon -v 15;   # install for PG 15
pig ext install pg_jobmon -v 14;   # install for PG 14
pig ext install pg_jobmon -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_jobmon;
```

