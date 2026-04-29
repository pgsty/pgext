---
title: "pgtt"
linkTitle: "pgtt"
description: "Extension to add Global Temporary Tables feature to PostgreSQL"
weight: 9110
categories: ["SIM"]
width: full
---

[**pgtt**](https://github.com/darold/pgtt) : Extension to add Global Temporary Tables feature to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9110** | {{< badge content="pgtt" link="https://github.com/darold/pgtt" >}} | {{< ext "pgtt" >}} | `4.4` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgtt_schema` |
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "orafce" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "periods" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgtt` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.4` | {{< bg "18" "pgtt_18" "green" >}} {{< bg "17" "pgtt_17" "green" >}} {{< bg "16" "pgtt_16" "green" >}} {{< bg "15" "pgtt_15" "green" >}} {{< bg "14" "pgtt_14" "green" >}} | `pgtt_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.4` | {{< bg "18" "postgresql-18-pgtt" "green" >}} {{< bg "17" "postgresql-17-pgtt" "green" >}} {{< bg "16" "postgresql-16-pgtt" "green" >}} {{< bg "15" "postgresql-15-pgtt" "green" >}} {{< bg "14" "postgresql-14-pgtt" "green" >}} | `postgresql-$v-pgtt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 10" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 9" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 8" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.4" "pgtt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.4" "pgtt_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 4.4" "postgresql-18-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-17-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-16-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-15-pgtt : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.4" "postgresql-14-pgtt : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtt_18` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.7 KiB | [pgtt_18-4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgtt_18-4.4-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_18` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.2 KiB | [pgtt_18-4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgtt_18-4.4-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_18` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.7 KiB | [pgtt_18-4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgtt_18-4.4-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_18` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [pgtt_18-4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgtt_18-4.4-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_18` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [pgtt_18-4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgtt_18-4.4-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [pgtt_18-4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgtt_18-4.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgtt` | `4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.3 KiB | [postgresql-18-pgtt_4.4-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgtt` | `4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-18-pgtt_4.4-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgtt` | `4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.8 KiB | [postgresql-18-pgtt_4.4-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgtt` | `4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.8 KiB | [postgresql-18-pgtt_4.4-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 54.7 KiB | [postgresql-18-pgtt_4.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 52.9 KiB | [postgresql-18-pgtt_4.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 53.6 KiB | [postgresql-18-pgtt_4.4-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 51.9 KiB | [postgresql-18-pgtt_4.4-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.7 KiB | [postgresql-18-pgtt_4.4-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgtt` | `4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 51.5 KiB | [postgresql-18-pgtt_4.4-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-18-pgtt_4.4-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtt_17` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.6 KiB | [pgtt_17-4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtt_17-4.4-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_17` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.2 KiB | [pgtt_17-4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtt_17-4.3-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_17` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.9 KiB | [pgtt_17-4.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtt_17-4.0-3PGDG.rhel8.x86_64.rpm) |
| `pgtt_17` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.8 KiB | [pgtt_17-4.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtt_17-4.0-2PGDG.rhel8.x86_64.rpm) |
| `pgtt_17` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.2 KiB | [pgtt_17-4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtt_17-4.4-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_17` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.7 KiB | [pgtt_17-4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtt_17-4.3-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_17` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.5 KiB | [pgtt_17-4.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtt_17-4.0-3PGDG.rhel8.aarch64.rpm) |
| `pgtt_17` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.4 KiB | [pgtt_17-4.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtt_17-4.0-2PGDG.rhel8.aarch64.rpm) |
| `pgtt_17` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.7 KiB | [pgtt_17-4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtt_17-4.4-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_17` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.2 KiB | [pgtt_17-4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtt_17-4.3-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_17` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.2 KiB | [pgtt_17-4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtt_17-4.0-3PGDG.rhel9.x86_64.rpm) |
| `pgtt_17` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.3 KiB | [pgtt_17-4.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtt_17-4.0-2PGDG.rhel9.x86_64.rpm) |
| `pgtt_17` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [pgtt_17-4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtt_17-4.4-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_17` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.9 KiB | [pgtt_17-4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtt_17-4.3-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_17` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.9 KiB | [pgtt_17-4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtt_17-4.0-3PGDG.rhel9.aarch64.rpm) |
| `pgtt_17` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.0 KiB | [pgtt_17-4.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtt_17-4.0-2PGDG.rhel9.aarch64.rpm) |
| `pgtt_17` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [pgtt_17-4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgtt_17-4.4-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_17` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.6 KiB | [pgtt_17-4.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgtt_17-4.3-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_17` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.7 KiB | [pgtt_17-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgtt_17-4.1-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.1 KiB | [pgtt_17-4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgtt_17-4.4-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.6 KiB | [pgtt_17-4.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgtt_17-4.3-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.7 KiB | [pgtt_17-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgtt_17-4.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgtt` | `4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.2 KiB | [postgresql-17-pgtt_4.4-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgtt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 57.8 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgtt` | `4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.6 KiB | [postgresql-17-pgtt_4.4-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgtt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.2 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgtt` | `4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.8 KiB | [postgresql-17-pgtt_4.4-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgtt` | `4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.7 KiB | [postgresql-17-pgtt_4.4-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.8 KiB | [postgresql-17-pgtt_4.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgtt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 56.8 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.1 KiB | [postgresql-17-pgtt_4.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgtt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 55.6 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 53.5 KiB | [postgresql-17-pgtt_4.4-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgtt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 49.4 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 51.8 KiB | [postgresql-17-pgtt_4.4-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgtt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.3 KiB | [postgresql-17-pgtt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-17-pgtt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.7 KiB | [postgresql-17-pgtt_4.4-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgtt` | `4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 51.4 KiB | [postgresql-17-pgtt_4.4-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-17-pgtt_4.4-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtt_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.7 KiB | [pgtt_16-4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-4.4-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.2 KiB | [pgtt_16-4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-4.3-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.9 KiB | [pgtt_16-4.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-4.0-3PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.7 KiB | [pgtt_16-4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.7 KiB | [pgtt_16-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [pgtt_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtt_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.2 KiB | [pgtt_16-4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-4.4-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.8 KiB | [pgtt_16-4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-4.3-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.5 KiB | [pgtt_16-4.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-4.0-3PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.3 KiB | [pgtt_16-4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [pgtt_16-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.1 KiB | [pgtt_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtt_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.8 KiB | [pgtt_16-4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-4.4-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.3 KiB | [pgtt_16-4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-4.3-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.2 KiB | [pgtt_16-4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-4.0-3PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.2 KiB | [pgtt_16-4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.2 KiB | [pgtt_16-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [pgtt_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtt_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.4 KiB | [pgtt_16-4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-4.4-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.9 KiB | [pgtt_16-4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-4.3-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.9 KiB | [pgtt_16-4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-4.0-3PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.9 KiB | [pgtt_16-4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [pgtt_16-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [pgtt_16-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtt_16-3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_16` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [pgtt_16-4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgtt_16-4.4-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_16` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.6 KiB | [pgtt_16-4.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgtt_16-4.3-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_16` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.7 KiB | [pgtt_16-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgtt_16-4.1-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.1 KiB | [pgtt_16-4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgtt_16-4.4-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.6 KiB | [pgtt_16-4.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgtt_16-4.3-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.7 KiB | [pgtt_16-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgtt_16-4.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgtt` | `4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.4 KiB | [postgresql-16-pgtt_4.4-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgtt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 57.6 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgtt` | `4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-16-pgtt_4.4-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgtt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.3 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgtt` | `4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.8 KiB | [postgresql-16-pgtt_4.4-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgtt` | `4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.8 KiB | [postgresql-16-pgtt_4.4-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.9 KiB | [postgresql-16-pgtt_4.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgtt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 56.8 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.2 KiB | [postgresql-16-pgtt_4.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgtt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 55.6 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 53.6 KiB | [postgresql-16-pgtt_4.4-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgtt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 49.6 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 52.0 KiB | [postgresql-16-pgtt_4.4-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgtt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.4 KiB | [postgresql-16-pgtt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-16-pgtt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.9 KiB | [postgresql-16-pgtt_4.4-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgtt` | `4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 51.7 KiB | [postgresql-16-pgtt_4.4-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-16-pgtt_4.4-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtt_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.9 KiB | [pgtt_15-4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-4.4-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pgtt_15-4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-4.3-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.2 KiB | [pgtt_15-4.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-4.0-3PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [pgtt_15-4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.0 KiB | [pgtt_15-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.7 KiB | [pgtt_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_15` | `2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [pgtt_15-2.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-2.10-1.rhel8.x86_64.rpm) |
| `pgtt_15` | `2.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.3 KiB | [pgtt_15-2.9-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtt_15-2.9-1.rhel8.x86_64.rpm) |
| `pgtt_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.5 KiB | [pgtt_15-4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-4.4-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.0 KiB | [pgtt_15-4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-4.3-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.7 KiB | [pgtt_15-4.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-4.0-3PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.5 KiB | [pgtt_15-4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [pgtt_15-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [pgtt_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_15` | `2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.7 KiB | [pgtt_15-2.10-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-2.10-1.rhel8.aarch64.rpm) |
| `pgtt_15` | `2.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.7 KiB | [pgtt_15-2.9-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtt_15-2.9-1.rhel8.aarch64.rpm) |
| `pgtt_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.7 KiB | [pgtt_15-4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-4.4-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.2 KiB | [pgtt_15-4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-4.3-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.1 KiB | [pgtt_15-4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-4.0-3PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.1 KiB | [pgtt_15-4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.1 KiB | [pgtt_15-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.0 KiB | [pgtt_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_15` | `2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.4 KiB | [pgtt_15-2.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-2.10-1.rhel9.x86_64.rpm) |
| `pgtt_15` | `2.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.9 KiB | [pgtt_15-2.9-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtt_15-2.9-1.rhel9.x86_64.rpm) |
| `pgtt_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.2 KiB | [pgtt_15-4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-4.4-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.8 KiB | [pgtt_15-4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-4.3-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [pgtt_15-4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-4.0-3PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [pgtt_15-4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.7 KiB | [pgtt_15-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.5 KiB | [pgtt_15-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_15` | `2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.2 KiB | [pgtt_15-2.10-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-2.10-1.rhel9.aarch64.rpm) |
| `pgtt_15` | `2.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.3 KiB | [pgtt_15-2.9-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtt_15-2.9-1.rhel9.aarch64.rpm) |
| `pgtt_15` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.9 KiB | [pgtt_15-4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgtt_15-4.4-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_15` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.5 KiB | [pgtt_15-4.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgtt_15-4.3-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_15` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.6 KiB | [pgtt_15-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgtt_15-4.1-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.9 KiB | [pgtt_15-4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgtt_15-4.4-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.5 KiB | [pgtt_15-4.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgtt_15-4.3-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.6 KiB | [pgtt_15-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgtt_15-4.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgtt` | `4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pgtt_4.4-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgtt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.1 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgtt` | `4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.1 KiB | [postgresql-15-pgtt_4.4-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgtt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.6 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgtt` | `4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.5 KiB | [postgresql-15-pgtt_4.4-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgtt` | `4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-15-pgtt_4.4-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 61.8 KiB | [postgresql-15-pgtt_4.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgtt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.9 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 60.2 KiB | [postgresql-15-pgtt_4.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgtt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.7 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 54.3 KiB | [postgresql-15-pgtt_4.4-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgtt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.4 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 52.6 KiB | [postgresql-15-pgtt_4.4-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgtt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.4 KiB | [postgresql-15-pgtt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-15-pgtt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 54.3 KiB | [postgresql-15-pgtt_4.4-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgtt` | `4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 52.3 KiB | [postgresql-15-pgtt_4.4-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-15-pgtt_4.4-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtt_14` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.9 KiB | [pgtt_14-4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-4.4-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pgtt_14-4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-4.3-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.2 KiB | [pgtt_14-4.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-4.0-3PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [pgtt_14-4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.0 KiB | [pgtt_14-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.7 KiB | [pgtt_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgtt_14` | `2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [pgtt_14-2.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-2.10-1.rhel8.x86_64.rpm) |
| `pgtt_14` | `2.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pgtt_14-2.9-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-2.9-1.rhel8.x86_64.rpm) |
| `pgtt_14` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.9 KiB | [pgtt_14-2.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-2.8-1.rhel8.x86_64.rpm) |
| `pgtt_14` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.3 KiB | [pgtt_14-2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtt_14-2.6-1.rhel8.x86_64.rpm) |
| `pgtt_14` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.5 KiB | [pgtt_14-4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-4.4-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.0 KiB | [pgtt_14-4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-4.3-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.7 KiB | [pgtt_14-4.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-4.0-3PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.5 KiB | [pgtt_14-4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [pgtt_14-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.3 KiB | [pgtt_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgtt_14` | `2.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.7 KiB | [pgtt_14-2.10-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-2.10-1.rhel8.aarch64.rpm) |
| `pgtt_14` | `2.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.5 KiB | [pgtt_14-2.9-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtt_14-2.9-1.rhel8.aarch64.rpm) |
| `pgtt_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.7 KiB | [pgtt_14-4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-4.4-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.2 KiB | [pgtt_14-4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-4.3-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.2 KiB | [pgtt_14-4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-4.0-3PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.1 KiB | [pgtt_14-4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.1 KiB | [pgtt_14-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.0 KiB | [pgtt_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgtt_14` | `2.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.4 KiB | [pgtt_14-2.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-2.10-1.rhel9.x86_64.rpm) |
| `pgtt_14` | `2.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.7 KiB | [pgtt_14-2.9-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-2.9-1.rhel9.x86_64.rpm) |
| `pgtt_14` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.5 KiB | [pgtt_14-2.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtt_14-2.8-1.rhel9.x86_64.rpm) |
| `pgtt_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.2 KiB | [pgtt_14-4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-4.4-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.8 KiB | [pgtt_14-4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-4.3-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [pgtt_14-4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-4.0-3PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [pgtt_14-4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.7 KiB | [pgtt_14-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.5 KiB | [pgtt_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgtt_14` | `2.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.2 KiB | [pgtt_14-2.10-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-2.10-1.rhel9.aarch64.rpm) |
| `pgtt_14` | `2.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.2 KiB | [pgtt_14-2.9-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtt_14-2.9-1.rhel9.aarch64.rpm) |
| `pgtt_14` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.9 KiB | [pgtt_14-4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgtt_14-4.4-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_14` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.5 KiB | [pgtt_14-4.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgtt_14-4.3-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_14` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.5 KiB | [pgtt_14-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgtt_14-4.1-1PGDG.rhel10.x86_64.rpm) |
| `pgtt_14` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.9 KiB | [pgtt_14-4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgtt_14-4.4-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_14` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.5 KiB | [pgtt_14-4.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgtt_14-4.3-1PGDG.rhel10.aarch64.rpm) |
| `pgtt_14` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.6 KiB | [pgtt_14-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgtt_14-4.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgtt` | `4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 58.2 KiB | [postgresql-14-pgtt_4.4-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgtt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 57.9 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgtt` | `4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.0 KiB | [postgresql-14-pgtt_4.4-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgtt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.6 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgtt` | `4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.6 KiB | [postgresql-14-pgtt_4.4-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgtt` | `4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.1 KiB | [postgresql-14-pgtt_4.4-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 61.7 KiB | [postgresql-14-pgtt_4.4-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgtt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.8 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 60.0 KiB | [postgresql-14-pgtt_4.4-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgtt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.7 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 54.2 KiB | [postgresql-14-pgtt_4.4-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgtt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.4 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 52.6 KiB | [postgresql-14-pgtt_4.4-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgtt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.4 KiB | [postgresql-14-pgtt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtt/postgresql-14-pgtt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 54.2 KiB | [postgresql-14-pgtt_4.4-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgtt` | `4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 52.3 KiB | [postgresql-14-pgtt_4.4-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtt/postgresql-14-pgtt_4.4-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/darold/pgtt" title="Repository" icon="github" subtitle="github.com/darold/pgtt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgtt-4.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgtt;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgtt;		# install via package name, for the active PG version

pig install pgtt -v 18;   # install for PG 18
pig install pgtt -v 17;   # install for PG 17
pig install pgtt -v 16;   # install for PG 16
pig install pgtt -v 15;   # install for PG 15
pig install pgtt -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgtt';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgtt;
```



## Usage

> [pgtt: Extension to add Global Temporary Tables feature to PostgreSQL](https://github.com/darold/pgtt)

### Creating a Global Temporary Table

```sql
CREATE EXTENSION pgtt;

-- ON COMMIT PRESERVE ROWS: data persists across transactions within a session
CREATE GLOBAL TEMPORARY TABLE test_gtt (
    id integer,
    lbl text
) ON COMMIT PRESERVE ROWS;

-- ON COMMIT DELETE ROWS: data is deleted at transaction commit
CREATE GLOBAL TEMPORARY TABLE session_data (
    id integer,
    value text
) ON COMMIT DELETE ROWS;
```

The `GLOBAL` keyword can also be used as a comment to avoid warnings:

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE test_gtt (
    LIKE other_table INCLUDING DEFAULTS INCLUDING CONSTRAINTS INCLUDING INDEXES
) ON COMMIT PRESERVE ROWS;
```

### CREATE AS Form

```sql
CREATE /*GLOBAL*/ TEMPORARY TABLE gtt_copy
AS SELECT * FROM source_table WITH DATA;
```

### Using Global Temporary Tables

```sql
INSERT INTO test_gtt VALUES (1, 'one'), (2, 'two');
SELECT * FROM test_gtt;  -- visible only in current session
```

### Creating Indexes

```sql
CREATE INDEX ON test_gtt (id);
```

### Constraints

All constraints except FOREIGN KEYS are supported:

```sql
CREATE GLOBAL TEMPORARY TABLE t2 (
    c1 serial PRIMARY KEY,
    c2 VARCHAR(50) UNIQUE NOT NULL,
    c3 boolean DEFAULT false
);
```

### Dropping

```sql
DROP TABLE test_gtt;  -- can be dropped even while used by other sessions
```

### Configuration

```sql
SET pgtt.enabled TO off;   -- disable GTT rerouting
SET pgtt.enabled TO on;    -- re-enable GTT rerouting
```

### Key Behaviors

- GTT content is session-local; other sessions cannot see your data
- The table structure is persistent (visible to all users), but data is per-session
- Requires loading via `session_preload_libraries = 'pgtt'`
- Partitioning is not supported on GTTs
