---
title: "toastinfo"
linkTitle: "toastinfo"
description: "show details on toasted datums"
weight: 6360
categories: ["STAT"]
width: full
---

[**toastinfo**](https://github.com/credativ/toastinfo)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6360** | {{< badge content="toastinfo" link="https://github.com/credativ/toastinfo" >}} | {{< ext "toastinfo" >}} | `1.5` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "pg_relusage" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/toastinfo" >}} | `1.4` | {{< bg "18" "toastinfo_18*" "green" >}} {{< bg "17" "toastinfo_17*" "green" >}} {{< bg "16" "toastinfo_16*" "green" >}} {{< bg "15" "toastinfo_15*" "green" >}} {{< bg "14" "toastinfo_14*" "green" >}} {{< bg "13" "toastinfo_13*" "green" >}} | `toastinfo_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/toastinfo" >}} | `1.5` | {{< bg "18" "postgresql-18-toastinfo" "green" >}} {{< bg "17" "postgresql-17-toastinfo" "green" >}} {{< bg "16" "postgresql-16-toastinfo" "green" >}} {{< bg "15" "postgresql-15-toastinfo" "green" >}} {{< bg "14" "postgresql-14-toastinfo" "green" >}} {{< bg "13" "postgresql-13-toastinfo" "green" >}} | `postgresql-$v-toastinfo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5" "toastinfo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "toastinfo_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5" "postgresql-18-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-17-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-16-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-15-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-14-toastinfo : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "postgresql-13-toastinfo : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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
| `toastinfo_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_17-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_17-1.4-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_17` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_17-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_17-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_17-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_17-1.4-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_17-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_17-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [toastinfo_17-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_17-1.4-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_17-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_17-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [toastinfo_17-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_17-1.4-1PIGSTY.el9.aarch64.rpm) |
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
| `toastinfo_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_16-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_16-1.4-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_16` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_16-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_16-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_16-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_16-1.4-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_16-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_16-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_16-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_16-1.4-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [toastinfo_16-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_16-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [toastinfo_16-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_16-1.4-1PIGSTY.el9.aarch64.rpm) |
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
| `toastinfo_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_15-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_15-1.4-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_15-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_15-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [toastinfo_15-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_15-1.4-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [toastinfo_15-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_15-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [toastinfo_15-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_15-1.4-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_15-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_15-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_15-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_15-1.4-1PIGSTY.el9.aarch64.rpm) |
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
| `toastinfo_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_14-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_14-1.4-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_14-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_14-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_14-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_14-1.4-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [toastinfo_14-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_14-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [toastinfo_14-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_14-1.4-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [toastinfo_14-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_14-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_14-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_14-1.4-1PIGSTY.el9.aarch64.rpm) |
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_13` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [toastinfo_13-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_13-1.5-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [toastinfo_13-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_13-1.4-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_13` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [toastinfo_13-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_13-1.5-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [toastinfo_13-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_13-1.4-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_13` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.3 KiB | [toastinfo_13-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_13-1.5-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_13` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_13-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_13-1.4-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_13` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [toastinfo_13-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_13-1.5-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [toastinfo_13-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_13-1.4-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_13` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.2 KiB | [toastinfo_13-1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_13-1.5-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_13` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [toastinfo_13-1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_13-1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.1 KiB | [postgresql-13-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-13-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.1 KiB | [postgresql-13-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-13-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.4 KiB | [postgresql-13-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.3 KiB | [postgresql-13-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.2 KiB | [postgresql-13-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.9 KiB | [postgresql-13-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-13-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/toastinfo" title="Repository" icon="github" subtitle="github.com/credativ/toastinfo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="toastinfo-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get toastinfo; # get toastinfo source code
pig build dep toastinfo; # install build dependencies
pig build pkg toastinfo; # build extension rpm or deb
pig build ext toastinfo; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install toastinfo; # install by extension name, for the current active PG version
pig ext install toastinfo; # install via package alias, for the active PG version
pig ext install toastinfo -v 18;   # install for PG 18
pig ext install toastinfo -v 17;   # install for PG 17
pig ext install toastinfo -v 16;   # install for PG 16
pig ext install toastinfo -v 15;   # install for PG 15
pig ext install toastinfo -v 14;   # install for PG 14
pig ext install toastinfo -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION toastinfo;
```

