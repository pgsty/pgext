---
title: "unit"
linkTitle: "unit"
description: "SI units extension"
weight: 3520
categories: ["TYPE"]
width: full
---

[**pgunit**](https://github.com/df7cb/postgresql-unit) : SI units extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3520** | {{< badge content="unit" link="https://github.com/df7cb/postgresql-unit" >}} | {{< ext "unit" "pgunit" >}} | `7.10` | {{< category "TYPE" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pgmp" >}} {{< ext "numeral" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `7.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgunit` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `7.10` | {{< bg "18" "postgresql-unit_18*" "green" >}} {{< bg "17" "postgresql-unit_17*" "green" >}} {{< bg "16" "postgresql-unit_16*" "green" >}} {{< bg "15" "postgresql-unit_15*" "green" >}} {{< bg "14" "postgresql-unit_14*" "green" >}} {{< bg "13" "postgresql-unit_13*" "green" >}} | `postgresql-unit_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `7.10` | {{< bg "18" "postgresql-18-unit" "green" >}} {{< bg "17" "postgresql-17-unit" "green" >}} {{< bg "16" "postgresql-16-unit" "green" >}} {{< bg "15" "postgresql-15-unit" "green" >}} {{< bg "14" "postgresql-14-unit" "green" >}} {{< bg "13" "postgresql-13-unit" "green" >}} | `postgresql-$v-unit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-unit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-unit_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 7.10" "postgresql-18-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-17-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-16-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-15-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-14-unit : AVAIL 1" "blue" >}} | {{< bg "PGDG 7.10" "postgresql-13-unit : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_18` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 128.4 KiB | [postgresql-unit_18-7.10-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql-unit_18-7.10-4PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_18` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.0 KiB | [postgresql-unit_18-7.10-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql-unit_18-7.10-4PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_18` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.4 KiB | [postgresql-unit_18-7.10-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql-unit_18-7.10-4PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_18` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 122.2 KiB | [postgresql-unit_18-7.10-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql-unit_18-7.10-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_18` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 123.9 KiB | [postgresql-unit_18-7.10-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresql-unit_18-7.10-4PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_18` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.3 KiB | [postgresql-unit_18-7.10-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresql-unit_18-7.10-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 158.6 KiB | [postgresql-18-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-18-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.1 KiB | [postgresql-18-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-18-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 158.4 KiB | [postgresql-18-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-18-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.2 KiB | [postgresql-18-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-18-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 160.6 KiB | [postgresql-18-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 158.5 KiB | [postgresql-18-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 158.4 KiB | [postgresql-18-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 157.0 KiB | [postgresql-18-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-18-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_17` | `7.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 90.6 KiB | [postgresql-unit_17-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-unit_17-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_17` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 128.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-unit_17-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_17` | `7.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.2 KiB | [postgresql-unit_17-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-unit_17-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_17` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.0 KiB | [postgresql-unit_17-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-unit_17-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_17` | `7.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.5 KiB | [postgresql-unit_17-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-unit_17-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_17` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-unit_17-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_17` | `7.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.4 KiB | [postgresql-unit_17-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-unit_17-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_17` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 122.4 KiB | [postgresql-unit_17-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-unit_17-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_17` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 123.9 KiB | [postgresql-unit_17-7.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresql-unit_17-7.10-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_17` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.2 KiB | [postgresql-unit_17-7.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresql-unit_17-7.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 158.5 KiB | [postgresql-17-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-17-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.1 KiB | [postgresql-17-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-17-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 158.5 KiB | [postgresql-17-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-17-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.2 KiB | [postgresql-17-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-17-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 164.3 KiB | [postgresql-17-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 162.7 KiB | [postgresql-17-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 158.5 KiB | [postgresql-17-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 157.0 KiB | [postgresql-17-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-17-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_16` | `7.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 90.6 KiB | [postgresql-unit_16-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-unit_16-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_16` | `7.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 90.2 KiB | [postgresql-unit_16-7.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-unit_16-7.7-1PIGSTY.el8.x86_64.rpm) |
| `postgresql-unit_16` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 128.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-unit_16-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_16` | `7.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.2 KiB | [postgresql-unit_16-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-unit_16-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_16` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.0 KiB | [postgresql-unit_16-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-unit_16-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_16` | `7.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 88.5 KiB | [postgresql-unit_16-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-unit_16-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_16` | `7.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.6 KiB | [postgresql-unit_16-7.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-unit_16-7.7-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-unit_16` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-unit_16-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_16` | `7.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.4 KiB | [postgresql-unit_16-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-unit_16-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_16` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 122.4 KiB | [postgresql-unit_16-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-unit_16-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_16` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 123.9 KiB | [postgresql-unit_16-7.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresql-unit_16-7.10-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_16` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.2 KiB | [postgresql-unit_16-7.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresql-unit_16-7.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 158.5 KiB | [postgresql-16-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-16-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.0 KiB | [postgresql-16-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-16-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 158.5 KiB | [postgresql-16-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-16-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.1 KiB | [postgresql-16-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-16-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 164.3 KiB | [postgresql-16-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 162.6 KiB | [postgresql-16-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 158.6 KiB | [postgresql-16-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 157.1 KiB | [postgresql-16-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-16-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_15` | `7.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 91.5 KiB | [postgresql-unit_15-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | `7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 134.9 KiB | [postgresql-unit_15-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.3 KiB | [postgresql-unit_15-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-unit_15-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_15` | `7.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.9 KiB | [postgresql-unit_15-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | `7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.5 KiB | [postgresql-unit_15-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.7 KiB | [postgresql-unit_15-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-unit_15-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_15` | `7.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 90.2 KiB | [postgresql-unit_15-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_15` | `7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 136.3 KiB | [postgresql-unit_15-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_15` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 125.1 KiB | [postgresql-unit_15-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-unit_15-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_15` | `7.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.3 KiB | [postgresql-unit_15-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | `7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 134.9 KiB | [postgresql-unit_15-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 124.1 KiB | [postgresql-unit_15-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-unit_15-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_15` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 125.8 KiB | [postgresql-unit_15-7.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresql-unit_15-7.10-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_15` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 124.7 KiB | [postgresql-unit_15-7.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresql-unit_15-7.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 159.7 KiB | [postgresql-15-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-15-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.7 KiB | [postgresql-15-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-15-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 159.6 KiB | [postgresql-15-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-15-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.9 KiB | [postgresql-15-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-15-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 165.5 KiB | [postgresql-15-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 163.8 KiB | [postgresql-15-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 159.8 KiB | [postgresql-15-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 158.1 KiB | [postgresql-15-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-15-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_14` | `7.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 91.5 KiB | [postgresql-unit_14-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | `7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 134.9 KiB | [postgresql-unit_14-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.3 KiB | [postgresql-unit_14-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql-unit_14-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_14` | `7.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.0 KiB | [postgresql-unit_14-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | `7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.4 KiB | [postgresql-unit_14-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.7 KiB | [postgresql-unit_14-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-unit_14-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_14` | `7.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 90.2 KiB | [postgresql-unit_14-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_14` | `7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 136.2 KiB | [postgresql-unit_14-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_14` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 125.0 KiB | [postgresql-unit_14-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-unit_14-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_14` | `7.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.2 KiB | [postgresql-unit_14-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | `7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 134.9 KiB | [postgresql-unit_14-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 124.1 KiB | [postgresql-unit_14-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-unit_14-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_14` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 125.8 KiB | [postgresql-unit_14-7.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresql-unit_14-7.10-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_14` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 124.7 KiB | [postgresql-unit_14-7.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresql-unit_14-7.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 159.6 KiB | [postgresql-14-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-14-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.6 KiB | [postgresql-14-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-14-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 159.3 KiB | [postgresql-14-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-14-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.9 KiB | [postgresql-14-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-14-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 165.5 KiB | [postgresql-14-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 163.7 KiB | [postgresql-14-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 159.8 KiB | [postgresql-14-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 158.1 KiB | [postgresql-14-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-14-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-unit_13` | `7.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 91.1 KiB | [postgresql-unit_13-7.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.9-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | `7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 134.7 KiB | [postgresql-unit_13-7.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.4-1.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | `7.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 128.9 KiB | [postgresql-unit_13-7.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql-unit_13-7.10-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-unit_13` | `7.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.9 KiB | [postgresql-unit_13-7.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.9-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | `7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.1 KiB | [postgresql-unit_13-7.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.4-1.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | `7.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.7 KiB | [postgresql-unit_13-7.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql-unit_13-7.10-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-unit_13` | `7.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 90.0 KiB | [postgresql-unit_13-7.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.9-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | `7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 135.9 KiB | [postgresql-unit_13-7.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.4-1.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | `7.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 125.0 KiB | [postgresql-unit_13-7.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql-unit_13-7.10-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-unit_13` | `7.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.2 KiB | [postgresql-unit_13-7.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.9-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_13` | `7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 134.6 KiB | [postgresql-unit_13-7.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.4-1.rhel9.aarch64.rpm) |
| `postgresql-unit_13` | `7.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 124.1 KiB | [postgresql-unit_13-7.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql-unit_13-7.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-unit_13` | `7.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 125.7 KiB | [postgresql-unit_13-7.10-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/postgresql-unit_13-7.10-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-unit_13` | `7.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 124.8 KiB | [postgresql-unit_13-7.10-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/postgresql-unit_13-7.10-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-unit` | `7.10` | [d12.x86_64](/os/d12.x86_64) | pgdg | 159.3 KiB | [postgresql-13-unit_7.10-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg12+1_amd64.deb) |
| `postgresql-13-unit` | `7.10` | [d12.aarch64](/os/d12.aarch64) | pgdg | 157.8 KiB | [postgresql-13-unit_7.10-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg12+1_arm64.deb) |
| `postgresql-13-unit` | `7.10` | [d13.x86_64](/os/d13.x86_64) | pgdg | 159.4 KiB | [postgresql-13-unit_7.10-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg13+1_amd64.deb) |
| `postgresql-13-unit` | `7.10` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.8 KiB | [postgresql-13-unit_7.10-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg13+1_arm64.deb) |
| `postgresql-13-unit` | `7.10` | [u22.x86_64](/os/u22.x86_64) | pgdg | 165.4 KiB | [postgresql-13-unit_7.10-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-unit` | `7.10` | [u22.aarch64](/os/u22.aarch64) | pgdg | 163.5 KiB | [postgresql-13-unit_7.10-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-unit` | `7.10` | [u24.x86_64](/os/u24.x86_64) | pgdg | 159.8 KiB | [postgresql-13-unit_7.10-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-unit` | `7.10` | [u24.aarch64](/os/u24.aarch64) | pgdg | 158.1 KiB | [postgresql-13-unit_7.10-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-unit/postgresql-13-unit_7.10-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/postgresql-unit" title="Repository" icon="github" subtitle="github.com/df7cb/postgresql-unit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql-unit-7.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgunit;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgunit;		# install via package name, for the active PG version
pig install unit;		# install by extension name, for the current active PG version

pig install unit -v 18;   # install for PG 18
pig install unit -v 17;   # install for PG 17
pig install unit -v 16;   # install for PG 16
pig install unit -v 15;   # install for PG 15
pig install unit -v 14;   # install for PG 14
pig install unit -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION unit CASCADE; -- requires plpgsql
```
