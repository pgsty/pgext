---
title: "pgl_ddl_deploy"
linkTitle: "pgl_ddl_deploy"
description: "automated ddl deployment using pglogical"
weight: 9520
categories: ["ETL"]
width: full
---

[**pgl_ddl_deploy**](https://github.com/enova/pgl_ddl_deploy) : automated ddl deployment using pglogical


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9520** | {{< badge content="pgl_ddl_deploy" link="https://github.com/enova/pgl_ddl_deploy" >}} | {{< ext "pgl_ddl_deploy" >}} | `2.2.1` | {{< category "ETL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pglogical" >}} |
|   **See Also**    | {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "ddlx" >}} {{< ext "pg_permissions" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2.1` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgl_ddl_deploy` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2.1` | {{< bg "18" "pgl_ddl_deploy_18*" "red" >}} {{< bg "17" "pgl_ddl_deploy_17*" "green" >}} {{< bg "16" "pgl_ddl_deploy_16*" "green" >}} {{< bg "15" "pgl_ddl_deploy_15*" "green" >}} {{< bg "14" "pgl_ddl_deploy_14*" "green" >}} {{< bg "13" "pgl_ddl_deploy_13*" "green" >}} | `pgl_ddl_deploy_$v*` | `pglogical_$v` |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2.1` | {{< bg "18" "postgresql-18-pgl-ddl-deploy" "red" >}} {{< bg "17" "postgresql-17-pgl-ddl-deploy" "green" >}} {{< bg "16" "postgresql-16-pgl-ddl-deploy" "green" >}} {{< bg "15" "postgresql-15-pgl-ddl-deploy" "green" >}} {{< bg "14" "postgresql-14-pgl-ddl-deploy" "green" >}} {{< bg "13" "postgresql-13-pgl-ddl-deploy" "green" >}} | `postgresql-$v-pgl-ddl-deploy` | `postgresql-$v-pglogical` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.0" "pgl_ddl_deploy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.0" "pgl_ddl_deploy_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.0" "pgl_ddl_deploy_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pgl_ddl_deploy_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "pgl_ddl_deploy_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgl-ddl-deploy : MISS 0" "red" >}}      | {{< bg "PGDG 2.2.1" "postgresql-17-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-16-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-15-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-14-pgl-ddl-deploy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2.1" "postgresql-13-pgl-ddl-deploy : AVAIL 1" "blue" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgl_ddl_deploy_17` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [pgl_ddl_deploy_17-2.2.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgl_ddl_deploy_17-2.2.1-2PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_17` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pgl_ddl_deploy_17-2.2.1-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgl_ddl_deploy_17-2.2.1-2PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_17` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.2 KiB | [pgl_ddl_deploy_17-2.2.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgl_ddl_deploy_17-2.2.1-2PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_17` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.2 KiB | [pgl_ddl_deploy_17-2.2.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgl_ddl_deploy_17-2.2.1-2PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_17` | `2.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.4 KiB | [pgl_ddl_deploy_17-2.2.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgl_ddl_deploy_17-2.2.1-3PGDG.rhel10.x86_64.rpm) |
| `pgl_ddl_deploy_17` | `2.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [pgl_ddl_deploy_17-2.2.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgl_ddl_deploy_17-2.2.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.4 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.8 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.4 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.9 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.0 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.0 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.6 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgl-ddl-deploy` | `2.2.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.9 KiB | [postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-17-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgl_ddl_deploy_16` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [pgl_ddl_deploy_16-2.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgl_ddl_deploy_16-2.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [pgl_ddl_deploy_16-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgl_ddl_deploy_16-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.6 KiB | [pgl_ddl_deploy_16-2.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgl_ddl_deploy_16-2.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.4 KiB | [pgl_ddl_deploy_16-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgl_ddl_deploy_16-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 38.1 KiB | [pgl_ddl_deploy_16-2.2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgl_ddl_deploy_16-2.2.1-1PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.7 KiB | [pgl_ddl_deploy_16-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgl_ddl_deploy_16-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [pgl_ddl_deploy_16-2.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgl_ddl_deploy_16-2.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.6 KiB | [pgl_ddl_deploy_16-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgl_ddl_deploy_16-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.4 KiB | [pgl_ddl_deploy_16-2.2.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgl_ddl_deploy_16-2.2.1-3PGDG.rhel10.x86_64.rpm) |
| `pgl_ddl_deploy_16` | `2.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [pgl_ddl_deploy_16-2.2.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgl_ddl_deploy_16-2.2.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.4 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.8 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.4 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.9 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.0 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.0 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.5 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgl-ddl-deploy` | `2.2.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.9 KiB | [postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-16-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgl_ddl_deploy_15` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [pgl_ddl_deploy_15-2.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgl_ddl_deploy_15-2.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pgl_ddl_deploy_15-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgl_ddl_deploy_15-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pgl_ddl_deploy_15-2.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgl_ddl_deploy_15-2.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [pgl_ddl_deploy_15-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgl_ddl_deploy_15-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.6 KiB | [pgl_ddl_deploy_15-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgl_ddl_deploy_15-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.0 KiB | [pgl_ddl_deploy_15-2.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgl_ddl_deploy_15-2.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.5 KiB | [pgl_ddl_deploy_15-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgl_ddl_deploy_15-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pgl_ddl_deploy_15-2.2.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgl_ddl_deploy_15-2.2.1-3PGDG.rhel10.x86_64.rpm) |
| `pgl_ddl_deploy_15` | `2.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.6 KiB | [pgl_ddl_deploy_15-2.2.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgl_ddl_deploy_15-2.2.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.0 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.4 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.0 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.4 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.6 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 39.5 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.3 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgl-ddl-deploy` | `2.2.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.6 KiB | [postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-15-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgl_ddl_deploy_14` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [pgl_ddl_deploy_14-2.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgl_ddl_deploy_14-2.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.3 KiB | [pgl_ddl_deploy_14-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgl_ddl_deploy_14-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pgl_ddl_deploy_14-2.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgl_ddl_deploy_14-2.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [pgl_ddl_deploy_14-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgl_ddl_deploy_14-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.6 KiB | [pgl_ddl_deploy_14-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgl_ddl_deploy_14-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.0 KiB | [pgl_ddl_deploy_14-2.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgl_ddl_deploy_14-2.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.5 KiB | [pgl_ddl_deploy_14-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgl_ddl_deploy_14-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pgl_ddl_deploy_14-2.2.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgl_ddl_deploy_14-2.2.1-3PGDG.rhel10.x86_64.rpm) |
| `pgl_ddl_deploy_14` | `2.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.6 KiB | [pgl_ddl_deploy_14-2.2.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgl_ddl_deploy_14-2.2.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.0 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.4 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.9 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.4 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.6 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 39.5 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.3 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgl-ddl-deploy` | `2.2.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.6 KiB | [postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-14-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgl_ddl_deploy_13` | `2.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [pgl_ddl_deploy_13-2.2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgl_ddl_deploy_13-2.2.1-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.2 KiB | [pgl_ddl_deploy_13-2.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgl_ddl_deploy_13-2.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pgl_ddl_deploy_13-2.2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgl_ddl_deploy_13-2.2.1-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.4 KiB | [pgl_ddl_deploy_13-2.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgl_ddl_deploy_13-2.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.6 KiB | [pgl_ddl_deploy_13-2.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgl_ddl_deploy_13-2.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.0 KiB | [pgl_ddl_deploy_13-2.2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgl_ddl_deploy_13-2.2.1-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.5 KiB | [pgl_ddl_deploy_13-2.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgl_ddl_deploy_13-2.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pgl_ddl_deploy_13-2.2.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgl_ddl_deploy_13-2.2.1-3PGDG.rhel10.x86_64.rpm) |
| `pgl_ddl_deploy_13` | `2.2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.6 KiB | [pgl_ddl_deploy_13-2.2.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgl_ddl_deploy_13-2.2.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 38.7 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg120+1_amd64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 39.0 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg120+1_arm64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 38.7 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg130+2_amd64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 39.0 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg130+2_arm64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.4 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 39.1 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 38.1 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgl-ddl-deploy` | `2.2.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 38.2 KiB | [postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgl-ddl-deploy/postgresql-13-pgl-ddl-deploy_2.2.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/enova/pgl_ddl_deploy" title="Repository" icon="github" subtitle="github.com/enova/pgl_ddl_deploy" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgl_ddl_deploy;		# install via package name, for the active PG version

pig install pgl_ddl_deploy -v 17;   # install for PG 17
pig install pgl_ddl_deploy -v 16;   # install for PG 16
pig install pgl_ddl_deploy -v 15;   # install for PG 15
pig install pgl_ddl_deploy -v 14;   # install for PG 14
pig install pgl_ddl_deploy -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgl_ddl_deploy CASCADE; -- requires pglogical
```
