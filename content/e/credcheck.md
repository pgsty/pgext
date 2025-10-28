---
title: "credcheck"
linkTitle: "credcheck"
description: "credcheck - postgresql plain text credential checker"
weight: 7110
categories: ["SEC"]
width: full
---

credcheck - postgresql plain text credential checker


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7110** | {{< badge content="credcheck" link="https://github.com/MigOpsRepos/credcheck" >}} | {{< ext "credcheck" >}} | `3.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "login_hook" >}} {{< ext "passwordcheck" >}} {{< ext "pgaudit" >}} {{< ext "pg_auth_mon" >}} {{< ext "set_user" >}} {{< ext "auth_delay" >}} {{< ext "pg_permissions" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/credcheck" >}} | `3.0` | {{< bg "18" "credcheck_18*" "green" >}} {{< bg "17" "credcheck_17*" "green" >}} {{< bg "16" "credcheck_16*" "green" >}} {{< bg "15" "credcheck_15*" "green" >}} {{< bg "14" "credcheck_14*" "green" >}} {{< bg "13" "credcheck_13*" "green" >}} | `credcheck_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/credcheck" >}} | `3.0` | {{< bg "18" "postgresql-18-credcheck" "green" >}} {{< bg "17" "postgresql-17-credcheck" "green" >}} {{< bg "16" "postgresql-16-credcheck" "green" >}} {{< bg "15" "postgresql-15-credcheck" "green" >}} {{< bg "14" "postgresql-14-credcheck" "green" >}} {{< bg "13" "postgresql-13-credcheck" "green" >}} | `postgresql-$v-credcheck` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 12" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 11" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 10" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 11" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 4.1" "credcheck_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1" "credcheck_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 3.0" "postgresql-18-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-17-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-16-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-15-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-14-credcheck : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "postgresql-13-credcheck : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_18` | 4.1 | `el8.x86_64` | pgdg | 39.4 KiB | [credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | 3.0 | `el8.x86_64` | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | 4.1 | `el8.aarch64` | pgdg | 38.8 KiB | [credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | 3.0 | `el8.aarch64` | pgdg | 35.1 KiB | [credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | 4.1 | `el9.x86_64` | pgdg | 39.2 KiB | [credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | 3.0 | `el9.x86_64` | pgdg | 35.9 KiB | [credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | 4.1 | `el9.aarch64` | pgdg | 38.7 KiB | [credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | 3.0 | `el9.aarch64` | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | 4.1 | `el10.x86_64` | pgdg | 39.7 KiB | [credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | 3.0 | `el10.x86_64` | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | 4.1 | `el10.aarch64` | pgdg | 39.5 KiB | [credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | 3.0 | `el10.aarch64` | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.3 KiB | [postgresql-18-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.8 KiB | [postgresql-18-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.6 KiB | [postgresql-18-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.9 KiB | [postgresql-18-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | 3.0 | `u22.x86_64` | pgdg | 59.0 KiB | [postgresql-18-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | 3.0 | `u22.aarch64` | pgdg | 58.1 KiB | [postgresql-18-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.8 KiB | [postgresql-18-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | 3.0 | `u24.aarch64` | pgdg | 58.1 KiB | [postgresql-18-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_17` | 4.1 | `el8.x86_64` | pgdg | 39.5 KiB | [credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | 3.0 | `el8.x86_64` | pgdg | 35.5 KiB | [credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | 2.8 | `el8.x86_64` | pgdg | 35.1 KiB | [credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | 4.1 | `el8.aarch64` | pgdg | 38.9 KiB | [credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | 3.0 | `el8.aarch64` | pgdg | 35.0 KiB | [credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | 2.8 | `el8.aarch64` | pgdg | 34.7 KiB | [credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | 4.1 | `el9.x86_64` | pgdg | 39.2 KiB | [credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | 3.0 | `el9.x86_64` | pgdg | 35.9 KiB | [credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | 2.8 | `el9.x86_64` | pgdg | 35.6 KiB | [credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | 4.1 | `el9.aarch64` | pgdg | 38.8 KiB | [credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | 3.0 | `el9.aarch64` | pgdg | 35.7 KiB | [credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | 2.8 | `el9.aarch64` | pgdg | 35.4 KiB | [credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | 4.1 | `el10.x86_64` | pgdg | 39.8 KiB | [credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | 3.0 | `el10.x86_64` | pgdg | 36.5 KiB | [credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | 4.1 | `el10.aarch64` | pgdg | 39.6 KiB | [credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | 3.0 | `el10.aarch64` | pgdg | 36.4 KiB | [credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.3 KiB | [postgresql-17-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.9 KiB | [postgresql-17-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.5 KiB | [postgresql-17-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.7 KiB | [postgresql-17-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | 3.0 | `u22.x86_64` | pgdg | 65.8 KiB | [postgresql-17-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | 3.0 | `u22.aarch64` | pgdg | 64.9 KiB | [postgresql-17-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.9 KiB | [postgresql-17-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | 3.0 | `u24.aarch64` | pgdg | 58.2 KiB | [postgresql-17-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_16` | 4.1 | `el8.x86_64` | pgdg | 39.5 KiB | [credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 3.0 | `el8.x86_64` | pgdg | 35.5 KiB | [credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 2.7 | `el8.x86_64` | pgdg | 34.7 KiB | [credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 2.6 | `el8.x86_64` | pgdg | 34.3 KiB | [credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 2.2 | `el8.x86_64` | pgdg | 32.8 KiB | [credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 2.1 | `el8.x86_64` | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | 4.1 | `el8.aarch64` | pgdg | 38.9 KiB | [credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 3.0 | `el8.aarch64` | pgdg | 35.1 KiB | [credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 2.7 | `el8.aarch64` | pgdg | 34.2 KiB | [credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 2.6 | `el8.aarch64` | pgdg | 33.9 KiB | [credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 2.2 | `el8.aarch64` | pgdg | 32.5 KiB | [credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 2.1 | `el8.aarch64` | pgdg | 31.3 KiB | [credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | 4.1 | `el9.x86_64` | pgdg | 39.2 KiB | [credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 3.0 | `el9.x86_64` | pgdg | 36.2 KiB | [credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 2.7 | `el9.x86_64` | pgdg | 35.1 KiB | [credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 2.6 | `el9.x86_64` | pgdg | 34.7 KiB | [credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 2.2 | `el9.x86_64` | pgdg | 33.5 KiB | [credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 2.1 | `el9.x86_64` | pgdg | 32.3 KiB | [credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | 4.1 | `el9.aarch64` | pgdg | 38.8 KiB | [credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 3.0 | `el9.aarch64` | pgdg | 35.7 KiB | [credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 2.7 | `el9.aarch64` | pgdg | 34.8 KiB | [credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 2.6 | `el9.aarch64` | pgdg | 34.5 KiB | [credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 2.2 | `el9.aarch64` | pgdg | 32.9 KiB | [credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 2.1 | `el9.aarch64` | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | 4.1 | `el10.x86_64` | pgdg | 39.8 KiB | [credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | 3.0 | `el10.x86_64` | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | 4.1 | `el10.aarch64` | pgdg | 39.6 KiB | [credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | 3.0 | `el10.aarch64` | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.3 KiB | [postgresql-16-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.8 KiB | [postgresql-16-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.5 KiB | [postgresql-16-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.7 KiB | [postgresql-16-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | 3.0 | `u22.x86_64` | pgdg | 65.4 KiB | [postgresql-16-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | 3.0 | `u22.aarch64` | pgdg | 64.5 KiB | [postgresql-16-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.7 KiB | [postgresql-16-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | 3.0 | `u24.aarch64` | pgdg | 58.2 KiB | [postgresql-16-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_15` | 4.1 | `el8.x86_64` | pgdg | 39.6 KiB | [credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 3.0 | `el8.x86_64` | pgdg | 35.6 KiB | [credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 2.7 | `el8.x86_64` | pgdg | 34.7 KiB | [credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 2.6 | `el8.x86_64` | pgdg | 34.4 KiB | [credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 2.2 | `el8.x86_64` | pgdg | 33.0 KiB | [credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 2.1 | `el8.x86_64` | pgdg | 31.9 KiB | [credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | 2.0 | `el8.x86_64` | pgdg | 31.1 KiB | [credcheck_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | 1.2 | `el8.x86_64` | pgdg | 27.7 KiB | [credcheck_15-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_15` | 1.0 | `el8.x86_64` | pgdg | 27.1 KiB | [credcheck_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | 0.2.0 | `el8.x86_64` | pgdg | 35.0 KiB | [credcheck_15-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | 0.2.0 | `el8.x86_64` | pgdg | 18.6 KiB | [credcheck_15-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_15` | 4.1 | `el8.aarch64` | pgdg | 38.8 KiB | [credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 3.0 | `el8.aarch64` | pgdg | 35.0 KiB | [credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 2.7 | `el8.aarch64` | pgdg | 34.2 KiB | [credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 2.6 | `el8.aarch64` | pgdg | 33.9 KiB | [credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 2.2 | `el8.aarch64` | pgdg | 32.5 KiB | [credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 2.1 | `el8.aarch64` | pgdg | 31.3 KiB | [credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | 2.0 | `el8.aarch64` | pgdg | 30.5 KiB | [credcheck_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | 1.2 | `el8.aarch64` | pgdg | 27.2 KiB | [credcheck_15-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_15` | 1.0 | `el8.aarch64` | pgdg | 26.6 KiB | [credcheck_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | 0.2.0 | `el8.aarch64` | pgdg | 18.3 KiB | [credcheck_15-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_15` | 0.2.0 | `el8.aarch64` | pgdg | 34.9 KiB | [credcheck_15-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | 4.1 | `el9.x86_64` | pgdg | 39.3 KiB | [credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 3.0 | `el9.x86_64` | pgdg | 36.1 KiB | [credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 2.7 | `el9.x86_64` | pgdg | 35.2 KiB | [credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 2.6 | `el9.x86_64` | pgdg | 34.9 KiB | [credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 2.2 | `el9.x86_64` | pgdg | 33.4 KiB | [credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 2.1 | `el9.x86_64` | pgdg | 32.5 KiB | [credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | 2.0 | `el9.x86_64` | pgdg | 31.6 KiB | [credcheck_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | 1.2 | `el9.x86_64` | pgdg | 28.1 KiB | [credcheck_15-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_15` | 1.0 | `el9.x86_64` | pgdg | 27.5 KiB | [credcheck_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | 0.2.0 | `el9.x86_64` | pgdg | 18.8 KiB | [credcheck_15-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_15` | 0.2.0 | `el9.x86_64` | pgdg | 35.9 KiB | [credcheck_15-0.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | 4.1 | `el9.aarch64` | pgdg | 38.7 KiB | [credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 3.0 | `el9.aarch64` | pgdg | 35.8 KiB | [credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 2.7 | `el9.aarch64` | pgdg | 34.8 KiB | [credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 2.6 | `el9.aarch64` | pgdg | 34.5 KiB | [credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 2.2 | `el9.aarch64` | pgdg | 32.9 KiB | [credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 2.1 | `el9.aarch64` | pgdg | 31.8 KiB | [credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | 2.0 | `el9.aarch64` | pgdg | 30.9 KiB | [credcheck_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | 1.2 | `el9.aarch64` | pgdg | 27.5 KiB | [credcheck_15-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_15` | 1.0 | `el9.aarch64` | pgdg | 26.9 KiB | [credcheck_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | 0.2.0 | `el9.aarch64` | pgdg | 18.1 KiB | [credcheck_15-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_15` | 0.2.0 | `el9.aarch64` | pgdg | 35.5 KiB | [credcheck_15-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | 4.1 | `el10.x86_64` | pgdg | 39.8 KiB | [credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | 3.0 | `el10.x86_64` | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | 4.1 | `el10.aarch64` | pgdg | 39.5 KiB | [credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | 3.0 | `el10.aarch64` | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.5 KiB | [postgresql-15-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.9 KiB | [postgresql-15-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.6 KiB | [postgresql-15-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.7 KiB | [postgresql-15-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | 3.0 | `u22.x86_64` | pgdg | 65.5 KiB | [postgresql-15-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | 3.0 | `u22.aarch64` | pgdg | 64.5 KiB | [postgresql-15-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.8 KiB | [postgresql-15-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | 3.0 | `u24.aarch64` | pgdg | 58.2 KiB | [postgresql-15-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_14` | 4.1 | `el8.x86_64` | pgdg | 39.6 KiB | [credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 3.0 | `el8.x86_64` | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 2.7 | `el8.x86_64` | pgdg | 34.6 KiB | [credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 2.6 | `el8.x86_64` | pgdg | 34.3 KiB | [credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 2.2 | `el8.x86_64` | pgdg | 32.9 KiB | [credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 2.1 | `el8.x86_64` | pgdg | 31.8 KiB | [credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | 2.0 | `el8.x86_64` | pgdg | 31.0 KiB | [credcheck_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | 1.2 | `el8.x86_64` | pgdg | 27.7 KiB | [credcheck_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_14` | 1.0 | `el8.x86_64` | pgdg | 27.1 KiB | [credcheck_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | 0.2.0 | `el8.x86_64` | pgdg | 35.3 KiB | [credcheck_14-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | 0.2.0 | `el8.x86_64` | pgdg | 18.6 KiB | [credcheck_14-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_14` | 4.1 | `el8.aarch64` | pgdg | 38.7 KiB | [credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 3.0 | `el8.aarch64` | pgdg | 35.0 KiB | [credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 2.7 | `el8.aarch64` | pgdg | 34.1 KiB | [credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 2.6 | `el8.aarch64` | pgdg | 33.8 KiB | [credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 2.2 | `el8.aarch64` | pgdg | 32.4 KiB | [credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 2.1 | `el8.aarch64` | pgdg | 31.2 KiB | [credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | 2.0 | `el8.aarch64` | pgdg | 30.4 KiB | [credcheck_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | 1.2 | `el8.aarch64` | pgdg | 27.2 KiB | [credcheck_14-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_14` | 1.0 | `el8.aarch64` | pgdg | 26.6 KiB | [credcheck_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | 0.2.0 | `el8.aarch64` | pgdg | 18.3 KiB | [credcheck_14-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_14` | 0.2.0 | `el8.aarch64` | pgdg | 34.8 KiB | [credcheck_14-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | 4.1 | `el9.x86_64` | pgdg | 39.3 KiB | [credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 3.0 | `el9.x86_64` | pgdg | 36.0 KiB | [credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 2.7 | `el9.x86_64` | pgdg | 35.1 KiB | [credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 2.6 | `el9.x86_64` | pgdg | 34.8 KiB | [credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 2.2 | `el9.x86_64` | pgdg | 33.3 KiB | [credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 2.1 | `el9.x86_64` | pgdg | 32.2 KiB | [credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | 2.0 | `el9.x86_64` | pgdg | 31.3 KiB | [credcheck_14-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | 1.2 | `el9.x86_64` | pgdg | 28.0 KiB | [credcheck_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_14` | 1.0 | `el9.x86_64` | pgdg | 27.4 KiB | [credcheck_14-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | 0.2.0 | `el9.x86_64` | pgdg | 18.8 KiB | [credcheck_14-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_14` | 4.1 | `el9.aarch64` | pgdg | 38.6 KiB | [credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 3.0 | `el9.aarch64` | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 2.7 | `el9.aarch64` | pgdg | 34.8 KiB | [credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 2.6 | `el9.aarch64` | pgdg | 34.4 KiB | [credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 2.2 | `el9.aarch64` | pgdg | 32.8 KiB | [credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 2.1 | `el9.aarch64` | pgdg | 31.7 KiB | [credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | 2.0 | `el9.aarch64` | pgdg | 30.8 KiB | [credcheck_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | 1.2 | `el9.aarch64` | pgdg | 27.4 KiB | [credcheck_14-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_14` | 1.0 | `el9.aarch64` | pgdg | 26.8 KiB | [credcheck_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | 0.2.0 | `el9.aarch64` | pgdg | 18.0 KiB | [credcheck_14-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_14` | 0.2.0 | `el9.aarch64` | pgdg | 35.4 KiB | [credcheck_14-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | 4.1 | `el10.x86_64` | pgdg | 39.8 KiB | [credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | 3.0 | `el10.x86_64` | pgdg | 36.5 KiB | [credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | 4.1 | `el10.aarch64` | pgdg | 39.2 KiB | [credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_14` | 3.0 | `el10.aarch64` | pgdg | 36.4 KiB | [credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.2 KiB | [postgresql-14-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.6 KiB | [postgresql-14-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.4 KiB | [postgresql-14-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-14-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.5 KiB | [postgresql-14-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-14-credcheck` | 3.0 | `u22.x86_64` | pgdg | 65.2 KiB | [postgresql-14-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-credcheck` | 3.0 | `u22.aarch64` | pgdg | 64.1 KiB | [postgresql-14-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.6 KiB | [postgresql-14-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-credcheck` | 3.0 | `u24.aarch64` | pgdg | 58.0 KiB | [postgresql-14-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_13` | 4.1 | `el8.x86_64` | pgdg | 39.2 KiB | [credcheck_13-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 3.0 | `el8.x86_64` | pgdg | 35.2 KiB | [credcheck_13-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 2.7 | `el8.x86_64` | pgdg | 34.2 KiB | [credcheck_13-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 2.6 | `el8.x86_64` | pgdg | 33.9 KiB | [credcheck_13-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 2.2 | `el8.x86_64` | pgdg | 32.4 KiB | [credcheck_13-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 2.1 | `el8.x86_64` | pgdg | 31.4 KiB | [credcheck_13-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_13` | 2.0 | `el8.x86_64` | pgdg | 30.6 KiB | [credcheck_13-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_13` | 1.2 | `el8.x86_64` | pgdg | 27.3 KiB | [credcheck_13-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_13` | 1.0 | `el8.x86_64` | pgdg | 26.7 KiB | [credcheck_13-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_13` | 0.2.0 | `el8.x86_64` | pgdg | 18.5 KiB | [credcheck_13-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_13` | 0.2.0 | `el8.x86_64` | pgdg | 35.0 KiB | [credcheck_13-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_13` | 0.1.1 | `el8.x86_64` | pgdg | 34.6 KiB | [credcheck_13-0.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/credcheck_13-0.1.1-1.rhel8.x86_64.rpm) |
| `credcheck_13` | 4.1 | `el8.aarch64` | pgdg | 38.7 KiB | [credcheck_13-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 3.0 | `el8.aarch64` | pgdg | 34.9 KiB | [credcheck_13-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 2.7 | `el8.aarch64` | pgdg | 34.0 KiB | [credcheck_13-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 2.6 | `el8.aarch64` | pgdg | 33.7 KiB | [credcheck_13-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 2.2 | `el8.aarch64` | pgdg | 32.3 KiB | [credcheck_13-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 2.1 | `el8.aarch64` | pgdg | 31.1 KiB | [credcheck_13-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_13` | 2.0 | `el8.aarch64` | pgdg | 30.3 KiB | [credcheck_13-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_13` | 1.2 | `el8.aarch64` | pgdg | 27.0 KiB | [credcheck_13-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_13` | 1.0 | `el8.aarch64` | pgdg | 26.4 KiB | [credcheck_13-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_13` | 0.2.0 | `el8.aarch64` | pgdg | 18.3 KiB | [credcheck_13-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_13` | 0.2.0 | `el8.aarch64` | pgdg | 34.5 KiB | [credcheck_13-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/credcheck_13-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_13` | 4.1 | `el9.x86_64` | pgdg | 39.1 KiB | [credcheck_13-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 3.0 | `el9.x86_64` | pgdg | 35.8 KiB | [credcheck_13-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 2.7 | `el9.x86_64` | pgdg | 35.0 KiB | [credcheck_13-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 2.6 | `el9.x86_64` | pgdg | 34.6 KiB | [credcheck_13-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 2.2 | `el9.x86_64` | pgdg | 33.2 KiB | [credcheck_13-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 2.1 | `el9.x86_64` | pgdg | 32.1 KiB | [credcheck_13-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_13` | 2.0 | `el9.x86_64` | pgdg | 31.2 KiB | [credcheck_13-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_13` | 1.2 | `el9.x86_64` | pgdg | 27.8 KiB | [credcheck_13-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_13` | 1.0 | `el9.x86_64` | pgdg | 27.2 KiB | [credcheck_13-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_13` | 0.2.0 | `el9.x86_64` | pgdg | 18.7 KiB | [credcheck_13-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/credcheck_13-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_13` | 4.1 | `el9.aarch64` | pgdg | 38.4 KiB | [credcheck_13-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 3.0 | `el9.aarch64` | pgdg | 35.6 KiB | [credcheck_13-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 2.7 | `el9.aarch64` | pgdg | 34.6 KiB | [credcheck_13-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 2.6 | `el9.aarch64` | pgdg | 34.3 KiB | [credcheck_13-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 2.2 | `el9.aarch64` | pgdg | 32.7 KiB | [credcheck_13-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 2.1 | `el9.aarch64` | pgdg | 31.5 KiB | [credcheck_13-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_13` | 2.0 | `el9.aarch64` | pgdg | 30.6 KiB | [credcheck_13-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_13` | 1.2 | `el9.aarch64` | pgdg | 27.3 KiB | [credcheck_13-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_13` | 1.0 | `el9.aarch64` | pgdg | 26.7 KiB | [credcheck_13-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_13` | 0.2.0 | `el9.aarch64` | pgdg | 18.1 KiB | [credcheck_13-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_13` | 0.2.0 | `el9.aarch64` | pgdg | 35.2 KiB | [credcheck_13-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/credcheck_13-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_13` | 4.1 | `el10.x86_64` | pgdg | 39.7 KiB | [credcheck_13-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/credcheck_13-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_13` | 3.0 | `el10.x86_64` | pgdg | 36.3 KiB | [credcheck_13-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/credcheck_13-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_13` | 4.1 | `el10.aarch64` | pgdg | 39.3 KiB | [credcheck_13-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/credcheck_13-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_13` | 3.0 | `el10.aarch64` | pgdg | 36.2 KiB | [credcheck_13-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/credcheck_13-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-credcheck` | 3.0 | `d12.x86_64` | pgdg | 63.0 KiB | [postgresql-13-credcheck_3.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-credcheck` | 3.0 | `d12.aarch64` | pgdg | 62.3 KiB | [postgresql-13-credcheck_3.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-credcheck` | 3.0 | `d13.x86_64` | pgdg | 63.2 KiB | [postgresql-13-credcheck_3.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg13+1_amd64.deb) |
| `postgresql-13-credcheck` | 3.0 | `d13.aarch64` | pgdg | 62.3 KiB | [postgresql-13-credcheck_3.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg13+1_arm64.deb) |
| `postgresql-13-credcheck` | 3.0 | `u22.x86_64` | pgdg | 64.7 KiB | [postgresql-13-credcheck_3.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-credcheck` | 3.0 | `u22.aarch64` | pgdg | 63.7 KiB | [postgresql-13-credcheck_3.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-credcheck` | 3.0 | `u24.x86_64` | pgdg | 58.3 KiB | [postgresql-13-credcheck_3.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-credcheck` | 3.0 | `u24.aarch64` | pgdg | 57.6 KiB | [postgresql-13-credcheck_3.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-13-credcheck_3.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MigOpsRepos/credcheck" title="Repository" icon="github" subtitle="github.com/MigOpsRepos/credcheck" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install credcheck; # install by extension name, for the current active PG version
pig ext install credcheck; # install via package alias, for the active PG version
pig ext install credcheck -v 18;   # install for PG 18
pig ext install credcheck -v 17;   # install for PG 17
pig ext install credcheck -v 16;   # install for PG 16
pig ext install credcheck -v 15;   # install for PG 15
pig ext install credcheck -v 14;   # install for PG 14
pig ext install credcheck -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION credcheck;
```

