---
title: "numeral"
linkTitle: "numeral"
description: "numeral datatypes extension"
weight: 3710
categories: ["TYPE"]
languages: ["C"]
licenses: ["GPL-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**numeral**](https://github.com/df7cb/postgresql-numeral) : numeral datatypes extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3710** | {{< badge content="numeral" link="https://github.com/df7cb/postgresql-numeral" >}} | {{< ext "numeral" >}} | `1.3` | {{< category "TYPE" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "unit" >}} {{< ext "pgmp" >}} {{< ext "pg_rational" >}} {{< ext "uint" >}} {{< ext "uint128" >}} {{< ext "seg" >}} {{< ext "cube" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `numeral` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "postgresql-numeral_18" "green" >}} {{< bg "17" "postgresql-numeral_17" "green" >}} {{< bg "16" "postgresql-numeral_16" "green" >}} {{< bg "15" "postgresql-numeral_15" "green" >}} {{< bg "14" "postgresql-numeral_14" "green" >}} | `postgresql-numeral_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3` | {{< bg "18" "postgresql-18-numeral" "green" >}} {{< bg "17" "postgresql-17-numeral" "green" >}} {{< bg "16" "postgresql-16-numeral" "green" >}} {{< bg "15" "postgresql-15-numeral" "green" >}} {{< bg "14" "postgresql-14-numeral" "green" >}} | `postgresql-$v-numeral` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-numeral_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.3" "postgresql-18-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-17-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-16-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-15-numeral : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3" "postgresql-14-numeral : AVAIL 2" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-numeral_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.7 KiB | [postgresql-numeral_18-1.3-6PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-numeral_18-1.3-6PGSTY.el8.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.5 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql-numeral_18-1.3-3PGDG.rhel8.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.8 KiB | [postgresql-numeral_18-1.3-6PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresql-numeral_18-1.3-6PGSTY.el8.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql-numeral_18-1.3-3PGDG.rhel8.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.9 KiB | [postgresql-numeral_18-1.3-6PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-numeral_18-1.3-6PGSTY.el9.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [postgresql-numeral_18-1.3-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql-numeral_18-1.3-5PGDG.rhel9.8.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.4 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql-numeral_18-1.3-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.5 KiB | [postgresql-numeral_18-1.3-6PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresql-numeral_18-1.3-6PGSTY.el9.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.4 KiB | [postgresql-numeral_18-1.3-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql-numeral_18-1.3-5PGDG.rhel9.8.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.9 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql-numeral_18-1.3-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.0 KiB | [postgresql-numeral_18-1.3-6PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresql-numeral_18-1.3-6PGSTY.el10.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.7 KiB | [postgresql-numeral_18-1.3-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresql-numeral_18-1.3-5PGDG.rhel10.2.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresql-numeral_18-1.3-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.1 KiB | [postgresql-numeral_18-1.3-6PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresql-numeral_18-1.3-6PGSTY.el10.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.0 KiB | [postgresql-numeral_18-1.3-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresql-numeral_18-1.3-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-numeral_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [postgresql-numeral_18-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresql-numeral_18-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg12+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.0 KiB | [postgresql-18-numeral_1.3-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg12+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.2 KiB | [postgresql-18-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.9 KiB | [postgresql-18-numeral_1.3-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg13+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.0 KiB | [postgresql-18-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-18-numeral_1.3-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg13+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-18-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.4 KiB | [postgresql-18-numeral_1.3-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.5 KiB | [postgresql-18-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.1 KiB | [postgresql-18-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-18-numeral_1.3-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-18-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.2 KiB | [postgresql-18-numeral_1.3-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg24.04+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.2 KiB | [postgresql-18-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg24.04+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.3 KiB | [postgresql-18-numeral_1.3-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg26.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.7 KiB | [postgresql-18-numeral_1.3-8.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg26.04+1_amd64.deb) |
| `postgresql-18-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.4 KiB | [postgresql-18-numeral_1.3-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-9.pgdg26.04+1_arm64.deb) |
| `postgresql-18-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.8 KiB | [postgresql-18-numeral_1.3-8.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-18-numeral_1.3-8.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-numeral_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.7 KiB | [postgresql-numeral_17-1.3-6PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-numeral_17-1.3-6PGSTY.el8.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.4 KiB | [postgresql-numeral_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql-numeral_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.8 KiB | [postgresql-numeral_17-1.3-6PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresql-numeral_17-1.3-6PGSTY.el8.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [postgresql-numeral_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql-numeral_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [postgresql-numeral_17-1.3-6PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-numeral_17-1.3-6PGSTY.el9.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [postgresql-numeral_17-1.3-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-numeral_17-1.3-5PGDG.rhel9.8.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [postgresql-numeral_17-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql-numeral_17-1.3-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.5 KiB | [postgresql-numeral_17-1.3-6PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresql-numeral_17-1.3-6PGSTY.el9.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.4 KiB | [postgresql-numeral_17-1.3-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-numeral_17-1.3-5PGDG.rhel9.8.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.0 KiB | [postgresql-numeral_17-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql-numeral_17-1.3-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.0 KiB | [postgresql-numeral_17-1.3-6PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresql-numeral_17-1.3-6PGSTY.el10.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.7 KiB | [postgresql-numeral_17-1.3-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresql-numeral_17-1.3-5PGDG.rhel10.2.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [postgresql-numeral_17-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresql-numeral_17-1.3-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.1 KiB | [postgresql-numeral_17-1.3-6PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresql-numeral_17-1.3-6PGSTY.el10.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.0 KiB | [postgresql-numeral_17-1.3-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresql-numeral_17-1.3-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-numeral_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [postgresql-numeral_17-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresql-numeral_17-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.0 KiB | [postgresql-17-numeral_1.3-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg12+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.1 KiB | [postgresql-17-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.1 KiB | [postgresql-17-numeral_1.3-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg12+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.2 KiB | [postgresql-17-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.0 KiB | [postgresql-17-numeral_1.3-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg13+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.0 KiB | [postgresql-17-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-17-numeral_1.3-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg13+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-17-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 77.3 KiB | [postgresql-17-numeral_1.3-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 77.4 KiB | [postgresql-17-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.1 KiB | [postgresql-17-numeral_1.3-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.1 KiB | [postgresql-17-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.7 KiB | [postgresql-17-numeral_1.3-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg24.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-17-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.2 KiB | [postgresql-17-numeral_1.3-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg24.04+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.2 KiB | [postgresql-17-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg24.04+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.2 KiB | [postgresql-17-numeral_1.3-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg26.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.9 KiB | [postgresql-17-numeral_1.3-8.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg26.04+1_amd64.deb) |
| `postgresql-17-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.4 KiB | [postgresql-17-numeral_1.3-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-9.pgdg26.04+1_arm64.deb) |
| `postgresql-17-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.9 KiB | [postgresql-17-numeral_1.3-8.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-17-numeral_1.3-8.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-numeral_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.7 KiB | [postgresql-numeral_16-1.3-6PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-numeral_16-1.3-6PGSTY.el8.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.4 KiB | [postgresql-numeral_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql-numeral_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.8 KiB | [postgresql-numeral_16-1.3-6PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresql-numeral_16-1.3-6PGSTY.el8.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.1 KiB | [postgresql-numeral_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql-numeral_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [postgresql-numeral_16-1.3-6PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-numeral_16-1.3-6PGSTY.el9.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [postgresql-numeral_16-1.3-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-numeral_16-1.3-5PGDG.rhel9.8.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [postgresql-numeral_16-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql-numeral_16-1.3-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.5 KiB | [postgresql-numeral_16-1.3-6PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresql-numeral_16-1.3-6PGSTY.el9.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.3 KiB | [postgresql-numeral_16-1.3-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-numeral_16-1.3-5PGDG.rhel9.8.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.0 KiB | [postgresql-numeral_16-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql-numeral_16-1.3-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.0 KiB | [postgresql-numeral_16-1.3-6PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresql-numeral_16-1.3-6PGSTY.el10.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.7 KiB | [postgresql-numeral_16-1.3-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresql-numeral_16-1.3-5PGDG.rhel10.2.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [postgresql-numeral_16-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresql-numeral_16-1.3-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.1 KiB | [postgresql-numeral_16-1.3-6PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresql-numeral_16-1.3-6PGSTY.el10.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.0 KiB | [postgresql-numeral_16-1.3-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresql-numeral_16-1.3-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-numeral_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [postgresql-numeral_16-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresql-numeral_16-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.0 KiB | [postgresql-16-numeral_1.3-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg12+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.1 KiB | [postgresql-16-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.0 KiB | [postgresql-16-numeral_1.3-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg12+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.1 KiB | [postgresql-16-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.9 KiB | [postgresql-16-numeral_1.3-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg13+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.0 KiB | [postgresql-16-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-16-numeral_1.3-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg13+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.3 KiB | [postgresql-16-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 77.3 KiB | [postgresql-16-numeral_1.3-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 77.3 KiB | [postgresql-16-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.1 KiB | [postgresql-16-numeral_1.3-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 77.1 KiB | [postgresql-16-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.7 KiB | [postgresql-16-numeral_1.3-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg24.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-16-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.1 KiB | [postgresql-16-numeral_1.3-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg24.04+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 73.2 KiB | [postgresql-16-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg24.04+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.2 KiB | [postgresql-16-numeral_1.3-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg26.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.9 KiB | [postgresql-16-numeral_1.3-8.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg26.04+1_amd64.deb) |
| `postgresql-16-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.4 KiB | [postgresql-16-numeral_1.3-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-9.pgdg26.04+1_arm64.deb) |
| `postgresql-16-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.9 KiB | [postgresql-16-numeral_1.3-8.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-16-numeral_1.3-8.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-numeral_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.8 KiB | [postgresql-numeral_15-1.3-6PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-numeral_15-1.3-6PGSTY.el8.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [postgresql-numeral_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql-numeral_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.7 KiB | [postgresql-numeral_15-1.3-6PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresql-numeral_15-1.3-6PGSTY.el8.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.8 KiB | [postgresql-numeral_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql-numeral_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.5 KiB | [postgresql-numeral_15-1.3-6PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-numeral_15-1.3-6PGSTY.el9.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.0 KiB | [postgresql-numeral_15-1.3-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-numeral_15-1.3-5PGDG.rhel9.8.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.8 KiB | [postgresql-numeral_15-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql-numeral_15-1.3-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.4 KiB | [postgresql-numeral_15-1.3-6PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresql-numeral_15-1.3-6PGSTY.el9.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.9 KiB | [postgresql-numeral_15-1.3-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-numeral_15-1.3-5PGDG.rhel9.8.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.9 KiB | [postgresql-numeral_15-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql-numeral_15-1.3-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.3 KiB | [postgresql-numeral_15-1.3-6PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresql-numeral_15-1.3-6PGSTY.el10.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [postgresql-numeral_15-1.3-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresql-numeral_15-1.3-5PGDG.rhel10.2.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.2 KiB | [postgresql-numeral_15-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresql-numeral_15-1.3-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.5 KiB | [postgresql-numeral_15-1.3-6PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresql-numeral_15-1.3-6PGSTY.el10.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.1 KiB | [postgresql-numeral_15-1.3-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresql-numeral_15-1.3-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-numeral_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.2 KiB | [postgresql-numeral_15-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresql-numeral_15-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.0 KiB | [postgresql-15-numeral_1.3-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg12+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.0 KiB | [postgresql-15-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-15-numeral_1.3-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg12+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-15-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.4 KiB | [postgresql-15-numeral_1.3-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg13+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.4 KiB | [postgresql-15-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.0 KiB | [postgresql-15-numeral_1.3-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg13+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.1 KiB | [postgresql-15-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.8 KiB | [postgresql-15-numeral_1.3-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.9 KiB | [postgresql-15-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 78.8 KiB | [postgresql-15-numeral_1.3-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 78.7 KiB | [postgresql-15-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.7 KiB | [postgresql-15-numeral_1.3-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg24.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.8 KiB | [postgresql-15-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.6 KiB | [postgresql-15-numeral_1.3-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg24.04+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.6 KiB | [postgresql-15-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg24.04+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.1 KiB | [postgresql-15-numeral_1.3-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg26.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.7 KiB | [postgresql-15-numeral_1.3-8.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg26.04+1_amd64.deb) |
| `postgresql-15-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.0 KiB | [postgresql-15-numeral_1.3-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-9.pgdg26.04+1_arm64.deb) |
| `postgresql-15-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.4 KiB | [postgresql-15-numeral_1.3-8.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-15-numeral_1.3-8.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-numeral_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.8 KiB | [postgresql-numeral_14-1.3-6PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresql-numeral_14-1.3-6PGSTY.el8.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.7 KiB | [postgresql-numeral_14-1.3-6PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresql-numeral_14-1.3-6PGSTY.el8.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.8 KiB | [postgresql-numeral_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql-numeral_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 35.5 KiB | [postgresql-numeral_14-1.3-6PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresql-numeral_14-1.3-6PGSTY.el9.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.8 KiB | [postgresql-numeral_14-1.3-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-numeral_14-1.3-5PGDG.rhel9.8.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [postgresql-numeral_14-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql-numeral_14-1.3-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 35.4 KiB | [postgresql-numeral_14-1.3-6PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresql-numeral_14-1.3-6PGSTY.el9.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.9 KiB | [postgresql-numeral_14-1.3-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-numeral_14-1.3-5PGDG.rhel9.8.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.9 KiB | [postgresql-numeral_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql-numeral_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 36.3 KiB | [postgresql-numeral_14-1.3-6PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresql-numeral_14-1.3-6PGSTY.el10.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [postgresql-numeral_14-1.3-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresql-numeral_14-1.3-5PGDG.rhel10.2.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.2 KiB | [postgresql-numeral_14-1.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresql-numeral_14-1.3-3PGDG.rhel10.x86_64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 35.5 KiB | [postgresql-numeral_14-1.3-6PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresql-numeral_14-1.3-6PGSTY.el10.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.1 KiB | [postgresql-numeral_14-1.3-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresql-numeral_14-1.3-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-numeral_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.2 KiB | [postgresql-numeral_14-1.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresql-numeral_14-1.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.1 KiB | [postgresql-14-numeral_1.3-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg12+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 76.0 KiB | [postgresql-14-numeral_1.3-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg12+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-14-numeral_1.3-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg12+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-14-numeral_1.3-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg12+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.3 KiB | [postgresql-14-numeral_1.3-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg13+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.3 KiB | [postgresql-14-numeral_1.3-8.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg13+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.0 KiB | [postgresql-14-numeral_1.3-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg13+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 75.1 KiB | [postgresql-14-numeral_1.3-8.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg13+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.8 KiB | [postgresql-14-numeral_1.3-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 79.8 KiB | [postgresql-14-numeral_1.3-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg22.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 78.7 KiB | [postgresql-14-numeral_1.3-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 78.7 KiB | [postgresql-14-numeral_1.3-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg22.04+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.6 KiB | [postgresql-14-numeral_1.3-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg24.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.7 KiB | [postgresql-14-numeral_1.3-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg24.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.6 KiB | [postgresql-14-numeral_1.3-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg24.04+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 74.6 KiB | [postgresql-14-numeral_1.3-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg24.04+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.1 KiB | [postgresql-14-numeral_1.3-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg26.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.6 KiB | [postgresql-14-numeral_1.3-8.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg26.04+1_amd64.deb) |
| `postgresql-14-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.0 KiB | [postgresql-14-numeral_1.3-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-9.pgdg26.04+1_arm64.deb) |
| `postgresql-14-numeral` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.3 KiB | [postgresql-14-numeral_1.3-8.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-numeral/postgresql-14-numeral_1.3-8.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/postgresql-numeral" title="Repository" icon="github" subtitle="github.com/df7cb/postgresql-numeral" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql-numeral-1.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg numeral;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install numeral;		# install via package name, for the active PG version

pig install numeral -v 18;   # install for PG 18
pig install numeral -v 17;   # install for PG 17
pig install numeral -v 16;   # install for PG 16
pig install numeral -v 15;   # install for PG 15
pig install numeral -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION numeral;
```




## Usage

> [numeral: text numeral data types (English, German, Roman)](https://github.com/df7cb/postgresql-numeral)

The `numeral` extension provides three custom numeric data types that use textual numerals instead of digits.

```sql
CREATE EXTENSION numeral;
```

### Data Types

- **`numeral`**: English numerals using short scale (10^9 = billion)
- **`zahl`**: German numerals using long scale (10^9 = Milliarde)
- **`roman`**: Roman numerals

All three are internally binary-compatible with `bigint` and implicitly cast to it.

### Examples

```sql
-- English numerals
SELECT 'thirty'::numeral + 'twelve'::numeral;
-- forty-two

-- German numerals
SELECT 'siebzehn'::zahl * 'dreiundzwanzig'::zahl;
-- dreihunderteinundneunzig

-- Roman numerals
SELECT 'MCMLV'::roman + 'II'::roman * 'XXX'::roman;
-- MMXV
```

### Operators

Standard arithmetic operators (`+`, `-`, `*`, `/`) work through the implicit `bigint` cast. All existing bigint operators and functions are available.
