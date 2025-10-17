---
title: "unit"
linkTitle: "unit"
description: "SI units extension"
weight: 3520
categories: ["Type"]
width: full
---

SI units extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3520** | {{< badge content="unit" link="https://github.com/df7cb/postgresql-unit" >}} | {{< ext "unit" "pgunit" >}} | `7.10` | {{< category "TYPE" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pgmp" >}} {{< ext "numeral" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/unit" >}} | `7.10` | {{< badge content="18" color="red" alt="postgresql-unit_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-unit_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/unit" >}} | `7.10` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-unit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "postgresql-unit_18" "7.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql-unit_18-7.10-4PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgresql-unit_17" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-unit_17-7.9-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgresql-unit_16" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-unit_16-7.9-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgresql-unit_15" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.9-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "postgresql-unit_14" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.9-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "postgresql-unit_18" "7.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql-unit_18-7.10-4PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgresql-unit_17" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-unit_17-7.9-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgresql-unit_16" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-unit_16-7.9-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgresql-unit_15" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.9-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "postgresql-unit_14" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.9-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "postgresql-unit_18" "7.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql-unit_18-7.10-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgresql-unit_17" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-unit_17-7.9-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgresql-unit_16" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-unit_16-7.9-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgresql-unit_15" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.9-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "postgresql-unit_14" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.9-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "postgresql-unit_18" "7.10" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql-unit_18-7.10-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgresql-unit_17" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-unit_17-7.9-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgresql-unit_16" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-unit_16-7.9-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgresql-unit_15" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.9-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "postgresql-unit_14" "7.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.9-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-unit" "7.10" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_18` | 7.10 | `el8.aarch64` | pgdg | 127.0 KiB | [postgresql-unit_18-7.10-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql-unit_18-7.10-4PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_18` | 7.10 | `el8.x86_64` | pgdg | 128.4 KiB | [postgresql-unit_18-7.10-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql-unit_18-7.10-4PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_18` | 7.10 | `el9.x86_64` | pgdg | 123.4 KiB | [postgresql-unit_18-7.10-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql-unit_18-7.10-4PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_18` | 7.10 | `el9.aarch64` | pgdg | 122.2 KiB | [postgresql-unit_18-7.10-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql-unit_18-7.10-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-unit` | 7.10 | `d12.x86_64` | pgdg | 158.6 KiB | [postgresql-18-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-18-unit` | 7.10 | `d12.aarch64` | pgdg | 157.1 KiB | [postgresql-18-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-18-unit` | 7.10 | `u22.aarch64` | pgdg | 158.5 KiB | [postgresql-18-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-unit` | 7.10 | `u22.x86_64` | pgdg | 160.6 KiB | [postgresql-18-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-unit` | 7.10 | `u24.x86_64` | pgdg | 158.4 KiB | [postgresql-18-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-unit` | 7.10 | `u24.aarch64` | pgdg | 157.0 KiB | [postgresql-18-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_17` | 7.9 | `el8.aarch64` | pgdg | 89.2 KiB | [postgresql-unit_17-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-unit_17-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_17` | 7.9 | `el8.x86_64` | pgdg | 90.6 KiB | [postgresql-unit_17-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-unit_17-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_17` | 7.10 | `el8.x86_64` | pgdg | 128.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-unit_17-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_17` | 7.10 | `el8.aarch64` | pgdg | 127.0 KiB | [postgresql-unit_17-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-unit_17-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_17` | 7.9 | `el9.x86_64` | pgdg | 88.5 KiB | [postgresql-unit_17-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-unit_17-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_17` | 7.9 | `el9.aarch64` | pgdg | 87.4 KiB | [postgresql-unit_17-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-unit_17-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_17` | 7.10 | `el9.x86_64` | pgdg | 123.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-unit_17-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_17` | 7.10 | `el9.aarch64` | pgdg | 122.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-unit_17-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-unit` | 7.10 | `d12.x86_64` | pgdg | 158.5 KiB | [postgresql-17-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-17-unit` | 7.10 | `d12.aarch64` | pgdg | 157.1 KiB | [postgresql-17-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-17-unit` | 7.10 | `u22.aarch64` | pgdg | 162.7 KiB | [postgresql-17-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-unit` | 7.10 | `u22.x86_64` | pgdg | 164.3 KiB | [postgresql-17-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-unit` | 7.10 | `u24.aarch64` | pgdg | 157.0 KiB | [postgresql-17-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-unit` | 7.10 | `u24.x86_64` | pgdg | 158.5 KiB | [postgresql-17-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_16` | 7.9 | `el8.x86_64` | pgdg | 90.6 KiB | [postgresql-unit_16-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-unit_16-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_16` | 7.9 | `el8.aarch64` | pgdg | 89.2 KiB | [postgresql-unit_16-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-unit_16-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_16` | 7.7 | `el8.x86_64` | pigsty | 90.2 KiB | [postgresql-unit_16-7.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-unit_16-7.7-1PIGSTY.el8.x86_64.rpm) |
| `postgresql-unit_16` | 7.10 | `el8.aarch64` | pgdg | 127.0 KiB | [postgresql-unit_16-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-unit_16-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_16` | 7.10 | `el8.x86_64` | pgdg | 128.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-unit_16-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_16` | 7.9 | `el9.aarch64` | pgdg | 87.4 KiB | [postgresql-unit_16-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-unit_16-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_16` | 7.9 | `el9.x86_64` | pgdg | 88.5 KiB | [postgresql-unit_16-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-unit_16-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_16` | 7.7 | `el9.x86_64` | pigsty | 88.6 KiB | [postgresql-unit_16-7.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-unit_16-7.7-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-unit_16` | 7.10 | `el9.x86_64` | pgdg | 123.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-unit_16-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_16` | 7.10 | `el9.aarch64` | pgdg | 122.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-unit_16-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-unit` | 7.10 | `d12.x86_64` | pgdg | 158.5 KiB | [postgresql-16-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-16-unit` | 7.10 | `d12.aarch64` | pgdg | 157.0 KiB | [postgresql-16-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-16-unit` | 7.10 | `u22.aarch64` | pgdg | 162.6 KiB | [postgresql-16-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-unit` | 7.10 | `u22.x86_64` | pgdg | 164.3 KiB | [postgresql-16-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-unit` | 7.10 | `u24.aarch64` | pgdg | 157.1 KiB | [postgresql-16-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-unit` | 7.10 | `u24.x86_64` | pgdg | 158.6 KiB | [postgresql-16-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_15` | 7.9 | `el8.x86_64` | pgdg | 91.5 KiB | [postgresql-unit_15-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | 7.9 | `el8.aarch64` | pgdg | 89.9 KiB | [postgresql-unit_15-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | 7.4 | `el8.x86_64` | pgdg | 134.9 KiB | [postgresql-unit_15-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | 7.4 | `el8.aarch64` | pgdg | 133.5 KiB | [postgresql-unit_15-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | 7.10 | `el8.x86_64` | pgdg | 129.3 KiB | [postgresql-unit_15-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | 7.10 | `el8.aarch64` | pgdg | 127.7 KiB | [postgresql-unit_15-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | 7.9 | `el9.aarch64` | pgdg | 89.3 KiB | [postgresql-unit_15-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | 7.9 | `el9.x86_64` | pgdg | 90.2 KiB | [postgresql-unit_15-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_15` | 7.4 | `el9.aarch64` | pgdg | 134.9 KiB | [postgresql-unit_15-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | 7.4 | `el9.x86_64` | pgdg | 136.3 KiB | [postgresql-unit_15-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_15` | 7.10 | `el9.aarch64` | pgdg | 124.1 KiB | [postgresql-unit_15-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | 7.10 | `el9.x86_64` | pgdg | 125.1 KiB | [postgresql-unit_15-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-15-unit` | 7.10 | `d12.x86_64` | pgdg | 159.7 KiB | [postgresql-15-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-15-unit` | 7.10 | `d12.aarch64` | pgdg | 157.7 KiB | [postgresql-15-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-15-unit` | 7.10 | `u22.aarch64` | pgdg | 163.8 KiB | [postgresql-15-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-unit` | 7.10 | `u22.x86_64` | pgdg | 165.5 KiB | [postgresql-15-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-unit` | 7.10 | `u24.x86_64` | pgdg | 159.8 KiB | [postgresql-15-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-unit` | 7.10 | `u24.aarch64` | pgdg | 158.1 KiB | [postgresql-15-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_14` | 7.9 | `el8.aarch64` | pgdg | 90.0 KiB | [postgresql-unit_14-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | 7.9 | `el8.x86_64` | pgdg | 91.5 KiB | [postgresql-unit_14-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | 7.4 | `el8.x86_64` | pgdg | 134.9 KiB | [postgresql-unit_14-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | 7.4 | `el8.aarch64` | pgdg | 133.4 KiB | [postgresql-unit_14-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | 7.10 | `el8.aarch64` | pgdg | 127.7 KiB | [postgresql-unit_14-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | 7.10 | `el8.x86_64` | pgdg | 129.3 KiB | [postgresql-unit_14-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | 7.9 | `el9.aarch64` | pgdg | 89.2 KiB | [postgresql-unit_14-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | 7.9 | `el9.x86_64` | pgdg | 90.2 KiB | [postgresql-unit_14-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_14` | 7.4 | `el9.x86_64` | pgdg | 136.2 KiB | [postgresql-unit_14-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_14` | 7.4 | `el9.aarch64` | pgdg | 134.9 KiB | [postgresql-unit_14-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | 7.10 | `el9.aarch64` | pgdg | 124.1 KiB | [postgresql-unit_14-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | 7.10 | `el9.x86_64` | pgdg | 125.0 KiB | [postgresql-unit_14-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-unit` | 7.10 | `d12.x86_64` | pgdg | 159.6 KiB | [postgresql-14-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-14-unit` | 7.10 | `d12.aarch64` | pgdg | 157.6 KiB | [postgresql-14-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-14-unit` | 7.10 | `u22.aarch64` | pgdg | 163.7 KiB | [postgresql-14-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-unit` | 7.10 | `u22.x86_64` | pgdg | 165.5 KiB | [postgresql-14-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-unit` | 7.10 | `u24.x86_64` | pgdg | 159.8 KiB | [postgresql-14-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-unit` | 7.10 | `u24.aarch64` | pgdg | 158.1 KiB | [postgresql-14-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-unit_13` | 7.9 | `el8.x86_64` | pgdg | 91.1 KiB | [postgresql-unit_13-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | 7.9 | `el8.aarch64` | pgdg | 89.9 KiB | [postgresql-unit_13-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | 7.4 | `el8.aarch64` | pgdg | 133.1 KiB | [postgresql-unit_13-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | 7.4 | `el8.x86_64` | pgdg | 134.7 KiB | [postgresql-unit_13-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | 7.10 | `el8.aarch64` | pgdg | 127.7 KiB | [postgresql-unit_13-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | 7.10 | `el8.x86_64` | pgdg | 128.9 KiB | [postgresql-unit_13-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | 7.9 | `el9.aarch64` | pgdg | 89.2 KiB | [postgresql-unit_13-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_13` | 7.9 | `el9.x86_64` | pgdg | 90.0 KiB | [postgresql-unit_13-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | 7.4 | `el9.x86_64` | pgdg | 135.9 KiB | [postgresql-unit_13-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | 7.4 | `el9.aarch64` | pgdg | 134.6 KiB | [postgresql-unit_13-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_13` | 7.10 | `el9.x86_64` | pgdg | 125.0 KiB | [postgresql-unit_13-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | 7.10 | `el9.aarch64` | pgdg | 124.1 KiB | [postgresql-unit_13-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-unit` | 7.10 | `d12.aarch64` | pgdg | 157.8 KiB | [postgresql-13-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-13-unit` | 7.10 | `d12.x86_64` | pgdg | 159.3 KiB | [postgresql-13-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-13-unit` | 7.10 | `u22.x86_64` | pgdg | 165.4 KiB | [postgresql-13-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-unit` | 7.10 | `u22.aarch64` | pgdg | 163.5 KiB | [postgresql-13-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-unit` | 7.10 | `u24.x86_64` | pgdg | 159.8 KiB | [postgresql-13-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-unit` | 7.10 | `u24.aarch64` | pgdg | 158.1 KiB | [postgresql-13-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/postgresql-unit" title="Repository" icon="github" subtitle="github.com/df7cb/postgresql-unit" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="postgresql-unit-7.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build get unit; # get unit source code
pig build dep unit; # install build dependencies
pig build pkg unit; # build extension rpm or deb
pig build ext unit; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install unit; # install by extension name, for the current active PG version
pig ext install pgunit; # install via package alias, for the active PG version
pig ext install unit -v 18;   # install for PG 18
pig ext install unit -v 17;   # install for PG 17
pig ext install unit -v 16;   # install for PG 16
pig ext install unit -v 15;   # install for PG 15
pig ext install unit -v 14;   # install for PG 14
pig ext install unit -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION unit;
```

