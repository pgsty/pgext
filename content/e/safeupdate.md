---
title: "safeupdate"
linkTitle: "safeupdate"
description: "Require criteria for UPDATE and DELETE"
weight: 5820
categories: ["ADMIN"]
width: full
---

Require criteria for UPDATE and DELETE


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5820** | {{< badge content="safeupdate" link="https://github.com/eradman/pg-safeupdate" >}} | {{< ext "safeupdate" >}} | `1.5` | {{< category "ADMIN" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_upless" >}} {{< ext "pg_savior" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "login_hook" >}} {{< ext "noset" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/safeupdate" >}} | `1.5` | {{< bg "18" "safeupdate_18*" "red" >}} {{< bg "17" "safeupdate_17*" "green" >}} {{< bg "16" "safeupdate_16*" "green" >}} {{< bg "15" "safeupdate_15*" "green" >}} {{< bg "14" "safeupdate_14*" "green" >}} | `safeupdate_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/safeupdate" >}} | `1.5` | {{< bg "18" "postgresql-18-pg-safeupdate" "red" >}} {{< bg "17" "postgresql-17-pg-safeupdate" "green" >}} {{< bg "16" "postgresql-16-pg-safeupdate" "green" >}} {{< bg "15" "postgresql-15-pg-safeupdate" "green" >}} {{< bg "14" "postgresql-14-pg-safeupdate" "green" >}} | `postgresql-$v-pg-safeupdate` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 2" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 3" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 3" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-safeupdate : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_18` | 1.5 | `el8.x86_64` | pgdg | 13.4 KiB | [safeupdate_18-1.5-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/safeupdate_18-1.5-2PGDG.rhel8.x86_64.rpm) |
| `safeupdate_18` | 1.5 | `el8.aarch64` | pgdg | 13.4 KiB | [safeupdate_18-1.5-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/safeupdate_18-1.5-2PGDG.rhel8.aarch64.rpm) |
| `safeupdate_18` | 1.5 | `el9.x86_64` | pgdg | 12.8 KiB | [safeupdate_18-1.5-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/safeupdate_18-1.5-2PGDG.rhel9.x86_64.rpm) |
| `safeupdate_18` | 1.5 | `el9.aarch64` | pgdg | 12.6 KiB | [safeupdate_18-1.5-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/safeupdate_18-1.5-2PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_17` | 1.5 | `el8.x86_64` | pgdg | 13.3 KiB | [safeupdate_17-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/safeupdate_17-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_17` | 1.5 | `el8.aarch64` | pgdg | 13.3 KiB | [safeupdate_17-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/safeupdate_17-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_17` | 1.5 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_17-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/safeupdate_17-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_17` | 1.5 | `el9.aarch64` | pgdg | 12.8 KiB | [safeupdate_17-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/safeupdate_17-1.5-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-safeupdate` | 1.5 | `d12.x86_64` | pigsty | 9.4 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-safeupdate` | 1.5 | `d12.aarch64` | pigsty | 9.6 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-safeupdate` | 1.5 | `u22.x86_64` | pigsty | 9.4 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-safeupdate` | 1.5 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-safeupdate` | 1.5 | `u24.x86_64` | pigsty | 9.1 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-safeupdate` | 1.5 | `u24.aarch64` | pigsty | 9.1 KiB | [postgresql-17-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_16` | 1.5 | `el8.x86_64` | pgdg | 13.3 KiB | [safeupdate_16-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/safeupdate_16-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_16` | 1.4.2 | `el8.x86_64` | pgdg | 13.4 KiB | [safeupdate_16-1.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/safeupdate_16-1.4.2-2PGDG.rhel8.x86_64.rpm) |
| `safeupdate_16` | 1.5 | `el8.aarch64` | pgdg | 13.3 KiB | [safeupdate_16-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/safeupdate_16-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_16` | 1.4.2 | `el8.aarch64` | pgdg | 13.4 KiB | [safeupdate_16-1.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/safeupdate_16-1.4.2-2PGDG.rhel8.aarch64.rpm) |
| `safeupdate_16` | 1.5 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_16-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/safeupdate_16-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_16` | 1.4.2 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_16-1.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/safeupdate_16-1.4.2-2PGDG.rhel9.x86_64.rpm) |
| `safeupdate_16` | 1.5 | `el9.aarch64` | pgdg | 12.8 KiB | [safeupdate_16-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/safeupdate_16-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_16` | 1.4.2 | `el9.aarch64` | pgdg | 12.6 KiB | [safeupdate_16-1.4.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/safeupdate_16-1.4.2-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-safeupdate` | 1.5 | `d12.x86_64` | pigsty | 9.4 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-safeupdate` | 1.5 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-safeupdate` | 1.5 | `u22.x86_64` | pigsty | 9.4 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-safeupdate` | 1.5 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-safeupdate` | 1.5 | `u24.x86_64` | pigsty | 9.1 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-safeupdate` | 1.5 | `u24.aarch64` | pigsty | 9.1 KiB | [postgresql-16-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_15` | 1.5 | `el8.x86_64` | pgdg | 13.3 KiB | [safeupdate_15-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/safeupdate_15-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_15` | 1.4 | `el8.x86_64` | pgdg | 17.8 KiB | [safeupdate_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/safeupdate_15-1.4-1.rhel8.x86_64.rpm) |
| `safeupdate_15` | 1.5 | `el8.aarch64` | pgdg | 13.3 KiB | [safeupdate_15-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_15` | 1.4.2 | `el8.aarch64` | pgdg | 13.3 KiB | [safeupdate_15-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.4.2-1.rhel8.aarch64.rpm) |
| `safeupdate_15` | 1.4 | `el8.aarch64` | pgdg | 17.8 KiB | [safeupdate_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.4-1.rhel8.aarch64.rpm) |
| `safeupdate_15` | 1.5 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_15-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_15` | 1.4.2 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_15-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.4.2-1.rhel9.x86_64.rpm) |
| `safeupdate_15` | 1.4 | `el9.x86_64` | pgdg | 17.9 KiB | [safeupdate_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.4-1.rhel9.x86_64.rpm) |
| `safeupdate_15` | 1.5 | `el9.aarch64` | pgdg | 12.8 KiB | [safeupdate_15-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_15` | 1.4.2 | `el9.aarch64` | pgdg | 12.6 KiB | [safeupdate_15-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.4.2-1.rhel9.aarch64.rpm) |
| `safeupdate_15` | 1.4 | `el9.aarch64` | pgdg | 17.7 KiB | [safeupdate_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.4-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-safeupdate` | 1.5 | `d12.x86_64` | pigsty | 9.4 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-safeupdate` | 1.5 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-safeupdate` | 1.5 | `u22.x86_64` | pigsty | 9.3 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-safeupdate` | 1.5 | `u22.aarch64` | pigsty | 9.5 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-safeupdate` | 1.5 | `u24.x86_64` | pigsty | 9.1 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-safeupdate` | 1.5 | `u24.aarch64` | pigsty | 9.1 KiB | [postgresql-15-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_14` | 1.5 | `el8.x86_64` | pgdg | 13.3 KiB | [safeupdate_14-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/safeupdate_14-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_14` | 1.4 | `el8.x86_64` | pgdg | 18.3 KiB | [safeupdate_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/safeupdate_14-1.4-1.rhel8.x86_64.rpm) |
| `safeupdate_14` | 1.5 | `el8.aarch64` | pgdg | 13.3 KiB | [safeupdate_14-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_14` | 1.4.2 | `el8.aarch64` | pgdg | 13.2 KiB | [safeupdate_14-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.4.2-1.rhel8.aarch64.rpm) |
| `safeupdate_14` | 1.4 | `el8.aarch64` | pgdg | 17.7 KiB | [safeupdate_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.4-1.rhel8.aarch64.rpm) |
| `safeupdate_14` | 1.5 | `el9.x86_64` | pgdg | 12.9 KiB | [safeupdate_14-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/safeupdate_14-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_14` | 1.4.2 | `el9.x86_64` | pgdg | 12.8 KiB | [safeupdate_14-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/safeupdate_14-1.4.2-1.rhel9.x86_64.rpm) |
| `safeupdate_14` | 1.5 | `el9.aarch64` | pgdg | 12.8 KiB | [safeupdate_14-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_14` | 1.4.2 | `el9.aarch64` | pgdg | 12.6 KiB | [safeupdate_14-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.4.2-1.rhel9.aarch64.rpm) |
| `safeupdate_14` | 1.4 | `el9.aarch64` | pgdg | 17.6 KiB | [safeupdate_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.4-1.rhel9.aarch64.rpm) |
| `postgresql-14-pg-safeupdate` | 1.5 | `d12.x86_64` | pigsty | 9.4 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-safeupdate` | 1.5 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-safeupdate` | 1.5 | `u22.x86_64` | pigsty | 9.3 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-safeupdate` | 1.5 | `u22.aarch64` | pigsty | 9.4 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-safeupdate` | 1.5 | `u24.x86_64` | pigsty | 9.0 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-safeupdate` | 1.5 | `u24.aarch64` | pigsty | 9.1 KiB | [postgresql-14-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eradman/pg-safeupdate" title="Repository" icon="github" subtitle="github.com/eradman/pg-safeupdate" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-safeupdate-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get safeupdate; # get safeupdate source code
pig build dep safeupdate; # install build dependencies
pig build pkg safeupdate; # build extension rpm or deb
pig build ext safeupdate; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install safeupdate; # install by extension name, for the current active PG version
pig ext install safeupdate; # install via package alias, for the active PG version
pig ext install safeupdate -v 17;   # install for PG 17
pig ext install safeupdate -v 16;   # install for PG 16
pig ext install safeupdate -v 15;   # install for PG 15
pig ext install safeupdate -v 14;   # install for PG 14
pig ext install safeupdate -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION safeupdate;
```

