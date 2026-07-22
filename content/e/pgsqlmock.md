---
title: "pgsqlmock"
linkTitle: "pgsqlmock"
description: "Mocking and faking helpers for PostgreSQL unit tests"
weight: 3130
categories: ["LANG"]
width: full
---

[**pgsqlmock**](https://github.com/v-maliutin/pgSQLMock) : Mocking and faking helpers for PostgreSQL unit tests


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3130** | {{< badge content="pgsqlmock" link="https://github.com/v-maliutin/pgSQLMock" >}} | {{< ext "pgsqlmock" >}} | `1.0.1` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "pgtap" >}} |
|   **See Also**    | {{< ext "pgtap" >}} {{< ext "pg_mockable" >}} {{< ext "faker" >}} {{< ext "unit" >}} |

> [!Note] Packaging corrects the upstream control dependency name from pgTap to pgtap and requires pgTAP 1.3.4 or newer.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgsqlmock` | `plpgsql`, `pgtap` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "pgsqlmock_18" "green" >}} {{< bg "17" "pgsqlmock_17" "green" >}} {{< bg "16" "pgsqlmock_16" "green" >}} {{< bg "15" "pgsqlmock_15" "green" >}} {{< bg "14" "pgsqlmock_14" "green" >}} | `pgsqlmock_$v` | `pgtap_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-pgsqlmock" "green" >}} {{< bg "17" "postgresql-17-pgsqlmock" "green" >}} {{< bg "16" "postgresql-16-pgsqlmock" "green" >}} {{< bg "15" "postgresql-15-pgsqlmock" "green" >}} {{< bg "14" "postgresql-14-pgsqlmock" "green" >}} | `postgresql-$v-pgsqlmock` | `postgresql-$v-pgtap` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pgsqlmock_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pgsqlmock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pgsqlmock : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsqlmock_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsqlmock_18-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.7 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsqlmock_18-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsqlmock_18-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsqlmock_18-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsqlmock_18-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `pgsqlmock_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgsqlmock_18-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsqlmock_18-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pgsqlmock` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-18-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsqlmock_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsqlmock_17-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.7 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsqlmock_17-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsqlmock_17-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsqlmock_17-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsqlmock_17-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `pgsqlmock_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgsqlmock_17-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsqlmock_17-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pgsqlmock` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-17-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsqlmock_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsqlmock_16-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.7 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsqlmock_16-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsqlmock_16-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsqlmock_16-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsqlmock_16-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `pgsqlmock_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgsqlmock_16-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsqlmock_16-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pgsqlmock` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-16-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsqlmock_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsqlmock_15-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.7 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsqlmock_15-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsqlmock_15-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsqlmock_15-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsqlmock_15-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `pgsqlmock_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgsqlmock_15-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsqlmock_15-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pgsqlmock` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-15-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-15-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsqlmock_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsqlmock_14-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.7 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsqlmock_14-1.0.1-1PIGSTY.el8.noarch.rpm) |
| `pgsqlmock_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsqlmock_14-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsqlmock_14-1.0.1-1PIGSTY.el9.noarch.rpm) |
| `pgsqlmock_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsqlmock_14-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `pgsqlmock_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgsqlmock_14-1.0.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsqlmock_14-1.0.1-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.1 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.1 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pgsqlmock` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsqlmock/postgresql-14-pgsqlmock_1.0.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/v-maliutin/pgSQLMock" title="Repository" icon="github" subtitle="github.com/v-maliutin/pgSQLMock" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsqlmock-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgsqlmock;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgsqlmock;		# install via package name, for the active PG version

pig install pgsqlmock -v 18;   # install for PG 18
pig install pgsqlmock -v 17;   # install for PG 17
pig install pgsqlmock -v 16;   # install for PG 16
pig install pgsqlmock -v 15;   # install for PG 15
pig install pgsqlmock -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgsqlmock CASCADE; -- requires plpgsql, pgtap
```

## Usage

Sources:

- [pgSQLMock 1.0.1 documentation](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/doc/pgsqlmock.md)
- [pgSQLMock README](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/README.md)
- [pgSQLMock control file](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/pgsqlmock.control)
- [pgSQLMock 1.0.1 release](https://github.com/v-maliutin/pgSQLMock/releases/tag/Release_v1.0.1)

`pgsqlmock` extends pgTAP with table fakes, function and view mocks, call-count assertions, and debugging helpers. Its helpers alter or replace real database objects, so upstream requires using them inside pgTAP's transaction-based test context, where the changes are rolled back after the test.

```sql
CREATE EXTENSION pgtap;
CREATE EXTENSION pgsqlmock;
```

### Fake Tables

`fake_table(text[], ...)` can isolate a test from foreign keys, primary keys, `NOT NULL` constraints, partitions, or pre-existing rows. Pass schema-qualified table names as a `text[]`:

```sql
SELECT plan(2);

SELECT fake_table(
  _table_ident       => ARRAY['app.accounts', 'app.transactions'],
  _make_table_empty  => true,
  _leave_primary_key => false,
  _drop_not_null     => true
);

INSERT INTO app.transactions(account_id, amount)
VALUES (999, 42.00);

SELECT is(
  (SELECT sum(amount) FROM app.transactions WHERE account_id = 999),
  42.00::numeric,
  'transaction logic is isolated from account fixtures'
);

SELECT * FROM finish();
```

Important options include `make_table_empty`, `leave_primary_key`, `drop_not_null`, `drop_collation`, and `drop_partitions`. Keeping a primary key while dropping the participating columns' `NOT NULL` constraints is contradictory; remove or recreate the key explicitly for that test shape.

### Mock Functions

`mock_func(schema, name, signature, ...)` temporarily replaces a routine while preserving its identity. Supply either a scalar value or SQL/prepared-statement text for a set result:

```sql
CREATE OR REPLACE FUNCTION app.current_business_time()
RETURNS time LANGUAGE sql AS $$ SELECT current_time $$;

SELECT mock_func(
  'app',
  'current_business_time',
  '()',
  _return_scalar_value => '13:00'::time
);

SELECT is(app.current_business_time(), '13:00'::time, 'clock is deterministic');
```

For set-returning routines, pass `_return_set_value` as a SQL query or the name of a prepared statement. Use `get_routine_signature()` when overloaded or defaulted arguments make the stored signature unclear.

### Mock Views

`mock_view(schema, view_name, return_set_sql)` replaces a view with controlled rows:

```sql
SELECT mock_view(
  'app',
  'active_accounts',
  $$SELECT * FROM (VALUES (1, 'test')) AS v(id, name)$$
);

SELECT results_eq(
  'SELECT id, name FROM app.active_accounts',
  $$VALUES (1, 'test')$$,
  'view consumer sees only the fixture'
);
```

### Call Counts and Diagnostics

Set `track_functions = 'all'` before using `call_count()` to assert how often a routine was invoked:

```sql
SET LOCAL track_functions = 'all';

SELECT call_count(
  1,
  'app',
  'current_business_time',
  '()'
);
```

`print_table_as_json()` and `print_query_as_json()` emit reproducible SQL/JSON-style snapshots through `NOTICE`, which is useful when pgTAP's rollback would otherwise hide the state created during a failed test.

### Caveats

- Run mocks and fakes only inside isolated test transactions; they issue real `ALTER`, `DROP`, and replacement DDL.
- pgSQLMock depends on PL/pgSQL and pgTAP. Load pgTAP before running its assertions.
- `call_count()` depends on PostgreSQL function statistics and therefore requires `track_functions = 'all'`.
- Release 1.0.1 fixes `fake_table()` dropping `NOT NULL` constraints on tables without a primary key.
