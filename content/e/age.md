---
title: "age"
linkTitle: "age"
description: "AGE graph database extension"
weight: 2760
categories: ["Feat"]
width: full
---

AGE graph database extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2760** | {{< badge content="age" link="https://github.com/apache/age" >}} | {{< ext "age" "age" >}} | `1.5.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_graphql" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "ltree" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "citus" >}} |

> [!Note] no longer maintained


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/age" >}} | `1.5.0` | {{< badge content="18" color="red" alt="apache-age_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `apache-age_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/age" >}} | `1.5.0` | {{< badge content="18" color="red" alt="postgresql-18-age" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-age` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "apache-age_18" >}}     | {{< pkg "apache-age_17" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_17-1.5.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "apache-age_16" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_16-1.5.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "apache-age_15" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_15-1.5.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "apache-age_14" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_14-1.5.0-2PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "apache-age_18" >}}     | {{< pkg "apache-age_17" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_17-1.5.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "apache-age_16" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_16-1.5.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "apache-age_15" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_15-1.5.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "apache-age_14" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_14-1.5.0-2PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "apache-age_18" >}}     | {{< pkg "apache-age_17" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_17-1.5.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "apache-age_16" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_16-1.5.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "apache-age_15" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_15-1.5.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "apache-age_14" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_14-1.5.0-2PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "apache-age_18" >}}     | {{< pkg "apache-age_17" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_17-1.5.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "apache-age_16" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_16-1.5.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "apache-age_15" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_15-1.5.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "apache-age_14" "1.5.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_14-1.5.0-2PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_amd64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_arm64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-age" >}}     | {{< pkg "postgresql-17-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-age" "1.6.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb" >}} | {{< pkg "postgresql-14-age" "1.5.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `apache-age_17` | 1.5.0 | `el8.x86_64` | pigsty | 219.4 KiB | [apache-age_17-1.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_17-1.5.0-2PIGSTY.el8.x86_64.rpm) |
| `apache-age_17` | 1.5.0 | `el8.aarch64` | pigsty | 201.7 KiB | [apache-age_17-1.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_17-1.5.0-2PIGSTY.el8.aarch64.rpm) |
| `apache-age_17` | 1.5.0 | `el9.aarch64` | pigsty | 208.1 KiB | [apache-age_17-1.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_17-1.5.0-2PIGSTY.el9.aarch64.rpm) |
| `apache-age_17` | 1.5.0 | `el9.x86_64` | pigsty | 216.6 KiB | [apache-age_17-1.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_17-1.5.0-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-age` | 1.6.0 | `d12.x86_64` | pgdg | 677.4 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-age` | 1.6.0 | `d12.aarch64` | pgdg | 657.3 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-age` | 1.6.0 | `u22.x86_64` | pgdg | 793.4 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-age` | 1.6.0 | `u22.aarch64` | pgdg | 773.4 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-age` | 1.6.0 | `u24.x86_64` | pgdg | 677.0 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-age` | 1.6.0 | `u24.aarch64` | pgdg | 654.9 KiB | [postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-17-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `apache-age_16` | 1.5.0 | `el8.x86_64` | pigsty | 205.1 KiB | [apache-age_16-1.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_16-1.5.0-2PIGSTY.el8.x86_64.rpm) |
| `apache-age_16` | 1.5.0 | `el8.aarch64` | pigsty | 188.7 KiB | [apache-age_16-1.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_16-1.5.0-2PIGSTY.el8.aarch64.rpm) |
| `apache-age_16` | 1.5.0 | `el9.x86_64` | pigsty | 202.3 KiB | [apache-age_16-1.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_16-1.5.0-2PIGSTY.el9.x86_64.rpm) |
| `apache-age_16` | 1.5.0 | `el9.aarch64` | pigsty | 195.5 KiB | [apache-age_16-1.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_16-1.5.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-age` | 1.6.0 | `d12.x86_64` | pgdg | 681.0 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-age` | 1.6.0 | `d12.aarch64` | pgdg | 657.9 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-age` | 1.6.0 | `u22.aarch64` | pgdg | 769.2 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-age` | 1.6.0 | `u22.x86_64` | pgdg | 790.9 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-age` | 1.6.0 | `u24.x86_64` | pgdg | 677.1 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-age` | 1.6.0 | `u24.aarch64` | pgdg | 657.7 KiB | [postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-16-age/postgresql-16-age_1.6.0~rc0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `apache-age_15` | 1.5.0 | `el8.x86_64` | pigsty | 207.9 KiB | [apache-age_15-1.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_15-1.5.0-2PIGSTY.el8.x86_64.rpm) |
| `apache-age_15` | 1.5.0 | `el8.aarch64` | pigsty | 191.5 KiB | [apache-age_15-1.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_15-1.5.0-2PIGSTY.el8.aarch64.rpm) |
| `apache-age_15` | 1.5.0 | `el9.x86_64` | pigsty | 208.4 KiB | [apache-age_15-1.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_15-1.5.0-2PIGSTY.el9.x86_64.rpm) |
| `apache-age_15` | 1.5.0 | `el9.aarch64` | pigsty | 200.2 KiB | [apache-age_15-1.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_15-1.5.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-age` | 1.5.0 | `d12.aarch64` | pgdg | 710.0 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg120+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_arm64.deb) |
| `postgresql-15-age` | 1.5.0 | `d12.x86_64` | pgdg | 728.5 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg120+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg120+2_amd64.deb) |
| `postgresql-15-age` | 1.5.0 | `u22.aarch64` | pgdg | 711.3 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-age` | 1.5.0 | `u22.x86_64` | pgdg | 730.8 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-age` | 1.5.0 | `u24.x86_64` | pgdg | 638.2 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-age` | 1.5.0 | `u24.aarch64` | pgdg | 620.6 KiB | [postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-15-age/postgresql-15-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `apache-age_14` | 1.5.0 | `el8.x86_64` | pigsty | 207.9 KiB | [apache-age_14-1.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_14-1.5.0-2PIGSTY.el8.x86_64.rpm) |
| `apache-age_14` | 1.5.0 | `el8.aarch64` | pigsty | 191.4 KiB | [apache-age_14-1.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_14-1.5.0-2PIGSTY.el8.aarch64.rpm) |
| `apache-age_14` | 1.5.0 | `el9.x86_64` | pigsty | 211.6 KiB | [apache-age_14-1.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_14-1.5.0-2PIGSTY.el9.x86_64.rpm) |
| `apache-age_14` | 1.5.0 | `el9.aarch64` | pigsty | 200.2 KiB | [apache-age_14-1.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_14-1.5.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-age` | 1.5.0 | `d12.x86_64` | pgdg | 728.5 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg120+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_amd64.deb) |
| `postgresql-14-age` | 1.5.0 | `d12.aarch64` | pgdg | 710.0 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg120+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg120+2_arm64.deb) |
| `postgresql-14-age` | 1.5.0 | `u22.x86_64` | pgdg | 730.2 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_amd64.deb) |
| `postgresql-14-age` | 1.5.0 | `u22.aarch64` | pgdg | 711.6 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg22.04+2_arm64.deb) |
| `postgresql-14-age` | 1.5.0 | `u24.x86_64` | pgdg | 638.4 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_amd64.deb) |
| `postgresql-14-age` | 1.5.0 | `u24.aarch64` | pgdg | 619.2 KiB | [postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-14-age/postgresql-14-age_1.5.0~rc0-1.pgdg24.04+2_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `apache-age_13` | 1.5.0 | `el8.aarch64` | pigsty | 191.4 KiB | [apache-age_13-1.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/apache-age_13-1.5.0-2PIGSTY.el8.aarch64.rpm) |
| `apache-age_13` | 1.5.0 | `el8.x86_64` | pigsty | 204.9 KiB | [apache-age_13-1.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/apache-age_13-1.5.0-2PIGSTY.el8.x86_64.rpm) |
| `apache-age_13` | 1.5.0 | `el9.aarch64` | pigsty | 199.4 KiB | [apache-age_13-1.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/apache-age_13-1.5.0-2PIGSTY.el9.aarch64.rpm) |
| `apache-age_13` | 1.5.0 | `el9.x86_64` | pigsty | 209.4 KiB | [apache-age_13-1.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/apache-age_13-1.5.0-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-age` | 1.5.0 | `d12.aarch64` | pgdg | 708.6 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg120+1_arm64.deb) |
| `postgresql-13-age` | 1.5.0 | `d12.x86_64` | pgdg | 729.1 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg120+1_amd64.deb) |
| `postgresql-13-age` | 1.5.0 | `u22.aarch64` | pgdg | 706.8 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-age` | 1.5.0 | `u22.x86_64` | pgdg | 727.0 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-age` | 1.5.0 | `u24.aarch64` | pgdg | 617.4 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-13-age` | 1.5.0 | `u24.x86_64` | pgdg | 636.1 KiB | [postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-13-age/postgresql-13-age_1.5.0~rc0-1.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/apache/age" title="Repository" icon="github" subtitle="github.com/apache/age" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="age-1.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get age; # get age source code
pig build dep age; # install build dependencies
pig build pkg age; # build extension rpm or deb
pig build ext age; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install age; # install by extension name, for the current active PG version
pig ext install age; # install via package alias, for the active PG version
pig ext install age -v 17;   # install for PG 17
pig ext install age -v 16;   # install for PG 16
pig ext install age -v 15;   # install for PG 15
pig ext install age -v 14;   # install for PG 14
pig ext install age -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION age CASCADE SCHEMA ag_catalog;
```

