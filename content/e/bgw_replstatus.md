---
title: "bgw_replstatus"
linkTitle: "bgw_replstatus"
description: "Small PostgreSQL background worker to report whether a node is a replication master or standby"
weight: 6340
categories: ["STAT"]
width: full
---

[**bgw_replstatus**](https://github.com/mhagander/bgw_replstatus) : Small PostgreSQL background worker to report whether a node is a replication master or standby


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6340** | {{< badge content="bgw_replstatus" link="https://github.com/mhagander/bgw_replstatus" >}} | {{< ext "bgw_replstatus" >}} | `1.0.8` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgautofailover" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgpool_recovery" >}} {{< ext "pgsentinel" >}} {{< ext "pglogical_origin" >}} {{< ext "repmgr" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `bgw_replstatus` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.8` | {{< bg "18" "bgw_replstatus_18*" "green" >}} {{< bg "17" "bgw_replstatus_17*" "green" >}} {{< bg "16" "bgw_replstatus_16*" "green" >}} {{< bg "15" "bgw_replstatus_15*" "green" >}} {{< bg "14" "bgw_replstatus_14*" "green" >}} {{< bg "13" "bgw_replstatus_13*" "green" >}} | `bgw_replstatus_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.8` | {{< bg "18" "postgresql-18-bgw-replstatus" "green" >}} {{< bg "17" "postgresql-17-bgw-replstatus" "green" >}} {{< bg "16" "postgresql-16-bgw-replstatus" "green" >}} {{< bg "15" "postgresql-15-bgw-replstatus" "green" >}} {{< bg "14" "postgresql-14-bgw-replstatus" "green" >}} {{< bg "13" "postgresql-13-bgw-replstatus" "green" >}} | `postgresql-$v-bgw-replstatus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "bgw_replstatus_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.8" "postgresql-18-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-17-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-16-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-15-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-14-bgw-replstatus : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.8" "postgresql-13-bgw-replstatus : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_18` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_18` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_18` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_18` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.8 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_18` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/bgw_replstatus_18-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_18` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_18-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/bgw_replstatus_18-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.7 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 14.9 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.8 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 14.9 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.0 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 14.6 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.8 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 14.9 KiB | [postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-18-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_17` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_17` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.7 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/bgw_replstatus_17-1.0.6-4PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_17` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_17` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.6 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/bgw_replstatus_17-1.0.6-4PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_17` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_17` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.5 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/bgw_replstatus_17-1.0.6-4PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_17` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.0 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_17` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [bgw_replstatus_17-1.0.6-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/bgw_replstatus_17-1.0.6-4PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_17` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/bgw_replstatus_17-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_17` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_17-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/bgw_replstatus_17-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.7 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 14.8 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.7 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 14.8 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.4 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.0 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.7 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 14.8 KiB | [postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-17-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_16` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.6-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/bgw_replstatus_16-1.0.6-2.rhel8.1.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_16` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.6-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/bgw_replstatus_16-1.0.6-2.rhel8.1.aarch64.rpm) |
| `bgw_replstatus_16` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.3 KiB | [bgw_replstatus_16-1.0.6-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.6-2.rhel9.1.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.5 KiB | [bgw_replstatus_16-1.0.6-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/bgw_replstatus_16-1.0.6-3PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.9 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_16` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.2 KiB | [bgw_replstatus_16-1.0.6-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/bgw_replstatus_16-1.0.6-2.rhel9.1.aarch64.rpm) |
| `bgw_replstatus_16` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/bgw_replstatus_16-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_16` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_16-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/bgw_replstatus_16-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.7 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 14.8 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.7 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 14.8 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.4 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.0 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.7 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 14.8 KiB | [postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-16-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_15` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_15` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.4 KiB | [bgw_replstatus_15-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/bgw_replstatus_15-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_15` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_15` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.3 KiB | [bgw_replstatus_15-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/bgw_replstatus_15-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_15` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_15` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_15-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/bgw_replstatus_15-1.0.6-1.rhel9.x86_64.rpm) |
| `bgw_replstatus_15` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.0 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_15` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.9 KiB | [bgw_replstatus_15-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/bgw_replstatus_15-1.0.6-1.rhel9.aarch64.rpm) |
| `bgw_replstatus_15` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/bgw_replstatus_15-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_15` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_15-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/bgw_replstatus_15-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.7 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 14.8 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.7 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 14.8 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.4 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.0 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.7 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 14.8 KiB | [postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-15-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_14` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.4 KiB | [bgw_replstatus_14-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [bgw_replstatus_14-1.0.3-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/bgw_replstatus_14-1.0.3-3.rhel8.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_14` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.3 KiB | [bgw_replstatus_14-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/bgw_replstatus_14-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_14` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_14-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/bgw_replstatus_14-1.0.6-1.rhel9.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 15.0 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_14` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.9 KiB | [bgw_replstatus_14-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/bgw_replstatus_14-1.0.6-1.rhel9.aarch64.rpm) |
| `bgw_replstatus_14` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.5 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/bgw_replstatus_14-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_14` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_14-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/bgw_replstatus_14-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.7 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 14.8 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.7 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 14.8 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.4 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.0 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.7 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 14.8 KiB | [postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-14-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `bgw_replstatus_13` | `1.0.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/bgw_replstatus_13-1.0.8-1PGDG.rhel8.x86_64.rpm) |
| `bgw_replstatus_13` | `1.0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.3 KiB | [bgw_replstatus_13-1.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/bgw_replstatus_13-1.0.6-1.rhel8.x86_64.rpm) |
| `bgw_replstatus_13` | `1.0.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/bgw_replstatus_13-1.0.8-1PGDG.rhel8.aarch64.rpm) |
| `bgw_replstatus_13` | `1.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.3 KiB | [bgw_replstatus_13-1.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/bgw_replstatus_13-1.0.6-1.rhel8.aarch64.rpm) |
| `bgw_replstatus_13` | `1.0.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/bgw_replstatus_13-1.0.8-1PGDG.rhel9.x86_64.rpm) |
| `bgw_replstatus_13` | `1.0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 15.2 KiB | [bgw_replstatus_13-1.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/bgw_replstatus_13-1.0.6-1.rhel9.x86_64.rpm) |
| `bgw_replstatus_13` | `1.0.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.9 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/bgw_replstatus_13-1.0.8-1PGDG.rhel9.aarch64.rpm) |
| `bgw_replstatus_13` | `1.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.9 KiB | [bgw_replstatus_13-1.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/bgw_replstatus_13-1.0.6-1.rhel9.aarch64.rpm) |
| `bgw_replstatus_13` | `1.0.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.4 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/bgw_replstatus_13-1.0.8-1PGDG.rhel10.x86_64.rpm) |
| `bgw_replstatus_13` | `1.0.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.5 KiB | [bgw_replstatus_13-1.0.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/bgw_replstatus_13-1.0.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 14.6 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.0 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg12+1_arm64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 14.6 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg13+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.1 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg13+1_arm64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.2 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 14.9 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 14.7 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-bgw-replstatus` | `1.0.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.1 KiB | [postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/b/bgw-replstatus/postgresql-13-bgw-replstatus_1.0.8-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/mhagander/bgw_replstatus" title="Repository" icon="github" subtitle="github.com/mhagander/bgw_replstatus" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install bgw_replstatus;		# install via package name, for the active PG version

pig install bgw_replstatus -v 18;   # install for PG 18
pig install bgw_replstatus -v 17;   # install for PG 17
pig install bgw_replstatus -v 16;   # install for PG 16
pig install bgw_replstatus -v 15;   # install for PG 15
pig install bgw_replstatus -v 14;   # install for PG 14
pig install bgw_replstatus -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'bgw_replstatus';
```


This extension does not need `CREATE EXTENSION` DDL command


