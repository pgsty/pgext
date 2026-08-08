---
title: "cat_tools"
linkTitle: "cat_tools"
description: "Tools for interfacing with the PostgreSQL catalog"
weight: 5290
categories: ["ADMIN"]
width: full
---

[**cat_tools**](https://github.com/Postgres-Extensions/cat_tools) : Tools for interfacing with the PostgreSQL catalog


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5290** | {{< badge content="cat_tools" link="https://github.com/Postgres-Extensions/cat_tools" >}} | {{< ext "cat_tools" >}} | `0.3.0` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `cat_tools` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "extension_drop" >}} {{< ext "object_reference" >}} |
|   **See Also**    | {{< ext "pg_catalog_get_defs" >}} {{< ext "pg_global_catalog" >}} {{< ext "meta_triggers" >}} {{< ext "pg_catcheck" >}} {{< ext "pgdd" >}} {{< ext "ddlx" >}} {{< ext "meta" >}} {{< ext "object_reference" >}} |

> [!Note] Promoted from a source-only universe row to PIGSTY RPM and DEB packages at 0.3.0; control fixes schema cat_tools and META declares the plpgsql runtime dependency.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `cat_tools` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "cat_tools_18" "green" >}} {{< bg "17" "cat_tools_17" "green" >}} {{< bg "16" "cat_tools_16" "green" >}} {{< bg "15" "cat_tools_15" "green" >}} {{< bg "14" "cat_tools_14" "green" >}} | `cat_tools_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-cat-tools" "green" >}} {{< bg "17" "postgresql-17-cat-tools" "green" >}} {{< bg "16" "postgresql-16-cat-tools" "green" >}} {{< bg "15" "postgresql-15-cat-tools" "green" >}} {{< bg "14" "postgresql-14-cat-tools" "green" >}} | `postgresql-$v-cat-tools` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Postgres-Extensions/cat_tools" title="Repository" icon="github" subtitle="github.com/Postgres-Extensions/cat_tools" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="cat_tools-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg cat_tools;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install cat_tools;		# install via package name, for the active PG version

pig install cat_tools -v 18;   # install for PG 18
pig install cat_tools -v 17;   # install for PG 17
pig install cat_tools -v 16;   # install for PG 16
pig install cat_tools -v 15;   # install for PG 15
pig install cat_tools -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION cat_tools CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [Official extension control file (cat_tools.control)](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/cat_tools.control)
- [Official extension SQL (cat_tools--0.1.0--0.1.3.sql)](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/sql/cat_tools--0.1.0--0.1.3.sql)

`cat_tools` — Tools for interfacing with the catalog. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION cat_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `__cat_tools.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `__cat_tools.exec(sql text)` is an extension function and returns `void`.
- `pg_temp.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `pg_temp.exec(sql text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
