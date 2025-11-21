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
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltclu" >}} |

> [!Note] missing debian/ubuntu pg18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pljava` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.10` | {{< bg "18" "pljava_18*" "green" >}} {{< bg "17" "pljava_17*" "green" >}} {{< bg "16" "pljava_16*" "green" >}} {{< bg "15" "pljava_15*" "green" >}} {{< bg "14" "pljava_14*" "green" >}} {{< bg "13" "pljava_13*" "green" >}} | `pljava_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.9` | {{< bg "18" "postgresql-18-pljava" "red" >}} {{< bg "17" "postgresql-17-pljava" "green" >}} {{< bg "16" "postgresql-16-pljava" "green" >}} {{< bg "15" "postgresql-15-pljava" "green" >}} {{< bg "14" "postgresql-14-pljava" "green" >}} {{< bg "13" "postgresql-13-pljava" "green" >}} | `postgresql-$v-pljava` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.10" "pljava_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.10" "pljava_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pljava : MISS 0" "red" >}}      | {{< bg "PGDG 1.6.9" "postgresql-17-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-16-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-15-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-14-pljava : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.9" "postgresql-13-pljava : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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
| `pljava_17` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.1 KiB | [pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.5 KiB | [pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pljava_17-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_17` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.0 KiB | [pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.4 KiB | [pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pljava_17-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_17` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.2 KiB | [pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pljava_17-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_17` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.1 KiB | [pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.4 KiB | [pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pljava_17-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_17` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.5 KiB | [pljava_17-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pljava_17-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_17` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 918.2 KiB | [pljava_17-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pljava_17-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_17` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.1 KiB | [pljava_17-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pljava_17-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `pljava_17` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.6 KiB | [pljava_17-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pljava_17-1.6.10-1PGDG.rhel10.aarch64.rpm) |
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
| `pljava_16` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 913.9 KiB | [pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.8 KiB | [pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pljava_16-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_16` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.1 KiB | [pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.7 KiB | [pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pljava_16-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_16` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.8 KiB | [pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pljava_16-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_16` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.1 KiB | [pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 888.2 KiB | [pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.3 KiB | [pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pljava_16-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_16` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.7 KiB | [pljava_16-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pljava_16-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_16` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 918.3 KiB | [pljava_16-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pljava_16-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_16` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.3 KiB | [pljava_16-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pljava_16-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `pljava_16` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.9 KiB | [pljava_16-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pljava_16-1.6.10-1PGDG.rhel10.aarch64.rpm) |
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
| `pljava_15` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.1 KiB | [pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pljava_15-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_15` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 909.9 KiB | [pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.5 KiB | [pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pljava_15-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_15` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.6 KiB | [pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.5 KiB | [pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.4 KiB | [pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pljava_15-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_15` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.0 KiB | [pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 887.6 KiB | [pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.3 KiB | [pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pljava_15-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_15` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.6 KiB | [pljava_15-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pljava_15-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_15` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 917.9 KiB | [pljava_15-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pljava_15-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_15` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.3 KiB | [pljava_15-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pljava_15-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `pljava_15` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.6 KiB | [pljava_15-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pljava_15-1.6.10-1PGDG.rhel10.aarch64.rpm) |
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
| `pljava_14` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 914.2 KiB | [pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.6 KiB | [pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pljava_14-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_14` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 910.0 KiB | [pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.4 KiB | [pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pljava_14-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_14` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 895.1 KiB | [pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.4 KiB | [pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.3 KiB | [pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pljava_14-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_14` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.3 KiB | [pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 887.8 KiB | [pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.2 KiB | [pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pljava_14-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_14` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.5 KiB | [pljava_14-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pljava_14-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_14` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 917.9 KiB | [pljava_14-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pljava_14-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_14` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.4 KiB | [pljava_14-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pljava_14-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `pljava_14` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.7 KiB | [pljava_14-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pljava_14-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 910.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.8 KiB | [postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 910.9 KiB | [postgresql-14-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 906.5 KiB | [postgresql-14-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.7 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.2 KiB | [postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.1 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.6 KiB | [postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-14-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljava_13` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 912.8 KiB | [pljava_13-1.6.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pljava_13-1.6.8-1PGDG.rhel8.x86_64.rpm) |
| `pljava_13` | `1.6.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 926.4 KiB | [pljava_13-1.6.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pljava_13-1.6.10-1PGDG.rhel8.x86_64.rpm) |
| `pljava_13` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 909.7 KiB | [pljava_13-1.6.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pljava_13-1.6.8-1PGDG.rhel8.aarch64.rpm) |
| `pljava_13` | `1.6.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 923.3 KiB | [pljava_13-1.6.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pljava_13-1.6.10-1PGDG.rhel8.aarch64.rpm) |
| `pljava_13` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 894.8 KiB | [pljava_13-1.6.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.8-1PGDG.rhel9.x86_64.rpm) |
| `pljava_13` | `1.6.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 891.4 KiB | [pljava_13-1.6.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.6-1PGDG.rhel9.x86_64.rpm) |
| `pljava_13` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 883.8 KiB | [pljava_13-1.6.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.4-1.rhel9.x86_64.rpm) |
| `pljava_13` | `1.6.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 917.0 KiB | [pljava_13-1.6.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pljava_13-1.6.10-1PGDG.rhel9.x86_64.rpm) |
| `pljava_13` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 892.2 KiB | [pljava_13-1.6.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.8-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | `1.6.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 888.0 KiB | [pljava_13-1.6.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.6-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 879.9 KiB | [pljava_13-1.6.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.4-1.rhel9.aarch64.rpm) |
| `pljava_13` | `1.6.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 914.1 KiB | [pljava_13-1.6.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pljava_13-1.6.10-1PGDG.rhel9.aarch64.rpm) |
| `pljava_13` | `1.6.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 914.1 KiB | [pljava_13-1.6.9-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pljava_13-1.6.9-1PGDG.rhel10.x86_64.rpm) |
| `pljava_13` | `1.6.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 917.5 KiB | [pljava_13-1.6.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pljava_13-1.6.10-1PGDG.rhel10.x86_64.rpm) |
| `pljava_13` | `1.6.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 911.2 KiB | [pljava_13-1.6.9-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pljava_13-1.6.9-1PGDG.rhel10.aarch64.rpm) |
| `pljava_13` | `1.6.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 914.5 KiB | [pljava_13-1.6.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pljava_13-1.6.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pljava` | `1.6.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 910.9 KiB | [postgresql-13-pljava_1.6.9-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg120+1_amd64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 906.0 KiB | [postgresql-13-pljava_1.6.9-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg120+1_arm64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 911.4 KiB | [postgresql-13-pljava_1.6.9-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg130+1_amd64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 906.6 KiB | [postgresql-13-pljava_1.6.9-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg130+1_arm64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 901.3 KiB | [postgresql-13-pljava_1.6.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 897.2 KiB | [postgresql-13-pljava_1.6.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 908.1 KiB | [postgresql-13-pljava_1.6.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pljava` | `1.6.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 904.8 KiB | [postgresql-13-pljava_1.6.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pljava/postgresql-13-pljava_1.6.9-1.pgdg24.04+1_arm64.deb) |

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
pig install pljava -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pljava;
```
