---
title: "age"
linkTitle: "age"
description: "AGE graph database extension"
weight: 2600
categories: ["FEAT"]
width: full
---

[**age**](https://github.com/apache/age) : AGE graph database extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2600** | {{< badge content="age" link="https://github.com/apache/age" >}} | {{< ext "age" >}} | `1.7.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `ag_catalog` |
|   **See Also**    | {{< ext "pg_graphql" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "ltree" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "citus" >}} |

> [!Note] pg18/17 = 1.7.0, pg 14-17 with 1.6.0 support, rename apache_age to age on el since 1.7


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `age` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "age_18" "green" >}} {{< bg "17" "age_17" "green" >}} {{< bg "16" "age_16" "red" >}} {{< bg "15" "age_15" "red" >}} {{< bg "14" "age_14" "red" >}} | `age_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "postgresql-18-age" "green" >}} {{< bg "17" "postgresql-17-age" "green" >}} {{< bg "16" "postgresql-16-age" "red" >}} {{< bg "15" "postgresql-15-age" "red" >}} {{< bg "14" "postgresql-14-age" "red" >}} | `postgresql-$v-age` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "age_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.0" "age_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.6.0" "age_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-age : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-age : AVAIL 2" "green" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-14-age : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `age_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 247.6 KiB | [age_18-1.7.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/age_18-1.7.0-2PIGSTY.el8.x86_64.rpm) |
| `age_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 227.0 KiB | [age_18-1.7.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/age_18-1.7.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 229.8 KiB | [age_18-1.7.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/age_18-1.7.0-2PIGSTY.el8.aarch64.rpm) |
| `age_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 209.6 KiB | [age_18-1.7.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/age_18-1.7.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 229.2 KiB | [age_18-1.7.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/age_18-1.7.0-2PIGSTY.el9.x86_64.rpm) |
| `age_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.4 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/age_18-1.7.0-rc0_1PGDG.rhel9.8.x86_64.rpm) |
| `age_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.4 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/age_18-1.7.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.5 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/age_18-1.7.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 221.0 KiB | [age_18-1.7.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/age_18-1.7.0-2PIGSTY.el9.aarch64.rpm) |
| `age_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.0 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/age_18-1.7.0-rc0_1PGDG.rhel9.8.aarch64.rpm) |
| `age_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.0 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/age_18-1.7.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.2 KiB | [age_18-1.7.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/age_18-1.7.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 231.5 KiB | [age_18-1.7.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/age_18-1.7.0-2PIGSTY.el10.x86_64.rpm) |
| `age_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 227.9 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/age_18-1.7.0-rc0_1PGDG.rhel10.2.x86_64.rpm) |
| `age_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 227.9 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/age_18-1.7.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 228.2 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/age_18-1.7.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 221.8 KiB | [age_18-1.7.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/age_18-1.7.0-2PIGSTY.el10.aarch64.rpm) |
| `age_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.2 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/age_18-1.7.0-rc0_1PGDG.rhel10.2.aarch64.rpm) |
| `age_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.3 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/age_18-1.7.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.3 KiB | [age_18-1.7.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/age_18-1.7.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-age` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 682.6 KiB | [postgresql-18-age_1.7.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.5 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 661.9 KiB | [postgresql-18-age_1.7.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 661.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 683.5 KiB | [postgresql-18-age_1.7.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 682.3 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.6 KiB | [postgresql-18-age_1.7.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 662.7 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 764.5 KiB | [postgresql-18-age_1.7.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 702.8 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 753.4 KiB | [postgresql-18-age_1.7.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 682.6 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 729.0 KiB | [postgresql-18-age_1.7.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 681.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 717.0 KiB | [postgresql-18-age_1.7.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 660.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 722.3 KiB | [postgresql-18-age_1.7.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 680.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 710.1 KiB | [postgresql-18-age_1.7.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/age/postgresql-18-age_1.7.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 656.3 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `age_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 247.8 KiB | [age_17-1.7.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/age_17-1.7.0-2PIGSTY.el8.x86_64.rpm) |
| `age_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 227.3 KiB | [age_17-1.7.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/age_17-1.7.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_17` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 224.5 KiB | [age_17-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/age_17-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 229.9 KiB | [age_17-1.7.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/age_17-1.7.0-2PIGSTY.el8.aarch64.rpm) |
| `age_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 209.7 KiB | [age_17-1.7.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/age_17-1.7.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_17` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 207.4 KiB | [age_17-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/age_17-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 229.3 KiB | [age_17-1.7.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/age_17-1.7.0-2PIGSTY.el9.x86_64.rpm) |
| `age_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.5 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/age_17-1.7.0-rc0_1PGDG.rhel9.8.x86_64.rpm) |
| `age_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.5 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/age_17-1.7.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 225.7 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/age_17-1.7.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_17` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 221.2 KiB | [age_17-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/age_17-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_17` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 221.3 KiB | [age_17-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/age_17-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 220.5 KiB | [age_17-1.7.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/age_17-1.7.0-2PIGSTY.el9.aarch64.rpm) |
| `age_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.3 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/age_17-1.7.0-rc0_1PGDG.rhel9.8.aarch64.rpm) |
| `age_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.2 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/age_17-1.7.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 216.4 KiB | [age_17-1.7.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/age_17-1.7.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_17` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.5 KiB | [age_17-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/age_17-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_17` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 212.6 KiB | [age_17-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/age_17-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 231.7 KiB | [age_17-1.7.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/age_17-1.7.0-2PIGSTY.el10.x86_64.rpm) |
| `age_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 227.8 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/age_17-1.7.0-rc0_1PGDG.rhel10.2.x86_64.rpm) |
| `age_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 227.8 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/age_17-1.7.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 228.2 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/age_17-1.7.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_17` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 224.1 KiB | [age_17-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/age_17-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_17` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 224.4 KiB | [age_17-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/age_17-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 222.0 KiB | [age_17-1.7.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/age_17-1.7.0-2PIGSTY.el10.aarch64.rpm) |
| `age_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.4 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/age_17-1.7.0-rc0_1PGDG.rhel10.2.aarch64.rpm) |
| `age_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.5 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/age_17-1.7.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 218.5 KiB | [age_17-1.7.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/age_17-1.7.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `age_17` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 214.7 KiB | [age_17-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/age_17-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_17` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 214.7 KiB | [age_17-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/age_17-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-age` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 682.7 KiB | [postgresql-17-age_1.7.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.8 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 663.1 KiB | [postgresql-17-age_1.7.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 660.8 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 684.3 KiB | [postgresql-17-age_1.7.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 681.8 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.3 KiB | [postgresql-17-age_1.7.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 661.9 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 862.9 KiB | [postgresql-17-age_1.7.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 800.3 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 853.8 KiB | [postgresql-17-age_1.7.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 778.6 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 727.1 KiB | [postgresql-17-age_1.7.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 680.8 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 717.9 KiB | [postgresql-17-age_1.7.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 660.7 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 722.3 KiB | [postgresql-17-age_1.7.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 679.2 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-age` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 710.4 KiB | [postgresql-17-age_1.7.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/age/postgresql-17-age_1.7.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-17-age` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 655.3 KiB | [postgresql-17-age_1.7.0~rc0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.7.0~rc0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `age_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 246.7 KiB | [age_16-1.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/age_16-1.6.0-2PIGSTY.el8.x86_64.rpm) |
| `age_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 225.9 KiB | [age_16-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/age_16-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 229.0 KiB | [age_16-1.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/age_16-1.6.0-2PIGSTY.el8.aarch64.rpm) |
| `age_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 208.6 KiB | [age_16-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/age_16-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 227.1 KiB | [age_16-1.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/age_16-1.6.0-2PIGSTY.el9.x86_64.rpm) |
| `age_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 223.1 KiB | [age_16-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/age_16-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 223.1 KiB | [age_16-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/age_16-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.2 KiB | [age_16-1.6.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/age_16-1.6.0-2PIGSTY.el9.aarch64.rpm) |
| `age_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.8 KiB | [age_16-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/age_16-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 213.9 KiB | [age_16-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/age_16-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_16` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 230.0 KiB | [age_16-1.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/age_16-1.6.0-2PIGSTY.el10.x86_64.rpm) |
| `age_16` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 226.1 KiB | [age_16-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/age_16-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_16` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 226.4 KiB | [age_16-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/age_16-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_16` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.9 KiB | [age_16-1.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/age_16-1.6.0-2PIGSTY.el10.aarch64.rpm) |
| `age_16` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.9 KiB | [age_16-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/age_16-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_16` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.9 KiB | [age_16-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/age_16-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-age` | `1.6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.5 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 657.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 678.5 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.1 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 789.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 769.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 677.0 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 656.4 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 674.8 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 654.0 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `age_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 251.0 KiB | [age_15-1.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/age_15-1.6.0-2PIGSTY.el8.x86_64.rpm) |
| `age_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 229.2 KiB | [age_15-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/age_15-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 233.2 KiB | [age_15-1.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/age_15-1.6.0-2PIGSTY.el8.aarch64.rpm) |
| `age_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 211.7 KiB | [age_15-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/age_15-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.4 KiB | [age_15-1.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/age_15-1.6.0-2PIGSTY.el9.x86_64.rpm) |
| `age_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 228.3 KiB | [age_15-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/age_15-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 228.4 KiB | [age_15-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/age_15-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 224.4 KiB | [age_15-1.6.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/age_15-1.6.0-2PIGSTY.el9.aarch64.rpm) |
| `age_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [age_15-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/age_15-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [age_15-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/age_15-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_15` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 236.8 KiB | [age_15-1.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/age_15-1.6.0-2PIGSTY.el10.x86_64.rpm) |
| `age_15` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 231.0 KiB | [age_15-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/age_15-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_15` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 231.4 KiB | [age_15-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/age_15-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_15` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 227.2 KiB | [age_15-1.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/age_15-1.6.0-2PIGSTY.el10.aarch64.rpm) |
| `age_15` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 221.5 KiB | [age_15-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/age_15-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_15` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 221.7 KiB | [age_15-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/age_15-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-15-age` | `1.6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.1 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-age` | `1.6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 660.2 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-age` | `1.6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 681.5 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-age` | `1.6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 663.0 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-age` | `1.6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 792.5 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-age` | `1.6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 771.1 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-age` | `1.6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 679.2 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-age` | `1.6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 661.5 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-age` | `1.6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 679.3 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-age` | `1.6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 658.2 KiB | [postgresql-15-age_1.6.0~rc0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.6.0~rc0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `age_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 250.8 KiB | [age_14-1.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/age_14-1.6.0-2PIGSTY.el8.x86_64.rpm) |
| `age_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 228.9 KiB | [age_14-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/age_14-1.6.0-rc0_1PGDG.rhel8.10.x86_64.rpm) |
| `age_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 233.1 KiB | [age_14-1.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/age_14-1.6.0-2PIGSTY.el8.aarch64.rpm) |
| `age_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 211.5 KiB | [age_14-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/age_14-1.6.0-rc0_1PGDG.rhel8.10.aarch64.rpm) |
| `age_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.4 KiB | [age_14-1.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/age_14-1.6.0-2PIGSTY.el9.x86_64.rpm) |
| `age_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 228.2 KiB | [age_14-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/age_14-1.6.0-rc0_1PGDG.rhel9.7.x86_64.rpm) |
| `age_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 228.5 KiB | [age_14-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/age_14-1.6.0-rc0_1PGDG.rhel9.6.x86_64.rpm) |
| `age_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [age_14-1.6.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/age_14-1.6.0-2PIGSTY.el9.aarch64.rpm) |
| `age_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.4 KiB | [age_14-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/age_14-1.6.0-rc0_1PGDG.rhel9.7.aarch64.rpm) |
| `age_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 218.6 KiB | [age_14-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/age_14-1.6.0-rc0_1PGDG.rhel9.6.aarch64.rpm) |
| `age_14` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 235.9 KiB | [age_14-1.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/age_14-1.6.0-2PIGSTY.el10.x86_64.rpm) |
| `age_14` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 230.8 KiB | [age_14-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/age_14-1.6.0-rc0_1PGDG.rhel10.1.x86_64.rpm) |
| `age_14` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 231.2 KiB | [age_14-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/age_14-1.6.0-rc0_1PGDG.rhel10.0.x86_64.rpm) |
| `age_14` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 227.0 KiB | [age_14-1.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/age_14-1.6.0-2PIGSTY.el10.aarch64.rpm) |
| `age_14` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 221.5 KiB | [age_14-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/age_14-1.6.0-rc0_1PGDG.rhel10.1.aarch64.rpm) |
| `age_14` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 221.5 KiB | [age_14-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/age_14-1.6.0-rc0_1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-14-age` | `1.6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.9 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-age` | `1.6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 660.7 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-age` | `1.6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 681.5 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-age` | `1.6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 661.4 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-age` | `1.6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 793.0 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-age` | `1.6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 771.4 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-age` | `1.6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 679.2 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-age` | `1.6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 660.7 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-age` | `1.6.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 678.1 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-age` | `1.6.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 657.7 KiB | [postgresql-14-age_1.6.0~rc0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.6.0~rc0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/apache/age" title="Repository" icon="github" subtitle="github.com/apache/age" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="age-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg age;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install age;		# install via package name, for the active PG version

pig install age -v 18;   # install for PG 18
pig install age -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'age';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION age;
```




## Usage

Sources: [Apache AGE repository](https://github.com/apache/age), [PG17 v1.7.0 branch](https://github.com/apache/age/tree/PG17/v1.7.0-rc0), [PG18 v1.7.0 branch](https://github.com/apache/age/tree/PG18/v1.7.0-rc0)

Apache AGE brings graph database capabilities to PostgreSQL using the openCypher query language. It enables hybrid querying that combines SQL and Cypher, property indexes on vertices and edges, and the ability to query multiple graphs.

Each session requires loading the extension:

```sql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
```

### Graph Operations

Create a graph:

```sql
SELECT create_graph('my_graph');
```

Create vertices:

```sql
SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Alice', age: 30})
$$) AS (v agtype);

SELECT * FROM cypher('my_graph', $$
    CREATE (:Person {name: 'Bob', age: 25})
$$) AS (v agtype);
```

Create edges:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person), (b:Person)
    WHERE a.name = 'Alice' AND b.name = 'Bob'
    CREATE (a)-[e:KNOWS {since: 2020}]->(b)
    RETURN e
$$) AS (e agtype);
```

Query the graph:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (v)-[r]-(v2)
    RETURN v, r, v2
$$) AS (v agtype, r agtype, v2 agtype);
```

### Cypher Query Features

AGE supports standard Cypher clauses including `MATCH`, `CREATE`, `SET`, `DELETE`, `RETURN`, `WITH`, `WHERE`, `ORDER BY`, `SKIP`, and `LIMIT`. Data is stored using the `agtype` data type, which extends JSON with graph-specific types for vertices, edges, and paths.

Pattern matching with variable-length paths:

```sql
SELECT * FROM cypher('my_graph', $$
    MATCH (a:Person)-[:KNOWS*1..3]->(b:Person)
    RETURN a.name, b.name
$$) AS (source agtype, target agtype);
```

Hybrid SQL/Cypher queries allow joining graph results with relational tables:

```sql
SELECT t.*, c.* FROM my_table t
JOIN cypher('my_graph', $$
    MATCH (n:Person) RETURN n.name, id(n)
$$) AS c(name agtype, id agtype)
ON t.graph_id = c.id;
```

### Version Notes

Apache AGE 1.7.0 is published through PostgreSQL-major-specific branches for PostgreSQL 17 and 18. The README surface remains the same graph/Cypher workflow, with support focused on the newer PostgreSQL majors rather than PG14-16.
