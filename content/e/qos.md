---
title: "qos"
linkTitle: "qos"
description: "QoS resource governor extension for PostgreSQL sessions and queries"
weight: 5240
categories: ["ADMIN"]
languages: ["C"]
licenses: ["GPL-3.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_qos**](https://github.com/appstonia/pg_qos) : QoS resource governor extension for PostgreSQL sessions and queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5240** | {{< badge content="qos" link="https://github.com/appstonia/pg_qos" >}} | {{< ext "qos" "pg_qos" >}} | `1.1.0` | {{< category "ADMIN" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plan_filter" >}} {{< ext "pg_kpart" >}} {{< ext "pg_readonly" >}} {{< ext "prioritize" >}} {{< ext "block_copy_command" >}} {{< ext "safeupdate" >}} {{< ext "pg_command_fw" >}} {{< ext "pg_strict" >}} {{< ext "pg_hint_plan" >}} |

> [!Note] Upstream PG15-18; RPM also carries PG14. Requires preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_qos` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_qos_18" "green" >}} {{< bg "17" "pg_qos_17" "green" >}} {{< bg "16" "pg_qos_16" "green" >}} {{< bg "15" "pg_qos_15" "green" >}} {{< bg "14" "pg_qos_14" "green" >}} | `pg_qos_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-qos" "green" >}} {{< bg "17" "postgresql-17-qos" "green" >}} {{< bg "16" "postgresql-16-qos" "green" >}} {{< bg "15" "postgresql-15-qos" "green" >}} {{< bg "14" "postgresql-14-qos" "red" >}} | `postgresql-$v-qos` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_qos_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-qos : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pg_qos_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.3 KiB | [pg_qos_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.3 KiB | [pg_qos_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.7 KiB | [pg_qos_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.6 KiB | [pg_qos_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.3 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.5 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.6 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.6 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.7 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.1 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.7 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.4 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.8 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.4 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pg_qos_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.5 KiB | [pg_qos_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.5 KiB | [pg_qos_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_qos_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pg_qos_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.3 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.6 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.6 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.7 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.3 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 80.9 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.9 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.5 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 72.0 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.6 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.9 KiB | [pg_qos_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.4 KiB | [pg_qos_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.4 KiB | [pg_qos_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pg_qos_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pg_qos_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.2 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.3 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.5 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.4 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 79.9 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 79.5 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.8 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.3 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.9 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.4 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pg_qos_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.3 KiB | [pg_qos_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.2 KiB | [pg_qos_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [pg_qos_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.6 KiB | [pg_qos_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.5 KiB | [pg_qos_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.5 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 80.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 80.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.9 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.8 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.5 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/appstonia/pg_qos" title="Repository" icon="github" subtitle="github.com/appstonia/pg_qos" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_qos-1.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_qos;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_qos;		# install via package name, for the active PG version
pig install qos;		# install by extension name, for the current active PG version

pig install qos -v 18;   # install for PG 18
pig install qos -v 17;   # install for PG 17
pig install qos -v 16;   # install for PG 16
pig install qos -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'qos';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION qos;
```

## Usage

Sources:

- [README.md](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/README.md)
- [qos.control](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos.control)
- [qos--1.0--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.0--1.1.sql)
- [qos--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.1.sql)

`qos` 1.1 (distribution 1.1.0) applies per-role and per-database resource limits on PostgreSQL 15+. Merge `qos` into `shared_preload_libraries` and restart before creating its SQL objects as an administrator. CPU affinity limits require Linux.

### Configure Limits

```sql
CREATE EXTENSION qos;
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
ALTER ROLE app_user SET qos.max_select_rate = '10/500ms';
SELECT * FROM qos_stat_rate;
```

### Limit Semantics

`qos.work_mem_limit` caps effective work memory; `qos.cpu_core_limit` controls CPU affinity on Linux and limits parallel workers on other platforms. `qos.max_concurrent_tx`, `qos.max_concurrent_select`, `qos.max_concurrent_update`, `qos.max_concurrent_delete` and `qos.max_concurrent_insert` cap concurrent operations.

`qos.max_tx_rate`, `qos.max_select_rate`, `qos.max_update_rate`, `qos.max_delete_rate` and `qos.max_insert_rate` use count/window pairs such as 100/1s. The default -1 disables each rate limit. Windows range from 100 ms to one day. Rate and concurrency violations raise SQLSTATE 54000; clients should use the retry hint. The most restrictive applicable role/database setting wins, and rate pairs are compared by normalized rate.

### Observability and Upgrade

`qos_stat_rate` exposes live windows; the other `qos_stat` views expose activity and counters. `qos_prometheus_metrics()` renders Prometheus exposition text. Counters reset at server restart. Version 1.1 replaces the old nonfunctional `qos_get_stats()` with these views.

Upgrading requires replacing the library and restarting PostgreSQL because the shared-memory layout changes, followed by `ALTER EXTENSION qos UPDATE TO '1.1'` in each database. New rate limits stay disabled until configured. These controls do not replace application admission limits or operating-system isolation.
