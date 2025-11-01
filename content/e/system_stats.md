---
title: "system_stats"
linkTitle: "system_stats"
description: "EnterpriseDB system statistics for PostgreSQL"
weight: 6290
categories: ["STAT"]
width: full
---

EnterpriseDB system statistics for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6290** | {{< badge content="system_stats" link="https://github.com/EnterpriseDB/system_stats" >}} | {{< ext "system_stats" >}} | `3.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "pg_proctab" >}} {{< ext "pgmeminfo" >}} {{< ext "pgfincore" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_background" >}} {{< ext "pg_cooldown" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/system_stats" >}} | `3.2` | {{< bg "18" "system_stats_18*" "green" >}} {{< bg "17" "system_stats_17*" "green" >}} {{< bg "16" "system_stats_16*" "green" >}} {{< bg "15" "system_stats_15*" "green" >}} {{< bg "14" "system_stats_14*" "green" >}} {{< bg "13" "system_stats_13*" "green" >}} | `system_stats_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/system_stats" >}} | `3.2` | {{< bg "18" "postgresql-18-system-stats" "green" >}} {{< bg "17" "postgresql-17-system-stats" "green" >}} {{< bg "16" "postgresql-16-system-stats" "green" >}} {{< bg "15" "postgresql-15-system-stats" "green" >}} {{< bg "14" "postgresql-14-system-stats" "green" >}} {{< bg "13" "postgresql-13-system-stats" "green" >}} | `postgresql-$v-system-stats` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 3" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 3" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-system-stats : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-system-stats : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-system-stats : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-system-stats : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-13-system-stats : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_18` | 3.2 | `el8.x86_64` | pgdg | 28.9 KiB | [system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm) |
| `system_stats_18` | 3.2 | `el8.aarch64` | pgdg | 28.7 KiB | [system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm) |
| `system_stats_18` | 3.2 | `el9.x86_64` | pgdg | 28.4 KiB | [system_stats_18-3.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/system_stats_18-3.2-2PGDG.rhel9.x86_64.rpm) |
| `system_stats_18` | 3.2 | `el9.aarch64` | pgdg | 28.3 KiB | [system_stats_18-3.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/system_stats_18-3.2-2PGDG.rhel9.aarch64.rpm) |
| `system_stats_18` | 3.2 | `el10.x86_64` | pgdg | 29.6 KiB | [system_stats_18-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/system_stats_18-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_18` | 3.2 | `el10.aarch64` | pgdg | 28.8 KiB | [system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_17` | 3.2 | `el8.x86_64` | pgdg | 28.8 KiB | [system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_17` | 3.2 | `el8.aarch64` | pgdg | 28.6 KiB | [system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_17` | 3.2 | `el9.x86_64` | pgdg | 28.5 KiB | [system_stats_17-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/system_stats_17-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_17` | 3.2 | `el9.aarch64` | pgdg | 28.5 KiB | [system_stats_17-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/system_stats_17-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_17` | 3.2 | `el10.x86_64` | pgdg | 29.6 KiB | [system_stats_17-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/system_stats_17-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_17` | 3.2 | `el10.aarch64` | pgdg | 28.9 KiB | [system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-system-stats` | 3.2 | `d12.x86_64` | pigsty | 66.1 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-system-stats` | 3.2 | `d12.aarch64` | pigsty | 65.5 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-system-stats` | 3.2 | `u22.x86_64` | pigsty | 68.9 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-system-stats` | 3.2 | `u22.aarch64` | pigsty | 68.9 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-system-stats` | 3.2 | `u24.x86_64` | pigsty | 61.6 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-system-stats` | 3.2 | `u24.aarch64` | pigsty | 61.7 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_16` | 3.2 | `el8.x86_64` | pgdg | 28.8 KiB | [system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | 2.1 | `el8.x86_64` | pgdg | 28.1 KiB | [system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | 3.2 | `el8.aarch64` | pgdg | 28.6 KiB | [system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | 2.1 | `el8.aarch64` | pgdg | 27.9 KiB | [system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | 3.2 | `el9.x86_64` | pgdg | 28.5 KiB | [system_stats_16-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_16` | 2.1 | `el9.x86_64` | pgdg | 27.7 KiB | [system_stats_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_16` | 3.2 | `el9.aarch64` | pgdg | 28.5 KiB | [system_stats_16-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_16` | 2.1 | `el9.aarch64` | pgdg | 27.4 KiB | [system_stats_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_16` | 3.2 | `el10.x86_64` | pgdg | 29.6 KiB | [system_stats_16-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/system_stats_16-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_16` | 3.2 | `el10.aarch64` | pgdg | 28.9 KiB | [system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-system-stats` | 3.2 | `d12.x86_64` | pigsty | 65.9 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-system-stats` | 3.2 | `d12.aarch64` | pigsty | 65.4 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-system-stats` | 3.2 | `u22.x86_64` | pigsty | 68.8 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-system-stats` | 3.2 | `u22.aarch64` | pigsty | 68.8 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-system-stats` | 3.2 | `u24.x86_64` | pigsty | 61.6 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-system-stats` | 3.2 | `u24.aarch64` | pigsty | 61.7 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_15` | 3.2 | `el8.x86_64` | pgdg | 29.6 KiB | [system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | 2.1 | `el8.x86_64` | pgdg | 28.9 KiB | [system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | 1.0 | `el8.x86_64` | pgdg | 74.7 KiB | [system_stats_15-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-1.0-2.rhel8.x86_64.rpm) |
| `system_stats_15` | 3.2 | `el8.aarch64` | pgdg | 29.3 KiB | [system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | 2.1 | `el8.aarch64` | pgdg | 28.6 KiB | [system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | 1.0 | `el8.aarch64` | pgdg | 74.0 KiB | [system_stats_15-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_15` | 3.2 | `el9.x86_64` | pgdg | 30.4 KiB | [system_stats_15-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_15` | 2.1 | `el9.x86_64` | pgdg | 29.4 KiB | [system_stats_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_15` | 1.0 | `el9.x86_64` | pgdg | 76.4 KiB | [system_stats_15-1.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-1.0-2.rhel9.x86_64.rpm) |
| `system_stats_15` | 3.2 | `el9.aarch64` | pgdg | 30.2 KiB | [system_stats_15-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_15` | 2.1 | `el9.aarch64` | pgdg | 29.0 KiB | [system_stats_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_15` | 1.0 | `el9.aarch64` | pgdg | 75.7 KiB | [system_stats_15-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-1.0-2.rhel9.aarch64.rpm) |
| `system_stats_15` | 3.2 | `el10.x86_64` | pgdg | 31.2 KiB | [system_stats_15-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/system_stats_15-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_15` | 3.2 | `el10.aarch64` | pgdg | 30.7 KiB | [system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-system-stats` | 3.2 | `d12.x86_64` | pigsty | 66.7 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-system-stats` | 3.2 | `d12.aarch64` | pigsty | 66.1 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-system-stats` | 3.2 | `u22.x86_64` | pigsty | 70.2 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-system-stats` | 3.2 | `u22.aarch64` | pigsty | 70.2 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-system-stats` | 3.2 | `u24.x86_64` | pigsty | 62.9 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-system-stats` | 3.2 | `u24.aarch64` | pigsty | 62.8 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_14` | 3.2 | `el8.x86_64` | pgdg | 29.6 KiB | [system_stats_14-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_14` | 2.1 | `el8.x86_64` | pgdg | 28.9 KiB | [system_stats_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_14` | 1.0 | `el8.x86_64` | pgdg | 75.0 KiB | [system_stats_14-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-1.0-2.rhel8.x86_64.rpm) |
| `system_stats_14` | 3.2 | `el8.aarch64` | pgdg | 29.3 KiB | [system_stats_14-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_14` | 2.1 | `el8.aarch64` | pgdg | 28.6 KiB | [system_stats_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_14` | 1.0 | `el8.aarch64` | pgdg | 74.0 KiB | [system_stats_14-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_14` | 3.2 | `el9.x86_64` | pgdg | 30.2 KiB | [system_stats_14-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/system_stats_14-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_14` | 2.1 | `el9.x86_64` | pgdg | 29.4 KiB | [system_stats_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/system_stats_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_14` | 3.2 | `el9.aarch64` | pgdg | 30.3 KiB | [system_stats_14-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_14` | 2.1 | `el9.aarch64` | pgdg | 29.0 KiB | [system_stats_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_14` | 1.0 | `el9.aarch64` | pgdg | 75.7 KiB | [system_stats_14-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-1.0-2.rhel9.aarch64.rpm) |
| `system_stats_14` | 3.2 | `el10.x86_64` | pgdg | 31.2 KiB | [system_stats_14-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/system_stats_14-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_14` | 3.2 | `el10.aarch64` | pgdg | 30.6 KiB | [system_stats_14-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/system_stats_14-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-system-stats` | 3.2 | `d12.x86_64` | pigsty | 66.8 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-system-stats` | 3.2 | `d12.aarch64` | pigsty | 66.1 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-system-stats` | 3.2 | `u22.x86_64` | pigsty | 70.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-system-stats` | 3.2 | `u22.aarch64` | pigsty | 70.1 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-system-stats` | 3.2 | `u24.x86_64` | pigsty | 62.9 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-system-stats` | 3.2 | `u24.aarch64` | pigsty | 62.7 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_13` | 3.2 | `el8.x86_64` | pgdg | 29.3 KiB | [system_stats_13-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/system_stats_13-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_13` | 2.1 | `el8.x86_64` | pgdg | 28.6 KiB | [system_stats_13-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/system_stats_13-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_13` | 1.0 | `el8.x86_64` | pgdg | 73.8 KiB | [system_stats_13-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/system_stats_13-1.0-1.rhel8.x86_64.rpm) |
| `system_stats_13` | 3.2 | `el8.aarch64` | pgdg | 29.3 KiB | [system_stats_13-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/system_stats_13-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_13` | 2.1 | `el8.aarch64` | pgdg | 28.6 KiB | [system_stats_13-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/system_stats_13-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_13` | 1.0 | `el8.aarch64` | pgdg | 73.9 KiB | [system_stats_13-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/system_stats_13-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_13` | 3.2 | `el9.x86_64` | pgdg | 30.4 KiB | [system_stats_13-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/system_stats_13-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_13` | 2.1 | `el9.x86_64` | pgdg | 29.4 KiB | [system_stats_13-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/system_stats_13-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_13` | 3.2 | `el9.aarch64` | pgdg | 30.2 KiB | [system_stats_13-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/system_stats_13-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_13` | 2.1 | `el9.aarch64` | pgdg | 29.0 KiB | [system_stats_13-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/system_stats_13-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_13` | 1.0 | `el9.aarch64` | pgdg | 75.5 KiB | [system_stats_13-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/system_stats_13-1.0-2.rhel9.aarch64.rpm) |
| `system_stats_13` | 3.2 | `el10.x86_64` | pgdg | 31.2 KiB | [system_stats_13-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/system_stats_13-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_13` | 3.2 | `el10.aarch64` | pgdg | 30.6 KiB | [system_stats_13-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/system_stats_13-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-system-stats` | 3.2 | `d12.x86_64` | pigsty | 66.6 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-system-stats` | 3.2 | `d12.aarch64` | pigsty | 66.0 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-system-stats` | 3.2 | `u22.x86_64` | pigsty | 70.0 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-system-stats` | 3.2 | `u22.aarch64` | pigsty | 70.1 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-system-stats` | 3.2 | `u24.x86_64` | pigsty | 62.8 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-system-stats` | 3.2 | `u24.aarch64` | pigsty | 62.6 KiB | [postgresql-13-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-13-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/system_stats" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/system_stats" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="system_stats-3.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get system_stats; # get system_stats source code
pig build dep system_stats; # install build dependencies
pig build pkg system_stats; # build extension rpm or deb
pig build ext system_stats; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install system_stats; # install by extension name, for the current active PG version
pig ext install system_stats; # install via package alias, for the active PG version
pig ext install system_stats -v 18;   # install for PG 18
pig ext install system_stats -v 17;   # install for PG 17
pig ext install system_stats -v 16;   # install for PG 16
pig ext install system_stats -v 15;   # install for PG 15
pig ext install system_stats -v 14;   # install for PG 14
pig ext install system_stats -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION system_stats;
```

