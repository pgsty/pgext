---
title: "icu_ext"
linkTitle: "icu_ext"
description: "Access ICU functions"
weight: 4240
categories: ["UTIL"]
width: full
---

[**icu_ext**](https://github.com/dverite/icu_ext) : Access ICU functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4240** | {{< badge content="icu_ext" link="https://github.com/dverite/icu_ext" >}} | {{< ext "icu_ext" >}} | `1.10.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgpcre" >}} {{< ext "pg_xenophile" >}} {{< ext "unaccent" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.10.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `icu_ext` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.0` | {{< bg "18" "icu_ext_18" "green" >}} {{< bg "17" "icu_ext_17" "green" >}} {{< bg "16" "icu_ext_16" "green" >}} {{< bg "15" "icu_ext_15" "green" >}} {{< bg "14" "icu_ext_14" "green" >}} {{< bg "13" "icu_ext_13" "green" >}} | `icu_ext_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.10.0` | {{< bg "18" "postgresql-18-icu-ext" "green" >}} {{< bg "17" "postgresql-17-icu-ext" "green" >}} {{< bg "16" "postgresql-16-icu-ext" "green" >}} {{< bg "15" "postgresql-15-icu-ext" "green" >}} {{< bg "14" "postgresql-14-icu-ext" "green" >}} {{< bg "13" "postgresql-13-icu-ext" "green" >}} | `postgresql-$v-icu-ext` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10.0" "icu_ext_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.10.0" "postgresql-18-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-17-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-16-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-15-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-14-icu-ext : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.10.0" "postgresql-13-icu-ext : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_18` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 51.0 KiB | [icu_ext_18-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_18-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.4 KiB | [icu_ext_18-1.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/icu_ext_18-1.10.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [icu_ext_18-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_18-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_18` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.1 KiB | [icu_ext_18-1.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/icu_ext_18-1.10.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_18` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 49.6 KiB | [icu_ext_18-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_18-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.6 KiB | [icu_ext_18-1.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/icu_ext_18-1.10.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.9 KiB | [icu_ext_18-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_18-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_18` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.7 KiB | [icu_ext_18-1.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/icu_ext_18-1.10.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_18` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.0 KiB | [icu_ext_18-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_18-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.2 KiB | [icu_ext_18-1.10.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/icu_ext_18-1.10.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_18` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.2 KiB | [icu_ext_18-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_18-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_18` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.6 KiB | [icu_ext_18-1.10.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/icu_ext_18-1.10.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.5 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.0 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.3 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.5 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 95.4 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 92.8 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.5 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.7 KiB | [postgresql-18-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-18-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_17` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 51.0 KiB | [icu_ext_17-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_17-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_17` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.4 KiB | [icu_ext_17-1.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/icu_ext_17-1.9.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_17` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [icu_ext_17-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_17-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_17` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.2 KiB | [icu_ext_17-1.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/icu_ext_17-1.9.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_17` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 49.6 KiB | [icu_ext_17-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_17-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_17` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.7 KiB | [icu_ext_17-1.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/icu_ext_17-1.9.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_17` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.9 KiB | [icu_ext_17-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_17-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_17` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.0 KiB | [icu_ext_17-1.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/icu_ext_17-1.9.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_17` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.0 KiB | [icu_ext_17-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_17-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_17` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.0 KiB | [icu_ext_17-1.9.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/icu_ext_17-1.9.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_17` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.2 KiB | [icu_ext_17-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_17-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_17` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.7 KiB | [icu_ext_17-1.9.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/icu_ext_17-1.9.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.3 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.3 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.5 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.4 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.1 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.4 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.4 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.7 KiB | [postgresql-17-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-17-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_16` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 50.9 KiB | [icu_ext_16-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_16-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_16` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.4 KiB | [icu_ext_16-1.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/icu_ext_16-1.9.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_16` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [icu_ext_16-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_16-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_16` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.2 KiB | [icu_ext_16-1.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/icu_ext_16-1.9.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_16` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 49.6 KiB | [icu_ext_16-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_16-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_16` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.7 KiB | [icu_ext_16-1.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/icu_ext_16-1.9.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_16` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.9 KiB | [icu_ext_16-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_16-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_16` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.9 KiB | [icu_ext_16-1.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/icu_ext_16-1.9.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_16` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.0 KiB | [icu_ext_16-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_16-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_16` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.0 KiB | [icu_ext_16-1.9.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/icu_ext_16-1.9.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_16` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.3 KiB | [icu_ext_16-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_16-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_16` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.7 KiB | [icu_ext_16-1.9.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/icu_ext_16-1.9.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.2 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.2 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.3 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.4 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.6 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.8 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.4 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.6 KiB | [postgresql-16-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-16-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_15` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 51.0 KiB | [icu_ext_15-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_15-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_15` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.5 KiB | [icu_ext_15-1.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/icu_ext_15-1.9.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_15` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.7 KiB | [icu_ext_15-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_15-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_15` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.2 KiB | [icu_ext_15-1.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/icu_ext_15-1.9.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_15` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.0 KiB | [icu_ext_15-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_15-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_15` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.8 KiB | [icu_ext_15-1.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/icu_ext_15-1.9.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_15` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 48.0 KiB | [icu_ext_15-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_15-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_15` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.2 KiB | [icu_ext_15-1.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/icu_ext_15-1.9.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_15` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 49.9 KiB | [icu_ext_15-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_15-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_15` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.1 KiB | [icu_ext_15-1.9.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/icu_ext_15-1.9.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_15` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.3 KiB | [icu_ext_15-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_15-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_15` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.7 KiB | [icu_ext_15-1.9.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/icu_ext_15-1.9.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.3 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.2 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.3 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.0 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.2 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.4 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.3 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.3 KiB | [postgresql-15-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-15-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_14` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 51.2 KiB | [icu_ext_14-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_14-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_14` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.6 KiB | [icu_ext_14-1.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/icu_ext_14-1.9.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_14` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.9 KiB | [icu_ext_14-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_14-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_14` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.3 KiB | [icu_ext_14-1.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/icu_ext_14-1.9.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_14` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 50.0 KiB | [icu_ext_14-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_14-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_14` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.0 KiB | [icu_ext_14-1.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/icu_ext_14-1.9.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_14` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 48.2 KiB | [icu_ext_14-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_14-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_14` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.3 KiB | [icu_ext_14-1.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/icu_ext_14-1.9.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_14` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 49.9 KiB | [icu_ext_14-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_14-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_14` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.2 KiB | [icu_ext_14-1.9.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/icu_ext_14-1.9.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_14` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.4 KiB | [icu_ext_14-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_14-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_14` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.8 KiB | [icu_ext_14-1.9.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/icu_ext_14-1.9.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.8 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.3 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.8 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.5 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.1 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.1 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.0 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.6 KiB | [postgresql-14-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-14-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `icu_ext_13` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 50.1 KiB | [icu_ext_13-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/icu_ext_13-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `icu_ext_13` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [icu_ext_13-1.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/icu_ext_13-1.9.0-1PGDG.rhel8.x86_64.rpm) |
| `icu_ext_13` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.8 KiB | [icu_ext_13-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/icu_ext_13-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `icu_ext_13` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.3 KiB | [icu_ext_13-1.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/icu_ext_13-1.9.0-1PGDG.rhel8.aarch64.rpm) |
| `icu_ext_13` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 49.9 KiB | [icu_ext_13-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/icu_ext_13-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `icu_ext_13` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.0 KiB | [icu_ext_13-1.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/icu_ext_13-1.9.0-1PGDG.rhel9.x86_64.rpm) |
| `icu_ext_13` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 48.2 KiB | [icu_ext_13-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/icu_ext_13-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `icu_ext_13` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.3 KiB | [icu_ext_13-1.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/icu_ext_13-1.9.0-1PGDG.rhel9.aarch64.rpm) |
| `icu_ext_13` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 50.0 KiB | [icu_ext_13-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/icu_ext_13-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `icu_ext_13` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.2 KiB | [icu_ext_13-1.9.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/icu_ext_13-1.9.0-1PGDG.rhel10.x86_64.rpm) |
| `icu_ext_13` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 48.4 KiB | [icu_ext_13-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/icu_ext_13-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `icu_ext_13` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.9 KiB | [icu_ext_13-1.9.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/icu_ext_13-1.9.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-icu-ext` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.4 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg12+1_amd64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 92.3 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg12+1_arm64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.6 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg13+1_amd64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 92.4 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg13+1_arm64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.7 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.8 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.6 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-icu-ext` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 92.7 KiB | [postgresql-13-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/icu-ext/postgresql-13-icu-ext_1.10.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dverite/icu_ext" title="Repository" icon="github" subtitle="github.com/dverite/icu_ext" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="icu_ext-1.10.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg icu_ext;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install icu_ext;		# install via package name, for the active PG version

pig install icu_ext -v 18;   # install for PG 18
pig install icu_ext -v 17;   # install for PG 17
pig install icu_ext -v 16;   # install for PG 16
pig install icu_ext -v 15;   # install for PG 15
pig install icu_ext -v 14;   # install for PG 14
pig install icu_ext -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION icu_ext;
```
