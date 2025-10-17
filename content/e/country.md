---
title: "country"
linkTitle: "country"
description: "Country data type, ISO 3166-1"
weight: 3600
categories: ["Type"]
width: full
---

Country data type, ISO 3166-1

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3600** | {{< badge content="country" link="https://github.com/adjust/pg-country" >}} | {{< ext "country" "pg_country" >}} | `0.0.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "l10n_table_dependent_extension" >}} {{< ext "pg_xenophile" >}} {{< ext "currency" >}} {{< ext "geoip" >}} {{< ext "icu_ext" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/country" >}} | `0.0.3` | {{< badge content="18" color="red" alt="pg_country_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_country_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/country" >}} | `0.0.3` | {{< badge content="18" color="red" alt="postgresql-18-pg-country" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-country` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_country_18" >}}     | {{< pkg "pg_country_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_17-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_country_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_16-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_country_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_15-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_country_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_14-0.0.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_country_18" >}}     | {{< pkg "pg_country_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_17-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_country_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_16-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_country_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_15-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_country_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_14-0.0.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_country_18" >}}     | {{< pkg "pg_country_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_17-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_country_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_16-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_country_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_15-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_country_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_14-0.0.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_country_18" >}}     | {{< pkg "pg_country_17" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_17-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_country_16" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_16-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_country_15" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_15-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_country_14" "0.0.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_14-0.0.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-country" >}}     | {{< pkg "postgresql-17-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-country" "0.0.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_country_17` | 0.0.3 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_country_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_country_17` | 0.0.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_country_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_country_17` | 0.0.3 | `el9.aarch64` | pigsty | 17.7 KiB | [pg_country_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_country_17` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_country_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-country` | 0.0.3 | `d12.x86_64` | pigsty | 28.4 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-country` | 0.0.3 | `d12.aarch64` | pigsty | 28.8 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-country` | 0.0.3 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-country` | 0.0.3 | `u22.aarch64` | pigsty | 32.9 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-country` | 0.0.3 | `u24.x86_64` | pigsty | 30.4 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-country` | 0.0.3 | `u24.aarch64` | pigsty | 31.1 KiB | [postgresql-17-pg-country_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-17-pg-country_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_country_16` | 0.0.3 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_country_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_country_16` | 0.0.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_country_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_country_16` | 0.0.3 | `el9.x86_64` | pigsty | 17.4 KiB | [pg_country_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_country_16` | 0.0.3 | `el9.aarch64` | pigsty | 17.7 KiB | [pg_country_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-country` | 0.0.3 | `d12.x86_64` | pigsty | 28.3 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-country` | 0.0.3 | `d12.aarch64` | pigsty | 28.9 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-country` | 0.0.3 | `u22.aarch64` | pigsty | 32.9 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-country` | 0.0.3 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-country` | 0.0.3 | `u24.x86_64` | pigsty | 30.4 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-country` | 0.0.3 | `u24.aarch64` | pigsty | 31.1 KiB | [postgresql-16-pg-country_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-16-pg-country_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_country_15` | 0.0.3 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_country_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_country_15` | 0.0.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_country_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_country_15` | 0.0.3 | `el9.x86_64` | pigsty | 17.4 KiB | [pg_country_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_country_15` | 0.0.3 | `el9.aarch64` | pigsty | 17.7 KiB | [pg_country_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-country` | 0.0.3 | `d12.aarch64` | pigsty | 28.9 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-country` | 0.0.3 | `d12.x86_64` | pigsty | 28.3 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-country` | 0.0.3 | `u22.aarch64` | pigsty | 32.9 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-country` | 0.0.3 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-country` | 0.0.3 | `u24.x86_64` | pigsty | 30.4 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-country` | 0.0.3 | `u24.aarch64` | pigsty | 31.1 KiB | [postgresql-15-pg-country_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-15-pg-country_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_country_14` | 0.0.3 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_country_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_country_14` | 0.0.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_country_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_country_14` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_country_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_country_14` | 0.0.3 | `el9.aarch64` | pigsty | 17.7 KiB | [pg_country_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-country` | 0.0.3 | `d12.x86_64` | pigsty | 28.2 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-country` | 0.0.3 | `d12.aarch64` | pigsty | 28.8 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-country` | 0.0.3 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-country` | 0.0.3 | `u22.aarch64` | pigsty | 32.8 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-country` | 0.0.3 | `u24.x86_64` | pigsty | 30.4 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-country` | 0.0.3 | `u24.aarch64` | pigsty | 31.1 KiB | [postgresql-14-pg-country_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-14-pg-country_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_country_13` | 0.0.3 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_country_13-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_country_13-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_country_13` | 0.0.3 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_country_13-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_country_13-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_country_13` | 0.0.3 | `el9.aarch64` | pigsty | 17.7 KiB | [pg_country_13-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_country_13-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_country_13` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_country_13-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_country_13-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-country` | 0.0.3 | `d12.aarch64` | pigsty | 28.4 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-country` | 0.0.3 | `d12.x86_64` | pigsty | 28.4 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-country` | 0.0.3 | `u22.aarch64` | pigsty | 32.5 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-country` | 0.0.3 | `u22.x86_64` | pigsty | 32.2 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-country` | 0.0.3 | `u24.aarch64` | pigsty | 30.8 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-country` | 0.0.3 | `u24.x86_64` | pigsty | 30.6 KiB | [postgresql-13-pg-country_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-country/postgresql-13-pg-country_0.0.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-country" title="Repository" icon="github" subtitle="github.com/adjust/pg-country" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg-country-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get country; # get country source code
pig build dep country; # install build dependencies
pig build pkg country; # build extension rpm or deb
pig build ext country; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install country; # install by extension name, for the current active PG version
pig ext install pg_country; # install via package alias, for the active PG version
pig ext install country -v 18;   # install for PG 18
pig ext install country -v 17;   # install for PG 17
pig ext install country -v 16;   # install for PG 16
pig ext install country -v 15;   # install for PG 15
pig ext install country -v 14;   # install for PG 14
pig ext install country -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION country;
```

