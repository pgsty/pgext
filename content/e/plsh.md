---
title: "plsh"
linkTitle: "plsh"
description: "PL/sh procedural language"
weight: 3080
categories: ["Lang"]
width: full
---

PL/sh procedural language

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3080** | {{< badge content="plsh" link="https://github.com/petere/plsh" >}} | {{< ext "plsh" "plsh" >}} | `1.20220917` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_tle" >}} {{< ext "plperl" >}} {{< ext "plperlu" >}} {{< ext "plpython3u" >}} {{< ext "plv8" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/plsh" >}} | `1.20220917` | {{< badge content="18" color="red" alt="plsh_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `plsh_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/plsh" >}} | `1.20220917` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-plsh` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "plsh_18" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plsh_18-1.20220917-7PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "plsh_17" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plsh_17-1.20220917-6PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "plsh_16" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plsh_16-1.20220917-4PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "plsh_15" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plsh_15-1.20220917-1.rhel8.x86_64.rpm" >}} | {{< pkg "plsh_14" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plsh_14-1.20220917-1.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "plsh_18" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plsh_18-1.20220917-7PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "plsh_17" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plsh_17-1.20220917-6PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "plsh_16" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plsh_16-1.20220917-4PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "plsh_15" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plsh_15-1.20220917-1.rhel8.aarch64.rpm" >}} | {{< pkg "plsh_14" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plsh_14-1.20220917-1.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "plsh_18" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plsh_18-1.20220917-7PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "plsh_17" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plsh_17-1.20220917-6PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "plsh_16" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plsh_16-1.20220917-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "plsh_15" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plsh_15-1.20220917-1.rhel9.x86_64.rpm" >}} | {{< pkg "plsh_14" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plsh_14-1.20220917-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "plsh_18" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plsh_18-1.20220917-7PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "plsh_17" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plsh_17-1.20220917-6PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "plsh_16" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plsh_16-1.20220917-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "plsh_15" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plsh_15-1.20220917-1.rhel9.aarch64.rpm" >}} | {{< pkg "plsh_14" "1.20220917" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plsh_14-1.20220917-1.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-plsh" "1.20220917" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_18` | 1.20220917 | `el8.x86_64` | pgdg | 22.4 KiB | [plsh_18-1.20220917-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plsh_18-1.20220917-7PGDG.rhel8.x86_64.rpm) |
| `plsh_18` | 1.20220917 | `el8.aarch64` | pgdg | 22.1 KiB | [plsh_18-1.20220917-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plsh_18-1.20220917-7PGDG.rhel8.aarch64.rpm) |
| `plsh_18` | 1.20220917 | `el9.x86_64` | pgdg | 21.5 KiB | [plsh_18-1.20220917-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plsh_18-1.20220917-7PGDG.rhel9.x86_64.rpm) |
| `plsh_18` | 1.20220917 | `el9.aarch64` | pgdg | 20.9 KiB | [plsh_18-1.20220917-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plsh_18-1.20220917-7PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.6 KiB | [postgresql-18-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-18-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 27.1 KiB | [postgresql-18-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-18-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 29.5 KiB | [postgresql-18-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 28.7 KiB | [postgresql-18-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 27.2 KiB | [postgresql-18-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.6 KiB | [postgresql-18-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_17` | 1.20220917 | `el8.x86_64` | pgdg | 22.4 KiB | [plsh_17-1.20220917-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plsh_17-1.20220917-6PGDG.rhel8.x86_64.rpm) |
| `plsh_17` | 1.20220917 | `el8.aarch64` | pgdg | 22.2 KiB | [plsh_17-1.20220917-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plsh_17-1.20220917-6PGDG.rhel8.aarch64.rpm) |
| `plsh_17` | 1.20220917 | `el9.aarch64` | pgdg | 21.3 KiB | [plsh_17-1.20220917-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plsh_17-1.20220917-6PGDG.rhel9.aarch64.rpm) |
| `plsh_17` | 1.20220917 | `el9.x86_64` | pgdg | 21.7 KiB | [plsh_17-1.20220917-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plsh_17-1.20220917-6PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 27.0 KiB | [postgresql-17-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-17-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.5 KiB | [postgresql-17-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-17-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 33.9 KiB | [postgresql-17-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 33.2 KiB | [postgresql-17-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 27.2 KiB | [postgresql-17-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.6 KiB | [postgresql-17-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_16` | 1.20220917 | `el8.x86_64` | pgdg | 22.2 KiB | [plsh_16-1.20220917-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plsh_16-1.20220917-4PGDG.rhel8.x86_64.rpm) |
| `plsh_16` | 1.20220917 | `el8.aarch64` | pgdg | 22.0 KiB | [plsh_16-1.20220917-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plsh_16-1.20220917-4PGDG.rhel8.aarch64.rpm) |
| `plsh_16` | 1.20220917 | `el9.aarch64` | pgdg | 20.9 KiB | [plsh_16-1.20220917-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plsh_16-1.20220917-4PGDG.rhel9.aarch64.rpm) |
| `plsh_16` | 1.20220917 | `el9.x86_64` | pgdg | 21.4 KiB | [plsh_16-1.20220917-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plsh_16-1.20220917-4PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 27.1 KiB | [postgresql-16-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-16-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.5 KiB | [postgresql-16-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-16-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 33.6 KiB | [postgresql-16-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 32.8 KiB | [postgresql-16-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 27.1 KiB | [postgresql-16-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.6 KiB | [postgresql-16-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_15` | 1.20220917 | `el8.aarch64` | pgdg | 21.5 KiB | [plsh_15-1.20220917-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plsh_15-1.20220917-1.rhel8.aarch64.rpm) |
| `plsh_15` | 1.20220917 | `el8.x86_64` | pgdg | 21.8 KiB | [plsh_15-1.20220917-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plsh_15-1.20220917-1.rhel8.x86_64.rpm) |
| `plsh_15` | 1.20220917 | `el9.x86_64` | pgdg | 21.6 KiB | [plsh_15-1.20220917-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plsh_15-1.20220917-1.rhel9.x86_64.rpm) |
| `plsh_15` | 1.20220917 | `el9.aarch64` | pgdg | 20.9 KiB | [plsh_15-1.20220917-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plsh_15-1.20220917-1.rhel9.aarch64.rpm) |
| `postgresql-15-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.2 KiB | [postgresql-15-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-15-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 26.8 KiB | [postgresql-15-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-15-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 32.6 KiB | [postgresql-15-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 33.3 KiB | [postgresql-15-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 26.9 KiB | [postgresql-15-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.3 KiB | [postgresql-15-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_14` | 1.20220917 | `el8.aarch64` | pgdg | 21.5 KiB | [plsh_14-1.20220917-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plsh_14-1.20220917-1.rhel8.aarch64.rpm) |
| `plsh_14` | 1.20220917 | `el8.x86_64` | pgdg | 21.8 KiB | [plsh_14-1.20220917-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plsh_14-1.20220917-1.rhel8.x86_64.rpm) |
| `plsh_14` | 1.20200522 | `el8.x86_64` | pgdg | 41.8 KiB | [plsh_14-1.20200522-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plsh_14-1.20200522-3.rhel8.x86_64.rpm) |
| `plsh_14` | 1.20220917 | `el9.aarch64` | pgdg | 20.9 KiB | [plsh_14-1.20220917-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plsh_14-1.20220917-1.rhel9.aarch64.rpm) |
| `plsh_14` | 1.20220917 | `el9.x86_64` | pgdg | 21.6 KiB | [plsh_14-1.20220917-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plsh_14-1.20220917-1.rhel9.x86_64.rpm) |
| `postgresql-14-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 26.8 KiB | [postgresql-14-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-14-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-14-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 33.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 32.5 KiB | [postgresql-14-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 26.8 KiB | [postgresql-14-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `plsh_13` | 1.20220917 | `el8.x86_64` | pgdg | 21.6 KiB | [plsh_13-1.20220917-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plsh_13-1.20220917-1.rhel8.x86_64.rpm) |
| `plsh_13` | 1.20220917 | `el8.aarch64` | pgdg | 21.5 KiB | [plsh_13-1.20220917-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plsh_13-1.20220917-1.rhel8.aarch64.rpm) |
| `plsh_13` | 1.20220917 | `el9.x86_64` | pgdg | 21.5 KiB | [plsh_13-1.20220917-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plsh_13-1.20220917-1.rhel9.x86_64.rpm) |
| `plsh_13` | 1.20220917 | `el9.aarch64` | pgdg | 20.9 KiB | [plsh_13-1.20220917-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plsh_13-1.20220917-1.rhel9.aarch64.rpm) |
| `postgresql-13-plsh` | 1.20220917 | `d12.aarch64` | pgdg | 27.0 KiB | [postgresql-13-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-13-plsh` | 1.20220917 | `d12.x86_64` | pgdg | 27.3 KiB | [postgresql-13-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-13-plsh` | 1.20220917 | `u22.x86_64` | pgdg | 32.9 KiB | [postgresql-13-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-plsh` | 1.20220917 | `u22.aarch64` | pgdg | 32.5 KiB | [postgresql-13-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-plsh` | 1.20220917 | `u24.aarch64` | pgdg | 27.0 KiB | [postgresql-13-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-13-plsh` | 1.20220917 | `u24.x86_64` | pgdg | 27.3 KiB | [postgresql-13-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-13-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/plsh" title="Repository" icon="github" subtitle="github.com/petere/plsh" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plsh; # install by extension name, for the current active PG version
pig ext install plsh; # install via package alias, for the active PG version
pig ext install plsh -v 18;   # install for PG 18
pig ext install plsh -v 17;   # install for PG 17
pig ext install plsh -v 16;   # install for PG 16
pig ext install plsh -v 15;   # install for PG 15
pig ext install plsh -v 14;   # install for PG 14
pig ext install plsh -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plsh;
```

