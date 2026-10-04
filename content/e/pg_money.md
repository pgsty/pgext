---
title: "pg_money"
linkTitle: "pg_money"
description: "Exact currency-aware monetary types, arithmetic, and allocation"
weight: 3690
categories: ["TYPE"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_money**](https://github.com/RustedBytes/pg-money) : Exact currency-aware monetary types, arithmetic, and allocation


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3690** | {{< badge content="pg_money" link="https://github.com/RustedBytes/pg-money" >}} | {{< ext "pg_money" >}} | `0.3.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "currency" >}} {{< ext "financial" >}} {{< ext "pg_rational" >}} {{< ext "pgmp" >}} |

> [!Note] Currency metadata is built in; the currency extension is related, not required.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_money` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_money_18" "green" >}} {{< bg "17" "pg_money_17" "green" >}} {{< bg "16" "pg_money_16" "green" >}} {{< bg "15" "pg_money_15" "green" >}} {{< bg "14" "pg_money_14" "green" >}} | `pg_money_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-money" "green" >}} {{< bg "17" "postgresql-17-pg-money" "green" >}} {{< bg "16" "postgresql-16-pg-money" "green" >}} {{< bg "15" "postgresql-15-pg-money" "green" >}} {{< bg "14" "postgresql-14-pg-money" "green" >}} | `postgresql-$v-pg-money` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_money_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-money : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-money : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_money_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1017.7 KiB | [pg_money_18-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_money_18-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_money_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 970.7 KiB | [pg_money_18-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_money_18-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_money_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_money_18-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_money_18-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_money_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_money_18-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_money_18-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_money_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_money_18-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_money_18-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_money_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1018.1 KiB | [pg_money_18-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_money_18-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-money` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 853.1 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 790.0 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 853.2 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 791.2 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 951.8 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 921.0 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 927.8 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 914.2 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 924.0 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-money` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 910.7 KiB | [postgresql-18-pg-money_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-18-pg-money_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_money_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1014.9 KiB | [pg_money_17-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_money_17-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_money_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 966.9 KiB | [pg_money_17-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_money_17-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_money_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_money_17-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_money_17-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_money_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_money_17-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_money_17-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_money_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_money_17-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_money_17-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_money_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1017.6 KiB | [pg_money_17-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_money_17-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-money` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 851.0 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 788.6 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 851.4 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 789.2 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 947.9 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 920.5 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 925.5 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 912.6 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 923.4 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-money` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 908.6 KiB | [postgresql-17-pg-money_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-17-pg-money_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_money_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1013.7 KiB | [pg_money_16-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_money_16-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_money_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 965.8 KiB | [pg_money_16-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_money_16-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_money_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_money_16-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_money_16-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_money_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_money_16-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_money_16-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_money_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_money_16-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_money_16-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_money_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1019.0 KiB | [pg_money_16-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_money_16-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-money` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 850.8 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 787.8 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 850.8 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 788.6 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 948.9 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 920.5 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 927.9 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 913.2 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 923.5 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-money` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 908.8 KiB | [postgresql-16-pg-money_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-16-pg-money_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_money_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1005.9 KiB | [pg_money_15-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_money_15-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_money_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 955.8 KiB | [pg_money_15-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_money_15-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_money_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1022.4 KiB | [pg_money_15-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_money_15-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_money_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_money_15-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_money_15-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_money_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1023.4 KiB | [pg_money_15-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_money_15-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_money_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1015.3 KiB | [pg_money_15-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_money_15-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-money` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 845.1 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 782.7 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 845.1 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 783.5 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 939.6 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 917.5 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 921.1 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 911.1 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 918.7 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-money` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 908.1 KiB | [postgresql-15-pg-money_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-15-pg-money_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_money_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1003.2 KiB | [pg_money_14-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_money_14-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_money_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 953.8 KiB | [pg_money_14-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_money_14-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_money_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1019.8 KiB | [pg_money_14-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_money_14-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_money_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1021.9 KiB | [pg_money_14-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_money_14-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_money_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1021.1 KiB | [pg_money_14-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_money_14-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_money_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1013.7 KiB | [pg_money_14-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_money_14-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-money` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 845.0 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 783.5 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 845.2 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 784.1 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 938.9 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 918.0 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 921.9 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 911.5 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 919.1 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-money` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 908.9 KiB | [postgresql-14-pg-money_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-money/postgresql-14-pg-money_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RustedBytes/pg-money" title="Repository" icon="github" subtitle="github.com/RustedBytes/pg-money" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_money-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_money;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_money;		# install via package name, for the active PG version

pig install pg_money -v 18;   # install for PG 18
pig install pg_money -v 17;   # install for PG 17
pig install pg_money -v 16;   # install for PG 16
pig install pg_money -v 15;   # install for PG 15
pig install pg_money -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_money;
```

## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/README.md)
- [Control file](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/pg_money.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/docs/SECURITY.md)

`pg_money` provides exact currency-aware amounts, arithmetic, rounding, and allocation on PostgreSQL 14–18. Its types are distinct from the locale-sensitive built-in money type.

### Core Workflow

```sql
CREATE EXTENSION pg_money;
CREATE TABLE invoices (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       total money_with_currency NOT NULL);
INSERT INTO invoices(total) VALUES ('USD 19.95'), (money_make(5.25, 'USD'));
SELECT sum(total), avg(total), money_format(sum(total)) FROM invoices;
SELECT money_split('USD 10.00', 3);
SELECT money_exchange('USD 100', 'EUR', 0.85);
```

### Objects and Semantics

`money_with_currency` stores a currency with an exact decimal amount; `money_minor` works with minor-unit amounts. `money_make`, `money_from_minor`, and `money_minor_make` construct values. `money_round` accepts an explicit rounding mode, and `money_split` allocates residual units while conserving the total.

Arithmetic and aggregation reject incompatible currencies instead of converting implicitly. Exchange rates are supplied explicitly or read from an application-owned table with `money_exchange_at`; the extension fetches no market rates. Check rounding and amount bounds at application boundaries.

### Operation

The 0.3.0 control file requires superuser installation and allows relocation. No preload is required. Treat exchange-rate tables and mutation privileges separately from read-only arithmetic; caller-supplied SQL data does not establish the correctness or freshness of a rate.
