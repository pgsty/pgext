---
title: "plv8"
linkTitle: "plv8"
description: "PL/JavaScript (v8) trusted procedural language"
weight: 3010
categories: ["LANG"]
width: full
---

[**plv8**](https://github.com/plv8/plv8) : PL/JavaScript (v8) trusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3010** | {{< badge content="plv8" link="https://github.com/plv8/plv8" >}} | {{< ext "plv8" >}} | `3.2.4` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plv8` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.4` | {{< bg "18" "plv8_18" "green" >}} {{< bg "17" "plv8_17" "green" >}} {{< bg "16" "plv8_16" "green" >}} {{< bg "15" "plv8_15" "green" >}} {{< bg "14" "plv8_14" "green" >}} | `plv8_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.4` | {{< bg "18" "postgresql-18-plv8" "green" >}} {{< bg "17" "postgresql-17-plv8" "green" >}} {{< bg "16" "postgresql-16-plv8" "green" >}} {{< bg "15" "postgresql-15-plv8" "green" >}} {{< bg "14" "postgresql-14-plv8" "green" >}} | `postgresql-$v-plv8` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_18` | `3.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.5 MiB | [plv8_18-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_18-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_18` | `3.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [plv8_18-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_18-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_18` | `3.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_18-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_18-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_18` | `3.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_18-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_18-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `plv8_18` | `3.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.8 MiB | [plv8_18-3.2.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_18-3.2.4-2PIGSTY.el10.x86_64.rpm) |
| `plv8_18` | `3.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [plv8_18-3.2.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_18-3.2.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-plv8` | `3.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.8 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plv8` | `3.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.6 MiB | [postgresql-18-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-18-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_17` | `3.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.5 MiB | [plv8_17-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_17-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_17` | `3.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [plv8_17-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_17-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_17` | `3.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_17-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_17-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_17` | `3.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_17-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_17-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `plv8_17` | `3.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.8 MiB | [plv8_17-3.2.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_17-3.2.4-2PIGSTY.el10.x86_64.rpm) |
| `plv8_17` | `3.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [plv8_17-3.2.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_17-3.2.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-plv8` | `3.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.8 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plv8` | `3.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_16` | `3.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.5 MiB | [plv8_16-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_16-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_16` | `3.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [plv8_16-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_16-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_16` | `3.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_16-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_16-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_16` | `3.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_16-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_16-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `plv8_16` | `3.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.8 MiB | [plv8_16-3.2.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_16-3.2.4-2PIGSTY.el10.x86_64.rpm) |
| `plv8_16` | `3.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [plv8_16-3.2.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_16-3.2.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-plv8` | `3.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.8 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plv8` | `3.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.6 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_15` | `3.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.5 MiB | [plv8_15-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_15-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_15` | `3.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [plv8_15-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_15-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_15` | `3.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_15-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_15-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_15` | `3.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_15-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_15-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `plv8_15` | `3.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.2 MiB | [plv8_15-3.2.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_15-3.2.4-2PIGSTY.el10.x86_64.rpm) |
| `plv8_15` | `3.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [plv8_15-3.2.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_15-3.2.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-plv8` | `3.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.8 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plv8` | `3.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.6 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_14` | `3.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.5 MiB | [plv8_14-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_14-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_14` | `3.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [plv8_14-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_14-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_14` | `3.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_14-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_14-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_14` | `3.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_14-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_14-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `plv8_14` | `3.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.2 MiB | [plv8_14-3.2.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_14-3.2.4-2PIGSTY.el10.x86_64.rpm) |
| `plv8_14` | `3.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [plv8_14-3.2.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_14-3.2.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-plv8` | `3.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.8 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plv8` | `3.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plv8/plv8" title="Repository" icon="github" subtitle="github.com/plv8/plv8" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plv8-3.2.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plv8;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plv8;		# install via package name, for the active PG version

pig install plv8 -v 18;   # install for PG 18
pig install plv8 -v 17;   # install for PG 17
pig install plv8 -v 16;   # install for PG 16
pig install plv8 -v 15;   # install for PG 15
pig install plv8 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plv8;
```

## Usage

Source: [README](https://github.com/plv8/plv8/blob/r3.2/README.md), [Docs site](https://plv8.github.io/), [Built-ins](https://github.com/plv8/plv8/blob/r3.2/docs/BUILTINS.md), [Runtime configuration](https://github.com/plv8/plv8/blob/r3.2/docs/CONFIGURATION.md), [Tag v3.2.4](https://github.com/plv8/plv8/tree/v3.2.4)

`plv8` is a trusted JavaScript procedural language for PostgreSQL, powered by the V8 engine. Upstream currently tags the extension as `v3.2.4`; Pigsty's `3.2.4-2` package version is a packaging revision rather than a new upstream extension release.

### Basic use

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;

CREATE FUNCTION plv8_test(keys text[], vals text[]) RETURNS json AS $$
  let out = {};
  for (let i = 0; i < keys.length; i++) out[keys[i]] = vals[i];
  return out;
$$ LANGUAGE plv8 IMMUTABLE STRICT;
```

### Common built-ins

- `plv8.elog(level, ...)`: emit PostgreSQL log or client messages.
- `plv8.execute(sql [, args])`: run SQL and return rows or affected-row count.
- `plv8.prepare(...)`, `PreparedPlan.execute()`, `PreparedPlan.cursor()`: prepared SPI access.
- `plv8.subtransaction(fn)`: run a group of SPI operations atomically.
- `plv8.find_function(...)`: call another PLV8 function by name.
- `plv8.memory_usage()`: inspect V8 heap usage for the current session.
- `plv8.run_script(source, name)`: evaluate named script text.

### Runtime settings

```sql
SET plv8.start_proc = 'plv8_init';
SET plv8.execution_timeout = 60;
SET plv8.memory_limit = 512;
```

- `plv8.start_proc`
- `plv8.v8_flags`
- `plv8.execution_timeout`
- `plv8.memory_limit`

### Caveats

- Current docs state support for PostgreSQL 13 and above.
- Each session has its own global JavaScript runtime; switching roles initializes a separate runtime context.
- `plv8.execution_timeout` only applies when the extension is compiled with execution-timeout support.
