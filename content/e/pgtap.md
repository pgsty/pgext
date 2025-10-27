---
title: "pgtap"
linkTitle: "pgtap"
description: "Unit testing for PostgreSQL"
weight: 3200
categories: ["LANG"]
width: full
---

Unit testing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3200** | {{< badge content="pgtap" link="https://github.com/theory/pgtap" >}} | {{< ext "pgtap" >}} | `1.3.3` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql_check" >}} {{< ext "plpgsql" >}} {{< ext "pldbgapi" >}} {{< ext "plprofiler" >}} {{< ext "faker" >}} {{< ext "unit" >}} {{< ext "dbt2" >}} {{< ext "plperl" >}} |

> [!Note] missing pg17 el9, breaking perl deps


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgtap" >}} | `1.3.3` | {{< bg "18" "pgtap_18*" "green" >}} {{< bg "17" "pgtap_17*" "green" >}} {{< bg "16" "pgtap_16*" "green" >}} {{< bg "15" "pgtap_15*" "green" >}} {{< bg "14" "pgtap_14*" "green" >}} | `pgtap_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgtap" >}} | `1.3.3` | {{< bg "18" "postgresql-18-pgtap" "green" >}} {{< bg "17" "postgresql-17-pgtap" "green" >}} {{< bg "16" "postgresql-16-pgtap" "green" >}} {{< bg "15" "postgresql-15-pgtap" "green" >}} {{< bg "14" "postgresql-14-pgtap" "green" >}} | `postgresql-$v-pgtap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PGDG 1.3.4" "pgtap_18 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_17 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_16 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_15 : HIDE 3" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_14 : HIDE 4" >}}   |
|    `el8.aarch64`    |  {{< bg "PGDG 1.3.4" "pgtap_18 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_17 : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_16 : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_15 : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_14 : HIDE 1" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 1.3.4" "pgtap_18 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_17 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_16 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_15 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_14 : HIDE 2" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 1.3.4" "pgtap_18 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_17 : HIDE 2" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_16 : HIDE 3" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_15 : HIDE 3" >}}   |  {{< bg "PGDG 1.3.4" "pgtap_14 : HIDE 3" >}}   |
|    `d12.x86_64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : HIDE 1" >}}   |  {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : HIDE 1" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_18` | 1.3.4 | `el8.x86_64` | pgdg | 118.1 KiB | [pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | 1.3.3 | `el8.x86_64` | pgdg | 117.3 KiB | [pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | 1.3.4 | `el8.aarch64` | pgdg | 118.1 KiB | [pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | 1.3.3 | `el8.aarch64` | pgdg | 117.3 KiB | [pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | 1.3.4 | `el9.x86_64` | pgdg | 106.8 KiB | [pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | 1.3.3 | `el9.x86_64` | pgdg | 106.5 KiB | [pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | 1.3.4 | `el9.aarch64` | pgdg | 106.7 KiB | [pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | 1.3.3 | `el9.aarch64` | pgdg | 106.4 KiB | [pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `postgresql-18-pgtap` | 1.3.4 | `d12.x86_64` | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-18-pgtap` | 1.3.4 | `d12.aarch64` | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-18-pgtap` | 1.3.4 | `u22.x86_64` | pgdg | 46.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pgtap` | 1.3.4 | `u22.aarch64` | pgdg | 46.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pgtap` | 1.3.4 | `u24.x86_64` | pgdg | 44.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-18-pgtap` | 1.3.4 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_17` | 1.3.4 | `el8.x86_64` | pgdg | 118.1 KiB | [pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | 1.3.3 | `el8.x86_64` | pgdg | 117.3 KiB | [pgtap_17-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtap_17-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | 1.3.4 | `el8.aarch64` | pgdg | 118.1 KiB | [pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | 1.3.4 | `el9.x86_64` | pgdg | 106.8 KiB | [pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | 1.3.3 | `el9.x86_64` | pgdg | 106.5 KiB | [pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | 1.3.4 | `el9.aarch64` | pgdg | 106.7 KiB | [pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | 1.3.3 | `el9.aarch64` | pgdg | 106.5 KiB | [pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pgtap` | 1.3.4 | `d12.x86_64` | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-17-pgtap` | 1.3.4 | `d12.aarch64` | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-17-pgtap` | 1.3.4 | `u22.x86_64` | pgdg | 46.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pgtap` | 1.3.4 | `u22.aarch64` | pgdg | 46.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pgtap` | 1.3.4 | `u24.x86_64` | pgdg | 44.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-17-pgtap` | 1.3.4 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_16` | 1.3.4 | `el8.x86_64` | pgdg | 118.1 KiB | [pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | 1.3.3 | `el8.x86_64` | pgdg | 117.3 KiB | [pgtap_16-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtap_16-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | 1.3.4 | `el8.aarch64` | pgdg | 118.1 KiB | [pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | 1.3.4 | `el9.x86_64` | pgdg | 106.8 KiB | [pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | 1.3.1 | `el9.x86_64` | pgdg | 109.5 KiB | [pgtap_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtap_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_16` | 1.3.4 | `el9.aarch64` | pgdg | 106.8 KiB | [pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | 1.3.3 | `el9.aarch64` | pgdg | 106.5 KiB | [pgtap_16-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | 1.3.1 | `el9.aarch64` | pgdg | 109.3 KiB | [pgtap_16-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pgtap` | 1.3.4 | `d12.x86_64` | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-16-pgtap` | 1.3.4 | `d12.aarch64` | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-16-pgtap` | 1.3.4 | `u22.x86_64` | pgdg | 46.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pgtap` | 1.3.4 | `u22.aarch64` | pgdg | 46.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pgtap` | 1.3.4 | `u24.x86_64` | pgdg | 44.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-16-pgtap` | 1.3.4 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_15` | 1.3.4 | `el8.x86_64` | pgdg | 118.1 KiB | [pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | 1.3.3 | `el8.x86_64` | pgdg | 117.3 KiB | [pgtap_15-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | 1.2.0 | `el8.x86_64` | pgdg | 113.9 KiB | [pgtap_15-1.2.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.2.0-1.rhel8.noarch.rpm) |
| `pgtap_15` | 1.3.4 | `el8.aarch64` | pgdg | 118.1 KiB | [pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | 1.3.4 | `el9.x86_64` | pgdg | 106.8 KiB | [pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | 1.3.1 | `el9.x86_64` | pgdg | 109.5 KiB | [pgtap_15-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtap_15-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_15` | 1.3.4 | `el9.aarch64` | pgdg | 106.7 KiB | [pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | 1.3.3 | `el9.aarch64` | pgdg | 106.5 KiB | [pgtap_15-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | 1.3.1 | `el9.aarch64` | pgdg | 109.3 KiB | [pgtap_15-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pgtap` | 1.3.4 | `d12.x86_64` | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-15-pgtap` | 1.3.4 | `d12.aarch64` | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-15-pgtap` | 1.3.4 | `u22.x86_64` | pgdg | 46.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pgtap` | 1.3.4 | `u22.aarch64` | pgdg | 46.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pgtap` | 1.3.4 | `u24.x86_64` | pgdg | 44.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-15-pgtap` | 1.3.4 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_14` | 1.3.4 | `el8.x86_64` | pgdg | 118.1 KiB | [pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | 1.3.3 | `el8.x86_64` | pgdg | 117.3 KiB | [pgtap_14-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | 1.2.0 | `el8.x86_64` | pgdg | 113.9 KiB | [pgtap_14-1.2.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.2.0-1.rhel8.noarch.rpm) |
| `pgtap_14` | 1.1.0 | `el8.x86_64` | pgdg | 111.0 KiB | [pgtap_14-1.1.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.1.0-3.rhel8.noarch.rpm) |
| `pgtap_14` | 1.3.4 | `el8.aarch64` | pgdg | 118.1 KiB | [pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | 1.3.4 | `el9.x86_64` | pgdg | 106.8 KiB | [pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | 1.3.1 | `el9.x86_64` | pgdg | 109.5 KiB | [pgtap_14-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtap_14-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_14` | 1.3.4 | `el9.aarch64` | pgdg | 106.7 KiB | [pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | 1.3.3 | `el9.aarch64` | pgdg | 106.5 KiB | [pgtap_14-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | 1.3.1 | `el9.aarch64` | pgdg | 109.3 KiB | [pgtap_14-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-pgtap` | 1.3.4 | `d12.x86_64` | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-14-pgtap` | 1.3.4 | `d12.aarch64` | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-14-pgtap` | 1.3.4 | `u22.x86_64` | pgdg | 46.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-14-pgtap` | 1.3.4 | `u22.aarch64` | pgdg | 46.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-14-pgtap` | 1.3.4 | `u24.x86_64` | pgdg | 44.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-14-pgtap` | 1.3.4 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pgtap" title="Repository" icon="github" subtitle="github.com/theory/pgtap" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgtap; # install by extension name, for the current active PG version
pig ext install pgtap; # install via package alias, for the active PG version
pig ext install pgtap -v 18;   # install for PG 18
pig ext install pgtap -v 17;   # install for PG 17
pig ext install pgtap -v 16;   # install for PG 16
pig ext install pgtap -v 15;   # install for PG 15
pig ext install pgtap -v 14;   # install for PG 14
pig ext install pgtap -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgtap;
```

