---
title: "pg_wait_sampling"
linkTitle: "pg_wait_sampling"
description: "sampling based statistics of wait events"
weight: 6270
categories: ["STAT"]
width: full
---

[**pg_wait_sampling**](https://github.com/postgrespro/pg_wait_sampling) : sampling based statistics of wait events


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6270** | {{< badge content="pg_wait_sampling" link="https://github.com/postgrespro/pg_wait_sampling" >}} | {{< ext "pg_wait_sampling" >}} | `1.1.9` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_background" >}} {{< ext "pg_stat_kcache" >}} {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "pg_profile" >}} {{< ext "pgsentinel" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_wait_sampling` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "pg_wait_sampling_18" "green" >}} {{< bg "17" "pg_wait_sampling_17" "green" >}} {{< bg "16" "pg_wait_sampling_16" "green" >}} {{< bg "15" "pg_wait_sampling_15" "green" >}} {{< bg "14" "pg_wait_sampling_14" "green" >}} {{< bg "13" "pg_wait_sampling_13" "green" >}} | `pg_wait_sampling_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "postgresql-18-pg-wait-sampling" "green" >}} {{< bg "17" "postgresql-17-pg-wait-sampling" "green" >}} {{< bg "16" "postgresql-16-pg-wait-sampling" "green" >}} {{< bg "15" "postgresql-15-pg-wait-sampling" "green" >}} {{< bg "14" "postgresql-14-pg-wait-sampling" "green" >}} {{< bg "13" "postgresql-13-pg-wait-sampling" "green" >}} | `postgresql-$v-pg-wait-sampling` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.9" "pg_wait_sampling_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-pg-wait-sampling : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-13-pg-wait-sampling : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_18` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_wait_sampling_18-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_18` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_wait_sampling_18-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_18` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.0 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_wait_sampling_18-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_18` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_wait_sampling_18-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_18` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.5 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_wait_sampling_18-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_18` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pg_wait_sampling_18-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_wait_sampling_18-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.2 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.4 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.0 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.8 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.5 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.4 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.0 KiB | [postgresql-18-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-18-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_17` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_wait_sampling_17-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_wait_sampling_17-1.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pg_wait_sampling_17-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_wait_sampling_17-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_17-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_wait_sampling_17-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_wait_sampling_17-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_wait_sampling_17-1.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.8 KiB | [pg_wait_sampling_17-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_wait_sampling_17-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pg_wait_sampling_17-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_wait_sampling_17-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.9 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_wait_sampling_17-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.8 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_wait_sampling_17-1.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pg_wait_sampling_17-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_wait_sampling_17-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_17-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_wait_sampling_17-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_wait_sampling_17-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_wait_sampling_17-1.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pg_wait_sampling_17-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_wait_sampling_17-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_17-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_wait_sampling_17-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_wait_sampling_17-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_wait_sampling_17-1.1.8-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_17` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pg_wait_sampling_17-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_wait_sampling_17-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `pg_wait_sampling_17` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pg_wait_sampling_17-1.1.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_wait_sampling_17-1.1.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.4 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.2 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.4 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.2 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.3 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.8 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.4 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-17-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-17-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_16` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_wait_sampling_16-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_wait_sampling_16-1.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pg_wait_sampling_16-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_wait_sampling_16-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_16-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_wait_sampling_16-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.3 KiB | [pg_wait_sampling_16-1.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_wait_sampling_16-1.1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_wait_sampling_16-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_wait_sampling_16-1.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.8 KiB | [pg_wait_sampling_16-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_wait_sampling_16-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pg_wait_sampling_16-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_wait_sampling_16-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.9 KiB | [pg_wait_sampling_16-1.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_wait_sampling_16-1.1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.0 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_wait_sampling_16-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.8 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_wait_sampling_16-1.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pg_wait_sampling_16-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_wait_sampling_16-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_16-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_wait_sampling_16-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.0 KiB | [pg_wait_sampling_16-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_wait_sampling_16-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_wait_sampling_16-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_wait_sampling_16-1.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pg_wait_sampling_16-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_wait_sampling_16-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_16-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_wait_sampling_16-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.9 KiB | [pg_wait_sampling_16-1.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_wait_sampling_16-1.1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_wait_sampling_16-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_wait_sampling_16-1.1.8-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_16` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pg_wait_sampling_16-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_wait_sampling_16-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `pg_wait_sampling_16` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pg_wait_sampling_16-1.1.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_wait_sampling_16-1.1.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.4 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.5 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.0 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.4 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.9 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.4 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 37.9 KiB | [postgresql-16-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-16-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_15` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.6 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_wait_sampling_15-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_wait_sampling_15-1.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_15-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_wait_sampling_15-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_15-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_wait_sampling_15-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.0 KiB | [pg_wait_sampling_15-1.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_wait_sampling_15-1.1.4-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.4 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_wait_sampling_15-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_wait_sampling_15-1.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_15-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_wait_sampling_15-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_15-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_wait_sampling_15-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.7 KiB | [pg_wait_sampling_15-1.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_wait_sampling_15-1.1.4-1.rhel8.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_wait_sampling_15-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_wait_sampling_15-1.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pg_wait_sampling_15-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_wait_sampling_15-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_15-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_wait_sampling_15-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.0 KiB | [pg_wait_sampling_15-1.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_wait_sampling_15-1.1.4-1.rhel9.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_wait_sampling_15-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_wait_sampling_15-1.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pg_wait_sampling_15-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_wait_sampling_15-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pg_wait_sampling_15-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_wait_sampling_15-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.7 KiB | [pg_wait_sampling_15-1.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_wait_sampling_15-1.1.4-1.rhel9.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_wait_sampling_15-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_wait_sampling_15-1.1.8-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_15` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.0 KiB | [pg_wait_sampling_15-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_wait_sampling_15-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `pg_wait_sampling_15` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.9 KiB | [pg_wait_sampling_15-1.1.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_wait_sampling_15-1.1.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.1 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.6 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.6 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 43.0 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.6 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.3 KiB | [postgresql-15-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-15-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_14` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.6 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_14-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_14-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.9 KiB | [pg_wait_sampling_14-1.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.4-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [pg_wait_sampling_14-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_wait_sampling_14-1.1.3-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.4 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_wait_sampling_14-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.3 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_wait_sampling_14-1.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_14-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_wait_sampling_14-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_14-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_wait_sampling_14-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.7 KiB | [pg_wait_sampling_14-1.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_wait_sampling_14-1.1.4-1.rhel8.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_wait_sampling_14-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_wait_sampling_14-1.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pg_wait_sampling_14-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_wait_sampling_14-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [pg_wait_sampling_14-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_wait_sampling_14-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.0 KiB | [pg_wait_sampling_14-1.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_wait_sampling_14-1.1.4-1.rhel9.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_wait_sampling_14-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_wait_sampling_14-1.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_14-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_wait_sampling_14-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pg_wait_sampling_14-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_wait_sampling_14-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.7 KiB | [pg_wait_sampling_14-1.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_wait_sampling_14-1.1.4-1.rhel9.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_wait_sampling_14-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_wait_sampling_14-1.1.8-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_14` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.0 KiB | [pg_wait_sampling_14-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_wait_sampling_14-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `pg_wait_sampling_14` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.9 KiB | [pg_wait_sampling_14-1.1.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_wait_sampling_14-1.1.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.5 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.2 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.6 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.5 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.9 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.6 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.3 KiB | [postgresql-14-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-14-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_wait_sampling_13` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.1 KiB | [pg_wait_sampling_13-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.1 KiB | [pg_wait_sampling_13-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.6 KiB | [pg_wait_sampling_13-1.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.4-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.4 KiB | [pg_wait_sampling_13-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.3-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.3 KiB | [pg_wait_sampling_13-1.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_wait_sampling_13-1.1.2-1.rhel8.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.4 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_wait_sampling_13-1.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_wait_sampling_13-1.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_13-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_wait_sampling_13-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_wait_sampling_13-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [pg_wait_sampling_13-1.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_wait_sampling_13-1.1.4-1.rhel8.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_wait_sampling_13-1.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_wait_sampling_13-1.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.0 KiB | [pg_wait_sampling_13-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_wait_sampling_13-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_wait_sampling_13-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.6 KiB | [pg_wait_sampling_13-1.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_wait_sampling_13-1.1.4-1.rhel9.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_wait_sampling_13-1.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_wait_sampling_13-1.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pg_wait_sampling_13-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_wait_sampling_13-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pg_wait_sampling_13-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_wait_sampling_13-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.6 KiB | [pg_wait_sampling_13-1.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_wait_sampling_13-1.1.4-1.rhel9.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_wait_sampling_13-1.1.9-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.6 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_wait_sampling_13-1.1.8-1PGDG.rhel10.x86_64.rpm) |
| `pg_wait_sampling_13` | `1.1.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.0 KiB | [pg_wait_sampling_13-1.1.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_wait_sampling_13-1.1.9-1PGDG.rhel10.aarch64.rpm) |
| `pg_wait_sampling_13` | `1.1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.9 KiB | [pg_wait_sampling_13-1.1.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_wait_sampling_13-1.1.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.0 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.1 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.1 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.1 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.5 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.1 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-wait-sampling` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.2 KiB | [postgresql-13-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-wait-sampling/postgresql-13-pg-wait-sampling_1.1.9-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/pg_wait_sampling" title="Repository" icon="github" subtitle="github.com/postgrespro/pg_wait_sampling" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_wait_sampling;		# install via package name, for the active PG version

pig install pg_wait_sampling -v 18;   # install for PG 18
pig install pg_wait_sampling -v 17;   # install for PG 17
pig install pg_wait_sampling -v 16;   # install for PG 16
pig install pg_wait_sampling -v 15;   # install for PG 15
pig install pg_wait_sampling -v 14;   # install for PG 14
pig install pg_wait_sampling -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_wait_sampling';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_wait_sampling;
```
