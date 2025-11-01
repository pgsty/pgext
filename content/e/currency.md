---
title: "currency"
linkTitle: "currency"
description: "Custom PostgreSQL currency type in 1Byte"
weight: 3620
categories: ["TYPE"]
width: full
---

Custom PostgreSQL currency type in 1Byte


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3620** | {{< badge content="currency" link="https://github.com/adjust/pg-currency" >}} | {{< ext "currency" "pg_currency" >}} | `0.0.3` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "l10n_table_dependent_extension" >}} {{< ext "country" >}} {{< ext "pg_xenophile" >}} {{< ext "numeral" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/currency" >}} | `0.0.3` | {{< bg "18" "pg_currency_18*" "green" >}} {{< bg "17" "pg_currency_17*" "green" >}} {{< bg "16" "pg_currency_16*" "green" >}} {{< bg "15" "pg_currency_15*" "green" >}} {{< bg "14" "pg_currency_14*" "green" >}} {{< bg "13" "pg_currency_13*" "green" >}} | `pg_currency_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/currency" >}} | `0.0.3` | {{< bg "18" "postgresql-18-pg-currency" "green" >}} {{< bg "17" "postgresql-17-pg-currency" "green" >}} {{< bg "16" "postgresql-16-pg-currency" "green" >}} {{< bg "15" "postgresql-15-pg-currency" "green" >}} {{< bg "14" "postgresql-14-pg-currency" "green" >}} {{< bg "13" "postgresql-13-pg-currency" "green" >}} | `postgresql-$v-pg-currency` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-currency : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-currency : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-currency : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-currency : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-currency : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_18` | 0.0.3 | `el8.x86_64` | pigsty | 16.8 KiB | [pg_currency_18-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_18-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_18` | 0.0.3 | `el8.aarch64` | pigsty | 17.4 KiB | [pg_currency_18-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_18-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_18` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_currency_18-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_18-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_18` | 0.0.3 | `el9.aarch64` | pigsty | 17.6 KiB | [pg_currency_18-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_18-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_18` | 0.0.3 | `el10.x86_64` | pigsty | 17.0 KiB | [pg_currency_18-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_18-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_18` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_18-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_18-0.0.3-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_17` | 0.0.3 | `el8.x86_64` | pigsty | 16.8 KiB | [pg_currency_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_17` | 0.0.3 | `el8.aarch64` | pigsty | 17.4 KiB | [pg_currency_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_17` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_currency_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_17` | 0.0.3 | `el9.aarch64` | pigsty | 17.6 KiB | [pg_currency_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_17` | 0.0.3 | `el10.x86_64` | pigsty | 17.0 KiB | [pg_currency_17-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_17-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_17` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_17-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_17-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-currency` | 0.0.3 | `d12.x86_64` | pigsty | 21.1 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-currency` | 0.0.3 | `d12.aarch64` | pigsty | 21.2 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-currency` | 0.0.3 | `u22.x86_64` | pigsty | 22.5 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-currency` | 0.0.3 | `u22.aarch64` | pigsty | 22.4 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-currency` | 0.0.3 | `u24.x86_64` | pigsty | 20.4 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-currency` | 0.0.3 | `u24.aarch64` | pigsty | 20.8 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_16` | 0.0.3 | `el8.x86_64` | pigsty | 16.8 KiB | [pg_currency_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_16` | 0.0.3 | `el8.aarch64` | pigsty | 17.4 KiB | [pg_currency_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_16` | 0.0.3 | `el9.x86_64` | pigsty | 17.2 KiB | [pg_currency_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_16` | 0.0.3 | `el9.aarch64` | pigsty | 17.6 KiB | [pg_currency_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_16` | 0.0.3 | `el10.x86_64` | pigsty | 17.0 KiB | [pg_currency_16-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_16-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_16` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_16-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_16-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-currency` | 0.0.3 | `d12.x86_64` | pigsty | 21.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-currency` | 0.0.3 | `d12.aarch64` | pigsty | 21.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-currency` | 0.0.3 | `u22.x86_64` | pigsty | 22.4 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-currency` | 0.0.3 | `u22.aarch64` | pigsty | 22.4 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-currency` | 0.0.3 | `u24.x86_64` | pigsty | 20.4 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-currency` | 0.0.3 | `u24.aarch64` | pigsty | 20.8 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_15` | 0.0.3 | `el8.x86_64` | pigsty | 16.7 KiB | [pg_currency_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_15` | 0.0.3 | `el8.aarch64` | pigsty | 17.3 KiB | [pg_currency_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_15` | 0.0.3 | `el9.x86_64` | pigsty | 17.0 KiB | [pg_currency_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_15` | 0.0.3 | `el9.aarch64` | pigsty | 17.6 KiB | [pg_currency_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_15` | 0.0.3 | `el10.x86_64` | pigsty | 16.9 KiB | [pg_currency_15-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_15-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_15` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_15-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_15-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-currency` | 0.0.3 | `d12.x86_64` | pigsty | 21.1 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-currency` | 0.0.3 | `d12.aarch64` | pigsty | 21.2 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-currency` | 0.0.3 | `u22.x86_64` | pigsty | 22.4 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-currency` | 0.0.3 | `u22.aarch64` | pigsty | 22.4 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-currency` | 0.0.3 | `u24.x86_64` | pigsty | 20.3 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-currency` | 0.0.3 | `u24.aarch64` | pigsty | 20.8 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_14` | 0.0.3 | `el8.x86_64` | pigsty | 16.7 KiB | [pg_currency_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_14` | 0.0.3 | `el8.aarch64` | pigsty | 17.3 KiB | [pg_currency_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_14` | 0.0.3 | `el9.x86_64` | pigsty | 17.0 KiB | [pg_currency_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_14` | 0.0.3 | `el9.aarch64` | pigsty | 17.6 KiB | [pg_currency_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_14` | 0.0.3 | `el10.x86_64` | pigsty | 16.9 KiB | [pg_currency_14-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_14-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_14` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_14-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_14-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-currency` | 0.0.3 | `d12.x86_64` | pigsty | 20.9 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-currency` | 0.0.3 | `d12.aarch64` | pigsty | 21.0 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-currency` | 0.0.3 | `u22.x86_64` | pigsty | 22.3 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-currency` | 0.0.3 | `u22.aarch64` | pigsty | 22.3 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-currency` | 0.0.3 | `u24.x86_64` | pigsty | 20.2 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-currency` | 0.0.3 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_13` | 0.0.3 | `el8.x86_64` | pigsty | 16.7 KiB | [pg_currency_13-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_13-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_13` | 0.0.3 | `el8.aarch64` | pigsty | 17.4 KiB | [pg_currency_13-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_13-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_13` | 0.0.3 | `el9.x86_64` | pigsty | 16.9 KiB | [pg_currency_13-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_13-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_13` | 0.0.3 | `el9.aarch64` | pigsty | 17.5 KiB | [pg_currency_13-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_13-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_13` | 0.0.3 | `el10.x86_64` | pigsty | 16.8 KiB | [pg_currency_13-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_13-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_13` | 0.0.3 | `el10.aarch64` | pigsty | 17.7 KiB | [pg_currency_13-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_13-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-currency` | 0.0.3 | `d12.x86_64` | pigsty | 20.6 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-currency` | 0.0.3 | `d12.aarch64` | pigsty | 20.6 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-currency` | 0.0.3 | `u22.x86_64` | pigsty | 22.1 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-currency` | 0.0.3 | `u22.aarch64` | pigsty | 22.3 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-currency` | 0.0.3 | `u24.x86_64` | pigsty | 20.2 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-currency` | 0.0.3 | `u24.aarch64` | pigsty | 20.4 KiB | [postgresql-13-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-13-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-currency" title="Repository" icon="github" subtitle="github.com/adjust/pg-currency" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-currency-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get currency; # get currency source code
pig build dep currency; # install build dependencies
pig build pkg currency; # build extension rpm or deb
pig build ext currency; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install currency; # install by extension name, for the current active PG version
pig ext install pg_currency; # install via package alias, for the active PG version
pig ext install currency -v 18;   # install for PG 18
pig ext install currency -v 17;   # install for PG 17
pig ext install currency -v 16;   # install for PG 16
pig ext install currency -v 15;   # install for PG 15
pig ext install currency -v 14;   # install for PG 14
pig ext install currency -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION currency;
```

