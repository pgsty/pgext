---
title: "external_file"
linkTitle: "external_file"
description: "Access external server-side files through PostgreSQL functions"
weight: 4285
categories: ["UTIL"]
width: full
---

[**external_file**](https://github.com/darold/external_file) : Access external server-side files through PostgreSQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4285** | {{< badge content="external_file" link="https://github.com/darold/external_file" >}} | {{< ext "external_file" >}} | `1.2` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `external_file` |

> [!Note] Fixed schema external_file; superuser required.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `external_file` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "external_file_18" "green" >}} {{< bg "17" "external_file_17" "green" >}} {{< bg "16" "external_file_16" "green" >}} {{< bg "15" "external_file_15" "green" >}} {{< bg "14" "external_file_14" "green" >}} | `external_file_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "postgresql-18-external-file" "green" >}} {{< bg "17" "postgresql-17-external-file" "green" >}} {{< bg "16" "postgresql-16-external-file" "green" >}} {{< bg "15" "postgresql-15-external-file" "green" >}} {{< bg "14" "postgresql-14-external-file" "green" >}} | `postgresql-$v-external-file` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2" "external_file_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "external_file_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-external-file : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-external-file : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `external_file_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [external_file_18-1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/external_file_18-1.2-1PIGSTY.el8.x86_64.rpm) |
| `external_file_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.9 KiB | [external_file_18-1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/external_file_18-1.2-1PIGSTY.el8.aarch64.rpm) |
| `external_file_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.9 KiB | [external_file_18-1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/external_file_18-1.2-1PIGSTY.el9.x86_64.rpm) |
| `external_file_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [external_file_18-1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/external_file_18-1.2-1PIGSTY.el9.aarch64.rpm) |
| `external_file_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [external_file_18-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/external_file_18-1.2-1PIGSTY.el10.x86_64.rpm) |
| `external_file_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [external_file_18-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/external_file_18-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-external-file` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-external-file` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-external-file` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-external-file` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-external-file` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-external-file` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-external-file` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-external-file` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 KiB | [postgresql-18-external-file_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-18-external-file_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `external_file_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [external_file_17-1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/external_file_17-1.2-1PIGSTY.el8.x86_64.rpm) |
| `external_file_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.9 KiB | [external_file_17-1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/external_file_17-1.2-1PIGSTY.el8.aarch64.rpm) |
| `external_file_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.9 KiB | [external_file_17-1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/external_file_17-1.2-1PIGSTY.el9.x86_64.rpm) |
| `external_file_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [external_file_17-1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/external_file_17-1.2-1PIGSTY.el9.aarch64.rpm) |
| `external_file_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [external_file_17-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/external_file_17-1.2-1PIGSTY.el10.x86_64.rpm) |
| `external_file_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [external_file_17-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/external_file_17-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-external-file` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-external-file` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-external-file` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-external-file` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-external-file` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-external-file` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.1 KiB | [postgresql-17-external-file_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-external-file` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.2 KiB | [postgresql-17-external-file_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-external-file` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.2 KiB | [postgresql-17-external-file_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-17-external-file_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `external_file_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [external_file_16-1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/external_file_16-1.2-1PIGSTY.el8.x86_64.rpm) |
| `external_file_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.9 KiB | [external_file_16-1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/external_file_16-1.2-1PIGSTY.el8.aarch64.rpm) |
| `external_file_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.9 KiB | [external_file_16-1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/external_file_16-1.2-1PIGSTY.el9.x86_64.rpm) |
| `external_file_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [external_file_16-1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/external_file_16-1.2-1PIGSTY.el9.aarch64.rpm) |
| `external_file_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [external_file_16-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/external_file_16-1.2-1PIGSTY.el10.x86_64.rpm) |
| `external_file_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [external_file_16-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/external_file_16-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-external-file` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-external-file` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-external-file` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-external-file` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-external-file` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-external-file` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-external-file` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-external-file` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 KiB | [postgresql-16-external-file_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-16-external-file_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `external_file_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [external_file_15-1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/external_file_15-1.2-1PIGSTY.el8.x86_64.rpm) |
| `external_file_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.9 KiB | [external_file_15-1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/external_file_15-1.2-1PIGSTY.el8.aarch64.rpm) |
| `external_file_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.9 KiB | [external_file_15-1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/external_file_15-1.2-1PIGSTY.el9.x86_64.rpm) |
| `external_file_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [external_file_15-1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/external_file_15-1.2-1PIGSTY.el9.aarch64.rpm) |
| `external_file_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [external_file_15-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/external_file_15-1.2-1PIGSTY.el10.x86_64.rpm) |
| `external_file_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [external_file_15-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/external_file_15-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-external-file` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-external-file` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-external-file` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-external-file` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-external-file` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-external-file` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-external-file` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-external-file` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 KiB | [postgresql-15-external-file_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-15-external-file_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `external_file_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [external_file_14-1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/external_file_14-1.2-1PIGSTY.el8.x86_64.rpm) |
| `external_file_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.9 KiB | [external_file_14-1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/external_file_14-1.2-1PIGSTY.el8.aarch64.rpm) |
| `external_file_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.9 KiB | [external_file_14-1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/external_file_14-1.2-1PIGSTY.el9.x86_64.rpm) |
| `external_file_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [external_file_14-1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/external_file_14-1.2-1PIGSTY.el9.aarch64.rpm) |
| `external_file_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [external_file_14-1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/external_file_14-1.2-1PIGSTY.el10.x86_64.rpm) |
| `external_file_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [external_file_14-1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/external_file_14-1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-external-file` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-external-file` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-external-file` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-external-file` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-external-file` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-external-file` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-external-file` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-external-file` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 KiB | [postgresql-14-external-file_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/external-file/postgresql-14-external-file_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/darold/external_file" title="Repository" icon="github" subtitle="github.com/darold/external_file" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="external_file-1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg external_file;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install external_file;		# install via package name, for the active PG version

pig install external_file -v 18;   # install for PG 18
pig install external_file -v 17;   # install for PG 17
pig install external_file -v 16;   # install for PG 16
pig install external_file -v 15;   # install for PG 15
pig install external_file -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION external_file;
```

## Usage

- Source: [GitHub repo](https://github.com/darold/external_file), [README](https://github.com/darold/external_file/blob/master/README.md)
- `external_file` provides server-side access to external files through a PostgreSQL extension, similar to Oracle BFILE-style locators.

```sql
CREATE EXTENSION external_file;
```

The README states that objects are created in the `external_file` schema by default and that creating the extension requires a PostgreSQL superuser.

## Core Workflow

The extension uses a directory alias plus a file name to identify an external file. The upstream README shows this sequence:

```sql
INSERT INTO directories(directory_name, directory_path)
VALUES ('temporary', '/tmp/');

INSERT INTO directory_roles(directory_name, directory_role, directory_read, directory_write)
VALUES ('temporary', 'a_role', true, false);

SELECT writeEfile('\x48656c6c6f2c0a0a596f75206172652072656164696e67206120746578742066696c652e0a0a526567617264732c0a',
                  ('temporary', 'blahblah.txt'));
SELECT readefile(the_file) FROM efile_test;
SELECT copyefile(('temporary', 'blahblah.txt'), ('temporary', 'copy_blahblah.txt'));
```

The main exported helpers are `efilename`, `readEfile`, `writeEfile`, `copyEfile`, and `getEfilePath`.

## Notes

The extension does not read files directly from the server filesystem. It uses the server-side `lo_*` family and enforces access through directory and role tables.
