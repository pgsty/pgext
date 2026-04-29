---
title: "plsh"
linkTitle: "plsh"
description: "PL/sh procedural language"
weight: 3080
categories: ["LANG"]
width: full
---

[**plsh**](https://github.com/petere/plsh) : PL/sh procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3080** | {{< badge content="plsh" link="https://github.com/petere/plsh" >}} | {{< ext "plsh" >}} | `1.20220917` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_tle" >}} {{< ext "plperl" >}} {{< ext "plperlu" >}} {{< ext "plpython3u" >}} {{< ext "plv8" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20220917` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plsh` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20220917` | {{< bg "18" "plsh_18" "green" >}} {{< bg "17" "plsh_17" "green" >}} {{< bg "16" "plsh_16" "green" >}} {{< bg "15" "plsh_15" "green" >}} {{< bg "14" "plsh_14" "green" >}} | `plsh_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20220917` | {{< bg "18" "postgresql-18-plsh" "green" >}} {{< bg "17" "postgresql-17-plsh" "green" >}} {{< bg "16" "postgresql-16-plsh" "green" >}} {{< bg "15" "postgresql-15-plsh" "green" >}} {{< bg "14" "postgresql-14-plsh" "green" >}} | `postgresql-$v-plsh` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.20220917" "plsh_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "plsh_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.20220917" "postgresql-18-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-17-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-16-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-15-plsh : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20220917" "postgresql-14-plsh : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plsh_18` | `1.20220917` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.4 KiB | [plsh_18-1.20220917-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plsh_18-1.20220917-7PGDG.rhel8.x86_64.rpm) |
| `plsh_18` | `1.20220917` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.1 KiB | [plsh_18-1.20220917-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plsh_18-1.20220917-7PGDG.rhel8.aarch64.rpm) |
| `plsh_18` | `1.20220917` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [plsh_18-1.20220917-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plsh_18-1.20220917-7PGDG.rhel9.x86_64.rpm) |
| `plsh_18` | `1.20220917` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [plsh_18-1.20220917-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plsh_18-1.20220917-7PGDG.rhel9.aarch64.rpm) |
| `plsh_18` | `1.20220917` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.9 KiB | [plsh_18-1.20220917-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plsh_18-1.20220917-7PGDG.rhel10.x86_64.rpm) |
| `plsh_18` | `1.20220917` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [plsh_18-1.20220917-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plsh_18-1.20220917-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plsh` | `1.20220917` | [d12.x86_64](/os/d12.x86_64) | pgdg | 27.6 KiB | [postgresql-18-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [d12.aarch64](/os/d12.aarch64) | pgdg | 27.1 KiB | [postgresql-18-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [d13.x86_64](/os/d13.x86_64) | pgdg | 27.7 KiB | [postgresql-18-plsh_1.20220917-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg13+1_amd64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [d13.aarch64](/os/d13.aarch64) | pgdg | 27.2 KiB | [postgresql-18-plsh_1.20220917-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg13+1_arm64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u22.x86_64](/os/u22.x86_64) | pgdg | 29.5 KiB | [postgresql-18-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u22.aarch64](/os/u22.aarch64) | pgdg | 28.7 KiB | [postgresql-18-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.6 KiB | [postgresql-18-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u24.aarch64](/os/u24.aarch64) | pgdg | 27.2 KiB | [postgresql-18-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u26.x86_64](/os/u26.x86_64) | pgdg | 27.8 KiB | [postgresql-18-plsh_1.20220917-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plsh` | `1.20220917` | [u26.aarch64](/os/u26.aarch64) | pgdg | 27.3 KiB | [postgresql-18-plsh_1.20220917-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-18-plsh_1.20220917-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plsh_17` | `1.20220917` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.4 KiB | [plsh_17-1.20220917-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plsh_17-1.20220917-6PGDG.rhel8.x86_64.rpm) |
| `plsh_17` | `1.20220917` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [plsh_17-1.20220917-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plsh_17-1.20220917-6PGDG.rhel8.aarch64.rpm) |
| `plsh_17` | `1.20220917` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.7 KiB | [plsh_17-1.20220917-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plsh_17-1.20220917-6PGDG.rhel9.x86_64.rpm) |
| `plsh_17` | `1.20220917` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [plsh_17-1.20220917-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plsh_17-1.20220917-6PGDG.rhel9.aarch64.rpm) |
| `plsh_17` | `1.20220917` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.9 KiB | [plsh_17-1.20220917-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plsh_17-1.20220917-7PGDG.rhel10.x86_64.rpm) |
| `plsh_17` | `1.20220917` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [plsh_17-1.20220917-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plsh_17-1.20220917-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plsh` | `1.20220917` | [d12.x86_64](/os/d12.x86_64) | pgdg | 27.5 KiB | [postgresql-17-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [d12.aarch64](/os/d12.aarch64) | pgdg | 27.0 KiB | [postgresql-17-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [d13.x86_64](/os/d13.x86_64) | pgdg | 27.6 KiB | [postgresql-17-plsh_1.20220917-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg13+1_amd64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [d13.aarch64](/os/d13.aarch64) | pgdg | 27.1 KiB | [postgresql-17-plsh_1.20220917-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg13+1_arm64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u22.x86_64](/os/u22.x86_64) | pgdg | 33.9 KiB | [postgresql-17-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u22.aarch64](/os/u22.aarch64) | pgdg | 33.2 KiB | [postgresql-17-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.6 KiB | [postgresql-17-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u24.aarch64](/os/u24.aarch64) | pgdg | 27.2 KiB | [postgresql-17-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u26.x86_64](/os/u26.x86_64) | pgdg | 27.7 KiB | [postgresql-17-plsh_1.20220917-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plsh` | `1.20220917` | [u26.aarch64](/os/u26.aarch64) | pgdg | 27.3 KiB | [postgresql-17-plsh_1.20220917-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-17-plsh_1.20220917-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plsh_16` | `1.20220917` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.2 KiB | [plsh_16-1.20220917-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plsh_16-1.20220917-4PGDG.rhel8.x86_64.rpm) |
| `plsh_16` | `1.20220917` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [plsh_16-1.20220917-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plsh_16-1.20220917-4PGDG.rhel8.aarch64.rpm) |
| `plsh_16` | `1.20220917` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.4 KiB | [plsh_16-1.20220917-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plsh_16-1.20220917-4PGDG.rhel9.x86_64.rpm) |
| `plsh_16` | `1.20220917` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [plsh_16-1.20220917-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plsh_16-1.20220917-4PGDG.rhel9.aarch64.rpm) |
| `plsh_16` | `1.20220917` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [plsh_16-1.20220917-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plsh_16-1.20220917-7PGDG.rhel10.x86_64.rpm) |
| `plsh_16` | `1.20220917` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [plsh_16-1.20220917-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plsh_16-1.20220917-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plsh` | `1.20220917` | [d12.x86_64](/os/d12.x86_64) | pgdg | 27.5 KiB | [postgresql-16-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [d12.aarch64](/os/d12.aarch64) | pgdg | 27.1 KiB | [postgresql-16-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [d13.x86_64](/os/d13.x86_64) | pgdg | 27.6 KiB | [postgresql-16-plsh_1.20220917-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg13+1_amd64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [d13.aarch64](/os/d13.aarch64) | pgdg | 27.1 KiB | [postgresql-16-plsh_1.20220917-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg13+1_arm64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u22.x86_64](/os/u22.x86_64) | pgdg | 33.6 KiB | [postgresql-16-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u22.aarch64](/os/u22.aarch64) | pgdg | 32.8 KiB | [postgresql-16-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.6 KiB | [postgresql-16-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u24.aarch64](/os/u24.aarch64) | pgdg | 27.1 KiB | [postgresql-16-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u26.x86_64](/os/u26.x86_64) | pgdg | 27.7 KiB | [postgresql-16-plsh_1.20220917-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plsh` | `1.20220917` | [u26.aarch64](/os/u26.aarch64) | pgdg | 27.3 KiB | [postgresql-16-plsh_1.20220917-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-16-plsh_1.20220917-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plsh_15` | `1.20220917` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.8 KiB | [plsh_15-1.20220917-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plsh_15-1.20220917-1.rhel8.x86_64.rpm) |
| `plsh_15` | `1.20220917` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.5 KiB | [plsh_15-1.20220917-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plsh_15-1.20220917-1.rhel8.aarch64.rpm) |
| `plsh_15` | `1.20220917` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.6 KiB | [plsh_15-1.20220917-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plsh_15-1.20220917-1.rhel9.x86_64.rpm) |
| `plsh_15` | `1.20220917` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [plsh_15-1.20220917-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plsh_15-1.20220917-1.rhel9.aarch64.rpm) |
| `plsh_15` | `1.20220917` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.8 KiB | [plsh_15-1.20220917-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plsh_15-1.20220917-7PGDG.rhel10.x86_64.rpm) |
| `plsh_15` | `1.20220917` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [plsh_15-1.20220917-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plsh_15-1.20220917-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plsh` | `1.20220917` | [d12.x86_64](/os/d12.x86_64) | pgdg | 27.2 KiB | [postgresql-15-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [d12.aarch64](/os/d12.aarch64) | pgdg | 26.8 KiB | [postgresql-15-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [d13.x86_64](/os/d13.x86_64) | pgdg | 27.4 KiB | [postgresql-15-plsh_1.20220917-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg13+1_amd64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [d13.aarch64](/os/d13.aarch64) | pgdg | 26.9 KiB | [postgresql-15-plsh_1.20220917-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg13+1_arm64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u22.x86_64](/os/u22.x86_64) | pgdg | 33.3 KiB | [postgresql-15-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u22.aarch64](/os/u22.aarch64) | pgdg | 32.6 KiB | [postgresql-15-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.3 KiB | [postgresql-15-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.9 KiB | [postgresql-15-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u26.x86_64](/os/u26.x86_64) | pgdg | 27.5 KiB | [postgresql-15-plsh_1.20220917-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plsh` | `1.20220917` | [u26.aarch64](/os/u26.aarch64) | pgdg | 27.0 KiB | [postgresql-15-plsh_1.20220917-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-15-plsh_1.20220917-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plsh_14` | `1.20220917` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.8 KiB | [plsh_14-1.20220917-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plsh_14-1.20220917-1.rhel8.x86_64.rpm) |
| `plsh_14` | `1.20200522` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [plsh_14-1.20200522-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plsh_14-1.20200522-3.rhel8.x86_64.rpm) |
| `plsh_14` | `1.20220917` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.5 KiB | [plsh_14-1.20220917-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plsh_14-1.20220917-1.rhel8.aarch64.rpm) |
| `plsh_14` | `1.20220917` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.6 KiB | [plsh_14-1.20220917-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plsh_14-1.20220917-1.rhel9.x86_64.rpm) |
| `plsh_14` | `1.20220917` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [plsh_14-1.20220917-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plsh_14-1.20220917-1.rhel9.aarch64.rpm) |
| `plsh_14` | `1.20220917` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.9 KiB | [plsh_14-1.20220917-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plsh_14-1.20220917-7PGDG.rhel10.x86_64.rpm) |
| `plsh_14` | `1.20220917` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.6 KiB | [plsh_14-1.20220917-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plsh_14-1.20220917-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plsh` | `1.20220917` | [d12.x86_64](/os/d12.x86_64) | pgdg | 27.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_amd64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [d12.aarch64](/os/d12.aarch64) | pgdg | 26.8 KiB | [postgresql-14-plsh_1.20220917-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg12+1_arm64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [d13.x86_64](/os/d13.x86_64) | pgdg | 27.3 KiB | [postgresql-14-plsh_1.20220917-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg13+1_amd64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [d13.aarch64](/os/d13.aarch64) | pgdg | 26.8 KiB | [postgresql-14-plsh_1.20220917-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg13+1_arm64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u22.x86_64](/os/u22.x86_64) | pgdg | 33.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u22.aarch64](/os/u22.aarch64) | pgdg | 32.5 KiB | [postgresql-14-plsh_1.20220917-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.2 KiB | [postgresql-14-plsh_1.20220917-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.8 KiB | [postgresql-14-plsh_1.20220917-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u26.x86_64](/os/u26.x86_64) | pgdg | 27.5 KiB | [postgresql-14-plsh_1.20220917-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plsh` | `1.20220917` | [u26.aarch64](/os/u26.aarch64) | pgdg | 27.0 KiB | [postgresql-14-plsh_1.20220917-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plsh/postgresql-14-plsh_1.20220917-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/plsh" title="Repository" icon="github" subtitle="github.com/petere/plsh" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plsh;		# install via package name, for the active PG version

pig install plsh -v 18;   # install for PG 18
pig install plsh -v 17;   # install for PG 17
pig install plsh -v 16;   # install for PG 16
pig install plsh -v 15;   # install for PG 15
pig install plsh -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plsh;
```




## Usage

> [plsh: PL/sh procedural language](https://github.com/petere/plsh)

`plsh` allows writing PostgreSQL functions as shell scripts. The function body must start with a shebang line specifying the interpreter.

```sql
CREATE EXTENSION plsh;
```

### Create Shell Functions

```sql
CREATE FUNCTION concat(text, text) RETURNS text AS '
#!/bin/sh
echo "$1$2"
' LANGUAGE plsh;

SELECT concat('hello ', 'world');  -- 'hello world'
```

### Argument Passing

Function arguments are passed as positional shell variables (`$1`, `$2`, etc.):

```sql
CREATE FUNCTION file_line_count(filename text) RETURNS int AS '
#!/bin/sh
wc -l < "$1"
' LANGUAGE plsh;
```

### Return Values

- Standard output becomes the return value (trailing newlines stripped)
- Empty output returns NULL
- Standard error output causes the function to abort with that error message
- Non-zero exit status triggers an error

```sql
CREATE FUNCTION system_uptime() RETURNS text AS '
#!/bin/sh
uptime
' LANGUAGE plsh;
```

### Database Access

Direct SPI access is not available, but `psql` can be used since libpq environment variables are preconfigured:

```sql
CREATE FUNCTION query_db(x int) RETURNS text AS $$
#!/bin/sh
psql -At -c "SELECT name FROM users WHERE id = $1"
$$  LANGUAGE plsh;
```

### Trigger Functions

Trigger context is available via environment variables:

| Variable | Description |
|----------|-------------|
| `PLSH_TG_NAME` | Trigger name |
| `PLSH_TG_WHEN` | `BEFORE`, `INSTEAD OF`, or `AFTER` |
| `PLSH_TG_LEVEL` | `ROW` or `STATEMENT` |
| `PLSH_TG_OP` | `DELETE`, `INSERT`, `UPDATE`, or `TRUNCATE` |
| `PLSH_TG_TABLE_NAME` | Target table name |
| `PLSH_TG_TABLE_SCHEMA` | Target table schema |

Event trigger variables: `PLSH_TG_EVENT`, `PLSH_TG_TAG`.

### Inline Execution

```sql
DO E'#!/bin/sh\necho "running shell command"' LANGUAGE plsh;
```

### Security

`plsh` should not be declared as `TRUSTED` since shell scripts have full OS-level access under the PostgreSQL user. Only superusers should create `plsh` functions.
