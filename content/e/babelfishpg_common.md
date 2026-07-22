---
title: "babelfishpg_common"
linkTitle: "babelfishpg_common"
description: "SQL Server Transact SQL Datatype Support"
weight: 9300
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server Transact SQL Datatype Support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9300** | {{< badge content="babelfishpg_common" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_common" "babelfish" >}} | `5.4.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "tds_fdw" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "uuid-ossp" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |
|    **Siblings**   | {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `babelfish` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `libantlr4-runtime413` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `babelfish-18` | `6.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 MiB | [babelfish-18-6.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/babelfish-18-6.0.0-1PIGSTY.el8.x86_64.rpm) |
| `babelfish-18` | `6.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.3 MiB | [babelfish-18-6.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/babelfish-18-6.0.0-1PIGSTY.el8.aarch64.rpm) |
| `babelfish-18` | `6.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.9 MiB | [babelfish-18-6.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/babelfish-18-6.0.0-1PIGSTY.el9.x86_64.rpm) |
| `babelfish-18` | `6.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.7 MiB | [babelfish-18-6.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/babelfish-18-6.0.0-1PIGSTY.el9.aarch64.rpm) |
| `babelfish-18` | `6.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.2 MiB | [babelfish-18-6.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/babelfish-18-6.0.0-1PIGSTY.el10.x86_64.rpm) |
| `babelfish-18` | `6.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 MiB | [babelfish-18-6.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/babelfish-18-6.0.0-1PIGSTY.el10.aarch64.rpm) |
| `babelfish-18` | `6.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 MiB | [babelfish-18_6.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~bookworm_amd64.deb) |
| `babelfish-18` | `6.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [babelfish-18_6.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~bookworm_arm64.deb) |
| `babelfish-18` | `6.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.2 MiB | [babelfish-18_6.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~trixie_amd64.deb) |
| `babelfish-18` | `6.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.7 MiB | [babelfish-18_6.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~trixie_arm64.deb) |
| `babelfish-18` | `6.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.7 MiB | [babelfish-18_6.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~jammy_amd64.deb) |
| `babelfish-18` | `6.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.5 MiB | [babelfish-18_6.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~jammy_arm64.deb) |
| `babelfish-18` | `6.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.6 MiB | [babelfish-18_6.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~noble_amd64.deb) |
| `babelfish-18` | `6.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.4 MiB | [babelfish-18_6.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~noble_arm64.deb) |
| `babelfish-18` | `6.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.6 MiB | [babelfish-18_6.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~resolute_amd64.deb) |
| `babelfish-18` | `6.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.2 MiB | [babelfish-18_6.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/babelfish-18/babelfish-18_6.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `babelfish-17` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.3 MiB | [babelfish-17-5.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/babelfish-17-5.4.0-1PIGSTY.el8.x86_64.rpm) |
| `babelfish-17` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.8 MiB | [babelfish-17-5.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/babelfish-17-5.4.0-1PIGSTY.el8.aarch64.rpm) |
| `babelfish-17` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 MiB | [babelfish-17-5.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/babelfish-17-5.4.0-1PIGSTY.el9.x86_64.rpm) |
| `babelfish-17` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 MiB | [babelfish-17-5.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/babelfish-17-5.4.0-1PIGSTY.el9.aarch64.rpm) |
| `babelfish-17` | `5.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 MiB | [babelfish-17-5.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/babelfish-17-5.4.0-1PIGSTY.el10.x86_64.rpm) |
| `babelfish-17` | `5.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 MiB | [babelfish-17-5.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/babelfish-17-5.4.0-1PIGSTY.el10.aarch64.rpm) |
| `babelfish-17` | `5.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 MiB | [babelfish-17_5.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~bookworm_amd64.deb) |
| `babelfish-17` | `5.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.2 MiB | [babelfish-17_5.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~bookworm_arm64.deb) |
| `babelfish-17` | `5.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 MiB | [babelfish-17_5.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~trixie_amd64.deb) |
| `babelfish-17` | `5.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.4 MiB | [babelfish-17_5.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~trixie_arm64.deb) |
| `babelfish-17` | `5.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.3 MiB | [babelfish-17_5.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~jammy_amd64.deb) |
| `babelfish-17` | `5.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 MiB | [babelfish-17_5.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~jammy_arm64.deb) |
| `babelfish-17` | `5.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 MiB | [babelfish-17_5.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~noble_amd64.deb) |
| `babelfish-17` | `5.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 MiB | [babelfish-17_5.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~noble_arm64.deb) |
| `babelfish-17` | `5.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.2 MiB | [babelfish-17_5.4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~resolute_amd64.deb) |
| `babelfish-17` | `5.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.8 MiB | [babelfish-17_5.4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/babelfish-17/babelfish-17_5.4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="babelfish-17-17.7-5.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg babelfish;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install babelfish;		# install via package name, for the active PG version
pig install babelfishpg_common;		# install by extension name, for the current active PG version

pig install babelfishpg_common -v 18;   # install for PG 18
pig install babelfishpg_common -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_common;
```




## Usage

> [babelfishpg_common: SQL Server Transact SQL Datatype Support](https://babelfishpg.org/)

The `babelfishpg_common` extension provides SQL Server-compatible data type support for PostgreSQL as part of the Babelfish project. It enables PostgreSQL to understand and work with Microsoft SQL Server data types.

### Enabling

```sql
CREATE EXTENSION babelfishpg_common;
```

### SQL Server Data Types

The extension adds the following SQL Server-compatible data types:

- **TINYINT** - 1-byte unsigned integer (0 to 255)
- **SMALLMONEY** - Small monetary value
- **MONEY** - Monetary value (see also `babelfishpg_money`)
- **DATETIME** - SQL Server-style datetime
- **DATETIME2** - Extended precision datetime
- **SMALLDATETIME** - Reduced precision datetime
- **DATETIMEOFFSET** - Date and time with timezone offset
- **BIT** - SQL Server-compatible boolean
- **NCHAR** / **NVARCHAR** - Unicode character types
- **UNIQUEIDENTIFIER** - SQL Server-style UUID
- **VARBINARY** - Variable-length binary data
- **IMAGE** - Legacy binary data type
- **SQL_VARIANT** - Generic data type container
- **XML** - SQL Server-compatible XML type
- **SYSNAME** - System name type (nvarchar(128))

### Key Features

- Provides implicit and explicit type casting between SQL Server and PostgreSQL types
- Supports SQL Server-style collation behavior
- Handles SQL Server-specific type coercion rules
- Works in conjunction with `babelfishpg_tsql` for full T-SQL compatibility

This extension is typically deployed as part of a full Babelfish for PostgreSQL installation and is a prerequisite for `babelfishpg_tsql`.
