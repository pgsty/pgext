---
title: "system_stats"
linkTitle: "system_stats"
description: "EnterpriseDB system statistics for PostgreSQL"
weight: 6420
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**system_stats**](https://github.com/EnterpriseDB/system_stats) : EnterpriseDB system statistics for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6420** | {{< badge content="system_stats" link="https://github.com/EnterpriseDB/system_stats" >}} | {{< ext "system_stats" >}} | `4.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pgmeminfo" >}} {{< ext "pg_stat_ch" >}} {{< ext "pgmonitor" >}} {{< ext "pg_datasentinel" >}} {{< ext "pgsentinel" >}} {{< ext "pgexporter_ext" >}} {{< ext "pg_statvfs" >}} |

> [!Note] PGDG RPM and PIGSTY DEB package/source version 4.1 cover PostgreSQL 14-18; SQL/control version is 4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `4.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `system_stats` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.1` | {{< bg "19" "system_stats_19" "green" >}} {{< bg "18" "system_stats_18" "green" >}} {{< bg "17" "system_stats_17" "green" >}} {{< bg "16" "system_stats_16" "green" >}} {{< bg "15" "system_stats_15" "green" >}} | `system_stats_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.1` | {{< bg "19" "postgresql-19-system-stats" "green" >}} {{< bg "18" "postgresql-18-system-stats" "green" >}} {{< bg "17" "postgresql-17-system-stats" "green" >}} {{< bg "16" "postgresql-16-system-stats" "green" >}} {{< bg "15" "postgresql-15-system-stats" "green" >}} | `postgresql-$v-system-stats` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "system_stats_19 : N/A 0" "gray" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "system_stats_19 : N/A 0" "gray" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 7" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.1" "system_stats_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.1" "system_stats_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.1" "system_stats_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.1" "system_stats_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_18 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.1" "system_stats_15 : AVAIL 10" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.1" "postgresql-19-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-18-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-17-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-16-system-stats : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.1" "postgresql-15-system-stats : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_19` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [system_stats_19-4.1-5PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/system_stats_19-4.1-5PGDG.rhel9.9.x86_64.rpm) |
| `system_stats_19` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [system_stats_19-4.1-5PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/system_stats_19-4.1-5PGDG.rhel9.9.aarch64.rpm) |
| `system_stats_19` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [system_stats_19-4.1-5PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/system_stats_19-4.1-5PGDG.rhel10.3.x86_64.rpm) |
| `system_stats_19` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_19-4.1-5PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/system_stats_19-4.1-5PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-system-stats` | `4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.8 KiB | [postgresql-19-system-stats_4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-system-stats` | `4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.8 KiB | [postgresql-19-system-stats_4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-system-stats` | `4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.8 KiB | [postgresql-19-system-stats_4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-system-stats` | `4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 66.8 KiB | [postgresql-19-system-stats_4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.6 KiB | [postgresql-19-system-stats_4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.5 KiB | [postgresql-19-system-stats_4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.3 KiB | [postgresql-19-system-stats_4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.9 KiB | [postgresql-19-system-stats_4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 70.2 KiB | [postgresql-19-system-stats_4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-system-stats` | `4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.2 KiB | [postgresql-19-system-stats_4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-19-system-stats_4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_18` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [system_stats_18-4.1-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-4.1-5PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_18` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.7 KiB | [system_stats_18-4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_18` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [system_stats_18-4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_18` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_18-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_18` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/system_stats_18-3.2-2PGDG.rhel8.x86_64.rpm) |
| `system_stats_18` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.1 KiB | [system_stats_18-4.1-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-4.1-5PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [system_stats_18-4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_18` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.0 KiB | [system_stats_18-4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [system_stats_18-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_18` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.7 KiB | [system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/system_stats_18-3.2-2PGDG.rhel8.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [system_stats_18-4.1-5PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/system_stats_18-4.1-5PGDG.rhel9.9.x86_64.rpm) |
| `system_stats_18` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [system_stats_18-4.1-5PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/system_stats_18-4.1-5PGDG.rhel9.9.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [system_stats_18-4.1-5PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/system_stats_18-4.1-5PGDG.rhel10.3.x86_64.rpm) |
| `system_stats_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_18-4.1-5PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.1-5PGDG.rhel10.3.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_18-4.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.1-5PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_18-4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_18-4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.7 KiB | [system_stats_18-4.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_18` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.5 KiB | [system_stats_18-4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_18` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.5 KiB | [system_stats_18-4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_18-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_18` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_18-3.2.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_18` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.8 KiB | [system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/system_stats_18-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-system-stats` | `4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.8 KiB | [postgresql-18-system-stats_4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-system-stats` | `4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.8 KiB | [postgresql-18-system-stats_4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-system-stats` | `4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.8 KiB | [postgresql-18-system-stats_4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-system-stats` | `4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 66.8 KiB | [postgresql-18-system-stats_4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.6 KiB | [postgresql-18-system-stats_4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.4 KiB | [postgresql-18-system-stats_4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.5 KiB | [postgresql-18-system-stats_4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.9 KiB | [postgresql-18-system-stats_4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 70.2 KiB | [postgresql-18-system-stats_4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-system-stats` | `4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.2 KiB | [postgresql-18-system-stats_4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-18-system-stats_4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_17` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [system_stats_17-4.1-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-4.1-5PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_17` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.7 KiB | [system_stats_17-4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_17` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [system_stats_17-4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_17` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_17-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_17` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.8 KiB | [system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/system_stats_17-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_17` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.1 KiB | [system_stats_17-4.1-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-4.1-5PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [system_stats_17-4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_17` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.0 KiB | [system_stats_17-4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_17` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.0 KiB | [system_stats_17-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_17` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/system_stats_17-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [system_stats_17-4.1-5PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/system_stats_17-4.1-5PGDG.rhel9.9.x86_64.rpm) |
| `system_stats_17` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [system_stats_17-4.1-5PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/system_stats_17-4.1-5PGDG.rhel9.9.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [system_stats_17-4.1-5PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/system_stats_17-4.1-5PGDG.rhel10.3.x86_64.rpm) |
| `system_stats_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_17-4.1-5PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.1-5PGDG.rhel10.3.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_17-4.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.1-5PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_17-4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_17-4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_17-4.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_17` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [system_stats_17-4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_17` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [system_stats_17-4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_17` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_17-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_17` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_17-3.2.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_17` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.9 KiB | [system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/system_stats_17-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-system-stats` | `4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.6 KiB | [postgresql-17-system-stats_4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-system-stats` | `4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.6 KiB | [postgresql-17-system-stats_4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-system-stats` | `4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.6 KiB | [postgresql-17-system-stats_4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-system-stats` | `4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 66.7 KiB | [postgresql-17-system-stats_4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 77.9 KiB | [postgresql-17-system-stats_4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 77.7 KiB | [postgresql-17-system-stats_4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.4 KiB | [postgresql-17-system-stats_4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.7 KiB | [postgresql-17-system-stats_4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 70.1 KiB | [postgresql-17-system-stats_4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-system-stats` | `4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.1 KiB | [postgresql-17-system-stats_4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-17-system-stats_4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_16` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [system_stats_16-4.1-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-4.1-5PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_16` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.7 KiB | [system_stats_16-4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_16` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [system_stats_16-4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_16` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.3 KiB | [system_stats_16-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_16` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.8 KiB | [system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.1 KiB | [system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/system_stats_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_16` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.1 KiB | [system_stats_16-4.1-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-4.1-5PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [system_stats_16-4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_16` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.0 KiB | [system_stats_16-4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_16` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.9 KiB | [system_stats_16-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_16` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.9 KiB | [system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/system_stats_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [system_stats_16-4.1-5PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/system_stats_16-4.1-5PGDG.rhel9.9.x86_64.rpm) |
| `system_stats_16` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [system_stats_16-4.1-5PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/system_stats_16-4.1-5PGDG.rhel9.9.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [system_stats_16-4.1-5PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/system_stats_16-4.1-5PGDG.rhel10.3.x86_64.rpm) |
| `system_stats_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_16-4.1-5PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.1-5PGDG.rhel10.3.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [system_stats_16-4.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.1-5PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_16-4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_16-4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [system_stats_16-4.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_16` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [system_stats_16-4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_16` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [system_stats_16-4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_16` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_16-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_16` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [system_stats_16-3.2.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_16` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.9 KiB | [system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/system_stats_16-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-system-stats` | `4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.5 KiB | [postgresql-16-system-stats_4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-system-stats` | `4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.7 KiB | [postgresql-16-system-stats_4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-system-stats` | `4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.6 KiB | [postgresql-16-system-stats_4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-system-stats` | `4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 66.7 KiB | [postgresql-16-system-stats_4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 77.8 KiB | [postgresql-16-system-stats_4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 77.6 KiB | [postgresql-16-system-stats_4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.1 KiB | [postgresql-16-system-stats_4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.7 KiB | [postgresql-16-system-stats_4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 69.9 KiB | [postgresql-16-system-stats_4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-system-stats` | `4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.1 KiB | [postgresql-16-system-stats_4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-16-system-stats_4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `system_stats_15` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [system_stats_15-4.1-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-4.1-5PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_15` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [system_stats_15-4.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-4.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_15` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.2 KiB | [system_stats_15-4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_15` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [system_stats_15-3.2.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-3.2.1-1PGDG.rhel8.10.x86_64.rpm) |
| `system_stats_15` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-3.2-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `system_stats_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.7 KiB | [system_stats_15-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/system_stats_15-1.0-2.rhel8.x86_64.rpm) |
| `system_stats_15` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [system_stats_15-4.1-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-4.1-5PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.1 KiB | [system_stats_15-4.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-4.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_15` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [system_stats_15-4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_15` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.7 KiB | [system_stats_15-3.2.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-3.2.1-1PGDG.rhel8.10.aarch64.rpm) |
| `system_stats_15` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-3.2-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.6 KiB | [system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `system_stats_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [system_stats_15-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/system_stats_15-1.0-2.rhel8.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.1 KiB | [system_stats_15-4.1-5PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/system_stats_15-4.1-5PGDG.rhel9.9.x86_64.rpm) |
| `system_stats_15` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.8 KiB | [system_stats_15-4.1-5PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/system_stats_15-4.1-5PGDG.rhel9.9.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.9 KiB | [system_stats_15-4.1-5PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/system_stats_15-4.1-5PGDG.rhel10.3.x86_64.rpm) |
| `system_stats_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.2 KiB | [system_stats_15-4.1-5PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.1-5PGDG.rhel10.3.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.2 KiB | [system_stats_15-4.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.1-5PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.6 KiB | [system_stats_15-4.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.1-1PGDG.rhel10.2.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.6 KiB | [system_stats_15-4.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.6 KiB | [system_stats_15-4.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_15` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.4 KiB | [system_stats_15-4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_15` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.4 KiB | [system_stats_15-4.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-4.0-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_15` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.4 KiB | [system_stats_15-3.2.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2.1-1PGDG.rhel10.1.aarch64.rpm) |
| `system_stats_15` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.4 KiB | [system_stats_15-3.2.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2.1-1PGDG.rhel10.0.aarch64.rpm) |
| `system_stats_15` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/system_stats_15-3.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-system-stats` | `4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 68.4 KiB | [postgresql-15-system-stats_4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-system-stats` | `4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.3 KiB | [postgresql-15-system-stats_4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-system-stats` | `4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 68.5 KiB | [postgresql-15-system-stats_4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-system-stats` | `4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.4 KiB | [postgresql-15-system-stats_4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 79.6 KiB | [postgresql-15-system-stats_4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 79.4 KiB | [postgresql-15-system-stats_4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.7 KiB | [postgresql-15-system-stats_4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.1 KiB | [postgresql-15-system-stats_4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.4 KiB | [postgresql-15-system-stats_4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-system-stats` | `4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 70.6 KiB | [postgresql-15-system-stats_4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/system-stats/postgresql-15-system-stats_4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/system_stats" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/system_stats" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="system_stats-4.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg system_stats;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install system_stats;		# install via package name, for the active PG version

pig install system_stats -v 19;   # install for PG 19
pig install system_stats -v 18;   # install for PG 18
pig install system_stats -v 17;   # install for PG 17
pig install system_stats -v 16;   # install for PG 16
pig install system_stats -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION system_stats;
```




## Usage

Source: [README](https://github.com/EnterpriseDB/system_stats/blob/master/README.md), [Release v4.0](https://github.com/EnterpriseDB/system_stats/releases/tag/v4.0), [SQL install script](https://github.com/EnterpriseDB/system_stats/blob/master/system_stats--4.0.sql)

`system_stats` exposes operating-system metrics through SQL functions. It supports Linux, macOS, and Windows, returning `NULL` for fields that are not meaningful on the current platform.

### Main functions

```sql
CREATE EXTENSION system_stats;

SELECT * FROM pg_sys_os_info();
SELECT * FROM pg_sys_cpu_info();
SELECT * FROM pg_sys_cpu_usage_info();
SELECT * FROM pg_sys_memory_info();
SELECT * FROM pg_sys_io_analysis_info();
SELECT * FROM pg_sys_disk_info();
SELECT * FROM pg_sys_load_avg_info();
SELECT * FROM pg_sys_process_info();
SELECT * FROM pg_sys_network_info();
SELECT * FROM pg_sys_cpu_memory_by_process();
```

These functions cover OS identity, CPU inventory and usage, memory, block-device I/O, disks, load average, process counts, network interfaces, and per-process CPU and memory usage.

### Access control

```sql
GRANT monitor_system_stats TO nagios;
GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;
```

- The extension creates a `monitor_system_stats` role and grants execution on the shipped functions to that role.
- Functions are revoked from `PUBLIC`.

### Caveats

- The `monitor_system_stats` role is not dropped automatically when the extension is removed.
- macOS cannot expose full per-process details for processes owned by other users; those rows may contain only PID and process name.
- Pigsty metadata tracks `system_stats` 4.0 for PostgreSQL 14-18; the RPM and DEB package names differ (`system_stats_$v` vs. `postgresql-$v-system-stats`).
- Current v4.0 upstream docs keep the same user-facing function family and security model; the refresh here mainly aligns names, privileges, and platform notes with the current README and SQL script.
