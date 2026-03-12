---
title: "system_stats"
linkTitle: "system_stats"
description: "EnterpriseDB system statistics for PostgreSQL"
weight: 6420
categories: ["STAT"]
width: full
---

[**system_stats**](https://github.com/EnterpriseDB/system_stats) : EnterpriseDB system statistics for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6420** | {{< badge content="system_stats" link="https://github.com/EnterpriseDB/system_stats" >}} | {{< ext "system_stats" >}} | `3.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "pg_proctab" >}} {{< ext "pgmeminfo" >}} {{< ext "pgfincore" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_background" >}} {{< ext "pg_cooldown" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `3.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `system_stats` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.2` | {{< bg "18" "system_stats_18" "green" >}} {{< bg "17" "system_stats_17" "green" >}} {{< bg "16" "system_stats_16" "green" >}} {{< bg "15" "system_stats_15" "green" >}} {{< bg "14" "system_stats_14" "green" >}} | `system_stats_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2` | {{< bg "18" "postgresql-18-system-stats" "green" >}} {{< bg "17" "postgresql-17-system-stats" "green" >}} {{< bg "16" "postgresql-16-system-stats" "green" >}} {{< bg "15" "postgresql-15-system-stats" "green" >}} {{< bg "14" "postgresql-14-system-stats" "green" >}} | `postgresql-$v-system-stats` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2.1" "system_stats_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.2" "system_stats_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "system_stats_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-system-stats : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_18` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm) |
| `system_stats_18` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_18-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_18` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.7 KiB | [system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [system_stats_18-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_18` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.4 KiB | [system_stats_18-3.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/system_stats_18-3.2-2PGDG.rhel9.x86_64.rpm) |
| `system_stats_18` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [system_stats_18-3.2.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/system_stats_18-3.2.1-1PGDG.rhel9.7.x86_64.rpm) |
| `system_stats_18` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.3 KiB | [system_stats_18-3.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/system_stats_18-3.2-2PGDG.rhel9.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [system_stats_18-3.2.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/system_stats_18-3.2.1-1PGDG.rhel9.7.aarch64.rpm) |
| `system_stats_18` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [system_stats_18-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/system_stats_18-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_18` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.3 KiB | [system_stats_18-3.2.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/system_stats_18-3.2.1-1PGDG.rhel10.1.x86_64.rpm) |
| `system_stats_18` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.8 KiB | [system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_18-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-system-stats` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.2 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-system-stats` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.6 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-system-stats` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.4 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-system-stats` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.7 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-system-stats` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.8 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-system-stats` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.9 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-system-stats` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-system-stats` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.2 KiB | [postgresql-18-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-18-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_17` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_17-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_17` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.8 KiB | [system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_17` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [system_stats_17-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_17` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_17` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [system_stats_17-3.2.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/system_stats_17-3.2.1-1PGDG.rhel9.7.x86_64.rpm) |
| `system_stats_17` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.5 KiB | [system_stats_17-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/system_stats_17-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_17` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [system_stats_17-3.2.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/system_stats_17-3.2.1-1PGDG.rhel9.7.aarch64.rpm) |
| `system_stats_17` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.5 KiB | [system_stats_17-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/system_stats_17-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_17` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [system_stats_17-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/system_stats_17-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_17` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [system_stats_17-3.2.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/system_stats_17-3.2.1-1PGDG.rhel10.1.x86_64.rpm) |
| `system_stats_17` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.9 KiB | [system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm) |
| `system_stats_17` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_17-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-system-stats` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.2 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-system-stats` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.5 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-system-stats` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.4 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-system-stats` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.6 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-system-stats` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 68.9 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-system-stats` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.0 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-system-stats` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-system-stats` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.2 KiB | [postgresql-17-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_16` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_16-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_16` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.8 KiB | [system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.1 KiB | [system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.9 KiB | [system_stats_16-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_16` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.9 KiB | [system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [system_stats_16-3.2.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-3.2.1-1PGDG.rhel9.7.x86_64.rpm) |
| `system_stats_16` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.5 KiB | [system_stats_16-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_16` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.7 KiB | [system_stats_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_16` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [system_stats_16-3.2.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-3.2.1-1PGDG.rhel9.7.aarch64.rpm) |
| `system_stats_16` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.5 KiB | [system_stats_16-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_16` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.4 KiB | [system_stats_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_16` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 29.6 KiB | [system_stats_16-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/system_stats_16-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_16` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.3 KiB | [system_stats_16-3.2.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/system_stats_16-3.2.1-1PGDG.rhel10.1.x86_64.rpm) |
| `system_stats_16` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.9 KiB | [system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm) |
| `system_stats_16` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_16-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-16-system-stats` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.2 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-system-stats` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.5 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-system-stats` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.3 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-system-stats` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.6 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-system-stats` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 68.8 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-system-stats` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 68.9 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-system-stats` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 61.4 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-system-stats` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.2 KiB | [postgresql-16-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_15` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [system_stats_15-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_15` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.7 KiB | [system_stats_15-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-1.0-2.rhel8.x86_64.rpm) |
| `system_stats_15` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.7 KiB | [system_stats_15-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_15` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [system_stats_15-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_15` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.3 KiB | [system_stats_15-3.2.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-3.2.1-1PGDG.rhel9.7.x86_64.rpm) |
| `system_stats_15` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.4 KiB | [system_stats_15-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.4 KiB | [system_stats_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.4 KiB | [system_stats_15-1.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-1.0-2.rhel9.x86_64.rpm) |
| `system_stats_15` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.2 KiB | [system_stats_15-3.2.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-3.2.1-1PGDG.rhel9.7.aarch64.rpm) |
| `system_stats_15` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [system_stats_15-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.0 KiB | [system_stats_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 75.7 KiB | [system_stats_15-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-1.0-2.rhel9.aarch64.rpm) |
| `system_stats_15` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.2 KiB | [system_stats_15-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/system_stats_15-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_15` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.0 KiB | [system_stats_15-3.2.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/system_stats_15-3.2.1-1PGDG.rhel10.1.x86_64.rpm) |
| `system_stats_15` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm) |
| `system_stats_15` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.4 KiB | [system_stats_15-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-15-system-stats` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.0 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-system-stats` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.3 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-system-stats` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.2 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-system-stats` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.3 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-system-stats` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.2 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-system-stats` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.2 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-system-stats` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.8 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-system-stats` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.4 KiB | [postgresql-15-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_14` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [system_stats_14-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_14` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [system_stats_14-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [system_stats_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.0 KiB | [system_stats_14-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/system_stats_14-1.0-2.rhel8.x86_64.rpm) |
| `system_stats_14` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.7 KiB | [system_stats_14-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_14` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [system_stats_14-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [system_stats_14-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/system_stats_14-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_14` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.3 KiB | [system_stats_14-3.2.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/system_stats_14-3.2.1-1PGDG.rhel9.7.x86_64.rpm) |
| `system_stats_14` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.2 KiB | [system_stats_14-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/system_stats_14-3.2-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.4 KiB | [system_stats_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/system_stats_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `system_stats_14` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.2 KiB | [system_stats_14-3.2.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-3.2.1-1PGDG.rhel9.7.aarch64.rpm) |
| `system_stats_14` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [system_stats_14-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-3.2-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.0 KiB | [system_stats_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `system_stats_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 75.7 KiB | [system_stats_14-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/system_stats_14-1.0-2.rhel9.aarch64.rpm) |
| `system_stats_14` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.2 KiB | [system_stats_14-3.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/system_stats_14-3.2-2PGDG.rhel10.x86_64.rpm) |
| `system_stats_14` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.0 KiB | [system_stats_14-3.2.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/system_stats_14-3.2.1-1PGDG.rhel10.1.x86_64.rpm) |
| `system_stats_14` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [system_stats_14-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/system_stats_14-3.2-2PGDG.rhel10.aarch64.rpm) |
| `system_stats_14` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [system_stats_14-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/system_stats_14-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-14-system-stats` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.0 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-system-stats` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-system-stats` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-system-stats` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-system-stats` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-system-stats` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.2 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-system-stats` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.7 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-system-stats` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.3 KiB | [postgresql-14-system-stats_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-14-system-stats_3.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/system_stats" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/system_stats" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="system_stats-3.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg system_stats;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install system_stats;		# install via package name, for the active PG version

pig install system_stats -v 18;   # install for PG 18
pig install system_stats -v 17;   # install for PG 17
pig install system_stats -v 16;   # install for PG 16
pig install system_stats -v 15;   # install for PG 15
pig install system_stats -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION system_stats;
```



## Usage

> [system_stats: system-level statistics for PostgreSQL](https://github.com/EnterpriseDB/system_stats)

system_stats provides SQL functions to access OS-level statistics for monitoring. Supports Linux, macOS, and Windows.

### Functions

```sql
-- Operating system info (name, version, hostname, uptime, etc.)
SELECT * FROM pg_sys_os_info();

-- CPU info (vendor, model, cores, clock speed, cache sizes)
SELECT * FROM pg_sys_cpu_info();

-- CPU usage (user, system, idle, iowait percentages)
SELECT * FROM pg_sys_cpu_usage_info();

-- Memory info (total, used, free, swap, cache in bytes)
SELECT * FROM pg_sys_memory_info();

-- Disk I/O analysis (reads, writes, bytes, time per device)
SELECT * FROM pg_sys_io_analysis_info();

-- Disk info (filesystem, mount point, total/used/available space)
SELECT * FROM pg_sys_disk_info();

-- Load average (1, 5, 10, 15 minute intervals)
SELECT * FROM pg_sys_load_avg_info();

-- Process info (total, running, sleeping, stopped, zombie counts)
SELECT * FROM pg_sys_process_info();

-- Network info (interface, IP, bytes/packets sent/received, errors)
SELECT * FROM pg_sys_network_info();

-- Per-process CPU and memory usage
SELECT * FROM pg_sys_cpu_memory_by_process();
```

### Security

Access is restricted to superusers and members of the `monitor_system_stats` role:

```sql
-- Grant access to a monitoring user
GRANT monitor_system_stats TO nagios;

-- Or grant per-function to pg_monitor
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

### Example: System Overview

```sql
SELECT * FROM pg_sys_load_avg_info();  -- load averages
SELECT memfree, memused, swapfree FROM pg_sys_memory_info();  -- memory
SELECT * FROM pg_sys_cpu_usage_info();  -- CPU usage breakdown
```
