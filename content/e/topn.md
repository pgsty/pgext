---
title: "topn"
linkTitle: "topn"
description: "type for top-n JSONB"
weight: 4600
categories: ["Func"]
width: full
---

type for top-n JSONB

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4600** | {{< badge content="topn" link="https://github.com/citusdata/postgresql-topn" >}} | {{< ext "topn" "topn" >}} | `2.7.0` | {{< category "FUNC" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "count_distinct" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "first_last_agg" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/topn" >}} | `2.7.0` | {{< badge content="18" color="red" alt="topn_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `topn_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/topn" >}} | `2.7.0` | {{< badge content="18" color="red" alt="postgresql-18-topn" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-topn` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "topn_18" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/topn_18-2.7.0-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "topn_17" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/topn_17-2.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "topn_16" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/topn_16-2.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "topn_15" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/topn_15-2.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "topn_14" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/topn_14-2.7.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "topn_18" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/topn_18-2.7.0-2PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "topn_17" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/topn_17-2.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "topn_16" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/topn_16-2.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "topn_15" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/topn_15-2.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "topn_14" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/topn_14-2.7.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "topn_18" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/topn_18-2.7.0-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "topn_17" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/topn_17-2.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "topn_16" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/topn_16-2.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "topn_15" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/topn_15-2.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "topn_14" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/topn_14-2.7.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "topn_18" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/topn_18-2.7.0-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "topn_17" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/topn_17-2.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "topn_16" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/topn_16-2.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "topn_15" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/topn_15-2.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "topn_14" "2.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/topn_14-2.7.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-topn" >}}     | {{< pkg "postgresql-17-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-topn" "2.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_18` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_18-2.7.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/topn_18-2.7.0-2PGDG.rhel8.aarch64.rpm) |
| `topn_18` | 2.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [topn_18-2.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/topn_18-2.7.0-2PGDG.rhel8.x86_64.rpm) |
| `topn_18` | 2.7.0 | `el9.x86_64` | pgdg | 23.3 KiB | [topn_18-2.7.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/topn_18-2.7.0-2PGDG.rhel9.x86_64.rpm) |
| `topn_18` | 2.7.0 | `el9.aarch64` | pgdg | 22.8 KiB | [topn_18-2.7.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/topn_18-2.7.0-2PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_17` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_17-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/topn_17-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_17` | 2.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [topn_17-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/topn_17-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_17` | 2.7.0 | `el9.x86_64` | pgdg | 23.4 KiB | [topn_17-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/topn_17-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `topn_17` | 2.7.0 | `el9.aarch64` | pgdg | 23.0 KiB | [topn_17-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/topn_17-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-topn` | 2.7.0 | `d12.aarch64` | pigsty | 31.6 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-topn` | 2.7.0 | `d12.x86_64` | pigsty | 32.0 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-topn` | 2.7.0 | `u22.aarch64` | pigsty | 32.8 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-topn` | 2.7.0 | `u22.x86_64` | pigsty | 33.0 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-topn` | 2.7.0 | `u24.aarch64` | pigsty | 29.6 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-topn` | 2.7.0 | `u24.x86_64` | pigsty | 29.6 KiB | [postgresql-17-topn_2.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-17-topn_2.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_16` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_16-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/topn_16-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_16` | 2.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [topn_16-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/topn_16-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_16` | 2.6.0 | `el8.x86_64` | pgdg | 23.7 KiB | [topn_16-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/topn_16-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_16` | 2.6.0 | `el8.aarch64` | pgdg | 23.5 KiB | [topn_16-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/topn_16-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_16` | 2.7.0 | `el9.x86_64` | pgdg | 23.4 KiB | [topn_16-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/topn_16-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `topn_16` | 2.7.0 | `el9.aarch64` | pgdg | 23.2 KiB | [topn_16-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/topn_16-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `topn_16` | 2.6.0 | `el9.aarch64` | pgdg | 22.7 KiB | [topn_16-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/topn_16-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `topn_16` | 2.6.0 | `el9.x86_64` | pgdg | 23.1 KiB | [topn_16-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/topn_16-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-topn` | 2.7.0 | `d12.x86_64` | pigsty | 31.8 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-topn` | 2.7.0 | `d12.aarch64` | pigsty | 31.6 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-topn` | 2.7.0 | `u22.aarch64` | pigsty | 32.7 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-topn` | 2.7.0 | `u22.x86_64` | pigsty | 32.9 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-topn` | 2.7.0 | `u24.x86_64` | pigsty | 29.5 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-topn` | 2.7.0 | `u24.aarch64` | pigsty | 29.6 KiB | [postgresql-16-topn_2.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-16-topn_2.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_15` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_15-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/topn_15-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_15` | 2.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [topn_15-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/topn_15-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_15` | 2.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [topn_15-2.4.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/topn_15-2.4.0-2.rhel8.aarch64.rpm) |
| `topn_15` | 2.4.0 | `el8.x86_64` | pgdg | 24.1 KiB | [topn_15-2.4.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/topn_15-2.4.0-2.rhel8.x86_64.rpm) |
| `topn_15` | 2.7.0 | `el9.x86_64` | pgdg | 23.4 KiB | [topn_15-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/topn_15-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `topn_15` | 2.7.0 | `el9.aarch64` | pgdg | 23.0 KiB | [topn_15-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/topn_15-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `topn_15` | 2.4.0 | `el9.aarch64` | pgdg | 23.5 KiB | [topn_15-2.4.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/topn_15-2.4.0-2.rhel9.aarch64.rpm) |
| `topn_15` | 2.4.0 | `el9.x86_64` | pgdg | 24.0 KiB | [topn_15-2.4.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/topn_15-2.4.0-2.rhel9.x86_64.rpm) |
| `postgresql-15-topn` | 2.7.0 | `d12.aarch64` | pigsty | 31.6 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-topn` | 2.7.0 | `d12.x86_64` | pigsty | 31.8 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-topn` | 2.7.0 | `u22.aarch64` | pigsty | 32.7 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-topn` | 2.7.0 | `u22.x86_64` | pigsty | 32.8 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-topn` | 2.7.0 | `u24.aarch64` | pigsty | 29.5 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-topn` | 2.7.0 | `u24.x86_64` | pigsty | 29.5 KiB | [postgresql-15-topn_2.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-15-topn_2.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_14` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_14-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/topn_14-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_14` | 2.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [topn_14-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/topn_14-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_14` | 2.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [topn_14-2.4.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/topn_14-2.4.0-2.rhel8.aarch64.rpm) |
| `topn_14` | 2.4.0 | `el8.x86_64` | pgdg | 43.0 KiB | [topn_14-2.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/topn_14-2.4.0-1.rhel8.x86_64.rpm) |
| `topn_14` | 2.7.0 | `el9.x86_64` | pgdg | 23.4 KiB | [topn_14-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/topn_14-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `topn_14` | 2.7.0 | `el9.aarch64` | pgdg | 23.0 KiB | [topn_14-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/topn_14-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `topn_14` | 2.4.0 | `el9.aarch64` | pgdg | 23.5 KiB | [topn_14-2.4.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/topn_14-2.4.0-2.rhel9.aarch64.rpm) |
| `postgresql-14-topn` | 2.7.0 | `d12.aarch64` | pigsty | 31.5 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-topn` | 2.7.0 | `d12.x86_64` | pigsty | 31.7 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-topn` | 2.7.0 | `u22.x86_64` | pigsty | 32.8 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-topn` | 2.7.0 | `u22.aarch64` | pigsty | 32.6 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-topn` | 2.7.0 | `u24.x86_64` | pigsty | 29.6 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-topn` | 2.7.0 | `u24.aarch64` | pigsty | 29.5 KiB | [postgresql-14-topn_2.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-14-topn_2.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `topn_13` | 2.7.0 | `el8.x86_64` | pgdg | 23.7 KiB | [topn_13-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/topn_13-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `topn_13` | 2.7.0 | `el8.aarch64` | pgdg | 23.8 KiB | [topn_13-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/topn_13-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `topn_13` | 2.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [topn_13-2.4.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/topn_13-2.4.0-2.rhel8.aarch64.rpm) |
| `topn_13` | 2.3.1 | `el8.x86_64` | pgdg | 41.5 KiB | [topn_13-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/topn_13-2.3.1-1.rhel8.x86_64.rpm) |
| `topn_13` | 2.7.0 | `el9.x86_64` | pgdg | 23.4 KiB | [topn_13-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/topn_13-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `topn_13` | 2.7.0 | `el9.aarch64` | pgdg | 23.2 KiB | [topn_13-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/topn_13-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `topn_13` | 2.4.0 | `el9.aarch64` | pgdg | 23.5 KiB | [topn_13-2.4.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/topn_13-2.4.0-2.rhel9.aarch64.rpm) |
| `postgresql-13-topn` | 2.7.0 | `d12.x86_64` | pigsty | 31.8 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-topn` | 2.7.0 | `d12.aarch64` | pigsty | 31.3 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-topn` | 2.7.0 | `u22.aarch64` | pigsty | 32.3 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-topn` | 2.7.0 | `u22.x86_64` | pigsty | 32.5 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-topn` | 2.7.0 | `u24.aarch64` | pigsty | 29.2 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-topn` | 2.7.0 | `u24.x86_64` | pigsty | 29.3 KiB | [postgresql-13-topn_2.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/topn/postgresql-13-topn_2.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/postgresql-topn" title="Repository" icon="github" subtitle="github.com/citusdata/postgresql-topn" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="postgresql-topn-2.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get topn; # get topn source code
pig build dep topn; # install build dependencies
pig build pkg topn; # build extension rpm or deb
pig build ext topn; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install topn; # install by extension name, for the current active PG version
pig ext install topn; # install via package alias, for the active PG version
pig ext install topn -v 17;   # install for PG 17
pig ext install topn -v 16;   # install for PG 16
pig ext install topn -v 15;   # install for PG 15
pig ext install topn -v 14;   # install for PG 14
pig ext install topn -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION topn;
```

