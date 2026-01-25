---
title: "age"
linkTitle: "age"
description: "AGE graph database extension"
weight: 2730
categories: ["FEAT"]
width: full
---

[**age**](https://github.com/apache/age) : AGE graph database extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2730** | {{< badge content="age" link="https://github.com/apache/age" >}} | {{< ext "age" >}} | `1.6.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `ag_catalog` |
|   **See Also**    | {{< ext "pg_graphql" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "ltree" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "citus" >}} |

> [!Note] pg18 = 1.7.0


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.6.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `age` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.0` | {{< bg "18" "apache-age_18" "green" >}} {{< bg "17" "apache-age_17" "green" >}} {{< bg "16" "apache-age_16" "green" >}} {{< bg "15" "apache-age_15" "green" >}} {{< bg "14" "apache-age_14" "green" >}} {{< bg "13" "apache-age_13" "green" >}} | `apache-age_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.0` | {{< bg "18" "postgresql-18-age" "green" >}} {{< bg "17" "postgresql-17-age" "green" >}} {{< bg "16" "postgresql-16-age" "green" >}} {{< bg "15" "postgresql-15-age" "green" >}} {{< bg "14" "postgresql-14-age" "green" >}} {{< bg "13" "postgresql-13-age" "green" >}} | `postgresql-$v-age` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.0" "apache-age_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.7.0" "postgresql-18-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-17-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "postgresql-16-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-15-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-14-age : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.0" "postgresql-13-age : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_18` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 245.8 KiB | [apache-age_18-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_18-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_18` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 228.2 KiB | [apache-age_18-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_18-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_18` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 226.6 KiB | [apache-age_18-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_18-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_18` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 220.9 KiB | [apache-age_18-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_18-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_18` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 228.6 KiB | [apache-age_18-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_18-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_18` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.3 KiB | [apache-age_18-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_18-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-age` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.5 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 661.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 682.3 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 662.7 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 702.8 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 682.6 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 681.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-age` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 660.4 KiB | [postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-18-age/postgresql-18-age_1.7.0~rc0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_17` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 245.1 KiB | [apache-age_17-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_17-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_17` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 227.7 KiB | [apache-age_17-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_17-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_17` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 226.0 KiB | [apache-age_17-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_17-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_17` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 216.8 KiB | [apache-age_17-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_17-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_17` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 228.4 KiB | [apache-age_17-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_17-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_17` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 218.6 KiB | [apache-age_17-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_17-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-age` | `1.6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 677.4 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-age` | `1.6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 655.2 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-age` | `1.6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 676.4 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-age` | `1.6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 657.7 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-age` | `1.6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 791.0 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-age` | `1.6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 770.3 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-age` | `1.6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 675.4 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-age` | `1.6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 654.3 KiB | [postgresql-17-age_1.6.0~rc0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-17-age/postgresql-17-age_1.6.0~rc0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 246.5 KiB | [apache-age_16-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_16-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 228.9 KiB | [apache-age_16-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_16-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 226.9 KiB | [apache-age_16-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_16-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.1 KiB | [apache-age_16-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_16-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_16` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 230.1 KiB | [apache-age_16-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_16-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_16` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.9 KiB | [apache-age_16-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_16-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-age` | `1.6.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 680.5 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 657.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 678.5 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 658.1 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 789.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 769.6 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-age` | `1.6.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 677.0 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-age` | `1.6.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 656.4 KiB | [postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 250.8 KiB | [apache-age_15-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_15-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 233.1 KiB | [apache-age_15-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_15-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.3 KiB | [apache-age_15-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_15-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 224.4 KiB | [apache-age_15-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_15-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_15` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 236.9 KiB | [apache-age_15-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_15-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_15` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 227.1 KiB | [apache-age_15-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_15-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-age` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 728.5 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg120+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_amd64.deb) |
| `postgresql-15-age` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 710.0 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg120+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_arm64.deb) |
| `postgresql-15-age` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 624.9 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg130+2_amd64.deb) |
| `postgresql-15-age` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 608.9 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg130+2_arm64.deb) |
| `postgresql-15-age` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 730.8 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-age` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 711.3 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-age` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 638.2 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-age` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 620.6 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 250.7 KiB | [apache-age_14-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_14-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 232.9 KiB | [apache-age_14-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_14-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.4 KiB | [apache-age_14-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_14-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 222.9 KiB | [apache-age_14-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_14-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_14` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 235.9 KiB | [apache-age_14-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_14-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_14` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 227.0 KiB | [apache-age_14-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_14-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-age` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 728.5 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg120+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_amd64.deb) |
| `postgresql-14-age` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 710.0 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg120+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_arm64.deb) |
| `postgresql-14-age` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 625.1 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg130+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg130+2_amd64.deb) |
| `postgresql-14-age` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 608.0 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg130+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg130+2_arm64.deb) |
| `postgresql-14-age` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 730.2 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-14-age` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 711.6 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-14-age` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 638.4 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-14-age` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 619.2 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `apache-age_13` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 246.7 KiB | [apache-age_13-1.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_13-1.6.0-1PIGSTY.el8.x86_64.rpm) |
| `apache-age_13` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 231.7 KiB | [apache-age_13-1.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_13-1.6.0-1PIGSTY.el8.aarch64.rpm) |
| `apache-age_13` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.0 KiB | [apache-age_13-1.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_13-1.6.0-1PIGSTY.el9.x86_64.rpm) |
| `apache-age_13` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 222.7 KiB | [apache-age_13-1.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_13-1.6.0-1PIGSTY.el9.aarch64.rpm) |
| `apache-age_13` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 236.3 KiB | [apache-age_13-1.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/apache-age_13-1.6.0-1PIGSTY.el10.x86_64.rpm) |
| `apache-age_13` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 226.7 KiB | [apache-age_13-1.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/apache-age_13-1.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-age` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 729.1 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg120+1_amd64.deb) |
| `postgresql-13-age` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 708.6 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg120+1_arm64.deb) |
| `postgresql-13-age` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 726.4 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg130+1_amd64.deb) |
| `postgresql-13-age` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 706.2 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg130+1_arm64.deb) |
| `postgresql-13-age` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 727.0 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-age` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 706.8 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-age` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 636.1 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-age` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 617.4 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/apache/age" title="Repository" icon="github" subtitle="github.com/apache/age" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="age-1.6.0.tar.gz" >}}
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
pig install age -v 16;   # install for PG 16
pig install age -v 15;   # install for PG 15
pig install age -v 14;   # install for PG 14
pig install age -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION age;
```
