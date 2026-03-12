---
title: "toastinfo"
linkTitle: "toastinfo"
description: "show details on toasted datums"
weight: 6530
categories: ["STAT"]
width: full
---

[**toastinfo**](https://github.com/credativ/toastinfo) : show details on toasted datums


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6530** | {{< badge content="toastinfo" link="https://github.com/credativ/toastinfo" >}} | {{< ext "toastinfo" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "pg_relusage" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `toastinfo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "toastinfo_18" "green" >}} {{< bg "17" "toastinfo_17" "green" >}} {{< bg "16" "toastinfo_16" "green" >}} {{< bg "15" "toastinfo_15" "green" >}} {{< bg "14" "toastinfo_14" "green" >}} | `toastinfo_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "postgresql-18-toastinfo" "green" >}} {{< bg "17" "postgresql-17-toastinfo" "green" >}} {{< bg "16" "postgresql-16-toastinfo" "green" >}} {{< bg "15" "postgresql-15-toastinfo" "green" >}} {{< bg "14" "postgresql-14-toastinfo" "green" >}} | `postgresql-$v-toastinfo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [toastinfo_18-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_18-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_18-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_18-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_18-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_18-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [toastinfo_18-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_18-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [toastinfo_18-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_18-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [toastinfo_18-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_18-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_17` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [toastinfo_17-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_17-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_17` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_17-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_17-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_17-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_17-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_17-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_17-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_17` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [toastinfo_17-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_17-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_17` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [toastinfo_17-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_17-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.1 KiB | [postgresql-17-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.8 KiB | [postgresql-17-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_16` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [toastinfo_16-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_16-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_16` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_16-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_16-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_16-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_16-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [toastinfo_16-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_16-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_16` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [toastinfo_16-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_16-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_16` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [toastinfo_16-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_16-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.1 KiB | [postgresql-16-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.8 KiB | [postgresql-16-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [toastinfo_15-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_15-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_15-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_15-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [toastinfo_15-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_15-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_15-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_15-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_15` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [toastinfo_15-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_15-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_15` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [toastinfo_15-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_15-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-15-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-15-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-15-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-15-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [toastinfo_14-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_14-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_14-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_14-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [toastinfo_14-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_14-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_14-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_14-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_14` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [toastinfo_14-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_14-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_14` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [toastinfo_14-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_14-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-14-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-14-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/toastinfo" title="Repository" icon="github" subtitle="github.com/credativ/toastinfo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="toastinfo-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg toastinfo;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install toastinfo;		# install via package name, for the active PG version

pig install toastinfo -v 18;   # install for PG 18
pig install toastinfo -v 17;   # install for PG 17
pig install toastinfo -v 16;   # install for PG 16
pig install toastinfo -v 15;   # install for PG 15
pig install toastinfo -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION toastinfo;
```



## Usage

> [toastinfo: inspect TOAST storage details of varlena columns](https://github.com/credativ/toastinfo)

toastinfo exposes the internal storage form of variable-length (varlena) data types, showing how PostgreSQL stores each datum.

### Functions

**`pg_toastinfo(anyelement)`** -- describes the storage form of a datum:

```sql
SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;

        a         | length  | pg_column_size |              pg_toastinfo              | pg_toastpointer
------------------+---------+----------------+----------------------------------------+-----------------
 null             |       * |              * | null                                   |               *
 default          |       7 |              8 | short inline varlena                   |               *
 external-200     |     200 |            204 | long inline varlena, uncompressed      |               *
 external-10000   |   10000 |          10000 | toasted varlena, uncompressed          |           16427
 extended-10000   |   10000 |            125 | long inline varlena, compressed (pglz) |               *
 extended-1000000 | 1000000 |          11452 | toasted varlena, compressed (pglz)     |           16429
 extended-1000000 | 1000000 |           3936 | toasted varlena, compressed (lz4)      |           16430
```

Possible storage forms:
- `null` -- NULL values
- `ordinary` -- non-varlena datatypes
- `short inline varlena` -- up to 126 bytes (1-byte header)
- `long inline varlena, (un)compressed` -- up to 1GiB (4-byte header)
- `toasted varlena, (un)compressed` -- up to 1GiB stored in TOAST table
- Compressed varlenas show method (`pglz`, `lz4`) on PG14+

**`pg_toastpointer(anyelement)`** -- returns the `chunk_id` OID in the TOAST table, or NULL for non-toasted data:

```sql
SELECT pg_toastpointer(large_column) FROM my_table;
```
