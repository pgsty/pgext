---
title: "pgq"
linkTitle: "pgq"
description: "Generic queue for PostgreSQL"
weight: 2910
categories: ["FEAT"]
width: full
---

Generic queue for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2910** | {{< badge content="pgq" link="https://github.com/pgq/pgq" >}} | {{< ext "pgq" >}} | `3.5.1` | {{< category "FEAT" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgq" >}} | `3.5.1` | {{< bg "18" "pgq_18*" "green" >}} {{< bg "17" "pgq_17*" "green" >}} {{< bg "16" "pgq_16*" "green" >}} {{< bg "15" "pgq_15*" "green" >}} {{< bg "14" "pgq_14*" "green" >}} | `pgq_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgq" >}} | `3.5.1` | {{< bg "18" "postgresql-18-pgq3" "green" >}} {{< bg "17" "postgresql-17-pgq3" "green" >}} {{< bg "16" "postgresql-16-pgq3" "green" >}} {{< bg "15" "postgresql-15-pgq3" "green" >}} {{< bg "14" "postgresql-14-pgq3" "green" >}} | `postgresql-$v-pgq3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 4" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 3" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_18` | 3.5.1 | `el8.x86_64` | pgdg | 54.2 KiB | [pgq_18-3.5.1-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgq_18-3.5.1-4PGDG.rhel8.x86_64.rpm) |
| `pgq_18` | 3.5.1 | `el8.aarch64` | pgdg | 53.9 KiB | [pgq_18-3.5.1-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgq_18-3.5.1-4PGDG.rhel8.aarch64.rpm) |
| `pgq_18` | 3.5.1 | `el9.x86_64` | pgdg | 51.9 KiB | [pgq_18-3.5.1-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgq_18-3.5.1-4PGDG.rhel9.x86_64.rpm) |
| `pgq_18` | 3.5.1 | `el9.aarch64` | pgdg | 51.3 KiB | [pgq_18-3.5.1-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgq_18-3.5.1-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pgq3` | 3.5.1 | `d12.x86_64` | pgdg | 123.9 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgq3` | 3.5.1 | `d12.aarch64` | pgdg | 122.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgq3` | 3.5.1 | `u22.x86_64` | pgdg | 125.7 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgq3` | 3.5.1 | `u22.aarch64` | pgdg | 124.3 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgq3` | 3.5.1 | `u24.x86_64` | pgdg | 123.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgq3` | 3.5.1 | `u24.aarch64` | pgdg | 122.6 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_17` | 3.5.1 | `el8.x86_64` | pgdg | 54.0 KiB | [pgq_17-3.5.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgq_17-3.5.1-3PGDG.rhel8.x86_64.rpm) |
| `pgq_17` | 3.5.1 | `el8.aarch64` | pgdg | 53.8 KiB | [pgq_17-3.5.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgq_17-3.5.1-3PGDG.rhel8.aarch64.rpm) |
| `pgq_17` | 3.5.1 | `el9.x86_64` | pgdg | 51.9 KiB | [pgq_17-3.5.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgq_17-3.5.1-3PGDG.rhel9.x86_64.rpm) |
| `pgq_17` | 3.5.1 | `el9.aarch64` | pgdg | 51.6 KiB | [pgq_17-3.5.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgq_17-3.5.1-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgq3` | 3.5.1 | `d12.x86_64` | pgdg | 123.6 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgq3` | 3.5.1 | `d12.aarch64` | pgdg | 122.6 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgq3` | 3.5.1 | `u22.x86_64` | pgdg | 145.3 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgq3` | 3.5.1 | `u22.aarch64` | pgdg | 143.9 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgq3` | 3.5.1 | `u24.x86_64` | pgdg | 123.5 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgq3` | 3.5.1 | `u24.aarch64` | pgdg | 122.3 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_16` | 3.5.1 | `el8.x86_64` | pgdg | 54.8 KiB | [pgq_16-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgq_16-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_16` | 3.5.1 | `el8.aarch64` | pgdg | 55.1 KiB | [pgq_16-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgq_16-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_16` | 3.5.1 | `el9.x86_64` | pgdg | 52.5 KiB | [pgq_16-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgq_16-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_16` | 3.5.1 | `el9.aarch64` | pgdg | 52.3 KiB | [pgq_16-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgq_16-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pgq3` | 3.5.1 | `d12.x86_64` | pgdg | 123.6 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgq3` | 3.5.1 | `d12.aarch64` | pgdg | 122.6 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgq3` | 3.5.1 | `u22.x86_64` | pgdg | 143.7 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgq3` | 3.5.1 | `u22.aarch64` | pgdg | 142.4 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgq3` | 3.5.1 | `u24.x86_64` | pgdg | 123.5 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgq3` | 3.5.1 | `u24.aarch64` | pgdg | 122.4 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_15` | 3.5.1 | `el8.x86_64` | pgdg | 55.3 KiB | [pgq_15-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgq_15-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_15` | 3.5 | `el8.x86_64` | pgdg | 54.8 KiB | [pgq_15-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgq_15-3.5-1.rhel8.x86_64.rpm) |
| `pgq_15` | 3.5.1 | `el8.aarch64` | pgdg | 55.5 KiB | [pgq_15-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgq_15-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_15` | 3.5 | `el8.aarch64` | pgdg | 55.0 KiB | [pgq_15-3.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgq_15-3.5-1.rhel8.aarch64.rpm) |
| `pgq_15` | 3.5.1 | `el9.x86_64` | pgdg | 53.6 KiB | [pgq_15-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgq_15-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_15` | 3.5 | `el9.x86_64` | pgdg | 53.3 KiB | [pgq_15-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgq_15-3.5-1.rhel9.x86_64.rpm) |
| `pgq_15` | 3.5.1 | `el9.aarch64` | pgdg | 53.3 KiB | [pgq_15-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgq_15-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgq_15` | 3.5 | `el9.aarch64` | pgdg | 53.1 KiB | [pgq_15-3.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgq_15-3.5-1.rhel9.aarch64.rpm) |
| `postgresql-15-pgq3` | 3.5.1 | `d12.x86_64` | pgdg | 124.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgq3` | 3.5.1 | `d12.aarch64` | pgdg | 123.0 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgq3` | 3.5.1 | `u22.x86_64` | pgdg | 144.2 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgq3` | 3.5.1 | `u22.aarch64` | pgdg | 142.9 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgq3` | 3.5.1 | `u24.x86_64` | pgdg | 124.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgq3` | 3.5.1 | `u24.aarch64` | pgdg | 123.0 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_14` | 3.5.1 | `el8.x86_64` | pgdg | 55.3 KiB | [pgq_14-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_14` | 3.5 | `el8.x86_64` | pgdg | 54.8 KiB | [pgq_14-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.5-1.rhel8.x86_64.rpm) |
| `pgq_14` | 3.4.2 | `el8.x86_64` | pgdg | 54.4 KiB | [pgq_14-3.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.4.2-1.rhel8.x86_64.rpm) |
| `pgq_14` | 3.4.1 | `el8.x86_64` | pgdg | 108.6 KiB | [pgq_14-3.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.4.1-2.rhel8.x86_64.rpm) |
| `pgq_14` | 3.5.1 | `el8.aarch64` | pgdg | 55.5 KiB | [pgq_14-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgq_14-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_14` | 3.5 | `el8.aarch64` | pgdg | 55.0 KiB | [pgq_14-3.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgq_14-3.5-1.rhel8.aarch64.rpm) |
| `pgq_14` | 3.5.1 | `el9.x86_64` | pgdg | 53.5 KiB | [pgq_14-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_14` | 3.5 | `el9.x86_64` | pgdg | 53.3 KiB | [pgq_14-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.5-1.rhel9.x86_64.rpm) |
| `pgq_14` | 3.4.2 | `el9.x86_64` | pgdg | 107.8 KiB | [pgq_14-3.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.4.2-1.rhel9.x86_64.rpm) |
| `pgq_14` | 3.5.1 | `el9.aarch64` | pgdg | 53.3 KiB | [pgq_14-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgq_14-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgq_14` | 3.5 | `el9.aarch64` | pgdg | 53.1 KiB | [pgq_14-3.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgq_14-3.5-1.rhel9.aarch64.rpm) |
| `postgresql-14-pgq3` | 3.5.1 | `d12.x86_64` | pgdg | 124.1 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgq3` | 3.5.1 | `d12.aarch64` | pgdg | 123.0 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgq3` | 3.5.1 | `u22.x86_64` | pgdg | 134.8 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgq3` | 3.5.1 | `u22.aarch64` | pgdg | 133.5 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgq3` | 3.5.1 | `u24.x86_64` | pgdg | 124.0 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgq3` | 3.5.1 | `u24.aarch64` | pgdg | 122.9 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgq/pgq" title="Repository" icon="github" subtitle="github.com/pgq/pgq" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgq; # install by extension name, for the current active PG version
pig ext install pgq; # install via package alias, for the active PG version
pig ext install pgq -v 18;   # install for PG 18
pig ext install pgq -v 17;   # install for PG 17
pig ext install pgq -v 16;   # install for PG 16
pig ext install pgq -v 15;   # install for PG 15
pig ext install pgq -v 14;   # install for PG 14
pig ext install pgq -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgq CASCADE SCHEMA pg_catalog;
```

