---
title: "plr"
linkTitle: "plr"
description: "load R interpreter and execute R script from within a database"
weight: 3100
categories: ["LANG"]
width: full
---

[**plr**](https://github.com/postgres-plr/plr) : load R interpreter and execute R script from within a database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3100** | {{< badge content="plr" link="https://github.com/postgres-plr/plr" >}} | {{< ext "plr" >}} | `8.4.8` | {{< category "LANG" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pgml" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "pllua" >}} |

> [!Note] missing el10.x86_64


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plr` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8` | {{< bg "18" "plr_18" "green" >}} {{< bg "17" "plr_17" "green" >}} {{< bg "16" "plr_16" "green" >}} {{< bg "15" "plr_15" "green" >}} {{< bg "14" "plr_14" "green" >}} | `plr_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8` | {{< bg "18" "postgresql-18-plr" "green" >}} {{< bg "17" "postgresql-17-plr" "green" >}} {{< bg "16" "postgresql-16-plr" "green" >}} {{< bg "15" "postgresql-15-plr" "green" >}} {{< bg "14" "postgresql-14-plr" "green" >}} | `postgresql-$v-plr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "postgresql-14-plr : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_18` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.4 KiB | [plr_18-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_18` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.5 KiB | [plr_18-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_18-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.0 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_18` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.9 KiB | [plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.4 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.7 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.8 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.1 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.1 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.5 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.6 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.6 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.7 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.6 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.5 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.4 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.4 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.1 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.5 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.8 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.3 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.4 KiB | [postgresql-18-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.3 KiB | [postgresql-18-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_17` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_17` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.0 KiB | [plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.7 KiB | [plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_17-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_17-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.7 KiB | [plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_17` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.8 KiB | [plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.2 KiB | [plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.2 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.8 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.8 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.7 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.1 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.1 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 135.9 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 135.9 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.6 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.4 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.5 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.5 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.7 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.6 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.3 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.2 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.2 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.1 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.5 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.5 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.4 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.3 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.3 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.0 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-17-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.0 KiB | [postgresql-17-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_16` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_16` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.0 KiB | [plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.7 KiB | [plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_16-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_16-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.6 KiB | [plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.6 KiB | [plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_16` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.8 KiB | [plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.8 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.2 KiB | [plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.1 KiB | [plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.2 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.0 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.7 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.6 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.2 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.2 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.0 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 135.8 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 135.8 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.4 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.4 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.6 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.7 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.7 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.1 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.1 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.1 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.6 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.4 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.4 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.4 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_15` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.3 KiB | [plr_15-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [plr_15-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_15` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.5 KiB | [plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.4 KiB | [plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 167.4 KiB | [plr_15-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.9 KiB | [plr_15-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_15-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.9 KiB | [plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.1 KiB | [plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.0 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_15` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.4 KiB | [plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 167.8 KiB | [plr_15-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.1 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.1 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.3 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.3 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.4 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.8 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.8 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.2 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.5 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.4 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-15-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-15-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_14` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.4 KiB | [plr_14-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [plr_14-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_14` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.5 KiB | [plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.4 KiB | [plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 166.7 KiB | [plr_14-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 166.5 KiB | [plr_14-8.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.3-1.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.9 KiB | [plr_14-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_14-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.9 KiB | [plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.1 KiB | [plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_14` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.3 KiB | [plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 167.4 KiB | [plr_14-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.3 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.1 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.4 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.1 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.5 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.4 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.4 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-14-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.3 KiB | [postgresql-14-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgres-plr/plr" title="Repository" icon="github" subtitle="github.com/postgres-plr/plr" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plr;		# install via package name, for the active PG version

pig install plr -v 18;   # install for PG 18
pig install plr -v 17;   # install for PG 17
pig install plr -v 16;   # install for PG 16
pig install plr -v 15;   # install for PG 15
pig install plr -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plr;
```




## Usage

> [plr: load R interpreter and execute R script from within a database](https://github.com/postgres-plr/plr)

`plr` enables writing PostgreSQL functions in the R programming language, providing full access to R's statistical and data analysis capabilities.

```sql
CREATE EXTENSION plr;
```

### Create Functions

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (arg1 > arg2)
  return(arg1)
else
  return(arg2)
' LANGUAGE plr STRICT;

SELECT r_max(10, 20);  -- 20
```

With named arguments:

```sql
CREATE OR REPLACE FUNCTION sd(vals float8[]) RETURNS float AS '
sd(vals)
' LANGUAGE plr STRICT;

SELECT sd(ARRAY[1.0, 2.0, 3.0, 4.0, 5.0]);
```

### Argument Handling

- Arguments are available as `arg1`, `arg2`, ... or by named parameter
- NULL arguments become R `NA` values (unless function is `STRICT`)
- Composite types (rows) are passed as R data.frames
- Arrays are passed as R vectors

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (is.null(arg1) && is.null(arg2))
  return(NULL)
if (is.null(arg1))
  return(arg2)
if (is.null(arg2))
  return(arg1)
if (arg1 > arg2)
  return(arg1)
return(arg2)
' LANGUAGE plr;
```

### Database Access via SPI

```sql
CREATE OR REPLACE FUNCTION test_spi(text) RETURNS SETOF record AS '
pg.spi.exec(arg1)
' LANGUAGE plr;

SELECT * FROM test_spi('SELECT oid, typname FROM pg_type LIMIT 5')
  AS t(oid oid, typname name);
```

Prepared statements:

```sql
-- Prepare
sp <<- pg.spi.prepare('SELECT * FROM pg_type WHERE typname = $1', c(NAMEOID))
-- Execute
pg.spi.execp(sp, list('text'))
```

### Set-Returning Functions

Return a data.frame for set-returning functions:

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:arg1
' LANGUAGE plr;

SELECT * FROM get_numbers(5);
```

### Window Functions

```sql
CREATE OR REPLACE FUNCTION r_regr_slope(float8, float8, int)
RETURNS float8 AS '
slope <- NA
y <- farg1
x <- farg2
if (fnumrows == arg3 + 1L)
  try(slope <- lm(y ~ x)$coefficients[2])
return(slope)
' LANGUAGE plr WINDOW;
```

Window functions receive `farg1..fargN` (vectors of values in the window frame), `fnumrows` (frame size), and `prownum` (current row position in partition).

### Global Variables

Persist data across function calls using R's global environment:

```sql
CREATE OR REPLACE FUNCTION set_state(key text, val text) RETURNS void AS '
assign(arg1, arg2, env=.GlobalEnv)
' LANGUAGE plr;
```

### Useful Support Functions

```sql
SELECT load_r_typenames();  -- Load type OID variables
SELECT * FROM r_typenames(); -- List available type OIDs
SELECT plr_version();        -- PL/R version
```

### Trigger Functions

PL/R supports trigger functions with access to `pg.tg.name`, `pg.tg.relname`, `pg.tg.when`, `pg.tg.level`, `pg.tg.op`, `pg.tg.new`, and `pg.tg.old`.
