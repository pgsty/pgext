---
title: "currency"
linkTitle: "currency"
description: "Custom PostgreSQL currency type in 1Byte"
weight: 3680
categories: ["TYPE"]
width: full
---

[**pg_currency**](https://github.com/adjust/pg-currency) : Custom PostgreSQL currency type in 1Byte


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3680** | {{< badge content="currency" link="https://github.com/adjust/pg-currency" >}} | {{< ext "currency" "pg_currency" >}} | `0.0.3` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "l10n_table_dependent_extension" >}} {{< ext "country" >}} {{< ext "pg_xenophile" >}} {{< ext "numeral" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_currency` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "pg_currency_18" "green" >}} {{< bg "17" "pg_currency_17" "green" >}} {{< bg "16" "pg_currency_16" "green" >}} {{< bg "15" "pg_currency_15" "green" >}} {{< bg "14" "pg_currency_14" "green" >}} | `pg_currency_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "postgresql-18-pg-currency" "green" >}} {{< bg "17" "postgresql-17-pg-currency" "green" >}} {{< bg "16" "postgresql-16-pg-currency" "green" >}} {{< bg "15" "postgresql-15-pg-currency" "green" >}} {{< bg "14" "postgresql-14-pg-currency" "green" >}} | `postgresql-$v-pg-currency` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_currency_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-currency : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-currency : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_18` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_currency_18-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_18-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_18` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_currency_18-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_18-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_18` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_currency_18-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_18-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_18` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.6 KiB | [pg_currency_18-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_18-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_18` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_currency_18-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_18-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_18` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.7 KiB | [pg_currency_18-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_18-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-currency` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.1 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.0 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.3 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.3 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.3 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.7 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-currency` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.9 KiB | [postgresql-18-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-18-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_17` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_currency_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_17` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_currency_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_17` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_currency_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_17` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.6 KiB | [pg_currency_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_17` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_currency_17-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_17-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_17` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.7 KiB | [pg_currency_17-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_17-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-currency` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.1 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.0 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.1 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.5 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.3 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.7 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-currency` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.8 KiB | [postgresql-17-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-17-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_16` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_currency_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_16` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_currency_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_16` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_currency_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_16` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.6 KiB | [pg_currency_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_16` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_currency_16-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_16-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_16` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.7 KiB | [pg_currency_16-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_16-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-currency` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.0 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.3 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.7 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-currency` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.8 KiB | [postgresql-16-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-16-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_15` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pg_currency_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_15` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [pg_currency_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_15` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [pg_currency_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_15` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.6 KiB | [pg_currency_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_15` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.9 KiB | [pg_currency_15-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_15-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_15` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.7 KiB | [pg_currency_15-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_15-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-currency` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.3 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.8 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-currency` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.8 KiB | [postgresql-15-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-15-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_currency_14` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pg_currency_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_currency_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_currency_14` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [pg_currency_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_currency_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_currency_14` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [pg_currency_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_currency_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_currency_14` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.6 KiB | [pg_currency_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_currency_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_currency_14` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.9 KiB | [pg_currency_14-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_currency_14-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_currency_14` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.7 KiB | [pg_currency_14-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_currency_14-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-currency` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.2 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.3 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.3 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.3 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.8 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-currency` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.8 KiB | [postgresql-14-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-currency/postgresql-14-pg-currency_0.0.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-currency" title="Repository" icon="github" subtitle="github.com/adjust/pg-currency" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-currency-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_currency;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_currency;		# install via package name, for the active PG version
pig install currency;		# install by extension name, for the current active PG version

pig install currency -v 18;   # install for PG 18
pig install currency -v 17;   # install for PG 17
pig install currency -v 16;   # install for PG 16
pig install currency -v 15;   # install for PG 15
pig install currency -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION currency CASCADE; -- requires plpgsql
```



## Usage

> [currency: ISO 4217 currency code type](https://github.com/adjust/pg-currency)

The `currency` extension provides a data type for ISO 4217 currency codes using only a single byte of storage per value.

```sql
CREATE EXTENSION currency;

CREATE TABLE transactions (
    id                serial,
    payment_currency  currency
);

INSERT INTO transactions VALUES (1, 'USD'), (2, 'EUR'), (3, 'USD');

SELECT * FROM transactions ORDER BY payment_currency;
 id | payment_currency
----+------------------
  2 | EUR
  1 | USD
  3 | USD
```

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

B-tree index support is included for efficient ordering and lookups.

### Functions

```sql
-- List all supported currency codes
SELECT * FROM supported_currencies() currency ORDER BY currency;
```
