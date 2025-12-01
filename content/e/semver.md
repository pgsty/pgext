---
title: "semver"
linkTitle: "semver"
description: "Semantic version data type"
weight: 3510
categories: ["TYPE"]
width: full
---

[**pg_semver**](https://github.com/theory/pg-semver) : Semantic version data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3510** | {{< badge content="semver" link="https://github.com/theory/pg-semver" >}} | {{< ext "semver" "pg_semver" >}} | `0.40.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "ltree" >}} {{< ext "citext" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.40.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_semver` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.40.0` | {{< bg "18" "semver_18*" "green" >}} {{< bg "17" "semver_17*" "green" >}} {{< bg "16" "semver_16*" "green" >}} {{< bg "15" "semver_15*" "green" >}} {{< bg "14" "semver_14*" "green" >}} {{< bg "13" "semver_13*" "green" >}} | `semver_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.40.0` | {{< bg "18" "postgresql-18-semver" "green" >}} {{< bg "17" "postgresql-17-semver" "green" >}} {{< bg "16" "postgresql-16-semver" "green" >}} {{< bg "15" "postgresql-15-semver" "green" >}} {{< bg "14" "postgresql-14-semver" "green" >}} {{< bg "13" "postgresql-13-semver" "green" >}} | `postgresql-$v-semver` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 5" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.40.0" "semver_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.40.0" "semver_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.40.0" "postgresql-18-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-17-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-16-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-15-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-14-semver : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.40.0" "postgresql-13-semver : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_18` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [semver_18-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_18-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.9 KiB | [semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/semver_18-0.40.0-1PGDG.rhel8.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [semver_18-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_18-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.5 KiB | [semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/semver_18-0.40.0-1PGDG.rhel8.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [semver_18-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_18-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.9 KiB | [semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/semver_18-0.40.0-1PGDG.rhel9.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.5 KiB | [semver_18-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_18-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.5 KiB | [semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/semver_18-0.40.0-1PGDG.rhel9.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [semver_18-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_18-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_18-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/semver_18-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_18` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.9 KiB | [semver_18-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_18-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_18` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_18-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/semver_18-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-18-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-18-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.2 KiB | [postgresql-18-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.4 KiB | [postgresql-18-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.2 KiB | [postgresql-18-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.8 KiB | [postgresql-18-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.3 KiB | [postgresql-18-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-18-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-18-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_17` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [semver_17-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_17-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_17` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/semver_17-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_17` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [semver_17-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_17-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_17` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.8 KiB | [semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/semver_17-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_17` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [semver_17-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_17-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_17` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/semver_17-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_17` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.5 KiB | [semver_17-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_17-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_17` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/semver_17-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_17` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [semver_17-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_17-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_17` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_17-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/semver_17-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_17` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.9 KiB | [semver_17-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_17-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_17` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_17-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/semver_17-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-17-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-17-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.2 KiB | [postgresql-17-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-17-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.9 KiB | [postgresql-17-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.4 KiB | [postgresql-17-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.2 KiB | [postgresql-17-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-17-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-17-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_16` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [semver_16-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_16-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_16` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/semver_16-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_16` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.0 KiB | [semver_16-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_16-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_16` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.8 KiB | [semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/semver_16-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_16` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.6 KiB | [semver_16-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_16-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_16` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/semver_16-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_16` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.5 KiB | [semver_16-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_16-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_16` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/semver_16-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_16` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.7 KiB | [semver_16-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_16-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_16` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [semver_16-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/semver_16-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_16` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.9 KiB | [semver_16-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_16-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_16` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.4 KiB | [semver_16-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/semver_16-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-16-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-16-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.2 KiB | [postgresql-16-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-16-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-16-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-16-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.9 KiB | [postgresql-16-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.5 KiB | [postgresql-16-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.2 KiB | [postgresql-16-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-16-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-16-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_15` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.4 KiB | [semver_15-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_15-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_15` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_15` | `0.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [semver_15-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_15` | `0.31.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [semver_15-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/semver_15-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_15` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.1 KiB | [semver_15-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_15-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_15` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_15` | `0.32.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [semver_15-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/semver_15-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_15` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [semver_15-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_15-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_15` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_15` | `0.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [semver_15-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_15` | `0.31.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [semver_15-0.31.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/semver_15-0.31.2-1.rhel9.x86_64.rpm) |
| `semver_15` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.8 KiB | [semver_15-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_15-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_15` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_15` | `0.32.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [semver_15-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/semver_15-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_15` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.1 KiB | [semver_15-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_15-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_15` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_15-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/semver_15-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_15` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [semver_15-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_15-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_15` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_15-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/semver_15-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-15-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.2 KiB | [postgresql-15-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-15-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-15-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-15-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.1 KiB | [postgresql-15-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.6 KiB | [postgresql-15-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-15-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-15-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_14` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.4 KiB | [semver_14-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_14-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_14` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.3 KiB | [semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_14` | `0.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [semver_14-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_14` | `0.31.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.4 KiB | [semver_14-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_14` | `0.31.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.8 KiB | [semver_14-0.31.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/semver_14-0.31.1-2.rhel8.x86_64.rpm) |
| `semver_14` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.1 KiB | [semver_14-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_14-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_14` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_14` | `0.32.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [semver_14-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/semver_14-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_14` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [semver_14-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_14-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_14` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_14` | `0.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [semver_14-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/semver_14-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_14` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.8 KiB | [semver_14-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_14-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_14` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.0 KiB | [semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_14` | `0.32.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [semver_14-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/semver_14-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_14` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.1 KiB | [semver_14-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_14-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_14` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_14-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/semver_14-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_14` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [semver_14-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_14-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_14` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_14-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/semver_14-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.0 KiB | [postgresql-14-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.2 KiB | [postgresql-14-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-14-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-14-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-14-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.0 KiB | [postgresql-14-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.6 KiB | [postgresql-14-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.1 KiB | [postgresql-14-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-14-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `semver_13` | `0.40.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.3 KiB | [semver_13-0.40.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/semver_13-0.40.0-1PIGSTY.el8.x86_64.rpm) |
| `semver_13` | `0.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.2 KiB | [semver_13-0.32.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.32.1-1PGDG.rhel8.x86_64.rpm) |
| `semver_13` | `0.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [semver_13-0.32.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.32.0-1.rhel8.x86_64.rpm) |
| `semver_13` | `0.31.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.4 KiB | [semver_13-0.31.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.31.2-1.rhel8.x86_64.rpm) |
| `semver_13` | `0.31.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [semver_13-0.31.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/semver_13-0.31.1-1.rhel8.x86_64.rpm) |
| `semver_13` | `0.40.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.1 KiB | [semver_13-0.40.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/semver_13-0.40.0-1PIGSTY.el8.aarch64.rpm) |
| `semver_13` | `0.32.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.9 KiB | [semver_13-0.32.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/semver_13-0.32.1-1PGDG.rhel8.aarch64.rpm) |
| `semver_13` | `0.32.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [semver_13-0.32.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/semver_13-0.32.0-1.rhel8.aarch64.rpm) |
| `semver_13` | `0.40.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.0 KiB | [semver_13-0.40.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/semver_13-0.40.0-1PIGSTY.el9.x86_64.rpm) |
| `semver_13` | `0.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [semver_13-0.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/semver_13-0.32.1-1PGDG.rhel9.x86_64.rpm) |
| `semver_13` | `0.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.2 KiB | [semver_13-0.32.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/semver_13-0.32.0-1.rhel9.x86_64.rpm) |
| `semver_13` | `0.40.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.8 KiB | [semver_13-0.40.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/semver_13-0.40.0-1PIGSTY.el9.aarch64.rpm) |
| `semver_13` | `0.32.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [semver_13-0.32.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/semver_13-0.32.1-1PGDG.rhel9.aarch64.rpm) |
| `semver_13` | `0.32.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.5 KiB | [semver_13-0.32.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/semver_13-0.32.0-1.rhel9.aarch64.rpm) |
| `semver_13` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.1 KiB | [semver_13-0.40.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/semver_13-0.40.0-1PIGSTY.el10.x86_64.rpm) |
| `semver_13` | `0.40.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.8 KiB | [semver_13-0.40.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/semver_13-0.40.0-1PGDG.rhel10.x86_64.rpm) |
| `semver_13` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.2 KiB | [semver_13-0.40.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/semver_13-0.40.0-1PIGSTY.el10.aarch64.rpm) |
| `semver_13` | `0.40.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 27.7 KiB | [semver_13-0.40.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/semver_13-0.40.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-semver` | `0.40.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.5 KiB | [postgresql-13-semver_0.40.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-semver` | `0.40.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 38.3 KiB | [postgresql-13-semver_0.40.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-semver` | `0.40.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.3 KiB | [postgresql-13-semver_0.40.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg13+1_amd64.deb) |
| `postgresql-13-semver` | `0.40.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 38.3 KiB | [postgresql-13-semver_0.40.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg13+1_arm64.deb) |
| `postgresql-13-semver` | `0.40.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.0 KiB | [postgresql-13-semver_0.40.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-semver` | `0.40.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-13-semver_0.40.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-semver` | `0.40.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.2 KiB | [postgresql-13-semver_0.40.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-semver` | `0.40.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.2 KiB | [postgresql-13-semver_0.40.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-semver/postgresql-13-semver_0.40.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pg-semver" title="Repository" icon="github" subtitle="github.com/theory/pg-semver" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-semver-0.40.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_semver;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_semver;		# install via package name, for the active PG version
pig install semver;		# install by extension name, for the current active PG version

pig install semver -v 18;   # install for PG 18
pig install semver -v 17;   # install for PG 17
pig install semver -v 16;   # install for PG 16
pig install semver -v 15;   # install for PG 15
pig install semver -v 14;   # install for PG 14
pig install semver -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION semver;
```
