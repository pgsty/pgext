---
title: "rum"
linkTitle: "rum"
description: "RUM index access method"
weight: 2780
categories: ["FEAT"]
width: full
---

[**rum**](https://github.com/postgrespro/rum) : RUM index access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2780** | {{< badge content="rum" link="https://github.com/postgrespro/rum" >}} | {{< ext "rum" >}} | `1.3.14` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "btree_gist" >}} {{< ext "btree_gin" >}} {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pgroonga_database" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.14` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `rum` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.14` | {{< bg "18" "rum_18" "red" >}} {{< bg "17" "rum_17" "green" >}} {{< bg "16" "rum_16" "green" >}} {{< bg "15" "rum_15" "green" >}} {{< bg "14" "rum_14" "green" >}} {{< bg "13" "rum_13" "green" >}} | `rum_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.14` | {{< bg "18" "postgresql-18-rum" "red" >}} {{< bg "17" "postgresql-17-rum" "green" >}} {{< bg "16" "postgresql-16-rum" "green" >}} {{< bg "15" "postgresql-15-rum" "green" >}} {{< bg "14" "postgresql-14-rum" "green" >}} {{< bg "13" "postgresql-13-rum" "green" >}} | `postgresql-$v-rum` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "rum_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.14" "rum_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.14" "rum_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-14-rum : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-13-rum : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 233.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.4 KiB | [postgresql-18-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 233.5 KiB | [postgresql-18-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.2 KiB | [postgresql-18-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 240.5 KiB | [postgresql-18-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 232.1 KiB | [postgresql-18-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 234.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_17` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 93.1 KiB | [rum_17-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/rum_17-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_17` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.8 KiB | [rum_17-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/rum_17-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_17` | `1.3.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 91.6 KiB | [rum_17-1.3.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/rum_17-1.3.14-1PGDG.rhel9.x86_64.rpm) |
| `rum_17` | `1.3.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.8 KiB | [rum_17-1.3.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/rum_17-1.3.14-1PGDG.rhel9.aarch64.rpm) |
| `rum_17` | `1.3.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 93.0 KiB | [rum_17-1.3.14-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/rum_17-1.3.14-2PGDG.rhel10.x86_64.rpm) |
| `rum_17` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.8 KiB | [rum_17-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/rum_17-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.4 KiB | [postgresql-17-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.9 KiB | [postgresql-17-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.6 KiB | [postgresql-17-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.7 KiB | [postgresql-17-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 265.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 257.4 KiB | [postgresql-17-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 235.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_16` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 93.0 KiB | [rum_16-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_16` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 92.7 KiB | [rum_16-1.3.13-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.13-2.rhel8.1.x86_64.rpm) |
| `rum_16` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.8 KiB | [rum_16-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_16` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.4 KiB | [rum_16-1.3.13-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.13-2.rhel8.1.aarch64.rpm) |
| `rum_16` | `1.3.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 91.7 KiB | [rum_16-1.3.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/rum_16-1.3.14-1PGDG.rhel9.x86_64.rpm) |
| `rum_16` | `1.3.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 91.3 KiB | [rum_16-1.3.13-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/rum_16-1.3.13-2.rhel9.1.x86_64.rpm) |
| `rum_16` | `1.3.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.7 KiB | [rum_16-1.3.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/rum_16-1.3.14-1PGDG.rhel9.aarch64.rpm) |
| `rum_16` | `1.3.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.3 KiB | [rum_16-1.3.13-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/rum_16-1.3.13-2.rhel9.1.aarch64.rpm) |
| `rum_16` | `1.3.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.9 KiB | [rum_16-1.3.14-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/rum_16-1.3.14-2PGDG.rhel10.x86_64.rpm) |
| `rum_16` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.7 KiB | [rum_16-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/rum_16-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.4 KiB | [postgresql-16-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.5 KiB | [postgresql-16-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.6 KiB | [postgresql-16-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 264.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 256.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 234.9 KiB | [postgresql-16-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.9 KiB | [postgresql-16-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_15` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.5 KiB | [rum_15-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_15` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.1 KiB | [rum_15-1.3.13-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.13-1.rhel8.x86_64.rpm) |
| `rum_15` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.7 KiB | [rum_15-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_15` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.0 KiB | [rum_15-1.3.13-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.13-1.rhel8.aarch64.rpm) |
| `rum_15` | `1.3.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.8 KiB | [rum_15-1.3.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/rum_15-1.3.14-1PGDG.rhel9.x86_64.rpm) |
| `rum_15` | `1.3.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.7 KiB | [rum_15-1.3.13-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/rum_15-1.3.13-1.rhel9.x86_64.rpm) |
| `rum_15` | `1.3.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.4 KiB | [rum_15-1.3.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/rum_15-1.3.14-1PGDG.rhel9.aarch64.rpm) |
| `rum_15` | `1.3.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.5 KiB | [rum_15-1.3.13-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/rum_15-1.3.13-1.rhel9.aarch64.rpm) |
| `rum_15` | `1.3.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 112.7 KiB | [rum_15-1.3.14-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/rum_15-1.3.14-2PGDG.rhel10.x86_64.rpm) |
| `rum_15` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.3 KiB | [rum_15-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/rum_15-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 288.7 KiB | [postgresql-15-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 277.9 KiB | [postgresql-15-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 289.6 KiB | [postgresql-15-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 278.7 KiB | [postgresql-15-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 327.0 KiB | [postgresql-15-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 316.2 KiB | [postgresql-15-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 288.9 KiB | [postgresql-15-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 279.2 KiB | [postgresql-15-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_14` | `1.3.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 308.8 KiB | [rum_14-1.3.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/rum_14-1.3.8-1.rhel8.x86_64.rpm) |
| `rum_14` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.7 KiB | [rum_14-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/rum_14-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_14` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.1 KiB | [rum_14-1.3.13-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/rum_14-1.3.13-1.rhel8.x86_64.rpm) |
| `rum_14` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.3 KiB | [rum_14-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/rum_14-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_14` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.8 KiB | [rum_14-1.3.13-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/rum_14-1.3.13-1.rhel8.aarch64.rpm) |
| `rum_14` | `1.3.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.1 KiB | [rum_14-1.3.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/rum_14-1.3.14-1PGDG.rhel9.x86_64.rpm) |
| `rum_14` | `1.3.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.1 KiB | [rum_14-1.3.13-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/rum_14-1.3.13-1.rhel9.x86_64.rpm) |
| `rum_14` | `1.3.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.8 KiB | [rum_14-1.3.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/rum_14-1.3.14-1PGDG.rhel9.aarch64.rpm) |
| `rum_14` | `1.3.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.8 KiB | [rum_14-1.3.13-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/rum_14-1.3.13-1.rhel9.aarch64.rpm) |
| `rum_14` | `1.3.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.8 KiB | [rum_14-1.3.14-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/rum_14-1.3.14-2PGDG.rhel10.x86_64.rpm) |
| `rum_14` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.1 KiB | [rum_14-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/rum_14-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 287.6 KiB | [postgresql-14-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-14-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 277.0 KiB | [postgresql-14-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-14-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 287.9 KiB | [postgresql-14-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-14-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 277.5 KiB | [postgresql-14-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-14-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 324.6 KiB | [postgresql-14-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 314.5 KiB | [postgresql-14-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 287.2 KiB | [postgresql-14-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 277.8 KiB | [postgresql-14-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-14-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_13` | `1.3.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 306.8 KiB | [rum_13-1.3.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/rum_13-1.3.8-1.rhel8.x86_64.rpm) |
| `rum_13` | `1.3.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 305.0 KiB | [rum_13-1.3.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/rum_13-1.3.7-1.rhel8.x86_64.rpm) |
| `rum_13` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.5 KiB | [rum_13-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/rum_13-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_13` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.9 KiB | [rum_13-1.3.13-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/rum_13-1.3.13-1.rhel8.x86_64.rpm) |
| `rum_13` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.0 KiB | [rum_13-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/rum_13-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_13` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.4 KiB | [rum_13-1.3.13-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/rum_13-1.3.13-1.rhel8.aarch64.rpm) |
| `rum_13` | `1.3.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 110.6 KiB | [rum_13-1.3.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/rum_13-1.3.14-1PGDG.rhel9.x86_64.rpm) |
| `rum_13` | `1.3.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 110.6 KiB | [rum_13-1.3.13-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/rum_13-1.3.13-1.rhel9.x86_64.rpm) |
| `rum_13` | `1.3.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.7 KiB | [rum_13-1.3.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/rum_13-1.3.14-1PGDG.rhel9.aarch64.rpm) |
| `rum_13` | `1.3.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.7 KiB | [rum_13-1.3.13-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/rum_13-1.3.13-1.rhel9.aarch64.rpm) |
| `rum_13` | `1.3.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 111.5 KiB | [rum_13-1.3.14-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/rum_13-1.3.14-2PGDG.rhel10.x86_64.rpm) |
| `rum_13` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.8 KiB | [rum_13-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/rum_13-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 286.5 KiB | [postgresql-13-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-13-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 276.3 KiB | [postgresql-13-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-13-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 287.0 KiB | [postgresql-13-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-13-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 276.5 KiB | [postgresql-13-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-13-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 322.6 KiB | [postgresql-13-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 312.3 KiB | [postgresql-13-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 286.5 KiB | [postgresql-13-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 277.4 KiB | [postgresql-13-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-13-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/rum" title="Repository" icon="github" subtitle="github.com/postgrespro/rum" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install rum;		# install via package name, for the active PG version

pig install rum -v 17;   # install for PG 17
pig install rum -v 16;   # install for PG 16
pig install rum -v 15;   # install for PG 15
pig install rum -v 14;   # install for PG 14
pig install rum -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION rum;
```
