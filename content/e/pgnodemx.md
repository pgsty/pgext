---
title: "pgnodemx"
linkTitle: "pgnodemx"
description: "Capture node OS metrics via SQL queries"
weight: 6440
categories: ["STAT"]
width: full
---

[**pgnodemx**](https://github.com/CrunchyData/pgnodemx) : Capture node OS metrics via SQL queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6440** | {{< badge content="pgnodemx" link="https://github.com/CrunchyData/pgnodemx" >}} | {{< ext "pgnodemx" >}} | `1.7` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgsentinel" >}} {{< ext "pgmeminfo" >}} {{< ext "pgfincore" >}} {{< ext "prioritize" >}} {{< ext "pg_buffercache" >}} |
|    **Siblings**   | {{< ext "pg_proctab" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgnodemx` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "pgnodemx_18" "green" >}} {{< bg "17" "pgnodemx_17" "green" >}} {{< bg "16" "pgnodemx_16" "green" >}} {{< bg "15" "pgnodemx_15" "green" >}} {{< bg "14" "pgnodemx_14" "green" >}} | `pgnodemx_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "postgresql-18-pgnodemx" "green" >}} {{< bg "17" "postgresql-17-pgnodemx" "green" >}} {{< bg "16" "postgresql-16-pgnodemx" "green" >}} {{< bg "15" "postgresql-15-pgnodemx" "green" >}} {{< bg "14" "postgresql-14-pgnodemx" "green" >}} | `postgresql-$v-pgnodemx` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_18` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.5 KiB | [pgnodemx_18-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_18-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgnodemx_18-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgnodemx_18-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 37.0 KiB | [pgnodemx_18-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_18-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_18` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [pgnodemx_18-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgnodemx_18-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_18` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.8 KiB | [pgnodemx_18-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_18-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [pgnodemx_18-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgnodemx_18-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.5 KiB | [pgnodemx_18-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_18-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_18` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.2 KiB | [pgnodemx_18-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgnodemx_18-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_18` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.0 KiB | [pgnodemx_18-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_18-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.2 KiB | [pgnodemx_18-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgnodemx_18-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_18` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.6 KiB | [pgnodemx_18-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_18-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_18` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.7 KiB | [pgnodemx_18-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgnodemx_18-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgnodemx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 95.2 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 82.3 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.0 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.4 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 83.3 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.4 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 82.2 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.3 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.4 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 81.7 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 83.6 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 94.9 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.9 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 89.2 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 93.8 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.9 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 88.9 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.5 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 81.8 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.7 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 93.3 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 80.8 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.7 KiB | [postgresql-18-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 93.9 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 81.6 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 93.1 KiB | [postgresql-18-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 80.4 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_17` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.5 KiB | [pgnodemx_17-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_17-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgnodemx_17-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgnodemx_17-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 37.0 KiB | [pgnodemx_17-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_17-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_17` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [pgnodemx_17-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgnodemx_17-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_17` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.8 KiB | [pgnodemx_17-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_17-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [pgnodemx_17-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgnodemx_17-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.5 KiB | [pgnodemx_17-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_17-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_17` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.2 KiB | [pgnodemx_17-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgnodemx_17-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_17` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.0 KiB | [pgnodemx_17-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_17-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.3 KiB | [pgnodemx_17-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgnodemx_17-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_17` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.6 KiB | [pgnodemx_17-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_17-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_17` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.6 KiB | [pgnodemx_17-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgnodemx_17-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgnodemx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 95.4 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 82.2 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.2 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.1 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 83.3 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.3 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 82.2 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.2 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.3 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 81.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 83.6 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 103.1 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 88.9 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 96.1 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.8 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 95.6 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.6 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 81.8 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.6 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 93.3 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 80.7 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.7 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.2 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 81.6 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 93.0 KiB | [postgresql-17-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 80.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.5 KiB | [pgnodemx_16-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_16-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgnodemx_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgnodemx_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 37.0 KiB | [pgnodemx_16-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_16-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [pgnodemx_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgnodemx_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.8 KiB | [pgnodemx_16-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_16-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [pgnodemx_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgnodemx_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.4 KiB | [pgnodemx_16-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_16-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.1 KiB | [pgnodemx_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgnodemx_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_16` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.0 KiB | [pgnodemx_16-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_16-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.2 KiB | [pgnodemx_16-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgnodemx_16-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_16` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.6 KiB | [pgnodemx_16-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_16-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_16` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.7 KiB | [pgnodemx_16-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgnodemx_16-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgnodemx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 95.2 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 82.2 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.1 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 81.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 83.3 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.2 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 82.2 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.3 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.3 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 81.6 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 83.7 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 103.0 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 88.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 96.1 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.7 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 95.6 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.5 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 81.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.6 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 93.3 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 80.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.7 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.1 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 81.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 93.0 KiB | [postgresql-16-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 80.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 38.7 KiB | [pgnodemx_15-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_15-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.1 KiB | [pgnodemx_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgnodemx_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 38.1 KiB | [pgnodemx_15-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_15-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pgnodemx_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgnodemx_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 38.0 KiB | [pgnodemx_15-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_15-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.9 KiB | [pgnodemx_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgnodemx_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 37.5 KiB | [pgnodemx_15-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_15-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pgnodemx_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgnodemx_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_15` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.0 KiB | [pgnodemx_15-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_15-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.1 KiB | [pgnodemx_15-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgnodemx_15-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_15` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 37.7 KiB | [pgnodemx_15-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_15-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_15` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pgnodemx_15-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgnodemx_15-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgnodemx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 96.6 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 83.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 85.7 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 95.4 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 82.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 84.6 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 96.9 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 83.7 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 85.8 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.8 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 83.0 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 84.7 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.0 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 90.6 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 98.1 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.9 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 89.4 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 97.9 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 96.6 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.4 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.3 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 82.4 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 89.5 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 95.8 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 83.1 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.1 KiB | [postgresql-15-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 82.2 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 38.6 KiB | [pgnodemx_14-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_14-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pgnodemx_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgnodemx_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 38.0 KiB | [pgnodemx_14-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_14-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pgnodemx_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgnodemx_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 37.9 KiB | [pgnodemx_14-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_14-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.8 KiB | [pgnodemx_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgnodemx_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 37.4 KiB | [pgnodemx_14-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_14-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pgnodemx_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgnodemx_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_14` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 37.9 KiB | [pgnodemx_14-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_14-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pgnodemx_14-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgnodemx_14-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_14` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 37.5 KiB | [pgnodemx_14-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_14-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_14` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.0 KiB | [pgnodemx_14-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgnodemx_14-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgnodemx` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 95.4 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 83.3 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 85.3 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.2 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 82.3 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 84.4 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.5 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 83.5 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 85.5 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.3 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 82.7 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 84.6 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 103.4 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 90.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 97.7 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.4 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 89.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 97.5 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.2 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.0 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 88.9 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.3 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 82.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 89.1 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.6 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 82.8 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 93.6 KiB | [postgresql-14-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_2.0.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 81.8 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgnodemx" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgnodemx" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgnodemx-1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgnodemx;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgnodemx;		# install via package name, for the active PG version

pig install pgnodemx -v 18;   # install for PG 18
pig install pgnodemx -v 17;   # install for PG 17
pig install pgnodemx -v 16;   # install for PG 16
pig install pgnodemx -v 15;   # install for PG 15
pig install pgnodemx -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgnodemx';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgnodemx;
```



## Usage

> [pgnodemx: SQL functions to access node OS metrics from PostgreSQL](https://github.com/CrunchyData/pgnodemx)

pgnodemx provides SQL access to OS-level metrics including cgroup statistics, `/proc` filesystem data, and system information. Requires `pg_monitor` role membership.

### cgroup Functions

```sql
-- Check cgroup support status
SELECT current_setting('pgnodemx.cgroup_enabled');
SELECT cgroup_mode();  -- 'legacy', 'unified', 'hybrid', or 'disabled'

-- Read cgroup scalar values
SELECT cgroup_scalar_bigint('memory.current');
SELECT cgroup_scalar_float8('cpu.uclamp.max');
SELECT cgroup_scalar_text('cgroup.type');

-- Read cgroup key-value files
SELECT * FROM cgroup_setof_kv('memory.stat');
SELECT * FROM cgroup_setof_kv('cpu.stat');

-- Read cgroup nested key-value files
SELECT * FROM cgroup_setof_nkv('memory.pressure');
SELECT * FROM cgroup_setof_nkv('cpu.pressure');

-- Get cgroup paths and process count
SELECT * FROM cgroup_path();
SELECT cgroup_process_count();
```

### /proc Functions

```sql
SELECT * FROM proc_diskstats();       -- /proc/diskstats
SELECT * FROM proc_mountinfo();       -- /proc/self/mountinfo
SELECT * FROM proc_meminfo();         -- /proc/meminfo
SELECT * FROM proc_network_stats();   -- /proc/self/net/dev
SELECT * FROM proc_pid_io();          -- /proc/<pid>/io for all PG processes
SELECT * FROM proc_pid_cmdline();     -- command line for all PG processes
SELECT * FROM proc_pid_stat();        -- /proc/<pid>/stat for all PG processes
SELECT * FROM proc_cputime();         -- first line of /proc/stat
SELECT * FROM proc_loadavg();         -- /proc/loadavg
```

### System Information

```sql
SELECT * FROM fsinfo('/pgdata');      -- filesystem info for a path
SELECT fips_mode();                   -- OpenSSL FIPS mode status
SELECT openssl_version();             -- OpenSSL version
SELECT exec_path();                   -- current PostgreSQL executable path
SELECT kpages_to_bytes(1024);         -- convert kernel pages to bytes
SELECT * FROM stat_file('/path');     -- file uid, gid, mode info
```

### Environment Variables

```sql
SELECT envvar_text('PGDATA');
SELECT envvar_bigint('PGPORT');
```

### Kubernetes DownwardAPI

```sql
SELECT * FROM kdapi_setof_kv('labels');
SELECT kdapi_scalar_bigint('cpu_limit');
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgnodemx.cgroup_enabled` | `on` | Enable cgroup functions |
| `pgnodemx.containerized` | `off` | Force containerized cgroup paths |
| `pgnodemx.cgrouproot` | `/sys/fs/cgroup` | cgroup mount location |
| `pgnodemx.kdapi_enabled` | `on` | Enable Kubernetes DownwardAPI |
| `pgnodemx.kdapi_path` | `/etc/podinfo` | DownwardAPI files location |
