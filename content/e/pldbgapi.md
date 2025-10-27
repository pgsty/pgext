---
title: "pldbgapi"
linkTitle: "pldbgapi"
description: "server-side support for debugging PL/pgSQL functions"
weight: 3050
categories: ["LANG"]
width: full
---

server-side support for debugging PL/pgSQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3050** | {{< badge content="pldbgapi" link="https://github.com/EnterpriseDB/pldebugger" >}} | {{< ext "pldbgapi" "pldebugger" >}} | `1.9` | {{< category "LANG" >}} | {{< license "Artistic" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql_check" >}} {{< ext "plprofiler" >}} {{< ext "plpgsql" >}} {{< ext "pgtap" >}} {{< ext "pg_stat_statements" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pldbgapi" >}} | `1.9` | {{< bg "18" "pldebugger_18*" "green" >}} {{< bg "17" "pldebugger_17*" "green" >}} {{< bg "16" "pldebugger_16*" "green" >}} {{< bg "15" "pldebugger_15*" "green" >}} {{< bg "14" "pldebugger_14*" "green" >}} | `pldebugger_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pldbgapi" >}} | `1.9` | {{< bg "18" "postgresql-18-pldebugger" "green" >}} {{< bg "17" "postgresql-17-pldebugger" "green" >}} {{< bg "16" "postgresql-16-pldebugger" "green" >}} {{< bg "15" "postgresql-15-pldebugger" "green" >}} {{< bg "14" "postgresql-14-pldebugger" "green" >}} | `postgresql-$v-pldebugger` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.9" "pldebugger_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_14 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.9" "pldebugger_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_14 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.9" "pldebugger_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_14 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.9" "pldebugger_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.8" "pldebugger_14 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.9" "postgresql-18-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-17-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-16-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-15-pldebugger : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9" "postgresql-14-pldebugger : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pldebugger_18` | 1.9 | `el8.x86_64` | pgdg | 38.7 KiB | [pldebugger_18-1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pldebugger_18-1.9-1PGDG.rhel8.x86_64.rpm) |
| `pldebugger_18` | 1.9 | `el8.aarch64` | pgdg | 37.8 KiB | [pldebugger_18-1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pldebugger_18-1.9-1PGDG.rhel8.aarch64.rpm) |
| `pldebugger_18` | 1.9 | `el9.x86_64` | pgdg | 36.9 KiB | [pldebugger_18-1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pldebugger_18-1.9-1PGDG.rhel9.x86_64.rpm) |
| `pldebugger_18` | 1.9 | `el9.aarch64` | pgdg | 36.1 KiB | [pldebugger_18-1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pldebugger_18-1.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pldebugger` | 1.9 | `d12.x86_64` | pgdg | 71.4 KiB | [postgresql-18-pldebugger_1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pldebugger` | 1.9 | `d12.aarch64` | pgdg | 70.0 KiB | [postgresql-18-pldebugger_1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pldebugger` | 1.9 | `u22.x86_64` | pgdg | 72.4 KiB | [postgresql-18-pldebugger_1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pldebugger` | 1.9 | `u22.aarch64` | pgdg | 71.1 KiB | [postgresql-18-pldebugger_1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pldebugger` | 1.9 | `u24.x86_64` | pgdg | 70.5 KiB | [postgresql-18-pldebugger_1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pldebugger` | 1.9 | `u24.aarch64` | pgdg | 68.7 KiB | [postgresql-18-pldebugger_1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-18-pldebugger_1.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pldebugger_17` | 1.8 | `el8.x86_64` | pgdg | 38.5 KiB | [pldebugger_17-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pldebugger_17-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pldebugger_17` | 1.8 | `el8.aarch64` | pgdg | 37.5 KiB | [pldebugger_17-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pldebugger_17-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pldebugger_17` | 1.8 | `el9.x86_64` | pgdg | 37.0 KiB | [pldebugger_17-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pldebugger_17-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pldebugger_17` | 1.8 | `el9.aarch64` | pgdg | 36.3 KiB | [pldebugger_17-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pldebugger_17-1.8-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pldebugger` | 1.9 | `d12.x86_64` | pgdg | 71.4 KiB | [postgresql-17-pldebugger_1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pldebugger` | 1.9 | `d12.aarch64` | pgdg | 70.0 KiB | [postgresql-17-pldebugger_1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pldebugger` | 1.9 | `u22.x86_64` | pgdg | 83.0 KiB | [postgresql-17-pldebugger_1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pldebugger` | 1.9 | `u22.aarch64` | pgdg | 81.4 KiB | [postgresql-17-pldebugger_1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pldebugger` | 1.9 | `u24.x86_64` | pgdg | 70.5 KiB | [postgresql-17-pldebugger_1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pldebugger` | 1.9 | `u24.aarch64` | pgdg | 68.7 KiB | [postgresql-17-pldebugger_1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-17-pldebugger_1.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pldebugger_16` | 1.8 | `el8.x86_64` | pgdg | 38.5 KiB | [pldebugger_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pldebugger_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pldebugger_16` | 1.5 | `el8.x86_64` | pgdg | 38.2 KiB | [pldebugger_16-1.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pldebugger_16-1.5-3PGDG.rhel8.x86_64.rpm) |
| `pldebugger_16` | 1.8 | `el8.aarch64` | pgdg | 37.5 KiB | [pldebugger_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pldebugger_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pldebugger_16` | 1.5 | `el8.aarch64` | pgdg | 37.2 KiB | [pldebugger_16-1.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pldebugger_16-1.5-3PGDG.rhel8.aarch64.rpm) |
| `pldebugger_16` | 1.8 | `el9.x86_64` | pgdg | 37.1 KiB | [pldebugger_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pldebugger_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pldebugger_16` | 1.5 | `el9.x86_64` | pgdg | 36.9 KiB | [pldebugger_16-1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pldebugger_16-1.5-3PGDG.rhel9.x86_64.rpm) |
| `pldebugger_16` | 1.8 | `el9.aarch64` | pgdg | 36.3 KiB | [pldebugger_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pldebugger_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pldebugger_16` | 1.5 | `el9.aarch64` | pgdg | 35.9 KiB | [pldebugger_16-1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pldebugger_16-1.5-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pldebugger` | 1.9 | `d12.x86_64` | pgdg | 71.3 KiB | [postgresql-16-pldebugger_1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pldebugger` | 1.9 | `d12.aarch64` | pgdg | 69.9 KiB | [postgresql-16-pldebugger_1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pldebugger` | 1.9 | `u22.x86_64` | pgdg | 82.4 KiB | [postgresql-16-pldebugger_1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pldebugger` | 1.9 | `u22.aarch64` | pgdg | 81.0 KiB | [postgresql-16-pldebugger_1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pldebugger` | 1.9 | `u24.x86_64` | pgdg | 70.5 KiB | [postgresql-16-pldebugger_1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pldebugger` | 1.9 | `u24.aarch64` | pgdg | 68.6 KiB | [postgresql-16-pldebugger_1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-16-pldebugger_1.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pldebugger_15` | 1.8 | `el8.x86_64` | pgdg | 39.1 KiB | [pldebugger_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pldebugger_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pldebugger_15` | 1.5 | `el8.x86_64` | pgdg | 96.8 KiB | [pldebugger_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pldebugger_15-1.5-1.rhel8.x86_64.rpm) |
| `pldebugger_15` | 1.8 | `el8.aarch64` | pgdg | 38.0 KiB | [pldebugger_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pldebugger_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pldebugger_15` | 1.5 | `el8.aarch64` | pgdg | 95.1 KiB | [pldebugger_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pldebugger_15-1.5-1.rhel8.aarch64.rpm) |
| `pldebugger_15` | 1.8 | `el9.x86_64` | pgdg | 39.2 KiB | [pldebugger_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pldebugger_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pldebugger_15` | 1.5 | `el9.x86_64` | pgdg | 98.5 KiB | [pldebugger_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pldebugger_15-1.5-1.rhel9.x86_64.rpm) |
| `pldebugger_15` | 1.8 | `el9.aarch64` | pgdg | 38.5 KiB | [pldebugger_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pldebugger_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pldebugger_15` | 1.5 | `el9.aarch64` | pgdg | 97.3 KiB | [pldebugger_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pldebugger_15-1.5-1.rhel9.aarch64.rpm) |
| `postgresql-15-pldebugger` | 1.9 | `d12.x86_64` | pgdg | 71.8 KiB | [postgresql-15-pldebugger_1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pldebugger` | 1.9 | `d12.aarch64` | pgdg | 70.4 KiB | [postgresql-15-pldebugger_1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pldebugger` | 1.9 | `u22.x86_64` | pgdg | 83.9 KiB | [postgresql-15-pldebugger_1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pldebugger` | 1.9 | `u22.aarch64` | pgdg | 82.4 KiB | [postgresql-15-pldebugger_1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pldebugger` | 1.9 | `u24.x86_64` | pgdg | 71.8 KiB | [postgresql-15-pldebugger_1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pldebugger` | 1.9 | `u24.aarch64` | pgdg | 70.2 KiB | [postgresql-15-pldebugger_1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-15-pldebugger_1.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pldebugger_14` | 1.8 | `el8.x86_64` | pgdg | 39.1 KiB | [pldebugger_14-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pldebugger_14-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pldebugger_14` | 1.5 | `el8.x86_64` | pgdg | 95.2 KiB | [pldebugger_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pldebugger_14-1.5-1.rhel8.x86_64.rpm) |
| `pldebugger_14` | 1.4 | `el8.x86_64` | pgdg | 95.0 KiB | [pldebugger_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pldebugger_14-1.4-1.rhel8.x86_64.rpm) |
| `pldebugger_14` | 1.8 | `el8.aarch64` | pgdg | 37.9 KiB | [pldebugger_14-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pldebugger_14-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pldebugger_14` | 1.5 | `el8.aarch64` | pgdg | 93.6 KiB | [pldebugger_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pldebugger_14-1.5-1.rhel8.aarch64.rpm) |
| `pldebugger_14` | 1.8 | `el9.x86_64` | pgdg | 39.1 KiB | [pldebugger_14-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pldebugger_14-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pldebugger_14` | 1.5 | `el9.x86_64` | pgdg | 96.8 KiB | [pldebugger_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pldebugger_14-1.5-1.rhel9.x86_64.rpm) |
| `pldebugger_14` | 1.8 | `el9.aarch64` | pgdg | 38.4 KiB | [pldebugger_14-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pldebugger_14-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pldebugger_14` | 1.5 | `el9.aarch64` | pgdg | 95.8 KiB | [pldebugger_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pldebugger_14-1.5-1.rhel9.aarch64.rpm) |
| `postgresql-14-pldebugger` | 1.9 | `d12.x86_64` | pgdg | 71.6 KiB | [postgresql-14-pldebugger_1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pldebugger` | 1.9 | `d12.aarch64` | pgdg | 70.1 KiB | [postgresql-14-pldebugger_1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pldebugger` | 1.9 | `u22.x86_64` | pgdg | 82.3 KiB | [postgresql-14-pldebugger_1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pldebugger` | 1.9 | `u22.aarch64` | pgdg | 81.0 KiB | [postgresql-14-pldebugger_1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pldebugger` | 1.9 | `u24.x86_64` | pgdg | 71.6 KiB | [postgresql-14-pldebugger_1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pldebugger` | 1.9 | `u24.aarch64` | pgdg | 70.1 KiB | [postgresql-14-pldebugger_1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pldebugger/postgresql-14-pldebugger_1.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/pldebugger" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/pldebugger" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pldbgapi; # install by extension name, for the current active PG version
pig ext install pldebugger; # install via package alias, for the active PG version
pig ext install pldbgapi -v 18;   # install for PG 18
pig ext install pldbgapi -v 17;   # install for PG 17
pig ext install pldbgapi -v 16;   # install for PG 16
pig ext install pldbgapi -v 15;   # install for PG 15
pig ext install pldbgapi -v 14;   # install for PG 14
pig ext install pldbgapi -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pldbgapi;
```

