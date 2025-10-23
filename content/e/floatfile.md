---
title: "floatfile"
linkTitle: "floatfile"
description: "Simple file storage for arrays of floats"
weight: 4280
categories: ["UTIL"]
width: full
---

Simple file storage for arrays of floats


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4280** | {{< badge content="floatfile" link="https://github.com/pjungwir/floatfile" >}} | {{< ext "floatfile" >}} | `1.3.1` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_bulkload" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/floatfile" >}} | `1.3.1` | {{< bg "18" "floatfile_18*" "red" >}} {{< bg "17" "floatfile_17*" "green" >}} {{< bg "16" "floatfile_16*" "green" >}} {{< bg "15" "floatfile_15*" "green" >}} {{< bg "14" "floatfile_14*" "green" >}} | `floatfile_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/floatfile" >}} | `1.3.1` | {{< bg "18" "postgresql-18-floatfile" "red" >}} {{< bg "17" "postgresql-17-floatfile" "green" >}} {{< bg "16" "postgresql-16-floatfile" "green" >}} {{< bg "15" "postgresql-15-floatfile" "green" >}} {{< bg "14" "postgresql-14-floatfile" "green" >}} | `postgresql-$v-floatfile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "floatfile_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "floatfile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "floatfile_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "floatfile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "floatfile_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "floatfile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "floatfile_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "floatfile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "floatfile_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-floatfile : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.1" "postgresql-17-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-floatfile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-floatfile : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatfile_17` | 1.3.1 | `el8.x86_64` | pigsty | 26.3 KiB | [floatfile_17-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatfile_17-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `floatfile_17` | 1.3.1 | `el8.aarch64` | pigsty | 25.8 KiB | [floatfile_17-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatfile_17-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `floatfile_17` | 1.3.1 | `el9.x86_64` | pigsty | 27.5 KiB | [floatfile_17-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatfile_17-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `floatfile_17` | 1.3.1 | `el9.aarch64` | pigsty | 26.3 KiB | [floatfile_17-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatfile_17-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-floatfile` | 1.3.1 | `d12.x86_64` | pigsty | 45.7 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-floatfile` | 1.3.1 | `d12.aarch64` | pigsty | 44.6 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-floatfile` | 1.3.1 | `u22.x86_64` | pigsty | 48.0 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-floatfile` | 1.3.1 | `u22.aarch64` | pigsty | 47.5 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-floatfile` | 1.3.1 | `u24.x86_64` | pigsty | 46.8 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-floatfile` | 1.3.1 | `u24.aarch64` | pigsty | 45.8 KiB | [postgresql-17-floatfile_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-17-floatfile_1.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatfile_16` | 1.3.1 | `el8.x86_64` | pigsty | 26.3 KiB | [floatfile_16-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatfile_16-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `floatfile_16` | 1.3.1 | `el8.aarch64` | pigsty | 25.8 KiB | [floatfile_16-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatfile_16-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `floatfile_16` | 1.3.1 | `el9.x86_64` | pigsty | 27.5 KiB | [floatfile_16-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatfile_16-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `floatfile_16` | 1.3.1 | `el9.aarch64` | pigsty | 26.4 KiB | [floatfile_16-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatfile_16-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-floatfile` | 1.3.1 | `d12.x86_64` | pigsty | 45.7 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-floatfile` | 1.3.1 | `d12.aarch64` | pigsty | 44.5 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-floatfile` | 1.3.1 | `u22.x86_64` | pigsty | 48.0 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-floatfile` | 1.3.1 | `u22.aarch64` | pigsty | 47.5 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-floatfile` | 1.3.1 | `u24.x86_64` | pigsty | 46.8 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-floatfile` | 1.3.1 | `u24.aarch64` | pigsty | 45.8 KiB | [postgresql-16-floatfile_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-16-floatfile_1.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatfile_15` | 1.3.1 | `el8.x86_64` | pigsty | 26.5 KiB | [floatfile_15-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatfile_15-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `floatfile_15` | 1.3.1 | `el8.aarch64` | pigsty | 25.9 KiB | [floatfile_15-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatfile_15-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `floatfile_15` | 1.3.1 | `el9.x86_64` | pigsty | 27.6 KiB | [floatfile_15-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatfile_15-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `floatfile_15` | 1.3.1 | `el9.aarch64` | pigsty | 26.5 KiB | [floatfile_15-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatfile_15-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-floatfile` | 1.3.1 | `d12.x86_64` | pigsty | 45.8 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-floatfile` | 1.3.1 | `d12.aarch64` | pigsty | 44.7 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-floatfile` | 1.3.1 | `u22.x86_64` | pigsty | 48.1 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-floatfile` | 1.3.1 | `u22.aarch64` | pigsty | 47.7 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-floatfile` | 1.3.1 | `u24.x86_64` | pigsty | 46.6 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-floatfile` | 1.3.1 | `u24.aarch64` | pigsty | 45.8 KiB | [postgresql-15-floatfile_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-15-floatfile_1.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `floatfile_14` | 1.3.1 | `el8.x86_64` | pigsty | 26.4 KiB | [floatfile_14-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/floatfile_14-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `floatfile_14` | 1.3.1 | `el8.aarch64` | pigsty | 25.9 KiB | [floatfile_14-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/floatfile_14-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `floatfile_14` | 1.3.1 | `el9.x86_64` | pigsty | 27.6 KiB | [floatfile_14-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/floatfile_14-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `floatfile_14` | 1.3.1 | `el9.aarch64` | pigsty | 26.5 KiB | [floatfile_14-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/floatfile_14-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-floatfile` | 1.3.1 | `d12.x86_64` | pigsty | 45.7 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-floatfile` | 1.3.1 | `d12.aarch64` | pigsty | 44.6 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-floatfile` | 1.3.1 | `u22.x86_64` | pigsty | 48.1 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-floatfile` | 1.3.1 | `u22.aarch64` | pigsty | 47.6 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-floatfile` | 1.3.1 | `u24.x86_64` | pigsty | 46.5 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-floatfile` | 1.3.1 | `u24.aarch64` | pigsty | 45.7 KiB | [postgresql-14-floatfile_1.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/floatfile/postgresql-14-floatfile_1.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pjungwir/floatfile" title="Repository" icon="github" subtitle="github.com/pjungwir/floatfile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="floatfile-1.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get floatfile; # get floatfile source code
pig build dep floatfile; # install build dependencies
pig build pkg floatfile; # build extension rpm or deb
pig build ext floatfile; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install floatfile; # install by extension name, for the current active PG version
pig ext install floatfile; # install via package alias, for the active PG version
pig ext install floatfile -v 18;   # install for PG 18
pig ext install floatfile -v 17;   # install for PG 17
pig ext install floatfile -v 16;   # install for PG 16
pig ext install floatfile -v 15;   # install for PG 15
pig ext install floatfile -v 14;   # install for PG 14
pig ext install floatfile -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION floatfile;
```

