---
title: "pgq"
linkTitle: "pgq"
description: "Generic queue for PostgreSQL"
weight: 2650
categories: ["FEAT"]
width: full
---

[**pgq**](https://github.com/pgq/pgq) : Generic queue for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2650** | {{< badge content="pgq" link="https://github.com/pgq/pgq" >}} | {{< ext "pgq" >}} | `3.5.1` | {{< category "FEAT" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgq` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.5.1` | {{< bg "18" "pgq_18" "green" >}} {{< bg "17" "pgq_17" "green" >}} {{< bg "16" "pgq_16" "green" >}} {{< bg "15" "pgq_15" "green" >}} {{< bg "14" "pgq_14" "green" >}} | `pgq_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.5.1` | {{< bg "18" "postgresql-18-pgq3" "green" >}} {{< bg "17" "postgresql-17-pgq3" "green" >}} {{< bg "16" "postgresql-16-pgq3" "green" >}} {{< bg "15" "postgresql-15-pgq3" "green" >}} {{< bg "14" "postgresql-14-pgq3" "green" >}} | `postgresql-$v-pgq3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.5.1" "pgq_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "pgq_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 3.5.1" "postgresql-18-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-17-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-16-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-15-pgq3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.5.1" "postgresql-14-pgq3 : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_18` | `3.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.2 KiB | [pgq_18-3.5.1-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgq_18-3.5.1-4PGDG.rhel8.x86_64.rpm) |
| `pgq_18` | `3.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pgq_18-3.5.1-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgq_18-3.5.1-4PGDG.rhel8.aarch64.rpm) |
| `pgq_18` | `3.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.9 KiB | [pgq_18-3.5.1-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgq_18-3.5.1-4PGDG.rhel9.x86_64.rpm) |
| `pgq_18` | `3.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.3 KiB | [pgq_18-3.5.1-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgq_18-3.5.1-4PGDG.rhel9.aarch64.rpm) |
| `pgq_18` | `3.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.5 KiB | [pgq_18-3.5.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgq_18-3.5.1-4PGDG.rhel10.x86_64.rpm) |
| `pgq_18` | `3.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.2 KiB | [pgq_18-3.5.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgq_18-3.5.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgq3` | `3.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.9 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 122.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 122.9 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 125.7 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 124.3 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 122.6 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.8 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgq3` | `3.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-18-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-18-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_17` | `3.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.0 KiB | [pgq_17-3.5.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgq_17-3.5.1-3PGDG.rhel8.x86_64.rpm) |
| `pgq_17` | `3.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pgq_17-3.5.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgq_17-3.5.1-3PGDG.rhel8.aarch64.rpm) |
| `pgq_17` | `3.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.9 KiB | [pgq_17-3.5.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgq_17-3.5.1-3PGDG.rhel9.x86_64.rpm) |
| `pgq_17` | `3.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.6 KiB | [pgq_17-3.5.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgq_17-3.5.1-3PGDG.rhel9.aarch64.rpm) |
| `pgq_17` | `3.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.4 KiB | [pgq_17-3.5.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgq_17-3.5.1-4PGDG.rhel10.x86_64.rpm) |
| `pgq_17` | `3.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [pgq_17-3.5.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgq_17-3.5.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgq3` | `3.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.6 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 122.6 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.5 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 122.7 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 145.3 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.9 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.5 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 122.3 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.6 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgq3` | `3.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.0 KiB | [postgresql-17-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-17-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_16` | `3.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.8 KiB | [pgq_16-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgq_16-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_16` | `3.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.1 KiB | [pgq_16-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgq_16-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_16` | `3.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.5 KiB | [pgq_16-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgq_16-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_16` | `3.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.3 KiB | [pgq_16-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgq_16-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgq_16` | `3.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.4 KiB | [pgq_16-3.5.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgq_16-3.5.1-4PGDG.rhel10.x86_64.rpm) |
| `pgq_16` | `3.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [pgq_16-3.5.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgq_16-3.5.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgq3` | `3.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.6 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 122.6 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.5 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 122.7 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 143.7 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.4 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.5 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 122.4 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.7 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgq3` | `3.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.0 KiB | [postgresql-16-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-16-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_15` | `3.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pgq_15-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgq_15-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_15` | `3.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.8 KiB | [pgq_15-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgq_15-3.5-1.rhel8.x86_64.rpm) |
| `pgq_15` | `3.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.5 KiB | [pgq_15-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgq_15-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_15` | `3.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.0 KiB | [pgq_15-3.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgq_15-3.5-1.rhel8.aarch64.rpm) |
| `pgq_15` | `3.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.6 KiB | [pgq_15-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgq_15-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_15` | `3.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.3 KiB | [pgq_15-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgq_15-3.5-1.rhel9.x86_64.rpm) |
| `pgq_15` | `3.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pgq_15-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgq_15-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgq_15` | `3.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgq_15-3.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgq_15-3.5-1.rhel9.aarch64.rpm) |
| `pgq_15` | `3.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.2 KiB | [pgq_15-3.5.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgq_15-3.5.1-4PGDG.rhel10.x86_64.rpm) |
| `pgq_15` | `3.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.0 KiB | [pgq_15-3.5.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgq_15-3.5.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgq3` | `3.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 123.0 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 123.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 144.2 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.9 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.1 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.0 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.4 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgq3` | `3.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.5 KiB | [postgresql-15-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-15-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgq_14` | `3.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pgq_14-3.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgq_14` | `3.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.8 KiB | [pgq_14-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.5-1.rhel8.x86_64.rpm) |
| `pgq_14` | `3.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.4 KiB | [pgq_14-3.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.4.2-1.rhel8.x86_64.rpm) |
| `pgq_14` | `3.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.6 KiB | [pgq_14-3.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgq_14-3.4.1-2.rhel8.x86_64.rpm) |
| `pgq_14` | `3.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.5 KiB | [pgq_14-3.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgq_14-3.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgq_14` | `3.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.0 KiB | [pgq_14-3.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgq_14-3.5-1.rhel8.aarch64.rpm) |
| `pgq_14` | `3.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.5 KiB | [pgq_14-3.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgq_14` | `3.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.3 KiB | [pgq_14-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.5-1.rhel9.x86_64.rpm) |
| `pgq_14` | `3.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pgq_14-3.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgq_14-3.4.2-1.rhel9.x86_64.rpm) |
| `pgq_14` | `3.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [pgq_14-3.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgq_14-3.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgq_14` | `3.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.1 KiB | [pgq_14-3.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgq_14-3.5-1.rhel9.aarch64.rpm) |
| `pgq_14` | `3.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.2 KiB | [pgq_14-3.5.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgq_14-3.5.1-4PGDG.rhel10.x86_64.rpm) |
| `pgq_14` | `3.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.9 KiB | [pgq_14-3.5.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgq_14-3.5.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgq3` | `3.5.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.1 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 123.0 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.0 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 123.1 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 134.8 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.5 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.0 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 122.9 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.3 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgq3` | `3.5.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.4 KiB | [postgresql-14-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgq/postgresql-14-pgq3_3.5.1-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgq/pgq" title="Repository" icon="github" subtitle="github.com/pgq/pgq" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgq;		# install via package name, for the active PG version

pig install pgq -v 18;   # install for PG 18
pig install pgq -v 17;   # install for PG 17
pig install pgq -v 16;   # install for PG 16
pig install pgq -v 15;   # install for PG 15
pig install pgq -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgq;
```




## Usage

> [pgq: Generic high-performance lockless queue for PostgreSQL](https://github.com/pgq/pgq)

PgQ is a PostgreSQL extension that provides a generic, high-performance lockless queue with a simple SQL function API. It uses a producer-consumer model with batch-based event processing.

```sql
CREATE EXTENSION pgq;
```

### Core Concepts

- **Queue**: A named event stream. Events are inserted by producers and consumed in batches.
- **Consumer**: A named subscriber registered to a queue. Each consumer tracks its own position.
- **Batch**: A group of events retrieved together. Consumers process events batch by batch.
- **Ticker**: A background process that creates batch boundaries (ticks) at regular intervals.

### Queue Management

```sql
-- Create a queue
SELECT pgq.create_queue('myqueue');

-- Drop a queue
SELECT pgq.drop_queue('myqueue');

-- Get queue info
SELECT * FROM pgq.get_queue_info();
SELECT * FROM pgq.get_queue_info('myqueue');
```

### Consumer Registration

```sql
-- Register a consumer on a queue
SELECT pgq.register_consumer('myqueue', 'myconsumer');

-- Unregister a consumer
SELECT pgq.unregister_consumer('myqueue', 'myconsumer');

-- Get consumer info
SELECT * FROM pgq.get_consumer_info('myqueue');
```

### Producing Events

```sql
-- Insert an event into a queue
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data');

-- Insert with extra fields
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data',
                         'extra1', 'extra2', 'extra3', 'extra4');
```

### Consuming Events

```sql
-- Get the next batch of events (returns batch_id or NULL if no new batches)
SELECT pgq.next_batch('myqueue', 'myconsumer');

-- Get events from the batch
SELECT * FROM pgq.get_batch_events(:batch_id);

-- Retry a failed event (will reappear after the specified interval)
SELECT pgq.event_retry(:batch_id, :event_id, :retry_seconds);

-- Mark batch as done
SELECT pgq.finish_batch(:batch_id);
```

### Typical Consumer Loop

```sql
-- 1. Get next batch
SELECT pgq.next_batch('myqueue', 'myconsumer') AS batch_id;

-- 2. If batch_id is not NULL, get events
SELECT * FROM pgq.get_batch_events(:batch_id);

-- 3. Process events, retry failures
SELECT pgq.event_retry(:batch_id, :event_id, 60);

-- 4. Finish the batch
SELECT pgq.finish_batch(:batch_id);
```

### Maintenance

PgQ requires a ticker daemon (`pgqd`) to run in the background for creating batch boundaries and performing maintenance tasks like table rotation and retry event processing.

### Key Functions

| Function | Description |
|----------|-------------|
| `pgq.create_queue(name)` | Create a new queue |
| `pgq.drop_queue(name)` | Remove a queue |
| `pgq.register_consumer(queue, consumer)` | Register a consumer |
| `pgq.unregister_consumer(queue, consumer)` | Unregister a consumer |
| `pgq.insert_event(queue, type, data, ...)` | Insert an event |
| `pgq.next_batch(queue, consumer)` | Get next batch ID |
| `pgq.get_batch_events(batch_id)` | Get events from a batch |
| `pgq.event_retry(batch_id, event_id, seconds)` | Schedule event retry |
| `pgq.finish_batch(batch_id)` | Mark batch as processed |
| `pgq.get_queue_info([name])` | Get queue statistics |
| `pgq.get_consumer_info(queue)` | Get consumer statistics |
