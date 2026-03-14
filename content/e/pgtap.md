---
title: "pgtap"
linkTitle: "pgtap"
description: "Unit testing for PostgreSQL"
weight: 3200
categories: ["LANG"]
width: full
---

[**pgtap**](https://github.com/theory/pgtap) : Unit testing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3200** | {{< badge content="pgtap" link="https://github.com/theory/pgtap" >}} | {{< ext "pgtap" >}} | `1.3.4` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "plpgsql_check" >}} {{< ext "plpgsql" >}} {{< ext "pldbgapi" >}} {{< ext "plprofiler" >}} {{< ext "faker" >}} {{< ext "unit" >}} {{< ext "dbt2" >}} {{< ext "plperl" >}} |

> [!Note] missing pg17 el9, breaking perl deps


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgtap` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.4` | {{< bg "18" "pgtap_18" "green" >}} {{< bg "17" "pgtap_17" "green" >}} {{< bg "16" "pgtap_16" "green" >}} {{< bg "15" "pgtap_15" "green" >}} {{< bg "14" "pgtap_14" "green" >}} | `pgtap_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.4` | {{< bg "18" "postgresql-18-pgtap" "green" >}} {{< bg "17" "postgresql-17-pgtap" "green" >}} {{< bg "16" "postgresql-16-pgtap" "green" >}} {{< bg "15" "postgresql-15-pgtap" "green" >}} {{< bg "14" "postgresql-14-pgtap" "green" >}} | `postgresql-$v-pgtap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.3.4" "pgtap_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.4" "pgtap_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.4" "postgresql-18-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-17-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-16-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-15-pgtap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.4" "postgresql-14-pgtap : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_18` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 118.1 KiB | [pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 118.1 KiB | [pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgtap_18-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.3 KiB | [pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgtap_18-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_18` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.5 KiB | [pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.7 KiB | [pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgtap_18-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.4 KiB | [pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgtap_18-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_18` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.3 KiB | [pgtap_18-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgtap_18-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.0 KiB | [pgtap_18-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgtap_18-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `pgtap_18` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.3 KiB | [pgtap_18-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgtap_18-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_18` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [pgtap_18-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgtap_18-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pgtap` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 62.1 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-18-pgtap` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-18-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_17` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 118.1 KiB | [pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pgtap_17-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgtap_17-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 118.1 KiB | [pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgtap_17-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_17` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | `1.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.5 KiB | [pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.7 KiB | [pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtap_17-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgtap_17-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_17` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.3 KiB | [pgtap_17-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgtap_17-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_17` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.0 KiB | [pgtap_17-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgtap_17-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `pgtap_17` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.3 KiB | [pgtap_17-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgtap_17-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_17` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [pgtap_17-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgtap_17-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pgtap` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 62.1 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-17-pgtap` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-17-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_16` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 118.1 KiB | [pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pgtap_16-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgtap_16-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 118.1 KiB | [pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgtap_16-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_16` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.5 KiB | [pgtap_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgtap_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_16` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.8 KiB | [pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [pgtap_16-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.3 KiB | [pgtap_16-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgtap_16-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtap_16` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.3 KiB | [pgtap_16-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgtap_16-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_16` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.0 KiB | [pgtap_16-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgtap_16-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `pgtap_16` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.3 KiB | [pgtap_16-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgtap_16-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_16` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [pgtap_16-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgtap_16-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pgtap` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 62.1 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-16-pgtap` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-16-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_15` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 118.1 KiB | [pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pgtap_15-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [pgtap_15-1.2.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgtap_15-1.2.0-1.rhel8.noarch.rpm) |
| `pgtap_15` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 118.1 KiB | [pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgtap_15-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_15` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.5 KiB | [pgtap_15-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgtap_15-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_15` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.7 KiB | [pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [pgtap_15-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.3 KiB | [pgtap_15-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgtap_15-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtap_15` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.3 KiB | [pgtap_15-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgtap_15-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_15` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.0 KiB | [pgtap_15-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgtap_15-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `pgtap_15` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.3 KiB | [pgtap_15-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgtap_15-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_15` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [pgtap_15-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgtap_15-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pgtap` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 62.1 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-15-pgtap` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-15-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtap_14` | `1.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 118.1 KiB | [pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | `1.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 117.3 KiB | [pgtap_14-1.3.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.3.3-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [pgtap_14-1.2.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.2.0-1.rhel8.noarch.rpm) |
| `pgtap_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.0 KiB | [pgtap_14-1.1.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgtap_14-1.1.0-3.rhel8.noarch.rpm) |
| `pgtap_14` | `1.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 118.1 KiB | [pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgtap_14-1.3.4-1PGDG.rhel8.noarch.rpm) |
| `pgtap_14` | `1.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.5 KiB | [pgtap_14-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgtap_14-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgtap_14` | `1.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.7 KiB | [pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.4-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | `1.3.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 106.5 KiB | [pgtap_14-1.3.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.3-1PGDG.rhel9.noarch.rpm) |
| `pgtap_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.3 KiB | [pgtap_14-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgtap_14-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgtap_14` | `1.3.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.3 KiB | [pgtap_14-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgtap_14-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_14` | `1.3.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.0 KiB | [pgtap_14-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgtap_14-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `pgtap_14` | `1.3.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.3 KiB | [pgtap_14-1.3.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgtap_14-1.3.4-1PGDG.rhel10.noarch.rpm) |
| `pgtap_14` | `1.3.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [pgtap_14-1.3.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgtap_14-1.3.3-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pgtap` | `1.3.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg12+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 62.1 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg13+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg22.04+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |
| `postgresql-14-pgtap` | `1.3.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.9 KiB | [postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgtap/postgresql-14-pgtap_1.3.4-1.pgdg24.04+1_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pgtap" title="Repository" icon="github" subtitle="github.com/theory/pgtap" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgtap;		# install via package name, for the active PG version

pig install pgtap -v 18;   # install for PG 18
pig install pgtap -v 17;   # install for PG 17
pig install pgtap -v 16;   # install for PG 16
pig install pgtap -v 15;   # install for PG 15
pig install pgtap -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgtap CASCADE; -- requires plpgsql
```




## Usage

> [pgtap: Unit testing for PostgreSQL](https://github.com/theory/pgtap)

`pgtap` is a unit testing framework for PostgreSQL that produces TAP (Test Anything Protocol) output, providing hundreds of assertion functions for testing database objects and query results.

```sql
CREATE EXTENSION pgtap;
```

### Test Structure

```sql
BEGIN;
SELECT plan(3);  -- declare how many tests to run

SELECT ok(1 = 1, 'one equals one');
SELECT is(1 + 1, 2, 'addition works');
SELECT isnt(1, 2, 'one is not two');

SELECT * FROM finish();
ROLLBACK;
```

Use `no_plan()` when the test count is not known in advance:

```sql
BEGIN;
SELECT * FROM no_plan();
-- ... tests ...
SELECT * FROM finish();
ROLLBACK;
```

### Basic Assertions

```sql
SELECT ok(expression, description);           -- boolean test
SELECT is(got, expected, description);         -- equality test
SELECT isnt(got, unexpected, description);     -- inequality test
SELECT matches(value, regex, description);     -- regex match
```

### Schema Testing

```sql
SELECT has_table('users');
SELECT has_table('myschema', 'users', 'users table exists');
SELECT has_column('users', 'email');
SELECT col_type_is('users', 'email', 'text');
SELECT col_not_null('users', 'id');
SELECT col_has_default('users', 'created_at');
SELECT has_function('calculate_total');
SELECT has_function('calculate_total', ARRAY['integer', 'numeric']);
SELECT has_index('users', 'users_email_idx');
SELECT has_pk('users');
SELECT has_fk('orders');
```

### Error Testing

```sql
SELECT lives_ok('INSERT INTO t(id) VALUES (1)', 'insert succeeds');
SELECT throws_ok(
  'SELECT 1/0',
  '22012',          -- SQLSTATE for division by zero
  'division by zero'
);
```

### Query Result Testing

```sql
-- Compare ordered result sets
SELECT results_eq(
  'SELECT * FROM active_users()',
  'SELECT * FROM users WHERE active',
  'active_users returns correct rows'
);

-- Compare unordered result sets
SELECT set_eq(
  'SELECT * FROM active_ids()',
  ARRAY[2, 3, 4, 5]
);

-- Check query returns no rows
SELECT is_empty('SELECT * FROM users WHERE id = -1');

-- Compare bag (multiset) results
SELECT bag_eq(
  'SELECT color FROM items',
  $$VALUES ('red'), ('blue'), ('red')$$
);
```

### Running Tests with pg_prove

```bash
pg_prove -d mydb tests/*.sql
pg_prove -d mydb --ext .sql --recurse tests/
```

### xUnit Style

```sql
CREATE FUNCTION test_my_feature() RETURNS SETOF text AS $$
  RETURN NEXT ok(1 = 1, 'basic check');
  RETURN NEXT is(my_func(1), 42, 'function works');
$$ LANGUAGE plpgsql;

SELECT * FROM runtests('test_my_feature');
```
