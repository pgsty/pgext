---
title: "plv8"
linkTitle: "plv8"
description: "PL/JavaScript (v8) trusted procedural language"
weight: 3010
categories: ["LANG"]
languages: ["C++"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**plv8**](https://github.com/plv8/plv8) : PL/JavaScript (v8) trusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3010** | {{< badge content="plv8" link="https://github.com/plv8/plv8" >}} | {{< ext "plv8" >}} | `3.2.5` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "pljs" >}} {{< ext "pllua" >}} {{< ext "pgwasm" >}} {{< ext "pg_tle" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plv8` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.5` | {{< bg "18" "plv8_18" "green" >}} {{< bg "17" "plv8_17" "green" >}} {{< bg "16" "plv8_16" "green" >}} {{< bg "15" "plv8_15" "green" >}} {{< bg "14" "plv8_14" "green" >}} | `plv8_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.5` | {{< bg "18" "postgresql-18-plv8" "green" >}} {{< bg "17" "postgresql-17-plv8" "green" >}} {{< bg "16" "postgresql-16-plv8" "green" >}} {{< bg "15" "postgresql-15-plv8" "green" >}} {{< bg "14" "postgresql-14-plv8" "green" >}} | `postgresql-$v-plv8` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "plv8_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-18-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.5" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_18` | `3.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [plv8_18-3.2.5-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_18-3.2.5-1PGSTY.el8.x86_64.rpm) |
| `plv8_18` | `3.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [plv8_18-3.2.5-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_18-3.2.5-1PGSTY.el8.aarch64.rpm) |
| `plv8_18` | `3.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.5 MiB | [plv8_18-3.2.5-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_18-3.2.5-1PGSTY.el9.x86_64.rpm) |
| `plv8_18` | `3.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.4 MiB | [plv8_18-3.2.5-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_18-3.2.5-1PGSTY.el9.aarch64.rpm) |
| `plv8_18` | `3.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.8 MiB | [plv8_18-3.2.5-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_18-3.2.5-1PGSTY.el10.x86_64.rpm) |
| `plv8_18` | `3.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.6 MiB | [plv8_18-3.2.5-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_18-3.2.5-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-plv8` | `3.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.7 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.2 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.8 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.5 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.2 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.5 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-plv8` | `3.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.5 MiB | [postgresql-18-plv8_3.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-18-plv8_3.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_17` | `3.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [plv8_17-3.2.5-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_17-3.2.5-1PGSTY.el8.x86_64.rpm) |
| `plv8_17` | `3.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [plv8_17-3.2.5-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_17-3.2.5-1PGSTY.el8.aarch64.rpm) |
| `plv8_17` | `3.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_17-3.2.5-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_17-3.2.5-1PGSTY.el9.x86_64.rpm) |
| `plv8_17` | `3.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.4 MiB | [plv8_17-3.2.5-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_17-3.2.5-1PGSTY.el9.aarch64.rpm) |
| `plv8_17` | `3.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.9 MiB | [plv8_17-3.2.5-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_17-3.2.5-1PGSTY.el10.x86_64.rpm) |
| `plv8_17` | `3.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.6 MiB | [plv8_17-3.2.5-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_17-3.2.5-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-plv8` | `3.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.7 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.2 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.8 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.2 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-plv8` | `3.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.5 MiB | [postgresql-17-plv8_3.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-17-plv8_3.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_16` | `3.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [plv8_16-3.2.5-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_16-3.2.5-1PGSTY.el8.x86_64.rpm) |
| `plv8_16` | `3.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [plv8_16-3.2.5-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_16-3.2.5-1PGSTY.el8.aarch64.rpm) |
| `plv8_16` | `3.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_16-3.2.5-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_16-3.2.5-1PGSTY.el9.x86_64.rpm) |
| `plv8_16` | `3.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.4 MiB | [plv8_16-3.2.5-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_16-3.2.5-1PGSTY.el9.aarch64.rpm) |
| `plv8_16` | `3.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.9 MiB | [plv8_16-3.2.5-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_16-3.2.5-1PGSTY.el10.x86_64.rpm) |
| `plv8_16` | `3.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.6 MiB | [plv8_16-3.2.5-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_16-3.2.5-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-plv8` | `3.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.2 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.8 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.2 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-plv8` | `3.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.5 MiB | [postgresql-16-plv8_3.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-16-plv8_3.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_15` | `3.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [plv8_15-3.2.5-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_15-3.2.5-1PGSTY.el8.x86_64.rpm) |
| `plv8_15` | `3.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [plv8_15-3.2.5-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_15-3.2.5-1PGSTY.el8.aarch64.rpm) |
| `plv8_15` | `3.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.6 MiB | [plv8_15-3.2.5-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_15-3.2.5-1PGSTY.el9.x86_64.rpm) |
| `plv8_15` | `3.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_15-3.2.5-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_15-3.2.5-1PGSTY.el9.aarch64.rpm) |
| `plv8_15` | `3.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.9 MiB | [plv8_15-3.2.5-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_15-3.2.5-1PGSTY.el10.x86_64.rpm) |
| `plv8_15` | `3.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.6 MiB | [plv8_15-3.2.5-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_15-3.2.5-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-plv8` | `3.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.2 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.8 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.2 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.1 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-plv8` | `3.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 MiB | [postgresql-15-plv8_3.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-15-plv8_3.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_14` | `3.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [plv8_14-3.2.5-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_14-3.2.5-1PGSTY.el8.x86_64.rpm) |
| `plv8_14` | `3.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [plv8_14-3.2.5-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_14-3.2.5-1PGSTY.el8.aarch64.rpm) |
| `plv8_14` | `3.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.5 MiB | [plv8_14-3.2.5-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_14-3.2.5-1PGSTY.el9.x86_64.rpm) |
| `plv8_14` | `3.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.3 MiB | [plv8_14-3.2.5-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_14-3.2.5-1PGSTY.el9.aarch64.rpm) |
| `plv8_14` | `3.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.9 MiB | [plv8_14-3.2.5-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plv8_14-3.2.5-1PGSTY.el10.x86_64.rpm) |
| `plv8_14` | `3.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.5 MiB | [plv8_14-3.2.5-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plv8_14-3.2.5-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-plv8` | `3.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.7 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.2 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.8 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.2 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.1 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-plv8` | `3.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 MiB | [postgresql-14-plv8_3.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plv8/postgresql-14-plv8_3.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plv8/plv8" title="Repository" icon="github" subtitle="github.com/plv8/plv8" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plv8-3.2.5.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg plv8;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [README](https://github.com/plv8/plv8/blob/v3.2.5/README.md)
- [Built-ins](https://github.com/plv8/plv8/blob/v3.2.5/docs/BUILTINS.md)
- [Configuration](https://github.com/plv8/plv8/blob/v3.2.5/docs/CONFIGURATION.md)
- [Changes](https://github.com/plv8/plv8/blob/v3.2.5/Changes)
- [Control](https://github.com/plv8/plv8/blob/v3.2.5/plv8.control.common)
- [SQL template](https://github.com/plv8/plv8/blob/v3.2.5/plv8.sql.common)

`plv8` provides a trusted JavaScript procedural language powered by V8. This page follows upstream 3.2.5, including its effective-user crash fixes and PostgreSQL 19 support.

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

- The 3.2.5 CI matrix covers PostgreSQL 14–19. PostgreSQL-major support is separate from the versions available in downstream packages.
- Creating the extension requires a superuser; the installed JavaScript language is trusted, which does not make the extension itself installable by an unprivileged user. `plv8_info()` has its public execution privilege revoked by the installation SQL.
- Version 3.2.5 fixes crashes when a cached function runs under a different effective user via `SECURITY DEFINER` or `SET ROLE`, or after `plv8_reset()`, and when an exception message cannot be converted to a string.
- Each session has its own global JavaScript runtime; switching roles initializes a separate runtime context.
- `plv8.execution_timeout` only applies when the extension is compiled with execution-timeout support.
