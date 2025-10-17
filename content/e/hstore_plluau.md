---
title: "hstore_plluau"
linkTitle: "hstore_plluau"
description: "Hstore transform for untrusted Lua"
weight: 3031
categories: ["Lang"]
width: full
---

Hstore transform for untrusted Lua

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3031** | {{< badge content="hstore_plluau" link="https://github.com/pllua/pllua" >}} | {{< ext "hstore_plluau" "pllua" >}} | `2.0.12` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} {{< ext "plluau" >}} |
|   **See Also**    | {{< ext "hstore_plperl" >}} {{< ext "hstore_plperlu" >}} {{< ext "hstore_plpython3u" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "hstore_pllua" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} |

> [!Note] missing pg12-15 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **Debian** | {{< badge content="PGDG" link="/e/pllua" >}} | `2.0.12` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pllua` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pllua_18" >}}     | {{< pkg "pllua_17" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pllua_16" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pllua_15" "2.0.11" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.11-1.rhel8.x86_64.rpm" >}} | {{< pkg "pllua_14" "2.0.11" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.11-1.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pllua_18" >}}     | {{< pkg "pllua_17" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pllua_16" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm" >}} |    {{< pkg "pllua_15" >}}     |    {{< pkg "pllua_14" >}}     |
|    `el9.x86_64`    |    {{< pkg "pllua_18" >}}     | {{< pkg "pllua_17" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pllua_16" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pllua_15" "2.0.11" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.11-1.rhel9.x86_64.rpm" >}} | {{< pkg "pllua_14" "2.0.11" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pllua_14-2.0.11-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pllua_18" >}}     | {{< pkg "pllua_17" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pllua_16" "2.0.12" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm" >}} |    {{< pkg "pllua_15" >}}     |    {{< pkg "pllua_14" >}}     |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pllua" "2.0.12" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-18-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.1 KiB | [postgresql-18-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.4 KiB | [postgresql-18-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 354.6 KiB | [postgresql-18-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 342.1 KiB | [postgresql-18-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.3 KiB | [postgresql-18-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 347.2 KiB | [postgresql-18-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pllua_17` | 2.0.12 | `el8.aarch64` | pgdg | 110.9 KiB | [pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm) |
| `pllua_17` | 2.0.12 | `el8.x86_64` | pgdg | 119.4 KiB | [pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm) |
| `pllua_17` | 2.0.12 | `el9.aarch64` | pgdg | 115.3 KiB | [pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm) |
| `pllua_17` | 2.0.12 | `el9.x86_64` | pgdg | 120.4 KiB | [pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 346.9 KiB | [postgresql-17-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.7 KiB | [postgresql-17-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 378.7 KiB | [postgresql-17-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 391.2 KiB | [postgresql-17-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 347.0 KiB | [postgresql-17-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.3 KiB | [postgresql-17-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pllua_16` | 2.0.12 | `el8.aarch64` | pgdg | 110.6 KiB | [pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm) |
| `pllua_16` | 2.0.12 | `el8.x86_64` | pgdg | 119.2 KiB | [pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm) |
| `pllua_16` | 2.0.12 | `el9.x86_64` | pgdg | 120.1 KiB | [pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm) |
| `pllua_16` | 2.0.12 | `el9.aarch64` | pgdg | 115.5 KiB | [pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.9 KiB | [postgresql-16-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.3 KiB | [postgresql-16-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 377.1 KiB | [postgresql-16-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 389.6 KiB | [postgresql-16-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.1 KiB | [postgresql-16-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 346.9 KiB | [postgresql-16-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pllua_15` | 2.0.11 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_15-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_15` | 2.0.10 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_15-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_15` | 2.0.11 | `el9.x86_64` | pgdg | 123.5 KiB | [pllua_15-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_15` | 2.0.10 | `el9.x86_64` | pgdg | 123.8 KiB | [pllua_15-2.0.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.10-1.rhel9.x86_64.rpm) |
| `postgresql-15-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 349.0 KiB | [postgresql-15-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 392.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 379.9 KiB | [postgresql-15-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 337.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pllua_14` | 2.0.11 | `el8.x86_64` | pgdg | 121.1 KiB | [pllua_14-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_14` | 2.0.10 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_14-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_14` | 2.0.11 | `el9.x86_64` | pgdg | 123.2 KiB | [pllua_14-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pllua_14-2.0.11-1.rhel9.x86_64.rpm) |
| `postgresql-14-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.7 KiB | [postgresql-14-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 349.0 KiB | [postgresql-14-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 388.0 KiB | [postgresql-14-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 375.9 KiB | [postgresql-14-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.5 KiB | [postgresql-14-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 336.9 KiB | [postgresql-14-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pllua_13` | 2.0.9 | `el8.x86_64` | pgdg | 120.1 KiB | [pllua_13-2.0.9-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.9-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.11 | `el8.x86_64` | pgdg | 120.4 KiB | [pllua_13-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.10 | `el8.x86_64` | pgdg | 120.2 KiB | [pllua_13-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.11 | `el9.x86_64` | pgdg | 123.2 KiB | [pllua_13-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pllua_13-2.0.11-1.rhel9.x86_64.rpm) |
| `postgresql-13-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 386.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 374.4 KiB | [postgresql-13-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 336.5 KiB | [postgresql-13-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.0 KiB | [postgresql-13-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pllua/pllua" title="Repository" icon="github" subtitle="github.com/pllua/pllua" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hstore_plluau; # install by extension name, for the current active PG version
pig ext install pllua; # install via package alias, for the active PG version
pig ext install hstore_plluau -v 18;   # install for PG 18
pig ext install hstore_plluau -v 17;   # install for PG 17
pig ext install hstore_plluau -v 16;   # install for PG 16
pig ext install hstore_plluau -v 15;   # install for PG 15
pig ext install hstore_plluau -v 14;   # install for PG 14
pig ext install hstore_plluau -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hstore_plluau;
```

