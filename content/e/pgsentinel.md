---
title: "pgsentinel"
linkTitle: "pgsentinel"
description: "active session history"
weight: 6410
categories: ["STAT"]
width: full
---

[**pgsentinel**](https://github.com/pgsentinel/pgsentinel) : active session history


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6410** | {{< badge content="pgsentinel" link="https://github.com/pgsentinel/pgsentinel" >}} | {{< ext "pgsentinel" >}} | `1.4.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_wait_sampling" >}} {{< ext "bgw_replstatus" >}} {{< ext "pg_profile" >}} {{< ext "pg_proctab" >}} {{< ext "powa" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgsentinel` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.0` | {{< bg "18" "pgsentinel_18" "green" >}} {{< bg "17" "pgsentinel_17" "green" >}} {{< bg "16" "pgsentinel_16" "green" >}} {{< bg "15" "pgsentinel_15" "green" >}} {{< bg "14" "pgsentinel_14" "green" >}} {{< bg "13" "pgsentinel_13" "green" >}} | `pgsentinel_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.0` | {{< bg "18" "postgresql-18-pgsentinel" "green" >}} {{< bg "17" "postgresql-17-pgsentinel" "green" >}} {{< bg "16" "postgresql-16-pgsentinel" "green" >}} {{< bg "15" "postgresql-15-pgsentinel" "green" >}} {{< bg "14" "postgresql-14-pgsentinel" "green" >}} {{< bg "13" "postgresql-13-pgsentinel" "green" >}} | `postgresql-$v-pgsentinel` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.4.0" "pgsentinel_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4.0" "postgresql-18-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-17-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-16-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-15-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.0" "postgresql-14-pgsentinel : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-13-pgsentinel : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.1 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_18-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_18-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.4 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_18-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.5 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.9 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_18-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_18-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsentinel_18-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.0 KiB | [pgsentinel_18-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_18-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsentinel_18-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [pgsentinel_18-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsentinel_18-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.4 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.5 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.3 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.4 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.3 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.5 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.4 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.5 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 45.5 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.1 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 44.2 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 43.2 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.3 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.6 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.4 KiB | [postgresql-18-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.1 KiB | [postgresql-18-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.2 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_17-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_17-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.5 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_17-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_17-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_17-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [pgsentinel_17-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_17-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_17-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_17-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.4 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.4 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.6 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.4 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.5 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.3 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.6 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 53.6 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.4 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.3 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 51.4 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.4 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.7 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.4 KiB | [postgresql-17-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.1 KiB | [postgresql-17-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.2 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_16-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.2 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_16-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.5 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_16-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.0 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_16-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_16-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.5 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.7 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [pgsentinel_16-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_16-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_16-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_16-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.4 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.5 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.3 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.5 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.3 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.4 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.6 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 53.4 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.3 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.1 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 51.3 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.4 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.7 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.4 KiB | [postgresql-16-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.0 KiB | [postgresql-16-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.2 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_15-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.8 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_15-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.7 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_15-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_15-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.5 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.1 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_15-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [pgsentinel_15-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_15-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.6 KiB | [pgsentinel_15-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_15-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.9 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.5 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.3 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.5 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.4 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.7 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.4 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.6 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 53.4 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.2 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.1 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 51.1 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.5 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.7 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.5 KiB | [postgresql-15-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.2 KiB | [postgresql-15-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.1 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_14-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.7 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_14-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.3 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.7 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_14-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.6 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_14-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_14-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.6 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.4.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.0 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.3.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [pgsentinel_14-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_14-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsentinel_14-1.4.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.4.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [pgsentinel_14-1.3.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.3.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgsentinel_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.2 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.3 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.0 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.2 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.1 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.3 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.1 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.3 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.9 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.6 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.6 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.6 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.2 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.4 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.2 KiB | [postgresql-14-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 41.8 KiB | [postgresql-14-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsentinel_13` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_13-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsentinel_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_13` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_13-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.7 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsentinel_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_13` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.6 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_13-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsentinel_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_13` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_13-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsentinel_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_13` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.1 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsentinel_13-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgsentinel_13-1.2.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsentinel_13` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [pgsentinel_13-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsentinel_13-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `pgsentinel_13` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgsentinel_13-1.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgsentinel` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.3 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.5 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.3 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.4 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.2 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.8 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgsentinel` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 41.9 KiB | [postgresql-13-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsentinel/pgsentinel" title="Repository" icon="github" subtitle="github.com/pgsentinel/pgsentinel" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsentinel-1.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgsentinel;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgsentinel;		# install via package name, for the active PG version

pig install pgsentinel -v 18;   # install for PG 18
pig install pgsentinel -v 17;   # install for PG 17
pig install pgsentinel -v 16;   # install for PG 16
pig install pgsentinel -v 15;   # install for PG 15
pig install pgsentinel -v 14;   # install for PG 14
pig install pgsentinel -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgsentinel';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgsentinel;
```
