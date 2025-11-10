---
title: "tdigest"
linkTitle: "tdigest"
description: "Provides tdigest aggregate function."
weight: 4700
categories: ["FUNC"]
width: full
---

[**tdigest**](https://github.com/tvondra/tdigest)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4700** | {{< badge content="tdigest" link="https://github.com/tvondra/tdigest" >}} | {{< ext "tdigest" >}} | `1.4.3` | {{< category "FUNC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/tdigest" >}} | `1.4.2` | {{< bg "18" "tdigest_18*" "green" >}} {{< bg "17" "tdigest_17*" "green" >}} {{< bg "16" "tdigest_16*" "green" >}} {{< bg "15" "tdigest_15*" "green" >}} {{< bg "14" "tdigest_14*" "green" >}} {{< bg "13" "tdigest_13*" "green" >}} | `tdigest_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/tdigest" >}} | `1.4.3` | {{< bg "18" "postgresql-18-tdigest" "green" >}} {{< bg "17" "postgresql-17-tdigest" "green" >}} {{< bg "16" "postgresql-16-tdigest" "green" >}} {{< bg "15" "postgresql-15-tdigest" "green" >}} {{< bg "14" "postgresql-14-tdigest" "green" >}} {{< bg "13" "postgresql-13-tdigest" "green" >}} | `postgresql-$v-tdigest` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "tdigest_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.2" "tdigest_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "tdigest_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4.3" "postgresql-18-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-17-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-16-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-15-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-14-tdigest : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.3" "postgresql-13-tdigest : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_18` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [tdigest_18-1.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.2-2PGDG.rhel8.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.1 KiB | [tdigest_18-1.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.2-2PGDG.rhel8.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_18-1.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.2-2PGDG.rhel9.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [tdigest_18-1.4.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.2-2PGDG.rhel9.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_18-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_18-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.6 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.9 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 58.0 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 57.0 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.7 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.9 KiB | [postgresql-18-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_17` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [tdigest_17-1.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.2-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.2 KiB | [tdigest_17-1.4.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.1-3PGDG.rhel8.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.1 KiB | [tdigest_17-1.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.2-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [tdigest_17-1.4.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.1-3PGDG.rhel8.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_17-1.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.2-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.3 KiB | [tdigest_17-1.4.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.1-3PGDG.rhel9.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_17-1.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.2-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.1 KiB | [tdigest_17-1.4.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.1-3PGDG.rhel9.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_17-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_17-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.6 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.0 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.2 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.3 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.6 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.9 KiB | [postgresql-17-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_16-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_16-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_16-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_16-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_16-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_16-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.5 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.0 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.2 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.3 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.6 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.9 KiB | [postgresql-16-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_15-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.1 KiB | [tdigest_15-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.0-1.rhel8.x86_64.rpm) |
| `tdigest_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_15-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.6 KiB | [tdigest_15-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.0-1.rhel8.aarch64.rpm) |
| `tdigest_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_15-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.1 KiB | [tdigest_15-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.0-1.rhel9.x86_64.rpm) |
| `tdigest_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_15-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [tdigest_15-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.0-1.rhel9.aarch64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_15-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_15-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.6 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 57.0 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.5 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.5 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.6 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.9 KiB | [postgresql-15-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_14-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.3 KiB | [tdigest_14-1.2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.2.0-2.rhel8.x86_64.rpm) |
| `tdigest_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_14-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.5 KiB | [tdigest_14-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.0-1.rhel8.aarch64.rpm) |
| `tdigest_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_14-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.2 KiB | [tdigest_14-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.0-1.rhel9.x86_64.rpm) |
| `tdigest_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_14-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [tdigest_14-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.0-1.rhel9.aarch64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_14-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_14-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.5 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.7 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.9 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.4 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.4 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.5 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.9 KiB | [postgresql-14-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_13` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [tdigest_13-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tdigest_13-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.9 KiB | [tdigest_13-1.2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tdigest_13-1.2.0-2.rhel8.x86_64.rpm) |
| `tdigest_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.5 KiB | [tdigest_13-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tdigest_13-1.0.1-1.rhel8.x86_64.rpm) |
| `tdigest_13` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_13-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/tdigest_13-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_13` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.4 KiB | [tdigest_13-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/tdigest_13-1.4.0-1.rhel8.aarch64.rpm) |
| `tdigest_13` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_13-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/tdigest_13-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_13` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.7 KiB | [tdigest_13-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/tdigest_13-1.4.0-1.rhel9.x86_64.rpm) |
| `tdigest_13` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [tdigest_13-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/tdigest_13-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_13` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.5 KiB | [tdigest_13-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/tdigest_13-1.4.0-1.rhel9.aarch64.rpm) |
| `tdigest_13` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_13-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/tdigest_13-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_13` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_13-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/tdigest_13-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-tdigest` | `1.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.3 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg12+1_amd64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 56.9 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg12+1_arm64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.6 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg13+1_amd64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 56.9 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg13+1_arm64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 60.5 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.4 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.8 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-tdigest` | `1.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 56.8 KiB | [postgresql-13-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-13-tdigest_1.4.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/tdigest" title="Repository" icon="github" subtitle="github.com/tvondra/tdigest" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install tdigest; # install by extension name, for the current active PG version
pig ext install tdigest; # install via package alias, for the active PG version
pig ext install tdigest -v 18;   # install for PG 18
pig ext install tdigest -v 17;   # install for PG 17
pig ext install tdigest -v 16;   # install for PG 16
pig ext install tdigest -v 15;   # install for PG 15
pig ext install tdigest -v 14;   # install for PG 14
pig ext install tdigest -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION tdigest;
```

