---
title: "emailaddr"
linkTitle: "emailaddr"
description: "Email address type for PostgreSQL"
weight: 3800
categories: ["TYPE"]
languages: ["C"]
licenses: ["Unknown"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_emailaddr**](https://github.com/petere/pgemailaddr) : Email address type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3800** | {{< badge content="emailaddr" link="https://github.com/petere/pgemailaddr" >}} | {{< ext "emailaddr" "pg_emailaddr" >}} | `0` | {{< category "TYPE" >}} | {{< license "Unknown" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_html5_email_address" >}} {{< ext "uri" >}} {{< ext "pg_utl_smtp" >}} {{< ext "pg_smtp_client" >}} {{< ext "omni_email" >}} {{< ext "ip4r" >}} {{< ext "url_encode" >}} {{< ext "prefix" >}} |

> [!Note] +varatt.h, no valid license, distribution suspend


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_emailaddr` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0` | {{< bg "18" "pg_emailaddr_18" "green" >}} {{< bg "17" "pg_emailaddr_17" "green" >}} {{< bg "16" "pg_emailaddr_16" "green" >}} {{< bg "15" "pg_emailaddr_15" "green" >}} {{< bg "14" "pg_emailaddr_14" "green" >}} | `pg_emailaddr_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0` | {{< bg "18" "postgresql-18-pg-emailaddr" "green" >}} {{< bg "17" "postgresql-17-pg-emailaddr" "green" >}} {{< bg "16" "postgresql-16-pg-emailaddr" "green" >}} {{< bg "15" "postgresql-15-pg-emailaddr" "green" >}} {{< bg "14" "postgresql-14-pg-emailaddr" "green" >}} | `postgresql-$v-pg-emailaddr` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0" "postgresql-18-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_18` | `0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.9 KiB | [pg_emailaddr_18-0-3PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_18-0-3PGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_18` | `0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [pg_emailaddr_18-0-3PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_18-0-3PGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_18` | `0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_18-0-3PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_18-0-3PGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_18` | `0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pg_emailaddr_18-0-3PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_18-0-3PGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_18` | `0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_18-0-3PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_18-0-3PGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_18` | `0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_emailaddr_18-0-3PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_18-0-3PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-emailaddr` | `0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.8 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.1 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-emailaddr` | `0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.3 KiB | [postgresql-18-pg-emailaddr_0-3PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-18-pg-emailaddr_0-3PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_17` | `0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.9 KiB | [pg_emailaddr_17-0-3PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_17-0-3PGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_17` | `0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [pg_emailaddr_17-0-3PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_17-0-3PGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_17` | `0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_17-0-3PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_17-0-3PGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_17` | `0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pg_emailaddr_17-0-3PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_17-0-3PGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_17` | `0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_17-0-3PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_17-0-3PGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_17` | `0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_emailaddr_17-0-3PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_17-0-3PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-emailaddr` | `0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-emailaddr` | `0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.3 KiB | [postgresql-17-pg-emailaddr_0-3PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-3PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_16` | `0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.9 KiB | [pg_emailaddr_16-0-3PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_16-0-3PGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_16` | `0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [pg_emailaddr_16-0-3PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_16-0-3PGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_16` | `0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_16-0-3PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_16-0-3PGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_16` | `0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pg_emailaddr_16-0-3PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_16-0-3PGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_16` | `0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_16-0-3PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_16-0-3PGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_16` | `0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_emailaddr_16-0-3PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_16-0-3PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-emailaddr` | `0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-emailaddr` | `0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.3 KiB | [postgresql-16-pg-emailaddr_0-3PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-3PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_15` | `0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.9 KiB | [pg_emailaddr_15-0-3PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_15-0-3PGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_15` | `0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [pg_emailaddr_15-0-3PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_15-0-3PGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_15` | `0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_15-0-3PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_15-0-3PGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_15` | `0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pg_emailaddr_15-0-3PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_15-0-3PGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_15` | `0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_15-0-3PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_15-0-3PGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_15` | `0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_emailaddr_15-0-3PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_15-0-3PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-emailaddr` | `0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.1 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.1 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-emailaddr` | `0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.3 KiB | [postgresql-15-pg-emailaddr_0-3PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-3PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_14` | `0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.9 KiB | [pg_emailaddr_14-0-3PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_14-0-3PGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_14` | `0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.8 KiB | [pg_emailaddr_14-0-3PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_14-0-3PGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_14` | `0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_14-0-3PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_14-0-3PGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_14` | `0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pg_emailaddr_14-0-3PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_14-0-3PGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_14` | `0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_emailaddr_14-0-3PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_14-0-3PGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_14` | `0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_emailaddr_14-0-3PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_14-0-3PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-emailaddr` | `0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-emailaddr` | `0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.3 KiB | [postgresql-14-pg-emailaddr_0-3PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-3PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pgemailaddr" title="Repository" icon="github" subtitle="github.com/petere/pgemailaddr" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgemailaddr-0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_emailaddr;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_emailaddr;		# install via package name, for the active PG version
pig install emailaddr;		# install by extension name, for the current active PG version

pig install emailaddr -v 18;   # install for PG 18
pig install emailaddr -v 17;   # install for PG 17
pig install emailaddr -v 16;   # install for PG 16
pig install emailaddr -v 15;   # install for PG 15
pig install emailaddr -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION emailaddr;
```




## Usage

> [emailaddr: email address data type for PostgreSQL](https://github.com/petere/pgemailaddr)

The `emailaddr` extension provides a data type for storing and validating email addresses conforming to the `addr-spec` format defined in RFC 5322.

```sql
CREATE EXTENSION emailaddr;

CREATE TABLE accounts (
    id    int PRIMARY KEY,
    name  text,
    email emailaddr
);

INSERT INTO accounts VALUES (1, 'Peter Eisentraut', 'peter@eisentraut.org');
```

### Data Type

The `emailaddr` type validates email addresses on input according to RFC 5322 `addr-spec` rules. Simple formats like `user@domain.com` are accepted. Display name syntax such as `"User Name" <user@domain.com>` is not supported.

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Btree indexes are available for efficient lookups and sorting on `emailaddr` columns.
