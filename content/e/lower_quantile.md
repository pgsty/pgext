---
title: "lower_quantile"
linkTitle: "lower_quantile"
description: "Lower quantile aggregate function"
weight: 4620
categories: ["Func"]
width: full
---

Lower quantile aggregate function

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4620** | {{< badge content="lower_quantile" link="https://github.com/tvondra/lower_quantile" >}} | {{< ext "lower_quantile" "lower_quantile" >}} | `1.0.3` | {{< category "FUNC" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "quantile" >}} {{< ext "topn" >}} {{< ext "ddsketch" >}} {{< ext "omnisketch" >}} {{< ext "count_distinct" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/lower_quantile" >}} | `1.0.3` | {{< badge content="18" color="red" alt="lower_quantile_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `lower_quantile_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/lower_quantile" >}} | `1.0.3` | {{< badge content="18" color="red" alt="postgresql-18-lower-quantile" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-lower-quantile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "lower_quantile_18" >}}     | {{< pkg "lower_quantile_17" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "lower_quantile_16" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "lower_quantile_15" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "lower_quantile_14" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "lower_quantile_18" >}}     | {{< pkg "lower_quantile_17" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "lower_quantile_16" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "lower_quantile_15" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "lower_quantile_14" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "lower_quantile_18" >}}     | {{< pkg "lower_quantile_17" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "lower_quantile_16" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "lower_quantile_15" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "lower_quantile_14" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "lower_quantile_18" >}}     | {{< pkg "lower_quantile_17" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "lower_quantile_16" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "lower_quantile_15" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "lower_quantile_14" "1.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-lower-quantile" >}}     | {{< pkg "postgresql-17-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-lower-quantile" "1.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `lower_quantile_17` | 1.0.3 | `el8.x86_64` | pigsty | 15.4 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_17` | 1.0.3 | `el8.aarch64` | pigsty | 15.2 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_17` | 1.0.3 | `el9.aarch64` | pigsty | 15.2 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_17-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_17` | 1.0.3 | `el9.x86_64` | pigsty | 15.5 KiB | [lower_quantile_17-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_17-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-lower-quantile` | 1.0.3 | `d12.x86_64` | pigsty | 17.4 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-lower-quantile` | 1.0.3 | `d12.aarch64` | pigsty | 17.2 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-lower-quantile` | 1.0.3 | `u22.x86_64` | pigsty | 17.7 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-lower-quantile` | 1.0.3 | `u22.aarch64` | pigsty | 17.4 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-lower-quantile` | 1.0.3 | `u24.x86_64` | pigsty | 17.1 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-lower-quantile` | 1.0.3 | `u24.aarch64` | pigsty | 16.8 KiB | [postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-17-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `lower_quantile_16` | 1.0.3 | `el8.x86_64` | pigsty | 15.4 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_16` | 1.0.3 | `el8.aarch64` | pigsty | 15.2 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_16` | 1.0.3 | `el9.x86_64` | pigsty | 15.5 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_16-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_16` | 1.0.3 | `el9.aarch64` | pigsty | 15.2 KiB | [lower_quantile_16-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_16-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-lower-quantile` | 1.0.3 | `d12.x86_64` | pigsty | 17.2 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-lower-quantile` | 1.0.3 | `d12.aarch64` | pigsty | 17.0 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-lower-quantile` | 1.0.3 | `u22.aarch64` | pigsty | 17.1 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-lower-quantile` | 1.0.3 | `u22.x86_64` | pigsty | 17.5 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-lower-quantile` | 1.0.3 | `u24.x86_64` | pigsty | 16.9 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-lower-quantile` | 1.0.3 | `u24.aarch64` | pigsty | 16.6 KiB | [postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-16-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `lower_quantile_15` | 1.0.3 | `el8.x86_64` | pigsty | 15.4 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_15` | 1.0.3 | `el8.aarch64` | pigsty | 15.2 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_15` | 1.0.3 | `el9.x86_64` | pigsty | 15.5 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_15-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_15` | 1.0.3 | `el9.aarch64` | pigsty | 15.2 KiB | [lower_quantile_15-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_15-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-lower-quantile` | 1.0.3 | `d12.aarch64` | pigsty | 16.9 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-lower-quantile` | 1.0.3 | `d12.x86_64` | pigsty | 17.1 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-lower-quantile` | 1.0.3 | `u22.aarch64` | pigsty | 17.0 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-lower-quantile` | 1.0.3 | `u22.x86_64` | pigsty | 17.4 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-lower-quantile` | 1.0.3 | `u24.x86_64` | pigsty | 16.9 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-lower-quantile` | 1.0.3 | `u24.aarch64` | pigsty | 16.5 KiB | [postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-15-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `lower_quantile_14` | 1.0.3 | `el8.x86_64` | pigsty | 15.4 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_14` | 1.0.3 | `el8.aarch64` | pigsty | 15.1 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_14` | 1.0.3 | `el9.x86_64` | pigsty | 15.5 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_14-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `lower_quantile_14` | 1.0.3 | `el9.aarch64` | pigsty | 15.2 KiB | [lower_quantile_14-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_14-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-lower-quantile` | 1.0.3 | `d12.x86_64` | pigsty | 17.1 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-lower-quantile` | 1.0.3 | `d12.aarch64` | pigsty | 16.9 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-lower-quantile` | 1.0.3 | `u22.x86_64` | pigsty | 17.4 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-lower-quantile` | 1.0.3 | `u22.aarch64` | pigsty | 17.0 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-lower-quantile` | 1.0.3 | `u24.x86_64` | pigsty | 16.9 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-lower-quantile` | 1.0.3 | `u24.aarch64` | pigsty | 16.5 KiB | [postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-14-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `lower_quantile_13` | 1.0.3 | `el8.aarch64` | pigsty | 15.1 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lower_quantile_13-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `lower_quantile_13` | 1.0.3 | `el8.x86_64` | pigsty | 15.3 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lower_quantile_13-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `lower_quantile_13` | 1.0.3 | `el9.aarch64` | pigsty | 15.2 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lower_quantile_13-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `lower_quantile_13` | 1.0.3 | `el9.x86_64` | pigsty | 15.5 KiB | [lower_quantile_13-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lower_quantile_13-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-lower-quantile` | 1.0.3 | `d12.aarch64` | pigsty | 16.6 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-lower-quantile` | 1.0.3 | `d12.x86_64` | pigsty | 17.0 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-lower-quantile` | 1.0.3 | `u22.aarch64` | pigsty | 16.9 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-lower-quantile` | 1.0.3 | `u22.x86_64` | pigsty | 17.2 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-lower-quantile` | 1.0.3 | `u24.aarch64` | pigsty | 16.2 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-lower-quantile` | 1.0.3 | `u24.x86_64` | pigsty | 16.8 KiB | [postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lower-quantile/postgresql-13-lower-quantile_1.0.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/lower_quantile" title="Repository" icon="github" subtitle="github.com/tvondra/lower_quantile" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="lower_quantile-1.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get lower_quantile; # get lower_quantile source code
pig build dep lower_quantile; # install build dependencies
pig build pkg lower_quantile; # build extension rpm or deb
pig build ext lower_quantile; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install lower_quantile; # install by extension name, for the current active PG version
pig ext install lower_quantile; # install via package alias, for the active PG version
pig ext install lower_quantile -v 18;   # install for PG 18
pig ext install lower_quantile -v 17;   # install for PG 17
pig ext install lower_quantile -v 16;   # install for PG 16
pig ext install lower_quantile -v 15;   # install for PG 15
pig ext install lower_quantile -v 14;   # install for PG 14
pig ext install lower_quantile -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION lower_quantile;
```

