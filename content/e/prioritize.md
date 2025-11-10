---
title: "prioritize"
linkTitle: "prioritize"
description: "get and set the priority of PostgreSQL backends"
weight: 5090
categories: ["ADMIN"]
width: full
---

[**pg_prioritize**](https://github.com/schmiddy/pg_prioritize)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5090** | {{< badge content="prioritize" link="https://github.com/schmiddy/pg_prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | `1.0.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_proctab" >}} {{< ext "pg_background" >}} {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |

> [!Note] no pg 14,13,12 on el9


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/prioritize" >}} | `1.0.4` | {{< bg "18" "pg_prioritize_18*" "green" >}} {{< bg "17" "pg_prioritize_17*" "green" >}} {{< bg "16" "pg_prioritize_16*" "green" >}} {{< bg "15" "pg_prioritize_15*" "green" >}} {{< bg "14" "pg_prioritize_14*" "green" >}} {{< bg "13" "pg_prioritize_13*" "green" >}} | `pg_prioritize_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/prioritize" >}} | `1.0.4` | {{< bg "18" "postgresql-18-prioritize" "green" >}} {{< bg "17" "postgresql-17-prioritize" "green" >}} {{< bg "16" "postgresql-16-prioritize" "green" >}} {{< bg "15" "postgresql-15-prioritize" "green" >}} {{< bg "14" "postgresql-14-prioritize" "green" >}} {{< bg "13" "postgresql-13-prioritize" "green" >}} | `postgresql-$v-prioritize` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_prioritize_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_prioritize_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-13-prioritize : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_18` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.1 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.6 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.4 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.3 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.1 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_17` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.1 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_prioritize_17-1.0.4-5PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_prioritize_17-1.0.4-5PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.0 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_prioritize_17-1.0.4-5PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.7 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_prioritize_17-1.0.4-5PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_17-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_prioritize_17-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_17-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_prioritize_17-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.4 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_16` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_prioritize_16-1.0.4-4PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.9 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_prioritize_16-1.0.4-4PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.8 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_prioritize_16-1.0.4-4PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.3 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_prioritize_16-1.0.4-4PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_16-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_prioritize_16-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_16-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_prioritize_16-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.3 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 KiB | [pg_prioritize_15-1.0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_prioritize_15-1.0.4-2.rhel8.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_15-1.0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_prioritize_15-1.0.4-2.rhel8.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_prioritize_15-1.0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_prioritize_15-1.0.4-2.rhel9.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_15-1.0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_prioritize_15-1.0.4-2.rhel9.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_15-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_prioritize_15-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_15-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_prioritize_15-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.4 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_prioritize_14-1.0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_prioritize_14-1.0.4-2.rhel8.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_14-1.0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_prioritize_14-1.0.4-2.rhel8.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_14-1.0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_prioritize_14-1.0.4-2.rhel9.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_14-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_prioritize_14-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_14-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_prioritize_14-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.3 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_13` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.5 KiB | [pg_prioritize_13-1.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_prioritize_13-1.0.4-1.rhel8.x86_64.rpm) |
| `pg_prioritize_13` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_prioritize_13-1.0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_prioritize_13-1.0.4-2.rhel8.aarch64.rpm) |
| `pg_prioritize_13` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.0 KiB | [pg_prioritize_13-1.0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_prioritize_13-1.0.4-2.rhel9.aarch64.rpm) |
| `pg_prioritize_13` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_13-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_prioritize_13-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_13` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_13-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_prioritize_13-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.6 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.6 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.6 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.4 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.2 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.7 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-13-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.6 KiB | [postgresql-13-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-13-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/schmiddy/pg_prioritize" title="Repository" icon="github" subtitle="github.com/schmiddy/pg_prioritize" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install prioritize; # install by extension name, for the current active PG version
pig ext install pg_prioritize; # install via package alias, for the active PG version
pig ext install prioritize -v 18;   # install for PG 18
pig ext install prioritize -v 17;   # install for PG 17
pig ext install prioritize -v 16;   # install for PG 16
pig ext install prioritize -v 15;   # install for PG 15
pig ext install prioritize -v 14;   # install for PG 14
pig ext install prioritize -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION prioritize;
```

