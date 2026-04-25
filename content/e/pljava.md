---
title: "pljava"
linkTitle: "pljava"
description: "PL/Java procedural language"
weight: 3090
categories: ["LANG"]
width: full
---

[**pljava**](https://github.com/tada/pljava) : PL/Java procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3090** | {{< badge content="pljava" link="https://github.com/tada/pljava" >}} | {{< ext "pljava" >}} | `1.6.10` | {{< category "LANG" >}} | {{< license "BSD 3-Clause" >}} | {{< language "Java" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `sqlj` |
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltclu" >}} |

> [!Note] missing debian/ubuntu pg18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pljava` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.10` | {{< bg "18" "pljava_18" "green" >}} {{< bg "17" "pljava_17" "green" >}} {{< bg "16" "pljava_16" "green" >}} {{< bg "15" "pljava_15" "green" >}} {{< bg "14" "pljava_14" "green" >}} | `pljava_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.9` | {{< bg "18" "postgresql-18-pljava" "red" >}} {{< bg "17" "postgresql-17-pljava" "green" >}} {{< bg "16" "postgresql-16-pljava" "green" >}} {{< bg "15" "postgresql-15-pljava" "green" >}} {{< bg "14" "postgresql-14-pljava" "green" >}} | `postgresql-$v-pljava` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pljava : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pljava : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pljava : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_18` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.7 KiB | [pljava_18-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pljava_18-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_18` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.4 KiB | [pljava_18-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pljava_18-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_18` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_18-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pljava_18-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_18` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.2 KiB | [pljava_18-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pljava_18-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_18` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 918.3 KiB | [pljava_18-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pljava_18-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_18` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.7 KiB | [pljava_18-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pljava_18-1.6.10-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_17` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.5 KiB | [pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.1 KiB | [pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.4 KiB | [pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.0 KiB | [pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_17` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.2 KiB | [pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.1 KiB | [pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 918.2 KiB | [pljava_17-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pljava_17-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_17` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.5 KiB | [pljava_17-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pljava_17-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.6 KiB | [pljava_17-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pljava_17-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `pljava_17` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.1 KiB | [pljava_17-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pljava_17-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 911.3 KiB | [postgresql-17-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.1 KiB | [postgresql-17-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 911.5 KiB | [postgresql-17-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 906.2 KiB | [postgresql-17-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.7 KiB | [postgresql-17-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.4 KiB | [postgresql-17-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.5 KiB | [postgresql-17-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.7 KiB | [postgresql-17-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-17-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_16` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.8 KiB | [pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 913.9 KiB | [pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.7 KiB | [pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.1 KiB | [pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.8 KiB | [pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.3 KiB | [pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 888.2 KiB | [pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 918.3 KiB | [pljava_16-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pljava_16-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_16` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.7 KiB | [pljava_16-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pljava_16-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.9 KiB | [pljava_16-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pljava_16-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `pljava_16` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.3 KiB | [pljava_16-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pljava_16-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 911.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.3 KiB | [postgresql-16-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 911.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 906.1 KiB | [postgresql-16-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.7 KiB | [postgresql-16-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.5 KiB | [postgresql-16-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.8 KiB | [postgresql-16-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-16-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_15` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.1 KiB | [pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 909.9 KiB | [pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.6 KiB | [pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.5 KiB | [pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.3 KiB | [pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.0 KiB | [pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 887.6 KiB | [pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 917.9 KiB | [pljava_15-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pljava_15-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_15` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.6 KiB | [pljava_15-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pljava_15-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.6 KiB | [pljava_15-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pljava_15-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `pljava_15` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.3 KiB | [pljava_15-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pljava_15-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 911.5 KiB | [postgresql-15-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.0 KiB | [postgresql-15-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 911.5 KiB | [postgresql-15-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 905.9 KiB | [postgresql-15-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.9 KiB | [postgresql-15-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.4 KiB | [postgresql-15-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.2 KiB | [postgresql-15-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.8 KiB | [postgresql-15-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-15-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_14` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.6 KiB | [pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.2 KiB | [pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.4 KiB | [pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.0 KiB | [pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.3 KiB | [pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.1 KiB | [pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.4 KiB | [pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.2 KiB | [pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.3 KiB | [pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 887.8 KiB | [pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 917.9 KiB | [pljava_14-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pljava_14-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_14` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.5 KiB | [pljava_14-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pljava_14-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.7 KiB | [pljava_14-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pljava_14-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `pljava_14` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.4 KiB | [pljava_14-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pljava_14-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 910.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.8 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 910.9 KiB | [postgresql-14-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 906.5 KiB | [postgresql-14-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.2 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.1 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.6 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tada/pljava" title="Repository" icon="github" subtitle="github.com/tada/pljava" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pljava;		# install via package name, for the active PG version

pig install pljava -v 18;   # install for PG 18
pig install pljava -v 17;   # install for PG 17
pig install pljava -v 16;   # install for PG 16
pig install pljava -v 15;   # install for PG 15
pig install pljava -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pljava;
```




## Usage

> [pljava: PL/Java procedural language](https://github.com/tada/pljava)

`pljava` enables writing PostgreSQL functions, triggers, and types in Java using the standard JDBC API.

```sql
CREATE EXTENSION pljava;
```

### Deploy Java Code

Package your Java classes in a JAR file with an SQLJ deployment descriptor, then install it:

```sql
SELECT sqlj.install_jar('file:///path/to/my_functions.jar', 'myjar', true);
SELECT sqlj.set_classpath('public', 'myjar');
```

### Create Functions

Write a Java class with static methods:

```java
package com.example;

import org.postgresql.pljava.annotation.Function;

public class MyFunctions {
    @Function
    public static int add(int a, int b) {
        return a + b;
    }

    @Function
    public static String hello(String name) {
        return "Hello, " + name + "!";
    }
}
```

Declare the SQL function mapping:

```sql
CREATE FUNCTION add(int, int) RETURNS int
  AS 'com.example.MyFunctions.add'
  LANGUAGE java;

CREATE FUNCTION hello(varchar) RETURNS varchar
  AS 'com.example.MyFunctions.hello'
  LANGUAGE java;
```

### Set-Returning Functions

Implement `ResultSetProvider` for set-returning functions:

```java
import org.postgresql.pljava.ResultSetProvider;
import java.sql.ResultSet;
import java.sql.SQLException;

public class MySetFunction implements ResultSetProvider {
    public boolean assignRowValues(ResultSet receiver, int currentRow)
            throws SQLException {
        if (currentRow < 10) {
            receiver.updateInt(1, currentRow);
            receiver.updateString(2, "row " + currentRow);
            return true;
        }
        return false;
    }

    public void close() {}

    public static ResultSetProvider generate()
            throws SQLException {
        return new MySetFunction();
    }
}
```

### Trigger Functions

```java
import org.postgresql.pljava.TriggerData;
import org.postgresql.pljava.annotation.Trigger;

public class MyTrigger {
    @Trigger(called = Trigger.Called.BEFORE, table = "my_table",
             events = {Trigger.Event.INSERT, Trigger.Event.UPDATE})
    public static void auditTrigger(TriggerData td) throws SQLException {
        ResultSet newRow = td.getNew();
        newRow.updateTimestamp("modified_at",
            new java.sql.Timestamp(System.currentTimeMillis()));
    }
}
```

### Database Access via JDBC

```java
import java.sql.*;

public static int countUsers() throws SQLException {
    Connection conn = DriverManager.getConnection("jdbc:default:connection");
    PreparedStatement stmt = conn.prepareStatement("SELECT count(*) FROM users");
    ResultSet rs = stmt.executeQuery();
    rs.next();
    return rs.getInt(1);
}
```

### JAR Management

```sql
SELECT sqlj.install_jar('file:///path/to/jar', 'jarname', true);
SELECT sqlj.replace_jar('file:///path/to/new.jar', 'jarname', true);
SELECT sqlj.remove_jar('jarname', true);
SELECT sqlj.set_classpath('schemaname', 'jar1:jar2');
SELECT sqlj.get_classpath('schemaname');
```
