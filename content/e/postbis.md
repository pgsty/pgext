---
title: "postbis"
linkTitle: "postbis"
description: "Adds compressed DNA, RNA, amino-acid, and aligned sequence types with casts, operators, indexes, and bioinformatics functions."
weight: 3760
categories: ["TYPE"]
width: full
---

[**postbis**](https://github.com/no0p/postbis) : Adds compressed DNA, RNA, amino-acid, and aligned sequence types with casts, operators, indexes, and bioinformatics functions.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3760** | {{< badge content="postbis" link="https://github.com/no0p/postbis" >}} | {{< ext "postbis" >}} | `1.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "rdkit" >}} {{< ext "vector" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_trgm" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "imgsmlr" >}} |

> [!Note] The packaged repository is an untagged copy of PostBIS, inactive since 2019; Pigsty pins commit ce454ebf and patches PostgreSQL 14-18 compatibility plus alphabet output and indexed slice correctness.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postbis` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postbis_18" "green" >}} {{< bg "17" "postbis_17" "green" >}} {{< bg "16" "postbis_16" "green" >}} {{< bg "15" "postbis_15" "green" >}} {{< bg "14" "postbis_14" "green" >}} | `postbis_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-postbis" "green" >}} {{< bg "17" "postgresql-17-postbis" "green" >}} {{< bg "16" "postgresql-16-postbis" "green" >}} {{< bg "15" "postgresql-15-postbis" "green" >}} {{< bg "14" "postgresql-14-postbis" "green" >}} | `postgresql-$v-postbis` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "postbis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postbis_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-postbis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-postbis : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postbis_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [postbis_18-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postbis_18-1.0-2PIGSTY.el8.x86_64.rpm) |
| `postbis_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.7 KiB | [postbis_18-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postbis_18-1.0-2PIGSTY.el8.aarch64.rpm) |
| `postbis_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.8 KiB | [postbis_18-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postbis_18-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postbis_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.8 KiB | [postbis_18-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postbis_18-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postbis_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.6 KiB | [postbis_18-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postbis_18-1.0-2PIGSTY.el10.x86_64.rpm) |
| `postbis_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.5 KiB | [postbis_18-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postbis_18-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-postbis` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 152.2 KiB | [postgresql-18-postbis_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-postbis` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 147.1 KiB | [postgresql-18-postbis_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-postbis` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 153.2 KiB | [postgresql-18-postbis_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-postbis` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 147.8 KiB | [postgresql-18-postbis_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-postbis` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 162.7 KiB | [postgresql-18-postbis_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-postbis` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 160.9 KiB | [postgresql-18-postbis_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-postbis` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 160.6 KiB | [postgresql-18-postbis_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-postbis` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 157.8 KiB | [postgresql-18-postbis_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-postbis` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 160.4 KiB | [postgresql-18-postbis_1.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-postbis` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 157.3 KiB | [postgresql-18-postbis_1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-18-postbis_1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postbis_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [postbis_17-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postbis_17-1.0-2PIGSTY.el8.x86_64.rpm) |
| `postbis_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.7 KiB | [postbis_17-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postbis_17-1.0-2PIGSTY.el8.aarch64.rpm) |
| `postbis_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.8 KiB | [postbis_17-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postbis_17-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postbis_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [postbis_17-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postbis_17-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postbis_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.6 KiB | [postbis_17-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postbis_17-1.0-2PIGSTY.el10.x86_64.rpm) |
| `postbis_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.5 KiB | [postbis_17-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postbis_17-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-postbis` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 152.1 KiB | [postgresql-17-postbis_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-postbis` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 147.2 KiB | [postgresql-17-postbis_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-postbis` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 153.1 KiB | [postgresql-17-postbis_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-postbis` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 147.9 KiB | [postgresql-17-postbis_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-postbis` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.9 KiB | [postgresql-17-postbis_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-postbis` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 167.0 KiB | [postgresql-17-postbis_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-postbis` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 160.6 KiB | [postgresql-17-postbis_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-postbis` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 157.8 KiB | [postgresql-17-postbis_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-postbis` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 160.4 KiB | [postgresql-17-postbis_1.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-postbis` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 157.2 KiB | [postgresql-17-postbis_1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-17-postbis_1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postbis_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [postbis_16-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postbis_16-1.0-2PIGSTY.el8.x86_64.rpm) |
| `postbis_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.7 KiB | [postbis_16-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postbis_16-1.0-2PIGSTY.el8.aarch64.rpm) |
| `postbis_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.8 KiB | [postbis_16-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postbis_16-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postbis_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [postbis_16-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postbis_16-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postbis_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.6 KiB | [postbis_16-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postbis_16-1.0-2PIGSTY.el10.x86_64.rpm) |
| `postbis_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.5 KiB | [postbis_16-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postbis_16-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-postbis` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 152.1 KiB | [postgresql-16-postbis_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-postbis` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 147.2 KiB | [postgresql-16-postbis_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-postbis` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 153.2 KiB | [postgresql-16-postbis_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-postbis` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 148.0 KiB | [postgresql-16-postbis_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-postbis` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 169.0 KiB | [postgresql-16-postbis_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-postbis` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.9 KiB | [postgresql-16-postbis_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-postbis` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 160.6 KiB | [postgresql-16-postbis_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-postbis` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 157.8 KiB | [postgresql-16-postbis_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-postbis` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 160.3 KiB | [postgresql-16-postbis_1.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-postbis` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 157.2 KiB | [postgresql-16-postbis_1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-16-postbis_1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postbis_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 66.4 KiB | [postbis_15-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postbis_15-1.0-2PIGSTY.el8.x86_64.rpm) |
| `postbis_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [postbis_15-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postbis_15-1.0-2PIGSTY.el8.aarch64.rpm) |
| `postbis_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.3 KiB | [postbis_15-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postbis_15-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postbis_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.3 KiB | [postbis_15-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postbis_15-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postbis_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.9 KiB | [postbis_15-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postbis_15-1.0-2PIGSTY.el10.x86_64.rpm) |
| `postbis_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.5 KiB | [postbis_15-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postbis_15-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-postbis` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 153.6 KiB | [postgresql-15-postbis_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-postbis` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 148.4 KiB | [postgresql-15-postbis_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-postbis` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 154.6 KiB | [postgresql-15-postbis_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-postbis` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.1 KiB | [postgresql-15-postbis_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-postbis` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 170.6 KiB | [postgresql-15-postbis_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-postbis` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 168.0 KiB | [postgresql-15-postbis_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-postbis` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 161.3 KiB | [postgresql-15-postbis_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-postbis` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 158.6 KiB | [postgresql-15-postbis_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-postbis` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 161.7 KiB | [postgresql-15-postbis_1.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-postbis` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 158.2 KiB | [postgresql-15-postbis_1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-15-postbis_1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postbis_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 66.4 KiB | [postbis_14-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postbis_14-1.0-2PIGSTY.el8.x86_64.rpm) |
| `postbis_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.5 KiB | [postbis_14-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postbis_14-1.0-2PIGSTY.el8.aarch64.rpm) |
| `postbis_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.7 KiB | [postbis_14-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postbis_14-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postbis_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.4 KiB | [postbis_14-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postbis_14-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postbis_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.5 KiB | [postbis_14-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postbis_14-1.0-2PIGSTY.el10.x86_64.rpm) |
| `postbis_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.7 KiB | [postbis_14-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postbis_14-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-postbis` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 153.7 KiB | [postgresql-14-postbis_1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-postbis` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 148.3 KiB | [postgresql-14-postbis_1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-postbis` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 154.5 KiB | [postgresql-14-postbis_1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-postbis` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.2 KiB | [postgresql-14-postbis_1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-postbis` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 170.5 KiB | [postgresql-14-postbis_1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-postbis` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 168.0 KiB | [postgresql-14-postbis_1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-postbis` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 161.3 KiB | [postgresql-14-postbis_1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-postbis` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 158.6 KiB | [postgresql-14-postbis_1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-postbis` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 161.7 KiB | [postgresql-14-postbis_1.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-postbis` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 158.2 KiB | [postgresql-14-postbis_1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postbis/postgresql-14-postbis_1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/no0p/postbis" title="Repository" icon="github" subtitle="github.com/no0p/postbis" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postbis-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg postbis;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install postbis;		# install via package name, for the active PG version

pig install postbis -v 18;   # install for PG 18
pig install postbis -v 17;   # install for PG 17
pig install postbis -v 16;   # install for PG 16
pig install postbis -v 15;   # install for PG 15
pig install postbis -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postbis;
```

## Usage

Sources:

- [Project README](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/README.txt)
- [Extension control file](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/postbis.control)
- [Version 1.0 SQL API](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/sql/postbis--1.0.sql)
- [Sequence regression tests](https://github.com/no0p/postbis/tree/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/test/sql)

`postbis` 1.0 provides compact native types for DNA, RNA, amino-acid, and aligned sequences. It also provides configurable alphabets and type modifiers, casts, sequence operations, biological transformations, comparison operators, and B-tree and hash operator classes.

### Store typed sequences

```sql
CREATE EXTENSION postbis;

CREATE TABLE specimen (
  specimen_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  dna dna_sequence(SHORT, FLC, CASE_SENSITIVE) NOT NULL,
  rna rna_sequence(IUPAC, CASE_SENSITIVE),
  protein aa_sequence(IUPAC, CASE_SENSITIVE)
);

INSERT INTO specimen (dna, rna, protein)
VALUES ('AACCGGTT', 'AACGUU', 'ACDEFG');

SELECT specimen_id,
       char_length(dna) AS bases,
       substr(dna, 3, 4)::text AS fragment
FROM specimen;
```

Input validation depends on the selected alphabet, case-sensitivity, and type modifiers. Verify that casts reject symbols outside the required biological convention and that aligned and unaligned types are not mixed accidentally.

### Transform and translate sequences

```sql
SELECT complement('ACGTN'::dna_sequence)::text;
-- TGCAN

SELECT reverse_complement('ACGTN'::dna_sequence)::text;
-- NACGT

SELECT transcribe('AACGTT'::dna_sequence)::text;
-- AACGUU

SELECT translate('AUGGCCUAA'::rna_sequence)::text;
-- MA
```

The extension also exposes `reverse_transcribe()`, `six_frame()`, `get_alphabet()`, `entropy()`, `gc_content()`, and sequence generators. The translation functions accept explicit translation tables when the standard genetic code is not appropriate.

### Inspect compression and add indexes

```sql
SELECT char_length(sequence) AS symbols,
       octet_length(sequence) AS storage_bytes,
       compression_ratio(sequence) AS storage_ratio
FROM (
  SELECT repeat('ACGT', 256)::dna_sequence AS sequence
) AS sample;

CREATE INDEX specimen_dna_btree ON specimen USING btree (dna);
CREATE INDEX specimen_dna_hash  ON specimen USING hash  (dna);
```

Equality, ordering, concatenation, substring, search, and length functions are available for the sequence types. Check plans and realistic data distributions before relying on an index for a production workload.

### Packaging and durability risk

Pigsty applies a downstream compatibility patch and packages PostBIS 1.0 for PostgreSQL 14–18. That packaging result does not change the upstream lifecycle: the project is inactive and has no extension upgrade path beyond 1.0.

The custom types use native compressed on-disk representations. Treat stored values and indexes as tied to an exact tested build. Before adoption or migration, prove dump and restore, binary and logical upgrades, replication, driver decoding, index rebuilds, malformed input handling, and large-sequence memory behavior.

Functions such as `reverse()`, `char_length()`, and `substr()` overload familiar names, so schema qualification and controlled `search_path` settings matter. For new durable datasets, prefer maintained sequence tooling or plain PostgreSQL types unless the extension has been locally audited, packaged, and assigned an explicit long-term migration owner.
