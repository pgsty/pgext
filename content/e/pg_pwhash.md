---
title: "pg_pwhash"
linkTitle: "pg_pwhash"
description: "Advanced password hashing methods for PostgreSQL"
weight: 7330
categories: ["SEC"]
width: full
---

[**pg_pwhash**](https://github.com/cybertec-postgresql/pg_pwhash) : Advanced password hashing methods for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7330** | {{< badge content="pg_pwhash" link="https://github.com/cybertec-postgresql/pg_pwhash" >}} | {{< ext "pg_pwhash" >}} | `1.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] RPM metadata shows license=PostgreSQL, but packaged LICENSE file is MIT


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_pwhash` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "pg_pwhash_18" "green" >}} {{< bg "17" "pg_pwhash_17" "green" >}} {{< bg "16" "pg_pwhash_16" "green" >}} {{< bg "15" "pg_pwhash_15" "green" >}} {{< bg "14" "pg_pwhash_14" "green" >}} {{< bg "13" "pg_pwhash_13" "red" >}} | `pg_pwhash_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-pwhash" "green" >}} {{< bg "17" "postgresql-17-pg-pwhash" "green" >}} {{< bg "16" "postgresql-16-pg-pwhash" "green" >}} {{< bg "15" "postgresql-15-pg-pwhash" "green" >}} {{< bg "14" "postgresql-14-pg-pwhash" "green" >}} {{< bg "13" "postgresql-13-pg-pwhash" "green" >}} | `postgresql-$v-pg-pwhash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "pg_pwhash_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_pwhash_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_pwhash_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-pg-pwhash : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-pg-pwhash : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pwhash_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.2 KiB | [pg_pwhash_18-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_pwhash_18-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_pwhash_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_pwhash_18-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_pwhash_18-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_pwhash_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_18-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_pwhash_18-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_pwhash_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_pwhash_18-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_pwhash_18-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_pwhash_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [pg_pwhash_18-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_pwhash_18-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_pwhash_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_pwhash_18-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_pwhash_18-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 55.9 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 55.0 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.6 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.5 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 56.4 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 54.9 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 55.8 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 54.6 KiB | [postgresql-18-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-18-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pwhash_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.2 KiB | [pg_pwhash_17-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_pwhash_17-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_pwhash_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_pwhash_17-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_pwhash_17-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_pwhash_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.9 KiB | [pg_pwhash_17-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_pwhash_17-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_pwhash_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_pwhash_17-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_pwhash_17-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_pwhash_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_17-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_pwhash_17-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_pwhash_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_pwhash_17-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_pwhash_17-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 55.9 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 54.9 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.5 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.5 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.0 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 57.6 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 55.8 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 54.6 KiB | [postgresql-17-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-17-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pwhash_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.2 KiB | [pg_pwhash_16-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_pwhash_16-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_pwhash_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_pwhash_16-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_pwhash_16-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_pwhash_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.9 KiB | [pg_pwhash_16-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_pwhash_16-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_pwhash_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_pwhash_16-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_pwhash_16-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_pwhash_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_16-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_pwhash_16-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_pwhash_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_pwhash_16-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_pwhash_16-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-16-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 55.9 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 54.9 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.5 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.5 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.0 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 57.6 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 55.8 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 54.6 KiB | [postgresql-16-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-16-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pwhash_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.2 KiB | [pg_pwhash_15-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_pwhash_15-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_pwhash_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_pwhash_15-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_pwhash_15-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_pwhash_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_15-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_pwhash_15-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_pwhash_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_pwhash_15-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_pwhash_15-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_pwhash_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_15-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_pwhash_15-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_pwhash_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_pwhash_15-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_pwhash_15-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-15-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.3 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 55.1 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.8 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.7 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.3 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 58.0 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.1 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 55.1 KiB | [postgresql-15-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-15-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pwhash_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.2 KiB | [pg_pwhash_14-1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_pwhash_14-1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_pwhash_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_pwhash_14-1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_pwhash_14-1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_pwhash_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_14-1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_pwhash_14-1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_pwhash_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_pwhash_14-1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_pwhash_14-1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_pwhash_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [pg_pwhash_14-1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_pwhash_14-1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_pwhash_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_pwhash_14-1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_pwhash_14-1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-14-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.2 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 55.1 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.8 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.7 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.3 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 57.8 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.1 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 55.0 KiB | [postgresql-14-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-14-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-pg-pwhash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.0 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 55.1 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.7 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 55.6 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.1 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 57.7 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 55.9 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-pwhash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 54.9 KiB | [postgresql-13-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-pwhash/postgresql-13-pg-pwhash_1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_pwhash" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_pwhash" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_pwhash;		# install via package name, for the active PG version

pig install pg_pwhash -v 18;   # install for PG 18
pig install pg_pwhash -v 17;   # install for PG 17
pig install pg_pwhash -v 16;   # install for PG 16
pig install pg_pwhash -v 15;   # install for PG 15
pig install pg_pwhash -v 14;   # install for PG 14
pig install pg_pwhash -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_pwhash;
```
